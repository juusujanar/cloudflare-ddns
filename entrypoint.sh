#!/bin/sh

# First run
/srv/cloudflare-ddns.sh

# Start cron
/usr/sbin/crond -f -l 8
