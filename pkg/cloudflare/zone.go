package cloudflare

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/sirupsen/logrus"

	log "github.com/juusujanar/cloudflare-ddns/pkg/logging"
)

// GetZoneIdentifier pulls domain Zone Identifier from CloudFlare API
func GetZoneIdentifier(domain string) (string, bool) {
	client := &http.Client{}
	// Regex extracts second-level top level domain from the string
	// So it works for both example.com as well as example.co.uk domains
	r, _ := regexp.Compile(`[^.]*\.[^.]{2,3}(?:\.[^.]{2,3})?$`)
	match := r.FindString(domain)
	dest := "https://api.cloudflare.com/client/v4/zones?name=" + match
	req, err := http.NewRequest(http.MethodGet, dest, nil)
	if err != nil {
		log.Error("GetZoneIdentifier: Request preparing failed.")
		log.Error(err)
		return "", true
	}

	req.Header.Set("X-Auth-Email", Config.Email)
	req.Header.Set("X-Auth-Key", Config.ApiToken)
	req.Header.Set("Content-Type", "application/json")
	req.Close = true

	response, err := client.Do(req)
	if err != nil {
		log.Error("GetZoneIdentifier: Request failed.")
		log.Error(err)
		return "", true
	}

	defer response.Body.Close()

	res := CFMultiResultResponse{}
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&res)
	if err != nil {
		log.Error("GetZoneIdentifier: Response parsing failed.")
		log.Error(err)
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
