-- name: GetEmployeeFromIdentities :one
SELECT * FROM identities
WHERE name = 'HR-Admin' LIMIT 1;

-- name: GetHrAdminFromIdentities :one
SELECT * FROM identities
WHERE name = 'HR-Admin' LIMIT 1;