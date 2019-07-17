package problem0520

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tc struct {
	input
	output
}

type input struct {
	word string
}

type output struct {
	answer bool
}

var tcs = []tc{

	tc{
		input{
			word: "USA",
		},
		output{
			answer: true,
		},
	},

	tc{
		input{
			word: "FlaG",
		},
		output{
			answer: false,
		},
	},
	tc{
		input{
			word: "Taiwan",
		},
		output{
			answer: true,
		},
	},
	tc{
		input{
			word: "CHINA",
		},
		output{
			answer: true,
		},
	},
	tc{
		input{
			word: "HonKong",
		},
		output{
			answer: false,
		},
	},
}

func Test_Problem0520(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		output, input := tc.output, tc.input
		ast.Equal(output.answer, detectCapitalUse(input.word), "Input: %v", input)

	}
}

func Benchmark_Problem0520(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			input := tc.input
			detectCapitalUse(input.word)
		}
	}
}
