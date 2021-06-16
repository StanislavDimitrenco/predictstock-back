package payment

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"log"
	"net/url"
	"strconv"
	"strings"
)

const (
	queryOutSumm     = "OutSum"
	queryInvID       = "InvId"
	queryCRC         = "SignatureValue"
	queryDescription = "Desc"
	isTestMode       = "IsTest"
	queryLogin       = "MrchLogin"
	robokassaHost    = "auth.robokassa.ru"
	// robokassaHost = "test.robokassa.ru"
	robokassaPath = "Merchant/Index.aspx"
	// robokassaPath = "Index.aspx"
	scheme = "https"
	// scheme        = "http"
	delim = ":"
)

var (
	ErrBadRequest = errors.New("bad request")
)

type RobokassaResult struct {
	OutSum         string `json:"OutSum" xml:"OutSum" form:"OutSum"`
	InvId          string `json:"InvId" xml:"InvId" form:"InvId"`
	SignatureValue string `json:"SignatureValue" xml:"SignatureValue" form:"SignatureValue"`
}

// Robokassa для генерации URL и проверки уведомлений
type Robokassa struct {
	login          string
	firstPassword  string
	secondPassword string
}

// URL переадресации пользователя на оплату
func (client *Robokassa) URL(invoice int64, value int64, description string, isTest bool) string {
	return buildRedirectURL(client.login, client.firstPassword, invoice, value, description, isTest)
}

// CheckResult получение уведомления об исполнении операции (ResultURL)
func (client *Robokassa) CheckResult(r *RobokassaResult) bool {
	return verifyRequest(client.secondPassword, r)
}

func (client *Robokassa) ResultInvoice(r *RobokassaResult) (int, int, error) {
	return getInvoice(client.secondPassword, r)
}

// CheckSuccess проверка параметров в скрипте завершения операции (SuccessURL)
func (client *Robokassa) CheckSuccess(r *RobokassaResult) bool {
	return verifyRequest(client.firstPassword, r)
}

// NewRobokassa Robokassa
func NewRobokassa(login, password1, password2 string) *Robokassa {
	return &Robokassa{login, password1, password2}
}

// CRC of joint values with delimeter
func CRC(v ...interface{}) string {
	s := make([]string, len(v))
	for key, value := range v {
		s[key] = fmt.Sprintf("%v", value)
	}
	h := md5.New()
	io.WriteString(h, strings.Join(s, delim))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func buildRedirectURL(login, password string, invoice int64, value int64, description string, isTest bool) string {
	q := url.URL{}
	q.Host = robokassaHost
	q.Scheme = scheme
	q.Path = robokassaPath

	params := url.Values{}
	params.Add(queryLogin, login)
	params.Add(queryOutSumm, strconv.FormatInt(value, 10))
	params.Add(queryInvID, strconv.FormatInt(invoice, 10))
	params.Add(queryDescription, description)
	params.Add(queryCRC, CRC(login, value, invoice, password))
	if isTest {
		params.Add(isTestMode, "1")
	}

	q.RawQuery = params.Encode()
	return q.String()
}

func verifyResult(password string, invoice int, value string, crc string) bool {
	return strings.ToUpper(crc) == strings.ToUpper(CRC(value, invoice, password))
}

func getInvoice(password string, r *RobokassaResult) (int, int, error) {
	invoice, _ := strconv.Atoi(r.InvId)
	crc := r.SignatureValue
	price, _ := strconv.ParseFloat(r.OutSum, 10)

	if !verifyResult(password, invoice, r.OutSum, crc) {
		log.Println("result not verified")
		return 0, 0, ErrBadRequest
	}

	return invoice, int(price), nil
}

func verifyRequest(password string, r *RobokassaResult) bool {
	_, _, err := getInvoice(password, r)
	return err == nil
}
