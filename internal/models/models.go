// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Exercise struct {
	ID   pgtype.UUID `db:"id" json:"id"`
	Name string      `db:"name" json:"name"`
}