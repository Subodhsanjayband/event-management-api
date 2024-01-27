package models

import (
	"errors"

	"github.com/Subodhsanjayband/event_manager/db"
	"github.com/Subodhsanjayband/event_manager/utils"
)

type User struct {
	ID       int64
	Email    string `binding: "required"`
	Password string `binding: "required"`
}

func (u *User) Save() error {
	insertQuery := `
	INSERT INTO users (id,Email,Password) 
	VALUES (?,?,?)
	`
	stmt, err := db.DB.Prepare(insertQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	id, err := GetUserIDSeq()
	if err != nil {
		return err
	}
	u.ID = id
	_, err = stmt.Exec(u.ID, u.Email, hashedPass)
	if err != nil {
		return err
	}

	return err

}

func (u User) ValidateUser() error {
	selectQuery := `
SELECT  password FROM users WHERE email = ? `

	row := db.DB.QueryRow(selectQuery, u.Email)
	var retrievedPassword string
	err := row.Scan(&retrievedPassword)
	if err != nil {
		return errors.New("invalid credentials @FETCHING PASSWORD")
	}
	isValid := utils.ValidatePassword(u.Password, retrievedPassword)
	if !isValid {
		return errors.New("invalid credentials @VALIDATEUSER")
	}
	return nil
}

func (u User) GetUserID() (int64, error) {
	selectQuery := `SELECT id FROM users where email = ? `

	row := db.DB.QueryRow(selectQuery, u.Email)
	var id int64
	err := row.Scan(&id)

	if err != nil {
		return -1, errors.New("invalid credentials @GETUSERID" + err.Error())
	}
	return id, nil
}

func GetUserIDSeq() (int64, error) {
	seqQuery := `SELECT max(ID) FROM users`

	row := db.DB.QueryRow(seqQuery)
	var id int64
	err := row.Scan(&id)
	if err != nil {
		return -1, errors.New("unable to generate user id ")
	}

	return id + 1, nil

}

func GetAllUsers() ([]User, error) {
	query := ` SELECT ID, Email FROM USERS`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
