package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func find_min_max(a []int) (int, int) {
	if len(a) == 0 {
		return 0, 0
	}
	min := a[0]
	max := a[0]
	for _, e := range a {
		if e < min {
			min = e
		}
		if e > max {
			max = e
		}
	}
	return min, max
}

func main() {

	file, _ := os.Open("data.txt")
	defer file.Close()

	check_sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		int_strs := strings.Split(scanner.Text(), "\t")
		var ints []int
		for _, x := range int_strs {
			int_x, _ := strconv.Atoi(x)
			ints = append(ints, int_x)
		}

		min, max := find_min_max(ints)
		check_sum += max - min
	}
	fmt.Println("Answer:", check_sum)
}
