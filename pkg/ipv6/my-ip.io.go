package ipv6

import (
	log "github.com/juusujanar/cloudflare-ddns/pkg/logging"
	"io/ioutil"
	"net/http"
)

// GetMyIpIoV4 gets IPv6 from my-ip.io
func GetMyIpIoV6() (string, bool) {
	resp, err := http.Get("https://api6.my-ip.io/ip")
	if err != nil {
		log.Error("myip.com - could not get IPv6 address")
		return "", false
	} else {
		defer func() {
			_ = resp.Body.Close()
		}()
		body, _ := ioutil.ReadAll(resp.Body)
		return string(body), ValidateIPv6(string(body))
	}
}
