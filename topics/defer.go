/***************************************************************************************************
 *
 * Copyright (C) 2020 Fabrice Salvaire
 * Licensed under CC BY-NC-SA 4.0
 *
 **************************************************************************************************/

package main

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")

	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)   // executed in lifo order
	}
	fmt.Println("done")
}
