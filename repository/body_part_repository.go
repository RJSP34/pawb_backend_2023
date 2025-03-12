package repository

import (
	"pawb_backend/config"
	"pawb_backend/entity"
)

func GetAllBodyParts() []entity.BodyPart {
	var BodyParts []entity.BodyPart
	config.Db.Find(&BodyParts)
	return BodyParts
}
