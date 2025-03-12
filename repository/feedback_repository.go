package repository

import (
	"errors"
	"pawb_backend/config"
	"pawb_backend/dto"
	"pawb_backend/entity"
)

func GetFeedbackByClinician(userID uint64) ([]entity.Feedback, error) {
	var feedback []entity.Feedback
	err := config.Db.Where("clinician_id=? AND image_id IN (SELECT id FROM psoriasis_images WHERE deleted_at IS NULL)", userID).Find(&feedback).Error
	if err != nil {
		return nil, err
	}
	return feedback, nil
}

func GetFeedbackByImageAsUser(ImageId uint64, userId uint64, role string) ([]entity.Feedback, error) {
	imageEntity, err := GetImage(ImageId)
	if err != nil {
		return nil, err
	}
	if role == "user" && imageEntity.UserID != userId {
		return nil, errors.New("user does not have permission")
	} else if role == "clinician" {
		var patientclinician entity.PatientClinician

		err = config.Db.Where("clinician_id=? and patient_id=?", userId, imageEntity.UserID).Find(&patientclinician).Error

		if err != nil {
			return nil, err
		}

		if patientclinician.ClinicianID <= 0 {
			return nil, errors.New("user does not have permission")
		}
	} else if role != "user" && role != "clinician" && role != "admin" {
		return nil, errors.New("role is not recognized")
	}

	var feedback []entity.Feedback
	err = config.Db.Preload("Clinician").Where("image_id=?", ImageId).Find(&feedback).Error

	if err != nil {
		return nil, err
	}

	return feedback, nil
}

func SubmitFeedback(feedback dto.ClinicianFeedbackSubmissionDTO, clinicianID uint64) error {
	Feedbacksub := entity.Feedback{
		ClinicianID: clinicianID,
		Feedback:    feedback.Feedback,
		ImageID:     feedback.ImageID,
	}
	result := config.Db.Save(&Feedbacksub)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
