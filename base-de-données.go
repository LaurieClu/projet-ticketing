package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func createDB() (db *sql.DB) {
	db, err := sql.Open("mysql",
		"newproject:12345@tcp(localhost)/publications")
	if err != nil {
		fmt.Println("fonction createDB", err)
	}
	return db
}

func newUsers(name string) {
	db := createDB()
	defer db.Close()
	tsql := fmt.Sprintf("INSERT INTO users_ticketing(name) values('%s');",
		name)
	_, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("createtname", err)
		return

	}
	return
}

func newEmail(email string) {
	db := createDB()
	defer db.Close()
	tsql := fmt.Sprintf("INSERT INTO users_ticketing(email) values('%s');",
		email)
	_, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("createtemail", err)
		return

	}
	return
}

func newSurname(surname string) {
	db := createDB()
	defer db.Close()
	tsql := fmt.Sprintf("INSERT INTO users_ticketing(surname) values('%s');",
		surname)
	_, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("createtsurname", err)
		return

	}
	return
}

func newBirthday(birthday string) {
	db := createDB()
	defer db.Close()
	tsql := fmt.Sprintf("INSERT INTO users_ticketing(birthday) values('%s');",
		birthday)
	_, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("createbirthday", err)
		return

	}
	return
}

func newPassword(password string) {
	db := createDB()
	defer db.Close()
	tsql := fmt.Sprintf("INSERT INTO users_ticketing(password) values('%s');",
		password)
	_, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("createpassword", err)
		return

	}
	return
}

// printAllIDAllName parcour et affiche tous les ID & name de la BDD
func printAllIDAllName() {
	var (
		id   int
		name string
	)
	db := createDB()
	defer db.Close()
	rows, err := db.Query("select id, name from users_ticketing")
	if err != nil {
		fmt.Println("create rowsquery", err)
		return
	}

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Println("create rowsscan", err)
			return
		}
		fmt.Println(id, name)

	}
}
