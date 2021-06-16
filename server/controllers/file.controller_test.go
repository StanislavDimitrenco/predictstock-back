package controllers

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/Paramosch/predictstock-backend-eng/providers"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
	"time"
)

func setUp() *fiber.App {
	ctx := providers.Boot(context.Background())
	app := ctx.Value("fiber").(*fiber.App)

	app.Post("/upload", UploadSharesCSV)

	return app
}

func TestUploadSharesCSV(t *testing.T) {
	app := setUp()
	timestamp := time.Now().Unix()

	hasher := md5.New()
	hasher.Write([]byte(fmt.Sprintf("%v-%s", timestamp, os.Getenv("CONTROL_SUM"))))
	computedHash := hex.EncodeToString(hasher.Sum(nil))

	body, contentType, err := getRequestBody(computedHash, timestamp)
	if err != nil {
		t.Fail()
		return
	}
	// http.Request
	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Add("Content-Type", contentType)

	// http.Response
	resp, _ := app.Test(req)

	// Do something with results:
	if resp == nil {
		t.Error("Undefined response")
	} else if resp.StatusCode != 201 {
		t.Errorf("Wrond Status Code: %d, expected: 200", resp.StatusCode)
	}
}

func TestUploadSharesCSVInvalidTime(t *testing.T) {
	app := setUp()
	duration, _ := strconv.ParseInt(os.Getenv("FAIL_REQUEST_SECOND"), 0, 10)
	duration = int64(^uint64(duration - 1))
	durationConverted := (duration - 1) * int64(time.Second)
	timestamp := time.Now().Add(time.Duration(durationConverted)).Unix()
	hasher := md5.New()
	hasher.Write([]byte(fmt.Sprintf("%v-%s", timestamp, os.Getenv("CONTROL_SUM"))))
	computedHash := hex.EncodeToString(hasher.Sum(nil))

	body, contentType, err := getRequestBody(computedHash, timestamp)
	if err != nil {
		t.Fail()
		return
	}

	// http.Request
	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Add("Content-Type", contentType)

	// http.Response
	resp, _ := app.Test(req)

	// Do something with results:
	if resp == nil {
		t.Error("Undefined response")
	} else if resp.StatusCode != 412 {
		t.Errorf("Wrond Status Code: %d, expected: 200", resp.StatusCode)
	}
}

func getRequestBody(hash string, timestamp int64) (*bytes.Buffer, string, error) {
	fileContents, err := getFileContent()
	if err != nil {

	}
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("document", "test.csv")

	if err != nil {
		return nil, "", err
	}
	_, err = part.Write(fileContents)
	if err != nil {
		return nil, "", err
	}

	params := map[string]string{
		"hash":      hash,
		"timestamp": strconv.FormatInt(timestamp, 10),
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}

	return body, writer.FormDataContentType(), nil
}

func getFileContent() ([]byte, error) {
	file := getValidFile()
	if file == nil {
		return nil, errors.New("file undefined")
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}

	return fileContents, nil
}

func getValidFile() *os.File {
	file, err := os.Open("valid_test.csv")
	if err != nil {
		return nil
	}

	return file
}
