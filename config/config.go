package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type ProcMode int

const (
	PRINT_VERSION ProcMode = 1
	PING                   = 2
)

// Config is a ready to use, parsed structure, contrary to config from
// yaml or env variables.
type Config struct {
	IsHelpReq bool
	Mode      ProcMode
	Backends  []string
}

type ConfigYaml struct {
	backends Backends `yaml:"backends"`
}

type Backends struct {
	backend []Backend `yaml:"backend"`
}

type Backend struct {
	kind string `yaml:"kind"`
	env  Env    `yaml:"env"`
}

type Env struct {
	user     string `yaml:"user"`
	password string `yaml:"password"`
	dbname   string `yaml:"dbname"`
	host     string `yaml:"host"`
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
	c.Mode = PING
}

// loadConfigFile loads the config.yaml file overriding default config
func (c *Config) loadConfigYamlFile() error {
	yfl, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return err
	}

	var yc ConfigYaml
	err = yaml.Unmarshal(yfl, &yc)
	if err != nil {
		return err
	}

	// all side effects are stored in env variables
	for _, backend := range yc.backends.backend {
		c.Backends = append(c.Backends, backend.kind)
		c.setEnvVars(backend.kind, backend.env)
	}

	return nil
}

// loadEnvVars overrides the default values from config with the env
func (c *Config) setEnvVars(kind string, env Env) error {
	fmt.Printf("%+v", env)
	return nil
}

// loadCmdArgvs overrides config with command line switches
func (c *Config) loadCmdArgvs() {}
