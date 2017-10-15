package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	validator "gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

// App - app
type App struct {
	Name     string
	Conf     Conf
	DB       DB
	AuthData map[string]string
	// Collections map[string]Collection
	Projects map[string]Project
	Router   *gin.Engine
}

// Auth - auth data
type Auth struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

// Conf - conf
type Conf struct {
	Database       string
	AuthCollection string
	Projects       string
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
	// Load projects
	prjs := session.DB(app.Conf.Database).C(app.Conf.Projects)
	var projects []Project
	if ret := prjs.Find(nil).All(&projects); ret != nil {
		panic(ret)
	}
	app.Projects = make(map[string]Project)
	for _, prj := range projects {
		app.Projects[prj.Name] = prj
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
	authorized.POST("/projects", app.AdminCreate)
	authorized.GET("/projects/:project", app.AdminRead)
	authorized.PUT("/projects/:project", app.AdminUpdate)
	authorized.DELETE("/projects/:project", app.AdminDelete)
	authorized.GET("/projects", app.AdminList)
	// authorized.POST("/auth", app.AdminAuthCreate) # TODO: create auth data
	// authorized.PUT("/auth/:username", app.AdminAuthUpdate) # TODO: update auth data
	// --- setup api routes
	api := app.Router.Group("/api")
	api.POST("/:project/:collection", app.Create)
	api.GET("/:project/:collection/:id", app.Read)
	api.PUT("/:project/:collection/:id", app.Update)
	api.DELETE("/:project/:collection/:id", app.Delete)
	api.GET("/:project/:collection", app.List)
}
