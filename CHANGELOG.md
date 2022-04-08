# Changelog

All notable changes to this chart will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this chart adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0-pre-2.4] - 2022-05-08

### Added

- resource `metaport_cluster`
- resource `metaport_cluster_attachment`
- import functionality to `metaport_cluster_attachment` and `metaport_attachment`

## [1.0.0-pre-2.3] - 2022-03-10

### Fixed

- Creating `metanetworks_mapped_service_alias` would sometimes produce an inconsistent state

## [1.0.0-pre-2.2] - 2022-03-10

### Fixed

- Destroy `metanetworks_mapped_service_alias` from the state if `mapped_service` is removed outside of terraform

## [1.0.0-pre-2.1] - 2022-03-10

### Fixed

- Destroy `metanetworks_mapped_service_alias` from the state if they are removed outside of terraform

### Updated

- Documentation, has been updated. It's generated automatically based on the schemas using `tfplugindocs`

## [1.0.0-pre-2] - 2022-03-09

### Added

- data source `protocol_group`
- data source `protocol_groups`

### Updated

- Documentation, has been updated. It's generated automatically based on the schemas using `tfplugindocs`
