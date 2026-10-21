package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int8 }

func main() {
    v1 := []string{}
    var v2 float64 = 1e10
    var v3 bool = ((!false) && (false && true))
    var v4 int16 = -1
    var v5 int = 32768
    v6 := make(map[int]int)
    var v7 float32 = 3.14159
    if (((v7 * (3.14159 / 0.1)) / v7) >= ((v7 / 1.0) * v7)) {
        var v8 int32 = 0
        _ = v8
    }
    var v9 int = 48
    v10 := v6[-278]
    v6[(v10 * (v10 ^ (v5 % (((v5 & v5)) | 1))))] = v10
    var v11 int16 = (^(v4))
    v12 := GenS1{F0: 255, F1: v9, F2: 42}
    v11 += v11
    v9 = (v5 & (^(v9)))
    fmt.Println(v4)
    var v13 int16 = v11
    v14 := GenS1{F0: 42, F1: (v9 ^ v10), F2: -128}
    if v3 {
        for fi15 := 0; fi15 < 4; fi15++ {
            var v16 int32 = -70
            _ = fi15
            _ = v16
        }
        var v17 uint = 642
        v9 = v10
        _ = v17
    } else {
        var v18 uint32 = 1
        _ = v18
    }
    fmt.Println(v5, v3, v13)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
