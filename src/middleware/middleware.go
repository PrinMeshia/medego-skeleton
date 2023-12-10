package middleware

import (
	"myapp/src/data"

	"github.com/PrinMeshia/medego"
)

type Middleware struct {
	App    *medego.Medego
	Models data.Models
}
