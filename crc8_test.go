package crc8

import (
	"testing"
)

type crcData struct {
	in   []byte
	poly byte
	crc  byte
}

var testTable = []crcData{
	{[]byte{0xda}, 0xcc, 0xe4},
	{[]byte{0xda, 0x72}, 0xcc, 0x98},
	{[]byte{0x67, 0x54}, 0x85, 0x00},
	{[]byte{0x67, 0x55}, 0xf3, 0x01},
}

func TestCrc(t *testing.T) {
	for _, test := range testTable {
		tab := MakeTable(test.poly)
		crc := tab.Crc(test.in)
		if crc != byte(test.crc) {
			t.Errorf("Data: %x Poly: %x Expected: %x Received: %x\n", test.in, test.poly, test.crc, crc)
		}
	}
}

var result byte

func benchmarkCrc(s crcData, b *testing.B) {
	tab := MakeTable(s.poly)
	var r byte
	for i := 0; i < b.N; i++ {
		r = tab.Crc(s.in)
	}
	result = r
}

func BenchmarkCrc0(b *testing.B) { benchmarkCrc(testTable[0], b) }
func BenchmarkCrc1(b *testing.B) { benchmarkCrc(testTable[1], b) }
