package main

import (
	"log"
	"os"

	"gobackend/backend"
	"gobackend/config"
)

var cfg *config.Config

// main load config and does something with each backend requested
func main() {
	log.Println("Started")

	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	if cfg.IsHelpReq {
		printHelpAndExit()
	}

	// for each parameter value provided in the command line
	//for _, kind := range cfg.Backends {
	//	procBackend(kind, cfg.Mode)
	//}

	log.Println("Finished")
}

// procBackend creates connection and uses backend to do something
func procBackend(kind string, mode config.ProcMode) {
	be, err := backend.NewBackend(kind)
	if err != nil {
		log.Fatal(err)
	}
	defer be.Close()

	switch mode {
	case config.PRINT_VERSION:
		version, err := be.Version()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Backend", be.Kind(), "version is", version)

	case config.PING:
		err := be.Ping()
		if err != nil {
			log.Fatal(err)
		}
	}
}

// printHelpAndfExit prints help and exits
func printHelpAndExit() {
	os.Exit(0)
}
