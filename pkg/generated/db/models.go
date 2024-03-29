// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"github.com/google/uuid"
)

type Account struct {
	AccountUuid uuid.UUID `json:"account_uuid"`
	Name        string    `json:"name"`
}

type AccountToRole struct {
	AccountUuid uuid.UUID `json:"account_uuid"`
	RoleUuid    uuid.UUID `json:"role_uuid"`
}

type Resource struct {
	ResourceUuid uuid.UUID `json:"resource_uuid"`
	Content      string    `json:"content"`
}

type ResourceToRole struct {
	ResourceUuid uuid.UUID `json:"resource_uuid"`
	RoleUuid     uuid.UUID `json:"role_uuid"`
}

type Role struct {
	RoleUuid uuid.UUID `json:"role_uuid"`
	Name     string    `json:"name"`
}
