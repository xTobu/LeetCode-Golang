package problem0136

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tc struct {
	input
	output
}

type input struct {
	nums []int
}

type output struct {
	target int
}

var tcs = []tc{

	tc{
		input{[]int{2, 2, 1}},
		output{1},
	},

	tc{
		input{[]int{4, 1, 2, 1, 2}},
		output{4},
	},
}

// tcs is testcase slice
// var tcs = []struct {
// 	nums []int
// 	ans  int
// }{

// 	{
// 		[]int{2, 2, 1},
// 		1,
// 	},
// 	{
// 		[]int{4, 1, 2, 1, 2},
// 		4,
// 	},
// }

func Test_Problem0136(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs {
		output, input := tc.output, tc.input
		ast.Equal(output.target, singleNumber(input.nums), "Input: %v", input)
	}
}

func Benchmark_Problem0136(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			singleNumber(tc.input.nums)
		}
	}
}
