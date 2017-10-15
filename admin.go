package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// curl -X POST 'http://127.0.0.1:8080/admin/projects' -H 'Content-Type: application/json' --data '{"data":{"name":"hydra1","description":"Hydra 1","collections":[]}}'
func (app App) AdminCreate(c *gin.Context) {
	var data FormData
	if c.BindJSON(&data) == nil {
		data, code, err := app.DB.Create(app.Conf.Database, app.Conf.Projects, data)
		if err != nil {
			// fmt.Println(err) // TODO: log me !
			c.JSON(code, gin.H{
				"message": err.Error(),
				"status":  "error",
			})
		} else {
			c.JSON(code, data)
		}
	} else {
		// fmt.Println(err) // TODO: log me !
		c.JSON(400, gin.H{
			"message": "Bad request",
			"status":  "error",
		})
	}
	// TODO: update app Collections
}

// curl 'http://127.0.0.1:8080/admin/projects/hydra1'
func (app App) AdminRead(c *gin.Context) {
	appProject := c.Param("project")
	data, code, err := app.DB.FindOne(app.Conf.Database, app.Conf.Projects, "name", appProject)
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

// curl -X PUT 'http://127.0.0.1:8080/admin/projects/hydra1' -H 'Content-Type: application/json' --data '{"data":{"collections":[{"name":"articles","singular":"article","columns":{"title":{"type":"String"},"email":{"type":"String","validations":"required,email"},"description":{"type":"String"},"position":{"type":"Float"},"published":{"type":"Boolean"},"dt":{"type":"DateTime"}}}]}}'
func (app App) AdminUpdate(c *gin.Context) {
	appProject := c.Param("project")
	data, code, err := app.DB.FindOne(app.Conf.Database, app.Conf.Projects, "name", appProject)
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
			code, err := app.DB.UpdateByID(app.Conf.Database, app.Conf.Projects, id, data2.Data)
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
	// TODO: update app Collections
}

// curl -X DELETE 'http://127.0.0.1:8080/admin/projects/hydra1'
func (app App) AdminDelete(c *gin.Context) {
	appProject := c.Param("project")
	data, code, err := app.DB.FindOne(app.Conf.Database, app.Conf.Projects, "name", appProject)
	if err != nil {
		// fmt.Println(err) // TODO: log me !
		c.JSON(code, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
	} else {
		id := data["_id"].(bson.ObjectId).Hex()
		code, err := app.DB.DeleteByID(app.Conf.Database, app.Conf.Projects, id)
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
	// TODO: update app Collections
}

// curl 'http://127.0.0.1:8080/admin/projects'
func (app App) AdminList(c *gin.Context) {
	data, code, err := app.DB.FindAll(app.Conf.Database, app.Conf.Projects)
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
