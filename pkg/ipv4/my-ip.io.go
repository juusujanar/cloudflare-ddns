package ipv4

import (
	log "github.com/juusujanar/cloudflare-ddns/pkg/logging"
	"io/ioutil"
	"net/http"
)

// GetMyIpIoV4 gets IPv4 from my-ip.io
func GetMyIpIoV4() (string, bool) {
	resp, err := http.Get("https://api4.my-ip.io/ip")
	if err != nil {
		log.Error("myip.com - could not get IPv4 address")
		return "", false
	} else {
		defer func() {
			_ = resp.Body.Close()
		}()
		body, _ := ioutil.ReadAll(resp.Body)
		return string(body), ValidateIPv4(string(body))
	}
}