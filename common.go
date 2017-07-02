package main

import (
	"gopkg.in/mgo.v2/bson"
)

type FormData struct {
	// Data string `form:"data" json:"data" binding:"required"`
	Data bson.M `form:"data" json:"data" binding:"required"`
}
