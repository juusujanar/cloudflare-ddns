package ipv4

import (
	log "github.com/juusujanar/cloudflare-ddns/pkg/logging"
	"io/ioutil"
	"net/http"
)

// GetIpifyV4 gets IPv4 from ipify.org
func GetIpifyV4() (string, bool) {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		log.Error("ipify.org - could not get IPv4 address")
		return "", false
	} else {
		defer func() {
			_ = resp.Body.Close()
		}()
		body, _ := ioutil.ReadAll(resp.Body)
		return string(body), ValidateIPv4(string(body))
	}
}
