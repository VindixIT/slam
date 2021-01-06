package db

import (
	"database/sql"
	"log"
	hd "slam/handlers"
	"strconv"
	"strings"
)

var db *sql.DB

func Initialize() {
	db = hd.Db
	createSeq()
	createTable()
	createFKey()
	createPKey()
	createFeatures()
	createRoles()
	createRoleFeatures()
	createUsers()
}

func createSeq() {
	// Sequence USERS
	stmt := "CREATE SEQUENCE IF NOT EXISTS users_id_seq"
	log.Println(stmt)
	db.Exec(stmt)

	// Sequence FEATURES
	stmt = "CREATE SEQUENCE IF NOT EXISTS features_id_seq"
	log.Println(stmt)
	db.Exec(stmt)

	// Sequence FEATURES_ROLES
	stmt = "CREATE SEQUENCE IF NOT EXISTS features_roles_id_seq"
	log.Println(stmt)
	db.Exec(stmt)

	// Sequence ROLES
	stmt = "CREATE SEQUENCE IF NOT EXISTS roles_id_seq"
	log.Println(stmt)
	db.Exec(stmt)
}

func createTable() {
	// Table USERS
	stmt := "CREATE TABLE IF NOT EXISTS users (" +
		" id integer DEFAULT nextval('users_id_seq'::regclass) NOT NULL," +
		" name character varying(255)," +
		" username character varying(255) NOT NULL," +
		" password character varying(255) NOT NULL," +
		" email character varying(255)," +
		" mobile character varying(255)," +
		" role_id integer," +
		" author_id integer," +
		" criado_em timestamp without time zone," +
		" id_versao_origem integer," +
		" status_id integer)"
	log.Println(stmt)
	db.Exec(stmt)

	// Table FEATURES
	stmt = "CREATE TABLE IF NOT EXISTS features  (" +
		" id integer DEFAULT nextval('features_id_seq'::regclass) NOT NULL," +
		" name character varying(255) NOT NULL," +
		" code character varying(255) NOT NULL," +
		" description character varying(4000)," +
		" author_id integer," +
		" created_at timestamp without time zone," +
		" id_versao_origem integer," +
		" status_id integer)"
	log.Println(stmt)
	db.Exec(stmt)

	// Table FEATURES_ROLES
	stmt = "CREATE TABLE features_roles (" +
		" id integer DEFAULT nextval('features_roles_id_seq'::regclass)," +
		" feature_id integer," +
		" role_id integer)"
	log.Println(stmt)
	db.Exec(stmt)

	// Table ROLES
	stmt = "CREATE TABLE IF NOT EXISTS roles  (" +
		" id integer DEFAULT nextval('roles_id_seq'::regclass) NOT NULL," +
		" name character varying(255) NOT NULL," +
		" description character varying(4000)," +
		" author_id integer," +
		" created_at timestamp without time zone," +
		" id_versao_origem integer," +
		" status_id integer)"
	log.Println(stmt)
	db.Exec(stmt)
}

func createFeatures() {
	db.Exec("INSERT INTO features (name, code) SELECT 'Listar Usuários', 'listUsers' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'listUsers')")
	db.Exec("INSERT INTO features (name, code) SELECT 'Criar Usuário', 'createUser' WHERE NOT EXISTS (SELECT 1 FROM features WHERE code = 'createUser')")
}

func createRoles() {
	stmt := "INSERT INTO roles (name, description, created_at) " +
		" SELECT 'Admin', 'Admin' , now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM roles WHERE name = 'Admin')"
	db.Exec(stmt)
}

func createRoleFeatures() {
	stmt1 := "INSERT INTO features_roles (role_id, feature_id) "
	stmt2 := ""
	for j := 1; j <= 2; j++ {
		roleId := "1"
		featureId := strconv.Itoa(j)
		stmt2 = stmt2 + " SELECT " + roleId + ", " + featureId + " WHERE NOT EXISTS (SELECT 1 FROM features_roles WHERE feature_id = " + featureId + " AND role_id = " + roleId + ") UNION "
	}
	pos := strings.LastIndex(stmt2, "UNION")
	stmt2 = stmt2[:pos]
	log.Println(stmt1 + stmt2)
	db.Exec(stmt1 + stmt2)
}

func createUsers() {
	stmt := "INSERT INTO users (username, password, email, mobile, name, role_id, author_id, criado_em) " +
		" SELECT 'admin', '$2a$10$ZkkRxriHQGsGwbt1PV/k/eagAm/PVJO0DE5ApZNRC8HwX2.bbyo6G', " +
		" 'masaru@vindixit.com', '61 984385415', 'Administrador do Sistema', 1, 1, now()::timestamp " +
		" WHERE NOT EXISTS (SELECT id FROM users WHERE username = 'admin')"
	log.Println(stmt)
	db.Exec(stmt)
}

func createPKey() {
	db.Exec("ALTER TABLE ONLY users ADD CONSTRAINT users_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY features ADD CONSTRAINT features_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY roles ADD CONSTRAINT roles_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY features_roles ADD CONSTRAINT features_roles_pkey PRIMARY KEY (id)")
}

func createFKey() {
	// FEATURES_ROLES
	db.Exec("ALTER TABLE ONLY features_roles " +
		" ADD CONSTRAINT features_fkey FOREIGN KEY (feature_id)" +
		" REFERENCES features (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY features_roles " +
		" ADD CONSTRAINT roles_fkey FOREIGN KEY (role_id)" +
		" REFERENCES roles (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")
}
