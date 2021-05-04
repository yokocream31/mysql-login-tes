package article

import (
	"fmt"

	"github.com/greenteabiscuit/next-gin-mysql/backend/lib"
)

type Article struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Articles struct {
	Items []Article
}

func New() *Articles {
	return &Articles{}
}

func (r *Articles) Add(a Article) {
	r.Items = append(r.Items, a)
	db := lib.GetDBConn().DB
	if err := db.Create(a).Error; err != nil {
		fmt.Println("err!")
	}
}

func (r *Articles) GetAll() []Article {
	db := lib.GetDBConn().DB
	var articles []Article
	if err := db.Find(&articles).Error; err != nil {
		return nil
	}
	return articles
}
