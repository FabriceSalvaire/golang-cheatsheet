// DONT'T
// go run: cannot run non-main package
// package foo
package main

import "fmt"

// DONT'T
// runtime.main_main·f: function main is undeclared in the main package
// func foo () {
func main () {
    fmt.Println("run main ...")
}
