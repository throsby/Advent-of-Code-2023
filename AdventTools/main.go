package adventtools

import (
	"fmt"
	"reflect"
)

func Check(e error) {
	if e != nil {
		print("\n    ****    ****\n\n      * Oops! *\n\n    ****    ****\n")
		panic(e)
	}
}

var Print = fmt.Println
var P = Print

func PrintType(input any) {
	fmt.Println(reflect.TypeOf(input))
}
