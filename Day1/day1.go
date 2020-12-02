package main

/*
1 - Get it running -   
2 - Get it running correctly & submit -   
3 - Refactor - tbd
*/

//In this input list, there are two expense entries that sum to 2020.
//Multiplying them together produces the correct answer.

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

/* PART 1
func main() {

	lines := GetLines("./puzzle_input.txt")
	expenses := ConvertToInt(lines)
	required_sum := 2020

	for i1, _ := range expenses {
		for i2 := i1+1; i2<len(expenses); i2++ {
			if expenses[i1] + expenses[i2] == required_sum {
			fmt.Println("Expense Values are", expenses[i1], expenses[i2])
			fmt.Println("Product is", expenses[i1] * expenses[i2])
			}
		}
	}
}
*/
//PART 2
func main() {

	lines := GetLines("./puzzle_input.txt")
	expenses := ConvertToInt(lines)
	requiredSum := 2020
	var expenseProduct int

	for i1, _ := range expenses {
		for i2, _ := range expenses {
			for i3 := i2 + 1; i3 < len(expenses); i3++ {
					expenseProduct = expenses[i1] * expenses[i2] * expenses[i3]
					fmt.Println("Expense Values are", expenses[i1], expenses[i2], expenses[i3])
					fmt.Println("Product is", expenseProduct)
					return
				}
			}
		}

	}
}

func GetLines(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	return strings.Split(string(data), "\n")
}

func ConvertToInt(lines []string) []int {
	digits := []int{}
	for _, line := range lines {
		digit, _ := strconv.Atoi(line)
		digits = append(digits, digit)
	}
	return digits
}
