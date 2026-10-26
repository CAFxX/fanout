package main

import (
    "fmt"
)


func main() {
    var v1 int16 = 0
    var v2 uint8 = 127
    var v3 float64 = 1e10
    var v4 float64 = ((1e10 / 3.14159) * v3)
    var v5 float32 = 1e10
    v6 := []int{127, int(v1), 4294967295, 127}
    if len(v6) > 0 {
        v6[0] = -71
    }
    var v7 uint32 = 42
    if false {
        v3 -= (v4 / 0.1)
        v2 -= v2
        var v8 float32 = (v5 / (v5 / v5))
        _ = v8
    }
    v7 += (v7 ^ v7)
    fmt.Println(v2, v4, v5)
    var v9 uint64 = 0
    for fi10 := 0; fi10 < 14; fi10++ {
        v3 = v4
        fmt.Println(v5, fi10, fi10)
        _ = fi10
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
