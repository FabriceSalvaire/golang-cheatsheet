//**************************************************************************************************
//
// Go's basic types are
//     bool
//     string
//     int  int8  int16  int32  int64
//     uint uint8 uint16 uint32 uint64 uintptr
//     byte // alias for uint8
//     rune // alias for int32
//          // represents a Unicode code point
//     float32 float64
//     complex64 complex128
//
// The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
//
//
// The zero value is:
//     0 for numeric types,
//     false for the boolean type, and
//     "" (the empty string) for strings.

package main

import (
	"fmt"
    "math"
	"math/cmplx"
)

// global set to zero value
var global1, global2 bool

// global with initializers
var (
	a_bool    bool       = false // true
	MaxInt    uint64     = 1<<64 - 1
    a_float   float64    = 1.23    // 1.2e3
	z         complex128 = cmplx.Sqrt(-5 + 12i)
    a_string  string     = "foo"
    a_byte    byte       = 'f'
)

// more on string
var a_string3 = "foo ' bar"
var a_string4 = "foo ` bar"
var a_string5 = "foo \"bar\" blaz"
// raw string litteral
var a_raw_string = `foo "bar" blaz`
var a_long_string = `...
...
...
`

// Constant
const Pi = 3.14
// Numeric constants are high-precision values.
const (
	Big = 1 << 100
	Small = Big >> 99
)

func need_float(x float64) float64 {
	return x * 0.1
}

func main() {
    fmt.Println("global1, global2:", global1, global2)

	fmt.Printf("Type: %T Value: %v\n", a_bool, a_bool)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Println("Big float:", need_float(Big))

    // Local
    var i int
    fmt.Println("i:", i)

    // Type inference
    // If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
    var global1, global2 = true, "no!"
    fmt.Println("local global1 global2:", global1, global2)

    // Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
    // Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
    k := 3
    fmt.Printf("Type: %T Value: %v\n", k, k)

    an_inferred_bool := 1 < 2    // bool = true
    fmt.Println(an_inferred_bool)

    // Type conversion
    var x, y int = 3, 4
    var f float64 = math.Sqrt(float64(x*x + y*y))
	var u uint = uint(f)
    fmt.Println(x, y, u)

    // Overflow
    // constant 4294967296 overflows uint32
    // var max_uint32 uint32 = 4294967295 + 1
    // fmt.Println(max_uint32)
}
