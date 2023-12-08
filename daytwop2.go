package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CubeColors struct {
	rawString string
	green     int
	blue      int
	red       int
	valid     bool
}

const finalRedCount int = 12
const finalGreenCount int = 13
const finalBlueCount int = 14

func D2P2() {

	// opening the file
	file, err := os.Open("supportfiles/cubes.txt")

	// check for error
	if err != nil {
		fmt.Println("Error opening the entered file: ", err)
		return
	}

	defer file.Close()

	// creating a scanner for reading the file line by line
	scanner := bufio.NewScanner(file)

	// creating a string array to store all the values of the read line
	var cubeStoreArray []string

	// iterating over each line
	for scanner.Scan() {
		// checking for any potential error
		if err := scanner.Err(); err != nil {
			fmt.Println("Error scanning the doc: ", err)
			return
		}

		// if all good, append the string to the array's end
		cubeStoreArray = append(cubeStoreArray, scanner.Text())
	}

	// checking the number of strings
	fmt.Println("Number of strings: ", len(cubeStoreArray))

	// the cube color details will be stored in a map[int]CubeColors
	var cubeMaps map[int]CubeColors = make(map[int]CubeColors)

	cubeSum := 0

	for id, str := range cubeStoreArray {
		// split the original string
		parsedGameDetail := strings.Split(str, ":")[1]
		// fmt.Printf("\n\nParsed game %d detail: %s\n\n", id, parsedGameDetail)

		cubeMaps[id+1] = CubeColors{
			rawString: parsedGameDetail,
			green:     0,
			red:       0,
			blue:      0,
		}

		cubeColors := strings.Split(strings.Trim(parsedGameDetail, " "), "; ")

		fmt.Printf("For game %d: \n\n", id)
		for individualPickIndex, allColors := range cubeColors {
			individualColorPick := strings.Split(strings.Trim(allColors, " "), ",")
			fmt.Printf("Pick number %d:\n", individualPickIndex)
			for _, colorCountName := range individualColorPick {
				// fmt.Printf("%s\n", strings.Trim(colorCountName, " "))
				singleColorBreakdown := strings.Split(strings.Trim(colorCountName, " "), " ")
				if singleColorBreakdown[1] == "green" {
					greenCount, greenErr := strconv.Atoi(singleColorBreakdown[0])
					if greenErr == nil {
						fmt.Printf("%d Green\n", greenCount)
						if entry, ok := cubeMaps[id+1]; ok {
							if greenCount > entry.green {
								entry.green = greenCount
							}
							cubeMaps[id+1] = entry
						}
					}
				} else if singleColorBreakdown[1] == "red" {
					redCount, redErr := strconv.Atoi(singleColorBreakdown[0])
					if redErr == nil {
						fmt.Printf("%d Red\n", redCount)
						if entry, ok := cubeMaps[id+1]; ok {
							if redCount > entry.red {
								entry.red = redCount
							}
							cubeMaps[id+1] = entry
						}
					}
				} else if singleColorBreakdown[1] == "blue" {
					blueCount, blueErr := strconv.Atoi(singleColorBreakdown[0])
					if blueErr == nil {
						fmt.Printf("%d Blue\n", blueCount)
						if entry, ok := cubeMaps[id+1]; ok {
							if blueCount > entry.blue {
								entry.blue = blueCount
							}
							cubeMaps[id+1] = entry
						}
					}
				}
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}

	for id, cubeMap := range cubeMaps {
		fmt.Printf("Key is %d, value is %d green, %d red, %d blue\n\n", id, cubeMap.green, cubeMap.red, cubeMap.blue)
		getRedVal := cubeMap.red
		getBlueVal := cubeMap.blue
		getGreenVal := cubeMap.green

		if getRedVal == 0 {
			getRedVal = 1
		}

		if getBlueVal == 0 {
			getBlueVal = 1
		}

		if getGreenVal == 0 {
			getGreenVal = 1
		}

		cubeSum += getRedVal * getBlueVal * getGreenVal
	}

	fmt.Printf("Sum: %d\n", cubeSum)
}
