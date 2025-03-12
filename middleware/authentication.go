package middleware

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"net/http"
	"pawb_backend/config"
	"pawb_backend/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authorized(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header is required",
			})
			return
		}

		token, err := service.ValidateToken(authHeader)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token claims",
			})
			return
		}

		isBlocked, err := isTokenBlocked(c, token, config.RedisClient)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": err.Error(),
			})
			return
		}

		if isBlocked {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "User has already logged out",
			})
			return
		}

		if len(roles) > 0 && !isAuthorized(claims, roles) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "User not authorized to access this resource",
			})
			return
		}

		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}

func isAuthorized(claims jwt.MapClaims, roles []string) bool {
	role, ok := claims["role"].(string)
	if !ok {
		return false
	}
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}

func isTokenBlocked(c *gin.Context, token *jwt.Token, redisClient *redis.Client) (bool, error) {
	ctx := c.Request.Context()

	exists, err := redisClient.Exists(ctx, token.Raw).Result()

	if err != nil {
		return false, fmt.Errorf("failed to check token status: %v", err)
	}

	return exists == 1, nil
}
