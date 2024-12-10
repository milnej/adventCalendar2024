package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Split(r rune) bool {
	return r == '(' || r == ',' || r == ')'
}

func day3Q1() {
	dat, err := os.ReadFile("test_data/day3_1.txt")
	check(err)
	input := string(dat)

	// Not sure why this doesn't work but the other implementation is definitely better
	// candidates := strings.Split(input, "mul(")
	// answer := 0
	// for _, candidate := range candidates {
	// 	numbers := strings.Split(candidate, ",")

	// 	num1, err := strconv.Atoi(numbers[0])
	// 	if err != nil {
	// 		fmt.Println(candidate)
	// 		continue
	// 	}
	// 	number2Str := strings.Split(numbers[1], ")")
	// 	num2, err := strconv.Atoi(number2Str[0])
	// 	if err != nil {
	// 		fmt.Println(candidate)
	// 		continue
	// 	}
	// 	//fmt.Println(candidate, answer)
	// 	answer += (num1 * num2)
	// }
	// fmt.Println(answer)

	exp := regexp.MustCompile(`mul\([\d]{1,3},[\d]{1,3}\)`)
	matches := exp.FindAllString(input, 10000)

	answer := 0
	for _, match := range matches {
		numbers := strings.FieldsFunc(match, Split)

		num1, _ := strconv.Atoi(numbers[1])
		num2, _ := strconv.Atoi(numbers[2])
		answer += num1 * num2
	}
	fmt.Println(answer)
}
