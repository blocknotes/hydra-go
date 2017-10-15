package main

func main() {
	// App setup
	app := App{
		Name: "Hydra Go",
		Conf: Conf{
			Database:       "_hydra",
			Collection:     "_collections",
			AuthCollection: "_auth",
		},
		DB: DB{"localhost:27017"},
	}
	// App run
	app.Init()
	app.RouterSetup()
	app.Router.Run("localhost:8080")
	// app.router.Run()  // listen and serve on 0.0.0.0:8080
	// app.router.RunTLS( ":8080", "./server.pem", "./server.key" ) // listen and serve on 0.0.0.0:8080
}
