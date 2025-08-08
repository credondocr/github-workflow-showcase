package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/credondocr/github-workflow-showcase/controllers"
	"github.com/credondocr/github-workflow-showcase/models"
)

// SetupRouter configures and returns the router with all defined routes.
func SetupRouter() *gin.Engine {
	// Create in-memory repository
	userRepo := models.NewInMemoryUserRepository()

	// Create controller
	userController := controllers.NewUserController(userRepo)

	// Configure Gin in release mode for production (optional)
	// gin.SetMode(gin.ReleaseMode)

	// Create router
	router := gin.Default()

	// Logging middleware
	router.Use(gin.Logger())

	// Recovery middleware in case of panic
	router.Use(gin.Recovery())

	// Basic CORS middleware for development
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Health routes
	router.GET("/health", userController.HealthCheck)

	// API version 1 route group
	v1 := router.Group("/api/v1")
	{
		// User routes
		users := v1.Group("/users")
		users.GET("", userController.GetUsers)          // GET /api/v1/users
		users.GET("/stats", userController.GetStats)    // GET /api/v1/users/stats
		users.GET("/:id", userController.GetUser)       // GET /api/v1/users/:id
		users.POST("", userController.CreateUser)       // POST /api/v1/users
		users.PUT("/:id", userController.UpdateUser)    // PUT /api/v1/users/:id
		users.DELETE("/:id", userController.DeleteUser) // DELETE /api/v1/users/:id
	}

	// Root route that shows API information
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the example REST API!",
			"version": "1.0.0",
			"endpoints": gin.H{
				"health":     "GET /health",
				"users":      "GET /api/v1/users",
				"userStats":  "GET /api/v1/users/stats",
				"user":       "GET /api/v1/users/:id",
				"createUser": "POST /api/v1/users",
				"updateUser": "PUT /api/v1/users/:id",
				"deleteUser": "DELETE /api/v1/users/:id",
			},
		})
	})

	return router
}
