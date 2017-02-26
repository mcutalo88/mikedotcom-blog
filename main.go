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
)

var router *gin.Engine

func init() {
	config.ReadConfig("./config/config.toml")
	azlogger.SetLogger(config.Get().GraylogAddr)
	db.Connect(config.Get().Mongo_server, config.Get().Mongo_db)
}

func CORSMiddleware(c *gin.Context) {
  c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
  c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
  c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
  c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

  if c.Request.Method == "OPTIONS" {
      c.AbortWithStatus(204)
      return
  }

  c.Next()
}

func main() {
  router = gin.Default()
	router.Use(CORSMiddleware)

	router.GET("/blogs", blogs.GetAllBlogs)
	router.GET("/blogs/:id", blogs.GetBlog)
	router.POST("/blogs", blogs.CreateBlog)
	router.PUT("/blogs/:id", blogs.UpdateBlog)
	// router.DELETE("/blogs/:id", deleting)

  router.Run()
}
