package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Password string `json:"password"`
}

// var users = []User{
// 	{ID: "1", Username: "Adotel", Name: "Adrian", Lastname: "Dotel", Password: "holaquetal"},
// 	{ID: "2", Username: "Adotel15", Name: "Adri", Lastname: "Dotel", Password: "holaquetal"},
// 	{ID: "3", Username: "Adotel15p", Name: "ADP", Lastname: "Dotel", Password: "holaquetal"},
// }

func GetUsers(response http.ResponseWriter, request *http.Request) {

	db, errDB := sql.Open("mysql", "root:@tcp(127.0.0.1:3500)/testdb")

	if errDB != nil {
		log.Fatalf("Error connection DB, %v", errDB)
	}
	// Esta linea se ejecuta al final de la función independientemente si ha ido bien o ha habido algun error
	defer db.Close()

	rows, errQuery := db.Query("SELECT user_id as ID, user_handle as Username, first_name as Name, last_name as Lastname, email as Password FROM users")

	if errQuery != nil {
		log.Fatalf("Error Query, %v", errQuery)
		return
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var iUser User
		if err := rows.Scan(&iUser.ID, &iUser.Username, &iUser.Name, &iUser.Lastname, &iUser.Password); err != nil {
			http.Error(response, "Server error", http.StatusInternalServerError)
			return
		}
		users = append(users, iUser)
	}

	if err := rows.Err(); err != nil {
		http.Error(response, "Server error", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(users)
}
