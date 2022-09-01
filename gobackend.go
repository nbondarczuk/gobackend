package main

import (
	"gobackend/db"
	"log"
)

func main() {
	log.Println("Started")

	// for each parameter value
	for _, kind := range [2]string{"inmem", "mysql"} {
		printBackendVersion(kind)
	}

	log.Println("Finished")
}

// printBackendVersion creates connection and prints the version info from db
func printBackendVersion(kind string) {
	be, err := db.NewBackend(kind)
	if err != nil {
		log.Fatal(err)
	}
	defer be.Close()

	version, err := be.Version()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Backend", be.Kind(), "version is", version)
}
