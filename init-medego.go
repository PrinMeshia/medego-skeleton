package main

import (
	"log"
	"myapp/src/data"
	"myapp/src/handlers"
	"myapp/src/middleware"
	"os"

	"github.com/PrinMeshia/medego"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//init medego
	cel := &medego.Medego{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "MyApp"

	myMiddleware := &middleware.Middleware{
		App: cel,
	}

	myHandler := &handlers.Handlers{
		App: cel,
	}

	app := &application{
		App:        cel,
		Middleware: myMiddleware,
		Handlers:   myHandler,
	}

	app.App.Routes = app.routes()
	app.Models = data.New(app.App.DB.Pool)
	myHandler.Models = app.Models
	app.Middleware.Models = app.Models

	return app

}
