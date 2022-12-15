package entity

import (
	"time"

	"github.com/google/uuid"
)

type Promotion struct {
	ID             uuid.UUID `json:"id"`
	Price          float64   `json:"price"`
	ExpirationDate time.Time `json:"expiration_date"`
}
