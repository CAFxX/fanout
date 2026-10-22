package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 int; F2 uint8 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 int64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 984
    var v3 int64 = -9223372036854775808
    if ("" > "hello") {
        if true {
            v2 += (^(v2))
            v3 = 9223372036854775807
        }
    } else {
        if false {
            fmt.Println(v3)
            v4 := GenS1{F0: (v3 << ((967) & 63))}
            v5 := GenS2{F0: 0}
            _ = v4
            _ = v5
        } else {
            fmt.Println(v3, v2)
        }
        v6 := []int{128, 32767}
        _ = v6
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int = (-(fi7))
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := GenS1{F0: -448}
        v3 = -1
        _ = fi9
        _ = v10
    }
    v2 = (v2 + 999)
    for fi11 := 0; fi11 < 4; fi11++ {
        v12 := GenS3{F0: (-(fi11)), F1: fi11, F2: 127}
        _ = fi11
        _ = v12
    }
    var v13 int16 = 32767
    v3 += (v3 >> ((0) & 63))
    if ((0.0 + 0.1) < 1.5) {
        v2 = v2
        var v14 int64 = v3
        v13 += (v13 / (((v13 / ((v13) | 1))) | 1))
        _ = v14
    }
    switch 32768 {
    case 264:
        v15 := GenS1{F0: v3}
        _ = v15
    default:
        v13 = (v13 & (-(v13)))
    }
    var v16 uint32 = 65535
    var v17 uint = 0
    v18 := []int{127, -2147483648, 2147483647, -706}
    v19 := GenS4{F0: -1, F1: -692}
    _ = len(v1)
    v20 := GenS5{F0: -1}
    v21 := v1[-1]
    v21 -= v21
    v13 -= (v13 << (((v17 & v17)) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
