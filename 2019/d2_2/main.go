package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	code    []string
	intCode []int
)

const num = 19690720

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		code = strings.Split(scanner.Text(), ",")
	}

	for _, s := range code {
		v, err := strconv.Atoi(s)
		if err != nil {
			return
		}
		intCode = append(intCode, v)
	}

	length := len(intCode)

	original := make([]int, length)
	copied := copy(original, intCode)
	if copied != length {
		return
	}
	match := false
	noun := 0
	verb := 0

check:
	for noun <= 100 {
		for verb <= 100 {
			intCode[1] = noun
			intCode[2] = verb
			match = processCode()

			if match {
				break check
			} else {
				copied := copy(intCode, original)
				if copied != length {
					return
				}
			}
			verb++
		}
		noun++
		verb = 0
	}

	fmt.Println(strings.Join(code, ","))
}

func processCode() bool {
	for i := 0; i < len(intCode); i = i + 4 {
		if intCode[i] == 99 {
			break
		}

		pos2 := intCode[i+1]
		pos3 := intCode[i+2]
		output := intCode[i+3]

		switch intCode[i] {
		case 1:
			intCode[output] = intCode[pos2] + intCode[pos3]
		case 2:
			intCode[output] = intCode[pos2] * intCode[pos3]
		}
		code[output] = strconv.Itoa(intCode[output])
	}

	if intCode[0] == num {
		code[1] = strconv.Itoa(intCode[1])
		code[2] = strconv.Itoa(intCode[2])
		return true
	}
	return false
}
