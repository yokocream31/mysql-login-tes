package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greenteabiscuit/next-gin-mysql/backend/crypto"
	"github.com/greenteabiscuit/next-gin-mysql/backend/lib"
	"gorm.io/gorm"
)

// User モデルの宣言
type User struct {
	gorm.Model
	Username string `form:"username" binding:"required" gorm:"unique;not null"`
	Password string `form:"password" binding:"required"`
}

var db *lib.SQLHandler

// ユーザー登録処理
func createUser(username string, password string) {

	var err error
	passwordEncrypt, _ := crypto.PasswordEncrypt(password)
	// db := lib.NewSQLHandler()
	defer lib.DBClose()
	user := User{Username: username, Password: passwordEncrypt}
	fmt.Println(user)
	fmt.Println(user.Password, user.Username)
	fmt.Println("ss")
	// Insert処理
	if err = db.DB.Create(&user).Error; err != nil {
		fmt.Println("ssnoato")
		// return err
	}
	fmt.Println("nilnomae")
	// return nil
}

// ユーザーを一件取得
func getUser(username string) User {
	// db := lib.NewSQLHandler()
	var user User
	db.DB.First(&user, "username = ?", username)
	defer lib.DBClose()
	return user
}

func main() {

	// article := article.New()
	// user := user.New()
	db = lib.NewSQLHandler()
	defer lib.DBClose()

	db.DB.AutoMigrate(&User{})

	r := gin.Default()
	r.LoadHTMLGlob("views/*.html")
	// ユーザー登録画面
	r.GET("/signup", func(c *gin.Context) {

		c.HTML(200, "signup.html", gin.H{})
	})

	// ユーザー登録
	r.POST("/signup", func(c *gin.Context) {
		var form User
		var err error
		// バリデーション処理
		if err = c.Bind(&form); err != nil {
			fmt.Println("if desu")
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{"err": err})
			c.Abort()
		} else {
			fmt.Println("else desu")
			username := c.PostForm("username")
			password := c.PostForm("password")
			// 登録ユーザーが重複していた場合にはじく処理
			//var err error
			passwordEncrypt, _ := crypto.PasswordEncrypt(password)
			db := lib.NewSQLHandler()
			defer lib.DBClose()
			user := User{Username: username, Password: passwordEncrypt}
			fmt.Println(user)
			fmt.Println(user.Password, user.Username)
			fmt.Println("ss")
			// Insert処理
			if err = db.DB.Create(&user).Error; err != nil {
				fmt.Println("ssnoato")
				// return err
			}
			fmt.Println("nilnomae")
			//createUser(username, password)
			// if err = createUser(username, password); err != nil {
			// 	fmt.Println("err detemasu")
			// 	c.HTML(http.StatusBadRequest, "signup.html", gin.H{"err": err})
			// }
			fmt.Println("signup seiko desu")
			// c.Redirect(302, "/")
		}
	})

	// ユーザーログイン画面
	r.GET("/login", func(c *gin.Context) {

		c.HTML(200, "login.html", gin.H{})
	})
	// ユーザーログイン
	r.POST("/login", func(c *gin.Context) {

		// DBから取得したユーザーパスワード(Hash)
		dbPassword := getUser(c.PostForm("username")).Password
		log.Println(dbPassword)
		// フォームから取得したユーザーパスワード
		formPassword := c.PostForm("password")

		// ユーザーパスワードの比較
		if err := crypto.CompareHashAndPassword(dbPassword, formPassword); err != nil {
			log.Println("ログインできませんでした")
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"err": err})
			c.Abort()
		} else {
			log.Println("ログインできました")
			c.Redirect(302, "/")
		}
	})

	// r.GET("/article", handler.ArticlesGet(article))
	// r.POST("/article", handler.ArticlePost(article))
	// // userはとりあえずlogin処理のみ
	// r.GET("/user", handler.UsersGet(user))
	// r.POST("/user/login", handler.UserPost(user))

	// r.Run(os.Getenv("HTTP_HOST") + ":" + os.Getenv("HTTP_PORT")) // listen and serve on 0.0.0.0:8080
	r.Run()

}
