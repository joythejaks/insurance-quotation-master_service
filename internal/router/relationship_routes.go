package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jordisetiawan/insurance-master-service/internal/handler"
	"github.com/jordisetiawan/insurance-master-service/internal/middleware"
)

func SetupRelationshipRoutes(r *gin.Engine, h *handler.RelationshipHandler, secret string) {
	api := r.Group("/api/v1")
	{
		relationships := api.Group("/relationships")
		relationships.Use(middleware.AuthMiddleware(secret))
		{
			relationships.GET("", h.GetRelationships)
			relationships.GET("/:id", h.GetRelationship)
			relationships.POST("", middleware.RoleMiddleware("ADMIN"), middleware.PermissionMiddleware("manage_master_data"), h.CreateRelationship)
			relationships.PUT("/:id", middleware.RoleMiddleware("ADMIN"), middleware.PermissionMiddleware("manage_master_data"), h.UpdateRelationship)
			relationships.DELETE("/:id", middleware.RoleMiddleware("ADMIN"), middleware.PermissionMiddleware("manage_master_data"), h.DeleteRelationship)
		}
	}
}
