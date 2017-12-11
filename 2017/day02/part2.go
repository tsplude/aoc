package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func find_divisors(a []int) (int, int) {
	if len(a) == 0 {
		return 0, 0
	}
	div1, div2 := a[0], a[0]
	for pos1, elem1 := range a {
		for pos2, elem2 := range a {
			if pos1 == pos2 {
				continue
			}
			if elem1%elem2 == 0 {
				div1 = elem1
				div2 = elem2
			}
		}
	}
	return div1, div2
}

func main() {

	file, _ := os.Open("data.txt")
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		int_strs := strings.Split(scanner.Text(), "\t")
		var ints []int
		for _, x := range int_strs {
			int_x, _ := strconv.Atoi(x)
			ints = append(ints, int_x)
		}

		div1, div2 := find_divisors(ints)
		sum += div1 / div2
	}
	fmt.Println("Answer:", sum)

}
