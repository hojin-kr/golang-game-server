package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const port string = ":8888"

var db *sql.DB

// User basic info
type User struct {
	ID         int    `json:"id"`
	PlatformID string `json:"platform_id" binding:"required"`
	Platform   string `json:"platform" binding:"required"`
}

// Login user return game id
func Login(c *gin.Context) {
	var user User
	if err := c.ShouldBind(&user); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}
	row := db.QueryRow("SELECT id, platform_id, platform from user where platform_id = ? AND platform = ?", user.PlatformID, user.Platform)
	row.Scan(&user.ID, &user.Platform, &user.PlatformID)
	if !(user.ID > 0) {
		// sign up
		rs, err := db.Exec("INSERT INTO user(platform_id, platform) VALUES (?, ?)", user.PlatformID, user.Platform)
		if err != nil {
			log.Fatalln(err)
		}
		id, err := rs.LastInsertId()
		if err != nil {
			log.Fatalln(err)
		}
		user.ID = int(id)
	}
	c.JSON(http.StatusOK, user)
}

func main() {
	// database
	var err error
	db, err = sql.Open("mysql", "dev:dev@tcp(127.0.0.1:3306)/user")
	if err != nil {
		log.Println(err)
	}
	// main() 종료되면 디비 연결 헤제 / Then finished main close db connet
	defer db.Close()
	//router
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// sign up or sign in return game id
	r.POST("/login", Login)
	r.Run(port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
