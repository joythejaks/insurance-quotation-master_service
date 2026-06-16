package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/jordisetiawan/insurance-master-service/docs"
	"github.com/jordisetiawan/insurance-master-service/internal/config"
	"github.com/jordisetiawan/insurance-master-service/internal/database"
	"github.com/jordisetiawan/insurance-master-service/internal/handler"
	"github.com/jordisetiawan/insurance-master-service/internal/middleware"
	"github.com/jordisetiawan/insurance-master-service/internal/repository"
	"github.com/jordisetiawan/insurance-master-service/internal/router"
	"github.com/jordisetiawan/insurance-master-service/internal/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// @title Insurance Master Service API
// @version 1.0
// @description Master Data Service for Countries, Currencies, Occupations, etc.
// @host localhost:8081
// @BasePath /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default configurations")
	}

	cfg := config.LoadConfig()
	utils.InitLogger()

	db, err := database.NewPostgres(cfg) // Re-use Postgres logic from auth_service
	if err != nil {
		utils.Log.Fatal("Failed to connect to database", zap.Error(err))
	}

	countryRepo := repository.NewCountryRepository(db)
	countryHandler := handler.NewCountryHandler(countryRepo)

	currencyRepo := repository.NewCurrencyRepository(db)
	currencyHandler := handler.NewCurrencyHandler(currencyRepo)

	occupationRepo := repository.NewOccupationRepository(db)
	occupationHandler := handler.NewOccupationHandler(occupationRepo)

	relationshipRepo := repository.NewRelationshipRepository(db)
	relationshipHandler := handler.NewRelationshipHandler(relationshipRepo)

	r := gin.Default()
	r.Use(middleware.ErrorHandler())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.SetupCountryRoutes(r, countryHandler, cfg.JWTSecret)
	router.SetupCurrencyRoutes(r, currencyHandler, cfg.JWTSecret)
	router.SetupOccupationRoutes(r, occupationHandler, cfg.JWTSecret)
	router.SetupRelationshipRoutes(r, relationshipHandler, cfg.JWTSecret)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.AppPort),
		Handler: r,
	}

	// Graceful shutdown logic
	go func() {
		utils.Log.Info("Master Service starting", zap.String("port", cfg.AppPort))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.Log.Fatal("Listen failed", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	utils.Log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		utils.Log.Fatal("Server forced to shutdown", zap.Error(err))
	}

	utils.Log.Info("Server exiting")
}
