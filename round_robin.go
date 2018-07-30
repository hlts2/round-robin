package roundrobin

import (
	"errors"

	"github.com/hlts2/lock-free"
)

// ErrServersNotExists is the error that servers dose not exists
var ErrServersNotExists = errors.New("servers dose not exist")

// RoundRobin represents base round-robin interface
type RoundRobin interface {
	Next() string
}

type roundrobin struct {
	addresses []string
	lf        lockfree.LockFree
	idx       int
}

// New returns RoundRobin(*roundrobin) object
func New(addresses []string) (RoundRobin, error) {
	if len(addresses) == 0 {
		return nil, ErrServersNotExists
	}

	return &roundrobin{
		addresses: addresses,
		lf:        lockfree.New(),
		idx:       0,
	}, nil
}

// RoundRobin returns round-robin closure
func (r *roundrobin) Next() string {
	r.lf.Wait()

	if r.idx >= len(r.addresses) {
		r.idx = 0
	}

	address := r.addresses[r.idx]
	r.idx++

	r.lf.Signal()

	return address
}
