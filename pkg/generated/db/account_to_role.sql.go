// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: account_to_role.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const addRoleFor = `-- name: AddRoleFor :one
INSERT INTO account_to_role(account_uuid, role_uuid)
VALUES ($1, $2) 
RETURNING account_uuid, role_uuid
`

type AddRoleForParams struct {
	AccountUuid uuid.UUID `json:"account_uuid"`
	RoleUuid    uuid.UUID `json:"role_uuid"`
}

func (q *Queries) AddRoleFor(ctx context.Context, arg AddRoleForParams) (AccountToRole, error) {
	row := q.db.QueryRow(ctx, addRoleFor, arg.AccountUuid, arg.RoleUuid)
	var i AccountToRole
	err := row.Scan(&i.AccountUuid, &i.RoleUuid)
	return i, err
}

const getRolesFor = `-- name: GetRolesFor :many
SELECT role.name
FROM account_to_role
JOIN role ON role.role_uuid = account_to_role.role_uuid
WHERE account_uuid = $1
`

func (q *Queries) GetRolesFor(ctx context.Context, accountUuid uuid.UUID) ([]string, error) {
	rows, err := q.db.Query(ctx, getRolesFor, accountUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
