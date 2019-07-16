package problem0242

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
	t string
}

type output struct {
	answer bool
}

var tcs = []tc{

	tc{
		input{"anagram", "nagaram"},
		output{true},
	},

	tc{
		input{"rat", "car"},
		output{false},
	},
}

func Test_Problem0242(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		output, input := tc.output, tc.input
		ast.Equal(output.answer, isAnagram(input.s, input.t), "Input: %v", input)

	}
}

func Benchmark_Problem0242(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			input := tc.input
			isAnagram(input.s, input.t)
		}
	}
}
