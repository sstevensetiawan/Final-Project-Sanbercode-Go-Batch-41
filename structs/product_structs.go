package structs

type Product struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Quantity    int64  `json:"quantity"`
	Status      uint8  `json:"status"`
}
