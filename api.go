package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// curl http://localhost:8080/api/articles/articles
func (app App) List(c *gin.Context) {
	database := c.Param("database")
	collection := c.Param("collection")
	data, code, err := app.DB.FindAll(database, collection)
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

// curl http://localhost:8080/api/articles/articles/59468245cfba25329f3272db
func (app App) Read(c *gin.Context) {
	database := c.Param("database")
	collection := c.Param("collection")
	id := c.Param("id")
	data, code, err := app.DB.FindByID(database, collection, id)
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

// curl -X POST http://localhost:8080/api/articles/articles -H 'Content-Type: application/json' --data '{"data":{"title":"A test"}}'
func (app App) Create(c *gin.Context) {
	database := c.Param("database")
	collection := c.Param("collection")
	var data FormData
	if c.BindJSON(&data) == nil {
		ret, code, err := app.DB.Create(database, collection, data)
		if err != nil {
			// fmt.Println(err) // TODO: log me !
			c.JSON(code, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
		} else {
			c.JSON(code, gin.H{
				"status": "ok",
				"id":     ret["_id"],
			})
		}
	}
}

// curl -X PUT http://localhost:8080/api/articles/articles/59468245cfba25329f3272db -H 'Content-Type: application/json' --data '{"data":{"title":"A test 2"}}'
func (app App) Update(c *gin.Context) {
	database := c.Param("database")
	collection := c.Param("collection")
	id := c.Param("id")
	data, code, err := app.DB.FindByID(database, collection, id)
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
			if errs := app.Collections[collection].Validate(data2.Data); errs != nil {
				// fmt.Println(errs)
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "error",
					"message": errs[0].Error(), // TODO: list of errors!
				})
			} else {
				code, err := app.DB.UpdateByID(database, collection, id, data2.Data)
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
		}
	}
}

// curl -X DELETE http://localhost:8080/api/books/books/59462836cfba25329f3272d0
func (app App) Delete(c *gin.Context) {
	database := c.Param("database")
	collection := c.Param("collection")
	id := c.Param("id")
	data, code, err := app.DB.FindByID(database, collection, id)
	if err != nil {
		// fmt.Println(err) // TODO: log me !
		c.JSON(code, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
	} else {
		id := data["_id"].(bson.ObjectId).Hex()
		code, err := app.DB.DeleteByID(database, collection, id)
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
}
