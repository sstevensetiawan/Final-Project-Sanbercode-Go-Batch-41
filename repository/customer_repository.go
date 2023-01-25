package repository

import (
	"Final-Project-Sanbercode-Go-Batch-41/structs"
	"database/sql"
)

func GetAllCustomer(db *sql.DB) (err error, results []structs.Customer) {
	rows, err := db.Query("SELECT * FROM customer")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var customer = structs.Customer{}
		err = rows.Scan(&customer.ID, &customer.Name, &customer.Status)
		if err != nil {
			panic(err)
		}
		results = append(results, customer)
	}
	return
}

func InsertCustomer(db *sql.DB, customer structs.Customer) (err error) {
	errs := db.QueryRow("INSERT INTO customer (id, name, status) VALUES (?, ?, ?)",
		customer.ID, customer.Name, customer.Status)
	return errs.Err()
}

func UpdateCustomer(db *sql.DB, customer structs.Customer) (err error) {
	errs := db.QueryRow("UPDATE customer SET name = ? WHERE id = ?",
		customer.Name, customer.ID)
	return errs.Err()
}

func DeleteCustomer(db *sql.DB, customer structs.Customer) (err error) {
	errs := db.QueryRow("UPDATE customer SET status = 0 WHERE id = ?",
		customer.ID)
	return errs.Err()
}
