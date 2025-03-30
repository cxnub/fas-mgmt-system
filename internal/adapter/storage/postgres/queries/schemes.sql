-- db/query/schemes.sql

-- name: GetScheme :one
-- Used for GET /api/schemes/{id}
SELECT * FROM schemes
WHERE id = $1 AND deleted_at IS NULL;

-- name: ListSchemes :many
-- Used for GET /api/schemes
SELECT * FROM schemes
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: CreateScheme :one
-- Used for POST /api/schemes
INSERT INTO schemes (
    id,
    created_at,
    name
) VALUES (
            gen_random_uuid(), now(), $1
         )
RETURNING *;

-- name: UpdateScheme :one
-- Used for PUT /api/schemes/{id}
UPDATE schemes
SET
    name = $2
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteScheme :exec
-- Used for DELETE /api/schemes/{id}
UPDATE schemes
SET
    deleted_at = now()
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetSchemeWithCriteriaAndBenefits :many
-- Used for getting a scheme with its criteria
SELECT
    s.*,
    sc.id as criteria_id,
    sc.name as criteria_name,
    sc.value as criteria_value
FROM schemes s
         LEFT JOIN scheme_criteria sc ON s.id = sc.scheme_id AND sc.deleted_at IS NULL
WHERE s.id = $1 AND s.deleted_at IS NULL;

-- name: GetSchemeWithBenefits :many
-- Used for getting a scheme with its benefits
SELECT
    s.*,
    b.id as benefit_id,
    b.name as benefit_name,
    b.amount as benefit_amount
FROM schemes s
         LEFT JOIN benefits b ON s.id = b.scheme_id AND b.deleted_at IS NULL
WHERE s.id = $1 AND s.deleted_at IS NULL;