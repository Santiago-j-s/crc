package main

import (
	"fmt"

	"github.com/Santiago-j-s/crc8"
)

func main() {
	poly := byte(0xfd)
	check := byte(250) // msg. original
	
	tab := crc8.MakeTable(poly)
	crcCheck := crc8.Sum([]byte{check}, tab) // crc del msg. original
	
	fmt.Printf("Msg: %8b (%d)\n", check, check)
	fmt.Printf("CRC: %8b (%d)\n", crcCheck, crcCheck)
	
	cnt := 0
	fmt.Println("msg + checksum -> crc")
	for msg := 0; msg < 256; msg++ {
		for crc := 0; crc < 256; crc++ {
			sum := crc8.Sum([]byte{byte(msg), byte(crc)}, tab)
			if sum == 0 {
				cnt++
				fmt.Printf("%3d %08b %08b -> %08b\n", cnt, msg, crc, sum)
			}
		}
	}
	fmt.Println()
	fmt.Printf("%d falsos positivos de %d mensajes.\n", cnt-1, 65536)
}
