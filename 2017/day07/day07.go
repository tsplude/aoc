package main

import (
	"bufio"
	"fmt"
	"github.com/deckarep/golang-set"
	"os"
	"strings"
)

func part_one() string {
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)

	lhs := mapset.NewSet()
	rhs := mapset.NewSet()

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Count(line, "->") > 0 {
			// Handle LHS membership first
			lhs_raw := strings.SplitAfter(line, "->")[0]
			lhs_id := strings.Split(lhs_raw, " ")[0]
			if !lhs.Contains(lhs_id) {
				lhs.Add(lhs_id)
			}

			// RHS membership
			rhs_raw := strings.Trim(strings.SplitAfter(line, "->")[1], " ")
			for _, rhs_id := range strings.Split(rhs_raw, " ") {
				rhs_id = strings.Trim(rhs_id, ", ")
				if !rhs.Contains(rhs_id) {
					rhs.Add(rhs_id)
				}
			}
		} else {
			// Leaf nodes
			lhs_id := strings.Split(line, " ")[0]
			if !lhs.Contains(lhs_id) {
				lhs.Add(lhs_id)
			}
		}
	}

	return lhs.Difference(rhs).String()
}

func part_two() bool {
	return false
}

func main() {
	fmt.Println("Part 1:", part_one())
	fmt.Println("Part 2:", part_two())
}
