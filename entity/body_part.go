package entity

import (
	"gorm.io/gorm"
	"time"
)

type BodyPart struct {
	ID        uint   `gorm:"primary_key:auto_increment" json:"id"`
	Name      string `gorm:"type:varchar(255);unique;not null" json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
