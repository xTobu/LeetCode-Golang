package problem0189

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type input struct {
	nums []int
	k    int
}

type output struct {
	nums []int
}

type tc struct {
	input  input
	output output
}

var tcs = []tc{
	tc{
		input: input{
			nums: []int{1, 2, 3, 4, 5, 6},
			k:    2,
		},
		output: output{
			nums: []int{5, 6, 1, 2, 3, 4},
		},
	},
	tc{
		input: input{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    14,
		},
		output: output{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
		},
	},
}

func Test_Problem0189(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		output, input := tc.output, tc.input
		rotate(input.nums, input.k)
		ast.Equal(output.nums, input.nums, "Input :%v", input)
	}
}

func Benchmark_Problem0189(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			input := tc.input
			rotate(input.nums, input.k)
		}
	}
}
