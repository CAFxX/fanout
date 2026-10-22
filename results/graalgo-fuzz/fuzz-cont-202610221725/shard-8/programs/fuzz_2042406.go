package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 bool; F2 string }
type GenS4 struct { F0 uint16; F1 string; F2 uint32 }
func genFn0() int32 {
    var v7 uint16 = 2
    v7 = (v7 - (^(v7)))
    var v8 uint32 = 2
    if true {
        for fi9 := 0; fi9 < 4; fi9++ {
            for fi10 := 0; fi10 < 4; fi10++ {
                v11 := GenS1{F0: 464, F1: 0, F2: -128}
                v7 += v7
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
        fmt.Println(v8, v8)
        for fi12 := 0; fi12 < 16; fi12++ {
            v8 = v8
            var v13 uint64 = 4294967295
            var v14 int32 = 42
            _ = fi12
            _ = v13
            _ = v14
        }
    } else {
        var v15 int = -2147483648
        fmt.Println(v8, v8)
        var v16 uint16 = 811
        _ = v15
        _ = v16
    }
    v17 := []int{}
    _ = v17
    _ = v7
    _ = v8
    return -2147483648
}
func genFn1() int16 {
    var v18 float64 = (((1e10 * 1.5) * 1.5) * 0.0)
    v19 := GenS2{F0: (1 <= 42)}
    _ = v18
    _ = v19
    return 1000
}

func main() {
    var v1 uint32 = 316
    v2 := make(map[string]string)
    var v3 bool = false
    v4 := make(map[int]uint32)
    var v5 string = "\n\t"
    var v6 int32 = 2147483647
    v5 = v5
    for fi20 := 0; fi20 < 4; fi20++ {
        v21 := GenS1{F0: v1, F1: 742, F2: 77}
        v22 := GenS3{F0: 18446744073709551615, F1: ((fi20 + (fi20 << ((473) & 63))) > fi20), F2: v5}
        _ = fi20
        _ = v21
        _ = v22
    }
    v23 := GenS1{F0: v1, F1: 42, F2: 42}
    var v24 int16 = -614
    fmt.Println(v6)
    fmt.Println(v24)
    v1 = 447
    var v25 uint64 = 2
    switch 215 {
    case 2147483647:
        v26 := GenS4{F0: 2, F1: ("hello" + ""), F2: v1}
        _ = v26
    case 128:
        switch 128 {
        case -880:
            v25 += (v25 | v25)
        case 312:
            var v27 int = 892
            _ = v27
        case 128:
            var v28 string = v5
            _ = v28
        default:
            v29 := GenS1{F0: uint32(v25), F1: 0, F2: 1}
            _ = v29
        }
    }
    var v30 int16 = v24
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
