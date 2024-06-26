# :fishing_pole_and_fish: config [![GoDoc][doc-img]][doc]

Convenient, injection-friendly YAML configuration.

## Installation

```
go get -u github.com/forestnode-io/config
```

Note that config only supports the two most recent minor versions of Go.

## Quick Start

```golang
// Model your application's configuration using a Go struct.
type cfg struct {
    Parameter string
}

// Two sources of configuration to merge.
base := strings.NewReader("module: {parameter: foo}")
override := strings.NewReader("module: {parameter: bar}")

// Merge the two sources into a Provider. Later sources are higher-priority.
// See the top-level package documentation for details on the merging logic.
provider, err := config.NewYAML(config.Source(base), config.Source(override))
if err != nil {
    panic(err) // handle error
}

var c cfg
if err := provider.Get("module").Populate(&c); err != nil {
  panic(err) // handle error
}

fmt.Printf("%+v\n", c)
// Output:
// {Parameter:bar}
```

## Development Status: Stable

All APIs are finalized, and no breaking changes will be made in the 1.x series
of releases. Users of semver-aware dependency management systems should pin
config to `^1`.

---

Released under the [MIT License](LICENSE.txt).

[doc-img]: http://img.shields.io/badge/GoDoc-Reference-blue.svg
[doc]: https://godoc.org/github.com/forestnode-io/config

[ci-img]: https://github.com/forestnode-io/config/actions/workflows/go.yml/badge.svg
[ci]: https://github.com/forestnode-io/config/actions/workflows/go.yml

[cov-img]: https://codecov.io/gh/forestnode-io/config/branch/master/graph/badge.svg
[cov]: https://codecov.io/gh/forestnode-io/config/branch/master
