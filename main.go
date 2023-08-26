package main

// go mod init <nombre del proyecto> => parecido a npm init -y
// go get -u github.com/gorilla/mux => para instalar una dependencia(no es obligatorio, go al compilar lo añadira solo a go.mod)
// go run . => ejecuta sin compilar
// go build . => para compilar primero
// gofmt -w controllers/user_controllers.go => para formatear los archivos a los estándares de go
// el nombre del package tiene que ser el nombre de la carpeta que contiene los archivos

import (
	"database/sql"
	"fmt"
	"go_rest_api/api/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Puntero que apunta a una instancia de la base de datos similar a el db de Sequelize
var db *sql.DB

func main() {

	db, errDB := sql.Open("mysql", "root:@tcp(127.0.0.1:3500)/testdb")

	if (errDB != nil) {
		log.Fatalf("Connection error, %v", errDB)
	}

	defer db.Close()

	var server *mux.Router = routes.OpenServer(db)

	fmt.Println("Server running on :3000")
	err := http.ListenAndServe(":3000", server)

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
