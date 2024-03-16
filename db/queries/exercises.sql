-- name: GetExercise :one
SELECT * FROM exercises
WHERE id = $1 LIMIT 1;

-- name: ListExercises :many
SELECT * FROM exercises
ORDER BY name;

-- name: CreateExercise :one
INSERT INTO exercises (
  name
) VALUES (
  $1
)
RETURNING *;

-- name: UpdateExercise :exec
UPDATE exercises
  set name = $2
WHERE id = $1;

-- name: DeleteExercise :exec
DELETE FROM exercises
WHERE id = $1;
