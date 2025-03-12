package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"pawb_backend/config"
	"pawb_backend/dto"
	"pawb_backend/service"
)

func Register(c *gin.Context) {
	var user dto.RegisterDTO
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !service.IsValidEmail(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email",
		})
		return
	}

	if service.IsEmailRegistered(user.Email) {
		c.JSON(http.StatusConflict, gin.H{
			"error": "email already registered",
		})
		return
	}

	user.Password, err = service.SetPasswordHash(user.Password)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := service.Register(user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func Login(c *gin.Context) {
	loginDTO := dto.LoginDTO{}

	err := c.ShouldBind(&loginDTO)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "error binding from dto",
			"error":   err.Error(),
		})
		return
	}

	if !service.IsValidEmail(loginDTO.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email format",
		})
		return
	}

	token, err := service.Login(loginDTO)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "error - invalid email or password",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "login success",
		"token":   token,
	})
}

func Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	err := service.Logout(c, authHeader, config.RedisClient, os.Getenv("JWT_SECRET"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error logging out",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user has logged out",
	})
}
