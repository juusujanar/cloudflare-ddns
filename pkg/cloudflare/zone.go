package cloudflare

import (
	"encoding/json"
	log "github.com/juusujanar/cloudflare-ddns/pkg/logging"
	"io/ioutil"
	"net/http"
	"regexp"
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
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		log.Fatal(err)
	}

	if len(res.Result) == 0 {
		return "", true
	} else if len(res.Result) != 1 {
		log.Error("Multiple zones were found: " + string(body))
		return "", true
	} else {
		return res.Result[0].Id, false
	}
}
