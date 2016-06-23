// Copyright 2016
// Dibez Pablo pdibez@gmail.com
// Santana Santiago santana.santiago@gmail.com
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package analysis provides error detection effectiveness analysis for CRC-8 polynomials
package analysis

import (
	"github.com/Santiago-j-s/crc8"
	"github.com/Santiago-j-s/crc8/analysis/hamming"
)

func crcMap(poly byte) map[int]byte {
	crc := make(map[int]byte)
	tab := crc8.MakeTable(poly)

	for i := 1; i < 256; i++ {
		crc[i] = tab.Sum([]byte{byte(i)})
	}

	return crc
}

func HammingDistance(poly byte) map[int]int {
	cnt := make(map[int]int)
	crc := crcMap(poly)

	v1 := []byte{byte(0), byte(0)}
	for n := range crc {
		v2 := []byte{byte(n), crc[n]}
		d, err := hamming.Distance(v1, v2)
		if err != nil {
			panic("the lengths of v1 and v2 must be equal")
		}
		cnt[d]++
	}
	return cnt
}
