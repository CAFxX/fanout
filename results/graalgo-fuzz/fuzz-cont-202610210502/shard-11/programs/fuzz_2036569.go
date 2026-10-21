package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool }

func main() {
    v1 := []int{2147483647, 152, 2147483647}
    var v2 uint8 = 127
    v3 := make(map[int]int)
    var v4 float64 = ((3.14159 / 1e10) / ((1.5 / 1.0) + -2.25))
    fmt.Println(v4, v4)
    if (v4 >= float64(0.1)) {
        v2 = (^(v2))
    }
    var v5 int32 = 499
    v6 := GenS1{F0: ""}
    v7 := v3[2147483647]
    _ = len(v1)
    if true {
        var v8 uint64 = 23
        _ = v8
    } else {
        v5 -= 0
    }
    v9 := GenS2{F0: (700 != -1)}
    var v10 uint8 = v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
