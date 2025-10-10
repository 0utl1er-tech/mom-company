-- name: CreateCompany :one
INSERT INTO company (id, ceo, trademark, type, position, address, company_code)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetCompany :one
SELECT * FROM company
WHERE id = $1;

-- name: CreateCompany :one
INSERT INTO company (id, ceo, trademark, type, position, address, company_code)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

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

-- name: DeleteCompany :exec
DELETE FROM company
WHERE id = $1;