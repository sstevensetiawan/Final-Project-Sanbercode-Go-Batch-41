package structs

type User struct {
	ID       uint64 `json:"id"`
	UserType uint16 `json:"user_type"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Status   uint8  `json:"status"`
}
