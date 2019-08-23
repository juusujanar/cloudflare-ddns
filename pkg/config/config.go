package config

import (
	"encoding/json"
	"flag"
	log "github.com/juusujanar/cloudflare-ddns/pkg/logging"
	"os"
)

type ConfiguredDomain struct {
	Name 	string	`json:"name"`
	TTL 	int		`json:"ttl"`
	IPv6 	bool	`json:"ipv6"`
	Proxied bool	`json:"proxied"`
}

type FileConfiguration struct {
	Domains 	[]ConfiguredDomain	`json:"domains"`
	Email 		string				`json:"email"`
	ApiToken 	string				`json:"token"`
}

func ReadConfig() FileConfiguration {
	c := flag.String("c", "config.json", "Specify the configuration file.")
	flag.Parse()
	file, err := os.Open(*c)
	if err != nil {
		log.Fatal("Can't open config file: ", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := FileConfiguration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal("Can't decode config JSON: ", err)
	}
	return config
}
