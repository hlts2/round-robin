package roundrobin

import (
	"testing"
)

func TestRoundRobin(t *testing.T) {
	tests := []struct {
		servers   Servers
		count     int
		errExists bool
		expected  Servers
	}{
		{
			servers: Servers{
				"server-1",
				"server-2",
				"server-3",
			},
			count:     4,
			errExists: false,
			expected: Servers{
				"server-1",
				"server-2",
				"server-3",
				"server-1",
			},
		},
		{
			servers:   Servers{},
			count:     0,
			errExists: true,
			expected:  Servers{},
		},
	}

	for i, test := range tests {
		next, err := RoundRobin(test.servers)

		errExists := !(err == nil)

		if test.errExists != errExists {
			t.Errorf("tests[%d] - RoundRobin errExists is wrong. expected: %v, got: %v", i, test.errExists, errExists)
		}

		gots := make(Servers, 0, test.count)

		for j := 0; j < test.count; j++ {
			gots = append(gots, next())
		}

		for j, expected := range test.expected {
			got := gots[j]

			if got != expected {
				t.Errorf("tests[%d] - RoundRobin is wrong. expected: %v, got: %v", i, expected, got)
			}
		}
	}
}
