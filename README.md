# CloudFlare Dynamic DNS

Tired of having to update your DNS every time your dynamic IP changes?

This app ensures your DNS records get updated at an hourly interval. IPs are collected using https://www.ipify.org/.
Both IPv4 and IPv6 are supported, however IPv6 must be explicitly turned on for each domain.

If you enable IPv6 and run the Docker container, your
[Docker must have IPv6 connectivity](https://docs.docker.com/config/daemon/ipv6/).
Do take note IPv6 address might be different for the container and the service you want to access.

App is built using Golang 1.13 and final Docker container is based on Scratch.

Latest image (janarj/cloudflare-ddns:latest) may contain unstable versions.
Use versioned image tags for stable releases (janarj/cloudflare-ddns:2.0)

Versions 1.x are no longer supported (previously made in Bash and ran on Alpine).


## How to use

Config file must be in one of the following locations:
- /etc/cf-ddns
- $HOME/.cf-ddns
- same directory as the binary

and be named config.json, config.yml or config.toml.
JSON and YAML examples are in configs folder.
Detailed information is below.

### To use Docker image:
```
  docker run -d \
  -v $PWD/config.json:/etc/cf-ddns/config.json \ 
  janarj/cloudflare-ddns:2.0
```

### Or if you want to build it yourself:
- `go build -o ./app ./cmd/cloudflare-ddns`
- Run the app binary with `./app`.

## Configuration

### CloudFlare API token
To get CloudFlare API token, log in to CloudFlare dashboard, go to your profile, view Global API key.
It must be included as well as email or the program will error out immediately.

### Domains
**Domain name** is required to be set.

**TTL** is optional and defaults to 1, which is CloudFlare's automatic option.

**IPv6** turns on IPv6 support for that domain. Optional, defaults to false.

**Proxied** turns on CloudFlare proxying for that domain. Optional, defaults to false.

## License
Licensed under GNU General Public License v3.0.

## Contributing
I'm open to pull requests.
