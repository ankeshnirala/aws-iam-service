package api

import (
	"net/http"

	"github.com/ankeshnirala/go/aws-iam-service/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the value of the token with the name "x-jwt-token"
		tokenString := c.Request.Header.Get("x-jwt-token")
		if tokenString == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token, err := util.ValidateJWT(tokenString)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Check if the value of the token is valid
		if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// If the token is valid, continue with the next middleware/handler
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("userID", claims["userID"])
		c.Next()
	}
}
