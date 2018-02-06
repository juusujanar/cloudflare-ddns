#!/bin/sh

# Copyright (C) 2018 Janar Juusu

# IP is automagically detected
ip=$(curl https://api.ipify.org)

# This does the updating
curl -X PUT "https://api.cloudflare.com/client/v4/zones/"$ZONE_IDENTIFIER"/dns_records/"$DNS_RECORD \
     -H "X-Auth-Email: "$AUTH_EMAIL \
     -H "X-Auth-Key: "$AUTH_TOKEN \
     -H "Content-Type: application/json" \
     --data '{"type":"A","name":"'"$DOMAIN"'","content":"'"$ip"'","ttl":300,"proxied":"'"$PROXIED"'"}'
