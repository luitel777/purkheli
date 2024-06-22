package db

// create a db instance
// value from form is inserted here
import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

func (D *Database) StoreDB(title string, textarea string, imgpath string) {
	sql_query := "insert into purkheli (title, textarea, imgpath, time) values (?, ?, ?, ?)"
	currentTime := time.Now().Unix()
	stmt, err := D.db.Prepare(sql_query)
	if err != nil {
		log.Panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, textarea, imgpath, currentTime)
	if err != nil {
		log.Panic("StoreDB() store: ", err)
	}

	// since the rowid is auto incremented when we store
	// we know the last auto incremented rowid is the rowid
	// of the current latest post
	// and we want to create a new table for it as soon as
	// we make a column for the post
	sql_query = "select max(rowid) from purkheli"
	rows, err := D.db.Query(sql_query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	var lastRowId int
	for rows.Next() {
		err = rows.Scan(&lastRowId)
		if err != nil {
			log.Panic(err)
		}
	}

	// create new table of the id rowid
	sql_query = "create table post_%d (c_title string, c_textarea string, c_imgpath string, c_time int)"
	sql_query = fmt.Sprintf(sql_query, lastRowId)
	_, err = D.db.Exec(sql_query)
	if err != nil {
		log.Panic(err)
	}
}

func (D *Database) StoreComments(title string, textarea string, id int, imgpath string) {
	sql_query := "insert into post_%d (c_title, c_textarea, c_imgpath, c_time) values (?, ?, ?, ?)"
	currentTime := time.Now().Unix()
	sql_query = fmt.Sprintf(sql_query, id)
	stmt, err := D.db.Prepare(sql_query)
	if err != nil {
		log.Panic(err)
	}
	defer stmt.Close()

	// tx, err := D.db.Begin()
	// if err != nil {
	// 	log.Panic("StoreDB(): ", err)
	// }
	_, err = stmt.Exec(title, textarea, imgpath, currentTime)
	if err != nil {
		log.Panic("StoreDB() store: ", err)
	}
	// err = tx.Commit()
}

func (D *Database) CloseDB() {
	D.db.Close()
}
