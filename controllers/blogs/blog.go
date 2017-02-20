package blogs

import (
  // 3rd Party libs
  "gopkg.in/mgo.v2/bson"
  "github.com/gin-gonic/gin"

  // Private libs
  db "gitlab.azeroth.io/go-pkgs/az-mongo.git"

  // Go libs
  "log"
)

type Blog struct {
  Id      bson.ObjectId `json:"id" bson:"_id,omitempty"`
  Title   string        `json:"title"`
}

func GetAllBlogs(c *gin.Context) {
  blogs := []Blog{}

  err := db.Db.C("blogs").Find(nil).All(&blogs)

  if err != nil {
    log.Println(err)
    c.JSON(500, err)
  } else {
    c.JSON(200, blogs)
  }
}

func CreateBlog(c *gin.Context) {
  blog := Blog{}
  c.BindJSON(&blog)
  err := db.Db.C("blogs").Insert(&blog)

  if err != nil {
    log.Println(err)
    c.JSON(500, err)
  } else {
    c.Status(201)
  }
}
