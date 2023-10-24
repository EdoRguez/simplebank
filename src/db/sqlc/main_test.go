package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDrive  = "postgres"
	dbSource = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDrive, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
