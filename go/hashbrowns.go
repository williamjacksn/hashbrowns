package main

import "crypto/sha256"
import "flag"
import "fmt"
import "sync"

const base64_chars = "+/0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type ValueWithHash struct {
	value, hash string
}

func GenerateValueAndSend(prefix string, num_to_append int, channel chan string, close_channel bool) {
	for _, c := range base64_chars {
		value := prefix + string(c)
		if num_to_append < 2 {
			channel <- value
		} else {
			GenerateValueAndSend(value, num_to_append-1, channel, false)
		}
	}

	if close_channel {
		close(channel)
	}
}

func ReceiveValue(valueChannel chan string, valueWithHashChannel chan ValueWithHash) {
	var wg sync.WaitGroup
	for value := range valueChannel {
		wg.Add(1)
		go func(){
			defer wg.Done()
			CalculateHashAndSend(value, valueWithHashChannel)
		}()
	}
	wg.Wait()
	close(valueWithHashChannel)
}

func CalculateHashAndSend(value string, hashChannel chan ValueWithHash) {
	sha := sha256.Sum256([]byte(value))
	shaString := fmt.Sprintf("%x", sha)
	result := ValueWithHash{value, shaString}
	hashChannel <- result
}

func formatSha(sha string) string {
	formatted := sha[0:8] + " " + sha[8:16] + " " + sha[16:24] + " " + sha[24:32] + " " + sha[32:40] + " " + sha[40:48] + " " + sha[48:56] + " " + sha[56:]
	return formatted
}

func main() {
	var length int
	flag.IntVar(&length, "length", 1, "")
	var username string
	flag.StringVar(&username, "username", "williamjackson", "")
	flag.Parse()

	valueChannel := make(chan string)
	valueWithHashChannel := make(chan ValueWithHash)

	startingPrefix := username + "/"
	go GenerateValueAndSend(startingPrefix, length, valueChannel, true)
	go ReceiveValue(valueChannel, valueWithHashChannel)

	best := "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"

	count := 0
	for valueWithHash := range valueWithHashChannel {
		count++
		if count%100000 == 0 {
			fmt.Print(valueWithHash.value + "\r")
		}
		if valueWithHash.hash < best {
			fmt.Println(valueWithHash.value, formatSha(valueWithHash.hash))
			best = valueWithHash.hash
		}
	}

	fmt.Println()
}
