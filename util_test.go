package ratelimit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseIP(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			"ipv4-localhost-2",
			"127.0.0.1:12",
			"127.0.0.1",
		},
		{
			"ipv4-localhost-3",
			"127.0.0.1:123",
			"127.0.0.1",
		},
		{
			"ipv4-localhost-4",
			"127.0.0.1:1234",
			"127.0.0.1",
		},
		{
			"ipv4-localhost-5",
			"127.0.0.1:1234",
			"127.0.0.1",
		},
		{
			"ipv6-localhost",
			"::1:1234",
			"::1",
		},
		{
			"ipv6-localhost-bracketed",
			"[::1]:1234",
			"[::1]",
		},
		{
			"ipv6-random",
			"2a03:4000:32:274:84d2:4eff:fe00:bfe9:1234",
			"2a03:4000:32:274:84d2:4eff:fe00:bfe9",
		},
		{
			"ipv6-random-bracketed",
			"[2a03:4000:32:274:84d2:4eff:fe00:bfe9]:1234",
			"[2a03:4000:32:274:84d2:4eff:fe00:bfe9]",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.output, getAddrWithoutPort(test.input))
		})
	}
}
