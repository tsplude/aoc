package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, _ := os.Open("data.txt")
	defer file.Close()

	var ints []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		new_instruction, _ := strconv.Atoi(scanner.Text())
		ints = append(ints, new_instruction)
	}

	idx, step_count := 0, 0
	for {
		new_idx := ints[idx]
		if ints[idx] >= 3 {
			ints[idx] -= 1
		} else {
			ints[idx] += 1
		}
		if new_idx+idx >= len(ints) {
			step_count += 1
			fmt.Println("Answer:", step_count)
			break
		} else {
			idx = new_idx + idx
			step_count += 1
		}
	}
}
