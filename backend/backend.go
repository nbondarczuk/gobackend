package backend

import (
	"fmt"

	"gobackend/backend/mysql"
	"gobackend/backend/postgres"
)

// Backend is an interface providing access to specific kind of db if needed
type Backend interface {
	Kind() string
	Version() (string, error)
	Ping() error
	Close()
}

// BackendCredentials is an interface handling db specific connect string
type BackendCredentials interface {
	ConnectString() string
}

// NewBackendCredentials fills the structure with items required for login
func NewBackendCredentials(kind string) (BackendCredentials, error) {
	switch kind {
	case "inmem":
		return nil, nil
	case "mysql":
		return mysql.NewBackendCredentialsMySQL()
	case "postgres":
		return postgres.NewBackendCredentialsPostgres()
	}

	return nil, fmt.Errorf("Invalid kind of backend: " + kind)
}

// NewBackend is a factory producing specific kind of backend db handlers based on dispatch
func NewBackend(kind string) (Backend, error) {
	bc, err := NewBackendCredentials(kind)
	if err != nil {
		return nil, fmt.Errorf("Can't get backend credentials: " + err.Error())
	}

	switch kind {
	case "inmem":
		return nil, nil
	case "mysql":
		return mysql.NewBackendMySQL(bc.(mysql.BackendCredentialsMySQL))
	case "postgres":
		return postgres.NewBackendPostgres(bc.(postgres.BackendCredentialsPostgres))
	}

	return nil, fmt.Errorf("Invalid kind of backend: " + kind)
}
