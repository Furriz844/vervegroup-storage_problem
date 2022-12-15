package promotion

import (
	"encoding/csv"
	"log"
	"os"
	"storage-api/internal/app/entity"
	"strconv"
	"time"

	"github.com/google/uuid"
)

const (
	// Mon Jan 2 15:04:05 MST 2006
	layout = "2006-01-02 15:04:05 -0700 MST"
)

type PromotionService struct {
	repo *PromotionRepository
}

func NewService(repo *PromotionRepository) PromotionService {
	return PromotionService{repo}
}

func (s PromotionService) GetPromotionById(id string) (*entity.Promotion, error) {
	return s.repo.getPromotionById(id)

}

func (s PromotionService) LoadFromCsv(filepath string) error {
	if err := s.repo.clearStorage(); err != nil {
		panic(err)
	}
	f, err := os.Open(filepath)
	if err != nil {
		log.Print("Unable to read input file", filepath, err)
		return err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for", filepath, err)
		return err
	}

	for _, record := range records {
		if len(record) < 3 {
			log.Print("Record contains less then 3 argument:", record)
			continue
		}
		price, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Print("Can't parse record with id=", record[0], "Can't parse Price:", err)
		}
		id, err := uuid.Parse(record[0])
		if err != nil {
			log.Print("Can't parse record with id=", record[0], "Can't parse ID:", err)
		}
		promTime, err := time.Parse(layout, record[2])
		if err != nil {
			log.Print("Can't parse record with id=", record[0], "Can't parse ID:", err)
		}
		p := entity.Promotion{ID: id, Price: price, ExpirationDate: promTime}
		if _, err := s.repo.savePromotion(p); err != nil {
			log.Print("Can't save promotion with id=", record[0], err)
		} else {
			log.Print("Promotion ", p, " saved to repository.")
		}
	}
	return nil
}
