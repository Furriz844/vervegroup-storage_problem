package model

type Promotion struct {
	ID             string  `json:"id"` //todo uuid
	Price          float64 `json:"price"`
	ExpirationDate string  `json:"expiration_date"` //todo time
}
