package main

import "fmt"

type bit byte // no bits in golang, use bytes but pretend they are only bits

const bitLength = 8

func getBit(num uint8, index int) bit {
	return bit((num >> index) & 1)
}
func Transpose(input []uint8) [][]bit {
	originalLength := len(input)
	ret := make([][]bit, bitLength)
	for i := 0; i < bitLength; i++ {
		ret[i] = make([]bit, originalLength)
		for j := 0; j < originalLength; j++ {
			ret[i][j] = getBit(input[j], bitLength-i-1)
		}
	}
	return ret
}

func main() {
	test := []uint8{1, 2}
	fmt.Println(test)
	fmt.Println(Transpose(test))
}
