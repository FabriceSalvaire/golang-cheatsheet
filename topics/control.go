/***************************************************************************************************
 *
 * Copyright (C) 2020 Fabrice Salvaire
 * Licensed under CC BY-NC-SA 4.0
 *
 **************************************************************************************************/

package main

import (
    "fmt"
    "runtime"
    "time"
)

func main() {
    // condional If
    x := 1
    if x > 0 {
    }

    // If with a short statement
    if y := 1; y < 10 {
    }

    if x > 0 {
    } else {
    }

    if x > 0 {
    } else if x < 0 {
    } else {
    }

    // for loop
    sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

    // The init and post statements are optional.
    sum = 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)

    // while
    sum = 1
    for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

    // infinite loop
    // break / continue
    sum = 1
    for {
        fmt.Println(sum)
        if sum > 100 {
            break
        }
        if sum > 50 {
            sum += 20
            continue
        }
        sum += 10
    }

    // Switch
    // Go's switch is like the one in C ..., except that Go only runs the selected case, not all the cases that follow.
    // The break statement is provided automatically in Go.
    // Switch cases need not be constants.
    fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

    // Switch without a condition is the same as switch true.
    // This construct can be a clean way to write long if-then-else chains.
    t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
