package dto

import (
	"time"
)

type FeedbackResponseDTO struct {
	ID            uint64 `json:"feedback_id" form:"feedback_id"`
	ClinicianID   uint64 `json:"clinician_id" form:"clinician_id"`
	ClinicianName string `json:"clinician_name" form:"clinician_name"`
	Feedback      string `json:"feedback" form:"feedback"`
}

type ClinicianFeedbackDTO struct {
	ID        uint64    `json:"feedback_id" form:"feedback_id"`
	ImageID   uint64    `json:"image_id" form:"image_id"`
	Feedback  string    `json:"feedback" form:"feedback"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
}

type EditFeedbackDTO struct {
	ID       uint64 `json:"feedback_id" form:"feedback_id"`
	Feedback string `json:"feedback" form:"feedback"`
}

type ClinicianFeedbackSubmissionDTO struct {
	ImageID  uint64 `json:"image_id" form:"image_id"`
	Feedback string `json:"feedback" form:"feedback"`
}
