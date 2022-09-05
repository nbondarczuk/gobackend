package config

type ProcMode int

const (
	PRINT_VERSION ProcMode = 1
	PING                   = 2
)

type Config struct {
	IsHelpReq bool
	Mode      ProcMode
	Backends  []string
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
func setInitConfig() {
	cnf.Mode = PING
	cnf.Backends = []string{"mysql", "postgres"}
}

// loadConfigFile loads the config.yaml file overriding default config
func loadConfigFile() {}

// loadEnvVars checks for env variables overriding file config
func loadEnvVars() {}

// loadCmdArgvs overrides config with command line switches
func loadCmdArgvs() {}
