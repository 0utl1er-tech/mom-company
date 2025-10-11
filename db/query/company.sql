-- name: CreateCompany :one
INSERT INTO company (id, ceo, trademark, type, position, address, company_code, contact_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetCompany :one
SELECT 
  c.id, c.ceo, c.trademark, c.type, c.position, c.address, c.company_code, c.contact_id, c.created_at,
  co.email as contact_email, co.phone as contact_phone
FROM company c
LEFT JOIN contact co ON c.contact_id = co.id
WHERE c.id = $1;

-- name: ListCompanies :many
SELECT 
  c.id, c.ceo, c.trademark, c.type, c.position, c.address, c.company_code, c.contact_id, c.created_at,
  co.email as contact_email, co.phone as contact_phone
FROM company c
LEFT JOIN contact co ON c.contact_id = co.id
ORDER BY c.created_at DESC;

-- name: UpdateCompany :one
UPDATE company 
SET 
  trademark = $2,
  type = $3,
  position = $4, 
  address = $5, 
  company_code = $6
WHERE id = $1
RETURNING *;

-- name: UpdateCompanyCeo :one
UPDATE company 
SET ceo = $2
WHERE id = $1
RETURNING *;

-- name: DeleteCompany :exec
DELETE FROM company
WHERE id = $1;