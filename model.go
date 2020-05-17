package ratelimit

import (
	"golang.org/x/time/rate"
	"time"
)

type RateLimiter struct {
	LastUsed time.Time

	Limiter *rate.Limiter
}
