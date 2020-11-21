/***************************************************************************************************
 *
 * Copyright (C) 2020 Fabrice Salvaire
 * Licensed under CC BY-NC-SA 4.0
 *
 **************************************************************************************************/

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
