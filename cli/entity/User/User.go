package User

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/painterdrown/go-agenda/entities"
)

type User struct {
	UserName string
	UserPass string
}

func returnAllUser() ([]entities.User, error) {
	allUser, err := entities.GetAllUsers()
	return allUser, err
}

func IsUser(userName string) bool {
	allUser, _ := returnAllUser()
	for _, value := range allUser {
		if value.Username == userName {
			return true
		}
	}
	return false
}

//增加注册用户
func UserRegitser(body User) error {
	stream, _ := ioutil.ReadFile("curUser.txt")
	if string(stream) != "" {
		return errors.New("You should louout the current account")
	}
	err := entities.Register(body.UserName, body.UserPass)
	return err
}

//登录
func UserLogin(userName, password string) error {
	stream, _ := ioutil.ReadFile("curUser.txt")
	if string(stream) != "" {
		return errors.New("Please logout the current account first")
	}
	allUser, _ := returnAllUser()
	for index, value := range allUser {
		if value.Username == userName && value.Password == password {
			file, _ := os.OpenFile("curUser.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModeAppend|os.ModePerm)
			file.WriteString(value.Username)
			return nil
		} else if value.Username == userName && value.Password != password {
			return errors.New("userName or password is wrong")
		} else if index == len(allUser)-1 {
			return errors.New("You has been not registered a account")
		}
	}
	return errors.New("You has been not registered a account")
}

//登出
func UserLogout() error {
	stream, _ := ioutil.ReadFile("curUser.txt")
	if string(stream) == "" {
		return errors.New("You has been logout")
	}
	file, _ := os.OpenFile("curUser.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModeAppend|os.ModePerm)
	file.WriteString("")
	return nil
}

//检测登录状态
func UserState() string {
	stream, _ := ioutil.ReadFile("curUser.txt")
	if string(stream) == "" {
		fmt.Println("You haven't logged in yet, Please login at first")
		return ""
	} else {
		fmt.Println(string(stream))
		return string(stream)
	}
}

//已登录用户删除用户
func UserDelete() error {
	stream, _ := ioutil.ReadFile("curUser.txt")
	if string(stream) == "" {
		return errors.New("Please login first")
	}
	err := entities.DeleteUser(string(stream))
	file1, _ := os.OpenFile("curUser.txt", os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModeAppend|os.ModePerm)
	file1.WriteString("")
	return err
}

func QueryUserByUserName() {
	allUser, _ := returnAllUser()
	fmt.Println("The search result can be show as followed:")
	fmt.Println("UserName      ")
	for _, value := range allUser {
		fmt.Println(value.Username)
	}
}
