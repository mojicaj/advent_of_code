package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var code []string
	var intCode []int

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

	fmt.Println(strings.Join(code, ","))
}
