package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint8 }
func genFn0(p0 uint, p1 uint16, p2 string) int16 {
    switch 665 {
    case 127:
        var v4 bool = true
        _ = v4
    case 752:
        p1 = (p1 & 42)
    }
    v5 := GenS1{F0: 255, F1: (p0 & (p0 >> ((p0) & 63))), F2: uint8(p1)}
    p1 = p1
    p1 = p1
    v6 := []int{696, 2147483647, 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 1000
}
func genFn1(p0 int) int32 {
    p0 = 0
    if (!false) {
        fmt.Println(p0)
        p0 += p0
        fmt.Println(p0)
    } else {
        p0 -= -9223372036854775808
        var v7 int = (p0 ^ p0)
        var v8 int = p0
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        p0 = (fi9 >> ((2) & 63))
        p0 += p0
        _ = fi9
    }
    var v10 float32 = -2.25
    _ = p0
    _ = v10
    return -706
}

func main() {
    v1 := []int{-807}
    var v2 float32 = -2.25
    var v3 uint32 = 2147483648
    var v11 uint16 = 42
    v3 = (v3 % ((v3) | 1))
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 string = "a"
        switch (fi12 % ((fi12) | 1)) {
        case -9223372036854775808:
            var v14 int64 = 0
            _ = v14
        default:
            switch 42 {
            case 127:
                v11 += v11
            case 9223372036854775807:
                for fi15 := 0; fi15 < 7; fi15++ {
                    var v16 string = v13
                    if len(v1) > 0 {
                        v1[0] = fi12
                    }
                    var v17 int16 = 0
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
        }
        v18 := GenS1{F0: 127, F1: 42, F2: 0}
        _ = fi12
        _ = v13
        _ = v18
    }
    var v19 float32 = v2
    for fi20 := 0; fi20 < 8; fi20++ {
        v21 := GenS1{F0: 128, F1: 1, F2: 215}
        _ = fi20
        _ = v21
    }
    var v22 int64 = -9223372036854775808
    if len(v1) > 0 {
        v1[0] = -1
    }
    v11 -= v11
    v3 = (v3 | (v3 & (v3 % ((v3) | 1))))
    v22 = (v22 ^ (v22 + v22))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v22
    _ = v3
}
