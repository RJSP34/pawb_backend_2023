package entity

import (
	"gorm.io/gorm"
	"time"
)

type Role struct {
	ID          uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	Users       []User `gorm:"foreignKey:RoleID" json:"users"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
