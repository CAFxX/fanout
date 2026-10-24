package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 string; F1 int; F2 int64 }

func main() {
    var v1 uint = 694
    v2 := make(map[uint32]int)
    var v3 bool = true
    var v4 uint = 65535
    var v5 uint64 = 389
    switch 42 {
    case 606:
        var v6 uint32 = 406
        _ = v6
    }
    v3 = ((0.0 / 0.1) > (3.14159 * 0.1))
    v7 := GenS1{F0: 206, F1: 520}
    var v8 uint64 = (^(v5))
    v5 += (v8 ^ 1)
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 uint32 = 42
        v4 = (v4 % (((v1 / ((0) | 1))) | 1))
        var v11 float64 = (float64(float32(0.0)) - 1e10)
        _ = fi9
        _ = v10
        _ = v11
    }
    v5 = 42
    v12 := GenS1{F0: 42, F1: 357}
    switch -179 {
    case 9223372036854775807:
        var v13 uint32 = 42
        _ = v13
    default:
        v14 := GenS2{F0: "", F1: 42, F2: 768}
        _ = v14
    }
    for fi15 := 0; fi15 < 7; fi15++ {
        v16 := []int{1}
        var v17 int8 = 100
        _ = fi15
        _ = v16
        _ = v17
    }
    fmt.Println(v1, v8)
    var v18 float32 = (((0.1 + 1.5) * (((-2.25 - (1.5 / (3.14159 / (0.1 + 3.14159)))) / 3.14159) + 3.14159)) + float32(1.5))
    var v19 uint16 = 42
    var v20 int = 577
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
