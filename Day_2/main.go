package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"slices"

	"strconv"
	"strings"
	// TypeOf is in "reflect"
)

var printBool = false

func check(e error) {
	if e != nil {
		print("\n    ****    ****\n\n      * Oops! *\n\n    ****    ****\n")
		panic(e)
	}
}

var print = fmt.Println
var p = print

func printType(input any) {
	print(reflect.TypeOf(input))
}

var red int
var green int
var blue int

func main() {

	file, err := os.Open("day2.txt")
	check(err)

	buffer := bufio.NewScanner(file)

	var reColon, reColonErr = regexp.Compile(":")
	check(reColonErr)

	var reSemiColon, reSemiColonErr = regexp.Compile(";")
	check(reSemiColonErr)

	var reRed, reRedErr = regexp.Compile(`(\d+) red`)
	check(reRedErr)
	var reGreen, reGreenErr = regexp.Compile(`(\d+) green`)
	check(reGreenErr)
	var reBlue, reBlueErr = regexp.Compile(`(\d+) blue`)
	check(reBlueErr)

	counter := 1
	var possibleGamesSum int
	var r []int
	var g []int
	var b []int

	for buffer.Scan() {
		// Separate the "Game X:" from the string returns an array of a single string
		games := reColon.Split(buffer.Text(), -1)[1]
		if printBool {
			p("Games:", games)
		}
		gameNumber := strings.Split(reColon.Split(buffer.Text(), -1)[0], " ")[1]
		// Separate games into an array of strings, using semicolon as the delimiter
		gamesByLine := reSemiColon.Split(games, -1)
		if printBool {
			p("Games By Line:", gamesByLine)
		}
		// test := reRed.FindAllStringSubmatch(gamesByLine[0], -1)
		// print(test)

		if reRed.FindAllStringSubmatch(games, -1) != nil {
			reds := reRed.FindAllStringSubmatch(games, -1)
			for i, red := range reds {
				if printBool {
					print(i, "Reds: ", red[1])
				}
				strR, err := strconv.Atoi(red[1])
				check(err)
				r = append(r, strR)
			}
			if printBool {
				print(r)
			}
		}
		if reGreen.FindAllStringSubmatch(games, -1) != nil {
			greens := reGreen.FindAllStringSubmatch(games, -1)
			for i, green := range greens {
				if printBool {
					print(i, "Greens: ", green[1])
				}
				strG, err := strconv.Atoi(green[1])
				check(err)
				g = append(g, strG)
			}
			if printBool {
				print(g)
			}
		}
		if reBlue.FindAllStringSubmatch(games, -1) != nil {
			blues := reBlue.FindAllStringSubmatch(games, -1)
			for i, blue := range blues {
				if printBool {
					print(i, "Blues: ", blue[1])
				}
				strB, err := strconv.Atoi(blue[1])
				check(err)
				b = append(b, strB)
			}
			if printBool {
				print(b)
			}
		}
		p(" ", gameNumber)

		// for i, ValuesOfTheGameAsString := range gamesByLine {
		if slices.Max(r) <= 12 && slices.Max(g) <= 13 && slices.Max(b) <= 14 {
			// p(i, ValuesOfTheGameAsString)
			var gameNumInt, err = strconv.Atoi(gameNumber)
			check(err)
			possibleGamesSum += gameNumInt
			p(possibleGamesSum)
		} else {
			// print(gameNumber, "!")
		}
		// }
		counter = counter + 1
		r = nil
		g = nil
		b = nil
	}
	if possibleGamesSum == 28 {
		p("Done!")
	}
	print(possibleGamesSum)
}
