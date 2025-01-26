package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-ecommerce/config"
	"go-ecommerce/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	// Setup Gin dan koneksi database testing
	gin.SetMode(gin.TestMode)

	// Load environment testing
	config.ConnectDatabase()

	// Hapus semua data di database testing sebelum dan sesudah testing
	config.DB.Exec("DELETE FROM transaction_items")
	config.DB.Exec("DELETE FROM transactions")

	// Data dummy untuk request
	requestBody := map[string]interface{}{
		"user_id": 1,
		"items": []map[string]interface{}{
			{
				"product_id": 1,
				"quantity":   1,
				"subtotal":   9000000,
			},
			{
				"product_id": 2,
				"quantity":   2,
				"subtotal":   400000,
			},
		},
	}

	// Convert data dummy menjadi JSON
	jsonData, _ := json.Marshal(requestBody)

	// Buat permintaan HTTP
	req, _ := http.NewRequest(http.MethodPost, "/api/transactions", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	// Rekam respons
	w := httptest.NewRecorder()

	// Setup router dan handler
	router := gin.Default()
	router.POST("/api/transactions", CreateTransaction)

	// Jalankan handler
	router.ServeHTTP(w, req)

	// Validasi respons
	assert.Equal(t, http.StatusOK, w.Code)

	// Validasi respons JSON
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, "Transaction created successfully", response["message"])

	// Validasi data di database
	var transaction models.Transaction
	err := config.DB.First(&transaction, "user_id = ?", requestBody["user_id"]).Error

	// Validasi data transaction
	assert.Nil(t, err)
	assert.Equal(t, 9400000, int(transaction.Total))

	var items []models.TransactionItem
	config.DB.Where("transaction_id = ?", transaction.ID).Find(&items)

	// Validasi data transaction items
	assert.Len(t, items, 2)

	assert.Equal(t, uint(1), items[0].ProductID)
	assert.Equal(t, uint(1), items[0].Quantity)
	assert.Equal(t, 9000000, int(items[0].Subtotal))

	assert.Equal(t, uint(2), items[1].ProductID)
	assert.Equal(t, uint(2), items[1].Quantity)
	assert.Equal(t, 400000, int(items[1].Subtotal))
}
