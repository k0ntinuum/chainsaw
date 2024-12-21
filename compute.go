package main

import (
    "crypto/rand"
    "encoding/binary"
	"math/bits"
	//"fmt"
)


func secure_rand() (y uint64) {
	_ = binary.Read(rand.Reader, binary.BigEndian, &y)
	return
}
func secure_bit() (b int) {
	return bits.OnesCount64(secure_rand()) % 2
}
func rand_plain(l int) ( p []int ) {
	for _ = range(l) { p = append(p, secure_bit() )}
	return
}

func codeword(k uint64, b int) ( x uint64 ) {
	for _ = range(200) {
		x = secure_rand()
		//fmt.Printf("k^x = %064b [%d] seek %d \n",k^x, bits.OnesCount64(x) % 2, b)
		if bits.OnesCount64(x^k) % 2 == b { return x }
	}
	print("error\n")
	return uint64(1)
}

func enc(k uint64, p []int) ( c []uint64 ) {
	for _, b := range p {
		x := codeword(k, b)
		c = append(c, x)
		k = k^x
	}
	return
}
func dec(k uint64, c []uint64) ( p []int ) {
	for _, w := range c {
		k = k^w
		p = append(p, bits.OnesCount64(k) % 2)
	}
	return
}