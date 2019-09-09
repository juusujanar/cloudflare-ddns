package cloudflare

import (
	"github.com/juusujanar/cloudflare-ddns/pkg/config"
	log "github.com/juusujanar/cloudflare-ddns/pkg/logging"
)

type Request struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
	Ttl     int    `json:"ttl"`
	Proxied bool   `json:"proxied"`
}

type CFResult struct {
	Id string `json:"id"`
}

type CFMultiResultResponse struct {
	Success bool		`json:"success"`
	Result  []CFResult	`json:"result"`
}

type CFSingleResultResponse struct {
	Success bool		`json:"success"`
	Result  CFResult	`json:"result"`
}

type Domain struct {
	Name string
	ZoneIdentifier string
	ARecordIdentifier string
	AAAARecordIdentifier string
	TTL int
	Proxied bool
	IPv6 bool
}

type Configuration struct {
	Domains []Domain
	Email string
	ApiToken string
}

var Config Configuration

// GenerateConfig reads the configuration file and collects all domains' details
// Returns boolean whether we have any domains that need IPv6
func GenerateConfig() bool {
	// Flag for storing whether we have any IPv6 requests
	// If there are none, we can skip IPv6 requests
	// It will be returned to main
	hasIpv6 := false
	file := config.ReadConfig()
	conf := Configuration{
		Email: file.Email,
		ApiToken: file.ApiToken,
	}
	// First saving so that email and token are available for
	// identifier requests through the Config object
	// Probably bad practice, but works for now
	Config = conf
	for _, domain := range file.Domains {
		d := Domain{
			Name: domain.Name,
			// TTL defaults to 0 if not set
			TTL: domain.TTL,
			// Proxied defaults to false if not set
			Proxied: domain.Proxied,
			// IPv6 defaults to false if not set
			IPv6: domain.IPv6,
		}

		// Let's default TTL to 1 aka CF Automatic.
		if d.TTL == 0 {
			d.TTL = 1
		}

		zone, err := GetZoneIdentifier(domain.Name)
		if err {
			log.Error("Zone not found, skipping domain " + domain.Name)
			continue
		}
		log.Debug("Got Zone Identifier: " + zone)
		d.ZoneIdentifier = zone

		aRecordIdentifier, err := GetADNSRecordIdentifier(zone, domain.Name)
		if err {
			log.Info("A record not found on domain " + domain.Name + ", needs creation.")
		} else {
			d.ARecordIdentifier = aRecordIdentifier
		}

		if domain.IPv6 {
			hasIpv6 = true
			aaaaRecordIdentifier, err := GetAAAADNSRecordIdentifier(zone, domain.Name)
			if err {
				log.Info("AAAA record not found on domain " + domain.Name + ", needs creation.")
			} else {
				d.AAAARecordIdentifier = aaaaRecordIdentifier
			}
		}
		conf.Domains = append(conf.Domains, d)
	}

	// Apply the global configuration
	Config = conf
	return hasIpv6
}