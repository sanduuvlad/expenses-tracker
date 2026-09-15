package middleware

import (
	"expense-tracker/internal/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tm *token.TokenManager) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")

		hasBearerPrefix := strings.HasPrefix(header, "Bearer ")
		if !hasBearerPrefix {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"error": "Unauthorized"},
			)
			return
		}

		tokenStr := strings.TrimPrefix(header, "Bearer ")

		claims, err := tm.ValidateToken(tokenStr)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{"error": "Unauthorized"},
			)
			return
		}

		ctx.Set("user_id", claims.UserID)
		ctx.Next()
	}
}
