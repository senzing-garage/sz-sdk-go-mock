# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
[markdownlint](https://dlaa.me/markdownlint/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

-

## [0.8.0] - 2024-08-22

### Changed in 0.8.0

- Change from `g2` to `sz`/`er`

## [0.7.2] - 2024-06-26

### Changed in 0.7.2

- Synchronized with [sz-sdk-go-core](https://github.com/senzing-garage/sz-sdk-go-core) and [sz-sdk-go-grpc](https://github.com/senzing-garage/sz-sdk-go-grpc)
- Updated dependencies

## [0.7.1] - 2024-05-09

### Added in 0.7.1

- `SzDiagnostic.GetFeature`
- `SzEngine.FindInterestingEntitiesByEntityId`
- `SzEngine.FindInterestingEntitiesByRecordId`

### Deleted in 0.7.1

- `SzEngine.GetRepositoryLastModifiedTime`

## [0.7.0] - 2024-04-26

### Changed in 0.7.0

-

## [0.6.0] - 2024-02-27

### Changed in 0.6.0

- Updated dependencies
  - github.com/senzing-garage/g2-sdk-go v0.10.0
- Deleted methods not used in V4

## [0.5.0] - 2024-01-26

### Changed in 0.5.0

- Renamed module to `github.com/senzing-garage/g2-sdk-go-moc`
- Refactor to [template-go](https://github.com/senzing-garage/template-go)
- Update dependencies
  - github.com/senzing-garage/g2-sdk-go v0.9.0

## [0.4.0] - 2024-01-02

### Changed in 0.4.0

- Refactor to [template-go](https://github.com/senzing-garage/template-go)
- Update dependencies
  - github.com/senzing-garage/go-common v0.4.0
  - github.com/senzing-garage/go-logging v1.4.0
  - github.com/senzing-garage/go-observing v0.3.0
  - github.com/senzing/g2-sdk-go v0.8.0

## [0.3.3] - 2023-12-12

### Added in 0.3.3

- `ExportCSVEntityReportIterator` and `ExportJSONEntityReportIterator`

### Changed in 0.3.3

- Update dependencies
  - github.com/senzing/g2-sdk-go v0.7.6

## [0.3.2] - 2023-10-18

### Changed in 0.3.2

- Update dependencies
  - github.com/senzing/g2-sdk-go v0.7.4
  - github.com/senzing-garage/go-common v0.3.1
  - github.com/senzing-garage/go-logging v1.3.3
  - github.com/senzing-garage/go-observing v0.2.8

## [0.3.1] - 2023-10-12

### Changed in 0.3.1

- Changed from `int` to `int64` where required by the SenzingAPI
- Update dependencies
  - github.com/senzing/g2-sdk-go v0.7.3

### Deleted in 0.3.1

- `g2product.ValidateLicenseFile`
- `g2product.ValidateLicenseStringBase64`

## [0.3.0] - 2023-10-03

### Changed in 0.3.0

- Supports SenzingAPI 3.8.0
- Deprecated functions have been removed
- Update dependencies
  - github.com/senzing/g2-sdk-go v0.7.0

## [0.2.2] - 2023-09-01

### Changed in 0.2.2

- Last version before SenzingAPI 3.8.0

## [0.2.1] - 2023-08-06

### Changed in 0.2.1

- Refactor to `template-go`
- Migrate to `go-logging`
- Update dependencies
  - github.com/senzing/g2-sdk-go v0.6.8
  - github.com/senzing-garage/go-common v0.2.11
  - github.com/senzing-garage/go-logging v1.3.2
  - github.com/senzing-garage/go-observing v0.2.7
  - github.com/stretchr/testify v1.8.4

## [0.2.0] - 2023-05-26

### Changed in 0.2.0

- Modified `g2config.Load()` method signature
- Added `gosec`
- Update dependencies
  - github.com/senzing/g2-sdk-go v0.6.4

## [0.1.4] - 2023-05-10

### Changed in 0.1.4

- Added `GetObserverOrigin()` and `SetObserverOrigin()` to g2* packages
- Update dependencies
  - github.com/senzing/g2-sdk-go v0.6.2
  - github.com/senzing-garage/go-observing v0.2.2

## [0.1.3] - 2023-04-21

### Changed in 0.1.3

## [0.1.2] - 2023-04-21

### Changed in 0.1.2

- Changed `SetLogLevel(ctx context.Context, logLevel logger.Level)` to `SetLogLevel(ctx context.Context, logLevelName string)`

## [0.1.1] - 2023-02-21

### Changed in 0.1.1

- Change GetSdkId() signature.

## [0.1.0] - 2023-02-17

### Added to 0.1.0

- Initial artifacts
