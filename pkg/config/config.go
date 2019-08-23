package config

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

type Configuration struct {
	ZoneIdentifier string
	DNSRecord string
	Domains []string
	Domain string
	IPv6 bool
	TTL int
	Email string
	ApiToken string
	Proxied bool
}

func ReadConfig() Configuration {
	c := flag.String("c", "config.json", "Specify the configuration file.")
	flag.Parse()
	file, err := os.Open(*c)
	if err != nil {
		log.Fatal("Can't open config file: ", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	Config := Configuration{}
	err = decoder.Decode(&Config)
	if err != nil {
		log.Fatal("Can't decode config JSON: ", err)
	}
	log.Println(Config)
	return Config
}
