package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 int32 = -1
    v2 := make(map[int]int)
    var v3 uint64 = 18446744073709551615
    var v4 int = 846
    switch v4 {
    case -870:
        v5 := GenS1{F0: "", F1: (42 > (v1 - 979))}
        _ = v5
    }
    v6 := GenS2{F0: 482}
    v7 := GenS2{F0: -250}
    v3 -= (v3 % (((v3 / ((v3) | 1))) | 1))
    for fi8 := 0; fi8 < 4; fi8++ {
        fmt.Println(fi8, v3, v1)
        var v9 bool = false
        _ = fi8
        _ = v9
    }
    v1 -= (v1 / ((42) | 1))
    v10 := GenS3{F0: 239}
    fmt.Println(v1)
    if (!false) {
        var v11 bool = (!(!false))
        _ = v11
    }
    fmt.Println(v3, v3)
    v12 := GenS3{F0: 2}
    var v13 int8 = -21
    var v14 float32 = float32(float64(3.14159))
    var v15 uint = 65535
    if false {
        if (0.0 > -2.25) {
            var v16 uint32 = 651
            v13 = 42
            fmt.Println(v4, v14)
            _ = v16
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
