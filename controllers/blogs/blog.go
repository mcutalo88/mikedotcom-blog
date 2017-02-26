package blogs

import (
  // 3rd Party libs
  "gopkg.in/mgo.v2/bson"
  "github.com/gin-gonic/gin"

  // Private libs
  db "gitlab.azeroth.io/go-pkgs/az-mongo.git"

  // Go libs
  "log"
  "io/ioutil"
)

var bdoc interface{}
type Blog struct {
  Id      bson.ObjectId `json:"id" bson:"_id,omitempty"`
  Title   string        `json:"title"`
  Body    []string      `json:"body"`
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

func GetBlog(c *gin.Context) {
  blog := Blog{}
  err := db.Db.C("blogs").FindId(bson.ObjectIdHex(c.Param("id"))).One(&blog)

  if err != nil {
    c.JSON(500, err)
  } else {
    c.JSON(200, blog)
  }
}

// TODO: Atomic Insert plz ty
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
