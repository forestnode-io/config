package config_test

import (
	"strings"
	"testing"

	"github.com/forestnode-io/config"
	"github.com/raphaelreyna/policyauthor"
	"github.com/raphaelreyna/policyauthor/pkg/conditions"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMerge_YAML_V3(t *testing.T) {
	const (
		base = `
server:
  addr: :8080
  tls:
    cert: /path/to/cert
`
		update = `
server:
  addr: :8081
  routing:
  - value: "https://rreyna.dev"
    conditions:
    - type: exists
      spec:
        key: "remote_addr"
`
	)

	policyauthor.RegisterConditions(conditions.AllConditionsMap())

	provider, err := config.NewYAML(
		config.Source(strings.NewReader(base)),
		config.Source(strings.NewReader(update)),
	)
	require.NoError(t, err, "couldn't construct provider")

	myconfig := struct {
		Server struct {
			Addr    string
			TLS     any
			Routing *policyauthor.PolicyEngine
		}
	}{}

	err = provider.Get("").Populate(&myconfig)
	require.NoError(t, err, "couldn't populate config")

	assert.Equal(t, ":8081", myconfig.Server.Addr)
	assert.Equal(t, "/path/to/cert", myconfig.Server.TLS.(map[string]any)["cert"])

	m := map[string]any{
		"remote_addr": "1.2.3.4",
	}
	policyVal, hit, err := myconfig.Server.Routing.Evaluate(m)
	require.NoError(t, err, "couldn't evaluate policy")
	assert.True(t, hit, "expected policy to not hit")
	assert.Equal(t, "https://rreyna.dev", policyVal)

	m = map[string]any{
		"foo": "bar",
	}
	policyVal, hit, err = myconfig.Server.Routing.Evaluate(m)
	require.NoError(t, err, "couldn't evaluate policy")
	assert.False(t, hit, "expected policy to hit")
	assert.Nil(t, policyVal)
}
