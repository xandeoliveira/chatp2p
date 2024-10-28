package main

import (
	"fmt"
	cipher "packages/caesarcipher"
)

func main() {
	var k byte = 2
	message := "Alexandre"

	encrypted := cipher.Encoding(message, k)
	dencrypted := cipher.Decoding(encrypted, k)

	fmt.Println(encrypted)
	fmt.Println(dencrypted)
}
