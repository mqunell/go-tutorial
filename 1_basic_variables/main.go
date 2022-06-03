// Go apps are structured into packages - this decales which package the file is part of
package main

// Import libraries - "fmt" is used to format strings
import (
	"fmt"
	"strconv"
)

// Entry point into application
func main() {
	// declarationInitialization()
	// conversions()
	// booleans()
	// strings()
	constants()
}

func declarationInitialization() {
	// Separate declaration/initialization
	var a int
	a = 28

	// Explicit dec/init
	var b int = 29

	// Implicit dec/init (note the lack of "var")
	c := 30

	fmt.Println(a, b, c)
}

func conversions() {	
	var a int = 38							// int 38
	var b float32 = float32(a)			// float32 38
	var c string = string(a)			// string "&", 38 in UTF-8
	var d string = strconv.Itoa(a)	// string "38" 
	var e int = int('m')					// int 109
	
	// %T = type, %v = value
	fmt.Printf("%T %v\n", a, a)
	fmt.Printf("%T %v\n", b, b)
	fmt.Printf("%T %v\n", c, c)
	fmt.Printf("%T %v\n", d, d)
	fmt.Printf("%T %v\n", e, e)
}

func booleans() {
	var a bool = true	// true
	b := (1 == 2)		// false
	var c bool			// false (uninitialized defaults to false)

	fmt.Println(a, b, c)
}

func strings() {
	// Use array syntax with a string
	str := "This is a string"
	fmt.Printf("%T %v\n", str[2], str[2]) // uint8 105 (ASCII for "i")
	fmt.Printf("%T %v\n", string(str[2]), string(str[2])) // string "i"

	// str[2] = "u" --> error; wrong type and strings are immutable

	// Concatenation
	concat := "New " + "string"
	fmt.Println(concat)  // string "New string"

	// Byte array conversion
	bytes := []byte(str)
	fmt.Printf("%T %v\n", bytes, bytes) // []uint8 [84 104 105 115 ...]
}

func constants() {
	const myConst int = 4 // int 42
	// myConst = 5 --> error; cannot assign to a constant
	// const myConstTwo float64 = math.Sin(3.14) --> error; can't set constant to something that has to be determined at runtime

	const (
		a = iota // 0
		b // 1 - don't need to continue typing "= iota"
		c // 2
	)

	const (
		a2 = iota // 0 - iota values are scoped to const blocks
	)

	fmt.Println(myConst, a, b, c, a2)
}