package controllers

import (
	"go-ecommerce/config"
	"go-ecommerce/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var input struct {
		UserID uint `json:"user_id"`
		Items  []struct {
			ProductID uint    `json:"product_id"`
			Quantity  uint    `json:"quantity"`
			Subtotal  float64 `json:"subtotal"`
		} `json:"items"`
	}

	// Validasi input JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi item kosong
	if len(input.Items) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Items cannot be empty"})
		return
	}

	// Hitung total transaksi
	var total float64
	for _, item := range input.Items {
		total += item.Subtotal
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")

	// Buat transaksi baru
	transaction := models.Transaction{
		UserID:    input.UserID,
		Total:     total,
		Status:    "pending",
		CreatedAt: time.Now().In(loc),
	}

	// Simpan transaksi ke database
	if err := config.DB.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	// Buat item transaksi
	for _, item := range input.Items {
		transactionItem := models.TransactionItem{
			TransactionID: transaction.ID,
			ProductID:     item.ProductID,
			Quantity:      item.Quantity,
			Subtotal:      item.Subtotal,
		}

		// Simpan item transaksi ke database
		if err := config.DB.Create(&transactionItem).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction item"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Transaction created successfully",
		"transaction": transaction,
	})
}

func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction

	// Query transaksi beserta itemnya
	if err := config.DB.Preload("User").Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    transactions,
	})
}

func GetTransactionByID(c *gin.Context) {
	var transaction models.Transaction

	// Ambil ID transaksi dari URL
	id := c.Param("id")

	// Query Transaksi berdasarkan ID
	if err := config.DB.Preload("User").First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    transaction,
	})
}

func GetTransactionsByUser(c *gin.Context) {
	var transactions []models.Transaction

	// Ambil ID user dari URL
	userID := c.Param("user_id")

	// Query transaksi berdasarkan ID user
	if err := config.DB.Preload("User").Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data":    transactions,
	})
}
