package promotion

import (
	"database/sql"
	"log"
	"storage-api/internal/app/entity"
)

type PromotionRepository struct {
	db *sql.DB
}

func NewRepository(database *sql.DB) PromotionRepository {
	return PromotionRepository{database}
}

func (r PromotionRepository) getPromotionById(id string) (*entity.Promotion, error) {
	rows, err := r.db.Query("select * from promotions where id = $1", id)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer rows.Close()
	p := entity.Promotion{}
	for rows.Next() {
		if err := rows.Scan(&p.ID, &p.Price, &p.ExpirationDate); err != nil {
			log.Print(err)
			return nil, err
		}
	}
	return &p, nil
}

func (r PromotionRepository) savePromotion(p entity.Promotion) (entity.Promotion, error) {

	if _, err := r.db.Exec("INSERT INTO promotions (id, price, expiration_date) VALUES ($1, $2, $3)", p.ID, p.Price, p.ExpirationDate); err != nil {
		log.Print("Can't save promotion with id:", p.ID, ". Reason: ", err)
		return p, err
	}
	return p, nil

}

func (r PromotionRepository) clearStorage() error {
	tx, err := r.db.Begin()
	defer tx.Commit()
	if err != nil {
		log.Print(err)
		return err
	}
	if _, err := r.db.Exec("truncate table promotions"); err != nil {
		log.Print(err)
		return err
	}
	log.Print("Storage is cleared.")
	return nil
}
