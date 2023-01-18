package repository

import (
	"Final-Project-Sanbercode-Go-Batch-41/structs"
	"database/sql"
)

func GetAllUser(db *sql.DB) (err error, results []structs.User) {
	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user = structs.User{}
		err = rows.Scan(&user.ID, &user.Name, &user.UserType, &user.Email, &user.MobilePhone)
		if err != nil {
			panic(err)
		}
		results = append(results, user)
	}
	return
}

func InsertUser(db *sql.DB, user structs.User) (err error) {
	errs := db.QueryRow("INSERT INTO user (name, user_type, email, mobile_phone, username, password, created_date, created_user_id, created_user_type, updated_date, updated_user_id, updated_user_type VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)",
		user.Name, user.UserType, user.Email, user.MobilePhone, user.Username, user.Password, user.CreatedDate, user.CreatedUserID, user.CreatedUserType, user.UpdatedDate, user.UpdatedUserID, user.UpdatedUserType)
	return errs.Err()
}

func UpdateUser(db *sql.DB, user structs.User) (err error) {
	errs := db.QueryRow("UPDATE user SET name = $1, email = $2, mobile_phone = $3, username = $4, password = $5, updated_date = $6, updated_user_id = %7, updated_user_type = $8 WHERE id = $9",
		user.Name, user.Email, user.MobilePhone, user.Username, user.Password, user.UpdatedDate, user.UpdatedUserID, user.UpdatedUserType, user.ID)
	return errs.Err()
}

func DeleteUser(db *sql.DB, user structs.User) (err error) {
	errs := db.QueryRow("UPDATE user SET status = $1, updated_date = $2, updated_user_id = $3 updated_user_type = $4 WHERE id = $5",
		user.Status, user.UpdatedDate, user.UpdatedUserID, user.UpdatedUserType, user.ID)
	return errs.Err()
}
