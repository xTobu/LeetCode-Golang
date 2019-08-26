package problem0169

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
	answer int
}

var tcs = []tc{
	tc{
		input{
			nums: []int{3, 2, 3},
		},
		output{
			answer: 3,
		},
	},
	tc{
		input{
			nums: []int{2, 2, 1, 1, 1, 2, 2},
		},
		output{
			answer: 2,
		},
	},
}

func Test_Problem0206(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		output, input := tc.output, tc.input
		ast.Equal(output.answer, majorityElement(input.nums), "Input: %v", input)

	}
}
func Test_Problem0206_2(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		output, input := tc.output, tc.input
		ast.Equal(output.answer, majorityElement2(input.nums), "Input: %v", input)

	}
}

func Benchmark_Problem0206_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			majorityElement2(tc.input.nums)
		}
	}
}
