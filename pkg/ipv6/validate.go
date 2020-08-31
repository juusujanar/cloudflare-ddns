package ipv6

import (
	"net"
	"strings"
)

// ValidateIPv6 ensures that we got IPv6 address instead of IPv4
func ValidateIPv6(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && strings.Contains(str, ":")
}