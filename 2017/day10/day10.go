package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var DATA_FILE string = os.Args[1]

func knot_hash_round(elements []int, sequence []string, pos, skip int) ([]int, int, int) {
	n := len(elements)
	for _, len_str := range sequence {
		len, _ := strconv.Atoi(len_str)

		// copy current state
		elem_copy := make([]int, n)
		copy(elem_copy, elements)

		last_pos := (pos + len - 1) % n
		for i := len; i > 0; i-- {
			elem_copy[pos] = elements[last_pos]

			// adjust reference points
			if last_pos == 0 {
				last_pos = n - 1
			} else {
				last_pos--
			}
			pos = (pos + 1) % n
		}

		elements = elem_copy
		pos = (pos + skip) % n
		skip += 1
	}
	return elements, pos, skip
}

func part_one(data_file string) int {
	file, _ := os.Open(data_file)
	defer file.Close()
	reader := bufio.NewReader(file)

	// Get length sequence (puzzle input)
	input, _ := reader.ReadString('\n')
	sequence := strings.Split(strings.TrimSpace(input), ",")

	// elements size
	n := 256

	// build our initial state of elements [0...num_elems]
	elements := make([]int, n)
	for i := 0; i < n; i++ {
		elements[i] = i
	}

	elements, _, _ = knot_hash_round(elements, sequence, 0, 0)

	return elements[0] * elements[1]
}

func part_two(data_file string) bool {
	file, _ := os.Open(data_file)
	defer file.Close()
	reader := bufio.NewReader(file)

	// Get length sequence (puzzle input)
	input, _ := reader.ReadString('\n')

	// elements size
	n := 256

	// build our initial state of elements [0...num_elems]
	elements := make([]int, n)
	for i := 0; i < n; i++ {
		elements[i] = i
	}

	ascii_codes := []rune{}
	for _, c := range strings.TrimSpace(input) {
		ascii_codes = append(ascii_codes, c)
	}

	// Append standard suffix
	ascii_codes = append(ascii_codes, 17)
	ascii_codes = append(ascii_codes, 31)
	ascii_codes = append(ascii_codes, 73)
	ascii_codes = append(ascii_codes, 47)
	ascii_codes = append(ascii_codes, 23)

	fmt.Println("codes:", ascii_codes)

	return false
}

func main() {
	fmt.Println("Part 1:", part_one(DATA_FILE))
	fmt.Println("Part 2:", part_two(DATA_FILE))
}
