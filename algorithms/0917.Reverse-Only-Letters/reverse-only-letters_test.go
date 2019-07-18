package problem0917

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tc struct {
	input
	output
}

type input struct {
	S string
}

type output struct {
	answer string
}

var tcs = []tc{
	tc{
		input{
			S: "ab-cd",
		},
		output{
			answer: "dc-ba",
		},
	},
	tc{
		input{
			S: "a-bC-dEf-ghIj",
		},
		output{
			answer: "j-Ih-gfE-dCba",
		},
	},
	tc{
		input{
			S: "Test1ng-Leet=code-Q!",
		},
		output{
			answer: "Qedo1ct-eeLg=ntse-T!",
		},
	},
}

func Test_Problem0917(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		output, input := tc.output, tc.input
		ast.Equal(output.answer, reverseOnlyLetters(input.S), "Input: %v", input)

	}
}

func Benchmark_Problem0917(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			input := tc.input
			reverseOnlyLetters(input.S)
		}
	}
}
