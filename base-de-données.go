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

func newUsers(email, password, surname, name string) {
	db := createDB()
	defer db.Close()
	tsql := fmt.Sprintf("INSERT INTO users_ticketing(email,password,surname,name) VALUES ('%v', '%v', '%v', '%v');",
		email, password, surname, name)
	_, err := db.Exec(tsql)
	if err != nil {
		fmt.Println("createnewusers", err)
		return

	}
	return
}

// printAllIDAllName parcour et affiche tous les ID & name de la BDD
func printAllIDAllName() {
	var (
		id    int
		name  string
		email string
	)
	db := createDB()
	defer db.Close()
	rows, err := db.Query("select id, name, email  from users_ticketing")
	if err != nil {
		fmt.Println("create rowsquery", err)
		return
	}

	for rows.Next() {
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			fmt.Println("create rowsscan", err)
			return
		}
		fmt.Println(id, name, email)

	}
}
