package drivers

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// DB type details
type DB struct {
	SQL *sql.DB
}

// ConnectSQL
func ConnectSQL(host, port, uname, pass, dbname string) (*DB, error) {
	// DBConn
	var dbConn = &DB{}

	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		"postgres",
		"postgres",
		"db",
		"5432",
		"test_db")

	db, err := sql.Open("postgres", url)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("DB Connection started successfully")

	dbConn.SQL = db

	return dbConn, err
}
