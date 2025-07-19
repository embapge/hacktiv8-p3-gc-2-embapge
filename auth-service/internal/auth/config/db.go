package config

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func MySQLInit() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost)/p3_gc_2_embapge?parseTime=true")
	if err != nil {
		log.Fatal("Cannot connnect to mysql")
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot ping to mysql")
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
