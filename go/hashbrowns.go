package main

import "crypto/sha256"
import "fmt"

func main() {
	var username = "williamjackson"
	var base64_chars = "+/0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	best := "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
	for _, c := range base64_chars {
		for _, d := range base64_chars {
			value := fmt.Sprintf("%s/%c%c", username, c, d)
			sha := sha256.Sum256([]byte(value))
			sha_str := fmt.Sprintf("%x", sha)
			if (sha_str < best) {
				fmt.Println(value, sha_str)
				best = sha_str
			}
		}
	}
}
