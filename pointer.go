package main

import "fmt"

func main() {
	i := 1

    // a pointer to int
    // its zero value is: nil
    var p1 *int

    // Unlike C, Go has no pointer arithmetic

    // point to i
    p1 = &i
	p2 := &i
    // dereferencing / indirecting
	fmt.Println(*p1, *p2)
    // set i through the pointer
	*p2 = 2
	fmt.Println(i)
}
