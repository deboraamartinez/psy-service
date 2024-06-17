package middlewares

import (
	"net/http"
	"os"
	"psy-service/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JwtVerify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			utils.RespondWithError(ctx, http.StatusUnauthorized, "Missing authorization header")
			ctx.Abort()
			return
		}

		tokenString := strings.Split(authHeader, " ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			utils.RespondWithError(ctx, http.StatusUnauthorized, "Invalid token")
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx.Set("user_id", claims["user_id"])
		} else {
			utils.RespondWithError(ctx, http.StatusUnauthorized, "Invalid token claims")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
