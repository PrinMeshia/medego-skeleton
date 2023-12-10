package main

import (
	"myapp/src/data"
	"myapp/src/handlers"
	"myapp/src/middleware"

	"github.com/PrinMeshia/medego"
)

type application struct {
	App        *medego.Medego
	Middleware *middleware.Middleware
	Handlers   *handlers.Handlers
	Models     data.Models
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
