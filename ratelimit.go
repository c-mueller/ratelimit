package ratelimit

import (
	"context"
	"github.com/coredns/coredns/plugin"
	"github.com/miekg/dns"
	"time"
)

type RateLimitPlugin struct {
	LimitDuration time.Duration
	LimitCount    int

	Next plugin.Handler
}

func (e *RateLimitPlugin) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {

	return plugin.NextOrFailure(e.Name(), e.Next, ctx, w, r)
}

// Name implements the Handler interface.
func (e *RateLimitPlugin) Name() string { return "ratelimit" }
