package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }

func main() {
    var v1 uint16 = 65535
    var v2 int8 = 42
    var v3 float32 = float32(3.14159)
    var v4 uint32 = 459
    for fi5 := 0; fi5 < 5; fi5++ {
        if (!true) {
            v6 := []int{(fi5 / ((int(v2)) | 1)), (fi5 << ((uint(v1)) & 63))}
            if true {
                v6 = append(v6, -573)
            }
            v7 := GenS1{F0: 1, F1: fi5}
            _ = v6
            _ = v7
        }
        var v8 float64 = -2.25
        _ = fi5
        _ = v8
    }
    var v9 int16 = 1
    v2 = v2
    fmt.Println(v2, v2)
    var v10 int16 = 42
    v1 -= v1
    var v11 int8 = (-(v2))
    fmt.Println(v1, v3)
    v12 := []int{-920}
    var v13 bool = true
    var v14 uint32 = v4
    v15 := GenS1{F0: -9223372036854775808, F1: 127}
    v12 = append(v12, 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
