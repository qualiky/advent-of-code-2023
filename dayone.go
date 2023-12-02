package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func Day1() {
	// opening the file
	file, err := os.Open("supportfiles/secret_code.txt")

	// check for error
	if err != nil {
		fmt.Println("Error opening the entered file: ", err)
		return
	}

	defer file.Close()

	// creating a scanner for reading the file line by line
	scanner := bufio.NewScanner(file)

	// creating a string array to store all the values of the read line
	var lineStoreArray []string

	// iterating over each line
	for scanner.Scan() {
		// checking for any potential error
		if err := scanner.Err(); err != nil {
			fmt.Println("Error scanning the doc: ", err)
			return
		}

		// if all good, append the string to the array's end
		lineStoreArray = append(lineStoreArray, scanner.Text())
	}

	// checking the number of strings
	fmt.Println("Number of strings: ", len(lineStoreArray))

	// create a sum variable to hold the sums
	sum := 0

	// performing the number detection operation on the string array
	for _, str := range lineStoreArray {
		firstNum, lastNum := FindFirstAndLastNumber(str)

		// check if the values returned are -1, for edge case handling
		if firstNum != -1 && lastNum != -1 {

			// extracting the runes from the first and last digit instance indices
			firstValAsRune := rune(str[firstNum])
			lastValAsRune := rune(str[lastNum])

			// converting the runes to integer format
			digit1, _ := strconv.Atoi(string(firstValAsRune))
			digit2, _ := strconv.Atoi(string(lastValAsRune))

			// creating two digit numbers
			sum += digit1*10 + digit2

		}
	}

	fmt.Println("The sum is: ", sum)

}

// function to find first and last number in a string
func FindFirstAndLastNumber(text string) (first int, last int) {
	first = -1
	last = -1

	// traversing the length of string to find the numbers
	for i, char := range text {
		if unicode.IsDigit(char) {
			if first == -1 {
				first = i
			}
			last = i
		}
	}

	if last == -1 {
		last = first
	}

	return first, last
}
