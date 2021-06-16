package file_parser

import (
	"encoding/csv"
	"github.com/Paramosch/predictstock-backend-eng/database"
	"github.com/Paramosch/predictstock-backend-eng/logger"
	"io"
	"os"
	"strconv"
)

type ShareCreator interface {
	Create(share *database.Share) *database.Share
	FindBy(param map[string]interface{}) (*database.Share, bool)
	Save(share *database.Share) *database.Share
}

var Files = make(chan File)

func PushFileForParse(pathName string) {
	go func() { Files <- File{pathName: pathName} }()
}

func Parser(repo ShareCreator) {
	for {
		select {
		case file := <-Files:
			csvFile, err := os.Open(file.PathName())
			logFile := logger.NewLogger("csv_parser")
			if err != nil {
				logFile.LogError("Couldn't open the csv file", err)
			}

			r := csv.NewReader(csvFile)

			// Iterate through the records
			for {
				// Read each record from csv
				record, err := r.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					logFile.LogError(err)
				}
				var name = record[0]
				var symbol = record[1]
				rating, _ := strconv.Atoi(record[2])

				share, notFound := repo.FindBy(map[string]interface{}{"name": name})
				if notFound {
					share := &database.Share{Name: name, Symbol: symbol, Rating: rating}
					repo.Create(share)
				} else {
					share.Name = name
					share.Symbol = symbol
					share.Rating = rating
					share = repo.Save(share)
				}
			}
			logFile.LogInfo("Import success")
		}
	}
}
