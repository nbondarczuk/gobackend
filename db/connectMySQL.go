package db

import (
	"log"
	"os"
)

// MakeConnectStrimg
func MakeConnectStringForMySQL() string {
	user := os.Getenv("MYSQLDB_USER")
	password := os.Getenv("MYSQLDB_PASSWORD")
	dbname := os.Getenv("MYSQLDB_DBNAME")
	cs := user + ":" + password + "@/" + dbname
	log.Println("Using connect string:", cs)

	return cs
}
