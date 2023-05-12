-- name: CreateEmployee :one
INSERT INTO employees (
  identity_id, code, full_name
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: ListEmployees :many
SELECT * FROM employees
ORDER BY code
LIMIT $1
OFFSET $2;