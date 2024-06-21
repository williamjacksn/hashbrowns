package main

import "crypto/sha256"
import "flag"
import "fmt"
import "sync"

const alphabet = "+/0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

type valueWithHash struct {
	value, hash string
}

func generateValueAndSend(prefix string, numToAppend int, channel chan string, closeChannel bool) {
	for _, c := range alphabet {
		value := prefix + string(c)
		if numToAppend < 2 {
			channel <- value
		} else {
			generateValueAndSend(value, numToAppend-1, channel, false)
		}
	}

	if closeChannel {
		close(channel)
	}
}

func receiveValue(valueChannel chan string, valueWithHashChannel chan valueWithHash) {
	var wg sync.WaitGroup
	for value := range valueChannel {
		wg.Add(1)
		go func(){
			defer wg.Done()
			calculateHashAndSend(value, valueWithHashChannel)
		}()
	}
	wg.Wait()
	close(valueWithHashChannel)
}

func calculateHashAndSend(value string, hashChannel chan valueWithHash) {
	sha := sha256.Sum256([]byte(value))
	shaString := fmt.Sprintf("%x", sha)
	result := valueWithHash{value, shaString}
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
	valueWithHashChannel := make(chan valueWithHash)

	startingPrefix := username + "/"
	go generateValueAndSend(startingPrefix, length, valueChannel, true)
	go receiveValue(valueChannel, valueWithHashChannel)

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
