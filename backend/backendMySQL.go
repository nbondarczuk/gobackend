package backend

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Backend for MySQL DB
type BackendMySQL struct {
	BackendCommonInfo
	cs string
	db *sql.DB
}

type BackendCredentialsMySQL struct {
	user, password, dbname string
}

func NewBackendCredentialsMySQL() (BackendCredentials, error) {
	return BackendCredentialsMySQL{
			user:     os.Getenv("MYSQL_USER"),
			password: os.Getenv("MYSQL_PASSWORD"),
			dbname:   os.Getenv("MYSQL_DBNAME"),
		},
		nil
}

func (bc BackendCredentialsMySQL) ConnectString() string {
	return fmt.Sprintf("%s:%s@/%s", bc.user, bc.password, bc.dbname)
}

// NewBackendMySQL creates and opens new MySQL DB connection
func NewBackendMySQL(bc BackendCredentials) (BackendMySQL, error) {
	cs := bc.ConnectString()
	db, err := sql.Open("mysql", cs)
	if err != nil {
		return BackendMySQL{}, err
	}

	return BackendMySQL{BackendCommonInfo{"mysql"}, cs, db}, nil
}

// Kind return what kind of backend db is used
func (be BackendMySQL) Kind() string {
	return be.kind
}

// Version obtains the backend server version: it is highly database dependent
func (be BackendMySQL) Version() (string, error) {
	var version string
	err := be.db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		return "", err
	}

	return version, nil
}

// Ping backend db
func (be BackendMySQL) Ping() error {
	err := be.db.Ping()
	if err != nil {
		return err
	}

	return nil
}

// Close backend connection
func (be BackendMySQL) Close() {
	be.db.Close()
}
