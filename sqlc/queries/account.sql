-- name: CreateAccount :one
INSERT INTO account(name) VALUES($1)
RETURNING *;