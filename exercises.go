package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println("\nExercise 1")
	CBCDecrypter("140b41b22a29beb4061bda66b6747e14",
		"4ca00ff4c898d61e1edbf1800618fb2828a226d160dad07883d04e008a7897ee"+
			"2e4b7465d5290d0c0e6c6822236e1daafb94ffe0c5da05d9476be028ad7c1d81")
	fmt.Println("\nExercise 2")
	CBCDecrypter("140b41b22a29beb4061bda66b6747e14",
		"5b68629feb8606f9a6667670b75b38a5b4832d0f2"+
			"6e1ab7da33249de7d4afc48e713ac646ace36e872ad5fb8a512428a6e21364b0c374df45503473c5242a253")
	fmt.Println("\nExercise 3")
	CTRDecrypter("36f18357be4dbd77f050515c73fcf9f2",
		"69dda8455c7dd4254bf353b773304eec0ec7702330098ce7f7520d1cbbb20fc388d1b0adb5054dbd7370849dbf0b88d393f252e764f1f5f7ad97ef79d59ce29f5f51eeca32eabedd9afa9329")
	fmt.Println("\nExercise 4")
	CTRDecrypter("36f18357be4dbd77f050515c73fcf9f2",
		"770b80259ec33beb2561358a9f2dc617e46218c0a53cbeca695ae45faa8952aa0e311bde9d4e01726d3184c34451")

}

func CBCDecrypter(keyStr, cipherStr string) {
	key, _ := hex.DecodeString(keyStr)
	ciphertext, _ := hex.DecodeString(cipherStr)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV is at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// CBC mode always works in whole blocks.
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	// CryptBlocks can work in-place if the two arguments are the same.
	mode.CryptBlocks(ciphertext, ciphertext)

	fmt.Printf("Answer=%d %#q\n", len(ciphertext), ciphertext)
	pad := int(ciphertext[len(ciphertext)-1])
	fmt.Printf("pad=%d\n", pad)
	fmt.Printf("Answer=%d %#q\n", len(ciphertext[:len(ciphertext)-pad]), ciphertext[:len(ciphertext)-pad])
	for i := 0; i < pad; i++ {
		p := int(ciphertext[len(ciphertext)-1-i])
		if p != pad {
			fmt.Printf("Bad pad. pad=%d p=%d i=%d\n", pad, p, i)
			panic("Bad pad 1")
		}
	}
	p := int(ciphertext[len(ciphertext)-1-pad])
	if p == pad {
		fmt.Printf("Bad pad. pad=%d p=%d i=%d\n", pad, p, pad)
		panic("Bad pad 2")
	}
}

func CTRDecrypter(keyStr, cipherStr string) {
	key, _ := hex.DecodeString(keyStr)
	ciphertext, _ := hex.DecodeString(cipherStr)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV is at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	plaintext := make([]byte, len(ciphertext))
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plaintext, ciphertext)

	fmt.Printf("Answer=%d %#q\n", len(plaintext), plaintext)
}
