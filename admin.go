package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// curl 'http://localhost:8080/admin/collections'
func (app App) adminList(c *gin.Context) {
	data, code, err := app.DB.findAll(app.Conf.Database, app.Conf.Collection)
	if err != nil {
		// fmt.Println(err) // TODO: log me !
		c.JSON(code, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
	} else {
		c.JSON(code, data)
	}
}

// curl 'http://localhost:8080/admin/collections/articles'
func (app App) adminRead(c *gin.Context) {
	appCollection := c.Param("collection")
	data, code, err := app.DB.findOne(app.Conf.Database, app.Conf.Collection, "name", appCollection)
	if err != nil {
		// fmt.Println(err) // TODO: log me !
		c.JSON(code, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
	} else {
		// fmt.Println(data)
		c.JSON(code, data)
	}
}

// curl -X POST 'http://localhost:8080/admin/collections' -H 'Content-Type: application/json' --data '{"data":{"name":"articles","singular":"article","columns":{"title":{"type":"String"},"email":{"type":"String","validations":"required,email"},"description":{"type":"String"},"position":{"type":"Float"},"published":{"type":"Boolean"},"dt":{"type":"DateTime"}}}}'
func (app App) adminCreate(c *gin.Context) {
	var data FormData
	if c.BindJSON(&data) == nil {
		data, code, err := app.DB.create(app.Conf.Database, app.Conf.Collection, data)
		if err != nil {
			// fmt.Println(err) // TODO: log me !
			c.JSON(code, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
		} else {
			c.JSON(code, data)
		}
	}
	// TODO: update app Collections
}

// curl -X PUT 'http://localhost:8080/admin/collections/articles' -H 'Content-Type: application/json' --data '{"data":{"name":"articles","columns":{"subtitle":{"type":"String"},"email":{"type":"String","validations":"required,email"}}}}'
func (app App) adminUpdate(c *gin.Context) {
	appCollection := c.Param("collection")
	data, code, err := app.DB.findOne(app.Conf.Database, app.Conf.Collection, "name", appCollection)
	if err != nil {
		// fmt.Println(err) // TODO: log me !
		c.JSON(code, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
	} else {
		id := data["_id"].(bson.ObjectId).Hex()
		var data2 FormData
		if c.BindJSON(&data2) == nil {
			code, err := app.DB.updateByID(app.Conf.Database, app.Conf.Collection, id, data2.Data)
			if err != nil {
				c.JSON(code, gin.H{
					"status":  "error",
					"message": err.Error(),
				})
			} else {
				c.JSON(code, gin.H{
					"status": "ok",
				})
			}
		}
		// fmt.Println(data2)
	}
	// TODO: update app Collections
}

// curl -X DELETE 'http://localhost:8080/admin/collections/articles'
func (app App) adminDelete(c *gin.Context) {
	appCollection := c.Param("collection")
	data, code, err := app.DB.findOne(app.Conf.Database, app.Conf.Collection, "name", appCollection)
	if err != nil {
		// fmt.Println(err) // TODO: log me !
		c.JSON(code, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
	} else {
		id := data["_id"].(bson.ObjectId).Hex()
		code, err := app.DB.deleteByID(app.Conf.Database, app.Conf.Collection, id)
		if err != nil {
			c.JSON(code, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
		} else {
			c.JSON(code, "ok")
		}
	}
	// TODO: update app Collections
}
