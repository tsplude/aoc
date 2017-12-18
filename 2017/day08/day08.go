package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var DATA_FILE string = os.Args[1]

// Evaluate a conditional expression e.g. "a > 1" -> "reg_to_balance[`reg`] `op` `rhs`"
func condition_passes(reg, op string, rhs int, reg_to_balance map[string]int) bool {
	result := false
	switch op {
	case "<":
		result = reg_to_balance[reg] < rhs
	case "<=":
		result = reg_to_balance[reg] <= rhs
	case ">":
		result = reg_to_balance[reg] > rhs
	case ">=":
		result = reg_to_balance[reg] >= rhs
	case "==":
		result = reg_to_balance[reg] == rhs
	case "!=":
		result = reg_to_balance[reg] != rhs
	default:
		panic(fmt.Sprintf("Unsupported relational operator: %v", op))
	}
	return result
}

// Find the largest balance of all registers in this current state
func find_max(reg_to_balance map[string]int) int {
	max := math.Inf(-1)
	for _, val := range reg_to_balance {
		if float64(val) > max {
			max = float64(val)
		}
	}
	return int(max)
}

// Solve part1 & part2 in tandem
func solve(data_file string) (int, int) {
	file, _ := os.Open(data_file)
	scanner := bufio.NewScanner(file)

	reg_to_balance := make(map[string]int) // register's account balance

	part_two := 0
	for scanner.Scan() {
		// ["b", "inc", "5", "if", "a", ">", "1"]
		instruction := strings.Split(scanner.Text(), " ")

		// "b"
		reg_name := instruction[0]

		// -5 if "dec" else 5
		instruction_val, _ := strconv.Atoi(instruction[2])
		if instruction[1] == "dec" {
			instruction_val *= -1
		}

		condition_reg := instruction[4]                  // "a"
		condition_op := instruction[5]                   // ">"
		condition_rhs, _ := strconv.Atoi(instruction[6]) // 1

		// Possible that either of the two referenced registers not yet in map of balances
		if _, ok := reg_to_balance[reg_name]; !ok {
			reg_to_balance[reg_name] = 0
		}
		if _, ok := reg_to_balance[condition_reg]; !ok {
			reg_to_balance[condition_reg] = 0
		}

		// if balance of "a" > 1, add instruction_val to balance of "b"
		if condition_passes(condition_reg, condition_op, condition_rhs, reg_to_balance) {
			reg_to_balance[reg_name] += instruction_val
		}

		// Keep track of largest register balance seen thus far
		current_max := find_max(reg_to_balance)
		if current_max > part_two {
			part_two = current_max
		}
	}

	// Find end state max balance
	part_one := find_max(reg_to_balance)

	return part_one, part_two
}

func main() {
	part_one, part_two := solve(DATA_FILE)

	fmt.Println("Part 1:", part_one)
	fmt.Println("Part 2:", part_two)
}
