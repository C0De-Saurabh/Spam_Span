package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var SecretKey = []byte("secret-key") // Replace this with a real secret from an environment variable

// Claims struct to represent the structure of the JWT claims
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

// GenerateJWT creates a JWT token for the given user ID
func GenerateJWT(userID uint) (string, error) {
	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expiration (24 hours)
		},
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates the JWT token and returns the custom claims
func ValidateToken(tokenString string) (*Claims, error) {
	// Parse and validate the token with the claims structure
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil // Return the secret key for verification
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Extract and return the custom claims
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("could not extract claims")
	}

	return claims, nil
}
