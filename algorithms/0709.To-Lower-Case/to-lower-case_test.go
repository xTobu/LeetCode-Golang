package problem0709

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tc struct {
	input
	output
}

type input struct {
	str string
}

type output struct {
	answer string
}

var tcs = []tc{
	tc{
		input{
			str: "Hello",
		},
		output{
			answer: "hello",
		},
	},

	tc{
		input{
			str: "here",
		},
		output{
			answer: "here",
		},
	},
	tc{
		input{
			str: "LOVELY",
		},
		output{
			answer: "lovely",
		},
	},
}

func Test_Problem0709(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		output, input := tc.output, tc.input
		ast.Equal(output.answer, toLowerCase(input.str), "Input: %v", input)

	}
}

func Benchmark_Problem0709(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			input := tc.input
			toLowerCase(input.str)
		}
	}
}
