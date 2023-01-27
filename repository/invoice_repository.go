package repository

import (
	"Final-Project-Sanbercode-Go-Batch-41/structs"
	"database/sql"
)

func GetAllInvoice(db *sql.DB) (err error, results []structs.HeaderInvoice) {
	rows, err := db.Query("SELECT * FROM header_invoice WHERE status = 1")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var header_invoice = structs.HeaderInvoice{}
		err = rows.Scan(&header_invoice.InvoiceID, &header_invoice.CustomerID, &header_invoice.SubtotalPrice, &header_invoice.Status)
		if err != nil {
			panic(err)
		}
		results = append(results, header_invoice)
	}
	return
}

func GetInvoice(db *sql.DB, HeaderInvoice structs.HeaderInvoice) (err error, result structs.HeaderInvoice) {
	rows, err := db.Query("SELECT * FROM header_invoice WHERE invoice_id = ? AND status = 1", HeaderInvoice.InvoiceID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var header_invoice = structs.HeaderInvoice{}
		err = rows.Scan(&header_invoice.InvoiceID, &header_invoice.CustomerID, &header_invoice.SubtotalPrice, &header_invoice.Status)
		if err != nil {
			panic(err)
		}
		result = header_invoice
	}
	return
}

func InsertHeaderInvoice(db *sql.DB, HeaderInvoice structs.HeaderInvoice) (err error) {
	errs := db.QueryRow("INSERT INTO header_invoice (invoice_id, customer_id, subtotal_price, status) VALUES (?, ?, ?, ?)",
		HeaderInvoice.InvoiceID, HeaderInvoice.CustomerID, HeaderInvoice.SubtotalPrice, HeaderInvoice.Status)
	return errs.Err()
}

func UpdateHeaderInvoice(db *sql.DB, HeaderInvoice structs.HeaderInvoice) (err error) {
	errs := db.QueryRow("UPDATE header_invoice SET customer_id = ? and subtotal_price = ? WHERE invoice_id = ?",
		HeaderInvoice.CustomerID, HeaderInvoice.SubtotalPrice, HeaderInvoice.InvoiceID)
	return errs.Err()
}

func DeleteHeaderInvoice(db *sql.DB, HeaderInvoice structs.HeaderInvoice) (err error) {
	errs := db.QueryRow("UPDATE header_invoice SET status = 0 WHERE invoice_id = ?",
		HeaderInvoice.InvoiceID)
	return errs.Err()
}

func GetDetailInvoice(db *sql.DB, DetailInvoice structs.DetailInvoice) (err error, results []structs.DetailInvoice) {
	rows, err := db.Query("SELECT * FROM detail_invoice WHERE invoice_id = ? AND status = 1", DetailInvoice.InvoiceID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var detail_invoice = structs.DetailInvoice{}
		err = rows.Scan(&detail_invoice.InvoiceID, &detail_invoice.ProductID, &detail_invoice.ProductPrice, &detail_invoice.Quantity, &detail_invoice.TotalPrice, &detail_invoice.Status)
		if err != nil {
			panic(err)
		}
		results = append(results, detail_invoice)
	}
	return
}

func InsertDetailInvoice(db *sql.DB, DetailInvoice structs.DetailInvoice) (err error) {
	errs := db.QueryRow("INSERT INTO detail_invoice (invoice_id, product_id, product_price, quantity, total_price, status) VALUES (?, ?, ?, ?, ?, ?)",
		DetailInvoice.InvoiceID, DetailInvoice.ProductID, DetailInvoice.ProductPrice, DetailInvoice.Quantity, DetailInvoice.TotalPrice, DetailInvoice.Status)
	return errs.Err()
}

func UpdateDetailInvoice(db *sql.DB, DetailInvoice structs.DetailInvoice) (err error) {
	errs := db.QueryRow("UPDATE detail_invoice SET quantity = ?, total_price = ? WHERE invoice_id = ? AND product_id = ?",
		DetailInvoice.Quantity, DetailInvoice.TotalPrice, DetailInvoice.InvoiceID, DetailInvoice.ProductID)
	return errs.Err()
}

func DeleteDetailInvoice(db *sql.DB, DetailInvoice structs.DetailInvoice) (err error) {
	errs := db.QueryRow("UPDATE detail_invoice SET status = 0 WHERE invoice_id = ?",
		DetailInvoice.InvoiceID)
	return errs.Err()
}
