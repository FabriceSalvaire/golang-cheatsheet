/***************************************************************************************************
 *
 * Copyright (C) 2020 Fabrice Salvaire
 * Licensed under CC BY-NC-SA 4.0
 *
 **************************************************************************************************/

package main

import "fmt"

func main() {
    {
        type Vertex struct {
            lat, long float64
        }
        var m map[string]Vertex
        fmt.Println(m, m == nil)
        // panic: assignment to entry in nil map
        m = make(map[string]Vertex)

        m["Bell Labs"] = Vertex{
            40.68433, -74.39967,
        }
        fmt.Println(m["Bell Labs"])

        m = map[string]Vertex{
            "Bell Labs": Vertex{
                40.68433, -74.39967,
            },
            "Google": Vertex{
                37.42202, -122.08408,
            },
        }
        fmt.Println(m)

        m = map[string]Vertex{
            "Bell Labs": {40.68433, -74.39967},
            "Google":    {37.42202, -122.08408},
        }
        fmt.Println(m)
    }

    // Mutating Maps
    {
        m := make(map[string]int)

        m["Answer"] = 42
        fmt.Println("The value:", m["Answer"])

        delete(m, "Answer")
        fmt.Println("The value:", m["Answer"])

        v, ok := m["Answer"]
        fmt.Println("The value:", v, "Present?", ok)
    }
}
