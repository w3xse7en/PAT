package problem0205

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	s   string
	t   string
	ans bool
}{
	{"aba", "baa", false},
	{"ab", "aa", false},
	{"egg", "add", true},
	{"paper", "title", true},
	{"foo", "bar", false},

	// 可以有多个 testcase
}

func Test_rtos(t *testing.T) {
	fmt.Println(string([]rune{'a', 'b'}))
}
func Test_isIsomorphic(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, isIsomorphic(tc.s, tc.t), "输入:%v", tc)
	}
}

func Benchmark_isIsomorphic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isIsomorphic(tc.s, tc.t)
		}
	}
}
