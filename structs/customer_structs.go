package structs

type Customer struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	Status uint8  `json:"status"`
}
