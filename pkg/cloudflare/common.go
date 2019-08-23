package cloudflare

import "github.com/juusujanar/cloudflare-ddns/pkg/config"

type Request struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Ttl     int    `json:"ttl"`
	Proxied bool   `json:"proxied"`
}

var Config = config.ReadConfig()