package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func adminSetup() *App {
	gin.SetMode(gin.ReleaseMode)
	app := App{
		"Hydra Go",
		Conf{
			"_hydra",
			"_collections",
			"_auth",
		},
		DB{"localhost:27017"},
		map[string]string{},
		map[string]Collection{},
		gin.Default(),
	}
	app.init()
	return &app
}

func TestAdminCreate(t *testing.T) {
	app := adminSetup()
	data := `{"data":{"name":"posts","singular":"post","columns":{"title":{"type":"String","validations":"required"},"description":{"type":"String"},"position":{"type":"Float"},"published":{"type":"Boolean"},"dt":{"type":"DateTime"}}}}`
	body := bytes.NewBufferString(data)
	req, _ := http.NewRequest("POST", "/admin/collections", body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("mat:123")))
	res := httptest.NewRecorder()
	app.Router.ServeHTTP(res, req)
	// check status
	assert.Equal(t, res.Code, 200)
}

func TestAdminIndex(t *testing.T) {
	app := adminSetup()
	req, _ := http.NewRequest("GET", "/admin/collections", nil)
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("mat:123")))
	res := httptest.NewRecorder()
	app.Router.ServeHTTP(res, req)
	// check status
	assert.Equal(t, res.Code, 200)
	// check content
	var data []Collection
	json.Unmarshal(res.Body.Bytes(), &data)
	assert.True(t, len(data) > 0)
	// assert.Equal(t, data[0].Name, "posts")
}

func TestAdminUpdate(t *testing.T) {
	app := adminSetup()
	data := `{"data":{"description":"A test..."}}`
	body := bytes.NewBufferString(data)
	req, _ := http.NewRequest("PUT", "/admin/collections/posts", body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("mat:123")))
	res := httptest.NewRecorder()
	app.Router.ServeHTTP(res, req)
	// check status
	assert.Equal(t, res.Code, 200)
}

func TestAdminRead(t *testing.T) {
	app := adminSetup()
	req, _ := http.NewRequest("GET", "/admin/collections/posts", nil)
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("mat:123")))
	res := httptest.NewRecorder()
	app.Router.ServeHTTP(res, req)
	// check status
	assert.Equal(t, res.Code, 200)
	// check content
	var data Collection
	json.Unmarshal(res.Body.Bytes(), &data)
	assert.Equal(t, data.Name, "posts")
	assert.Equal(t, data.Description, "A test...")
}

func TestAdminDestroy(t *testing.T) {
	app := adminSetup()
	req, _ := http.NewRequest("DELETE", "/admin/collections/posts", nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte("mat:123")))
	res := httptest.NewRecorder()
	app.Router.ServeHTTP(res, req)
	// check status
	assert.Equal(t, res.Code, 200)
}
