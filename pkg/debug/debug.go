package debug

import (
	"robot/config"
)

const (
	AppEnvDev        = "dev"
	AppEnvProduction = "prod"
	AppEnvStaging    = "staging"
)

func DebugMessage(debugMsg, fallbackMsg string) string {
	if config.Get().AppEnv == AppEnvDev && config.Get().AppDebug == true {
		return debugMsg
	}
	return fallbackMsg
}
