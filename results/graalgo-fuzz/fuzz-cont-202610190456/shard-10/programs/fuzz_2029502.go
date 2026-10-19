package main

import (
    "fmt"
)


func main() {
    var v1 uint32 = 2147483647
    var v2 uint = 18446744073709551615
    var v3 string = "a"
    var v4 int8 = 28
    v5 := make(map[uint32]int)
    v6 := []int{32768, 32768}
    fmt.Println(v2, v3, v4)
    v4 = 0
    fmt.Println(v2, v4, v2)
    fmt.Println(v2)
    v3 = ""
    for fi7 := 0; fi7 < 6; fi7++ {
        fmt.Println(fi7, fi7)
        v4 -= int8(v1)
        _ = fi7
    }
    v2 -= (^(v2))
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 float32 = 3.14159
        var v10 float64 = float64((v9 / v9))
        _ = fi8
        _ = v10
        _ = v9
    }
    v4 = v4
    v3 = v3
    var v11 float64 = -2.25
    v12 := v5[(v1 / ((v1) | 1))]
    var v13 int64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
