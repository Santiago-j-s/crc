// Package analysis provides error detection effectiveness analysis for CRC-8 polynomials
package analysis

import (
	"github.com/Santiago-j-s/crc8"
	"github.com/Santiago-j-s/crc8/server/model/hamming"
)

func crcMap(poly byte) map[int]byte {
	crc := make(map[int]byte)
	tab := crc8.MakeTable(poly)

	for i := 1; i < 256; i++ {
		crc[i] = tab.Sum([]byte{byte(i)})
	}

	return crc
}

func HammingDistance(poly byte) (map[int]int, error) {
	cnt := make(map[int]int)
	crc := crcMap(poly)

	v1 := []byte{byte(0), byte(0)}
	for n := range crc {
		v2 := []byte{byte(n), crc[n]}
		d, err := hamming.Distance(v1, v2)
		if err != nil {
			return nil, err
		}
		cnt[d]++
	}
	return cnt, nil
}
