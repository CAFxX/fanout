package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int16 }

func main() {
    var v1 bool = (true || ((242 != 255) && (!(-1 == -389))))
    var v2 uint16 = 318
    v3 := make(map[uint32]int)
    v4 := []int{32767, 32767, 2147483647}
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 int = -1
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 127, F1: 224, F2: 0}
    var v8 int32 = -2147483648
    v2 = v2
    var v9 int8 = 1
    var v10 uint64 = 42
    if v1 {
        var v11 int32 = -2147483648
        var v12 uint16 = 2
        v1 = v1
        _ = v11
        _ = v12
    } else {
        fmt.Println(v10)
        var v13 float32 = (0.0 / (((-2.25 - 0.1) + (float32(3.14159) * 1e10)) * 1e10))
        fmt.Println(v2, v13)
        _ = v13
    }
    if false {
        v14 := v3[2147483648]
        _ = v14
    }
    _ = len(v3)
    if false {
        v9 -= (v9 | 0)
    }
    delete(v3, 886)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
