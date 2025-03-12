package service

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"pawb_backend/dto"
	"pawb_backend/entity"
	"pawb_backend/repository"
	"strings"
	"time"
)

func SetPasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func IsValidEmail(email string) bool {
	return govalidator.IsEmail(email)
}

func IsEmailRegistered(email string) bool {
	return repository.IsEmailRegistered(email)
}

func Register(userDTO dto.RegisterDTO) (string, error) {
	var user entity.User
	err := smapping.FillStruct(&user, smapping.MapFields(&userDTO))
	key, err := generateKey()
	if err != nil {
		return "", err
	}
	user.PrivateKey = key
	token, _ := CreateToken(user.ID, user.Email, user.Role.Description)
	user, err = repository.Register(user)
	if err != nil {
		return "", err
	}
	token, _ = CreateToken(user.ID, user.Email, user.Role.Description)
	return token, nil
}

func generateKey() ([]byte, error) {
	// Generate a new AES key
	key := make([]byte, 16) // AES-256 requires a 32-byte key
	_, err := rand.Read(key)
	if err != nil {
		return []byte(""), fmt.Errorf("failed to generate AES key: %w", err)
	}

	return key, nil
}

func Login(loginDTO dto.LoginDTO) (string, error) {
	token := ""
	user, err := repository.Login(loginDTO)
	if err != nil {
		return "", err
	}
	err = CheckPasswordHash(loginDTO.Password, user.Password)
	if err != nil {
		return "", err
	}
	token, err = CreateToken(user.ID, user.Email, user.Role.Description)
	return token, err
}

func Logout(c *gin.Context, token string, redisClient *redis.Client, secretKey string) error {
	token = strings.Replace(token, "Bearer ", "", -1)
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(secretKey), nil
	})

	// Check for token parsing errors
	if err != nil {
		return fmt.Errorf("failed to parse token: %v", err)
	}

	// Check if the token is valid
	if _, ok := parsedToken.Claims.(jwt.MapClaims); !ok || !parsedToken.Valid {
		return fmt.Errorf("invalid token")
	}

	// Extract the user ID or any other necessary information from the token claims
	exp := parsedToken.Claims.(jwt.MapClaims)["exp"].(float64)
	expirationTime := time.Unix(int64(exp), 0)

	// Add the user's token to the Redis blocklist
	err = redisClient.Set(c.Request.Context(), token, "blocked", expirationTime.Sub(time.Now())+1).Err()
	if err != nil {
		return fmt.Errorf("failed to block token: %v", err)
	}

	return nil
}
