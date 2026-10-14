package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int; F1 int8 }
func genFn0(p0 uint8, p1 int32) uint8 {
    v5 := []int{9223372036854775807, 4294967295}
    _ = len(v5)
    if true {
        p1 = (p1 << ((18446744073709551615) & 63))
        if true {
            p1 = (-(p1))
        }
    }
    _ = p0
    _ = p1
    _ = v5
    return (p0 * 42)
}

func main() {
    v1 := []int{}
    var v2 int8 = 42
    var v3 bool = false
    v4 := make(map[int]bool)
    if false {
        v3 = (0 < 1)
    } else {
        v3 = (1 < 255)
        fmt.Println(v3, v2)
    }
    var v6 uint8 = 1
    if (false || v3) {
        var v7 int16 = 32767
        _ = v7
    } else {
        v8 := GenS1{F0: (v2 / ((v2) | 1)), F1: 1, F2: 1}
        v9 := v4[-487]
        v10 := GenS1{F0: v2, F1: -703, F2: 2147483648}
        _ = v10
        _ = v8
        _ = v9
    }
    fmt.Println(v2, v2)
    var v11 uint32 = 29
    _ = len(v1)
    var v12 int32 = 2147483647
    var v13 uint16 = 585
    switch -2147483648 {
    case 1:
        v14 := GenS2{F0: 2147483647, F1: v11}
        _ = v14
    case 9223372036854775807:
        for fi15 := 0; fi15 < 17; fi15++ {
            fmt.Println(v2)
            v4[(-(fi15))] = (true || false)
            var v16 int64 = 1
            _ = fi15
            _ = v16
        }
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    switch 4294967295 {
    case 128:
        var v17 int64 = 154
        _ = v17
    default:
        for fi18 := 0; fi18 < 15; fi18++ {
            var v19 int32 = v12
            for fi20 := 0; fi20 < 16; fi20++ {
                fmt.Println(v19, v11)
                _ = fi20
            }
            _ = fi18
            _ = v19
        }
    }
    if ((v6 & 0) != v6) {
        for fi21 := 0; fi21 < 16; fi21++ {
            v6 += 0
            _ = fi21
        }
        switch -2147483648 {
        case -458:
            v22 := GenS1{F0: (v2 % ((v2) | 1)), F1: 32767, F2: v11}
            _ = v22
        case 676:
            v1 = append(v1, 905)
        case -2147483648:
            v23 := GenS3{F0: 957, F1: (v2 ^ v2)}
            _ = v23
        default:
            var v24 uint32 = uint32(v6)
            _ = v24
        }
    }
    v12 = v12
    v13 += 169
    v25 := GenS1{F0: (v2 ^ 0), F1: 0, F2: (v11 - (^(v11)))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v6
}
