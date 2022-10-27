package roundrobin

import (
	"errors"
	"net/url"
	"sync/atomic"
)

// ErrServersNotExists is the error that servers dose not exists
var ErrServersNotExists = errors.New("servers dose not exist")

// RoundRobin is an interface for representing round-robin balancing.
type RoundRobin interface {
	Next() *url.URL
}

type roundrobin struct {
	urls []*url.URL
	next uint32
}

// New returns RoundRobin implementation(*roundrobin).
func New(urls ...*url.URL) (RoundRobin, error) {
	if len(urls) == 0 {
		return nil, ErrServersNotExists
	}

	return &roundrobin{
		urls: urls,
	}, nil
}

// Next returns next address
func (r *roundrobin) Next() *url.URL {
	n := atomic.AddUint32(&r.next, 1)
	return r.urls[(int(n)%len(r.urls)]
}
