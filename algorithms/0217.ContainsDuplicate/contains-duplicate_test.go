package problem0217

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
	answer bool
}

var tcs = []tc{

	tc{
		input{[]int{1, 2, 3, 1}},
		output{true},
	},

	tc{
		input{[]int{1, 2, 3, 4}},
		output{false},
	},

	tc{
		input{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
		output{true},
	},
}

func Test_Problem0217(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		output, input := tc.output, tc.input
		ast.Equal(output.answer, containsDuplicate(input.nums), "Input: %v", input)

	}
}
func Test_Problem0217_2(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		output, input := tc.output, tc.input
		ast.Equal(output.answer, containsDuplicate2(input.nums), "Input: %v", input)

	}
}

func Benchmark_Problem0217(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			containsDuplicate(tc.input.nums)
		}
	}
}

func Benchmark_Problem0217_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			containsDuplicate2(tc.input.nums)
		}
	}
}
