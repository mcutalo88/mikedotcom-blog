package main

import (
	// 3rd Party
	"github.com/gin-gonic/gin"
	//"github.com/cactus/go-statsd-client/statsd"

	// Private
	"gitlab.azeroth.io/go-pkgs/az-logger.git"
	db "gitlab.azeroth.io/go-pkgs/az-mongo.git"

	// Local
	"blog-service/config"
	"blog-service/controllers/blogs"

	// Go
	"log"
)

var router *gin.Engine

func init() {
	config.ReadConfig("./config/config.toml")
	azlogger.SetLogger(config.Get().GraylogAddr)
	db.Connect(config.Get().Mongo_server, config.Get().Mongo_db)
}

func main() {
	log.Println("Starting WebServer ...")
  router = gin.Default()

	// Blog
	router.GET("/blogs", blogs.GetAllBlogs)
	// router.GET("/blogs/:id", getting)
	router.POST("/blogs", blogs.CreateBlog)
	// router.PUT("/bogs", putting)
	// router.DELETE("/blogs/:id", deleting)

  router.Run()
}
