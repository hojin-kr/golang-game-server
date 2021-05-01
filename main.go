package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/my/repo/models"
)

var sqlDBClient *sql.DB
var redisClient *redis.Client

const port string = ":8888"

func main() {
	// sqldb
	var err error
	sqlDBClient, err = sql.Open("mysql", "app:1q2w3e4r@tcp(rdb:3306)/user")
	if err != nil {
		log.Println(err)
	}
	defer sqlDBClient.Close()
	// redis
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	// router
	r := gin.Default()
	r.LoadHTMLFiles("templates/index.html")
	r.Static("/static", "./static")
	r.GET("/leaderboard", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Leaderboard",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Login user return game id
	r.POST("/user/login", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBind(&user); err != nil {
			c.String(http.StatusBadRequest, "bad request")
			return
		}
		row := sqlDBClient.QueryRow("SELECT id, platform_id, platform, device_id from user where platform_id = ? AND platform = ?", user.PlatformID, user.Platform)
		row.Scan(&user.ID, &user.PlatformID, &user.Platform, &user.DeviceID)
		if user.ID == 0 {
			// sign up
			rs, err := sqlDBClient.Exec("INSERT INTO user(platform_id, platform, device_id) VALUES (?, ?, ?)", user.PlatformID, user.Platform, user.DeviceID)
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
	})
	// ChangePlatfrom change platform
	r.POST("/user/changeplatfrom", func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBind(&user); err != nil {
			c.String(http.StatusBadRequest, "bad request")
			return
		}
		if user.ID != 0 {
			rs, err := sqlDBClient.Exec("UPDATE user SET platform_id = ?, platform = ? WHERE id = ? AND device_id = ?", user.PlatformID, user.Platform, user.ID, user.DeviceID)
			if err != nil || rs == nil {
				c.JSON(http.StatusBadRequest, err)
				return
			}
		}
		c.JSON(http.StatusOK, user)
	})
	// stage start
	r.POST("/stage/start", func(c *gin.Context) {
		var stage models.Stage
		if err := c.ShouldBind(&stage); err != nil {
			c.String(http.StatusBadRequest, "bad request")
			return
		}
		err := redisClient.ZIncrBy("stage:try", 1, strconv.Itoa(stage.ID)).Err()
		if err != nil {
			panic(err)
		}
		log.Println(stage.ID)
		tryStages, err := redisClient.ZRevRangeWithScores("stage:try", 0, -1).Result()
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, tryStages)
	})
	// stage clear
	r.POST("/stage/clear", func(c *gin.Context) {
		var stage models.Stage
		if err := c.ShouldBind(&stage); err != nil {
			c.String(http.StatusBadRequest, "bad request")
			return
		}
		log.Println(stage.ID)
		err := redisClient.ZIncrBy("stage:clear", 1, strconv.Itoa(stage.ID)).Err()
		if err != nil {
			panic(err)
		}
		clearStages, err := redisClient.ZRevRangeWithScores("stage:clear", 0, -1).Result()
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, clearStages)
	})
	// stage info get
	r.POST("/stage/get", func(c *gin.Context) {
		var stage models.Stage
		if err := c.ShouldBind(&stage); err != nil {
			c.String(http.StatusBadRequest, "bad request")
			return
		}
		try, err := redisClient.ZScore("stage:try", strconv.Itoa(stage.ID)).Result()
		if err != nil {
			try = 0
		}
		clear, err := redisClient.ZScore("stage:clear", strconv.Itoa(stage.ID)).Result()
		if err != nil {
			clear = 0
		}
		stage.Try = try
		stage.Clear = clear
		c.JSON(http.StatusOK, stage)
	})
	// user rank incr
	r.POST("/rank/incr", func(c *gin.Context) {
		var rank models.Rank
		if err := c.ShouldBind(&rank); err != nil {
			c.String(http.StatusBadRequest, "bad request")
			return
		}
		err := redisClient.ZIncrBy("rank", rank.Score, strconv.Itoa(rank.ID)).Err()
		if err != nil {
			panic(err)
		}
		ranks, err := redisClient.ZRevRangeWithScores("rank", 0, 500).Result()
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, ranks)
	})
	// user rank get
	r.POST("/rank/get", func(c *gin.Context) {
		ranks, err := redisClient.ZRevRangeWithScores("rank", 0, 500).Result()
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, ranks)
	})
	r.Run(port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
