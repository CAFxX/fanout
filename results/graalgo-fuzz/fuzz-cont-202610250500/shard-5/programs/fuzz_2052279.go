package main

import (
    "fmt"
)


func main() {
    var v1 float32 = (1.5 - 3.14159)
    v2 := make(map[int]string)
    var v3 float32 = v1
    v4 := make(map[int]bool)
    var v5 int = 379
    var v6 bool = true
    v7 := []int{(-(v5)), v5}
    var v8 bool = (42 == 255)
    v5 = 2147483647
    var v9 uint64 = 18446744073709551615
    for fi10 := 0; fi10 < 15; fi10++ {
        switch 32767 {
        case 42:
            v11 := v4[-2147483648]
            _ = v11
        case 608:
            v8 = true
        }
        v3 = (v1 + v3)
        _ = fi10
    }
    fmt.Println(v6, v6)
    v5 -= v5
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 float64 = -2.25
        v9 += v9
        _ = fi12
        _ = v13
    }
    var v14 int64 = 1
    var v15 bool = (!(255 != 38))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
