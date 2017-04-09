package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {

	blockSize := 1024
	p1 := "6.2.birthday.mp4"
	p2 := "6.1.intro.mp4"
	h10, _ := hex.DecodeString("03c08f4ee0b576fe319338139c045c89c3e8e9409633bea29442e21425006ea8")
	_, h1 := fileHashes(p1, blockSize)
	fmt.Printf("path=%20q hash=%x\n", p1, h1[0])
	fmt.Printf("path=%20q hash=%x\n", "", h10)
	_, h2 := fileHashes(p2, blockSize)
	fmt.Printf("path=%20q hash=%x\n", p2, h2[0])

	if !isEqual(h1[0][:], h10) {
		panic("bad hash")
	}
}

func fileHashes(path string, blockSize int) ([][]byte, [][32]byte) {
	blocks := fileBlocks(path, blockSize)
	return blockHashes(blocks)
}

func fileBlocks(path string, blockSize int) [][]byte {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Printf("%20q %8d\n", path, fileSize(*f))
	blocks := [][]byte{}
	for {
		// !@#$ Why does the buffer need to be created each pass through the loop
		buf := make([]byte, blockSize)
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

func blockHashes(blocks [][]byte) ([][]byte, [][32]byte) {
	n := len(blocks)
	bout := make([][]byte, n)
	hashes := make([][32]byte, n)
	bout[n-1] = blocks[n-1]
	hashes[n-1] = sha256.Sum256(blocks[n-1])
	fmt.Printf("%5d: %x\n", n-1, hashes[n-1])
	for i := n - 2; i >= 0; i-- {
		bout[i] = append(blocks[i], hashes[i+1][:]...)
		hashes[i] = sha256.Sum256(bout[i])
	}
	return bout, hashes
}

func printBlockStats(blocks [][]byte) {
	n := len(blocks)
	l := 0
	for i := n - 1; i >= 0; i-- {
		b := blocks[i]
		if len(b) == 0 {
			panic("empty block")
		}

		if i < n-1 {
			if isEqual(b, blocks[i+1]) {
				panic(fmt.Sprintf("equal blocks. i=%d n=%d\n\t% x\n\t% x",
					i, n, b[:10], blocks[i+1][:10]))
			}
		}
		l += len(b)
	}
	fmt.Printf("n=%d l=%d size=%d\n", n, l, (l+n-1)/n)
}

func fileSize(f os.File) int64 {
	fi, err := f.Stat()
	if err != nil {
		panic(err)
	}
	return fi.Size()
}

func isNull(b []byte) bool {
	for _, x := range b {
		if x != 0 {
			return false
		}
	}
	return true
}

func isEqual(b1, b2 []byte) bool {
	if len(b1) != len(b2) {
		return false
	}
	for i, x := range b1 {
		if x != b2[i] {
			return false
		}
	}
	return true
}
