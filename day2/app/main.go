package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}

type SqlResult struct {
	Result sql.Result
}

type SqlRow struct {
	Rows *sql.Rows
}

type SqlHandler struct {
	Conn *sql.DB
}

type User struct {
	Id    string `json:"user_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewSqlHandler() *SqlHandler {
	sql_url := os.Getenv("SQL_URL")

	conn, err := sql.Open("postgres", sql_url)

	if err != nil {
		fmt.Println(err.Error())
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (SqlResult, error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	row.Rows = rows
	return row, nil
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}

func main() {
	// SQLのインスタンスを作成
	sqlHandler := NewSqlHandler()
	// UUIDを作成
	uuidObj, _ := uuid.NewUUID()
	// データベースに値を登録
	sqlHandler.Conn.Exec(`insert into users(user_id,name,email) values($1,'hoge','piyo@example.com')`, uuidObj.String())

	// データベースから値を読み出し
	u := &User{}
	if err := sqlHandler.Conn.QueryRow("select user_id,name,email from users where user_id=$1", uuidObj.String()).Scan(&u.Id, &u.Name, &u.Email); err != nil {
		log.Fatal(err)
	}

	//読み出せたか値を確認
	fmt.Println(u)

}
