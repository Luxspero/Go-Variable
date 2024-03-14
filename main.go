package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var name type = expression
	var integer int32 = 1
	var float float32 = 0.5
	var boolean bool = false
	var string string = "Example String"

	//short declaration: VariableName := expression
	shortInt := 100
	shortFloat := 0.035
	shortBoolean := true
	shortStr := "Example Short String"

	fmt.Println(integer)
	fmt.Println(float)
	fmt.Println(boolean)
	fmt.Println(string)
	fmt.Println("")

	//Short Declaration Print
	fmt.Println("Short Declaration Print")
	fmt.Println(shortInt)
	fmt.Println(shortFloat)
	fmt.Println(shortBoolean)
	fmt.Println(shortStr)

	// *** POINTERS ****
	x := 1
	p := &x

	fmt.Println("")
	fmt.Println("POINTERS")
	fmt.Println(p)
	fmt.Println(*p)

	//*** TYPE ****

	//fahrenheit & celsius
	type fahrenheit int32
	type celsius int64

	var f fahrenheit = 60
	var c celsius = 0

	c = celsius((f - 32) * 5 / 9)

	fmt.Println("Fahrenheit: ", f, " Celcius= ", c)
	//print type data F and C
	fmt.Println("f = ", reflect.TypeOf(f))
	fmt.Println("c = ", reflect.TypeOf(c))
}
