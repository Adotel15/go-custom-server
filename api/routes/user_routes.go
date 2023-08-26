package routes

import (
	"go_rest_api/api/controllers"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	var server *mux.Router = mux.NewRouter()
	server.HandleFunc("/users", controllers.GetUsers).Methods("GET")
	return server
}
