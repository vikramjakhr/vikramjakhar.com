package config

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB(driverName, dataSource string) {
	var err error
	DB, err = sql.Open(driverName, dataSource)
	if err != nil {
		log.Panic(err)
	}

	if err = DB.Ping(); err != nil {
		log.Panic(err)
	}
}
