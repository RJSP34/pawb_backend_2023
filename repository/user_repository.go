package repository

import (
	"errors"
	"pawb_backend/config"
	"pawb_backend/entity"
)

func GetUser(userID uint64) entity.User {
	var user entity.User
	config.Db.Preload("Role").First(&user, userID)
	return user
}

func GetAllUsers() []entity.User {
	var users []entity.User
	config.Db.Find(&users)
	return users
}

func Profile(userID uint64) (entity.User, error) {
	var user entity.User
	config.Db.First(&user, userID)
	if user.ID != 0 {
		return user, nil
	}
	return user, errors.New("user does not exist")
}

func UpdateProfile(user entity.User) (entity.User, error) {
	if _, err := Profile(user.ID); err == nil {
		config.Db.Save(&user)
		config.Db.Find(&user)
		return user, nil
	}
	return user, errors.New("user does not exist")
}

func DeleteAccount(userID uint64) error {
	var user entity.User
	config.Db.First(&user, userID)
	if user.ID != 0 {
		config.Db.Delete(&user)
		return nil
	}
	return errors.New("user does not exist")
}

func UpdateCliniciansPermissions(insertPatientsClinicians []entity.PatientClinician, deletePatientsClinicians []entity.PatientClinician) error {
	if len(deletePatientsClinicians) > 0 {
		err := config.Db.Delete(deletePatientsClinicians).Error
		if err != nil {
			return errors.New("error deleting from DB")
		}
	}

	if len(insertPatientsClinicians) > 0 {
		err := config.Db.Save(insertPatientsClinicians).Error
		if err != nil {
			return errors.New("error inserting to DB")
		}
	}
	return nil
}

func GetAllowedClinicians(userID uint64) ([]entity.PatientClinician, error) {
	var patientClinicians []entity.PatientClinician
	err := config.Db.Where("patient_id=?", userID).Find(&patientClinicians).Error
	if err != nil {
		return nil, err
	}
	return patientClinicians, nil
}
