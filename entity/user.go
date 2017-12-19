package entity

import (
	"errors"
)

type User struct {
	Username string
	Password string
}

var createUserTable = `
CREATE TABLE IF NOT EXISTS user (
	username VARCHAR(32) NOT NULL,
	password VARCHAR(32) NOT NULL
);
`
var findUserByUsername = "SELECT * FROM user WHERE user.username = ?;"
var findAllUsers = "SELECT * FROM user;"
var createUser = "INSERT INTO user (username, password) VALUES (?, ?);"
var deleteUser = "DELETE FROM user WHERE user.username = ?;"

func Register(username, password string) error {
	var err error
	// 检查是否被注册
	user := User{}
	row := db.QueryRow(findUserByUsername, username)
	row.Scan(&user.Username, &user.Password)
	if user.Username != "" {
		return errors.New("该用户已被注册")
	}
	// 插入数据库
	_, err = db.Exec(createUser, username, password)
	checkErr(err)
	return err
}

func GetAllUsers() ([]User, error) {
	var err error
	rows, err := db.Query(findAllUsers)
	checkErr(err)
	defer rows.Close()
	users := make([]User, 0, 0)
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Username, &user.Password)
		checkErr(err)
		users = append(users, user)
	}
	return users, err
}

func DeleteUser(username string) error {
	var err error
	_, err = db.Exec(deleteUser, username)
	checkErr(err)
	return err
}
