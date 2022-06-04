package bittorrent

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasPiece(t *testing.T) {
	bf := Bitfield{0b01010100, 0b01010100}
	outputs := []bool{false, true, false, true, false, true, false, false, false, true, false, true, false, true, false, false, false, false, false, false}
	for i := 0; i < len(outputs); i++ {
		assert.Equal(t, outputs[i], bf.HasPiece(i))
	}
}

func TestSetPiece(t *testing.T) {
	tests := []struct {
		input  Bitfield
		index  int
		output Bitfield
	}{
		{
			input:  Bitfield{0b00000000},
			index:  7,
			output: Bitfield{0b00000001},
		},
		{
			input:  Bitfield{0b00001001},
			index:  6,
			output: Bitfield{0b00001011},
		},
		{
			input:  Bitfield{0b01010100, 0b01010100},
			index:  4,
			output: Bitfield{0b01011100, 0b01010100},
		},
		{
			input:  Bitfield{0b01010100, 0b01010100},
			index:  9,
			output: Bitfield{0b01010100, 0b01010100},
		},
		{
			input:  Bitfield{0b01010100, 0b01010100},
			index:  15,
			output: Bitfield{0b01010100, 0b01010101},
		},
		{
			input:  Bitfield{0b01010100, 0b01010100},
			index:  19,
			output: Bitfield{0b01010100, 0b01010100},
		},
	}
	for _, test := range tests {
		bf := test.input
		bf.SetPiece(test.index)
		assert.Equal(t, test.output, bf)
	}
}
