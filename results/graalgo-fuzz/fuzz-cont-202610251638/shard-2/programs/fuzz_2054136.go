package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 int64; F1 int; F2 bool }
func genFn0(p0 int) uint16 {
    v4 := []int{(p0 | (p0 & (p0 << ((991) & 63))))}
    if len(v4) > 0 {
        v4[0] = (p0 ^ p0)
    }
    _ = p0
    _ = v4
    return 255
}

func main() {
    var v1 float64 = (((3.14159 * ((-2.25 * (-2.25 - 3.14159)) + float64(1.5))) * 1.5) / ((1.5 - 0.0) - float64((0.1 + (((1e10 / (3.14159 - 1.5)) * float32(0.0)) + (0.0 * 0.0))))))
    var v2 uint64 = 2
    var v3 uint16 = 1
    v2 = (^(v2))
    switch 4294967295 {
    case -320:
        fmt.Println(v2, v1)
    case 128:
        v2 += (v2 << ((42) & 63))
    case 127:
        switch 334 {
        case -504:
            fmt.Println(v3)
        case 0:
            var v5 int32 = -193
            _ = v5
        case -681:
            v6 := GenS1{F0: 503, F1: 9223372036854775807, F2: (v3 + (^(v3)))}
            _ = v6
        }
    }
    var v7 uint = 1
    var v8 uint16 = (^(v3))
    var v9 uint32 = uint32(v8)
    var v10 float32 = float32(-2.25)
    for fi11 := 0; fi11 < 17; fi11++ {
        switch "a" {
        case "\n\t":
            fmt.Println(v1)
        }
        _ = fi11
    }
    switch 9223372036854775807 {
    case 2147483647:
        v7 = 42
    default:
        if (1 > uint64(v9)) {
            v3 += (v3 ^ uint16(v2))
        }
    }
    if ((^(v2)) >= uint64(v9)) {
        for fi12 := 0; fi12 < 13; fi12++ {
            v3 = genFn0(32767)
            _ = fi12
        }
        var v13 uint64 = v2
        _ = v13
    } else {
        v14 := GenS1{F0: (v9 * 1), F1: -9223372036854775808, F2: genFn0(-2147483648)}
        _ = v14
    }
    for fi15 := 0; fi15 < 4; fi15++ {
        v2 += (^(v2))
        v16 := []int{int(v7), -692, 32768, int(v7)}
        _ = fi15
        _ = v16
    }
    switch 344 {
    case -9223372036854775808:
        var v17 uint32 = 2147483647
        _ = v17
    case 127:
        v9 += (v9 | v9)
    case 0:
        v18 := []int{int(v3)}
        _ = v18
    }
    var v19 int = int(v2)
    var v20 uint32 = uint32(v7)
    for fi21 := 0; fi21 < 4; fi21++ {
        v22 := GenS1{F0: v20, F1: -346, F2: v3}
        v23 := GenS2{F0: 2, F1: 32767}
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 int8 = 127
    v25 := []int{-2147483648, (v19 >> (((v7 << (((v7 ^ v7)) & 63))) & 63)), (v19 * (v19 >> (((v7 / ((65535) | 1))) & 63))), (v19 + int(v7))}
    v26 := GenS3{F0: 1, F1: v19, F2: ((-(v19)) < (v19 ^ (v19 - (-(v19)))))}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
