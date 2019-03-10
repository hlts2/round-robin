package roundrobin

import (
	"net/url"
	"reflect"
	"testing"
)

func TestRoundRobin(t *testing.T) {
	tests := []struct {
		urls     []*url.URL
		iserr    bool
		expected []string
		want     []*url.URL
	}{
		{
			urls: []*url.URL{
				{Host: "192.168.33.10"},
				{Host: "192.168.33.11"},
				{Host: "192.168.33.12"},
			},
			iserr: false,
			want: []*url.URL{
				{Host: "192.168.33.10"},
				{Host: "192.168.33.11"},
				{Host: "192.168.33.12"},
				{Host: "192.168.33.10"},
			},
		},
		{
			urls:  []*url.URL{},
			iserr: true,
			want:  []*url.URL{},
		},
	}

	for i, test := range tests {
		rr, err := New(test.urls)

		if got, want := !(err == nil), test.iserr; got != want {
			t.Errorf("tests[%d] - RoundRobin iserr is wrong. want: %v, but got: %v", i, test.want, got)
		}

		gots := make([]*url.URL, 0, len(test.want))
		for j := 0; j < len(test.want); j++ {
			gots = append(gots, rr.Next())
		}

		if got, want := gots, test.want; !reflect.DeepEqual(got, want) {
			t.Errorf("tests[%d] - RoundRobin is wrong. want: %v, got: %v", i, want, got)
		}
	}
}
