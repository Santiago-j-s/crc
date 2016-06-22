package hamming

import (
	"testing"
)

type weightTab struct {
	in []byte
	out int
}

var tabW = []weightTab {
	{[]byte{0xff}, 8},
	{[]byte{0x00}, 0},
	{[]byte{0xf0}, 4},
    {[]byte{0xff, 0xff}, 16},
    {[]byte{0xaa, 0xaa}, 8},
}

func TestWeight(t *testing.T) {
	for _, test := range tabW {
		result := Weight(test.in...)
		if result != test.out {
			t.Errorf("Input: %x Expected: %x Received: %x\n", test.in, test.out, result)
		}
	}
}

type distanceTab struct {
    in1 []byte
    in2 []byte
    out int   
}

var tabD = []distanceTab {
    {[]byte{0xff}, []byte{0x00}, 8},
    {[]byte{0x00}, []byte{0x00}, 0},
    {[]byte{0x01}, []byte{0x00}, 1},
    {[]byte{0x00, 0x01}, []byte{0xff, 0xff}, 15},
}

func TestDistance(t *testing.T) {
	for _, test := range tabD {
		result, err := Distance(test.in1, test.in2)
        
        if err != nil {
            t.Errorf("Error for Input: %x, %x", test.in1, test.in2)
        }
        
		if result != test.out {
			t.Errorf("Input1: %x Input2: %x Expected: %x Received: %x\n", test.in1, test.in2, test.out, result)
		}
	}
}