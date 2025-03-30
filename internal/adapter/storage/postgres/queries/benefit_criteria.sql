-- name: GetBenefitCriteriaByID :one
SELECT *
FROM benefit_criteria
WHERE id = $1
LIMIT 1;

-- name: GetBenefitCriteriaByBenefitID :many
SELECT *
FROM benefit_criteria
WHERE benefit_id = $1;

-- name: CreateBenefitCriteria :exec
INSERT INTO benefit_criteria (id, created_at, name, value, benefit_id)
VALUES (gen_random_uuid(), now(), $1, $2, $3);

-- name: UpdateBenefitCriteria :exec
UPDATE benefit_criteria
SET name       = $1,
    value      = $2
WHERE id = $3;

-- name: DeleteBenefitCriteria :exec
UPDATE benefit_criteria
SET
    deleted_at = now()
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetAllBenefitCriteria :many
SELECT *
FROM benefit_criteria
WHERE deleted_at IS NULL;