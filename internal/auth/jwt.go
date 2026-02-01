// package auth

// import (
// 	"errors"
// 	"os"
// 	"time"

// 	"github.com/golang-jwt/jwt/v5"
// )

// type Claims struct {
// 	UserID uint   `json:"user_id"`
// 	Email  string `json:"email"`
// 	Role   string `json:"role"`
// 	jwt.RegisteredClaims
// }

// var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// // Generate JWT Token
// func GenerateToken(userID uint, email string, role string) (string, error) {
// 	claims := Claims{
// 		UserID: userID,
// 		Email:  email,
// 		Role:   role,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
// 			IssuedAt:  jwt.NewNumericDate(time.Now()),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString(jwtSecret)
// }

// // Validate JWT Token
// func ValidateToken(tokenString string) (*Claims, error) {
// 	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
// 		return jwtSecret, nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
// 		return claims, nil
// 	}

//		return nil, errors.New("invalid token")
//	}
package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// Init initializes JWT with secret
func Init(secret string) error {
	if secret == "" {
		return errors.New("JWT secret cannot be empty")
	}
	jwtSecret = []byte(secret)
	return nil
}

// GenerateToken creates a new JWT token for user
func GenerateToken(userID uint, email string, role string) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateToken validates and parses JWT token
func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
