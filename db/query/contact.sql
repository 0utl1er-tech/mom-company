-- name: CreateContact :one
INSERT INTO contact (id, email, phone)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateContact :one
UPDATE contact
SET email = $2, phone = $3
WHERE id = $1
RETURNING *;

-- name: DeleteContact :exec
DELETE FROM contact
WHERE id = $1;

-- name: GetContact :one
SELECT * FROM contact
WHERE id = $1;

-- name: ListContact :many
SELECT * FROM contact
WHERE id = $1;