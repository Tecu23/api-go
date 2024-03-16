// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Exercise struct {
	ID                    pgtype.UUID        `db:"id" json:"id"`
	Name                  string             `db:"name" json:"name"`
	Description           string             `db:"description" json:"description"`
	VideoUrl              pgtype.Text        `db:"video_url" json:"video_url"`
	Notes                 pgtype.Text        `db:"notes" json:"notes"`
	PrimaryMuscleGroup    pgtype.Text        `db:"primary_muscle_group" json:"primary_muscle_group"`
	SecondaryMuscleGroups []string           `db:"secondary_muscle_groups" json:"secondary_muscle_groups"`
	CreatedAt             pgtype.Timestamptz `db:"created_at" json:"created_at"`
	UpdatedAt             pgtype.Timestamptz `db:"updated_at" json:"updated_at"`
}
