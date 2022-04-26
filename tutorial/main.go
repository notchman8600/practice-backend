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

func readUser(id string) (user User, err error) {
	if id == "example-user-id" {
		user = permanentData
		return
	} else {
		err = errors.New(fmt.Sprintf("Error: %s", "Wrong user_id"))
		return
	}
}
func updateUser(user User) (err error) {
	permanentData = user
	return
}

func deleteUser() (err error) {
	permanentData = User{}
	return
}

func main() {
	//デフォルト値の確認
	fmt.Println("Default user data: ", permanentData)
	//ユーザーデータの作成
	userData := User{Name: "notch_man", Id: "example-user-id", Email: "notchman@example.com"}
	//ユーザーの作成
	createUser(userData)
	//結果の確認
	fmt.Println("Created user data: ", permanentData)
	//ユーザーデータの読み出し（ハッピーパス）
	getData, _ := readUser("example-user-id")
	fmt.Println("Read user data: ", getData)
	//ユーザーデータの読み出し（ネガティブパス）
	getData2, err := readUser("example-user-id-2")
	fmt.Println("Read user data: ", getData2)
	fmt.Println("Read user data: ", err)
	//更新用データの作成
	//ユーザーデータの作成
	updatedUserData := User{Name: "notch_man", Id: "example-user-id", Email: "notchman@example.jp"}
	_ = updateUser(updatedUserData)
	fmt.Println("Updated user data: ", permanentData)

	//データの削除
	_ = deleteUser()
	fmt.Println("Deleted user data: ", permanentData)
}
