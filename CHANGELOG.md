# Changelog

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [v0.2.0] - 2023-09-06

### Added

- Support HEAD, CONNECT, and TRACE HTTP methods.
- Support registering the same path for different HTTP methods.
- Support chaining `router` struct method calls.

### Changed

- Privatise `router.Route` struct method.

## [v0.1.0] - 2023-09-05

### Added

- Add router implementation.
