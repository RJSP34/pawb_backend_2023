package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pawb_backend/dto"
	"pawb_backend/service"
	"strconv"
)

func GetFeedbackByClinician(c *gin.Context) {
	userID, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID in token claims",
		})
		return
	}
	feedbackDTO, err := service.GetFeedbackbyClinician(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error getting feedback",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"feedback": feedbackDTO,
	})
}

func GetFeedbackByImage(c *gin.Context) {
	userID, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID in token claims",
		})
		return
	}

	imageID, err := strconv.ParseUint(c.Param("id"), 10, 64)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid image ID in parameters",
		})
		return
	}

	feedbackDTO, err := service.GetFeedbackByImage(imageID, userID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"feedback": feedbackDTO,
	})
}

func SubmitFeedback(c *gin.Context) {
	clinicianID, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID in token claims",
		})
		return
	}

	var feedbacksub dto.ClinicianFeedbackSubmissionDTO
	err = c.ShouldBind(&feedbacksub)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid image ID in parameters",
		})
		return
	}

	err = service.SubmitFeedback(feedbacksub, clinicianID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to submit feedback",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "feedback submitted successfully",
	})

}
