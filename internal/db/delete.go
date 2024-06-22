package db

import (
	"errors"
	"fmt"
)

func (D *Database) CheckPasscode(passcode string) error {
	var sql_query = "select * from passcodes where passcode = '%s'"
	sql_query = fmt.Sprintf(sql_query, passcode)

	rows, err := D.db.Query(sql_query)
	if err != nil {
		return errors.New("cannot query the database")
	}
	defer rows.Close()
	var pass string
	for rows.Next() {
		err = rows.Scan(&pass)
		if err != nil {
			return errors.New("passcode not found")
		}
	}
	if pass != passcode {
		return errors.New("passcode not found")
	}

	return nil
}

func (D *Database) DeletePost(id string, passcode string) error {
	err := D.CheckPasscode(passcode)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	sql_query := fmt.Sprintf("drop table if exists post_%s", id)
	_, err = D.db.Exec(sql_query)
	if err != nil {
		return errors.New("cannot delete the post")
	}

	sql_query = fmt.Sprintf("delete from purkheli where ROWID = %s", id)
	_, err = D.db.Exec(sql_query)
	if err != nil {
		return errors.New("cannot delete the post")
	}
	return nil
}
