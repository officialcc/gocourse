package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {

	password := "password123"

	hash256 := sha256.Sum256([]byte(password))
	hash512 := sha512.Sum512([]byte(password))

	fmt.Println(password)
	fmt.Println(hash256)
	fmt.Printf("SHA-256 Hash hex value: %x\n", hash256)

	fmt.Println(hash512)
	fmt.Printf("SHA-512 Hash hex value: %x\n", hash512)

}