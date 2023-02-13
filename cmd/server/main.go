package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/romeq/pac/cmd/server/api"
	"github.com/romeq/pac/pkg/db"
)

func envOr(key, def string) string {
	env := os.Getenv(key)
	if env != "" {
		return env
	}

	return def
}

func mustConnectDB(ctx context.Context) *pgx.Conn {
	connstring := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		envOr("DB_USERNAME", "pac"), envOr("DB_PASSWORD", "pac"),
		envOr("DB_HOST", "127.0.0.1"), envOr("DB_PORT", "5432"),
		envOr("DB_NAME", "pac"),
	)

	conn, err := pgx.Connect(ctx, connstring)
	if err != nil {
		panic(err)
	}

	return conn
}

func main() {
	dbctx := context.Background()
	conn := mustConnectDB(dbctx)
	defer conn.Close(dbctx)

	dbhandle := db.New(conn)
	api := api.NewAPI(dbhandle)
	router := echo.New()

	// I wish all of this could be done INSIDE A SIMPLE CONFIGURATION STRUCT SO THIS WOULD STAY CLEAN
	router.RouteNotFound("/*", api.NotFound)
	router.HideBanner = true
	router.HidePort = true
	router.Debug = false
	router.HTTPErrorHandler = api.ErrorHandler

	// I hate variable naming
	randomBytes := make([]byte, 16)
	if _, err := rand.Read(randomBytes); err != nil {
		log.Fatal(err)
	}

	adminPassword := envOr("PAC_PASSWORD", hex.EncodeToString(randomBytes))
	log.Printf("Authentication token: %s", adminPassword)
	adminAccess := api.AdminAuthentication(adminPassword)

	group := router.Group("/account")
	{
		group.POST("/create", api.CreateAccount)
		group.PUT("/add-role", adminAccess(api.AddAccountRole))
	}

	group = router.Group("/resource")
	{
		group.POST("/create", api.CreateResource)
		group.PUT("/add-role", adminAccess(api.AddResourceRole))
	}

	group = router.Group("/role")
	{
		group.POST("/create", api.CreateRole)
	}

	ladr := envOr("PAC_LISTENADDR", "127.0.0.1:8000")
	log.Println("starting server on", ladr)
	if err := router.Start(ladr); err != nil {
		log.Fatal(err)
	}
}
