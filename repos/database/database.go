package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	//host     = "localhost"
	host     = "pgdb"
	port     = 5432
	user     = "user"
	password = "pass"
	dbname   = "db"
)

func ConnectToDb() *sql.DB {
	pgInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, _ := sql.Open("postgres", pgInfo)

	//defer db.Close()

	err := db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	return db
}

/*func Close(db *sql.DB) {
	if db != nil && db.db != nil {
		err := conn.db.Close()
		if err != nil {
			log.Error("db close: ", err.Error())
		}

		conn = nil
	}
}*/
