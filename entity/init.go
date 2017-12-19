package entity

import (
	"database/sql"

	_ "github.com/painterdrown/service-computing/go-sqlite3"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./agenda.db")
	if err != nil {
		panic("数据库连接错误")
	}
	_, err = db.Exec(createUserTable)
	checkErr(err)
	_, err = db.Exec(createMeetingTable)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
