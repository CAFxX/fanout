package main

import (
    "fmt"
)


func main() {
    var v1 int32 = 1
    v2 := []int{-1, -160}
    var v3 uint16 = 522
    var v4 int32 = v1
    var v5 float64 = 1e10
    v2 = append(v2, 127)
    fmt.Println(v3, v4)
    if true {
        fmt.Println(v5, v1, v5)
        var v6 uint16 = v3
        v1 -= (v4 + (v1 * v4))
        _ = v6
    }
    v3 += 1
    fmt.Println(v4)
    var v7 int = 128
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 uint8 = 255
        v9 += (v9 - (v9 - 2))
        _ = fi8
        _ = v9
    }
    var v10 uint16 = (v3 / ((255) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
