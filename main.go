package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// App setup
	app := App{
		Name: "Hydra Go",
		Conf: Conf{
			"_hydra",
			"_collections",
			"_auth",
		},
		DB:          DB{"localhost:27017"},
		AuthData:    map[string]string{},
		Collections: map[string]Collection{},
		Router:      gin.Default(),
	}
	// App run
	app.init()
	app.Router.Run("localhost:8080")
	// app.router.Run()  // listen and serve on 0.0.0.0:8080
	// app.router.RunTLS( ":8080", "./server.pem", "./server.key" ) // listen and serve on 0.0.0.0:8080
}
