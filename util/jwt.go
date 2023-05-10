package util

import (
	"fmt"
	"os"
	"time"

	"github.com/ankeshnirala/go/aws-iam-service/types"
	"github.com/golang-jwt/jwt/v4"
)

var JWT_SECRET = os.Getenv("JWT_SECRET")

func CreateJWT(user *types.User) (string, error) {

	// Create JWT token
	expirationTime := time.Now().Add(50 * time.Minute)

	claims := &types.Claims{
		Email:  user.Email,
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(JWT_SECRET))
}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(JWT_SECRET), nil
	})
}
