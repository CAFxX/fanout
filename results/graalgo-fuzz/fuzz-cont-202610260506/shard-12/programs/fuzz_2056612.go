package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int; F2 int32 }

func main() {
    var v1 uint64 = 471
    v2 := make(map[string]int)
    v3 := make(map[uint32]uint32)
    var v4 uint = 255
    v5 := []int{}
    var v6 float64 = float64(1e10)
    v7 := GenS1{F0: v4}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = len(v3)
    v8 := GenS1{F0: 42}
    var v9 uint64 = (v1 ^ (^(v1)))
    v6 -= ((1.5 - (0.1 / v6)) / 3.14159)
    v1 += 42
    v10 := GenS2{F0: (v4 << (((v4 / (((v4 & (^(v4)))) | 1))) & 63))}
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 uint16 = 310
        v13 := v3[2]
        v14 := GenS3{F0: 65535, F1: fi11, F2: 840}
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    fmt.Println(v1, v6)
    v5 = append(v5, 42)
    for fi15 := 0; fi15 < 1; fi15++ {
        if len(v5) > 0 {
            v5[0] = (fi15 / (((fi15 ^ 128)) | 1))
        }
        fmt.Println(v4)
        v16 := GenS2{F0: 18446744073709551615}
        _ = fi15
        _ = v16
    }
    var v17 int32 = 678
    var v18 uint32 = 859
    var v19 uint = (v4 ^ v4)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
