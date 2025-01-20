package controllers

import (
	"go-ecommerce/config"
	"go-ecommerce/models"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categories []models.ProductCategory
	var wg sync.WaitGroup
	var err error

	wg.Add(1)
	go func() {
		defer wg.Done()
		// Query untuk mendapatkan semua Product Category
		if err = config.DB.Find(&categories).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
			return
		}
	}()

	wg.Wait()

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    categories,
	})
}

func CreateCategory(c *gin.Context) {
	var input models.ProductCategory

	// Validasi input JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Query untuk membuat Product Category baru
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category created",
		"data":    input,
	})
}
