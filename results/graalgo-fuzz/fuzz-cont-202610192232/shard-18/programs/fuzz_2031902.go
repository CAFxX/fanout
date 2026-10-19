package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 int; F1 uint64; F2 int }
func genFn0(p0 int16, p1 int32) uint8 {
    p1 += (p1 % ((p1) | 1))
    for fi4 := 0; fi4 < 8; fi4++ {
        for fi5 := 0; fi5 < 7; fi5++ {
            v6 := GenS1{F0: 4294967295, F1: -795, F2: 3}
            p0 = p0
            var v7 float64 = 1.5
            _ = fi5
            _ = v6
            _ = v7
        }
        for fi8 := 0; fi8 < 2; fi8++ {
            p1 = p1
            p1 -= p1
            _ = fi8
        }
        _ = fi4
    }
    _ = p0
    _ = p1
    return 46
}

func main() {
    var v1 float32 = 0.0
    var v2 string = ""
    v3 := make(map[int]int)
    v1 += (v1 / float32(1e10))
    if true {
        var v9 uint64 = 246
        v1 += v1
        var v10 float64 = 0.1
        _ = v10
        _ = v9
    } else {
        fmt.Println(v2)
    }
    v11 := []int{2147483647}
    v12 := GenS1{F0: 65535, F1: 377, F2: 2}
    var v13 uint64 = 498
    for fi14 := 0; fi14 < 19; fi14++ {
        v2 = v2
        var v15 int8 = 1
        _ = fi14
        _ = v15
    }
    var v16 bool = (float32(-2.25) < v1)
    switch -2147483648 {
    case 127:
        _ = len(v11)
    case -668:
        _ = len(v3)
    case 42:
        var v17 int = 374
        _ = v17
    }
    var v18 uint = 18446744073709551615
    var v19 uint32 = 222
    v16 = v16
    var v20 float64 = 0.1
    switch -737 {
    case -143:
        v1 += 1.5
    case 128:
        v21 := GenS1{F0: (v19 << ((773) & 63)), F1: -9223372036854775808, F2: (v13 & (^(v13)))}
        _ = v21
    case 0:
        switch 128 {
        case 536:
            v16 = v16
        case -1:
            v22 := GenS2{F0: v13, F1: 1}
            _ = v22
        default:
            var v23 int8 = 82
            _ = v23
        }
    }
    v24 := GenS2{F0: v13, F1: 42}
    var v25 float64 = float64(3.14159)
    for fi26 := 0; fi26 < 7; fi26++ {
        v19 += 0
        v27 := GenS3{F0: 32767, F1: v13, F2: (fi26 / (((fi26 / (((fi26 | (^(fi26)))) | 1))) | 1))}
        for fi28 := 0; fi28 < 14; fi28++ {
            _ = len(v3)
            fmt.Println(v13)
            fmt.Println(v18, fi26)
            _ = fi28
        }
        _ = fi26
        _ = v27
    }
    fmt.Println(v1, v20)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v3
}
