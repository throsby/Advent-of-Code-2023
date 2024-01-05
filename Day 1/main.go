package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var p = fmt.Println

func check(e error) {
	if e != nil {
		p("\n    ****    ****\n\n      * Oops! *\n\n    ****    ****\n")
		panic(e)
	}
}

func main() {

	b, bErr := os.ReadFile("day1.txt")
	check(bErr)

	var re, reErr = regexp.Compile("[a-zA-Z]")
	check(reErr)

	var bOnlyNumerics = re.ReplaceAllString(string(b), "")

	var runningSum = 0

	for i := 0; i < len(strings.Split(bOnlyNumerics, "\n")); i++ {
		var endOfLine = len(strings.Split((strings.Split(bOnlyNumerics, "\n")[i]), "")) - 1

		var int1, intErr = strconv.Atoi(strings.Split((strings.Split(bOnlyNumerics, "\n")[i]), "")[0] + strings.Split((strings.Split(bOnlyNumerics, "\n")[i]), "")[endOfLine])
		check(intErr)

		runningSum += int1
		p("Line:", i, "Adding: ", int1, "Running sum: ", runningSum)
	}

	// Iterate through line by line O(n) n = number of lines
	// Iterate through character by character O(m) m = characters in line
	// Starting from the beginning and starting from the end O(2m/2) = O(m) for the first and last ints respectively
	// Load into array or sum along the way? probably sum along the way because the individual values aren't useful. Array for debugging

}
