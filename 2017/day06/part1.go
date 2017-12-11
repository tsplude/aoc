package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func distribute_block(memory []int, idx int) []int {
	num_blocks := memory[idx]
	memory[idx] = 0
	for num_blocks > 0 {
		idx = (idx + 1) % len(memory)
		memory[idx] += 1
		num_blocks -= 1
	}
	return memory
}

func to_strings(ints []int) []string {
	var strings []string
	for _, x := range ints {
		x_int := strconv.Itoa(x)
		strings = append(strings, x_int)
	}
	return strings
}

func to_ints(strings []string) []int {
	var ints []int
	for _, str := range strings {
		x_str, _ := strconv.Atoi(str)
		ints = append(ints, x_str)
	}
	return ints
}

func find_max_idx(x []int) int {
	max, max_idx := x[0], 0
	for i, elem := range x {
		if elem > max {
			max = elem
			max_idx = i
		}
	}
	return max_idx
}

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		memory := to_ints(strings.Split(scanner.Text(), "\t"))

		count := 0
		state_map := make(map[string]int)
		for {
			max_idx := find_max_idx(memory)
			memory = distribute_block(memory, max_idx)
			count += 1
			state_str := strings.Join(to_strings(memory), "")
			_, ok := state_map[state_str]
			if ok {
				fmt.Println("Found duplicate state:", state_str)
				fmt.Println("Part 1:", count)
				fmt.Println("Part 2:", count-state_map[state_str])
				break
			} else {
				state_map[state_str] = count
			}
		}

	}
}
