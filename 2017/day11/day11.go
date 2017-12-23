package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var DATA_FILE string = os.Args[1]

func part_one(data_file string) bool {
	file, _ := os.Open(data_file)
	defer file.Close()
	reader := bufio.NewReader(file)

	input, _ := reader.ReadString('\n')
	steps := strings.Split(strings.TrimSpace(input), ",")
	fmt.Println("Steps:", steps)

	x, y, z := 0, 0, 0
	for _, step := range steps {
		fmt.Println(step)
		switch step {
		case "n":
			fmt.Println("case 'n'")
			y += 1
		case "ne":
			fmt.Println("case 'ne'")
			x += 1
		case "se":
			fmt.Println("case 'se'")
			z += 1
		case "s":
			fmt.Println("case 's'")
			y -= 1
		case "sw":
			fmt.Println("case 'sw'")
			x -= 1
		case "nw":
			fmt.Println("case 'nw'")
			z -= 1
		default:
			panic(fmt.Sprintf("Unsupported step %s", step))
		}
	}
	fmt.Printf("x: %d y: %d z: %d\n", x, y, z)
	return false
}

func main() {
	fmt.Println("Part 1:", part_one(DATA_FILE))
}
