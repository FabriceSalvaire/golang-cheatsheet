//**************************************************************************************************

package main
// a name is exported if it begins with a capital letter

//**************************************************************************************************
//
// import

// import "fmt"
// import "math"
// or
import (
    "fmt"
    "math"
)

//  math.Sqrt(7)

//**************************************************************************************************
//
// Variables

var c, python, java bool

//**************************************************************************************************
//
// Functions

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//**************************************************************************************************
//
// Main

func main() {
    var i int

    fmt.Println("Hello, 世界")

	fmt.Println(add(42, 13))

    a, b := swap("hello", "world")
	fmt.Println(a, b)

    fmt.Println(split(17))
}
