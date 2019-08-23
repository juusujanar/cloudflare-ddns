package main

import (
	"github.com/juusujanar/cloudflare-ddns/pkg/cloudflare"
	"github.com/juusujanar/cloudflare-ddns/pkg/ipv4"
	"github.com/juusujanar/cloudflare-ddns/pkg/ipv6"
	log "github.com/juusujanar/cloudflare-ddns/pkg/logging"
	"github.com/robfig/cron"
)

var hasIpv6 bool

func main() {
	// Generate config and get a boolean in return if there are domains that need IPv6
	hasIpv6 = cloudflare.GenerateConfig()
	// Check IPs immediately
	checkIPs()
	c := cron.New()
	_, _ = c.AddFunc("@every 1m", checkIPs)
	// Run after every 1 hour
	_, _ = c.AddFunc("@every 1h", checkIPs)
	c.Run()
}

func checkIPs() {
	// Check IPv4 address and carry out updates
	ipv4Addr, validV4 := ipv4.GetIPv4()
	if validV4 {
		log.Info("Valid IPv4 received, updating A records to " + ipv4Addr)
		cloudflare.UpdateARecords(ipv4Addr)
	} else {
		log.Error("Invalid IPv4 address was found: " + ipv4Addr)
	}

	if hasIpv6 {
		ipv6Addr, validV6 := ipv6.GetIPv6()
		if validV6 {
			log.Info("Valid IPv6 received, updating AAAA records to " + ipv6Addr)
			cloudflare.UpdateAAAARecords(ipv6Addr)
		} else {
			log.Error("Invalid IPv6 address was found: " + ipv4Addr)
		}
	}
}