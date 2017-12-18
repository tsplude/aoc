package main

import (
	"bufio"
	"fmt"
	"github.com/deckarep/golang-set"
	"os"
	"strconv"
	"strings"
)

var DATA_FILE string = os.Args[1]

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

	diff_str := lhs.Difference(rhs).String()
	if strings.Count(diff_str, ",") > 0 {
		panic("Input not valid! Could not find root node.")
	}
	return diff_str[strings.Index(diff_str, "{")+1 : strings.Index(diff_str, "}")]
}

func find_weight(id string, w_map map[string]int, n_map map[string][]string) int {
	if _, ok := n_map[id]; !ok {
		return w_map[id]
	}

	w_sum := w_map[id]
	for _, n := range n_map[id] {
		w_sum += find_weight(n, w_map, n_map)
	}
	return w_sum
}

func count(a []int, b int) int {
	c := 0
	for _, e := range a {
		if e == b {
			c += 1
		}
	}
	return c
}

func investigate(id string, w_map map[string]int, c_map map[string][]string) string {

	child_weights := []int{}
	for _, c := range c_map[id] {
		cw := find_weight(c, w_map, c_map)
		fmt.Println(id, c, cw)
		child_weights = append(child_weights, find_weight(c, w_map, c_map))
	}

	fmt.Println(id, c_map[id])
	fmt.Println(id, child_weights)

	found_different := false
	pos_to_investigate := 0
	for pos, cw := range child_weights {
		if count(child_weights, cw) == 1 {
			pos_to_investigate = pos
			found_different = true
			break
		}
	}

	if !found_different {
		return id
	} else {
		return investigate(c_map[id][pos_to_investigate], w_map, c_map)
	}
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

	root := part_one(DATA_FILE)
	return investigate(root, weight_map, neighbor_map)
}

func main() {
	fmt.Println("Part 1:", part_one(DATA_FILE))
	fmt.Println("Part 2:", part_two(DATA_FILE))
}
