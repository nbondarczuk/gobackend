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

// ConfigYamlFromFile gets the contents of the config file.
func ConfigYamlFromFile() ([]byte, error) {
	log.Println("Reading config file:", CONFIG_FILE_NAME)
	input, err := ioutil.ReadFile(CONFIG_FILE_NAME)
	if err != nil {
		return nil, err
	}

	return input, nil
}

// String makes Config a Stringer (implicitely).
func (c *Config) String() string {
	return fmt.Sprintf("%T%+v", *c, *c)
}

// NewConfig create a new configuration.
func NewConfig(input []byte) (*Config, error) {
	c := &Config{Mode: PING}
	err := c.load(input)
	log.Printf("Loaded: %s\n", c.String())

	return c, err
}

// load gets config from file, env vars (if found) and command line (if used).
func (c *Config) load(input []byte) error {
	c.initDefaultValues()
	err := c.loadConfigYaml(input)
	if err != nil {
		return err
	}

	return nil
}

// setInitConfig nitializes the config with initial profile.
func (c *Config) initDefaultValues() {
	c.Mode = PRINT_VERSION
}

// loadConfigFile loads the config.yaml file overriding default config.
func (c *Config) loadConfigYaml(input []byte) error {
	var doc Document
	log.Println("Parsing config file:", CONFIG_FILE_NAME)
	err := yaml.Unmarshal(input, &doc)
	if err != nil {
		return err
	}

	for _, backend := range doc.Backends {
		c.setEnvVars(backend.Kind, backend.Env)
	}

	return nil
}

// setEnvVars overrides the default values from config with the env.
func (c *Config) setEnvVars(kind string, env map[string]string) {
	c.Backends = append(c.Backends, kind)
	log.Printf("Using backend: %s -> %+v\n", kind, env)
	for key, envval := range env {
		// Add new env var.
		envvar := fmt.Sprintf("%s_%s", strings.ToUpper(kind), strings.ToUpper(key))
		// First use existing env var values, then config, then command line vars.
		if os.Getenv(envvar) == "" {
			flg := fmt.Sprintf("%s_%s", strings.ToLower(kind), strings.ToLower(key))
			flgval := checkFlagUsage(flg)
			if flgval == "" {
				os.Setenv(envvar, envval)
				log.Printf("Use cnf variable: %s = %s\n", envvar, envval)
			} else {
				os.Setenv(envvar, flgval)
				log.Printf("Use flg variable: %s = %s\n", envvar, envval)
			}
		} else {
			log.Printf("Use env variable: %s = %s\n", envvar, os.Getenv(envvar))
		}
	}
}

func checkFlagUsage(flg string) string {
	return ""
}
