package backend

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // add this
)

// Backend for Postgres DB
type BackendPostgres struct {
	BackendCommonInfo
	cs string
	db *sql.DB
}

type BackendCredentialsPostgres struct {
	user, password, dbname, host string
}

func NewBackendCredentialsPostgres() (BackendCredentials, error) {
	return BackendCredentialsPostgres{
			user:     os.Getenv("POSTGRES_USER"),
			password: os.Getenv("POSTGRES_PASSWORD"),
			dbname:   os.Getenv("POSTGRES_DBNAME"),
			host:     os.Getenv("POSTGRES_HOST"),
		},
		nil
}

func (bc BackendCredentialsPostgres) ConnectString() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", bc.host, bc.user, bc.password, bc.dbname)
}

// NewBackendPostgres creates and opens new Postgres DB connection
func NewBackendPostgres(bc BackendCredentials) (BackendPostgres, error) {
	cs := bc.ConnectString()
	db, err := sql.Open("postgres", cs)
	if err != nil {
		return BackendPostgres{}, err
	}

	return BackendPostgres{BackendCommonInfo{"mysql"}, cs, db}, nil
}

// Kind return what kind of backend db is used
func (be BackendPostgres) Kind() string {
	return be.kind
}

// Version obtains the backend server version: it is highly database dependent
func (be BackendPostgres) Version() (string, error) {
	var version string
	err := be.db.QueryRow("SELECT VERSION()").Scan(&version)
	if err != nil {
		return "", err
	}

	return version, nil
}

// Close backend connection
func (be BackendPostgres) Ping() error {
	err := be.db.Ping()
	if err != nil {
		return err
	}

	return nil
}

// Close backend connection
func (be BackendPostgres) Close() {
	be.db.Close()
}