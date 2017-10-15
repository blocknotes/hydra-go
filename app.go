package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	validator "gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

// App - app
type App struct {
	Name        string
	Conf        Conf
	DB          DB
	AuthData    map[string]string
	Collections map[string]Collection
	Router      *gin.Engine
}

// Auth - auth data
type Auth struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

// Conf - conf
type Conf struct {
	Database       string
	Collection     string
	AuthCollection string
}

func (app *App) Init() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("App-Init: %v", r)
		}
	}()
	validate = validator.New()
	session := app.DB.Connect()
	defer session.Close()
	// Load collections
	coll := session.DB(app.Conf.Database).C(app.Conf.Collection)
	var collections []Collection
	if ret := coll.Find(nil).All(&collections); ret != nil {
		panic(ret)
	}
	app.Collections = make(map[string]Collection)
	for _, c := range collections {
		app.Collections[c.Name] = c
	}
	// Load auth data
	authColl := session.DB(app.Conf.Database).C(app.Conf.AuthCollection)
	var authData []Auth
	if ret := authColl.Find(nil).All(&authData); ret != nil {
		panic(ret)
	}
	app.AuthData = make(map[string]string)
	for _, c := range authData {
		app.AuthData[c.Username] = c.Password
	}

	return
}

func (app *App) RouterSetup() {
	app.Router = gin.Default()
	// --- setup admin routes
	authorized := app.Router.Group("/admin", gin.BasicAuth(gin.Accounts(app.AuthData)))
	authorized.GET("/collections", app.AdminList)
	authorized.GET("/collections/:collection", app.AdminRead)
	authorized.POST("/collections", app.AdminCreate)
	authorized.PUT("/collections/:collection", app.AdminUpdate)
	authorized.DELETE("/collections/:collection", app.AdminDelete)
	// authorized.POST("/auth", app.AdminAuthCreate) # TODO: create auth data
	// authorized.PUT("/auth/:username", app.AdminAuthUpdate) # TODO: update auth data
	// --- setup api routes
	api := app.Router.Group("/api")
	api.GET("/:database/:collection", app.List)
	api.GET("/:database/:collection/:id", app.Read)
	api.POST("/:database/:collection", app.Create)
	api.PUT("/:database/:collection/:id", app.Update)
	api.DELETE("/:database/:collection/:id", app.Delete)
}
