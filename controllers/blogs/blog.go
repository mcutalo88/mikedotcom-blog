package blogs

import (
  "log"
  "io/ioutil"

  models "blog-service/models"

  "gopkg.in/mgo.v2/bson"
  "github.com/gin-gonic/gin"

  db "gitlab.azeroth.io/go-pkgs/az-mongo.git"
)

var bdoc interface{}

type Blog struct {
  Id    bson.ObjectId `json:"id" bson:"_id,omitempty"`
  Title string        `json:"title"`
  Desc  string        `json:"desc"`
  Body  []BlogBody    `json:"body"`
}

type BlogBody struct {
  Type string `json:"type"`
  Data string `json:"data"`
}

func GetAllBlogs(c *gin.Context) {
  pager := c.MustGet("pager").(models.Pager)
  // blogs := []Blog{}
  var maps []bson.M

  err := db.Db.C("blogs").
                Find(pager.Filter).
                Skip(pager.Skip).
                Sort(pager.Sort).
                Limit(pager.Limit).
                All(&maps)

  if err != nil {
    log.Println(err)
    c.JSON(500, err)
  } else {
    c.JSON(200, maps)
  }
}

func GetBlog(c *gin.Context) {
  // blog := Blog{}
  var maps bson.M
  // err := db.Db.C("blogs").FindId(bson.ObjectIdHex(c.Param("id"))).One(&blog)
  err := db.Db.C("blogs").FindId(bson.ObjectIdHex(c.Param("id"))).One(&maps)

  if err != nil {
    log.Println(err)
    c.JSON(500, err)
  } else {
    c.JSON(200, maps)
  }
}

// TODO: Atomic Insert plz ty
func CreateBlog(c *gin.Context) {
  // blog := Blog{}
  // c.BindJSON(&blog)
  // err := db.Db.C("blogs").Insert(&blog)

  var maps bson.M
  c.BindJSON(&maps)
  err := db.Db.C("blogs").Insert(&maps)

  if err != nil {
    log.Println(err)
    c.JSON(500, err)
  } else {
    c.Status(201)
  }
}

func UpdateBlog(c *gin.Context) {
  updateBody, _ := ioutil.ReadAll(c.Request.Body)
  bdocerror := bson.UnmarshalJSON([]byte(string(updateBody)), &bdoc)

  if bdocerror != nil {
    log.Println("Could not parse Request.Body")
    log.Println(bdocerror)
  }

  err := db.Db.C("blogs").Update(
    bson.M{"_id": bson.ObjectIdHex(c.Param("id"))},
    bson.M{"$set": &bdoc})

  if err != nil {
    log.Println(err)
    c.JSON(500, err)
  } else {
    c.Status(204)
  }
}
