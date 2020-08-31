package ipv4

import (
	"net"
	"strings"
)

// Just to ensure that the IP is IPv4 and not some error or other invalid data
func ValidateIPv4(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && strings.Contains(str, ".")
}