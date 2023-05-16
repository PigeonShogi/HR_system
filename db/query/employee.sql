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

-- name: GetEmployeeForUpdate :one
SELECT * FROM employees
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListEmployees :many
SELECT * FROM employees
ORDER BY code ASC
LIMIT $1
OFFSET $2;

-- name: UpdateEmployeeWithStock :one
UPDATE employees
SET stock = $2
WHERE id = $1
RETURNING *;

-- name: AddEmployeeStock :one
UPDATE employees
SET stock = stock + sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;