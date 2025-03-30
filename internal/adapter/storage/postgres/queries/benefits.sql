-- db/query/benefits.sql

-- name: GetBenefitsByScheme :many
-- Used for getting benefits for a scheme
SELECT * FROM benefits
WHERE scheme_id = $1 AND deleted_at IS NULL
ORDER BY created_at DESC;

-- name: GetBenefitByID :one
-- Used for getting benefits by id
SELECT * FROM benefits
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: CreateBenefit :one
-- Used when creating a scheme with benefits
INSERT INTO benefits (
    id,
    created_at,
    scheme_id,
    name,
    amount
) VALUES (
            gen_random_uuid(), now(), $1, $2, $3
         )
RETURNING *;

-- name: UpdateBenefit :one
-- Used when updating scheme benefits
UPDATE benefits
SET
    name = $2,
    amount = $3
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteBenefit :exec
-- Used when deleting scheme benefits
UPDATE benefits
SET
    deleted_at = now()
WHERE id = $1 AND deleted_at IS NULL;