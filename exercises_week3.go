package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func main() {
	text := "hello world\n"
	b := []byte(text)

	h := sha256.New()
	h.Write(b)
	fmt.Printf("[% x]\n", h.Sum(nil))

	sum := sha256.Sum256(b)
	fmt.Printf("[% x]\n", sum)

	blockSize := 1024
	p1 := "6.2.birthday.mp4"
	p2 := "6.1.intro.mp4"
	_, h1 := fileHashes(p1, blockSize)
	fmt.Printf("path=%20q hash=%x\n", p1, h1[0])
	_, h2 := fileHashes(p2, blockSize)
	fmt.Printf("path=%20q hash=%x\n", p2, h2[0])
}

func fileHashes(path string, blockSize int) ([][]byte, [][32]byte) {
	blocks := fileBlocks(path, blockSize)
	return blockHashes(blocks)
}

func blockHashes(blocks [][]byte) ([][]byte, [][32]byte) {
	n := len(blocks)
	l := 0
	for i := n - 1; i >= 0; i-- {
		b := blocks[i]
		if len(b) == 0 {
			panic("empty block")
		}
		l += len(b)
	}
	fmt.Printf("n=%d l=%d size=%d\n", n, l, (l+n-1)/n)

	hashes := make([][32]byte, n)
	hashes[n-1] = sha256.Sum256(blocks[n-1])
	fmt.Printf("%5d: %x\n", n-1, hashes[n-1])
	for i := n - 2; i >= 0; i-- {
		a := len(blocks[i])
		blocks[i] = append(blocks[i], hashes[i+1][:]...)
		if len(blocks[i]) != a+32 {
			panic("bad append")
		}
		for j := 0; j < 32; j++ {
			if blocks[i][a+j] != hashes[i+1][j] {
				panic("goooo")
			}
		}
		hashes[i] = sha256.Sum256(blocks[i])
		if i > n-10 {
			fmt.Printf("%5d: %x\n", i, hashes[i])
		}
	}
	return blocks, hashes
}

func fileBlocks(path string, blockSize int) [][]byte {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	blocks := [][]byte{}
	buf := make([]byte, blockSize)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		blocks = append(blocks, buf[:n])
	}
	return blocks
}
