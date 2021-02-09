package service

import (
	"fmt"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var bindings = map[string]string{
	"version":         "VERSION",
	"service.name":    "SERVICE_NAME",
	"log.filepath":    "LOG_FILE_PATH",
	"log.level":       "LOG_LEVEL",
	"log.max-size":    "LOG_MAX_SIZE",
	"log.max-backups": "LOG_MAX_BACKUP",
	"log.max-age":     "LOG_MAX_AGE",
	"log.compress":    "LOG_COMPRESS",
}

// BindEnvVars binds any environment variables that have been defined with the
// specified viper configuration. The default settings have been predefined so you
// only need to set additional bindings specific to your application.
// The prefix will be added to all defined environment variables. For example,
// if your chosen prefix is myapp, then any environment variable defined for the
// binding will be prefixed as MYAPP_<ENV_VAR>.
// E.g.
//
// SetEnvVarBinding("service.address", "SERVICE_ADDRESS")
//
// will bind the service.address configuration to the MYAPP_SERVICE_ADDRESS
// environment variable.
func BindEnvVars(prefix string) {
	log := zap.L()
	viper.SetEnvPrefix(prefix)
	for k, v := range bindings {
		if err := viper.BindEnv(k, fmt.Sprintf("%s_%s", prefix, v)); err != nil {
			log.Error("BindEnv", zap.String("config", k))
		}
	}
}

// SetEnvVarBinding adds an environment variable binding for viper configuration
// This function should be called for each binding you wish to add before calling
// the BindEnvVars function
func SetEnvVarBinding(config, env string) {
	bindings[config] = env
}
