package problem0771

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tc struct {
	input
	output
}

type input struct {
	J string
	S string
}

type output struct {
	answer int
}

var tcs = []tc{
	tc{
		input{
			J: "aA",
			S: "aAAbbbb",
		},
		output{
			answer: 3,
		},
	},

	tc{
		input{
			J: "z",
			S: "ZZ",
		},
		output{
			answer: 0,
		},
	},
}

func Test_Problem0771(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		output, input := tc.output, tc.input
		ast.Equal(output.answer, numJewelsInStones(input.J, input.S), "Input: %v", input)

	}
}

func Benchmark_Problem0771(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			input := tc.input
			numJewelsInStones(input.J, input.S)
		}
	}
}
