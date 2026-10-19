package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint }

func main() {
    v1 := make(map[string]int)
    var v2 int16 = -902
    var v3 int32 = -1
    var v4 int32 = (^(v3))
    var v5 bool = false
    var v6 uint32 = 4294967295
    var v7 float64 = (1.5 - (float64(0.0) / -2.25))
    var v8 uint = 2
    v3 += (v4 + (v3 & (^(v4))))
    var v9 int = 32768
    if (false && false) {
        v10 := GenS1{F0: v8, F1: 42, F2: v8}
        v3 += -410
        _ = v10
    } else {
        var v11 float64 = v7
        _ = v11
    }
    var v12 uint64 = 4294967295
    var v13 int16 = 32767
    _ = len(v1)
    v14 := []int{v9, (v9 % ((v9) | 1)), (v9 - -1), -178}
    var v15 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
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
