package repository

import (
	"pawb_backend/config"
	"pawb_backend/dto"
	"pawb_backend/entity"
)

func IsEmailRegistered(email string) bool {
	var user entity.User
	err := config.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return false
	}
	return user.ID > 0
}

func Register(user entity.User) (entity.User, error) {
	var role entity.Role
	err := config.Db.Where("description = ?", "user").First(&role).Error
	if err != nil {
		return entity.User{}, err
	}
	user.Role = role
	config.Db.Save(&user)
	return user, nil
}

func Login(loginDTO dto.LoginDTO) (entity.User, error) {
	var user entity.User
	err := config.Db.Preload("Role").Select("id, password, role_id").Where("email = ?", loginDTO.Email).First(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
