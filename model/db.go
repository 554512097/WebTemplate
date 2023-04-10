package model

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var gloableDB *sql.DB

func InitDataBase() (bool, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/goon")
	if err != nil {
		return false, err
	}
	gloableDB = db
	_, err2 := db.Exec(sql_user_create)
	if err2 != nil {
		return false, err2
	}
	log.Println("table user created!")
	return true, nil
}
