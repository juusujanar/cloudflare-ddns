package ipv4

import (
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

// GetIPv4 gets the IP from ipify.org, which only gives out IPv4
func GetIPv4() (string, bool) {
	resp, err := http.Get("https://api.ipify.org")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), ValidateIPv4(string(body))
}

// Just to ensure that the IP is IPv4 and not some error or other invalid data
func ValidateIPv4(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && strings.Contains(str, ".")
}