// Copyright 2016
// Dibez Pablo pdibez@gmail.com
// Santana Santiago santana.santiago@gmail.com
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hamming provides functions for calculate hamming weight and hamming distance of binary numbers
package hamming

import "fmt"

// Weight return the number of ones in a sequence of bits
func Weight(msg ...byte) int {
	cnt := 0
	for _, m := range msg {
		cnt += onesInByte(m)
	}
	return cnt
}

func onesInByte(msg byte) int {
	msg2 := msg
	cnt := 0
	for i := 0; i < 8; i++ {
		lastbit := msg2 & 0x01
		if lastbit == 1 {
			cnt++
		}
		msg2 >>= 1
	}
	return cnt
}

// Distance returns the Hamming Distance between m1 and m2.
// The Hamming Distance between two binary numbers is the number of
// bits that aren't equal in them
func Distance(m1, m2 []byte) (int, error) {
    l1 := len(m1)
    l2 := len(m2)
    
	if l1 != l2 {
		return 0, fmt.Errorf("the lengths of m1 and m2 must be equal")
	}
    
    dif := make([]byte, l1)
    for i := range dif {
        dif[i] = m2[i] ^ m1[i]
    }
    hw := Weight(dif...)
    
    return hw, nil
}