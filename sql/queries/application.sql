-- name: CreateApplication :one
INSERT INTO applications (id, project_id, name, description, repo_url, created_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: CheckApplicationExistsByName :one
SELECT EXISTS (
    SELECT 1 FROM applications WHERE name = $1
) AS exists;

-- name: GetApplicationByName :one
SELECT * FROM applications
WHERE name = $1
LIMIT 1;

-- name: DeleteApplicationByName :exec
DELETE FROM applications
WHERE name = $1;

-- name: ListAllProjectApplications :many
SELECT * FROM applications
WHERE project_id = (
  SELECT id FROM projects WHERE projects.name = $1
);

-- name: GetLatestApplicationVersionByApplicationName :one
SELECT av.* FROM application_versions av
JOIN applications a ON av.application_id = a.id
WHERE a.name = $1
ORDER BY av.created_at DESC
LIMIT 1;
