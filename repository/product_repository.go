package repository

import (
	"Final-Project-Sanbercode-Go-Batch-41/structs"
	"database/sql"
)

func GetAllProduct(db *sql.DB) (err error, results []structs.Product) {
	rows, err := db.Query("SELECT * FROM product")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var product = structs.Product{}
		err = rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity)
		if err != nil {
			panic(err)
		}
		results = append(results, product)
	}
	return
}

func InsertProduct(db *sql.DB, product structs.Product) (err error) {
	errs := db.QueryRow("INSERT INTO product (name, description, price, quantity) VALUES (?, ?, ?, ?)",
		product.Name, product.Description, product.Price, product.Quantity)
	return errs.Err()
}

func UpdateProduct(db *sql.DB, product structs.Product) (err error) {
	errs := db.QueryRow("UPDATE product SET name = ?, description = ?, price = ?, quantity = ? WHERE id = ?",
		product.Name, product.Description, product.Price, product.Quantity, product.ID)
	return errs.Err()
}

func DeleteProduct(db *sql.DB, product structs.Product) (err error) {
	errs := db.QueryRow("UPDATE product SET status = 0 WHERE id = ?",
		product.ID)
	return errs.Err()
}
