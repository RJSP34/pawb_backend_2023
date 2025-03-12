package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pawb_backend/dto"
	"pawb_backend/service"
	"strconv"
)

func GetClinicians(c *gin.Context) {
	Clinicians, err := service.GetClinicians()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Failed to get clinicians",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"clinicians": Clinicians,
	})
}

func GetAllPatientsImages(c *gin.Context) {
	clinicianID, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID in token claims",
		})
		return
	}

	patientID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid patient ID ",
		})
		return
	}

	imageDTO, err := service.GetAllPatientsImages(patientID, clinicianID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error getting feedback",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"images": imageDTO,
	})
}

func GetAllAuthorizedPatients(c *gin.Context) {
	clinicianID, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID in token claims",
		})
		return
	}

	Authopatients, err := service.GetAllAuthorizedPatients(clinicianID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error getting patients",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"patients": Authopatients,
	})
}
func GetImageByClinician(c *gin.Context) {
	clinicianID, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID in token claims",
		})
		return
	}

	imageID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid image ID ",
		})
		return
	}

	image, err := service.GetImageByClinician(imageID, clinicianID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error getting image",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"image": image,
	})

}

func GetAllImagesByClinician(c *gin.Context) {
	clinicianID, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID in token claims",
		})
		return
	}

	images, err := service.GetAllImagesByClinician(clinicianID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error getting image",
		})
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error getting image",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"images": images,
	})

}

func UpdateFeedback(c *gin.Context) {
	clinicianID, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID in token claims",
		})
		return
	}

	var feedbackDTO dto.EditFeedbackDTO
	err = c.ShouldBind(&feedbackDTO)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	err = service.UpdateFeedback(clinicianID, feedbackDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Error getting image",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Feedback updated",
	})

}

func RemoveFeedback(c *gin.Context) {
	feedbackID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	clinicianID, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID in token claims",
		})
		return
	}

	err = service.RemoveFeedback(feedbackID, clinicianID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Failed to remove",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Feedback removed",
	})
}
