// Package analysis provides error detection effectiveness analysis for CRC-8 polynomials
package analysis

import (
    "github.com/Santiago-j-s/crc8"
)

func calculate(poly byte) map[int]int {
    m := make(map[int]int)
    tab := crc8.MakeTable(poly)
    
    for i := 0; i < 256; i++ {
        m[i] = tab.Crc([]byte{byte(i)})
    }
    
    return m
}