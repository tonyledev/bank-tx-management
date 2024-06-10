package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

const (
	dbdriver = "postgres"
	dbsource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries

func testMain(m *testing.M) {
	conn, err := sql.Open(dbdriver, dbsource)

	if err != nil {
		log.Fatal("could not connect to db: ", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
