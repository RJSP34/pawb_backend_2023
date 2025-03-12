package repository

import (
	"errors"
	"pawb_backend/config"
	"pawb_backend/entity"
)

func GetClinicians() []entity.User {
	var clinicians []entity.User
	var role entity.Role
	config.Db.Where("description=?", "clinician").First(&role)
	config.Db.Where("role_id=?", role.ID).Find(&clinicians)
	return clinicians
}

func GetAllAuthorizedPatients(clinicianID uint64) []entity.PatientClinician {
	var PatientClinicians []entity.PatientClinician
	config.Db.Preload("Patient").Where("clinician_id", clinicianID).Find(&PatientClinicians)
	return PatientClinicians
}

func GetFeedbackByID(feedbackID uint64) (entity.Feedback, error) {
	var feedback entity.Feedback
	err := config.Db.Where("id=?", feedbackID).Find(&feedback).Error
	return feedback, err
}

func UpdateFeedback(feedback entity.Feedback) error {
	if err := config.Db.Save(&feedback).Error; err != nil {
		return err
	}
	return nil
}

func DeleteFeedback(feedbackID uint64) error {
	var feedback entity.Feedback
	config.Db.First(&feedback, feedbackID)
	if feedback.ID != 0 {
		config.Db.Delete(&feedback)
		return nil
	}
	return errors.New("feedback does not exist")
}
