-- name: CreatePunch :one
INSERT INTO punches (
  employee_id, working_day, working_hours, status_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdatePunch :one
UPDATE punches
  set working_hours = $2,
  status_id = $3
WHERE id = $1
RETURNING *;