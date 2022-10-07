package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/greenteabiscuit/next-gin-mysql/backend/crypto"
	"github.com/greenteabiscuit/next-gin-mysql/backend/lib"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User モデルの宣言
type User struct {
	gorm.Model
	Username string `form:"username" json:"username" binding:"required" gorm:"unique;not null"`
	Password string `gorm:"size:511" json:"password" form:"password" binding:"required"`
}

// ユーザーを一件取得
func getUser(username string) User {
	db := lib.NewSQLHandler().DB
	var user User
	db.First(&user, "username = ?", username)
	sqlDB, _ := db.DB()

	defer sqlDB.Close()
	return user
}

func main() {

	//初期化処理
	db := lib.NewSQLHandler().DB
	sqlDB, _ := db.DB()
	db.AutoMigrate(&User{})
	defer sqlDB.Close()

	r := gin.Default()
	r.LoadHTMLGlob("views/*.html")

	// ユーザー登録画面
	r.GET("/signup", func(c *gin.Context) {
		c.JSON(200, gin.H{"result": "ok"})

		c.HTML(200, "signup.html", gin.H{})
	})

	// ユーザー登録
	r.POST("/signup", func(c *gin.Context) {
		var form User
		// バリデーション処理
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			username := c.PostForm("username")
			password := c.PostForm("password")

			//エラーが出てます!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
			fmt.Println(form.Username)
			fmt.Println(username)

			passwordEncrypt, _ := crypto.PasswordEncrypt(password)
			db := lib.NewSQLHandler().DB
			// Insert処理
			db.Create(&User{Username: form.Username, Password: passwordEncrypt})

			sqlDB, _ := db.DB()
			defer sqlDB.Close()
			c.JSON(200, gin.H{"result": "ok"})

			// c.Redirect(302, "/")
		}
	})

	// ユーザーログイン画面
	r.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{"result": "ok"})
		c.HTML(200, "login.html", gin.H{})
	})

	// ユーザーログイン
	r.POST("/login", func(c *gin.Context) {

		// DBから取得したユーザーパスワード(Hash)
		dbPassword := getUser(c.PostForm("username")).Password
		log.Println(dbPassword)
		// フォームから取得したユーザーパスワード
		formPassword := c.PostForm("password")
		fmt.Println("送られたパスワード")
		fmt.Println(dbPassword, formPassword)

		// ユーザーパスワードの比較
		err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(formPassword))
		if err != nil {
			log.Println("ログインできませんでした")
			c.HTML(http.StatusBadRequest, "login.html", gin.H{"err": err})
			c.Abort()
		} else {
			log.Println("ログインできました")
			c.JSON(200, gin.H{"result": "ok"})
			// c.Redirect(302, "/")
		}
	})
	r.Run()

}
