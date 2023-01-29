-- name: AddRoleFor :one
INSERT INTO account_to_role(account_uuid, role_uuid)
VALUES ($1, $2) 
RETURNING *;

-- name: GetRolesFor :many
SELECT role.name
FROM account_to_role
JOIN role ON role.role_uuid = account_to_role.role_uuid
WHERE account_uuid = $1;