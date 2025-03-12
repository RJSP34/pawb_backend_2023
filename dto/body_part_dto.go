package dto

type BodyPartResponseDTO struct {
	ID   uint   `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}
