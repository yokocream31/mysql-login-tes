package main

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/greenteabiscuit/next-gin-mysql/backend/article"
	"github.com/greenteabiscuit/next-gin-mysql/backend/handler"
	"github.com/greenteabiscuit/next-gin-mysql/backend/lib"
	"github.com/greenteabiscuit/next-gin-mysql/backend/user"
	"github.com/joho/godotenv"

	"github.com/gin-contrib/cors"
)

func main() {
	if os.Getenv("USE_HEROKU") != "1" {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}

	article := article.New()
	user := user.New()

	lib.DBOpen()
	defer lib.DBClose()

	r := gin.Default()

	// フロントエンドの http://localhost:3000 からの通信は受け付ける
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	r.GET("/article", handler.ArticlesGet(article))
	r.POST("/article", handler.ArticlePost(article))
	// userはとりあえずlogin処理のみ
	r.GET("/user", handler.UsersGet(user))
	r.POST("/user/login", handler.UserPost(user))

	r.Run(os.Getenv("HTTP_HOST") + ":" + os.Getenv("HTTP_PORT")) // listen and serve on 0.0.0.0:8080
}
