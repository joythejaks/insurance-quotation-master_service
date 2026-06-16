package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jordisetiawan/insurance-master-service/internal/handler"
	"github.com/jordisetiawan/insurance-master-service/internal/middleware"
)

func SetupOccupationRoutes(r *gin.Engine, h *handler.OccupationHandler, secret string) {
	api := r.Group("/api/v1")
	{
		occupations := api.Group("/occupations")
		occupations.Use(middleware.AuthMiddleware(secret))
		{
			occupations.GET("", h.GetOccupations)
			occupations.GET("/:id", h.GetOccupation)
			occupations.POST("", middleware.RoleMiddleware("ADMIN"), middleware.PermissionMiddleware("manage_master_data"), h.CreateOccupation)
			occupations.PUT("/:id", middleware.RoleMiddleware("ADMIN"), middleware.PermissionMiddleware("manage_master_data"), h.UpdateOccupation)
			occupations.DELETE("/:id", middleware.RoleMiddleware("ADMIN"), middleware.PermissionMiddleware("manage_master_data"), h.DeleteOccupation)
		}
	}
}
