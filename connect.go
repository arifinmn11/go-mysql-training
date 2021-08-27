package main

import "database/sql"

func connect() (*sql.DB, error) {
	// format : username:password@tcp(host:port)/database-name
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/rakamin")
	if err != nil {
		return nil, err
	}
	return db, nil
}
