package problem0283

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
	nums []int
}

var tcs = []tc{

	tc{
		input{nums: []int{0, 1, 0, 3, 12}},
		output{nums: []int{1, 3, 12, 0, 0}},
	},
}

func Test_Problem0283(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		output, input := tc.output, tc.input
		moveZeroes(input.nums)
		ast.Equal(output.nums, input.nums, "Input: %v", input)

	}
}

func Test_Problem0283_2(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		output, input := tc.output, tc.input
		moveZeroes2(input.nums)
		ast.Equal(output.nums, input.nums, "Input: %v", input)

	}
}

func Benchmark_Problem0283(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			input := tc.input
			moveZeroes(input.nums)
		}
	}
}

func Benchmark_Problem0283_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			input := tc.input
			moveZeroes2(input.nums)
		}
	}
}
