package main

import (
    "fmt"
)


func main() {
    var v1 int8 = -128
    var v2 int8 = 100
    var v3 int32 = 0
    v4 := make(map[string]uint32)
    v2 -= v1
    var v5 int = 1
    v5 -= v5
    v1 += v2
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := []int{(fi6 / ((410) | 1)), (v5 * 2147483647), 42}
        _ = fi6
        _ = v7
    }
    v1 = (v1 + (v2 % ((v1) | 1)))
    v8 := []int{(v5 | v5), (^(v5)), 25}
    var v9 int8 = (v1 - 1)
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v2 += 86
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
