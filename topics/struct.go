/***************************************************************************************************
 *
 * Copyright (C) 2020 Fabrice Salvaire
 * Licensed under CC BY-NC-SA 4.0
 *
 **************************************************************************************************/

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
    v1 := Vertex{1, 2}
	v2 := Vertex{X: 1}  // Y:0 is implicit
	v3 := Vertex{}      // X:0 and Y:0
	fmt.Println(v1, v2, v3)

	v1.X = 4
	fmt.Println(v1.X)
	fmt.Println(v1)   // {4 2}

    // pointer to v
    p := &v1   // or := &Vertex(...)
	p.X = 1e9
	fmt.Println(v1)
}
