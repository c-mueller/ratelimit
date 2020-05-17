package ratelimit

import (
	"github.com/caddyserver/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	"time"
)

func init() {
	caddy.RegisterPlugin("ratelimit", caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	c.Next()

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {

		return &RateLimitPlugin{
			LimitDuration: time.Minute,
			LimitCount:    300,
		}
	})

	return nil
}
