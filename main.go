package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       int
	Username string
	Email    string
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO users (username, email) VALUES ('aryasukma', 'arya@gmail.com')")
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}

	users := []User{}
	for results.Next() {
		user := User{}
		err := results.Scan(&user.Id, &user.Username, &user.Email)

		if err != nil {
			panic(err.Error())
		}

		users = append(users, User{user.Id, user.Username, user.Email})
	}

	fmt.Println(users)

	defer results.Close()
}
