package repository

import (
	"Final-Project-Sanbercode-Go-Batch-41/structs"
	"database/sql"
)

func GetAllUser(db *sql.DB) (err error, results []structs.User) {
	rows, err := db.Query("SELECT * FROM user_account")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user = structs.User{}
		err = rows.Scan(&user.ID, &user.UserType, &user.Name, &user.Username, &user.Password, &user.Status)
		if err != nil {
			panic(err)
		}
		results = append(results, user)
	}
	return
}

func InsertUser(db *sql.DB, user structs.User) (err error) {
	errs := db.QueryRow("INSERT INTO user_account (id, name, user_type, username, password) VALUES (?, ?, ?, ?, ?)",
		user.ID, user.Name, user.UserType, user.Username, user.Password)
	return errs.Err()
}

func UpdateUser(db *sql.DB, user structs.User) (err error) {
	errs := db.QueryRow("UPDATE user_account SET name = ?, username = ?, password = ? WHERE id = ?",
		user.Name, user.Username, user.Password, user.ID)
	return errs.Err()
}

func DeleteUser(db *sql.DB, user structs.User) (err error) {
	errs := db.QueryRow("UPDATE user_account SET status = 0 WHERE id = ?", user.ID)
	return errs.Err()
}
