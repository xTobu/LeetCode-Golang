package problem0657

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tc struct {
	input
	output
}

type input struct {
	moves string
}

type output struct {
	answer bool
}

var tcs = []tc{
	tc{
		input{
			moves: "UD",
		},
		output{
			answer: true,
		},
	},

	tc{
		input{
			moves: "LL",
		},
		output{
			answer: false,
		},
	},
	tc{
		input{
			moves: "UDRL",
		},
		output{
			answer: true,
		},
	},
	tc{
		input{
			moves: "URRL",
		},
		output{
			answer: false,
		},
	},
	tc{
		input{
			moves: "RDL",
		},
		output{
			answer: false,
		},
	},
}

func Test_Problem0657(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		output, input := tc.output, tc.input
		ast.Equal(output.answer, judgeCircle(input.moves), "Input: %v", input)

	}
}

func Benchmark_Problem0657(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			input := tc.input
			judgeCircle(input.moves)
		}
	}
}
