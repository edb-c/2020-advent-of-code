package main

/*
1 - Get it running -   
2 - Get it running correctly & submit -   
3 - Refactor - tbd
*/

/*
TASK:
 Each line gives the password policy and then the password.
 The password policy indicates the lowest and highest number of times a given letter must appear for the password to be valid.
 For example, 1-3 a means that the password must contain a at least 1 time and at most 3 times.

Hmm, needed steps seem to be:
 read the file/get lines
 for each line 
  - assign to a variable - password policy length 3, letter length 1, password length varies [range based slicing & split?]
  - convert policy to ints - lowest_num / highest_num
  - check password according to policy
Refactor:
  - use type struct
  - break out into separate helper functions
*/

import (
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {

	lines := GetLines("./puzzle_input2.txt")
	validPasswordCount := 0

	for _, line := range lines {
	
		fields := strings.Split(line, " ")
		//fmt.Println(fields[0])
		//fmt.Println(fields[1])
		//fmt.Println(fields[2])
		
		passwordPolicy := strings.Split(fields[0],"-")
		passwordPolicyLow :=  ConvertToInt(passwordPolicy[0])
		passwordPolicyHigh :=  ConvertToInt(passwordPolicy[1])
		letter := fields[1]
		password := fields[2]

	    //fmt.Println("passwordPolicyLow is", passwordPolicyLow)
		//fmt.Println("passwordPolicyHigh is", passwordPolicyHigh)
		//fmt.Println("letter is", letter[0:1])
		//fmt.Println("password is", password)

		/* Part 1

		letterCount := 0


		for _, pwLetter := range password {			
			if string(pwLetter) == letter[0:1] {
				letterCount += 1
			}
		} 
		
		fmt.Println("letterCount is", letterCount)			
		
		if (letterCount >= passwordPolicyLow) && (letterCount <= passwordPolicyHigh) {
			validPasswordCount ++
		}
		fmt.Println("letterCount >= passwordPolicyLow", letterCount >= passwordPolicyLow)
		fmt.Println("letterCount <= passwordPolicyHigh", letterCount <= passwordPolicyHigh)
		*/

		/* Part 2 */
		validPositionCount := 0
		for i, pwLetter := range password {			
    	
			if (string(pwLetter) == letter[0:1]) && (i+1 == passwordPolicyLow) {
				validPositionCount += 1
			}
			if (string(pwLetter) == letter[0:1]) && (i+1 == passwordPolicyHigh) {
				validPositionCount =  validPositionCount + 1	
				//fmt.Println("validPositionCount is", validPositionCount)
			}			
		} 
		if validPositionCount == 1 {
			validPasswordCount ++
		}	
	}
	fmt.Println("validPasswordCount is", validPasswordCount)
}

func GetLines(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	return strings.Split(string(data), "\n")
}

func ConvertToInt(field string) int {
	digit, _ := strconv.Atoi(field)
	return digit
	}
