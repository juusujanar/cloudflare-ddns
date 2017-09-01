##################################################
# 
# Bash script that is ran every hour to update the
# DNS A record for your domain in CloudFlare
# Can be customized according to Cloudflare API
#
##################################################

FROM alpine:latest
MAINTAINER Janar Juusu <janar.juusu@gmail.com>

# Add necessary files to the container
ADD cloudflare-ddns.sh /cloudflare-ddns.sh
ADD crontab.txt /crontab.txt
ADD entrypoint.sh /entrypoint.sh

# Give permissions and run cron
RUN chmod 755 /cloudflare-ddns.sh /entrypoint.sh
RUN /usr/bin/crontab /crontab.txt

# Install curl
RUN apk update && apk add curl

# Run the script for the first time also
RUN /cloudflare-ddns.sh

CMD ["/entrypoint.sh"]