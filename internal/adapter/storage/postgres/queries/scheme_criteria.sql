-- db/query/scheme_criteria.sql

-- name: GetSchemeCriteria :many
-- Used for getting all criteria for a scheme
SELECT * FROM scheme_criteria
WHERE scheme_id = $1 AND deleted_at IS NULL;

-- name: GetSchemeCriteriaByID :one
-- Used for getting scheme criteria by ID
SELECT * FROM scheme_criteria
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: CreateSchemeCriteria :one
-- Used when creating a scheme with criteria
INSERT INTO scheme_criteria (
    id,
    created_at,
    name,
    value,
    scheme_id
) VALUES (
             gen_random_uuid(), now(), $1, $2, $3
         )
RETURNING *;

-- name: ListSchemeCriteria :many
-- Used to get a list of all scheme criteria
SELECT * FROM scheme_criteria
WHERE deleted_at is NULL;

-- name: UpdateSchemeCriteria :one
-- Used when updating scheme criteria
UPDATE scheme_criteria
SET
    name = $2,
    value = $3
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteSchemeCriteria :exec
-- Used when deleting scheme criteria
UPDATE scheme_criteria
SET
    deleted_at = now()
WHERE id = $1 AND deleted_at IS NULL;