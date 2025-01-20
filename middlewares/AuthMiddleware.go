package middlewares

import (
	"go-ecommerce/config"
	"go-ecommerce/models"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Format: "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer <token>"})
			c.Abort()
			return
		}

		// Parse token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validasi algoritma
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError("invalid signing method", jwt.ValidationErrorSignatureInvalid)
			}
			// Kembalikan secret key
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Ambil klaim token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Simpan userID ke context
		userID := uint(claims["user_id"].(float64)) // Pastikan klaim user_id ada
		c.Set("userID", userID)

		// Cek role pengguna dari database
		var user models.User
		if err := config.DB.Where("id = ?", userID).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// Periksa role pengguna
		if user.Role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to access this resource"})
			c.Abort()
			return
		}

		// Lanjutkan request
		c.Next()
	}
}
