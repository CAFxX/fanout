package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int }
func genFn0(p0 uint32, p1 uint32) int16 {
    v5 := []int{}
    p0 = 1000
    var v6 uint8 = 255
    var v7 uint64 = 4294967295
    v7 -= (v7 - (v7 | uint64(v6)))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 int32 = 1
    var v2 int = 151
    var v3 bool = true
    var v4 float64 = (-2.25 * -2.25)
    fmt.Println(v4)
    v8 := []int{-255, (v2 + (v2 + (v2 << ((4294967295) & 63))))}
    for fi9 := 0; fi9 < 2; fi9++ {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        var v10 int16 = -130
        for fi11 := 0; fi11 < 19; fi11++ {
            _ = len(v8)
            v12 := GenS1{F0: -9223372036854775808, F1: 1}
            var v13 uint8 = 2
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi9
        _ = v10
    }
    v14 := GenS1{F0: 0, F1: 255}
    var v15 int32 = -704
    var v16 uint16 = 42
    for fi17 := 0; fi17 < 13; fi17++ {
        if len(v8) > 0 {
            v8[0] = v2
        }
        fmt.Println(v3, v15, v3)
        v4 -= v4
        _ = fi17
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        v1 = (-(v1))
        _ = fi18
    }
    v19 := GenS1{F0: 42, F1: 1}
    for fi20 := 0; fi20 < 18; fi20++ {
        var v21 uint64 = 2
        _ = fi20
        _ = v21
    }
    v22 := GenS1{F0: int64(v2), F1: 2}
    switch v2 {
    case -2147483648:
        switch "" {
        case "a":
            fmt.Println(v3, v15, v15)
        case "":
            var v23 uint16 = 1
            _ = v23
        case "\n\t":
            var v24 uint = 2
            _ = v24
        }
    case 32767:
        var v25 int64 = -1
        _ = v25
    }
    v4 = v4
    v16 -= (v16 * (^(v16)))
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v26 := GenS2{F0: 255, F1: (v2 | (v2 % ((v2) | 1)))}
    switch (v2 + (^(v2))) {
    case -906:
        for fi27 := 0; fi27 < 17; fi27++ {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            var v28 int32 = v1
            _ = fi27
            _ = v28
        }
    case 667:
        if len(v8) > 0 {
            v8[0] = (v2 - 234)
        }
    default:
        var v29 uint8 = 2
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v8
}
