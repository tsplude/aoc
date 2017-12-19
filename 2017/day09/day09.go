package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var DATA_FILE string = os.Args[1]

func solve(datafile string) {
	file, _ := os.Open(datafile)
	scanner := bufio.NewScanner(file)

	re_ignored := regexp.MustCompile("!.")    // Regex to match all valid ignored characters
	re_garbage := regexp.MustCompile("<.*?>") // Regex to match all valid garbage characters

	for scanner.Scan() {
		// e.g. "<{!>}>"
		raw_text := scanner.Text()

		// "<{}>"
		purge_ignored := re_ignored.ReplaceAllLiteralString(raw_text, "")

		// len("<{!>}>") - len("<{}>") = 6 - 4 = 2
		num_ignored := len(raw_text) - len(purge_ignored)

		// "<>"
		purge_garbage := re_garbage.ReplaceAllLiteralString(purge_ignored, "<>")

		// len("<{!>}>") - len("<>") - 2 = 6 - 2 - 2 = 2
		num_garbage_chars := len(raw_text) - len(purge_garbage) - num_ignored

		total_score, parent_group_score := 0, 0
		for _, char := range purge_garbage {
			if char == '{' {
				parent_group_score++
			} else if char == '}' {
				total_score += parent_group_score
				parent_group_score--
			}
		}
		fmt.Println("Score....:", total_score)
		fmt.Println("# Garbage:", num_garbage_chars)
	}
}

func main() {
	solve(DATA_FILE)
}
