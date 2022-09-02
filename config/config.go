package config

type Config struct {
	IsHelpReq     bool
	Backends2Test []string
}

var cnf Config

// Get config from files (if exist), env vars (if found) and command line (if used)
func Load() Config {
	setInitConfig()
	loadConfigFile()
	loadEnvVars()
	loadCmdArgvs()

	return cnf
}

// setInitConfig nitializes the config with initial profile
func setInitConfig() {}

// loadConfigFile loads the config.yaml file overriding default config
func loadConfigFile() {}

// loadEnvVars checks for env variables overriding file config

func loadEnvVars() {}

// loadCmdArgvs overrides config with command line switches
func loadCmdArgvs() {}
