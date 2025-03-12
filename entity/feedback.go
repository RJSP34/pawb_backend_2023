package entity

import (
	"gorm.io/gorm"
	"time"
)

type Feedback struct {
	ID          uint64         `gorm:"primaryKey:auto_increment:true" json:"id"`
	ClinicianID uint64         `gorm:"" json:"clinician_id"`
	Clinician   User           `gorm:"foreignKey:ClinicianID" json:"clinician"`
	ImageID     uint64         `gorm:"" json:"image_id"`
	Image       PsoriasisImage `gorm:"foreignKey:ImageID" json:"image"`
	Feedback    string         `gorm:"type:varchar(255)" json:"feedback"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
