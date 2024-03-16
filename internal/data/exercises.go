package data

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Exercise struct {
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	name string
}

type ExerciseModel struct {
	DB *gorm.DB
}
