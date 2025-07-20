package drivers

import (
	"database/sql"
	// no-lint
	"fmt"
	// no-lint
	_ "github.com/go-sql-driver/mysql"
)

// DB type details
type DB struct {
	SQL *sql.DB
	// Mgo *mgo.database
}

// ConnectSQL
func ConnectSQL(host, port, uname, pass, dbname string) (*DB, error) {
	// DBConn
	var dbConn = &DB{}

	dbSource := fmt.Sprintf(
		"docker:%s@tcp(%s:%s)/%s?charset=utf8",
		pass,
		host,
		port,
		dbname,
	)

	d, err := sql.Open("mysql", dbSource)

	if err != nil {
		panic(err)
	}
	dbConn.SQL = d

	return dbConn, err
}
