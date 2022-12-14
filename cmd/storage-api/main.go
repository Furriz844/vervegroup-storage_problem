package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"storage-api/internal/app/store"

	"fmt"

	"database/sql"

	_ "github.com/lib/pq"

	"log"
)

var DB *sql.DB
var repo store.PromotionRepository

// todo to settings file
const (
	HOST     = "localhost"
	PORT     = 55000
	USER     = "postgres"
	PASSWORD = "postgrespw"
	DBNAME   = "storage_api"
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
	repo = store.NewRepository(DB)

	router := gin.Default()
	router.GET("/promotions", getPromotions)
	router.GET("/promotions/:id", GetPromotionById)
	router.GET("/admin/load", LoadPromotions)

	router.Run("localhost:8080")
}

func NewRepository(DB *sql.DB) {
	panic("unimplemented")
}

func getPromotions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, repo.GetPromotions())
}

func GetPromotionById(c *gin.Context) {
	id := c.Param("id")
	promotion, err := repo.GetPromotionById(id)
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
	repo.LoadFromCsv()
	c.IndentedJSON(http.StatusOK, gin.H{"status": "ok"})
}
