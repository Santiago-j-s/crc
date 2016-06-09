// Copyright 2016, Dibez Pablo, Santana Santiago. 
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package crc8 implements the 8-bit cyclic redundancy check,
// with some fixed parameters of the Rocksoft Model:
//    Width  : 8
//    Init   : 0000
//    RefIn  : False (MSB)
//    RefOut : False
//    XorOut : 0000
// 
// See http://chrisballance.com/wp-content/uploads/2015/10/CRC-Primer.html for more information.
package crc8

// Size of a CRC-8 checksum in bytes.
const Size = 2;

// Table is a 256-word table representing the polynomial for efficient processing.
type Table [256]byte

// MakeTable returns a Table generated from the specified poly.
func MakeTable(poly byte) *Table {
	t := new(Table)
	for i := 0; i < 256; i++ {
		crc := byte(i)
		for j := 0; j < 8; j++ {
			if (crc & 0x80) != 0 {
				crc = (crc << 1) ^ poly
			} else {
				crc <<= 1
			}
		}
		t[i] = crc
	}
	return t
}

// Sum performs the checksum of data using the polynomial defined by the Table
func Sum(data []byte, tab *Table) byte {
	crc := byte(0)
	for _, v := range data {
		crc = tab[byte(crc) ^ v]
	}
	return crc
}