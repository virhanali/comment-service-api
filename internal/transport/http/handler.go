package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler - stores pointer to comment services
type Handler struct{
	Router *mux.Router
}

// NewHandler - return a pointer to a handler
func NewHandler() *Handler{
	return &Handler{}
}


// Setup Routes - setups all the router for the application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting Up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Printf("I'am Alive")
	})
}

