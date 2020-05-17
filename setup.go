package ratelimit

import (
	"github.com/caddyserver/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
	clog "github.com/coredns/coredns/plugin/pkg/log"
	"strconv"
	"sync"
	"time"
)

func init() {
	caddy.RegisterPlugin("ratelimit", caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

var log = clog.NewWithPlugin("ads")

func setup(c *caddy.Controller) error {
	limit := 10

	if c.NextArg() && c.NextArg() {
		val := c.Val()
		if val != "" {
			l, err := strconv.Atoi(val)
			if err != nil {
				return plugin.Error("Failed to parse limit", err)
			}
			limit = l
		}
	}
	c.Next()

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {

		plugin := &RateLimitPlugin{
			LimitCount:  limit,
			Mutex:       &sync.RWMutex{},
			LimitQueues: make(map[string]*RateLimiter),
			Next:        next,
		}

		ticker := time.NewTicker(time.Minute)
		go func() {
			for range ticker.C {
				plugin.removeUnusedLimits()
			}
		}()

		return plugin
	})

	return nil
}
