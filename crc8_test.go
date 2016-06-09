package crc8

import (
	"testing"
)

type sum struct {
	in []byte
	poly byte
	out byte
}

var tableSum = []sum {
	{[]byte{0xda}, 0xcc, 0xe4},
	{[]byte{0xda, 0x72}, 0xcc, 0x98},
	{[]byte{0x67, 0x54}, 0x85, 0x00},
	{[]byte{0x67, 0x55}, 0xf3, 0x01},
}

func TestSum(t *testing.T) {
	for _, test := range tableSum {
		s := Sum(test.in, MakeTable(test.poly))
		if s != byte(test.out) {
			t.Errorf("Data: %x Poly: %x Expected: %x Received: %x\n", test.in, test.poly, test.out, s)
		}
	}
}

var result byte;

func benchmarkSum(s sum, b *testing.B) {
	tab := MakeTable(s.poly)
	var r byte;
	for i := 0; i < b.N; i++ {
		r = Sum(s.in, tab)
	}
	result = r
}

func BenchmarkSum0(b *testing.B) { benchmarkSum(tableSum[0], b) }
func BenchmarkSum1(b *testing.B) { benchmarkSum(tableSum[1], b) }