package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int64, p1 uint16) uint8 {
    v4 := []int{-2147483648}
    var v5 int64 = 0
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 42
}
func genFn1(p0 bool, p1 uint8) int32 {
    var v6 uint64 = 42
    switch -1 {
    case -317:
        v7 := []int{}
        _ = v7
    case 42:
        if p0 {
            var v8 int16 = 0
            _ = v8
        }
    case 259:
        fmt.Println(p0, v6)
    default:
        p0 = p0
    }
    p1 -= (^(p1))
    _ = p0
    _ = p1
    _ = v6
    return -844
}

func main() {
    var v1 float64 = float64((float32((1e10 * 0.1)) / 1e10))
    v2 := []int{1, -616}
    var v3 uint16 = 240
    fmt.Println(v3)
    for fi9 := 0; fi9 < 17; fi9++ {
        if false {
            var v10 int16 = 0
            fmt.Println(v3, v3, fi9)
            _ = v10
        }
        _ = fi9
    }
    fmt.Println(v1, v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v11 int8 = 42
    if (!false) {
        _ = len(v2)
    } else {
        _ = len(v2)
    }
    v12 := GenS1{F0: 1}
    var v13 string = "a"
    v3 += 1
    var v14 uint16 = (v3 / (((v3 * (^(v3)))) | 1))
    for fi15 := 0; fi15 < 5; fi15++ {
        v16 := GenS1{F0: -593}
        var v17 int64 = 1
        _ = fi15
        _ = v16
        _ = v17
    }
    v14 = (^(v3))
    v3 += (^(v14))
    var v18 uint8 = 202
    var v19 uint8 = (v18 << ((65535) & 63))
    v11 = v11
    v2 = append(v2, 32768)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
