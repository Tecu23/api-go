// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: exercises.sql

package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createExercise = `-- name: CreateExercise :one
INSERT INTO exercises (
  name
) VALUES (
  $1
)
RETURNING id, name, description, video_url, notes, primary_muscle_group, secondary_muscle_groups, created_at, updated_at
`

func (q *Queries) CreateExercise(ctx context.Context, name string) (Exercise, error) {
	row := q.db.QueryRow(ctx, createExercise, name)
	var i Exercise
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.VideoUrl,
		&i.Notes,
		&i.PrimaryMuscleGroup,
		&i.SecondaryMuscleGroups,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteExercise = `-- name: DeleteExercise :exec
DELETE FROM exercises
WHERE id = $1
`

func (q *Queries) DeleteExercise(ctx context.Context, id pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteExercise, id)
	return err
}

const getExercise = `-- name: GetExercise :one
SELECT id, name, description, video_url, notes, primary_muscle_group, secondary_muscle_groups, created_at, updated_at FROM exercises
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetExercise(ctx context.Context, id pgtype.UUID) (Exercise, error) {
	row := q.db.QueryRow(ctx, getExercise, id)
	var i Exercise
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.VideoUrl,
		&i.Notes,
		&i.PrimaryMuscleGroup,
		&i.SecondaryMuscleGroups,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listExercises = `-- name: ListExercises :many
SELECT id, name, description, video_url, notes, primary_muscle_group, secondary_muscle_groups, created_at, updated_at FROM exercises
ORDER BY name
`

func (q *Queries) ListExercises(ctx context.Context) ([]Exercise, error) {
	rows, err := q.db.Query(ctx, listExercises)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Exercise
	for rows.Next() {
		var i Exercise
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.VideoUrl,
			&i.Notes,
			&i.PrimaryMuscleGroup,
			&i.SecondaryMuscleGroups,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateExercise = `-- name: UpdateExercise :exec
UPDATE exercises
  set name = $2
WHERE id = $1
`

type UpdateExerciseParams struct {
	ID   pgtype.UUID `db:"id" json:"id"`
	Name string      `db:"name" json:"name"`
}

func (q *Queries) UpdateExercise(ctx context.Context, arg UpdateExerciseParams) error {
	_, err := q.db.Exec(ctx, updateExercise, arg.ID, arg.Name)
	return err
}
