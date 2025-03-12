package entity

import (
	"gorm.io/gorm"
	"time"
)

type PsoriasisImage struct {
	ID          uint64   `gorm:"primary_key:auto_increment" json:"id"`
	Description string   `gorm:"type:text" json:"description"`
	BodyPart    BodyPart `gorm:"foreignKey:BodyPartID" json:"body_part"`
	BodyPartID  uint     `gorm:"index" json:"body_part_id"`
	UserID      uint64   `gorm:"index" json:"user_id"`
	User        User     `gorm:"foreignKey:UserID" json:"user"`
	ImageData   []byte   `gorm:"type:longblob" json:"image_data"`
	Mime        string   `gorm:"type:text" json:"mime"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
