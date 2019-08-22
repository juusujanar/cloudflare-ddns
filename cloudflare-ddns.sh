#!/bin/sh

# Copyright (C) 2019 Janar Juusu

# IPv4 address is taken from ipify.org
ip=$(curl -s https://api.ipify.org)
printf "[+] IPv4 address found: $ip\n"

ttl="${TTL:-3600}"

# Update the A record in CloudFlare
curl -s -X PUT "https://api.cloudflare.com/client/v4/zones/"$ZONE_IDENTIFIER"/dns_records/"$DNS_RECORD \
     -H "X-Auth-Email: "$AUTH_EMAIL \
     -H "X-Auth-Key: "$AUTH_TOKEN \
     -H "Content-Type: application/json" \
     --data '{"type":"A","name":"'"$DOMAIN"'","content":"'"$ip"'","ttl":"'"$ttl"'","proxied":'"$PROXIED"'}'

check_ipv6="${IPv6:-true}"

if [[ "$check_ipv6" = 'true' ]]; then
  printf "\n\n[+] Checking IPv6\n"
  
  # icanhazip.com supports both IPv4 and IPv6
  ipv6=$(curl -s https://icanhazip.com)

  # Ensure that it is an IPv6 address
  hasIPv6=$(echo $ipv6 | awk -F: "NF>4")

  if [[ -z "${hasIPv6}" ]]; then
    printf "[-] Got an IPv4 address, ignoring: $ipv6\n"
  else
    printf "[+] IPv6 address found: $ipv6\n"

    if [[ -z "${IPv6_DNS_RECORD}" ]]; then
      printf "[-] No IPv6_DNS_RECORD found, not updating AAAA records.\n"
    else 
      curl -s -X PUT "https://api.cloudflare.com/client/v4/zones/"$ZONE_IDENTIFIER"/dns_records/"$IPv6_DNS_RECORD \
       -H "X-Auth-Email: "$AUTH_EMAIL \
       -H "X-Auth-Key: "$AUTH_TOKEN \
       -H "Content-Type: application/json" \
       --data '{"type":"AAAA","name":"'"$DOMAIN"'","content":"'"$ipv6"'","ttl":3600,"proxied":'"$PROXIED"'}'
    fi
  fi
fi

printf "\n\n[+] Done!\n"
