package cloudflare

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/juusujanar/cloudflare-ddns/pkg/logging"
)

func UpdateARecords(ip string) {
	client := &http.Client{}
	for _, domain := range Config.Domains {
		// If we have A Record Identifier
		if len(domain.ARecordIdentifier) != 0 {
			url := "https://api.cloudflare.com/client/v4/zones/" + domain.ZoneIdentifier + "/dns_records/" + domain.ARecordIdentifier

			data := Request{
				Type:    "A",
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

			req.Close = true

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
				log.Info("A Record update successful for domain " + domain.Name)
			} else {
				log.Info("A Record update failed for domain " + domain.Name)
				log.Info(string(body))
			}
			response.Body.Close()
		} else {
			// If we do not have A record identifier, we need to create A record and save it
			url := "https://api.cloudflare.com/client/v4/zones/" + domain.ZoneIdentifier + "/dns_records/"

			data := Request{
				Type:    "A",
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
				domain.ARecordIdentifier = res.Result.Id
				log.Info("A Record creation successful for domain " + domain.Name)
			} else {
				log.Info("A Record creation failed for domain " + domain.Name)
				log.Info(string(body))
			}
			response.Body.Close()
		}
	}
}
