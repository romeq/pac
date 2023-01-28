// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: accounts.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO account(name) VALUES($1)
RETURNING uuid, name
`

func (q *Queries) CreateAccount(ctx context.Context, name string) (Account, error) {
	row := q.db.QueryRow(ctx, createAccount, name)
	var i Account
	err := row.Scan(&i.Uuid, &i.Name)
	return i, err
}