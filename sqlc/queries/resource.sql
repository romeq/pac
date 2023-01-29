-- name: CreateResource :one
INSERT INTO resource(content)
VALUES ($1)
RETURNING *;

-- name: AddResourceRole :exec
INSERT INTO resource_to_role(resource_uuid, role_uuid) 
VALUES ($1, $2);

-- name: GetResource :one
SELECT * FROM resource
JOIN resource_to_role
ON resource_to_role.resource_uuid = resource.resource_uuid
WHERE resource.resource_uuid = $1;