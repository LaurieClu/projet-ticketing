package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql",
		"newproject:12345@tcp(localhost)/publications")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var (
		id   int
		name string
	)
	rows, err := db.Query("select id, name from users_ticketing where id = ?", 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer db.Close()
	}

	newusers(db, "riendutout")
}

func newusers(db *sql.DB, name string) {
	tsql := fmt.Sprintf("INSERT INTO users_ticketing(name) values('%s');",
		name)
	_, err := db.Exec(tsql)
	if err != nil {
		fmt.Println(err)
		return

	}
	return
}
