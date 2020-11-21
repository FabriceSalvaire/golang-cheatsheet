/***************************************************************************************************
 *
 * Copyright (C) 2020 Fabrice Salvaire
 * Licensed under CC BY-NC-SA 4.0
 *
 **************************************************************************************************/

package main

import (
    "fmt"
    "strings"
)

func main() {
    {
        // The type [n]T is an array of n values of type T.
        // An array's length is part of its type, so arrays cannot be resized.
        var a [2]string
        a[0] = "Hello"
        a[1] = "World"
        fmt.Println(a[0], a[1])
        fmt.Println(a)

        primes := [6]int{2, 3, 5, 7, 11, 13}
        fmt.Println(primes)

        // Slice literals
        primes2 := []int{2, 3, 5, 7, 11, 13}
        fmt.Println(primes2)

        s := []struct {
            i int
            b bool
        }{
            {2, true},
            {3, false},
            {5, true},
            {7, true},
            {11, false},
            {13, true},
        }
        fmt.Println(s)
    }

    // Slice
    {
        primes := [6]int{2, 3, 5, 7, 11, 13}
        fmt.Println(primes)

        var s []int = primes[1:4]
        fmt.Println(s)   // [3 5 7]

        s = s[:2]
        fmt.Println(s)

        s = s[1:]
        fmt.Println(s)

        names := [4]string{
            "John",
            "Paul",
            "George",
            "Ringo",
        }

        a := names[0:2]
        b := names[1:3]
        b[0] = "XXX"
        fmt.Println(a, b, names)
    }

    // Slice length and capacity
    // A slice has both a length and a capacity.
    // len() : The length of a slice is the number of elements it contains.
    // cap() : The capacity of a slice is the number of elements in the underlying array,
    //         counting from the first element in the slice.
    {
        s := []int{2, 3, 5, 7, 11, 13}
        println(len(s), cap(s))

        // Slice the slice to give it zero length.
        s = s[:0]
        printSlice(s)
        // len=0 cap=6 []

        // Extend its length.
        s = s[:4]
        printSlice(s)
        // len=4 cap=6 [2 3 5 7]
        // but
        // s = s[:10]
        // panic: runtime error: slice bounds out of range [:10] with capacity 4

        // Drop its first two values.
        s = s[2:]
        printSlice(s)
        // len=2 cap=4 [5 7]
    }

    // Nil slice
    {
        var s []int
        fmt.Println(s, len(s), cap(s))
        if s == nil {
            fmt.Println("nil!")
        }
    }

    // Creating a slice with make
    {
        a := make([]int, 5)
        printSlice(a)
        // len=5 cap=5 [0 0 0 0 0]

        b := make([]int, 0, 5)
        printSlice(b)
        // len=0 cap=5 []
    }

    {
        a := [][]string{
            []string{"1", "2", "3"},
            []string{"4", "5", "6"},
            []string{"7", "8", "9"},
        }

        for i := 0; i < len(a); i++ {
            fmt.Printf("%s\n", strings.Join(a[i], " "))
        }
    }

    // Appending to a slice
    {
        var s []int

        s = append(s, 0)
        printSlice(s)

        s = append(s, 1, 2, 3)
        printSlice(s)
    }

    // Range
    {
        var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

        for i, v := range pow {
            fmt.Printf("2**%d = %d\n", i, v)
        }
    }
    {
        // We can skip the index or value by assigning to _.
        // If we only want the index, you can omit the second variable.
        pow := make([]int, 10)
        // for i, _ := range pow {
        for i := range pow {
            pow[i] = 1 << uint(i) // == 2**i
        }
        for _, value := range pow {
            fmt.Printf("%d\n", value)
        }
    }
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
