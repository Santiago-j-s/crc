package analysis

import (
	"testing"
    "fmt"
)

type data struct {
	in   byte
	out  map[int]int
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
		m, err := HammingDistance(test.in)
        if err != nil {
            t.Errorf("Error")
        }
        
        b := (test.out[2] == m[2])
        b = b &&(test.out[3] == m[3])
        b = b &&(test.out[4] == m[4])
        b = b &&(test.out[5] == m[5])
        b = b &&(test.out[6] == m[6])
        b = b &&(test.out[7] == m[7])
        b = b &&(test.out[8] == m[8])
        
        if !b {
            s := fmt.Sprintf("Expected: %v\n", test.out)
            s += fmt.Sprintf("Received: %v", m)
            t.Errorf(s)
        }
	}
}