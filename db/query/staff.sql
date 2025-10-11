-- name: CreateStaff :one
INSERT INTO staff (id, name, role, contact_id, company_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateStaff :one
UPDATE staff
SET name = $2, role = $3, contact_id = $4
WHERE id = $1
RETURNING *;

-- name: UpdateStaffCompany :one
UPDATE staff
SET company_id = $2
WHERE id = $1
RETURNING *;

-- name: DeleteStaff :exec
DELETE FROM staff
WHERE id = $1;

-- name: GetStaff :one
SELECT * FROM staff
WHERE id = $1;

-- name: ListStaff :many
SELECT 
  s.id, s.name, s.role, s.contact_id, s.company_id, s.created_at,
  c.email as contact_email, c.phone as contact_phone
FROM staff s
LEFT JOIN contact c ON s.contact_id = c.id
WHERE s.company_id = $1;