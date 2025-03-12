package service

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"os"
	"strconv"
	"strings"
	"time"
)

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func CreateToken(id uint64, email string, role string) (string, error) {
	userID := strconv.FormatUint(id, 10)
	claims := &jwtCustomClaim{
		userID,
		role,
		jwt.StandardClaims{
			Issuer:    "com.example.pawb_project",
			Subject:   email,
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signature := os.Getenv("JWT_SECRET")

	t, err := token.SignedString([]byte(signature))
	return t, err
}

func ValidateToken(token string) (*jwt.Token, error) {
	token = strings.Replace(token, "Bearer ", "", -1)
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}

		signature := os.Getenv("JWT_SECRET")
		return []byte(signature), nil
	})
}
