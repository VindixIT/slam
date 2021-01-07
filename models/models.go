package models

import ()

var AppName = "Slam"

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Feature struct {
	Order          int
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Code           string `json:"code"`
	Description    string `json:"description"`
	AuthorId       int64  `json:"authorId"`
	AuthorName     string `json:"authorName"`
	CreatedAt      string `json:"createdAt"`
	C_CreatedAt    string `json:"c_createdAt"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
	Selected       bool
}

type LoggedUser struct {
	User          User
	HasPermission func(string) bool
}

type PageUsers struct {
	WarnMsg    string
	ErrMsg     string
	Msg        string
	AppName    string
	Title      string
	LoggedUser LoggedUser
	Users      []User
	Roles      []Role
}

type Role struct {
	Order          int
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	AuthorId       int64  `json:"autorId"`
	AuthorName     string `json:"autorNome"`
	CreatedAt      string `json:"createdAt"`
	C_CreatedAt    string `json:"c_createdAt"`
	IdVersaoOrigem int64  `json:"idVersaoOrigem"`
	StatusId       int64  `json:"statusId"`
	CStatus        string `json:"cStatus"`
	Selected       bool
	Features       []Feature
}

type User struct {
	Order          int       `json:"order"`
	Id             int64     `json:"id"`
	Name           string    `json:"name"`
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Email          string    `json:"email"`
	Mobile         string    `json:"mobile"`
	Role           int64     `json:"role"`
	RoleName       string    `json:"roleName"`
	AuthorId       int64     `json:"authorId"`
	AuthorName     string    `json:"authorName"`
	CriadoEm       string    `json:"criadoEm"`
	C_CriadoEm     string    `json:"c_criadoEm"`
	IdVersaoOrigem int64     `json:"idVersaoOrigem"`
	StatusId       int64     `json:"statusId"`
	CStatus        string    `json:"cStatus"`
	Features       []Feature `json:"features"`
	Selected       bool      `json:"selected"`
}
