#!/bin/sh

# Copyright (C) 2017 Janar Juusu

# IP is automagically detected
ip=$(curl https://api.ipify.org)

# Enter below information according to your account

# zone_identifier can be found on domain overview page
zone_identifier="023e105f4ecef8ad9ca31a8372d0c353"

# dns_record can be found via CloudFlare API
# https://api.cloudflare.com/#dns-records-for-a-zone-list-dns-records
dns_record="372e67954025e0ba6aaa6d586b9e0b59"

# your user account
auth_email="user@example.com"

# your cloudflare global api key, found in your profile
auth_key="c2547eb745079dac9320b638f5e225cf483cc5cfdda41"

# the domain that you are updating
domain="example.com"


# this does the updating
curl -X PUT "https://api.cloudflare.com/client/v4/zones/"$zone_identifier"/dns_records/"$dns_record \
     -H "X-Auth-Email: "$auth_email \
     -H "X-Auth-Key: "$auth_key \
     -H "Content-Type: application/json" \
     --data '{"type":"A","name":"'"$domain"'","content":"'"$ip"'","ttl":120,"proxied":false}'