package models

import (
  "gopkg.in/mgo.v2/bson"
)

type Pager struct {
	Sort   string
  Limit  int
  Skip   int
  Search bson.M
  Filter bson.M
}
