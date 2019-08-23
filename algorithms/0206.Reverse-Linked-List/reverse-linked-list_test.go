package problem0206

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
	answer []int
}

var tcs = []tc{

	tc{
		input{[]int{1, 2, 3, 4, 5}},
		output{[]int{5, 4, 3, 2, 1}},
	},
	tc{
		input{[]int{}},
		output{[]int{}},
	},
}

func Test_Problem0206(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		output, input := tc.output, tc.input
		ast.Equal(output.answer, l2s(reverseList(s2l(input.nums))), "Input: %v", input)

	}
}

func Benchmark_Problem0206(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			reverseList(s2l(tc.input.nums))
		}
	}
}
