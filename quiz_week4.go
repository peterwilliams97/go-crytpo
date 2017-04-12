package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	m1 := "Pay Bob 100$"
	m2 := "Pay Bob 500$"
	b := xor([]byte(m1), []byte(m2))
	m := string(b)
	// s := "20814804c1767293b99f1d9cab3bc3e7" + "ac1e37bfb15599e5f40eef805488281d"
	s := "ac1e37bfb15599e5f40eef805488281d"
	h, _ := hex.DecodeString(s)
	a := xorPad(h, b)
	fmt.Printf("str=%d %q\n", len(s), s)
	fmt.Printf("hash=%d %x\n", len(h), h)
	fmt.Printf("a   =%d %x\n", len(a), a)
	fmt.Printf("m1=%d %q %x\n", len(m1), m1, []byte(m1))
	fmt.Printf("m2=%d %q %x\n", len(m2), m2, []byte(m2))
	fmt.Printf("m =%d %q %x\n", len(m), m, []byte(m))

}

func xorPad(b1, b2 []byte) []byte {
	n1 := len(b1)
	n2 := len(b2)
	b1 = pad(b1, n2)
	b2 = pad(b2, n1)
	return xor(b1, b2)
}

func xor(b1, b2 []byte) []byte {
	n := len(b1)
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = b1[i] ^ b2[i]
	}
	return b
}

func pad(b []byte, n int) []byte {
	o := b
	for i := len(b); i < n; i++ {
		o = append(o, 0)
	}
	return o
}
