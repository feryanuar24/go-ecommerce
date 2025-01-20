package routes

import (
	"github.com/gin-gonic/gin"

	"go-ecommerce/controllers"
	"go-ecommerce/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Auth
	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/register", controllers.Register)
		authRoutes.POST("/login", controllers.Login)
	}

	// User
	userRoutes := r.Group("/api/users", middlewares.AuthMiddleware("Admin"))
	{
		userRoutes.GET("/", controllers.GetAllUsers)
		userRoutes.GET("/:id", controllers.GetProfile)
		userRoutes.PUT("/:id", controllers.UpdateProfile)
		userRoutes.DELETE("/:id", controllers.DeleteAccount)
	}

	// Product
	productRoutes := r.Group("/api/products", middlewares.AuthMiddleware("Admin"))
	{
		productRoutes.GET("/", controllers.GetProducts)
		productRoutes.GET("/:id", controllers.GetProductByID)
		productRoutes.GET("/category/:category_id", controllers.GetProductsByCategory)
		productRoutes.POST("/", controllers.CreateProduct)
		productRoutes.PUT("/:id", controllers.UpdateProduct)
		productRoutes.DELETE("/:id", controllers.DeleteProduct)
	}

	// Product Category
	productCategoryRoutes := r.Group("/api/product-categories", middlewares.AuthMiddleware("Admin"))
	{
		productCategoryRoutes.GET("/", controllers.GetCategories)
		productCategoryRoutes.POST("/", controllers.CreateCategory)
	}

	// Transaction
	transactionRoutes := r.Group("/api/transactions", middlewares.AuthMiddleware("Pengguna"))
	{
		transactionRoutes.POST("/", controllers.CreateTransaction)
		transactionRoutes.GET("/", controllers.GetTransactions)
		transactionRoutes.GET("/:id", controllers.GetTransactionByID)
		transactionRoutes.GET("/user/:user_id", controllers.GetTransactionsByUser)
	}

	return r
}
