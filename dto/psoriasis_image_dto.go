package dto

import "time"

type PsoriasisImageSubmitDTO struct {
	BodyPartID  uint   `json:"body_part_id" binding:"required"`
	Description string `json:"description" binding:"required"`
	Image       struct {
		Mime string `json:"mime" binding:"required"`
		Data string `json:"data" binding:"required"`
	} `json:"image" binding:"required"`
}

type GetPsoriasisImageDTO struct {
	ID          uint64 `json:"id"`
	UserID      uint64 `json:"user_id"`
	BodyPartID  uint   `json:"body_part_id"`
	BodyPart    string `json:"body_part"`
	Description string `json:"description"`
	Image       struct {
		Mime string `json:"mime"`
		Data string `json:"data"`
	} `json:"image"`
	CreatedAt time.Time `json:"created_at"`
}

type GetMyImageDTO struct {
	ID          uint64 `json:"id"`
	BodyPartID  uint   `json:"body_part_id"`
	BodyPart    string `json:"body_part"`
	Description string `json:"description"`
	Image       struct {
		Mime string `json:"mime"`
		Data string `json:"data"`
	} `json:"image"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateMyImageDescriptionDTO struct {
	ID          uint64 `json:"id" form:"id"`
	Description string `json:"description" form:"description"`
}

type GetMyImagesCliniciansDTO struct {
	ID          uint64 `json:"id"`
	BodyPartID  uint   `json:"body_part_id"`
	BodyPart    string `json:"body_part"`
	Description string `json:"description"`
	Image       struct {
		Mime string `json:"mime"`
		Data string `json:"data"`
	} `json:"image"`
	CreatedAt    time.Time `json:"created_at"`
	PatientID    uint64    `json:"patient_id" form:"patient_id"`
	PatientName  string    `json:"patient_name" form:"patient_name"`
	PatientEmail string    `json:"patient_email" form:"patient_email"`
}
