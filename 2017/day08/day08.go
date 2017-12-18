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

func split_instruction(instruction string) []string {
	parts := strings.Split(instruction, " ")
	if len(parts) != 7 {
		panic(fmt.Sprintf("Invalid instruction line: '%s'", instruction))
	}
	// TODO: verify recognized conditional operator
	return parts
}

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

func find_max(reg_to_balance map[string]int) int {
	max := math.Inf(-1)
	for _, val := range reg_to_balance {
		if float64(val) > max {
			max = float64(val)
		}
	}
	return int(max)
}

func solve(data_file string) (int, int) {
	file, _ := os.Open(data_file)
	scanner := bufio.NewScanner(file)

	reg_to_balance := make(map[string]int) // register's account balance

	part_two := 0
	for scanner.Scan() {
		instruction := strings.Split(scanner.Text(), " ")

		reg_name := instruction[0]
		instruction_val, _ := strconv.Atoi(instruction[2])
		if instruction[1] == "dec" {
			instruction_val *= -1
		}
		condition_reg := instruction[4]
		condition_op := instruction[5]
		condition_rhs, _ := strconv.Atoi(instruction[6])

		if _, ok := reg_to_balance[reg_name]; !ok {
			reg_to_balance[reg_name] = 0
		}
		if _, ok := reg_to_balance[condition_reg]; !ok {
			reg_to_balance[condition_reg] = 0
		}

		if condition_passes(condition_reg, condition_op, condition_rhs, reg_to_balance) {
			reg_to_balance[reg_name] += instruction_val
		}

		current_max := find_max(reg_to_balance)
		if current_max > part_two {
			part_two = current_max
		}
	}

	part_one := find_max(reg_to_balance)

	return part_one, part_two
}

func main() {

	part_one, part_two := solve(DATA_FILE)

	fmt.Println("Part 1:", part_one)
	fmt.Println("Part 2:", part_two)
}
