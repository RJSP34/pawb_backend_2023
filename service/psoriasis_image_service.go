package service

import (
	"errors"
	"pawb_backend/dto"
	"pawb_backend/entity"
	"pawb_backend/repository"
)

func SubmitImage(psoriasisImage *entity.PsoriasisImage) error {
	user := repository.GetUser(psoriasisImage.UserID)
	encryptedImageData, err := Encrypt(psoriasisImage.ImageData, user.PrivateKey)
	if err != nil {
		return err
	}
	psoriasisImage.ImageData = encryptedImageData
	err = repository.PsoriasisImageSubmit(*psoriasisImage)
	if err != nil {
		return err
	}
	return nil
}

func GetImage(imageID uint64) (dto.GetPsoriasisImageDTO, error) {
	image, err := repository.GetImage(imageID)
	if err != nil {
		return dto.GetPsoriasisImageDTO{}, err
	}
	user := repository.GetUser(image.UserID)
	imageBlob, err := Decrypt(image.ImageData, user.PrivateKey)
	if err != nil {
		return dto.GetPsoriasisImageDTO{}, err
	}

	imageDTO := dto.GetPsoriasisImageDTO{
		ID:          image.ID,
		UserID:      image.UserID,
		BodyPartID:  image.BodyPartID,
		BodyPart:    image.BodyPart.Name,
		Description: image.Description,
		Image: struct {
			Mime string `json:"mime"`
			Data string `json:"data"`
		}{
			Mime: image.Mime,
			Data: EncodeToBase64(imageBlob),
		},
		CreatedAt: image.CreatedAt,
	}

	return imageDTO, nil
}

func GetMyImages(userID uint64) ([]dto.GetMyImageDTO, error) {
	images, err := repository.GetMyImages(userID)
	if err != nil {
		return []dto.GetMyImageDTO{}, err
	}

	var imagesDTO []dto.GetMyImageDTO
	user := repository.GetUser(userID)
	for _, image := range images {
		imageBlob, err := Decrypt(image.ImageData, user.PrivateKey)
		if err != nil {
			return []dto.GetMyImageDTO{}, errors.New("problem occurred decrypting image")
		}

		imageDTO := dto.GetMyImageDTO{
			ID:          image.ID,
			BodyPartID:  image.BodyPartID,
			BodyPart:    image.BodyPart.Name,
			Description: image.Description,
			Image: struct {
				Mime string `json:"mime"`
				Data string `json:"data"`
			}{
				Mime: image.Mime,
				Data: EncodeToBase64(imageBlob),
			},
			CreatedAt: image.CreatedAt,
		}

		imagesDTO = append(imagesDTO, imageDTO)
	}

	return imagesDTO, nil
}

func UpdateDescriptionImage(ImageDto dto.UpdateMyImageDescriptionDTO, UserID uint64) error {
	image, err := repository.GetImage(ImageDto.ID)
	if err != nil {
		return err
	}
	if image.UserID != UserID {
		return errors.New("failed to get image")
	}

	image.Description = ImageDto.Description
	return repository.UpdateDescriptionImage(image)
}

func DeleteImage(ImageID uint64, UserID uint64) error {
	image, err := repository.GetImage(ImageID)
	if err != nil {
		return err
	}
	if image.UserID != UserID {
		return errors.New("failed to get image")
	}

	return repository.DeleteImage(ImageID)
}

func GetMyImagesClinicians(userID uint64) ([]dto.GetMyImagesCliniciansDTO, error) {
	images, err := repository.GetMyImages(userID)
	if err != nil {
		return []dto.GetMyImagesCliniciansDTO{}, err
	}

	var imagesDTO []dto.GetMyImagesCliniciansDTO
	user := repository.GetUser(userID)
	for _, image := range images {
		imageBlob, err := Decrypt(image.ImageData, user.PrivateKey)
		if err != nil {
			return []dto.GetMyImagesCliniciansDTO{}, errors.New("problem occurred decrypting image")
		}

		imageDTO := dto.GetMyImagesCliniciansDTO{
			ID:          image.ID,
			BodyPartID:  image.BodyPartID,
			BodyPart:    image.BodyPart.Name,
			Description: image.Description,
			Image: struct {
				Mime string `json:"mime"`
				Data string `json:"data"`
			}{
				Mime: image.Mime,
				Data: EncodeToBase64(imageBlob),
			},
			PatientID:    userID,
			PatientEmail: user.Email,
			PatientName:  user.Name,
			CreatedAt:    image.CreatedAt,
		}

		imagesDTO = append(imagesDTO, imageDTO)
	}

	return imagesDTO, nil
}
