package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jordisetiawan/insurance-master-service/internal/handler"
	"github.com/jordisetiawan/insurance-master-service/internal/middleware"
)

func SetupCurrencyRoutes(r *gin.Engine, h *handler.CurrencyHandler, secret string) {
	api := r.Group("/api/v1")
	{
		currencies := api.Group("/currencies")
		currencies.Use(middleware.AuthMiddleware(secret))
		{
			currencies.GET("", h.GetCurrencies)
			currencies.GET("/:id", h.GetCurrency)
			currencies.POST("", middleware.RoleMiddleware("ADMIN"), middleware.PermissionMiddleware("manage_master_data"), h.CreateCurrency)
			currencies.PUT("/:id", middleware.RoleMiddleware("ADMIN"), middleware.PermissionMiddleware("manage_master_data"), h.UpdateCurrency)
			currencies.DELETE("/:id", middleware.RoleMiddleware("ADMIN"), middleware.PermissionMiddleware("manage_master_data"), h.DeleteCurrency)
		}
	}
}
