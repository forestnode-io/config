# Changelog

## v1.0.0-rc2 (unreleased)

- **[Breaking]** `ValueType` and `GetType` functionality is removed in favor of using
  `reflect.Kind`.
- Skip populating function and value types instead of reporting errors.
- **[Breaking]** `Value.Timestamp` is private, use Value.LastUpdated instead.
- **[Breaking]** Use semantic version paths for yaml and validator packages.
- Let user to skip loading command line provider via `commandLine` parameter.

## v1.0.0-rc1 (26 Jun 2017)

- **[Breaking]** `Provider` interface was trimmed down to 2 methods: `Name` and `Get`