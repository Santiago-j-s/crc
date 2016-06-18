// Package analysis provides error detection effectiveness analysis for CRC-8 polynomials
package analysis

import (
	"fmt"

	"github.com/Santiago-j-s/crc8"
	"github.com/Santiago-j-s/crc8/server/model/hamming"
)

func crcMap(poly byte) map[int]byte {
	crc := make(map[int]byte)
	tab := crc8.MakeTable(poly)

	for i := 1; i < 256; i++ {
		crc[i] = tab.Crc([]byte{byte(i)})
	}

	return crc
}

func HammingDistance(poly byte) (map[int]int, error) {
	cnt := make(map[int]int)
	cnt[2] = 0
	cnt[3] = 0
	cnt[4] = 0
	cnt[5] = 0
	cnt[6] = 0
	cnt[7] = 0
	cnt[8] = 0
	cnt[9] = 0
	crc := crcMap(poly)
	v1 := []byte{byte(0), byte(0)}
	for n := range crc {
		v2 := []byte{byte(n), crc[n]}
		d, err := hamming.Distance(v1, v2)
		if err != nil {
			e := fmt.Errorf("Error")
			return nil, e
		}
		cnt[d]++
	}
	return cnt, nil
}
