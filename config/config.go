package config

import (
	"github.com/spf13/viper"
)

const (
	CONFIG_FILE_NAME = "config.ini"
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
	viper.SetDefault("ServerNamespace", "default namespace")
	viper.SetDefault("ServerSecret", "default secret")
	viper.SetDefault("ClientHeartbeatSecs", 3)

	viper.AddConfigPath("./" + CONFIG_FILE_NAME)
	viper.ReadInConfig()

	return Config{
		ServerNamespace: viper.GetString("ServerNamespace"),
		ServerSecret:    viper.GetString("ServerSecret"),

		ClientHeartbeat: viper.GetInt("ClientHeartbeatSecs"),

		APIPort: viper.GetInt(""),

		SMTPPort:       viper.GetInt(""),
		SMTPHostName:   viper.GetString("SMTPHostName"),
		SMTPWelcomeMsg: viper.GetString(""),
	}
}
