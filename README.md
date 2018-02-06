# CloudFlare Dynamic DNS
##### Bash script running as cronjob in Docker to update CloudFlare DNS records

Records are updated every hour at 38 minutes using https://www.ipify.org/ to get IP.
DNS record must be made beforehand, otherwise it's record ID is not available.

Currently only supports IPv4.

Container is based on latest Alpine Linux and uses cURL to carry out the operations.

### Configuration

Configuration is done via environment variables.

ZONE_IDENTIFIER - Check it on domain overview page on CloudFlare  
DNS_RECORD - You can find it via CloudFlare API - https://api.cloudflare.com/#dns-records-for-a-zone-list-dns-records  
AUTH_EMAIL - Your CloudFlare email  
AUTH_TOKEN - Your CloudFlare authentication token  
DOMAIN - Domain which A record you want to update  
PROXIED - Whether to proxy through CloudFlare or not (true/false)  

Example:
```
  docker run -d -e "ZONE_IDENTIFIER=875d12a39455fa2c2685856e45450ce" -e "DNS_RECORD=875d12a39455fa2c2685856e45450ce" -e "AUTH_EMAIL=user@example.com" -e "AUTH_TOKEN=875d12a39455fa2c2685856e45450ce" -e "DOMAIN=example.com" -e "PROXIED=true" janarj/cloudflare-ddns
```

## License
Licensed under GNU General Public License v3.0
