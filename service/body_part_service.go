package service

import (
	"github.com/mashingan/smapping"
	"pawb_backend/dto"
	"pawb_backend/repository"
)

func GetBodyParts() ([]dto.BodyPartResponseDTO, error) {
	bodyParts := repository.GetAllBodyParts()
	bodyPartsDTOlist := make([]dto.BodyPartResponseDTO, 0, len(bodyParts))

	for _, bodypart := range bodyParts {
		bodyPartDTO := dto.BodyPartResponseDTO{}
		err := smapping.FillStruct(&bodyPartDTO, smapping.MapFields(&bodypart))
		if err != nil {
			return nil, err
		}
		bodyPartsDTOlist = append(bodyPartsDTOlist, bodyPartDTO)
	}

	return bodyPartsDTOlist, nil
}
