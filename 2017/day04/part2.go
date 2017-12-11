package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	file, _ := os.Open("data.txt")
	defer file.Close()

	num_valid := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		phrase_valid := true
		passphrase := strings.Split(scanner.Text(), " ")
		phrase_map := make(map[string]bool)
		for _, word := range passphrase {
			chars := strings.Split(word, "")
			sort.Strings(chars)
			sorted_str := strings.Join(chars, "")
			_, ok := phrase_map[sorted_str]
			if ok {
				phrase_valid = false
				break
			} else {
				phrase_map[sorted_str] = true
			}
		}
		if phrase_valid {
			num_valid += 1
		}
	}
	fmt.Println("Answer:", num_valid)
}
