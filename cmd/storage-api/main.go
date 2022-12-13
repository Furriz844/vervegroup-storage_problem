package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"storage-api/internal/app/store"
)

func getPromotions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, store.GetPromotions())
}

func GetPromotionById(c *gin.Context) {
	id := c.Param("id")
	promotion := store.GetPromotionById(id)
	if promotion == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "promotion not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, promotion)
}

func main() {
	router := gin.Default()
	router.GET("/promotions", getPromotions)
	router.GET("/promotions/:id", GetPromotionById)

	router.Run("localhost:8080")
}
