package db

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbdriver = "postgres"
	dbsource = "postgresql://root:secret@localhost:5433/simplebank?sslmode=disable"
)

var testQueries *Queries

func Test_Main(m *testing.T) {
	conn, err := sql.Open(dbdriver, dbsource)

	if err != nil {
		log.Fatal("could not connect to db: ", err)
	}

	testQueries = New(conn)
}
