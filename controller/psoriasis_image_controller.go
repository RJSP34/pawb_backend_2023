package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pawb_backend/dto"
	"pawb_backend/entity"
	"pawb_backend/service"
	"strconv"
	"strings"
)

func SubmitImage(c *gin.Context) {
	userID, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID in token claims",
		})
		return
	}

	var psoriasisImageDTO dto.PsoriasisImageSubmitDTO
	err = c.ShouldBind(&psoriasisImageDTO)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	// Extract image MIME and data from request body
	imageMIME := psoriasisImageDTO.Image.Mime
	imageData, err := service.DecodeBase64(psoriasisImageDTO.Image.Data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid base64-encoded image data",
			"error":   err.Error(),
		})
		return
	}

	if !strings.HasPrefix(imageMIME, "image/") {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid image MIME type",
		})
		return
	}

	const MaxImageSize = 20 * 1024 * 1024 // 20 MB in bytes

	// Check if the image size is within limits
	if len(imageData) > MaxImageSize {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Image size exceeds limit",
		})
		return
	}

	psoriasisImage := entity.PsoriasisImage{
		Description: psoriasisImageDTO.Description,
		BodyPartID:  psoriasisImageDTO.BodyPartID,
		UserID:      userID,
		ImageData:   imageData,
		Mime:        imageMIME,
	}

	err = service.SubmitImage(&psoriasisImage)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to submit image",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Image submitted successfully",
	})
}

func GetImage(c *gin.Context) {
	userID, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	imageID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid image ID in parameters",
		})
		return
	}

	image, err := service.GetImage(imageID)
	if err != nil || image.UserID != userID {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Failed to get image",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"image": image,
	})
}

func GetMyImages(c *gin.Context) {
	userID, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID in token claims",
		})
		return
	}

	imagesDTO, err := service.GetMyImages(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"images": imagesDTO,
	})
}

func UpdateDescriptionImage(c *gin.Context) {
	userID, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	var imageDto dto.UpdateMyImageDescriptionDTO
	err = c.ShouldBind(&imageDto)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid image ID in parameters",
		})
		return
	}

	image, err := service.GetImage(imageDto.ID)
	if err != nil || image.UserID != userID {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Failed to get image",
		})
		return
	}

	err = service.UpdateDescriptionImage(imageDto, userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Failed to update description of image",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
	})
}

func DeleteImage(c *gin.Context) {
	imageID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	userID, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID in token claims",
		})
		return
	}

	image, err := service.GetImage(imageID)
	if err != nil || image.UserID != userID {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Failed to get image",
		})
		return
	}

	err = service.DeleteImage(imageID, userID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "Failed to delete image",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Image removed",
	})
}
