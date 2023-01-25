package structs

type DetailInvoice struct {
	InvoiceID    uint64 `json:"invoice_id"`
	ProductID    uint64 `json:"product_id"`
	ProductPrice uint64 `json:"ProductPrice"`
	Quantity     int64  `json:"quantity"`
	TotalPrice   uint64 `json:TotalPrice"`
	Status       uint8  `json:"status"`
}

type HeaderInvoice struct {
	InvoiceID     uint64 `json:"invoice_id"`
	CustomerID    uint64 `json:"customer_id"`
	SubtotalPrice uint64 `json:"SubtotalPrice"`
	Status        uint8  `json:"Status"`
}

type Invoice struct {
	HeaderInvoice
	DetailInvoice []DetailInvoice `json:"detail_invoice"`
}
