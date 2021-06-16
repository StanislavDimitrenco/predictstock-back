package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/Paramosch/predictstock-backend-eng/file_parser"
	"github.com/Paramosch/predictstock-backend-eng/logger"
	"github.com/gofiber/fiber/v2"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type FileUploadRequest struct {
	Hash      string `json:"hash" xml:"hash" form:"hash"`
	Timestamp int64  `json:"timestamp" xml:"timestamp" form:"timestamp"`
}

func UploadSharesCSV(c *fiber.Ctx) error {
	fileLog := logger.NewLogger("csv_requests")
	request := new(FileUploadRequest)
	if err := c.BodyParser(request); err != nil {
		fileLog.LogWarning(err)
		return c.Status(412).SendString("Undefined request")
	}

	nowDate := time.Now()
	requestTimestamp := time.Unix(request.Timestamp, 0)
	diff := nowDate.Sub(requestTimestamp).Seconds()

	hasher := md5.New()
	hasher.Write([]byte(fmt.Sprintf("%d-%s", request.Timestamp, os.Getenv("CONTROL_SUM"))))
	computedHash := hex.EncodeToString(hasher.Sum(nil))

	failRequestSeconds, err := strconv.ParseFloat(os.Getenv("FAIL_REQUEST_SECOND"), 64)
	if computedHash != request.Hash || diff >= failRequestSeconds || err != nil {
		fileLog.SetFields(logger.Fields{
			"ip": c.IP(),
		}).LogWarning("Access denied")

		return c.Status(412).SendString("Access denied!")
	}

	file, err := c.FormFile("document")
	if err != nil {
		return c.SendString("File undefined!")
	}

	var fileExt = filepath.Ext(file.Filename)
	if fileExt != ".csv" {
		fileLog.SetFields(logger.Fields{
			"ip":       c.IP(),
			"fileName": file.Filename,
		}).LogWarning("Invalid .csv file format")

		return c.Status(412).SendString("Support only .csv file format!")
	}

	currentTime := time.Now()
	fileName := fmt.Sprintf("%s%s", currentTime.Format("2006-01-02 15:04:05"), fileExt)
	var filePath = fmt.Sprintf("shares_files/%s", fileName)

	// Save file to root directory:
	err = c.SaveFile(file, filePath)
	if err != nil {
		fileLog.LogWarning(err)
		return c.Status(412).SendString("File cannot be uploaded!")
	}

	file_parser.PushFileForParse(filePath)

	fileLog.SetFields(logger.Fields{
		"ip":       c.IP(),
		"fileName": fileName,
	}).LogInfo("File uploaded")

	return c.Status(201).SendString("File uploaded")
}
