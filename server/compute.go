// Copyright 2016
// Dibez Pablo pdibez@gmail.com
// Santana Santiago santana.santiago@gmail.com
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package main

import (
	"fmt"

	"github.com/Santiago-j-s/crc8"
	"github.com/Santiago-j-s/crc8/analysis"
	"github.com/Santiago-j-s/stringutil"
)

// verifyLen returns an error if s is not of length l
func verifyLen(s string, l int) error {
	if len(s) != l {
		return fmt.Errorf("Only strings of length %d are allowed", l)
	}
	return nil
}

func verifyLen8(s string) error {
	return verifyLen(s, 8)
}

// verifyBinaryString returns an error if s is not composed only of characters '0' and '1'
func verifyBinaryString(s string) error {
	for _, c := range s {
		if c != '0' && c != '1' {
			return fmt.Errorf("%v isn't a binary number", s)
		}
	}
	return nil
}

func verifyBinaryStrings(str ...string) error {
	for _, s := range str {
		if err := verifyBinaryString(s); err != nil {
			return err
		}
	}
	return nil
}

func readByte(s string) (c byte, err error) {
	if len(s)%8 != 0 {
		return 0, fmt.Errorf("InputError. Length must be multiple of 8.")
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

func crc(poly string, msg string) (string, error) {
	if err := verifyLen(poly, 8); err != nil {
		return "", err
	}
	
	if err := verifyBinaryStrings(poly, msg); err != nil {
		return "", err
	}
	
	bPoly, err := readByte(poly)
	if err != nil {
		return "", err
	}

	bMsg, err := readByte(msg)
	if err != nil {
		return "", err
	}

	tab := crc8.MakeTable(bPoly)
	chk := tab.Sum([]byte{bMsg})
	res := fmt.Sprintf("%08b", chk)
	return res, nil
}

func hamming(poly string) (string, error) {
	if err := verifyLen(poly, 8); err != nil {
		return "", err
	}
	
	if err := verifyBinaryString(poly); err != nil {
		return "", err
	}
	
	bPoly, err := readByte(poly)

	if err != nil {
		return "", err
	}

	m := analysis.HammingDistance(bPoly)
	var s string

	for key := 2; key <= 10; key++ {
		s += fmt.Sprintf("%d errores de %d bits.\n", m[key], key)
	}

	return s, nil
}
