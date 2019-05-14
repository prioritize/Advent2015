package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	hash := md5.New()
	hash_begin := byte("ckczppom")
	hash.Sum(hash_begin)
	fmt.Println()
}
