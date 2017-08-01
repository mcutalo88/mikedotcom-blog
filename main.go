package main

import (
	"blog-service/config"
	"blog-service/controllers/blogs"
	middleware "blog-service/middleware"

	"github.com/gin-gonic/gin"
	"gitlab.azeroth.io/go-pkgs/az-logger.git"
	db "gitlab.azeroth.io/go-pkgs/az-mongo.git"
)

var router *gin.Engine

func init() {
	config.ReadConfig("./config/config.toml")
	azlogger.SetLogger(config.Get().GraylogAddr)
	db.Connect(config.Get().Mongo_server, config.Get().Mongo_db)
}

func main() {
  router = gin.Default()

	router.Use(middleware.CORSMiddleware)
	router.Use(middleware.GenericPageFilterSearchSortLimit)

	router.GET("/blogs", blogs.GetAllBlogs)
	router.GET("/blogs/:id", blogs.GetBlog)
	router.POST("/blogs", blogs.CreateBlog)
	router.PUT("/blogs/:id", blogs.UpdateBlog)
	// router.DELETE("/blogs/:id", deleting)

  router.Run()
}
