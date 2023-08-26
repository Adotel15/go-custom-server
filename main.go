package main

// go mod init <nombre del proyecto> => parecido a npm init -y
// go get -u github.com/gorilla/mux => para instalar una dependencia(no es obligatorio, go al compilar lo añadira solo a go.mod)
// go run . => ejecuta sin compilar
// go build . => para compilar primero
// gofmt -w controllers/user_controllers.go => para formatear los archivos a los estándares de go
// el nombre del package tiene que ser el nombre de la carpeta que contiene los archivos

import (
	"fmt"
	"go_rest_api/api/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	var server *mux.Router = routes.NewRouter()

	fmt.Println("Server running on :3000")
	err := http.ListenAndServe(":3000", server)

	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
