package ipv6

import (
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

// GetIPv6 gets IP from icanhazip.com, which supports both IPv4 and IPv6
func GetIPv6() (string, bool) {
	resp, err := http.Get("https://icanhazip.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body), ValidateIPv6(string(body))
}

// ValidateIPv6 ensures that we got IPv6 address instead of IPv4
func ValidateIPv6(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && strings.Contains(str, ":")
}