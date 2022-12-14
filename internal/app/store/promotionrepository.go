package store

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"
	"storage-api/internal/app/model"
	"strconv"
)

type PromotionRepository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) PromotionRepository {
	return PromotionRepository{database}
}

var promotions = []*model.Promotion{}

func (r PromotionRepository) GetPromotions() []*model.Promotion {
	return promotions
}

func (r PromotionRepository) GetPromotionById(id string) (*model.Promotion, error) {
	rows, err := r.db.Query("select * from promotion where id = $1", id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()
	p := model.Promotion{}
	for rows.Next() {
		if err := rows.Scan(&p.ID, &p.Price, &p.ExpirationDate); err != nil {
			log.Fatal(err)
			return nil, err
		}
	}
	return &p, nil
}

func (r PromotionRepository) LoadFromCsv() {
	csvFile := "./resources/promotions.csv"
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
	tx, err := r.db.Begin()
	for _, record := range records {
		//todo check 3 element in array
		price, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal("Can't parse record with id=", record[0], "Can't parse Price:", err)
		}
		//promotions = append(promotions, &model.Promotion{ID: record[0], Price: price, ExpirationDate: record[2]})
		res, err := r.db.Exec("INSERT INTO promotion (id, price, date) VALUES ($1, $2, $3)", record[0], price, record[2])
		if err != nil {
			log.Fatal(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("affected = %d\n", rowCnt)
	}
	tx.Commit()
}
