-- db/query/applications.sql

-- name: GetApplication :one
-- Used for GET /api/applications/{id}
SELECT * FROM applications
WHERE id = $1 AND deleted_at IS NULL;

-- name: ListApplications :many
-- Used for GET /api/applications
SELECT * FROM applications
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: GetApplicationsByApplicant :many
-- Used for getting applications for a specific applicant
SELECT * FROM applications
WHERE applicant_id = $1 AND deleted_at IS NULL
ORDER BY created_at DESC;

-- name: CreateApplication :one
-- Used for POST /api/applications
INSERT INTO applications (
    id,
    created_at,
    applicant_id,
    scheme_id
) VALUES (
             gen_random_uuid(), now(), $1, $2
         )
RETURNING *;

-- name: UpdateApplication :one
-- Used for PUT /api/applications/{id}
UPDATE applications
SET
    applicant_id = $2,
    scheme_id = $3
WHERE id = $1 AND deleted_at IS NULL
RETURNING *;

-- name: DeleteApplication :exec
-- Used for DELETE /api/applications/{id}
UPDATE applications
SET
    deleted_at = now()
WHERE id = $1 AND deleted_at IS NULL;

-- name: GetApplicationsWithDetails :many
-- Used for getting applications with applicant and scheme details
SELECT
    app.*,
    a.name as applicant_name,
    a.employment_status as applicant_employment_status,
    s.name as scheme_name
FROM applications app
         JOIN applicants a ON app.applicant_id = a.id AND a.deleted_at IS NULL
         JOIN schemes s ON app.scheme_id = s.id AND s.deleted_at IS NULL
WHERE app.deleted_at IS NULL
ORDER BY app.created_at DESC
LIMIT $1 OFFSET $2;