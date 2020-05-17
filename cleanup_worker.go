package ratelimit

import "time"

func (e *RateLimitPlugin) removeUnusedLimits() {
	e.Mutex.RLock()
	todelete := make([]string, 0)
	removalThreshold := time.Now().Add(-1 * time.Minute * 10)
	for ip, limiter := range e.LimitQueues {
		if limiter.LastUsed.Before(removalThreshold) {
			todelete = append(todelete, ip)
		}
	}
	e.Mutex.RUnlock()

	if len(todelete) == 0 {
		return
	}
	e.Mutex.Lock()
	defer e.Mutex.Unlock()

	removalThreshold = time.Now().Add(-1 * time.Minute * 1)
	for _, ip := range todelete {
		// Validate threshold again to ensure a limiter does not get removed if it was used
		if e.LimitQueues[ip] != nil && e.LimitQueues[ip].LastUsed.Before(removalThreshold) {
			delete(e.LimitQueues, ip)
		}
	}
}
