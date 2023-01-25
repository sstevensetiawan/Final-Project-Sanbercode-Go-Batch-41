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
		err = rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.Status)
		if err != nil {
			panic(err)
		}
		results = append(results, product)
	}
	return
}

func InsertProduct(db *sql.DB, product structs.Product) (err error) {
	errs := db.QueryRow("INSERT INTO product (id, name, description, price, quantity, status) VALUES (?, ?, ?, ?, ?, ?)",
		product.ID, product.Name, product.Description, product.Price, product.Quantity, product.Status)
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
