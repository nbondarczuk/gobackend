package db

import (
	"database/sql"

	_ "github.com/lib/pq" // add this
)

// Backend for Postgres DB
type BackendPostgres struct {
	BackendCommonInfo
	cs string
	db *sql.DB
}

// NewBackendPostgres creates and opens new Postgres DB connection
func NewBackendPostgres(cs string) (BackendPostgres, error) {
	db, err := sql.Open("postgres", cs)
	if err != nil {
		return BackendPostgres{}, err
	}

	err = db.Ping()
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
func (be BackendPostgres) Close() {
	be.db.Close()
}
