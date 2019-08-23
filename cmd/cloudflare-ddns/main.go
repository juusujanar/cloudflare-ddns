package main

import (
	"github.com/juusujanar/cloudflare-ddns/pkg/cloudflare"
	"github.com/juusujanar/cloudflare-ddns/pkg/ipv4"
	"github.com/juusujanar/cloudflare-ddns/pkg/ipv6"
	"log"
)
func main() {
	ipv4Addr, validV4 := ipv4.GetIPv4()
	ipv6Addr, validV6 := ipv6.GetIPv6()
	if validV4 {
		log.Print("Valid IPv4 received, updating A record to " + ipv4Addr)
		cloudflare.UpdateARecord(ipv4Addr)
	}
	if validV6 {
		log.Print("Valid IPv6 received, updating AAAA record to " + ipv6Addr)
		cloudflare.UpdateAAAARecord(ipv6Addr)
	}
}