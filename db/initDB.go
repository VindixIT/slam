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

	// Sequence ACTIONS
	stmt = "CREATE SEQUENCE IF NOT EXISTS actions_id_seq"
	log.Println(stmt)
	db.Exec(stmt)

	// Sequence ACTIONS_STATUS
	stmt = "CREATE SEQUENCE IF NOT EXISTS actions_status_id_seq"
	log.Println(stmt)
	db.Exec(stmt)

	// Sequence ACTIVITIES
	stmt = "CREATE SEQUENCE IF NOT EXISTS activities_id_seq"
	log.Println(stmt)
	db.Exec(stmt)

	// Sequence ACTIVITIES_ROLES
	stmt = "CREATE SEQUENCE IF NOT EXISTS activities_roles_id_seq"
	log.Println(stmt)
	db.Exec(stmt)

	// Sequence FEATURES
	stmt = "CREATE SEQUENCE IF NOT EXISTS features_activities_id_seq"
	log.Println(stmt)
	db.Exec(stmt)

	// Sequence STATUS
	stmt = "CREATE SEQUENCE IF NOT EXISTS status_id_seq"
	log.Println(stmt)
	db.Exec(stmt)

	// Sequence WORKFLOWS
	stmt = "CREATE SEQUENCE IF NOT EXISTS workflows_id_seq"
	log.Println(stmt)
	db.Exec(stmt)
}

func createTable() {

	// Table ACTIONS
	stmt := " CREATE TABLE IF NOT EXISTS actions (" +
		" id integer DEFAULT nextval('actions_id_seq'::regclass) NOT NULL, " +
		" name character varying(255) NOT NULL, " +
		" origin_status_id integer, " +
		" destination_status_id integer, " +
		" other_than boolean, " +
		" description character varying(4000)," +
		" author_id integer," +
		" created_at timestamp without time zone," +
		" id_versao_origem integer," +
		" status_id integer)"
	log.Println(stmt)
	db.Exec(stmt)

	// Table ACTIONS_STATUS
	stmt = " CREATE TABLE IF NOT EXISTS actions_status (" +
		" id integer DEFAULT nextval('actions_status_id_seq'::regclass)," +
		" action_id integer," +
		" origin_status_id integer," +
		" destination_status_id integer)"
	log.Println(stmt)
	db.Exec(stmt)

	// Table ACTIVITIES
	stmt = " CREATE TABLE IF NOT EXISTS activities (" +
		" id integer NOT NULL DEFAULT nextval('activities_id_seq'::regclass)," +
		" workflow_id integer," +
		" action_id integer," +
		" expiration_action_id integer," +
		" expiration_time_days integer," +
		" start_at timestamp without time zone," +
		" end_at timestamp without time zone)"
	log.Println(stmt)
	db.Exec(stmt)

	// Table ACTIVITIES_ROLES
	stmt = " CREATE TABLE IF NOT EXISTS activities_roles (" +
		" id integer DEFAULT nextval('activities_roles_id_seq'::regclass)," +
		" activity_id integer," +
		" role_id integer)"
	log.Println(stmt)
	db.Exec(stmt)

	// Table FEATURES_ROLES
	stmt = " CREATE TABLE IF NOT EXISTS features_roles (" +
		" id integer DEFAULT nextval('features_roles_id_seq'::regclass)," +
		" feature_id integer," +
		" role_id integer)"
	log.Println(stmt)
	db.Exec(stmt)

	// Table FEATURES_ACTIVITIES
	stmt = " CREATE TABLE IF NOT EXISTS features_activities (" +
		" id integer DEFAULT nextval('features_activities_id_seq'::regclass)," +
		" feature_id integer," +
		" activity_id integer)"
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
	stmt = "CREATE IF NOT EXISTS TABLE features_roles (" +
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

	// Table STATUS
	stmt = " CREATE TABLE IF NOT EXISTS status  (" +
		" id integer DEFAULT nextval('status_id_seq'::regclass) NOT NULL," +
		" name character varying(255) NOT NULL," +
		" description character varying(4000)," +
		" author_id integer," +
		" created_at timestamp without time zone," +
		" id_versao_origem integer," +
		" status_id integer," +
		" stereotype character varying(255))"
	log.Println(stmt)
	db.Exec(stmt)

	// Table USERS
	stmt = "CREATE TABLE IF NOT EXISTS users (" +
		" id integer DEFAULT nextval('users_id_seq'::regclass) NOT NULL," +
		" name character varying(255)," +
		" username character varying(255) NOT NULL," +
		" password character varying(255) NOT NULL," +
		" email character varying(255)," +
		" mobile character varying(255)," +
		" role_id integer," +
		" qtd_atendimentos int," +
		" tip_especialidade character varying(50)," +
		" outra_especialidade character varying(50)," +
		" author_id integer," +
		" criado_em timestamp without time zone," +
		" id_versao_origem integer," +
		" status_id integer)"
	log.Println(stmt)
	db.Exec(stmt)

	// Table WORKFLOWS
	stmt = " CREATE TABLE IF NOT EXISTS workflows  (" +
		" id integer DEFAULT nextval('workflows_id_seq'::regclass) NOT NULL," +
		" name character varying(255) NOT NULL," +
		" description character varying(4000)," +
		" entity_type character varying(50)," +
		" start_at timestamp without time zone," +
		" end_at timestamp without time zone," +
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
	db.Exec("ALTER TABLE ONLY actions ADD CONSTRAINT actions_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY actions_status ADD CONSTRAINT actions_status_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY activities ADD CONSTRAINT activities_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY activities_roles ADD CONSTRAINT activities_roles_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY features ADD CONSTRAINT features_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY features_activities ADD CONSTRAINT features_activities_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY features_roles ADD CONSTRAINT features_roles_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY roles ADD CONSTRAINT roles_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY status ADD CONSTRAINT status_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY users ADD CONSTRAINT users_pkey PRIMARY KEY (id)")
	db.Exec("ALTER TABLE ONLY workflows ADD CONSTRAINT workflows_pkey PRIMARY KEY (id)")
}

func createFKey() {
	// ACTIONS
	db.Exec("ALTER TABLE ONLY actions " +
		" ADD CONSTRAINT destination_status_fkey FOREIGN KEY (destination_status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY actions " +
		" ADD CONSTRAINT origin_status_fkey FOREIGN KEY (origin_status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY actions " +
		" ADD CONSTRAINT workflows_fkey FOREIGN KEY (workflow_id)" +
		" REFERENCES workflows (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	//  ACTIONS_STATUS
	db.Exec("ALTER TABLE ONLY actions_status " +
		" ADD CONSTRAINT actions_fkey FOREIGN KEY (action_id)" +
		" REFERENCES actions (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY actions_status " +
		" ADD CONSTRAINT origin_status_fkey FOREIGN KEY (origin_status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY actions_status " +
		" ADD CONSTRAINT destination_status_fkey FOREIGN KEY (destination_status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	// ACTIVITIES
	db.Exec("ALTER TABLE activities ADD CONSTRAINT action_fkey FOREIGN KEY (action_id)" +
		" REFERENCES actions (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE activities ADD CONSTRAINT expiration_action_fkey FOREIGN KEY (expiration_action_id)" +
		" REFERENCES actions (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE activities ADD CONSTRAINT workflow_fkey FOREIGN KEY (workflow_id)" +
		" REFERENCES workflows (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	// ACTIVITIES_ROLES
	db.Exec("ALTER TABLE ONLY activities_roles " +
		" ADD CONSTRAINT activities_fkey FOREIGN KEY (activity_id)" +
		" REFERENCES activities (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY activities_roles " +
		" ADD CONSTRAINT roles_fkey FOREIGN KEY (role_id)" +
		" REFERENCES roles (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

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

	// USERS
	db.Exec("ALTER TABLE ONLY users " +
		" ADD CONSTRAINT roles_fkey FOREIGN KEY (role_id)" +
		" REFERENCES roles (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT")

	db.Exec("ALTER TABLE ONLY users" +
		" ADD CONSTRAINT authors_fkey FOREIGN KEY (author_id)" +
		" REFERENCES users (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT ON DELETE RESTRICT" +
		" NOT VALID")

	db.Exec("ALTER TABLE ONLY users" +
		" ADD CONSTRAINT status_fkey FOREIGN KEY (status_id)" +
		" REFERENCES status (id) MATCH SIMPLE" +
		" ON UPDATE RESTRICT" +
		" ON DELETE RESTRICT" +
		" NOT VALID")

}
