/***************************************************************************************************
 *
 * Copyright (C) 2020 Fabrice Salvaire
 * Licensed under CC BY-NC-SA 4.0
 *
 **************************************************************************************************/

package main

import (
    "fmt"
)

func main() {
    fmt.Println("Please enter your name.")
    var name string
    fmt.Scanln(&name)
    fmt.Printf("Hi, %s! I'm Go!", name)

    fmt.Println("Hello, 世界")
}
