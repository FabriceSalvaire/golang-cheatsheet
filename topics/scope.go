/***************************************************************************************************
 *
 * Copyright (C) 2020 Fabrice Salvaire
 * Licensed under CC BY-NC-SA 4.0
 *
 **************************************************************************************************/

package main

import "fmt"

var v int

func main() {
    var v int
    fmt.Println(v)
    {
        var v int
        fmt.Println(v)
        {
            var v int
            fmt.Println(v)
        }
    }
}
