#!/bin/sh

# First run
/srv/main

# Start cron
/usr/sbin/crond -f -l 8
