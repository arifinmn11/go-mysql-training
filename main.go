package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// connect to db
	// format : username:password@tcp(host:port)/database-name
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/rakamin")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
}
