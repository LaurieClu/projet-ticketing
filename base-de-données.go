package main

import (
	"database/sql"
	 _ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/users_ticketing") => DSN
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
}