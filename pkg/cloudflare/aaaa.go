package cloudflare

import (
	"bytes"
	"encoding/json"
	log "github.com/juusujanar/cloudflare-ddns/pkg/logging"
	"io/ioutil"
	"net/http"
)

func UpdateAAAARecords(ip string) {
	client := &http.Client{}
	for _, domain := range Config.Domains {
		// If we have AAAA Record Identifier
		if domain.IPv6 && len(domain.AAAARecordIdentifier) != 0 {
			url := "https://api.cloudflare.com/client/v4/zones/" + domain.ZoneIdentifier + "/dns_records/" + domain.AAAARecordIdentifier

			data := Request{
				Type:    "AAAA",
				Name:    domain.Name,
				Content: ip,
				Ttl:     domain.TTL,
				Proxied: domain.Proxied,
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

			response, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}

			body, err = ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}

			res := CFSingleResultResponse{}
			err = json.Unmarshal(body, &res)
			if err != nil {
				log.Fatal(err)
			}
			if res.Success {
				log.Info("AAAA Record update successful for domain " + domain.Name)
			} else {
				log.Info("AAAA Record update failed for domain " + domain.Name)
				log.Info(string(body))
			}
			response.Body.Close()
		} else {
			// If we do not have A record identifier, we need to create AAAA record and save it
			url := "https://api.cloudflare.com/client/v4/zones/" + domain.ZoneIdentifier + "/dns_records/"

			data := Request{
				Type:    "AAAA",
				Name:    domain.Name,
				Content: ip,
				Ttl:     domain.TTL,
				Proxied: domain.Proxied,
			}
			body, err := json.Marshal(data)
			if err != nil {
				log.Fatal(err)
			}

			req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			if err != nil {
				log.Fatal(err)
			}

			req.Header.Set("X-Auth-Email", Config.Email)
			req.Header.Set("X-Auth-Key", Config.ApiToken)
			req.Header.Set("Content-Type", "application/json")

			response, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}

			body, err = ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}

			res := CFSingleResultResponse{}
			err = json.Unmarshal(body, &res)
			if err != nil {
				log.Fatal(err)
			}
			if res.Success {
				domain.AAAARecordIdentifier = res.Result.Id
				log.Info("AAAA Record creation successful for domain " + domain.Name)
			} else {
				log.Info("AAAA Record creation failed for domain " + domain.Name)
				log.Info(string(body))
			}
			response.Body.Close()
		}
	}
}
