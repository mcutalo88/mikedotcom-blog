package main

import (
	// 3rd Party libs
	"github.com/gin-gonic/gin"
	//"github.com/cactus/go-statsd-client/statsd"

	// Private Modules
	"gitlab.azeroth.io/go-pkgs/az-logger.git"

	// Local Modules
  "blog-service/config"
	"blog-service/controllers/blogs"

	// Go libs
	"log"
)

var router *gin.Engine

func init() {
	config.ReadConfig("./config/config.toml")
	azlogger.SetLogger(config.Get().GraylogAddr)
}

func main() {
	log.Println("Starting WebServer ...")
  router = gin.Default()

	// Blog
	router.GET("/blogs", blogs.GetAllBlogs)
	// router.GET("/blogs/:id", getting)
	// router.POST("/blogs", posting)
	// router.PUT("/bogs", putting)
	// router.DELETE("/blogs/:id", deleting)

  router.Run()
}
