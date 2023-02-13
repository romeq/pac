package api

import (
	"encoding/json"
	"io"

	"github.com/romeq/pac/pkg/generated/db"
)

type API struct {
	db *db.Queries
}

func NewAPI(conn *db.Queries) API {
	return API{
		db: conn,
	}
}

func parseBodyToStruct[T any](b io.Reader) (T, error) {
	var body T
	return body, json.NewDecoder(b).Decode(&body)
}
