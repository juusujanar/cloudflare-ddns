# CloudFlare Dynamic DNS
##### Bash script running as cronjob in Docker to update CloudFlare DNS records

Records are updated every hour at 38 minutes using https://www.ipify.org/ and https://icanhazip.com to get IP.
DNS record(s) must be created beforehand, otherwise it's record ID is not available.

Supports both IPv4 and IPv6, IPv6 can be turned off with the IPv6 variable. If you enable IPv6, you must also include IPv6_DNS_RECORD and your [Docker must have IPv6 connectivity](https://docs.docker.com/config/daemon/ipv6/).

TTL is set to 1 hour.

Container is based on Alpine Linux 3.8 and uses cURL to carry out the requests.


### Finding CloudFlare identifiers and DNS records

You can find your zone identifier in CloudFlare dashboard, under DNS section. Press the API link below your records table and it will be shown there.

Then you can run the following command to get list of the DNS records along with their IDs.

```
  curl -X GET https://api.cloudflare.com/client/v4/zones/ZONE_IDENTIFIER/dns_records \
  -H "X-Auth-Email: YOUR_EMAIL@example.com" \
  -H "X-Auth-Key: CLOUDFLARE_API_KEY" \
  -H "Content-Type: application/json"
```


### Configuration and running

Configuration is done via environment variables.

ZONE_IDENTIFIER - Check it on domain overview page on CloudFlare  
DNS_RECORD - You can find it via CloudFlare API - https://api.cloudflare.com/#dns-records-for-a-zone-list-dns-records  
IPv6_DNS_RECORD - Same as above, needed if you want to update AAAA records.
AUTH_EMAIL - Your CloudFlare email  
AUTH_TOKEN - Your CloudFlare API token  
DOMAIN - Domain which A/AAAA record(s) you want to update.
PROXIED - Whether to proxy through CloudFlare or not (true/false)  
IPv6 - Turn off IPv6 checking. true/false, defaults to true.

Example:
```
  docker run -d \
  -e "ZONE_IDENTIFIER=875d12a39455fa2c2685856e45450ce" \
  -e "DNS_RECORD=875d12a39455fa2c2685856e45450ce" \
  -e "AUTH_EMAIL=user@example.com" \
  -e "AUTH_TOKEN=875d12a39455fa2c2685856e45450ce" \
  -e "DOMAIN=example.com" \
  -e "IPv6=true" \
  -e "IPv6_DNS_RECORD=someothertoken1938175" \
  -e "PROXIED=true" janarj/cloudflare-ddns
```

## License
Licensed under GNU General Public License v3.0
