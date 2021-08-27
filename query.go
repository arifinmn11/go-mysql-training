package main

import (
	"database/sql"
	"fmt"
)

func insertQuery(db *sql.DB) error {
	// run query
	_, err := db.Exec("INSERT INTO Customers (FirstName, LastName, Phone, City, Birthday) VALUES (?, ?, ?, ?, ?)", "Lord", "Voldemort", "929-278-0564", "Hogwarts", "1887-09-10")
	if err != nil {
		return err
	}
	fmt.Println("insert success")
	return nil
}

func selectQuery(db *sql.DB) error {
	// run query
	rows, err := db.Query("SELECT CustomerId, FirstName, LastName, Phone, City, Birthday FROM Customers")
	if err != nil {
		return err
	}
	defer rows.Close()

	// parse query result to customer struct
	var result []customer

	// call Next to get each row
	for rows.Next() {
		var each = customer{}

		// map each column to customer's attribute
		var err = rows.Scan(&each.id, &each.firstName, &each.lastName, &each.phone, &each.city, &each.birthday)
		if err != nil {
			return err
		}

		result = append(result, each)
	}

	// print the result
	for _, each := range result {
		fmt.Println(each.id, each.firstName, each.lastName, each.phone, each.city, each.birthday)
	}

	return nil
}

func selectOneQuery(db *sql.DB, id int) error {
	var result = customer{}
	// run query
	err := db.QueryRow("SELECT CustomerId, FirstName, LastName, Phone, City, Birthday FROM Customers WHERE customerId = ?", id).
		Scan(&result.id, &result.firstName, &result.lastName, &result.phone, &result.city, &result.birthday)
	if err != nil {
		return err
	}

	// print the result
	fmt.Println(result.id, result.firstName, result.lastName, result.phone, result.city, result.birthday)
	return nil
}

func updateQuery(db *sql.DB, id int) error {
	// run query
	_, err := db.Exec("UPDATE Customers SET FirstName = ? WHERE CustomerId = ?", "Lanaya", id)
	if err != nil {
		return err
	}
	fmt.Println("update success")
	return nil
}

func deleteQuery(db *sql.DB, id int) error {
	// run query
	_, err := db.Exec("DELETE FROM Customers WHERE CustomerId = ?", id)
	if err != nil {
		return err
	}
	fmt.Println("delete success")
	return nil
}
