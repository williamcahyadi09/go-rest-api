package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB
var db_once sync.Once

func InitDBConnectionPool() *sqlx.DB {
	db_once.Do(func() {
		conn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.DB_HOST,
			config.DB_PORT,
			config.DB_USER,
			config.DB_PASS,
			config.DB_NAME,
		)

		var err error
		db, err = sqlx.Connect("postgres", conn)
		if nil != err {
			log.Fatal(err)
		}

		fmt.Println("DB connected")
	})

	err := db.Ping()
	if nil != err {
		log.Fatal(err)
	}
	return db
}
