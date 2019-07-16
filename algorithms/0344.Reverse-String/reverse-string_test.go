package problem0344

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tc struct {
	input
	output
}

type input struct {
	s []byte
}

type output struct {
	answer []byte
}

var tcs = []tc{

	tc{
		input{s: []byte{'h', 'e', 'l', 'l', 'o'}},
		output{answer: []byte{'o', 'l', 'l', 'e', 'h'}},
	},

	tc{
		input{s: []byte{'H', 'a', 'n', 'n', 'a', 'h'}},
		output{answer: []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	},
}

func Test_Problem0344(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		output, input := tc.output, tc.input
		reverseString(input.s)
		ast.Equal(output.answer, input.s, "Input: %v", input)

	}
}

func Benchmark_Problem0344(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			input := tc.input
			reverseString(input.s)
		}
	}
}
