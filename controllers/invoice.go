package controllers

import (
	"Final-Project-Sanbercode-Go-Batch-41/database"
	"Final-Project-Sanbercode-Go-Batch-41/repository"
	"Final-Project-Sanbercode-Go-Batch-41/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllInvoice(c *gin.Context) {
	var (
		result gin.H
	)

	invoices, err := repository.GetAllInvoice(database.DBConnection)

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": invoices,
		}
	}

	c.JSON(http.StatusOK, result)
}

func GetInvoice(c *gin.Context) {
	var header structs.HeaderInvoice
	var detail structs.DetailInvoice
	var invoice structs.Invoice
	invoice_id, err := strconv.Atoi(c.Param("invoice_id"))
	var (
		result gin.H
	)

	header.InvoiceID = uint64(invoice_id)
	detail.InvoiceID = uint64(invoice_id)

	err, header_invoice := repository.GetInvoice(database.DBConnection, header)
	err, detail_invoice := repository.GetDetailInvoice(database.DBConnection, detail)
	invoice.HeaderInvoice = header_invoice
	invoice.DetailInvoice = detail_invoice

	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": invoice,
		}
	}

	c.JSON(http.StatusOK, result)
}
func InsertInvoice(c *gin.Context) {
	var invoice structs.Invoice

	err := c.ShouldBindJSON(&invoice)
	if err != nil {
		panic(err)
	}

	err = repository.InsertHeaderInvoice(database.DBConnection, invoice.HeaderInvoice)
	if err != nil {
		panic(err)
	}
	for _, detail_invoice := range invoice.DetailInvoice {
		err = repository.InsertDetailInvoice(database.DBConnection, detail_invoice)
		if err != nil {
			panic(err)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Insert Invoice",
	})
}

func DeleteInvoice(c *gin.Context) {
	var header structs.HeaderInvoice
	var detail structs.DetailInvoice
	invoice_id, err := strconv.Atoi(c.Param("invoice_id"))

	if err != nil {
		panic(err)
	}

	header.InvoiceID = uint64(invoice_id)
	detail.InvoiceID = uint64(invoice_id)

	err = repository.DeleteHeaderInvoice(database.DBConnection, header)
	err = repository.DeleteDetailInvoice(database.DBConnection, detail)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Success Delete Invoice",
	})
}
