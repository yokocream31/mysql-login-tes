package handler

import (
	"net/http"

	"github.com/greenteabiscuit/next-gin-mysql/backend/user"

	"github.com/gin-gonic/gin"
)

func UsersGet(users *user.Users) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := users.GetAll()
		c.JSON(http.StatusOK, result)
	}
}

type UserPostRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func UserPost(post *user.Users) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := UserPostRequest{}
		c.Bind(&requestBody)

		item := user.User{
			Username: requestBody.Username,
			Password: requestBody.Password,
		}
		post.Add(item)

		c.Status(http.StatusNoContent)
	}
}
