package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"storage-api/internal/app/promotion"

	"fmt"

	"database/sql"

	_ "github.com/lib/pq"

	"log"
)

var DB *sql.DB
var service promotion.PromotionService

// todo to settings file
const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "storage"
	PASSWORD = "storage_pwd"
	DBNAME   = "storage_api"
	FILE     = "./resources/promotions.csv"
)

func main() {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DBNAME,
	)

	DB, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()
	repo := promotion.NewRepository(DB)
	service = promotion.NewService(&repo)

	router := gin.Default()
	router.GET("/promotions/:id", GetPromotionById)
	router.GET("/admin/load", LoadPromotions)

	router.Run("localhost:8080")
}

func NewRepository(DB *sql.DB) {
	panic("unimplemented")
}

func GetPromotionById(c *gin.Context) {
	id := c.Param("id")
	promotion, err := service.GetPromotionById(id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if promotion == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "promotion not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, promotion)
}
func LoadPromotions(c *gin.Context) {
	service.LoadFromCsv(FILE)
	c.IndentedJSON(http.StatusOK, gin.H{"status": "ok"})
}
