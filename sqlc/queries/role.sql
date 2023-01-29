-- name: CreateRole :one
INSERT INTO role(name) VALUES($1)
RETURNING *;