FROM alpine:3.8
LABEL maintainer="janar@juusujanar.eu"

# Add necessary files to the container
ADD cloudflare-ddns.sh crontab.txt entrypoint.sh /srv/

# Install cURL, give permissions and run cron
RUN apk add --no-cache curl && \
    chmod 755 /srv/cloudflare-ddns.sh /srv/entrypoint.sh && \
    /usr/bin/crontab /srv/crontab.txt

CMD ["/srv/entrypoint.sh"]
