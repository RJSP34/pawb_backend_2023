package repository

import (
	"errors"
	"pawb_backend/config"
	"pawb_backend/entity"
)

func PsoriasisImageSubmit(psoriasisImage entity.PsoriasisImage) error {
	result := config.Db.Save(&psoriasisImage)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}

func GetImage(imageID uint64) (entity.PsoriasisImage, error) {
	var image entity.PsoriasisImage
	err := config.Db.Preload("BodyPart").First(&image, imageID).Error
	if err != nil || image.ID == 0 {
		return image, errors.New("image does not exist")
	}
	return image, nil
}

func GetMyImages(userID uint64) ([]entity.PsoriasisImage, error) {
	var images []entity.PsoriasisImage
	if err := config.Db.Preload("BodyPart").Where("user_id = ?", userID).Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}

func UpdateDescriptionImage(Image entity.PsoriasisImage) error {
	if err := config.Db.Save(&Image).Error; err != nil {
		return err
	}
	return nil
}

func DeleteImage(ImageID uint64) error {
	var image entity.PsoriasisImage
	config.Db.First(&image, ImageID)
	if image.ID != 0 {
		config.Db.Delete(&image)
		return nil
	}
	return errors.New("image does not exist")
}
