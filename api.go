package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// curl -X POST 'http://127.0.0.1:8080/api/hydra1/articles' -H 'Content-Type: application/json' --data '{"data":{"title":"A test"}}'
func (app App) Create(c *gin.Context) {
	project := c.Param("project")
	collection := c.Param("collection")
	var data FormData
	if c.BindJSON(&data) == nil {
		ret, code, err := app.DB.Create(project, collection, data)
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

// curl 'http://127.0.0.1:8080/api/hydra1/articles/59e3b9b7d23d8028efd327c4'
func (app App) Read(c *gin.Context) {
	project := c.Param("project")
	collection := c.Param("collection")
	id := c.Param("id")
	data, code, err := app.DB.FindByID(project, collection, id)
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

// curl -X PUT 'http://127.0.0.1:8080/api/hydra1/articles/59e3b9b7d23d8028efd327c4' -H 'Content-Type: application/json' --data '{"data":{"title":"A test 2"}}'
func (app App) Update(c *gin.Context) {
	project := c.Param("project")
	collection := c.Param("collection")
	id := c.Param("id")
	data, code, err := app.DB.FindByID(project, collection, id)
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
			/// TODO: fix validations
			// if errs := app.Collections[collection].Validate(data2.Data); errs != nil {
			// 	// fmt.Println(errs)
			// 	c.JSON(http.StatusBadRequest, gin.H{
			// 		"status":  "error",
			// 		"message": errs[0].Error(), // TODO: list of errors!
			// 	})
			// } else {
			code, err := app.DB.UpdateByID(project, collection, id, data2.Data)
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
			// }
		}
	}
}

// curl -X DELETE 'http://127.0.0.1:8080/api/hydra1/articles/59e3b9b7d23d8028efd327c4'
func (app App) Delete(c *gin.Context) {
	project := c.Param("project")
	collection := c.Param("collection")
	id := c.Param("id")
	data, code, err := app.DB.FindByID(project, collection, id)
	if err != nil {
		// fmt.Println(err) // TODO: log me !
		c.JSON(code, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
	} else {
		id := data["_id"].(bson.ObjectId).Hex()
		code, err := app.DB.DeleteByID(project, collection, id)
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

// curl 'http://127.0.0.1:8080/api/hydra1/articles'
func (app App) List(c *gin.Context) {
	project := c.Param("project")
	collection := c.Param("collection")
	data, code, err := app.DB.FindAll(project, collection)
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
