package structs

type DetailInvoice struct {
	InvoiceID    uint64 `json:"invoice_id"`
	ProductID    uint64 `json:"product_id"`
	ProductPrice uint64 `json:"product_price"`
	Quantity     int64  `json:"quantity"`
	TotalPrice   uint64 `json:"total_price"`
	Status       uint8  `json:"status"`
}

type HeaderInvoice struct {
	InvoiceID     uint64 `json:"invoice_id"`
	CustomerID    uint64 `json:"customer_id"`
	SubtotalPrice uint64 `json:"subtotal_price"`
	Status        uint8  `json:"status"`
}

type Invoice struct {
	HeaderInvoice HeaderInvoice   `json:header_invoice`
	DetailInvoice []DetailInvoice `json:"detail_invoice"`
}
