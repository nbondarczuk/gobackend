package main

import (
	"gobackend/db"
	"log"
	"os"

	"gobackend/config"
)

func main() {
	log.Println("Started")

	cfg := config.Load()

	if cfg.IsHelpReq {
		printHelpAndExit()
	}

	// for each parameter value
	for _, kind := range cfg.Backends2Test {
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

// printHelpAndfExit prints help and exits
func printHelpAndExit() {
	os.Exit(0)
}
