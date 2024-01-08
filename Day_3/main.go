package main

// var summaryOfValues = "788+54+555+270+893+963+860+53+52+347+428+522+41+481+462+187+678+420+115+707+562+438+877+199+145+71+68+38+300=11574"
var test = 11574
var sumOfEngineParts = 0

func main() {

	// file, err := os.Open("day3subset.txt")
	// adventtools.Check(err)

	grid := [][]string{
		{"2", "0", "."},
		{"/", ".", "."},
		{".", "*", "9"},
	}

	gridTest := int(29)

	for y_val, interf := range grid {
		for x_val, value := range interf {
			print("(", x_val, ",", y_val, ") \"", value, "\" ")
			// Code here

		}
		print("\n")
	}

	print(grid[0][0], "\n")
	print(gridTest)

	// buffer := bufio.NewScanner(file)
	// wholeFile := ""
	// y_pos := 1
	// for buffer.Scan() {
	// 	wholeFile = wholeFile + buffer.Text() + "\n"
	// 	for x_pos, rune := range buffer.Text() {
	// 		adventtools.Print(x_pos, ",", y_pos, string(rune))
	// 	}
	// 	y_pos++
	// }
	// print(wholeFile)
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	adventtools.Print(scanner.Text())
	// }
	if sumOfEngineParts == test {
		print("Done!")
	}
}

// Collect the xy locations of symbols
//
