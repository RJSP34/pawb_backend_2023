package dto

import "time"

type UserUpdateDTO struct {
	ID             uint64 `json:"id" form:"id"`
	Name           string `json:"name" form:"name" binding:"required"`
	Email          string `json:"email" form:"email" binding:"required,email"`
	Password       string `json:"password" form:"password,omitempty"`
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
}

type UserResponseDTO struct {
	ID    uint64 `json:"id" form:"id"`
	Name  string `json:"name" form:"name" binding:"required"`
	Email string `json:"email" form:"email" binding:"required,email"`
}

type ProfileDTO struct {
	Name      string    `json:"name" form:"name" binding:"required"`
	Email     string    `json:"email" form:"email" binding:"required,email"`
	Role      string    `json:"role_description" form:"role_description"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}

type ClinicianDTO struct {
	ID             uint64 `json:"id" form:"id"`
	Name           string `json:"name" form:"name" binding:"required"`
	Email          string `json:"email" form:"email" binding:"required,email"`
	ProfilePicture string `json:"profile_picture" form:"profile_picture"`
}
