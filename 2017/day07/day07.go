package main

import (
	"bufio"
	"fmt"
	"github.com/deckarep/golang-set"
	"os"
	"strconv"
	"strings"
)

func part_one(input_file string) string {
	file, _ := os.Open(input_file)
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

func find_weight(w_map map[string]int, n_map map[string][]string, id string) int {
	if _, ok := n_map[id]; !ok {
		return w_map[id]
	}

	w_sum := w_map[id]
	for _, n := range n_map[id] {
		w_sum += find_weight(w_map, n_map, n)
	}
	return w_sum
}

func part_two(input_file string) string {
	file, _ := os.Open(input_file)
	scanner := bufio.NewScanner(file)

	weight_map := make(map[string]int)
	neighbor_map := make(map[string][]string)

	for scanner.Scan() {
		line := scanner.Text()

		// get LHS id
		lhs_id := strings.Split(line, " ")[0]

		// get LHS weight
		lidx := strings.Index(line, "(") + 1
		ridx := strings.Index(line, ")")
		w, _ := strconv.Atoi(line[lidx:ridx])

		_, ok := weight_map[lhs_id]
		if !ok {
			weight_map[lhs_id] = w
		}

		if strings.Count(line, "->") > 0 {
			neighbors := []string{}
			neighbors_str := strings.Split(line, "->")[1]
			for _, n := range strings.Split(neighbors_str, ",") {
				neighbors = append(neighbors, strings.Trim(n, ", "))
			}

			_, ok := neighbor_map[lhs_id]
			if !ok {
				neighbor_map[lhs_id] = neighbors
			}
		}
	}

	fmt.Println("mkxke:", find_weight(weight_map, neighbor_map, "mkxke"))

	return "not yet"
}

func main() {
	fmt.Println("Part 1:", part_one("input.txt"))
	fmt.Println("Part 2:", part_two("input.txt"))
}
