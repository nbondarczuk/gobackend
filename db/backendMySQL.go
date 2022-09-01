package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Backend for MySQL DB
type BackendMySQL struct {
	BackendCommonInfo
	cs string
	db *sql.DB
}

// NewBackendMySQL creates and opens new MySQL DB connection
func NewBackendMySQL(cs string) (BackendMySQL, error) {
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

// Close backend connection
func (be BackendMySQL) Close() {
	be.db.Close()
}