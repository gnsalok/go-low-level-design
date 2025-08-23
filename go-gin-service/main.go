package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/gnsalok/go-low-level-design/go-gin-service/config"
	"github.com/gnsalok/go-low-level-design/go-gin-service/handlers"
	"github.com/gnsalok/go-low-level-design/go-gin-service/middleware"
)

func main() {
	log.Println("Loading configuration...")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("FATAL: Failed to load configuration: %v", err)
	}
	log.Println("Configuration loaded successfully.")

	router := gin.Default()

	router.POST("/login", handlers.GenerateToken(cfg.JWT.SecretKey, cfg.JWT.Issuer))
	router.GET("/public", handlers.PublicEndpoint)

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware(cfg.JWT.SecretKey))
	{
		protected.GET("/data", handlers.ProtectedEndpoint)
	}

	serverAddr := cfg.Server.Port
	log.Printf("Starting service on %s...", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("FATAL: Failed to start server: %v", err)
	}
}
