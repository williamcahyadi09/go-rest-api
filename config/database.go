package config

import (
	"fmt"
	"sync"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB
var db_once sync.Once

func InitDBConnectionPool() {
	db_once.Do(func() {
		var err error

		conn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.DB_HOST,
			config.DB_PORT,
			config.DB_USER,
			config.DB_PASS,
			config.DB_NAME,
		)

		db, err = sqlx.Connect("postgres", conn)

		if nil != err {
			log.Fatal(err)
		}

		db.Ping()
		fmt.Println("DB connected")
	})
}


func CloseDBConnectionPool() {
	db.Close()
	fmt.Println("DB connection pool closed")
}