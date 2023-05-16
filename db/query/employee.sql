-- name: CreateEmployee :one
INSERT INTO employees (
  identity_id, code, full_name
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteEmployeeById :exec
DELETE FROM employees
WHERE id = $1;

-- name: GetEmployee :one
SELECT * FROM employees
WHERE id = $1 LIMIT 1;

-- name: ListEmployees :many
SELECT * FROM employees
ORDER BY code
LIMIT $1
OFFSET $2;

-- name: UpdateEmployeeWithStock :one
UPDATE employees
SET stock = $2
WHERE id = $1
RETURNING *;