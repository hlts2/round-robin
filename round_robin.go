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
	addrs []string
	lf    lockfree.LockFree
	idx   int
}

// New returns RoundRobin(*roundrobin) object
func New(addrs []string) (RoundRobin, error) {
	if len(addrs) == 0 {
		return nil, ErrServersNotExists
	}

	return &roundrobin{
		addrs: addrs,
		lf:    lockfree.New(),
		idx:   0,
	}, nil
}

// Next returns next address
func (r *roundrobin) Next() string {
	r.lf.Wait()

	if r.idx >= len(r.addrs) {
		r.idx = 0
	}

	address := r.addrs[r.idx]
	r.idx++

	r.lf.Signal()

	return address
}
