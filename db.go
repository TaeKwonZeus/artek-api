package main

import (
	"fmt"

	"github.com/TaeKwonZeus/artek-api/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE IF NOT EXISTS test (
    id SERIAL PRIMARY KEY
);
`

func database(config config.DBConfig) (*sqlx.DB, error) {
	connString := fmt.Sprintf("user=%s password=%s dbname=%s", config.User, config.Password, config.Name)

	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil, err
	}

	db.MustExec(schema)

	return db, nil
}
