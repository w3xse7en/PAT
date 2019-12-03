package problem0118

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	numRows int
	ans     [][]int
}{
	{5, [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}},

	// 可以有多个 testcase
}

func Test_generate(t *testing.T) {
	a := assert.New(t)

	for _, tc := range tcs {
		a.Equal(tc.ans, generate(tc.numRows), "输入:%v", tc)
	}
}

func Benchmark_generate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			generate(tc.numRows)
		}
	}
}
