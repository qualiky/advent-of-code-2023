package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func Day1P2() {

	// creating the array of all the individual list of words
	var onesList []string = []string{
		"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen",
	}

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
	for ran, str := range lineStoreArray {
		fmt.Println("Index ", ran, ": string is ", str)
		firstNum, lastNum, firstConv, lastConv := FindFirstAndLastNumberAndString(str, onesList)

		// check if the values returned are -1, for edge case handling
		if firstNum != -1 && lastNum != -1 {

			firstValAsRune := 'A'
			digit1 := -1

			lastValAsRune := 'A'
			digit2 := -1

			if firstConv == -1 {
				firstValAsRune = rune(str[firstNum])
				digit1, _ = strconv.Atoi(string(firstValAsRune))
			} else {
				digit1 = firstConv
			}

			if lastConv == -1 {
				lastValAsRune = rune(str[lastNum])
				digit2, _ = strconv.Atoi(string(lastValAsRune))
			} else {
				digit2 = lastConv
			}

			fmt.Printf("First number is %d and last number is %d\n", digit1, digit2)
			// creating two digit numbers
			sum += digit1*10 + digit2

		}
	}

	fmt.Println("The sum is: ", sum)

}

// function to find first and last number in a string
func FindFirstAndLastNumberAndString(text string, substrings []string) (first int, last int, firstConv int, lastConv int) {
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

	// fmt.Println("Before string work: first and last index are ", first, " and ", last)

	allSubstringWithIndices := findIndicesOfMultipleSubstrings(text, substrings)

	// get the value of first and last key-value from the map
	if len(allSubstringWithIndices) == 0 {
		// fmt.Println("Length of allSubstringWithIndices is 0")
		return first, last, -1, -1
	}

	finalMinMaxMap := make(map[int]int)

	newMin := 100000000
	newMax := -1

	for key, value := range allSubstringWithIndices {

		if key < newMin {
			newMin = key
			finalMinMaxMap[newMin] = firstIndexOf(substrings, value)
		}

		if key > newMax {
			newMax = key
			finalMinMaxMap[newMax] = firstIndexOf(substrings, value)
		}

		// fmt.Println("Key is: ", key, " and the value is", firstIndexOf(substrings, value))

	}

	// fmt.Printf("New min value is %d at key %d\n", finalMinMaxMap[newMin], newMin)
	// fmt.Printf("New max value is %d at key %d\n", finalMinMaxMap[newMax], newMax)

	sendFirstVal := -1
	sendLastVal := -1

	if newMin != -1 && newMin < first {
		first = newMin
		sendFirstVal = finalMinMaxMap[newMin]
	}

	if newMax != 100000000 && newMax > last {
		last = newMax
		sendLastVal = finalMinMaxMap[newMax]
	}

	return first, last, sendFirstVal, sendLastVal
}

// just store all the substrings, calculate their index, and map them to the index of their arrays

func findIndicesOfMultipleSubstrings(text string, substrings []string) map[int]string {

	// stores all the outputted strings along with their position
	result := make(map[int]string)

	// iterating through the substrings array
	for _, sub := range substrings {
		// using regex to find all the instances of the substring and mapping it to the index of the first character
		matches := regexp.MustCompile(sub).FindAllStringIndex(text, -1)

		// adding matches to the map with their first index as key, substring as value
		for _, match := range matches {
			result[match[0]] = sub
		}
	}

	return result
}

// array to get first position of a substring of an array
func firstIndexOf(arr []string, target string) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}
