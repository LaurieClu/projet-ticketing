package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	id   int
	name string
)

func createDB() (db *sql.DB) {
	db, err := sql.Open("mysql",
		"newproject:12345@tcp(localhost)/publications")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	return db
}

func newUsers(name string) {
	db := createDB()
	tsql := fmt.Sprintf("INSERT INTO users_ticketing(name) values('%s');",
		name)
	_, err := db.Exec(tsql)
	if err != nil {
		fmt.Println(err)
		return

	}
	return
}

// printAllIDAllName parcour et affiche tous les ID & name de la BDD
func printAllIDAllName() {
	db := createDB()
	rows, err := db.Query("select id, name from users_ticketing")
	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()
		fmt.Println(id, name)
	}
}
