package config

import (
	"log"

	"github.com/spf13/viper"
)

type ViperConfig struct {
	config *config
}

func ViperConfigInit() Config {
	viper.AddConfigPath("../")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error config file: %v", err)
	}

	return ViperConfig{config: &config{
		serverConfig: &ServerConfig{
			Host:         viper.GetString("app.host"),
			GrpcProtocol: viper.GetString("app.grpc_protocol"),
			GrpcPort:     viper.GetString("app.grpc_port"),
		},
	},
	}
}

func (v ViperConfig) GetServerConfig() *ServerConfig {
	return v.config.serverConfig
}

func (v ViperConfig) GetConfig() *config {
	return v.config
}
