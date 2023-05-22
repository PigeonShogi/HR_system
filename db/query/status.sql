-- name: GetStatusByName :one
SELECT * FROM statuses
WHERE name = $1 LIMIT 1;