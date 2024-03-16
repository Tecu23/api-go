package data

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Exercises ExerciseModel
}

func NewModels(db *gorm.DB) Models {
	return Models{
		Exercises: ExerciseModel{DB: db},
	}
}
