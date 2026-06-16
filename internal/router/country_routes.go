package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jordisetiawan/insurance-master-service/internal/handler"
	"github.com/jordisetiawan/insurance-master-service/internal/middleware"
)

func SetupCountryRoutes(r *gin.Engine, countryHandler *handler.CountryHandler, secret string) {
	api := r.Group("/api/v1")
	{
		countries := api.Group("/countries")
		{
			// Publicly accessible GET endpoints (or protected by AuthMiddleware only)
			countries.GET("", countryHandler.GetCountries)
			countries.GET("/:id", countryHandler.GetCountry)

			// Admin-only endpoints with specific permissions
			countries.POST("", // Create Country
				middleware.AuthMiddleware(secret),
				middleware.RoleMiddleware("ADMIN"),
				middleware.PermissionMiddleware("manage_countries"), // Contoh permission
				countryHandler.CreateCountry,
			)

			countries.PUT("/:id", // Update Country
				middleware.AuthMiddleware(secret),
				middleware.RoleMiddleware("ADMIN"),
				middleware.PermissionMiddleware("manage_countries"), // Contoh permission
				countryHandler.UpdateCountry,
			)

			countries.DELETE("/:id", // Delete Country
				middleware.AuthMiddleware(secret),
				middleware.RoleMiddleware("ADMIN"),
				middleware.PermissionMiddleware("manage_countries"), // Contoh permission
				countryHandler.DeleteCountry,
			)
		}
	}
}
