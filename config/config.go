package config

import (
	"fmt"
	"path/filepath"

	"github.com/fiffu/vprmail/utils"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	ServerNamespace string
	ServerSecret    string
	ClientHeartbeat int
	APIPort         int
	SMTPHostName    string
	SMTPPort        int
	SMTPWelcomeMsg  string
}

func SetupConfig() Config {
	here, err := utils.GetRuntimeDir()
	if err != nil {
		log.WithError(err).Error("failed to get current executable path")
	}
	abspath := filepath.Join(here, fmt.Sprintf("%s.%s", ConfigFileName, ConfigType))
	log.Infof("Searching for config=%s type=%s in path %s", ConfigFileName, ConfigType, abspath)
	viper.SetConfigFile(abspath)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Warnf("Config file not found: %s", ConfigFileName)
		} else {
			// Config file was found but another error was produced
			log.WithError(err).Errorf("Errored while trying to load config")
		}
	}

	viper.SetDefault(ServerNamespace, "default namespace")
	viper.SetDefault(ServerSecret, "default secret")
	viper.SetDefault(ClientHeartbeat, 3)

	return Config{
		ServerNamespace: viper.GetString(ServerNamespace),
		ServerSecret:    viper.GetString(ServerSecret),

		ClientHeartbeat: viper.GetInt(ClientHeartbeat),

		APIPort: viper.GetInt(APIPort),

		SMTPPort:       viper.GetInt(SMTPPort),
		SMTPHostName:   viper.GetString(SMTPHostName),
		SMTPWelcomeMsg: viper.GetString(SMTPWelcomeMsg),
	}
}
