package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
	"github.com/romeq/pac/pkg/db"
)

func envOr(key, def string) string {
	env := os.Getenv(key)
	if env != "" {
		return env
	}

	return def
}

func mustConnectDB() *pgx.Conn {
	connstring := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		envOr("DB_USERNAME", "pac"), envOr("DB_PASSWORD", "pac"),
		envOr("DB_HOST", "127.0.0.1"), envOr("DB_PORT", "5432"),
		envOr("DB_NAME", "pac"),
	)

	conn, err := pgx.Connect(context.Background(), connstring)
	if err != nil {
		panic(err)
	}

	return conn
}

func main() {
	conn := mustConnectDB()
	defer conn.Close(context.Background())

	dbhandle := db.New(conn)

	random := make([]byte, 2^4)
	if _, err := rand.Read(random); err != nil {
		panic(err)
	}

	// Create user
	username := hex.EncodeToString(random)
	_, err := dbhandle.CreateAccount(context.Background(), username)
	if err != nil {
		panic(err)
	}
}
