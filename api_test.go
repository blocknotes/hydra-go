package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var id string

func apiSetup() *App {
	gin.SetMode(gin.ReleaseMode)
	app := App{
		Name: "Hydra Go",
		Conf: Conf{
			Database:       "_hydra",
			Collection:     "_collections",
			AuthCollection: "_auth",
		},
		DB:     DB{"localhost:27017"},
		Router: gin.Default(),
	}
	app.Init()
	app.RouterSetup()
	return &app
}

func TestApiCreate(t *testing.T) {
	app := apiSetup()
	data := `{"data":{"title":"A test..."}}`
	body := bytes.NewBufferString(data)
	req, _ := http.NewRequest("POST", "/api/articles/articles", body)
	req.Header.Add("Content-Type", "application/json")
	res := httptest.NewRecorder()
	app.Router.ServeHTTP(res, req)
	// check status
	assert.Equal(t, res.Code, 200)
	// check content
	var ret bson.M
	json.Unmarshal(res.Body.Bytes(), &ret)
	// // result := ret["result"].(bson.M)
	id = ret["id"].(string)
	// fmt.Println(id)
}

func TestApiUpdate(t *testing.T) {
	app := apiSetup()
	data := `{"data":{"title":"A test!"}}`
	body := bytes.NewBufferString(data)
	req, _ := http.NewRequest("PUT", "/api/articles/articles/"+id, body)
	req.Header.Add("Content-Type", "application/json")
	res := httptest.NewRecorder()
	app.Router.ServeHTTP(res, req)
	// check status
	assert.Equal(t, res.Code, 200)
}

func TestApiIndex(t *testing.T) {
	app := apiSetup()
	req, _ := http.NewRequest("GET", "/api/articles/articles", nil)
	res := httptest.NewRecorder()
	app.Router.ServeHTTP(res, req)
	// check status
	assert.Equal(t, res.Code, 200)
	// check content
	var data []Collection
	json.Unmarshal(res.Body.Bytes(), &data)
	assert.True(t, len(data) > 0)
	// fmt.Println(res.Body.String())
}

func TestApiRead(t *testing.T) {
	app := apiSetup()
	req, _ := http.NewRequest("GET", "/api/articles/articles/"+id, nil)
	res := httptest.NewRecorder()
	app.Router.ServeHTTP(res, req)
	// check status
	assert.Equal(t, res.Code, 200)
	// check content
	var data bson.M
	json.Unmarshal(res.Body.Bytes(), &data)
	assert.Equal(t, data["title"], "A test!")
	// fmt.Println(data)
}

func TestApiDestroy(t *testing.T) {
	app := apiSetup()
	req, _ := http.NewRequest("DELETE", "/api/articles/articles/"+id, nil)
	req.Header.Add("Content-Type", "application/json")
	res := httptest.NewRecorder()
	app.Router.ServeHTTP(res, req)
	// check status
	assert.Equal(t, res.Code, 200)
}
