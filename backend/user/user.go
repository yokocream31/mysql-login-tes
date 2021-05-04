package user

import (
	"fmt"

	"github.com/greenteabiscuit/next-gin-mysql/backend/lib"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Users struct {
	Items []User
}

func New() *Users {
	return &Users{}
}

func (r *Users) Add(a User) {
	r.Items = append(r.Items, a)
	db := lib.GetDBConn().DB
	if err := db.Create(a).Error; err != nil {
		fmt.Println("err!")
	}
}

func (r *Users) GetAll() []User {
	db := lib.GetDBConn().DB
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil
	}
	return users
}
