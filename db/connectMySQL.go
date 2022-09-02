package db

import (
	"fmt"
	"log"
	"os"
)

// MakeConnectStrimg
func MakeConnectStringForMySQL() string {
	user := os.Getenv("MYSQLDB_USER")
	password := os.Getenv("MYSQLDB_PASSWORD")
	dbname := os.Getenv("MYSQLDB_DBNAME")

	cs := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

	log.Println("Using MySQL DB connect string:", cs)

	return cs
}
