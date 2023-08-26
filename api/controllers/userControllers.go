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
	Phone	 string `json:"phonenumber"`
}

type Controller struct {
	db *sql.DB
}

func InitDB(db *sql.DB) *Controller {
	return &Controller{
		db: db,
	}
}

func  (DBInstance *Controller) GetUsers(response http.ResponseWriter, request *http.Request) {

	rows, errQuery := DBInstance.db.Query("SELECT user_id as ID, user_handle as Username, first_name as Name, last_name as Lastname, email as Password,phonenumber as Phone FROM users")

	if errQuery != nil {
		log.Fatalf("Error Query, %v", errQuery)
		return
	}
	defer rows.Close()

	var user []User

	for rows.Next() {
		var iUser User
		if err := rows.Scan(&iUser.ID, &iUser.Username, &iUser.Name, &iUser.Lastname, &iUser.Password, &iUser.Phone); err != nil {
			http.Error(response, "Server error 1", http.StatusInternalServerError)
			return
		}
		user = append(user, iUser)
	}

	if err := rows.Err(); err != nil {
		http.Error(response, "Server error 2", http.StatusInternalServerError)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(user)
}

func (DBInstance *Controller) CreateUser(response http.ResponseWriter, request *http.Request) {

	var user User

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&user)

	if err != nil {
		http.Error(response, "Error al decodificar el JSON", http.StatusBadRequest)
		return
	}

	defer request.Body.Close()

	var query string = "INSERT INTO users (user_handle, email, first_name, last_name, phonenumber) VALUES (?, ?, ?, ?, ?)"
	_, err = DBInstance.db.Exec(query, user.Username, user.Password, user.Name, user.Lastname, user.Phone)

	if err != nil {
		http.Error(response, "Error al insertar en la base de datos", http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
	response.Write([]byte("Creado con éxito!"))

}
