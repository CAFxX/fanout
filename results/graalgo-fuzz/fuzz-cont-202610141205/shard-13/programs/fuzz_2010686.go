package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 bool }
func genFn0() uint8 {
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 string = "\n\t"
        _ = fi7
        _ = v8
    }
    v9 := []int{-204}
    var v10 bool = false
    var v11 uint64 = 780
    _ = v10
    _ = v11
    _ = v9
    return 10
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 42
    var v3 float64 = (0.1 + 1.5)
    v4 := make(map[uint32]string)
    var v5 uint16 = v2
    var v6 int16 = 1000
    for fi12 := 0; fi12 < 7; fi12++ {
        fmt.Println(v5)
        v13 := []int{}
        fmt.Println(v5, fi12, v6)
        _ = fi12
        _ = v13
    }
    var v14 float32 = 0.1
    switch 970 {
    case 717:
        v15 := GenS1{F0: 865, F1: v5}
        _ = v15
    case 2147483647:
        var v16 int64 = 42
        _ = v16
    case 94:
        var v17 uint32 = 2147483647
        _ = v17
    }
    v18 := GenS1{F0: 42, F1: 557}
    v19 := GenS1{F0: int64(v2), F1: v5}
    if true {
        v3 = (float64(v14) + ((-2.25 * (3.14159 / v3)) - 3.14159))
        v2 += 449
        var v20 uint8 = 0
        _ = v20
    }
    v21 := GenS1{F0: -571, F1: v2}
    var v22 float32 = 0.0
    v23 := GenS2{F0: -1}
    for fi24 := 0; fi24 < 19; fi24++ {
        switch (fi24 * fi24) {
        case -1:
            v22 -= v14
        case -9223372036854775808:
            v25 := []int{(fi24 * fi24)}
            _ = v25
        }
        _ = fi24
    }
    v26 := GenS3{F0: 1000, F1: false}
    var v27 uint8 = 42
    for fi28 := 0; fi28 < 9; fi28++ {
        var v29 int64 = -1
        _ = fi28
        _ = v29
    }
    if true {
        var v30 int = 127
        _ = v30
    }
    v31 := []int{9223372036854775807, 284}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
