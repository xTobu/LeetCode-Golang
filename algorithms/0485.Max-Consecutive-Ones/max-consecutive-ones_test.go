package problem0485

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
		input{[]int{1, 1, 0, 1, 1, 1}},
		output{3},
	},

	tc{
		input{[]int{0, 1, 0, 0, 0}},
		output{1},
	},
}

func Test_Problem0485(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		output, input := tc.output, tc.input
		ast.Equal(output.target, findMaxConsecutiveOnes(input.nums), "Input: %v", input)
	}
}

func Benchmark_Problem0485(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			input := tc.input
			findMaxConsecutiveOnes(input.nums)
		}
	}
}
