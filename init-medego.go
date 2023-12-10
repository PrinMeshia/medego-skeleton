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
	core := &medego.Medego{}
	err = core.New(path)
	if err != nil {
		log.Fatal(err)
	}

	core.AppName = "MyApp"

	myMiddleware := &middleware.Middleware{
		App: core,
	}

	myHandler := &handlers.Handlers{
		App: core,
	}

	app := &application{
		App:        core,
		Middleware: myMiddleware,
		Handlers:   myHandler,
	}

	app.App.Routes = app.routes()
	app.Models = data.New(app.App.DB.Pool)
	myHandler.Models = app.Models
	app.Middleware.Models = app.Models

	return app

}
