package ratelimit

import (
	"fmt"
	"strings"
)

func getAddrWithoutPort(addr string) string {
	addrPortions := strings.Split(addr, ":")
	port := addrPortions[len(addrPortions)-1]
	suffix := fmt.Sprintf(":%s", port)
	return strings.TrimSuffix(addr, suffix)
}
