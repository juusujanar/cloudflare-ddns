package ipv6

import (
	log "github.com/juusujanar/cloudflare-ddns/pkg/logging"
	"io/ioutil"
	"net"
	"net/http"
)

// GetIPv6 gets IP from api6.ipify.org, which gives out only IPv6
func GetIPv6() (string, bool) {
	// For some reason, just GET was using IPv4 and this is one way to force it to IPv6
	// No HTTPS though
	ip, err := net.ResolveIPAddr("ip6", "api6.ipify.org")
	if err != nil {
		// IPv6 not found
		log.Error("Could not get IPv6 connectivity.")
		return "", false
	}
	resp, err := http.Get("http://[" + ip.String() + "]")
	if err != nil {
		log.Error("Could not get IPv6 address.")
		return "", false
	} else {
		defer func() {
			_ = resp.Body.Close()
		}()
		body, _ := ioutil.ReadAll(resp.Body)
		return string(body), ValidateIPv6(string(body))
	}
}
