package problem0076

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	t   string
	ans string
}{
	{"a", "a", "a"},
	{"ADOBECODEBANC", "ABC", "BANC"},
	{"aa", "a", "a"},
	{"aa", "aaa", ""},

	// 可以有多个 testcase
}

func Test_minWindow(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, minWindow(tc.s, tc.t), "输入:%v", tc)
	}
}

func Benchmark_minWindow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minWindow(tc.s, tc.t)
		}
	}
}
