package handlers

import (
	"html/template"
	"log"
	"net/http"
	mdl "slam/models"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Create User")
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update User")
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete User")
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Users")
	sql := " SELECT id, name FROM users "
	log.Println("SQL: " + sql)
	rows, _ := Db.Query(sql)
	defer rows.Close()
	var users []mdl.User
	var user mdl.User
	for rows.Next() {
		rows.Scan(&user.Id,
			&user.Name)
		users = append(users, user)
	}
	var page mdl.PageUser
	page.Users = users
	page.AppName = mdl.AppName
	page.Title = "Usuários"
	var tmpl = template.Must(template.ParseGlob("tiles/users/*"))
	tmpl.ParseGlob("tiles/*")
	tmpl.ExecuteTemplate(w, "Main-Users", page)
}
