package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var DATA_FILE string = os.Args[1]

func regexp_replace(target, re_str, re_replace string) string {
	re := regexp.MustCompile(re_str)
	return re.ReplaceAllLiteralString(target, re_replace)
}

func part_one(data_file string) bool {
	file, _ := os.Open(data_file)
	defer file.Close()
	reader := bufio.NewReader(file)

	input, _ := reader.ReadString('\n')
	steps_str := strings.TrimSpace(input)
	fmt.Println("Steps:", steps_str)

	replacement_map := map[string]string{
		"n,s,|n,s$":     "",
		"n,se":          "ne",
		"n,sw":          "nw",
		"nw,se,|nw,se$": "",
		"nw,ne":         "n",
		"nw,s":          "sw",
		"sw,ne,|sw,ne$": "",
		"sw,n":          "nw",
		"sw,se":         "s",
		"s,n,|s,n$":     "",
		"s,nw":          "sw",
		"s,ne":          "se",
		"se,nw,|se,nw$": "",
		"se,sw":         "s",
		"se,n":          "ne",
		"ne,sw,|ne,sw$": "",
		"ne,nw":         "n",
		"ne,s":          "se",
	}

	before_len := len(steps_str)
	for {
		for re_str, repl := range replacement_map {
			after_str := regexp_replace(steps_str, re_str, repl)

			fmt.Printf("RE: %s  W: %s  BEF: %s  AFT: %s\n", re_str, repl, steps_str, after_str)
			steps_str = after_str
		}
		if before_len == len(steps_str) {
			break
		} else {
			before_len = len(steps_str)
		}
	}

	fmt.Println("After:", steps_str)

	return false
}

func main() {
	fmt.Println("Part 1:", part_one(DATA_FILE))
}
