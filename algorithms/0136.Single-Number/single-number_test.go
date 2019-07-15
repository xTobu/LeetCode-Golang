package problem0136

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcs is testcase slice
var tcs = []struct {
	nums []int
	ans  int
}{

	{
		[]int{2, 2, 1},
		1,
	},
	{
		[]int{4, 1, 2, 1, 2},
		4,
	},
}

func Test_Problem0136(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		ast.Equal(tc.ans, singleNumber(tc.nums), "Input: %v", tc)
	}
}

func Benchmark_Problem0136(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			singleNumber(tc.nums)
		}
	}
}
