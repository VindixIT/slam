package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	db "slam/db"
	hd "slam/handlers"
	routes "slam/routes"
)

func main() {
	hd.Db = dbConectar()
	log.Println("CONECTADO COM SUCESSO!")
	db.Initialize()
	defer hd.Db.Close()
	r := mux.NewRouter()
	// ----------------- HOME E SECURITY
	r.HandleFunc("/", hd.IndexHandler).Methods("GET")
	r.HandleFunc("/login", hd.LoginHandler)
	r.HandleFunc("/logout", hd.LogoutHandler).Methods("GET")
	// ----------------- USERS
	r.HandleFunc(routes.UsersRoute, hd.ListUsersHandler).Methods("GET")
	r.HandleFunc("/createUser", hd.CreateUserHandler).Methods("POST")
	r.HandleFunc("/updateUser", hd.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/deleteUser", hd.DeleteUserHandler).Methods("DELETE")

	http.Handle("/", r)
	http.Handle("/statics/",
		http.StripPrefix("/statics/", http.FileServer(http.Dir("./statics"))),
	)
	http.ListenAndServe(":5000", nil)
}

func dbConectar() *sql.DB {
	conexao, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5435/slam?sslmode=disable")
	if err != nil {
		log.Println(err.Error())
	}
	return conexao
}
