package roundrobin

// func TestRoundRobin(t *testing.T) {
// 	tests := []struct {
// 		servers   []string
// 		count     int
// 		errExists bool
// 		expected  []string
// 	}{
// 		{
// 			servers: []string{
// 				"server-1",
// 				"server-2",
// 				"server-3",
// 			},
// 			count:     4,
// 			errExists: false,
// 			expected: []string{
// 				"server-1",
// 				"server-2",
// 				"server-3",
// 				"server-1",
// 			},
// 		},
// 		{
// 			servers:   []string{},
// 			count:     0,
// 			errExists: true,
// 			expected:  []string{},
// 		},
// 	}
//
// 	for i, test := range tests {
// 		rr, err := New(test.servers)
//
// 		errExists := !(err == nil)
//
// 		if test.errExists != errExists {
// 			t.Errorf("tests[%d] - RoundRobin errExists is wrong. expected: %v, got: %v", i, test.errExists, errExists)
// 		}
//
// 		gots := make([]string, 0, test.count)
//
// 		for j := 0; j < test.count; j++ {
// 			gots = append(gots, rr.Next())
// 		}
//
// 		for j, expected := range test.expected {
// 			got := gots[j]
//
// 			if got != expected {
// 				t.Errorf("tests[%d] - RoundRobin is wrong. expected: %v, got: %v", i, expected, got)
// 			}
// 		}
// 	}
// }
