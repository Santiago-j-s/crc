package analysis

import (
	"fmt"
	"testing"
)

type data struct {
	in  byte
	out map[int]int
}

var testTable = []data{
	{0xd7, map[int]int{
		2: 0,
		3: 0,
		4: 0,
		5: 24,
		6: 44,
		7: 40, // diferente a Koopman
		8: 45,
		9: 40,
	}},
	{0xd5, map[int]int{
		2: 0,
		3: 0,
		4: 12,
		5: 0,
		6: 69,
		7: 0,
		8: 89,
	}},
}

func TestHammingDistance(t *testing.T) {
	for _, test := range testTable {
		m := HammingDistance(test.in)

		b := true
		for key, value := range test.out {
			b = b && (value == m[key])
		}

		if !b {
			s := fmt.Sprintf("\nExpected: %v\n", test.out)
			s += fmt.Sprintf("Received: %v", m)
			t.Errorf(s)
		}
	}
}
