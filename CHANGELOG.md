# Changelog

## [2.0] - Unreleased
### Changed
- Rewrote tool to Golang 1.12
### Added
- JSON configuration file support
### In progress
- Multiple domain support
- Read Zone and DNS record identifiers from CloudFlare API
- Create record if does not exist

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
