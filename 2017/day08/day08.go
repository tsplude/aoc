package main

import (
	"bufio"
	"fmt"
	"os"
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

func part_one(data_file string) bool {
	file, _ := os.Open(data_file)
	scanner := bufio.NewScanner(file)

	instruction_list := []string{}              // raw instruction string
	reg_to_balance := make(map[string]int)      // register's account balance
	reg_to_val := make(map[string]string)       // register's inc/dec amount
	reg_to_condition := make(map[string]string) // register's condition

	for scanner.Scan() {
		line := scanner.Text()

		instruction_list = append(instruction_list, line)

		parts := split_instruction(line)
		reg_name := parts[0]
		reg_val := fmt.Sprintf("%v:%v", parts[1], parts[2])
		reg_condition := fmt.Sprintf("%v:%v:%v", parts[4], parts[5], parts[6])

		reg_to_balance[reg_name] = 0
		reg_to_val[reg_name] = reg_val
		reg_to_condition[reg_name] = reg_condition

	}

	fmt.Println(reg_to_balance)
	fmt.Println(reg_to_val)
	fmt.Println(reg_to_condition)

	return false
}

func part_two(data_file string) bool {
	return false
}

func main() {
	fmt.Println("Part 1:", part_one(DATA_FILE))
	fmt.Println("Part 2:", part_two(DATA_FILE))
}
