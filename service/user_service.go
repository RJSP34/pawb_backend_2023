package service

import (
	"errors"
	"github.com/mashingan/smapping"
	"log"
	"pawb_backend/dto"
	"pawb_backend/entity"
	"pawb_backend/repository"
)

var failMapResponseStr = "failed to map to response "

func GetAllUsers() []dto.UserResponseDTO {
	users := repository.GetAllUsers()
	userDTOlist := make([]dto.UserResponseDTO, 0, len(users))

	for _, user := range users {
		userDTO := dto.UserResponseDTO{}
		err := smapping.FillStruct(&userDTO, smapping.MapFields(&user))
		if err != nil {
			return nil
		}
		userDTOlist = append(userDTOlist, userDTO)
	}

	return userDTOlist
}

func Profile(userID uint64) (dto.ProfileDTO, error) {
	user := repository.GetUser(userID)
	profileDTO := dto.ProfileDTO{
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role.Description,
		CreatedAt: user.CreatedAt,
	}
	return profileDTO, nil
}

func UpdateProfile(user entity.User) (dto.ProfileDTO, error) {
	user, err := repository.UpdateProfile(user)
	profileDTO := dto.ProfileDTO{}
	if err == nil {
		err = smapping.FillStruct(&profileDTO, smapping.MapFields(&user))
		if err != nil {
			return profileDTO, err
		}
		return profileDTO, nil
	}
	return profileDTO, errors.New("user does not exist")
}

func DeleteAccount(userID uint64) error {
	if err := repository.DeleteAccount(userID); err == nil {
		return nil
	}
	return errors.New("user does not exist")
}

func IsAllowed(userID uint64, loggedUserID uint64) bool {
	return userID == loggedUserID
}

func UpdateCliniciansPermissions(userID uint64, patientCliniciansDTO []dto.PatientClinicianDTO) error {
	requestedList := ConvertDTOToEntities(userID, patientCliniciansDTO)
	currentList, err := repository.GetAllowedClinicians(userID)
	if err != nil {
		return err
	}

	toInsertList := GetItemsToInsert(requestedList, currentList)
	toDeleteList := GetItemsToDelete(requestedList, currentList)

	if err := repository.UpdateCliniciansPermissions(toInsertList, toDeleteList); err != nil {
		return err
	}

	return nil
}

func ConvertDTOToEntities(userID uint64, patientClinicianDtos []dto.PatientClinicianDTO) []entity.PatientClinician {
	var entities []entity.PatientClinician
	for _, patientClinicianDto := range patientClinicianDtos {
		patientClinician := entity.PatientClinician{}
		if err := smapping.FillStruct(&patientClinician, smapping.MapFields(&patientClinicianDto)); err != nil {
			log.Fatal(failMapResponseStr, err)
		}
		patientClinician.PatientID = userID
		entities = append(entities, patientClinician)
	}
	return entities
}

func GetItemsToInsert(requestedList, currentList []entity.PatientClinician) []entity.PatientClinician {
	var toInsertList []entity.PatientClinician
	for _, requestedClinician := range requestedList {
		found := false
		for _, currentClinician := range currentList {
			if requestedClinician.ClinicianID == currentClinician.ClinicianID {
				found = true
				break
			}
		}
		if !found {
			toInsertList = append(toInsertList, requestedClinician)
		}
	}
	return toInsertList
}

func GetItemsToDelete(requestedList, currentList []entity.PatientClinician) []entity.PatientClinician {
	var toDeleteList []entity.PatientClinician
	for _, currentClinician := range currentList {
		found := false
		for _, requestedClinician := range requestedList {
			if requestedClinician.ClinicianID == currentClinician.ClinicianID {
				found = true
				break
			}
		}
		if !found {
			toDeleteList = append(toDeleteList, currentClinician)
		}
	}
	return toDeleteList
}

func GetAllowedClinicians(userID uint64) ([]dto.ClinicianDTO, error) {
	clinicians, err := repository.GetAllowedClinicians(userID)
	if err != nil {
		return nil, err
	}

	cliniciansDTOlist := make([]dto.ClinicianDTO, 0, len(clinicians))

	for _, clinician := range clinicians {
		clinicianDTO := dto.ClinicianDTO{}
		userclinician := repository.GetUser(clinician.ClinicianID)
		err := smapping.FillStruct(&clinicianDTO, smapping.MapFields(&userclinician))
		if err != nil {
			return nil, err
		}
		cliniciansDTOlist = append(cliniciansDTOlist, clinicianDTO)
	}

	return cliniciansDTOlist, nil
}
