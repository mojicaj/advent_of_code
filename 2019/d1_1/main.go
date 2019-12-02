package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var values []int
	var mass float64
	var err error

	scanner := bufio.NewScanner(os.Stdin)
	input := make([]float64, 0)

	for scanner.Scan() {
		mass, err = strconv.ParseFloat(scanner.Text(), 32)
		if err != nil {
			return
		}
		input = append(input, mass)
		values = append(values, calcFuel(mass))
	}

	if err = scanner.Err(); err != nil {
		return
	}

	fmt.Printf("I: %v - %v\n", input, len(input))
	fmt.Println("V: ", values)
	fmt.Println("S: ", sum(values))
}

func calcFuel(m float64) int {
	// casting to int truncates the value essentially rounding down
	return int(m/3 - 2)
}

func sum(v []int) (s int) {
	for _, f := range v {
		s += f
	}
	return
}
