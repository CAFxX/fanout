package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint64 }

func main() {
    var v1 int = -9223372036854775808
    var v2 string = ""
    var v3 float64 = (3.14159 - float64(0.0))
    var v4 uint64 = 42
    v5 := make(map[uint32]uint32)
    var v6 int64 = 11
    var v7 uint32 = 4294967295
    v5[(v7 + v7)] = (^(v7))
    v6 = v6
    v7 = 851
    fmt.Println(v1, v7, v7)
    v8 := v5[2147483648]
    var v9 int8 = -108
    v10 := GenS1{F0: 1, F1: v1, F2: v4}
    if false {
        for fi11 := 0; fi11 < 6; fi11++ {
            v12 := GenS1{F0: 2147483647, F1: (fi11 + int(v8)), F2: (v4 * v4)}
            _ = fi11
            _ = v12
        }
    }
    v13 := GenS1{F0: -1, F1: v1, F2: v4}
    v14 := []int{(v1 >> ((1) & 63)), v1}
    v9 += -128
    var v15 int64 = v6
    v16 := GenS1{F0: -2147483648, F1: v1, F2: (v4 * (v4 + 2))}
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
