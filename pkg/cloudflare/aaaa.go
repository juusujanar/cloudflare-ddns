package cloudflare

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func UpdateAAAARecord(ip string)  {
	client := &http.Client{}
	url := "https://api.cloudflare.com/client/v4/zones/" + Config.ZoneIdentifier + "/dns_records/" + Config.DNSRecord

	data := Request{
		Type:    "AAAA",
		Name:    Config.Domain,
		Content: ip,
		Ttl:     Config.TTL,
		Proxied: Config.Proxied,
	}
	body, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("X-Auth-Email", Config.Email)
	req.Header.Set("X-Auth-Key", Config.ApiToken)
	req.Header.Set("Content-Type", "application/json")

	_, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}