package conversion

import (
	"github.com/KaymeKaydex/algorithms-console-io.git/equals"
	"testing"

)

func TestSliceAtoi(t *testing.T) {
	cases := []struct {
		input       []string
		isException bool
		output      []int
	}{
		{
			[]string{"213", "321"},
			false,
			[]int{213,321},
		},
		{
			[]string{"5345", "231", "2135"},
			false,
			[]int{5345,231,2135},
		},
		{
			[]string{"-213", "-9321"},
			false,
			[]int{-213,-9321},

		},
	}

	for _, c := range cases {
		result, _ := SliceAtoi(c.input)
		if !equals.IntArrayEquals(result, c.output) {
			t.Errorf("Bad result by int conversation. Expected : %d. Result : %d", c.output, result )
		}
	}
}
