package service

import (
	"errors"
	"github.com/mashingan/smapping"
	"pawb_backend/dto"
	"pawb_backend/entity"
	"pawb_backend/repository"
)

func GetFeedbackbyClinician(userID uint64) ([]dto.ClinicianFeedbackDTO, error) {
	feedbackList, err := repository.GetFeedbackByClinician(userID)

	if err != nil {
		return nil, err
	}

	var feedbackDTOList []dto.ClinicianFeedbackDTO

	for _, feedback := range feedbackList {
		feedbackDTO := dto.ClinicianFeedbackDTO{}
		err := smapping.FillStruct(&feedbackDTO, smapping.MapFields(&feedback))
		if err != nil {
			return nil, err
		}
		feedbackDTOList = append(feedbackDTOList, feedbackDTO)
	}
	return feedbackDTOList, err
}

func GetFeedbackByImage(ImageId uint64, userId uint64) ([]dto.FeedbackResponseDTO, error) {

	user := repository.GetUser(userId)
	var feedbacklist []entity.Feedback
	feedbacklist, err := repository.GetFeedbackByImageAsUser(ImageId, userId, user.Role.Description)
	if err != nil {
		return nil, err
	}

	var feedbackListDTO []dto.FeedbackResponseDTO
	for _, feedback := range feedbacklist {

		feedbackDTO := dto.FeedbackResponseDTO{
			ClinicianID:   feedback.ClinicianID,
			ClinicianName: feedback.Clinician.Name,
			Feedback:      feedback.Feedback,
			ID:            feedback.ID,
		}

		feedbackListDTO = append(feedbackListDTO, feedbackDTO)
	}
	return feedbackListDTO, nil
}

func SubmitFeedback(feedback dto.ClinicianFeedbackSubmissionDTO, clinicianID uint64) error {
	ImageDTO, err := GetImage(feedback.ImageID)
	if err != nil {
		return err
	}
	Authopatients, err := GetAllAuthorizedPatients(clinicianID)

	for _, Authpant := range Authopatients {
		if Authpant.PatientID == ImageDTO.UserID {
			err = repository.SubmitFeedback(feedback, clinicianID)
			return err
		}
	}
	return errors.New("not ALLOWED")
}
