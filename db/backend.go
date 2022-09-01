package db

import (
	"fmt"
)

type BackendCommonInfo struct {
	kind string
}

// Backend is a generic type providing access to specific kind of db if needed
type Backend interface {
	Kind() string
	Version() (string, error)
	Close()
}

// NewBackend is a factory producing specific kind of backend db handlers
func NewBackend(kind string) (Backend, error) {
	if kind == "inmem" {
		return NewBackendInMem()
	} else if kind == "mysql" {
		return NewBackendMySQL(MakeConnectStringForMySQL())
	} else {
		return nil, fmt.Errorf("Invalid kind of backend: " + kind)
	}
}
