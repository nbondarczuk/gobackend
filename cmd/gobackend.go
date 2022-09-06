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

	LoadConfig()

	if cfg.IsHelpReq {
		printHelpAndExit()
	}

	// for each parameter value provided in the command line
	//for _, kind := range cfg.Backends {
	//	procBackend(kind, cfg.Mode)
	//}

	log.Println("Finished")
}

// LoadConfig gets the contents of file and uses it to make a config
func LoadConfig() {
	input, err := config.ConfigYamlFromFile()
	if err != nil {
		panic(err)
	}
	cfg, err = config.NewConfig(input)
	if err != nil {
		panic(err)
	}
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
