package main

import "fmt"
import _ "github.com/go-sql-driver/mysql"

func main() {
	// connect to db
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// run query
	err = selectOneQuery(db, 1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
