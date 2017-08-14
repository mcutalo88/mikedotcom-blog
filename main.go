package main

import (
	"os"
	"log"
	"blog-service/config"
	"blog-service/controllers/blogs"
	middleware "blog-service/middleware"

	"github.com/gin-gonic/gin"
	// "gitlab.azeroth.io/go-pkgs/az-logger.git"
	db "gitlab.azeroth.io/go-pkgs/az-mongo.git"
)

var router *gin.Engine

func init() {
	// azlogger.SetLogger(config.Get().GraylogAddr)
	if err := config.Load(); err != nil {
			log.Fatal(err)
	}

	db.Connect(config.Vip.GetString("mongo.server"), config.Vip.GetString("mongo.db"))
}

func main() {
  router = gin.Default()

	if os.Getenv("GO_ENV") != "" && os.Getenv("GO_ENV") == "dev" {
		router.Use(middleware.CORSMiddleware)
	}

	router.Use(middleware.GenericPageFilterSearchSortLimit)

	router.GET("/blogs", blogs.GetAllBlogs)
	router.GET("/blogs/:id", blogs.GetBlog)
	router.POST("/blogs", blogs.CreateBlog)
	router.PUT("/blogs/:id", blogs.UpdateBlog)
	// router.DELETE("/blogs/:id", deleting)

  router.Run()
}
