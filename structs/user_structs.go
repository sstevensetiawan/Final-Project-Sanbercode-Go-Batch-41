package structs

import "time"

type User struct {
	ID              uint64    `json:"id"`
	UserType        uint16    `json:"user_type"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	MobilePhone     string    `json:"mobile_phone"`
	Username        string    `json:"username"`
	Password        string    `json:"password"`
	Status          uint8     `json:"status"`
	CreatedDate     time.Time `json:"created_date"`
	CreatedUserID   uint16    `json:"created_user_id"`
	CreatedUserType uint64    `json:"created_user_type"`
	UpdatedDate     time.Time `json:"updated_date"`
	UpdatedUserID   uint16    `json:"updated_user_id"`
	UpdatedUserType uint64    `json:"updated_user_type"`
}
