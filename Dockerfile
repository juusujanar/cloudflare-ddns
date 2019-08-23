FROM golang:1.12 as builder
WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.9
LABEL maintainer="janar@juusujanar.eu"

WORKDIR /srv

# Add necessary files to the container
COPY --from=builder /app/main .
ADD crontab.txt entrypoint.sh /srv/

# Install cURL, give permissions and run cron
RUN apk add --no-cache curl ca-certificates && \
    chmod 755 /srv/main /srv/entrypoint.sh && \
    /usr/bin/crontab /srv/crontab.txt

CMD ["/srv/entrypoint.sh"]
