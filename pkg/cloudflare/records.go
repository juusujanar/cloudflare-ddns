package cloudflare

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	log "github.com/juusujanar/cloudflare-ddns/pkg/logging"
)

// GetADNSRecordIdentifier pulls domain's A (IPv4) DNS Record Identifier from CloudFlare API
func GetADNSRecordIdentifier(zoneIdentifier string, domain string) (string, bool) {
	client := &http.Client{}
	dest := "https://api.cloudflare.com/client/v4/zones/" + zoneIdentifier + "/dns_records?type=A&name=" + domain
	req, err := http.NewRequest(http.MethodGet, dest, nil)
	if err != nil {
		log.Error("Request preparing failed.")
		return "", true
	}

	req.Header.Set("X-Auth-Email", Config.Email)
	req.Header.Set("X-Auth-Key", Config.ApiToken)
	req.Header.Set("Content-Type", "application/json")
	req.Close = true

	response, err := client.Do(req)
	if err != nil {
		log.Error("GetADNSRecordIdentifier: Request failed.")
		return "", true
	}

	defer response.Body.Close()

	res := CFMultiResultResponse{}
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&res)
	if err != nil {
		log.Error("GetAAAADNSRecordIdentifier: Response parsing failed.")
		return "", true
	}

	if len(res.Result) == 0 {
		return "", true
	} else if len(res.Result) != 1 {
		log.WithFields(logrus.Fields{"response": res}).Error("Multiple records were found.")
		return "", true
	} else {
		return res.Result[0].Id, false
	}
}

// GetAAAADNSRecordIdentifier pulls domain's AAAA (IPv6) DNS Record Identifier from CloudFlare API
func GetAAAADNSRecordIdentifier(zoneIdentifier string, domain string) (string, bool) {
	client := &http.Client{}
	dest := "https://api.cloudflare.com/client/v4/zones/" + zoneIdentifier + "/dns_records?type=AAAA&name=" + domain
	req, err := http.NewRequest(http.MethodGet, dest, nil)
	if err != nil {
		log.Error("Request preparing failed.")
		return "", true
	}

	req.Header.Set("X-Auth-Email", Config.Email)
	req.Header.Set("X-Auth-Key", Config.ApiToken)
	req.Header.Set("Content-Type", "application/json")
	req.Close = true

	response, err := client.Do(req)
	if err != nil {
		log.Error("GetAAAADNSRecordIdentifier: Request failed.")
		return "", true
	}

	defer response.Body.Close()

	res := CFMultiResultResponse{}
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&res)
	if err != nil {
		log.Error("GetAAAADNSRecordIdentifier: Response parsing failed.")
		return "", true
	}

	if len(res.Result) == 0 {
		return "", true
	} else if len(res.Result) != 1 {
		log.WithFields(logrus.Fields{"response": res}).Error("Multiple records were found.")
		return "", true
	} else {
		return res.Result[0].Id, false
	}
}
