package main

import (
    "fmt"
)


func main() {
    v1 := []int{391, 1, -1, 42}
    v2 := make(map[int]int)
    var v3 int8 = 0
    var v4 uint16 = 657
    v3 = v3
    var v5 float64 = 1e10
    fmt.Println(v5, v5)
    if false {
        v3 -= (v3 * 74)
    }
    for fi6 := 0; fi6 < 17; fi6++ {
        v5 -= v5
        _ = fi6
    }
    v3 = v3
    v5 += (3.14159 * v5)
    var v7 int32 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
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
