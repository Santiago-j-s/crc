package main

import (
	"fmt"

	"github.com/Santiago-j-s/crc8"
)

func main() {
	poly := byte(0xfd)
	tab := crc8.MakeTable(poly)

	fmt.Printf("poly: %08b\n", poly)
	fmt.Println("\tmsg -> crc")
	for msg := 0; msg < 256; msg++ {
		crc := tab.Sum([]byte{byte(msg)})
		fmt.Printf("%3d %08b %08b\n", msg, msg, crc)
	}
}
