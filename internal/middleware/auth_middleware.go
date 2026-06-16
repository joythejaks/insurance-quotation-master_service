package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jordisetiawan/insurance-master-service/internal/utils"
)

type Claims struct {
	UserID      string   `json:"user_id"`
	Email       string   `json:"email"`
	FullName    string   `json:"full_name"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
	jwt.RegisteredClaims
}

func AuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Authorization header is required", nil)
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid authorization header format", nil)
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid or expired token", err.Error())
			c.Abort()
			return
		}

		claims, _ := token.Claims.(*Claims)
		c.Set("user_id", claims.UserID)
		c.Set("full_name", claims.FullName) // Menambahkan FullName ke context
		c.Set("role", claims.Role)
		c.Set("permissions", claims.Permissions)
		c.Next()
	}
}

func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, _ := c.Get("role")
		for _, role := range roles {
			if role == userRole {
				c.Next()
				return
			}
		}
		utils.ErrorResponse(c, http.StatusForbidden, "Access denied: insufficient role permissions", nil)
		c.Abort()
	}
}

func PermissionMiddleware(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		permissions, exists := c.Get("permissions")
		if !exists {
			utils.ErrorResponse(c, http.StatusForbidden, "Access denied: no permissions assigned", nil)
			c.Abort()
			return
		}

		userPerms := permissions.([]string)
		found := false
		for _, p := range userPerms {
			if p == permission {
				found = true
				break
			}
		}

		if !found {
			utils.ErrorResponse(c, http.StatusForbidden, "Access denied: missing required permission", permission)
			c.Abort()
			return
		}

		c.Next()
	}
}
