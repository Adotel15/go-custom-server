package routes

import (
	"database/sql"
	"go_rest_api/api/controllers"

	"github.com/gorilla/mux"
)

func OpenServer(db *sql.DB) *mux.Router {
	var server *mux.Router = mux.NewRouter()

	Controllers := controllers.InitDB(db)

	// USER ROUTES
	server.HandleFunc("/users", Controllers.GetUsers).Methods("GET")
	server.HandleFunc("/users", Controllers.CreateUser).Methods("POST")

	return server
}
