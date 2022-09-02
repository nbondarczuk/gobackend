package db

import (
	"fmt"
	"log"
	"os"
)

// MakeConnectStrimg
func MakeConnectStringForPostgres() string {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DBNAME")

	cs := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)

	log.Println("Using Postgres DB connect string:", cs)

	return cs
}
