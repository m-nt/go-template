package internal

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type DB struct {
	Conn *pgx.Conn
}

func inintialize_database(db *DB) {
	_settings := Get_Settings()
	conn, err := pgx.Connect(context.Background(), _settings.DATABASE_URL)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	db.Conn = conn
	// set up database tables and roles
}

func Close_database(db *DB) {
	db.Conn.Close(context.Background())
}

func Get_database() DB {
	var db DB
	inintialize_database(&db)
	return db
}
