package roundrobin

import (
	"errors"
	"sync/atomic"
)

// ErrServersNotExists is the error that servers dose not exists
var ErrServersNotExists = errors.New("servers dose not exist")

// RoundRobin is an interface for representing round-robin balancing.
type RoundRobin[T any] interface {
	Next() T
}

type roundrobin[T any] struct {
	urls []T
	next uint64
}

// New RoundRobin instance.
func New[T any](urls ...T) (RoundRobin[T], error) {
	if len(urls) == 0 {
		return nil, ErrServersNotExists
	}

	return &roundrobin[T]{
		urls: urls,
	}, nil
}

// Next item in list.
func (r *roundrobin[T]) Next() T {
	n := atomic.AddUint64(&r.next, 1)
	return r.urls[(int(n)-1)%len(r.urls)]
}
