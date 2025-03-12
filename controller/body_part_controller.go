package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pawb_backend/service"
)

func GetBodyParts(c *gin.Context) {
	BodyParts, err := service.GetBodyParts()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Failed to get body part",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"body_parts": BodyParts,
	})
}
