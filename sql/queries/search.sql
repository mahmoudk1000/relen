-- name: SearchProjectsByName :many
SELECT *
FROM projects
WHERE name = $1;

-- name: SearchProjectsByNamePattern :many
SELECT *
FROM projects
WHERE name ~ ($1::text)
ORDER BY name;

-- name: SearchProjectsByStatus :many
SELECT *
FROM projects
WHERE status = $1
ORDER BY name;

-- name: SearchProjectsByMetadata :many
SELECT *
FROM projects
WHERE metadata @> ($1::jsonb)
ORDER BY name;

-- name: SearchApplicationsByName :many
SELECT a.*, p.name AS project_name
FROM applications a
JOIN projects p ON a.project_id = p.id
WHERE a.name = $1;

-- name: SearchApplicationsByNamePattern :many
SELECT a.*, p.name AS project_name
FROM applications a
JOIN projects p ON a.project_id = p.id
WHERE a.name ~ ($1::text)
ORDER BY p.name, a.name;

-- name: SearchApplicationsByStatus :many
SELECT a.*, p.name AS project_name
FROM applications a
JOIN projects p ON a.project_id = p.id
WHERE a.status = $1
ORDER BY p.name, a.name;

-- name: SearchVersionsByName :many
SELECT av.*, a.name AS app_name, p.name AS project_name
FROM application_versions av
JOIN applications a ON av.application_id = a.id
JOIN projects p ON a.project_id = p.id
WHERE av.version = $1;

-- name: SearchVersionsByPattern :many
SELECT av.*, a.name AS app_name, p.name AS project_name
FROM application_versions av
JOIN applications a ON av.application_id = a.id
JOIN projects p ON a.project_id = p.id
WHERE av.version ~ ($1::text)
ORDER BY p.name, a.name, av.created_at DESC;

-- name: SearchVersionsByStatus :many
SELECT av.*, a.name AS app_name, p.name AS project_name
FROM application_versions av
JOIN applications a ON av.application_id = a.id
JOIN projects p ON a.project_id = p.id
WHERE av.status = $1
ORDER BY p.name, a.name, av.created_at DESC;

-- name: SearchVersionsByGitHash :many
SELECT av.*, a.name AS app_name, p.name AS project_name
FROM application_versions av
JOIN applications a ON av.application_id = a.id
JOIN projects p ON a.project_id = p.id
WHERE av.git_hash = $1;

-- name: SearchReleasesByName :many
SELECT pv.*, p.name AS project_name
FROM project_versions pv
JOIN projects p ON pv.project_id = p.id
WHERE pv.version = $1;

-- name: SearchReleasesByPattern :many
SELECT pv.*, p.name AS project_name
FROM project_versions pv
JOIN projects p ON pv.project_id = p.id
WHERE pv.version ~ ($1::text)
ORDER BY p.name, pv.created_at DESC;
