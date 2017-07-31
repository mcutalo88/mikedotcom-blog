package main

import (
	"strings"
	"strconv"

	"blog-service/config"
	"blog-service/controllers/blogs"
	models "blog-service/models"

	"gopkg.in/mgo.v2/bson"
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

/**
 * Generic Middleware to convert get params to mongo parsable structs
 *
 * Sort:
 * 	Syntax:
 * 		?sort=asc:attributeName
 * 	 	?sort=desc:attributeName
 *
 * Limit:
 * 	Syntax:
 * 	 ?limit=20
 *
 * Page AKA Skip:
 * 	Default behavior is always to get the first page. Notice index starts at 0.
 * 	Syntax:
 * 		?page=0
 *
 * Search:
 * 	Basicy regex search, you must specify what field to search against and then your query string.
 * 	Syntax:
 * 		?search=name:mikeischoice
 * 		?search=desc:FancyFeast
 *
 * Filter:
 * 	Basic top document filtering
 * 	Syntax:
 * 		?filter=type:post
 * 	 	?filter=type:pcbuild
 *
 * @type {[type]}
 */
func GenericPageFilterSearchSortLimit(c *gin.Context) {
	pager := models.Pager{}

	// Sort
	if c.Query("sort") != "" {
		for _, v := range strings.Split(c.Query("sort"), ",") {
			attributeSort := strings.Split(v, ":")

			if attributeSort[0] == "desc" {
				pager.Sort += "-" + attributeSort[1]
			} else {
				pager.Sort += attributeSort[1]
			}
		}
	} else {
		pager.Sort = "$_id"
	}

	// Limit
	if c.Query("limit") != "" {
		limit, _ := strconv.Atoi(c.Query("limit"))
		pager.Limit = int(limit)
	} else {
		pager.Limit = 25
	}

	// Page AKA Skip
	if c.Query("page") != "" {
		skip, _ := strconv.Atoi(c.Query("page"))
		pager.Skip = int(skip) * pager.Limit
	} else {
		pager.Skip = 0
	}

	// Search
	if c.Query("search") != "" {
		searchArgs := strings.Split(c.Query("search"), ":")
		pager.Search = bson.M{ searchArgs[0]: &bson.RegEx{Pattern: searchArgs[1], Options: "i"} }
	} else {
		pager.Search = nil
	}

	// Filter
	if c.Query("filter") != "" {
		filterArgs := strings.Split(c.Query("filter"), ":")
		pager.Filter = bson.M{ filterArgs[0]: filterArgs[1] }
	} else {
		pager.Filter = nil
	}

	c.Set("pager", pager)
	c.Next()
}

func main() {
  router = gin.Default()

	router.Use(CORSMiddleware)
	router.Use(GenericPageFilterSearchSortLimit)

	router.GET("/blogs", blogs.GetAllBlogs)
	router.GET("/blogs/:id", blogs.GetBlog)
	router.POST("/blogs", blogs.CreateBlog)
	router.PUT("/blogs/:id", blogs.UpdateBlog)
	// router.DELETE("/blogs/:id", deleting)

  router.Run()
}
