package db

import (
	"database/sql"
	"server/helper"
	"time"

	
)

type Database struct {
	DB *sql.DB
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_chat_app")
	helper.PanicIfError(err, "error sql open :")

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	return &Database{DB: db}, nil
}

func (d *Database) Close() {
	d.DB.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.DB
}
