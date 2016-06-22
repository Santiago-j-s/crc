package main

import (
	"fmt"

	"github.com/Santiago-j-s/crc8"
	"github.com/Santiago-j-s/stringutil"
)

func readByte(s string) (c byte, err error) {
	if len(s)%8 != 0 {
		return 0, fmt.Errorf("InputError")
	}

	var b byte
	reversed := stringutil.Reverse(s)

	for i, ch := range reversed {
		if ch == '1' {
			d := byte(0x01) << uint(i)
			b = b | d
		}
	}

	return b, nil
}

func crc(poly string, msg string) string {
	bPoly, err := readByte(poly)

	if err != nil {
		panic("InputError")
	}

	bMsg, err := readByte(msg)

	if err != nil {
		panic("InputError")
	}

	tab := crc8.MakeTable(bPoly)
	chk := tab.Sum([]byte{bMsg})
	res := fmt.Sprintf("%08b", chk)
	return res
}
