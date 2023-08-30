package main

import (
	"fmt"
	"testing"

	"slices"
)

func TestTranspose(t *testing.T) {
	tests := []struct {
		input  []uint8
		output [][]bit
	}{
		{
			input:  []uint8{},
			output: [][]bit{{}, {}, {}, {}, {}, {}, {}, {}},
		},
		{
			input:  []uint8{0},
			output: [][]bit{{0}, {0}, {0}, {0}, {0}, {0}, {0}, {0}},
		},
		{
			input:  []uint8{1},
			output: [][]bit{{0}, {0}, {0}, {0}, {0}, {0}, {0}, {1}},
		},
		{
			input:  []uint8{2},
			output: [][]bit{{0}, {0}, {0}, {0}, {0}, {0}, {1}, {0}},
		},
		{
			input:  []uint8{127},
			output: [][]bit{{0}, {1}, {1}, {1}, {1}, {1}, {1}, {1}},
		},
		{
			input:  []uint8{255},
			output: [][]bit{{1}, {1}, {1}, {1}, {1}, {1}, {1}, {1}},
		},
		{
			input:  []uint8{0, 0},
			output: [][]bit{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}},
		},
		{
			input:  []uint8{0, 1},
			output: [][]bit{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 1}},
		},
		{
			input:  []uint8{255, 0},
			output: [][]bit{{1, 0}, {1, 0}, {1, 0}, {1, 0}, {1, 0}, {1, 0}, {1, 0}, {1, 0}},
		},
	}
	for _, tc := range tests {
		got := Transpose(tc.input)
		if !slices.EqualFunc(tc.output, got, func(s1, s2 []bit) bool {
			return slices.Equal(s1, s2)
		}) {

			//fmt.P intln(cap(tc.output[0]))
			fmt.Println(cap(got[0]))
			t.Fatalf("test %v: expected %v, got %v", tc.input, tc.output, got)
		}
	}
}

func TestGetBit(t *testing.T) {
	tests := []struct {
		num    uint8
		index  int
		output bit
	}{
		{
			num:    0,
			index:  0,
			output: 0,
		},
		{
			num:    1,
			index:  0,
			output: 1,
		},
		{
			num:    2,
			index:  0,
			output: 0,
		},
		{
			num:    2,
			index:  1,
			output: 1,
		},
		{
			num:    255,
			index:  7,
			output: 1,
		},
		{
			num:    127,
			index:  7,
			output: 0,
		},
		{
			num:    127,
			index:  6,
			output: 1,
		},
	}
	for _, tc := range tests {
		got := getBit(tc.num, tc.index)
		if got != tc.output {
			t.Fatalf("test %dth bit of %d: expected %v, got %v", tc.index, tc.num, tc.output, got)
		}
	}
}
