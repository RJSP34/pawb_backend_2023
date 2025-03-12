package entity

import (
	"gorm.io/gorm"
	"time"
)

type PatientClinician struct {
	PatientID   uint64 `gorm:"primaryKey:auto_increment:false" json:"patient_id"`
	Patient     User   `gorm:"foreignKey:PatientID" json:"patient"`
	ClinicianID uint64 `gorm:"primaryKey:auto_increment:false" json:"clinician_id"`
	Clinician   User   `gorm:"foreignKey:ClinicianID" json:"clinician"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
