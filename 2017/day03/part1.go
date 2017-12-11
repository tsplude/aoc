package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func get_edge_mid(bound, length, step int) int {
	start := bound - length*(step-1)
	end := bound - length*(step)
	return (start + end) / 2
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Expected one command line argument of type <int>")
		return
	}

	x, _ := strconv.ParseFloat(os.Args[1], 64)

	for i := 1.0; ; i += 2.0 {
		odd_sq := math.Pow(i, 2.0)
		if odd_sq >= x {

			square := int(odd_sq)
			ring := int(i) / 2
			edge_length := ring * 2

			min := x
			edges := [4]int{1, 2, 3, 4}
			for _, e := range edges {
				mid := float64(get_edge_mid(square, edge_length, e))
				if math.Abs(x-mid) < min {
					min = math.Abs(x - mid)
				}
			}

			fmt.Println("Distance:", ring+int(min))

			break
		}
	}
}
