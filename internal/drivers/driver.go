package drivers

import (
	"database/sql"
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

	connStr := "postgres://postgres:password@localhost:5432/test_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Println("DB Ping Failed")
		log.Fatal(err)
	}
	log.Println("DB Connection started successfully")

	if err != nil {
		panic(err)
	}
	dbConn.SQL = db

	return dbConn, err
}
