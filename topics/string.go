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
    var a_string string = "foo bar"
    fmt.Println(a_string, len(a_string))

    var another_string string

    // another_string = a_string.FUNCTION()    // DON'T
    another_string = strings.ToUpper(a_string)
    another_string = strings.ToLower(a_string)
    strings.HasSuffix(a_string, "...")
    strings.HasPrefix(a_string, "...")
    strings.Contains(a_string, "...")
    strings.Count(a_string, "...")
    another_string = strings.TrimSpace(a_string)
    fmt.Println(another_string)
}
