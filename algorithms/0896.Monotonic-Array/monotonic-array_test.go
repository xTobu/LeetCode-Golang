package problem0896

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type tc struct {
	input
	output
}

type input struct {
	A []int
}

type output struct {
	answer bool
}

var tcs = []tc{
	tc{
		input{
			A: []int{1, 2, 2, 3},
		},
		output{
			answer: true,
		},
	},
	tc{
		input{
			A: []int{6, 5, 4, 4},
		},
		output{
			answer: true,
		},
	},
	tc{
		input{
			A: []int{1, 3, 2},
		},
		output{
			answer: false,
		},
	},
	tc{
		input{
			A: []int{1, 2, 4, 5},
		},
		output{
			answer: true,
		},
	},
	tc{
		input{
			A: []int{1, 1, 1},
		},
		output{
			answer: true,
		},
	},
}

func Test_Problem0896(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		output, input := tc.output, tc.input
		ast.Equal(output.answer, isMonotonic(input.A), "Input: %v", input)

	}
}

func Benchmark_Problem0896(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			input := tc.input
			isMonotonic(input.A)
		}
	}
}
