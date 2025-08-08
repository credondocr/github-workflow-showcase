package main

import (
	"log"
	"os"

	"github.com/credondocr/github-workflow-showcase/routes"
)

func main() {
	// Get port from environment or use 8080 as default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Setup router with all routes
	router := routes.SetupRouter()

	// Startup messages
	log.Printf("🚀 Server started on port %s", port)
	log.Printf("📝 API documentation available at: http://localhost:%s/", port)
	log.Printf("🏥 Health check available at: http://localhost:%s/health", port)
	log.Printf("👥 User endpoints available at: http://localhost:%s/api/v1/users", port)

	// Start the server
	if err := router.Run(":" + port); err != nil {
		log.Fatal("❌ Error starting server:", err)
	}
}
