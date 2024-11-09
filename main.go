package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struct untuk login request
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Handler untuk route GET
func getHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Ini adalah route GET",
	})
}

// Handler untuk route dengan parameter di URL
func getParamHandler(c *gin.Context) {
	param := c.Param("param")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Ini adalah route GET dengan parameter %s", param),
	})
}

// Handle untuk route login
func loginHandler(c *gin.Context) {
	var loginRequest LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if loginRequest.Username != "admin" || loginRequest.Password != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Username atau password salah",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Login berhasil dengan username %s", loginRequest.Username),
	})
}

// Handler untuk route dengan query parameter
func queryParamHandler(c *gin.Context) {
	param := c.Query("param")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Ini adalah route GET dengan query parameter %s", param),
	})
}

func main() {
	r := gin.Default()

	// Route GET
	r.GET("/", getHandler)

	// Route GET dengan parameter di URL
	r.GET("/:param", getParamHandler)

	// Route POST
	r.POST("/login", loginHandler)

	// Route GET dengan query parameter
	r.GET("/query", queryParamHandler)

	r.Run(":8080")
}
