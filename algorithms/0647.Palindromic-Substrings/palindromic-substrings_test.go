package problem0647

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tc struct {
	input
	output
}

type input struct {
	s string
}

type output struct {
	answer int
}

var tcs = []tc{
	tc{
		input{
			s: "abc",
		},
		output{
			answer: 3,
		},
	},
	tc{
		input{
			s: "aaa",
		},
		output{
			answer: 6,
		},
	},
}

func Test_Problem0647(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		output, input := tc.output, tc.input
		ast.Equal(output.answer, countSubstrings(input.s), "Input: %v", input)

	}
}

func Benchmark_Problem0647(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			input := tc.input
			countSubstrings(input.s)
		}
	}
}
