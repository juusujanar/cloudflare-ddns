FROM golang:1.12 as builder
WORKDIR /go/src/github.com/juusujanar/cloudflare-ddns

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/cloudflare-ddns

FROM alpine:3.9
LABEL maintainer="janar@juusujanar.eu"

WORKDIR /srv

# Add necessary files to the container
COPY --from=builder /go/src/github.com/juusujanar/cloudflare-ddns/main .
ADD crontab.txt entrypoint.sh /srv/

# Install cURL, give permissions and run cron
RUN apk add --no-cache curl ca-certificates && \
    chmod 755 /srv/main /srv/entrypoint.sh && \
    /usr/bin/crontab /srv/crontab.txt

CMD ["/srv/entrypoint.sh"]
