package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var input []string

	for scanner.Scan() {
		input = strings.Split(scanner.Text(), "-")
	}

	min, err := strconv.Atoi(input[0])
	if err != nil {
		log.Println("error converting min input")
	}

	max, err := strconv.Atoi(input[1])
	if err != nil {
		log.Println("error converting max input")
	}

	fmt.Println(pass(min, max))
}

func pass(min, max int) int {
	var passwords []int
	for min <= max {
		str := strconv.Itoa(min)

		valid := true
		double := false

		for i := 0; i+1 < len(str); i++ {
			if int(str[i]) > int(str[i+1]) {
				valid = false
			}
			if str[i] == str[i+1] {
				double = true
			}
		}

		if valid && double {
			passwords = append(passwords, min)
		}
		min++
	}
	return len(passwords)
}
