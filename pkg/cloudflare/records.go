package cloudflare

import (
	"encoding/json"
	log "github.com/juusujanar/cloudflare-ddns/pkg/logging"
	"io/ioutil"
	"net/http"
)

// GetADNSRecordIdentifier pulls domain's A (IPv4) DNS Record Identifier from CloudFlare API
func GetADNSRecordIdentifier(zoneIdentifier string, domain string) (string, bool) {
	client := &http.Client{}
	dest := "https://api.cloudflare.com/client/v4/zones/" + zoneIdentifier + "/dns_records?type=A&name=" + domain
	req, err := http.NewRequest(http.MethodGet, dest, nil)
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

	defer response.Body.Close()

	res := CFMultiResultResponse{}
	body, err := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal(err)
	}

	if len(res.Result) == 0 {
		return "", true
	} else if len(res.Result) != 1 {
		log.Error("Multiple records were found: " + string(body))
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
		log.Fatal(err)
	}

	req.Header.Set("X-Auth-Email", Config.Email)
	req.Header.Set("X-Auth-Key", Config.ApiToken)
	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	res := CFMultiResultResponse{}
	body, err := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal(err)
	}

	if len(res.Result) == 0 {
		return "", true
	} else if len(res.Result) != 1 {
		log.Error("Multiple records were found: " + string(body))
		return "", true
	} else {
		return res.Result[0].Id, false
	}
}
