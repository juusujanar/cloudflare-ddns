package config

import (
	"github.com/spf13/viper"

	log "github.com/juusujanar/cloudflare-ddns/pkg/logging"
)

type ConfiguredDomain struct {
	Name 	string	`mapstructure:"name"`
	TTL 	int		`mapstructure:"ttl"`
	IPv6 	bool	`mapstructure:"ipv6"`
	Proxied bool	`mapstructure:"proxied"`
}

type FileConfiguration struct {
	Domains 	[]ConfiguredDomain	`mapstructure:"domains"`
	Email 		string				`mapstructure:"email"`
	ApiToken 	string				`mapstructure:"token"`
}

func ReadConfig() FileConfiguration {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("/etc/cf-ddns/")   // path to look for the config file in
	viper.AddConfigPath("$HOME/.cf-ddns")  // call multiple times to add many search paths
	viper.AddConfigPath(".")               // optionally look for config in the working directory
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found.")
		} else {
			log.Fatal("Config file parsing failed.")
		}
	}
	config := FileConfiguration{}
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Unable to decode configuration file, %v", err)
	}
	if len(config.Email) <= 3 {
		log.Fatal("Email is too short or not configured.")
	}
	if len(config.ApiToken) <= 3 {
		log.Fatal("API token is too short or not configured.")
	}
	return config
}
