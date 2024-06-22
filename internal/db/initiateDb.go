package db

import (
	"database/sql"
	"log"
)

func (D *Database) InitiateDB() {
	var err error

	D.db, err = sql.Open("sqlite3", "./purkheli.db")
	if err != nil {
		log.Fatal(err)
	}
	_, err = D.db.Exec("create table purkheli (title string, textarea string, imgpath string, time int)")
	// if err != nil {
	// 	// log.Panic(err)
	// }

	_, err = D.db.Exec("create table passcodes (passcode string)")
}
