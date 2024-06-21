package main

import "crypto/sha256"
import "flag"
import "fmt"

const base64_chars = "+/0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func appendAndSend(prefix string, num_to_append int, channel chan string, close_channel bool) {
	for _, c := range base64_chars {
		value := prefix + string(c)
		if (num_to_append < 2) {
			channel <- value
		} else {
			appendAndSend(value, num_to_append-1, channel, false)
		}
	}

	if (close_channel) {
		close(channel)
	}
}

func main() {
	var username string
	var length int
	flag.IntVar(&length, "length", 1, "")
	flag.StringVar(&username, "username", "williamjackson", "")
	flag.Parse()

	value_channel := make(chan string)

	starting_prefix := username + "/"
	go appendAndSend(starting_prefix, length, value_channel, true)

	best := "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"

	for value := range value_channel {
		sha := sha256.Sum256([]byte(value))
		sha_str := fmt.Sprintf("%x", sha)
		if (sha_str < best) {
			fmt.Println(value, sha_str)
			best = sha_str
		}
	}
	fmt.Println()
}
