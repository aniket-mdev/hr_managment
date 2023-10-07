package querytest

import (
	"database/sql"
	"log"
	"os"
	"testing"

	sqlc_lib "github.com/aniket-mdev/hr_managment/sqlc_lib"
	_ "github.com/lib/pq"
)

const (
	db_driver = "postgres"
	db_source = "postgresql://postgres:root@localhost:5432/postgres?sslmode=disable"
)

var testQueries *sqlc_lib.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(db_driver, db_source)

	if err != nil {
		log.Fatal("cannot connect to db : ", err)
	}

	testQueries = sqlc_lib.New(testDB)
	os.Exit(m.Run())
}
