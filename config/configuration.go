package config

import (
	"errors"
	"fmt"
	"github.com/Netflix/go-env"
	"os"
	"robot/pkg/logger"

	"github.com/joho/godotenv"
)

var config *Configuration

//Configuration This struct holds config keys and values
// which are used in the mod_tracking. Only this struct must be used
// to hold any configuration values, no direct access to
// env, ini or any other config source should be made
type Configuration struct {
	// the rest of configuration for a mod_tracking MUST be defined in this struct
	AppEnv             string `env:"APP_ENV"`
	AppName            string `env:"APP_NAME"`
	AppDebug           bool   `env:"APP_DEBUG"`
	HttpListenAddr     string `env:"HTTP_LISTEN_ADDR"`
	HttpBaseRequestUrl string `env:"HTTP_BASE_REQUEST_URI"`

	AppBaseUrl string `env:"APP_BASE_URL"`

	PgConnUrl     string `env:"PG_CONNECT_URL"`
	PGUser        string `env:"PG_USER"`
	PGPassword    string `env:"PG_PASS"`
	PGDb          string `env:"PG_DATABASE_NAME"`
	PGHost        string `env:"PG_HOST"`
	PGPort        string `env:"PG_PORT"`
	PgConnTimeout int    `env:"PG_CONN_TIMEOUT"`
}

func InitConfig(envName string) error {
	c := &Configuration{}
	var err error
	if envName != "" {
		logger.Get().Info("trying to publish env from file", "file", envName)
		err = godotenv.Load(envName)
		if err != nil {
			return errors.New("failed to load configuration file " + envName)
		}
	}

	_, err = env.UnmarshalFromEnviron(c)

	if err != nil {
		return errors.New("failed to map env variables to Configuration object")
	}

	config = c

	return nil
}

func Get() *Configuration {
	if config == nil {
		fmt.Println("Config is not initialized")
		os.Exit(1)
	}
	return config
}
