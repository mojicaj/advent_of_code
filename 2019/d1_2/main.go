package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var values []int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		mass, err := strconv.ParseFloat(scanner.Text(), 32)
		if err != nil {
			return
		}

		fuel := calcFuel(mass)

		for fuel > 0 {
			values = append(values, fuel)
			fuel = calcFuel(float64(fuel))
		}
	}

	if err := scanner.Err(); err != nil {
		return
	}

	fmt.Println(sum(values))
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
