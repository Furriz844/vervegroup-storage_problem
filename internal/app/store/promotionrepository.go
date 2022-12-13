package store

import (
	"encoding/csv"
	"log"
	"os"
	"storage-api/internal/app/model"
	"strconv"
)

var promotions = []*model.Promotion{}

func GetPromotions() []*model.Promotion {
	return promotions
}

func GetPromotionById(id string) *model.Promotion {
	for _, a := range promotions {
		if a.ID == id {
			return a
		}
	}
	return nil
}

func LoadFromCsv() {
	csvFile := "./promotions.csv"
	f, err := os.Open(csvFile)
	if err != nil {
		log.Fatal("Unable to read input file "+csvFile, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+csvFile, err)
	}
	for _, record := range records {
		//todo check 3 element in array
		price, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal("Can't parse record with id=", record[0], "Can't parse Price:", err)
		}
		promotions = append(promotions, &model.Promotion{ID: record[0], Price: price, ExpirationDate: record[2]})
	}
}
