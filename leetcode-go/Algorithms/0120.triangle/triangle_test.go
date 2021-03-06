package problem0120

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	triangle [][]int
	ans      int
}{
	{[][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}, 11},

	{[][]int{
		{-2},
		{3, 4},
		{6, -5, 7},
		{4, 1, -8, 3},
	}, -12},

	// 可以有多个 testcase
}

func Test_minimumTotal(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, minimumTotal(tc.triangle), "输入:%v", tc)
	}
}

func Benchmark_minimumTotal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minimumTotal(tc.triangle)
		}
	}
}
