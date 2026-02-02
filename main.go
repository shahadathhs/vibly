package main

import (
	"log"
	"net/http"
	"time"

	"vibly/app/handlers"
	"vibly/app/middleware"
	_ "vibly/docs"
	"vibly/pkg/config"
	"vibly/pkg/utils"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Vibly API
// @version 1.0
// @description This is the backend API for Vibly, a local live streaming platform.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
func main() {
	mux := http.NewServeMux()

	// Default routes
	mux.HandleFunc("/", handlers.RootHandler)
	mux.HandleFunc("/api", handlers.ApiRootHandler)

	mux.HandleFunc("/health", handlers.HealthHandler)

	// Swagger
	mux.HandleFunc("/docs/", httpSwagger.WrapHandler)

	// Public routes
	mux.HandleFunc("/api/auth/register", handlers.RegisterHandler)
	mux.HandleFunc("/api/auth/login", handlers.LoginHandler)

	// Wrap with logger + recover
	handler := middleware.Recover(middleware.Logger(mux))

	// Load env variables
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load environment variables: ", err)
	}

	utils.InitJWTSecret(cfg.JWTSecret)

	server := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Server running in %s mode on port %s\n", cfg.Env, cfg.Port)
	log.Fatal(server.ListenAndServe())
}
