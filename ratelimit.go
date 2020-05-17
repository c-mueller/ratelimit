package ratelimit

import (
	"context"
	"fmt"
	"github.com/coredns/coredns/plugin"
	"github.com/miekg/dns"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

type RateLimitPlugin struct {
	LimitCount int

	Mutex       *sync.RWMutex
	LimitQueues map[string]*RateLimiter

	Next plugin.Handler
}

func (e *RateLimitPlugin) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	remoteAddr := getAddrWithoutPort(w.RemoteAddr().String())

	e.Mutex.Lock()
	if e.LimitQueues[remoteAddr] == nil {
		e.LimitQueues[remoteAddr] = &RateLimiter{
			Limiter: rate.NewLimiter(rate.Limit(e.LimitCount), e.LimitCount*2),
		}
	}
	e.LimitQueues[remoteAddr].LastUsed = time.Now()
	if e.LimitQueues[remoteAddr].Limiter.Allow() {
		e.Mutex.Unlock()
		return plugin.NextOrFailure(e.Name(), e.Next, ctx, w, r)
	} else {
		fmt.Printf("%s has exceeded the Ratelimit\n", remoteAddr)
		// Rate limit has been Exceeded, the query is refused
		e.Mutex.Unlock()
		return dns.RcodeServerFailure, nil
	}
}

// Name implements the Handler interface.
func (e *RateLimitPlugin) Name() string { return "ratelimit" }
