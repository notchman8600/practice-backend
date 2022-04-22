package main

import (
	"errors"
	"fmt"
)

type User struct {
	Id    string `json:"user_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var permanentData User = User{
	Id:    "",
	Name:  "",
	Email: "",
}

func createUser(user User) (err error) {
	permanentData = user
	return nil
}

func getUserData(id string) (user User, err error) {
	if id == "example-user-id" {
		user = permanentData
		return
	} else {
		err = errors.New(fmt.Sprintf("Error: %s", "Wrong user_id"))
		return
	}
}

func main() {
	fmt.Println("Default user data: ", permanentData)
	userData := User{Name: "notch_man", Id: "example-user-id", Email: "notchman@example.com"}
	createUser(userData)
	fmt.Println("Updated user data: ", permanentData)
	getData, _ := getUserData("example-user-id")
	fmt.Println("Updated user data: ", getData)
	getData2, err := getUserData("example-user-id-2")
	fmt.Println("Updated user data: ", getData2)
	fmt.Println("Updated user data: ", err)
}
