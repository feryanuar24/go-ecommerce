package controllers

import (
	"go-ecommerce/config"
	"go-ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product

	// Query untuk mendapatkan semua Product beserta relasi dengan Product Category
	if err := config.DB.Preload("ProductCategory").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data": products,
	})
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	// Query untuk mendapatkan Product berdasarkan ID beserta relasi dengan Product Category
	if err := config.DB.Preload("ProductCategory").First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data": product,
	})
}

func GetProductsByCategory(c *gin.Context) {
	categoryID := c.Param("category_id")

	var products []models.Product

	// Query untuk mendapatkan semua Product berdasarkan ID Product Category
	if err := config.DB.Preload("ProductCategory").Where("product_category_id = ?", categoryID).Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"data": products,
	})
}

func CreateProduct(c *gin.Context) {
	var input models.Product

	// Validasi input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Query untuk membuat Product baru dengan relasi Product Category
	if err := config.DB.Preload("ProductCategory").Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product created successfully",
		"data": input,
	})
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	// Query untuk mengambil Product berdasarkan ID, termasuk relasi Product Category
	if err := config.DB.Preload("ProductCategory").First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var input struct {
		Name              string  `json:"name"`
		Description       string  `json:"description"`
		Price             float64 `json:"price"`
		ProductCategoryID uint    `json:"product_category_id"`
	}

	// Validasi input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perbarui data Product dengan input baru
	product.Name = input.Name
	product.Description = input.Description
	product.Price = input.Price
	product.ProductCategoryID = input.ProductCategoryID

	// Simpan perubahan ke database
	if err := config.DB.Model(&product).Updates(product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	// Ambil data relasi ProductCategory yang diperbarui
	if err := config.DB.Preload("ProductCategory").First(&product, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated product with category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product updated successfully",
		"data": product,
	})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	// Query untuk mendapatkan Product berdasarkan ID, termasuk relasi ProductCategory
	if err := config.DB.Preload("ProductCategory").First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	// Query untuk menghapus Product berdasarkan ID
	if err := config.DB.Preload("ProductCategory").Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
		"data": product,
	})
}
