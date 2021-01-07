package handlers

import (
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	mdl "slam/models"
	sec "slam/security"
	"strconv"
	"time"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	sec.IsAuthenticated(w, r)
	log.Println("Create User")
	if r.Method == "POST" {
		currentUser := GetUserInCookie(w, r)
		name := r.FormValue("name")
		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")
		mobile := r.FormValue("mobile")
		role := r.FormValue("role")
		qtdAtendimentos := r.FormValue("qtdAtendimentos")
		tipoEspecialidade := r.FormValue("tipoEspecialidade")
		outraEspecialidade := r.FormValue("outraEspecialidade")
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		sqlStatement := "INSERT INTO Users(name, username, password, email, mobile, qtd_atendimentos, tip_especialidade, outra_especialidade, role_id, author_id, criado_em) " +
			" VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING id"
		id := 0
		err = Db.QueryRow(sqlStatement, name, username, hash, email, mobile, qtdAtendimentos, tipoEspecialidade, outraEspecialidade, role, currentUser.Id, time.Now()).Scan(&id)
		if err != nil {
			log.Println(err.Error())
			errMsg := "Erro ao criar Usuário."
			if role == "" {
				errMsg = errMsg + " Faltou informar o Perfil do Usuário."
			}
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("Erro ao criar Usuário."))
		} else {
			log.Println("INSERT: Id: " + strconv.Itoa(id) +
				" | Name: " + name + " | Username: " + username +
				" | Password: " + password + " | Email: " + email +
				" | Mobile: " + mobile + " | Role: " + role)
			w.Write([]byte("Usuário criado com sucesso."))
		}
	}
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Update User")
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete User")
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("List Users")
	currentUser := GetUserInCookie(w, r)
	if sec.IsAuthenticated(w, r) && HasPermission(currentUser, "listUsers") {
		msg := r.FormValue("msg")
		errMsg := r.FormValue("errMsg")
		sql := "SELECT " +
			" a.id, a.name, a.username, a.password, " +
			" a.email, a.mobile, COALESCE(a.role_id, 0), COALESCE(b.name,'') as role_name, " +
			" a.author_id, " +
			" e.name as author_name, " +
			" to_char(a.criado_em,'DD/MM/YYYY HH24:MI:SS'), " +
			" coalesce(d.name,'') as cstatus, " +
			" a.status_id, " +
			" a.id_versao_origem " +
			" FROM users a " +
			" LEFT JOIN roles b ON a.role_id = b.id " +
			" LEFT JOIN status d ON a.status_id = d.id " +
			" LEFT JOIN users e ON a.author_id = e.id " +
			" ORDER BY a.name ASC "
		log.Println("SQL: " + sql)
		rows, _ := Db.Query(sql)
		defer rows.Close()
		var users []mdl.User
		var user mdl.User
		var i = 1
		for rows.Next() {
			rows.Scan(&user.Id,
				&user.Name,
				&user.Username,
				&user.Password,
				&user.Email,
				&user.Mobile,
				&user.Role,
				&user.RoleName,
				&user.AuthorId,
				&user.AuthorName,
				&user.C_CriadoEm,
				&user.CStatus,
				&user.StatusId,
				&user.IdVersaoOrigem)
			user.Order = i
			i++
			log.Println(user)
			users = append(users, user)
		}
		sql = "SELECT id, name FROM roles ORDER BY name asc"
		log.Println("SQL Roles: " + sql)
		rows, _ = Db.Query(sql)
		defer rows.Close()
		var roles []mdl.Role
		var role mdl.Role
		i = 1
		for rows.Next() {
			rows.Scan(&role.Id,
				&role.Name)
			role.Order = i
			i++
			roles = append(roles, role)
		}
		var page mdl.PageUsers
		if errMsg != "" {
			page.ErrMsg = errMsg
		}
		if msg != "" {
			page.Msg = msg
		}
		page.Users = users
		page.Roles = roles
		page.AppName = mdl.AppName
		page.Title = "Usuários"
		page.LoggedUser = BuildLoggedUser(GetUserInCookie(w, r))
		var tmpl = template.Must(template.ParseGlob("tiles/users/*"))
		tmpl.ParseGlob("tiles/*")
		tmpl.ExecuteTemplate(w, "Main-Users", page)
	} else {
		http.Redirect(w, r, "/logout", 301)
	}
}
