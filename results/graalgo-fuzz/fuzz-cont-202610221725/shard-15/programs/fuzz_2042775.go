package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 uint64 = 4294967295
    var v2 int = 4294967295
    v3 := make(map[int]bool)
    var v4 float32 = 0.0
    if (float64(0.0) == (-2.25 + -2.25)) {
        var v5 int32 = -2147483648
        v6 := []int{128, (v2 / ((32767) | 1)), 42, v2}
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 float32 = float32(0.0)
        v9 := []int{v2}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint32 = 2147483648
    v1 = v1
    v4 -= 0.1
    v1 -= v1
    fmt.Println(v4, v1)
    var v11 int32 = -1
    if true {
        v12 := GenS1{F0: 1}
        fmt.Println(v11)
        var v13 uint8 = 1
        _ = v12
        _ = v13
    } else {
        v14 := GenS1{F0: 1}
        var v15 uint64 = v1
        _ = v14
        _ = v15
    }
    v16 := []int{(v2 / (((v2 & (v2 | v2))) | 1)), -1, -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
