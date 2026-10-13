package main

import (
    "fmt"
)


func main() {
    var v1 int16 = -59
    var v2 uint32 = 2
    var v3 string = "a"
    fmt.Println(v3, v3, v3)
    v2 = 43
    var v4 uint32 = (v2 | v2)
    var v5 float64 = 0.1
    v1 = v1
    for fi6 := 0; fi6 < 13; fi6++ {
        v5 += v5
        v2 = v2
        _ = fi6
    }
    var v7 string = v3
    fmt.Println(v2, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
