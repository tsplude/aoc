package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("data.txt")
	defer file.Close()

	num_valid := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		good_phrase := true
		passphrase := strings.Split(scanner.Text(), " ")
		phrase_map := make(map[string]string)
		for _, word := range passphrase {
			_, ok := phrase_map[word]
			if ok {
				good_phrase = false
				break
			} else {
				phrase_map[word] = "exists"
			}
		}
		if good_phrase {
			num_valid += 1
		}
	}
	fmt.Println("Answer:", num_valid)
}
