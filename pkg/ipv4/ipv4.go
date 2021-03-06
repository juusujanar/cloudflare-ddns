package ipv4

import (
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	log "github.com/juusujanar/cloudflare-ddns/pkg/logging"
)

// GetIPv4 gets the IP from ipify.org, which only gives out IPv4
func GetIPv4() (string, bool) {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		log.Error("Could not get IPv4 address.")
		return "", false
	} else {
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		return string(body), ValidateIPv4(string(body))
	}
}

// Just to ensure that the IP is IPv4 and not some error or other invalid data
func ValidateIPv4(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && strings.Contains(str, ".")
}