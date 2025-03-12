package dto

type PatientClinicianDTO struct {
	ClinicianID uint64 `json:"clinician_id" form:"clinician_id"`
}

type AllowedClinicianDTO struct {
	ClinicianID uint64 `json:"clinician_id" form:"clinician_id"`
}

type AllowedPatientsClinicianDTO struct {
	ClinicianID  uint64 `json:"clinician_id" form:"clinician_id"`
	PatientID    uint64 `json:"patient_id" form:"patient_id"`
	PatientName  string `json:"patient_name" form:"patient_name"`
	PatientEmail string `json:"patient_email" form:"patient_email"`
}
