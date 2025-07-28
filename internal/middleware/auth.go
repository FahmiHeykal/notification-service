package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		log.Printf("Authorization header: %s", tokenString)

		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header required"})
			ctx.Abort()
			return
		}

		if len(tokenString) > 7 && strings.ToUpper(tokenString[0:7]) == "BEARER " {
			tokenString = tokenString[7:]
		}

		log.Printf("Token after processing: %s", tokenString)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil {
			log.Printf("Token parse error: %v", err)
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		ctx.Set("userID", uint(claims["user_id"].(float64)))
		ctx.Set("username", claims["username"].(string))
		ctx.Next()
	}
}
