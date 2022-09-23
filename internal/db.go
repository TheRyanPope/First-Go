package internal

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "usr"
	password = "pswd"
	dbname   = "postgres"
)

func ConnectDB() *sql.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
