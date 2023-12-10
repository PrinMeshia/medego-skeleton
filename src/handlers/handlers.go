package handlers

import (
	"myapp/src/data"
	"net/http"
	"github.com/PrinMeshia/medego"
)

// Handlers is the type for handlers, and gives access to Medego and models
type Handlers struct {
	App    *medego.Medego
	Models data.Models
}

// Home is the handler to render the home page
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.render(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

