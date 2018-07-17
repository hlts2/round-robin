package roundrobin

import (
	"errors"
	"sync"
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

	mu := new(sync.Mutex)

	idx := 0

	var server string

	return func() string {
		defer mu.Unlock()
		mu.Lock()

		if idx >= len(servers) {
			idx = 0
		}

		server = servers[idx]

		idx++

		return server
	}, nil
}
