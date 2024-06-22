package db

import (
	"fmt"
	"log"
	"strconv"

	"github.com/luitel777/purkheli/config"
)

// retrive only returns map[int]data from database
// encoding and decoding is left to the api /api/getposts
func (D *Database) Retrive(sql_query string) (error, map[int][]string) {
	totalRetrivedData := make(map[int][]string)

	rows, err := D.db.Query(sql_query)
	if err != nil {
		fmt.Println(err)
		return err, make(map[int][]string)
	}
	defer rows.Close()

	var rowid int
	var title string
	var textarea string
	var imgpath string
	var time int
	for rows.Next() {
		err = rows.Scan(&rowid, &title, &textarea, &imgpath, &time)
		totalRetrivedData[rowid] = []string{title, textarea, imgpath, strconv.Itoa(time)}
		if err != nil {
			log.Panic(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Panic(err)
	}
	return err, totalRetrivedData

}

func (D *Database) RetriveDb() (error, map[int][]string) {
	var sql_query = "select ROWID, title, textarea, imgpath, time from purkheli order by ROWID desc limit %d"
	sql_query = fmt.Sprintf(sql_query, config.SQL_MAX_POSTS)
	return D.Retrive(sql_query)
}

func (D *Database) RetriveSinglePost(id int) (error, []string) {
	var sql_query = "select title, textarea, imgpath, time from purkheli where ROWID=%d"
	sql_query = fmt.Sprintf(sql_query, id)
	rows, err := D.db.Query(sql_query)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var title string
	var textarea string
	var imgpath string
	var time int
	for rows.Next() {
		err = rows.Scan(&title, &textarea, &imgpath, &time)
		if err != nil {
			log.Panic(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Panic(err)
	}
	fields := []string{title, textarea, imgpath, strconv.Itoa(time)}
	return err, fields

}

func (D *Database) RetriveCommentsDb(postid int) (error, map[int][]string) {
	var sql_query = "select ROWID, c_title, c_textarea, c_imgpath, c_time from post_%d order by ROWID desc"
	sql_query = fmt.Sprintf(sql_query, postid)
	return D.Retrive(sql_query)
}
