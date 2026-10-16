package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 bool; F1 uint8; F2 uint32 }
func genFn0(p0 int, p1 uint32) uint8 {
    var v4 string = "hello"
    v5 := GenS1{F0: 2, F1: 1}
    for fi6 := 0; fi6 < 9; fi6++ {
        v7 := []int{fi6, p0, p0}
        _ = fi6
        _ = v7
    }
    var v8 uint64 = 739
    v9 := []int{}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 bool = false
    var v2 int64 = 758
    v3 := make(map[int]uint32)
    for fi10 := 0; fi10 < 1; fi10++ {
        v2 = (v2 % (((v2 / ((v2) | 1))) | 1))
        switch -2147483648 {
        case 2147483647:
            var v11 int64 = (v2 * v2)
            _ = v11
        case 526:
            fmt.Println(v2, fi10, fi10)
        case -9223372036854775808:
            v3[-9223372036854775808] = 2147483647
        default:
            if v1 {
                fmt.Println(v2, v2, fi10)
                v12 := GenS2{F0: 259}
                v13 := GenS3{F0: 42, F1: 255}
                _ = v12
                _ = v13
            }
        }
        switch "a" {
        case "foo bar":
            var v14 float64 = (1e10 / 1.0)
            _ = v14
        case "":
            var v15 int16 = 42
            _ = v15
        }
        _ = fi10
    }
    var v16 uint32 = 42
    fmt.Println(v2, v16, v16)
    v17 := GenS1{F0: 32767, F1: 42}
    v16 = v16
    v18 := GenS1{F0: 120, F1: v2}
    v19 := GenS4{F0: v1, F1: 0, F2: 65535}
    v2 -= v2
    v1 = v1
    for fi20 := 0; fi20 < 16; fi20++ {
        var v21 uint16 = 65535
        v22 := v3[32768]
        _ = fi20
        _ = v21
        _ = v22
    }
    switch 4294967295 {
    case -9223372036854775808:
        for fi23 := 0; fi23 < 12; fi23++ {
            v2 += (v2 | v2)
            _ = fi23
        }
    case 1:
        for fi24 := 0; fi24 < 3; fi24++ {
            v1 = (791 != v16)
            v25 := []int{(fi24 / ((fi24) | 1)), (fi24 << ((65535) & 63))}
            var v26 uint32 = 835
            _ = fi24
            _ = v25
            _ = v26
        }
    default:
        if (235 > genFn0(32767, 221)) {
            var v27 int = 828
            fmt.Println(v2, v27, v16)
            _ = v27
        }
    }
    if ((0 != -32768) || true) {
        var v28 string = "a"
        v29 := v3[600]
        var v30 float64 = (0.1 + 1.5)
        _ = v28
        _ = v29
        _ = v30
    }
    v2 = v2
    for fi31 := 0; fi31 < 14; fi31++ {
        v16 = v16
        var v32 uint = 2
        if (v1 || v1) {
            v32 = (v32 ^ (v32 + v32))
        }
        _ = fi31
        _ = v32
    }
    fmt.Println(v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
