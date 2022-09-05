package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type ProcMode int

const (
	PRINT_VERSION    ProcMode = 1
	PING                      = 2
	CONFIG_FILE_NAME          = "config.yaml"
)

// Config is a ready to use, parsed structure, contrary to config from
// yaml or env variables.
type Config struct {
	IsHelpReq bool
	Mode      ProcMode
	Backends  []string
}

type Document struct {
	Backends []Backend `yaml:"backends"`
}

type Backend struct {
	Kind string            `yaml:"kind"`
	Env  map[string]string `yaml:"env"`
}

// NewConfig create a new configuration
func NewConfig() (*Config, error) {
	c := &Config{Mode: PING}
	err := c.load()
	return c, err
}

// Get config from files (if exist), env vars (if found) and command line (if used)
func (c *Config) load() error {
	c.initDefaultValues()
	err := c.loadConfigYamlFile()
	if err != nil {
		return err
	}
	c.loadCmdArgvs()

	return nil
}

// setInitConfig nitializes the config with initial profile
func (c *Config) initDefaultValues() {
	c.Mode = PRINT_VERSION
}

// loadConfigFile loads the config.yaml file overriding default config
func (c *Config) loadConfigYamlFile() error {
	log.Println("Reading config file:", CONFIG_FILE_NAME)
	yfl, err := ioutil.ReadFile(CONFIG_FILE_NAME)
	if err != nil {
		return err
	}

	var doc Document
	log.Println("Parsing config file:", CONFIG_FILE_NAME)
	err = yaml.Unmarshal(yfl, &doc)
	if err != nil {
		return err
	}

	for _, backend := range doc.Backends {
		c.setEnvVars(backend.Kind, backend.Env)
	}

	return nil
}

// setEnvVars overrides the default values from config with the env
func (c *Config) setEnvVars(kind string, env map[string]string) {
	c.Backends = append(c.Backends, kind)
	log.Printf("Using backend: %s -> %+v\n", kind, env)
	for key, value := range env {
		name := fmt.Sprintf("%s_%s", strings.ToUpper(kind), strings.ToUpper(key))
		os.Setenv(name, value)
		log.Printf("Add env variable: %s = %s\n", name, value)
	}
}

// loadCmdArgvs overrides config with command line switches
func (c *Config) loadCmdArgvs() {}
