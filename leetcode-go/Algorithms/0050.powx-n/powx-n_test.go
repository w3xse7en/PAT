package problem0050

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	n   int
	ans int
}{
	{25, 2},
	{9, 1},
	{12258, 5},
	{18580, 2},
	{1212429444, 8},
	// 可以有多个 testcase
}

func Test_myPow(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, translateNum(tc.n), "输入:%v", tc)
	}
}

func Benchmark_myPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			translateNum(tc.n)
		}
	}
}
