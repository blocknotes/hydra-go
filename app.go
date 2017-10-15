package main

import (
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

func (app *App) init() {
	defer func() {
		if r := recover(); r != nil {
			// err = fmt.Errorf("create: %v", r)
		}
	}()

	validate = validator.New()

	session := app.DB.connect()
	defer session.Close()

	coll := session.DB(app.Conf.Database).C(app.Conf.Collection)
	var collections []Collection
	if ret := coll.Find(nil).All(&collections); ret != nil {
		panic(ret)
	}
	for _, c := range collections {
		app.Collections[c.Name] = c
	}

	authColl := session.DB(app.Conf.Database).C(app.Conf.AuthCollection)
	var authData []Auth
	if ret := authColl.Find(nil).All(&authData); ret != nil {
		panic(ret)
	}
	for _, c := range authData {
		app.AuthData[c.Username] = c.Password
	}

	app.routerSetup()
}

func (app *App) routerSetup() {
	authorized := app.Router.Group("/admin", gin.BasicAuth(gin.Accounts(app.AuthData)))
	authorized.GET("/collections", app.adminList)
	authorized.GET("/collections/:collection", app.adminRead)
	authorized.POST("/collections", app.adminCreate)
	authorized.PUT("/collections/:collection", app.adminUpdate)
	authorized.DELETE("/collections/:collection", app.adminDelete)
	// authorized.POST("/auth", app.adminAuthCreate) # TODO: create auth data
	// authorized.PUT("/auth/:username", app.adminAuthUpdate) # TODO: update auth data

	api := app.Router.Group("/api")
	api.GET("/:database/:collection", app.list)
	api.GET("/:database/:collection/:id", app.read)
	api.POST("/:database/:collection", app.create)
	api.PUT("/:database/:collection/:id", app.update)
	api.DELETE("/:database/:collection/:id", app.delete)
}
