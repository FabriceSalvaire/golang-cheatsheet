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
