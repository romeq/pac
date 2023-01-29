package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"log"
	mathRand "math/rand"
	"testing"

	_ "github.com/lib/pq"
	"github.com/romeq/pac/pkg/db"
)

func mustNotT(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func mustGenerateRandomString(length int) string {
	random := make([]byte, length/2)
	if _, err := rand.Read(random); err != nil {
		panic(err)
	}
	return hex.EncodeToString(random)
}

func TestMain(t *testing.T) {
	dbctx := context.Background()
	conn := mustConnectDB(dbctx)
	defer conn.Close(dbctx)

	dbhandle := db.New(conn)

	// create user
	username := mustGenerateRandomString(10)
	user, err := dbhandle.CreateAccount(context.Background(), username)
	mustNotT(t, err)

	// create roles
	var roles []db.Role
	for i := 0; i < 3; i++ {
		rolename := mustGenerateRandomString(4)
		role, err := dbhandle.CreateRole(context.Background(), rolename)
		mustNotT(t, err)

		roles = append(roles, role)
	}
	role := roles[mathRand.Intn(len(roles))]

	// bind role to user
	_, err = dbhandle.AddRoleFor(context.Background(), db.AddRoleForParams{
		AccountUuid: user.AccountUuid,
		RoleUuid:    role.RoleUuid,
	})
	mustNotT(t, err)

	// get roles for user
	rolesForUser, err := dbhandle.GetRolesFor(context.Background(), user.AccountUuid)
	mustNotT(t, err)
	for _, role := range rolesForUser {
		log.Printf("%s has role %s\n", user.Name, role)
	}

	// create resource
	resource, err := dbhandle.CreateResource(context.Background(), "yes")
	mustNotT(t, err)
	log.Println("created new resource:", resource.ResourceUuid)

	// bind role to resource
	err = dbhandle.AddResourceRole(context.Background(), db.AddResourceRoleParams{
		ResourceUuid: resource.ResourceUuid,
		RoleUuid:     role.RoleUuid,
	})
	mustNotT(t, err)

	resourceByUser, err := dbhandle.GetResource(context.Background(), resource.ResourceUuid)
	mustNotT(t, err)

	log.Println("got content:", resourceByUser.Content)
}
