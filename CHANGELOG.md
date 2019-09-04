# Changelog

## [2.0] - Unreleased
**BREAKING CHANGE:** This release no longer uses environment variables. Create a configuration file instead.
### Changed
- Rewrote tool to Golang 1.13 (with go modules)
- Switch base image to scratch
- Run Docker as non-root user
- Use api6.ipify.org for IPv6
  - Currently does not use HTTPS due to an issue
### Added
- JSON/YAML/TOML configuration file support
  - Multiple domain support
  - Each domain now has individual TTL time & IPv6 and proxied flags
- Pull Zone and DNS record identifiers from CloudFlare API
  - No longer to fetch them manually beforehand
- Create A/AAAA record if does not exist
### Removed
- Environmental variable support


## [1.2] - 2019-08-22
### Added
- TTL can be configured via env variable.

### Changed
- Updated to Alpine 3.9


## [1.1] - 2019-01-13
### Added
- IPv6 and AAAA records support
- Run script on container start

### Changed
- Send cron job logs to stdout
- Re-wrote some parts of code for clarity
- Default TTL from 5 minutes to 1 hour
- Moved all files in container to /srv
- Dockerfile improvements
- Logging improvements
