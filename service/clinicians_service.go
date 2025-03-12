package service

import (
	"errors"
	"fmt"
	"github.com/mashingan/smapping"
	"pawb_backend/dto"
	"pawb_backend/repository"
)

func GetClinicians() ([]dto.ClinicianDTO, error) {
	clinicians := repository.GetClinicians()
	cliniciansDTOlist := make([]dto.ClinicianDTO, 0, len(clinicians))

	for _, clinician := range clinicians {
		clinicianDTO := dto.ClinicianDTO{}
		err := smapping.FillStruct(&clinicianDTO, smapping.MapFields(&clinician))
		if err != nil {
			return nil, err
		}
		cliniciansDTOlist = append(cliniciansDTOlist, clinicianDTO)
	}

	return cliniciansDTOlist, nil
}

func GetAllPatientsImages(patientID uint64, clinicianId uint64) ([]dto.GetMyImageDTO, error) {
	authPatients, err := GetAllAuthorizedPatients(clinicianId)

	if err != nil {
		return nil, err
	}

	for _, authPatient := range authPatients {
		if authPatient.PatientID == patientID {
			return GetMyImages(patientID)
		}
	}
	return nil, errors.New("patient didnt authorize clinician")
}

func GetAllAuthorizedPatients(clinicianId uint64) ([]dto.AllowedPatientsClinicianDTO, error) {
	patientClinicians := repository.GetAllAuthorizedPatients(clinicianId)

	patientsCliniciansDTOList := make([]dto.AllowedPatientsClinicianDTO, 0, len(patientClinicians))

	for _, patientClinician := range patientClinicians {
		patientClinicianDTO := dto.AllowedPatientsClinicianDTO{}
		err := smapping.FillStruct(&patientClinicianDTO, smapping.MapFields(&patientClinician))
		patientClinicianDTO.PatientName = patientClinician.Patient.Name
		patientClinicianDTO.PatientEmail = patientClinician.Patient.Email
		if err != nil {
			return nil, err
		}
		patientsCliniciansDTOList = append(patientsCliniciansDTOList, patientClinicianDTO)
	}

	return patientsCliniciansDTOList, nil
}

func GetImageByClinician(ImageID uint64, clinicianID uint64) (dto.GetPsoriasisImageDTO, error) {
	ImageDTO, err := GetImage(ImageID)
	if err != nil {
		return dto.GetPsoriasisImageDTO{}, err
	}
	authPatients, err := GetAllAuthorizedPatients(clinicianID)

	for _, authPatient := range authPatients {
		if authPatient.PatientID == ImageDTO.UserID {
			return ImageDTO, nil
		}
	}

	return dto.GetPsoriasisImageDTO{}, errors.New("patient didn't authorize it")

}

func GetAllImagesByClinician(clinicianID uint64) ([]dto.GetMyImagesCliniciansDTO, error) {
	authPatients, err := GetAllAuthorizedPatients(clinicianID)
	if err != nil {
		return []dto.GetMyImagesCliniciansDTO{}, err
	}
	var ImagesDTO []dto.GetMyImagesCliniciansDTO

	for _, authPatient := range authPatients {
		ImagesPatients, err := GetMyImagesClinicians(authPatient.PatientID)
		if err != nil {
			return []dto.GetMyImagesCliniciansDTO{}, err
		}

		ImagesDTO = append(ImagesDTO, ImagesPatients...)
	}

	return ImagesDTO, nil
}

func UpdateFeedback(clinicianID uint64, feedbackDTO dto.EditFeedbackDTO) error {
	feedback, err := repository.GetFeedbackByID(feedbackDTO.ID)
	if err != nil {
		return err
	}

	if feedback.ClinicianID != clinicianID {
		return fmt.Errorf("not authorized")
	}

	feedback.Feedback = feedbackDTO.Feedback

	err = repository.UpdateFeedback(feedback)

	if err != nil {
		return err
	}

	return nil
}

func RemoveFeedback(feedbackID uint64, clinicianID uint64) error {
	feedback, err := repository.GetFeedbackByID(feedbackID)
	if err != nil {
		return err
	}

	if feedback.ClinicianID != clinicianID {
		return fmt.Errorf("not authorized")
	}

	err = repository.DeleteFeedback(feedback.ID)

	if err != nil {
		return err
	}

	return nil
}
