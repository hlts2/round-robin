package roundrobin

import (
	"errors"
	"sync/atomic"
)

// ErrServersNotExists is the error that servers dose not exists
var ErrServersNotExists = errors.New("servers dose not exist")

// Servers is custom type of servers
type Servers []string

// RoundRobin returns round-robin closure
func RoundRobin(servers Servers) (func() string, error) {
	if len(servers) == 0 {
		return nil, ErrServersNotExists
	}

	var flg int32

	idx := 0

	var server string

	return func() string {
		for {
			if flg == 0 && atomic.CompareAndSwapInt32(&flg, 0, 1) {
				break
			}
		}

		if idx >= len(servers) {
			idx = 0
		}

		server = servers[idx]

		idx++

		// I do not use defer, decause defer is slow.
		flg = 0
		return server
	}, nil
}
