package controller

import (
	"net/http"
	"pawb_backend/dto"
	"pawb_backend/entity"
	"pawb_backend/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "select users",
		"users":   service.GetAllUsers(),
	})
}

func Profile(c *gin.Context) {
	userID, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid user ID in token claims",
		})
		return
	}

	profile, err := service.Profile(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get profile",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"profile": profile,
	})
}

func UpdateProfile(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	var user entity.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	user.ID = userID
	loggedUserID, _ := strconv.ParseUint(c.GetString("user_id"), 10, 64)

	if !service.IsAllowed(userID, loggedUserID) {
		c.JSON(401, gin.H{
			"message": "you do not have the permission - you are not the owner of this profile",
		})
		return
	}

	userDTO, err := service.UpdateProfile(user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "update user",
		"user":    userDTO,
	})
}

func DeleteAccount(c *gin.Context) {
	userID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	loggedUserID, _ := strconv.ParseUint(c.GetString("user_id"), 10, 64)

	if !service.IsAllowed(userID, loggedUserID) {
		c.JSON(401, gin.H{
			"message": "you do not have the permission - you are not the owner of this profile",
		})
		return
	}

	err := service.DeleteAccount(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted",
	})
}

func UpdateCliniciansPermissions(c *gin.Context) {
	userID, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	var cliniciansdto []dto.PatientClinicianDTO
	err = c.ShouldBind(&cliniciansdto)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	err = service.UpdateCliniciansPermissions(userID, cliniciansdto)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Clinicians Added",
	})
}

func GetAllowedClinicians(c *gin.Context) {
	userID, err := strconv.ParseUint(c.GetString("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}

	Clinicians, err := service.GetAllowedClinicians(userID)
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
