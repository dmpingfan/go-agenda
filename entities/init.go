package entities

import (
	"database/sql"

	_ "github.com/painterdrown/service-computing/go-sqlite3" //
)

var db *sql.DB

// InitDB .
func InitDB(dbpath string) {
	var err error
	// TODO
	db, err = sql.Open("sqlite3", dbpath)
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
