package model

import (
	"ecommerceGo/internal/database"
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID
	Username string
	Email    string
	Password string
	IsAdmin  bool
}

func NewUser() *User {
	return &User{}
}

func (u User) Create() error {
	_, err := database.Db.Exec("INSERT INTO users(username, email, password, isAdmin) VALUE ($1, $2, $3, $4)", u.Username, u.Email, u.Password, false)
	if err != nil {
		return fmt.Errorf("error while creating the user")
	}
	return nil
}

func GetUserById(id uuid.UUID) (User, error) {
	u := *NewUser()
	err := database.Db.Get(&u, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return *NewUser(), fmt.Errorf("error while getting the user")
	}
	return u, nil
}

func GetUserByEmail(email string) (User, error) {
	u := *NewUser()
	err := database.Db.Get(&u, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return *NewUser(), fmt.Errorf("error while getting the user")
	}
	return u, nil
}
