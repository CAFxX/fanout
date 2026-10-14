package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint8 }
func genFn0(p0 int, p1 uint, p2 uint) int16 {
    var v4 uint8 = 128
    v5 := GenS1{F0: 432, F1: 65535}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 0
}

func main() {
    var v1 uint32 = 314
    v2 := make(map[uint32]bool)
    v3 := make(map[string]int)
    switch "foo bar" {
    case "\n\t":
        var v6 uint32 = v1
        _ = v6
    case "":
        var v7 int32 = -190
        _ = v7
    default:
        var v8 uint64 = 876
        _ = v8
    }
    _ = len(v2)
    v1 = (v1 & (v1 ^ v1))
    switch "a" {
    case "":
        v1 += (v1 / (((v1 + 1)) | 1))
    default:
        v1 += 42
    }
    var v9 float32 = 1.5
    v10 := []int{127, 2147483647, 1}
    if true {
        if false {
            v1 += (v1 << ((1) & 63))
        }
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    } else {
        var v11 uint = 18446744073709551615
        _ = v11
    }
    var v12 int32 = 1
    v13 := GenS1{F0: 1, F1: 65535}
    v1 = 42
    if (18446744073709551615 > 0) {
        v14 := GenS2{F0: 128}
        v15 := GenS3{F0: 4294967295, F1: 127, F2: 16}
        v1 -= v1
        _ = v14
        _ = v15
    }
    var v16 int = 127
    if (((("hello" + ("foo bar" + "")) != "") || ((false && (!false)) && (-9 <= -60))) && (true || false)) {
        var v17 float64 = 3.14159
        var v18 uint = 386
        var v19 int32 = -2147483648
        _ = v17
        _ = v18
        _ = v19
    } else {
        v12 -= (v12 & -915)
        switch (v16 / (((v16 * -565)) | 1)) {
        case 0:
            fmt.Println(v16, v9)
        case 569:
            v20 := GenS3{F0: 18446744073709551615, F1: 42, F2: 127}
            _ = v20
        case 76:
            fmt.Println(v1)
        default:
            v3[("foo bar" + "")] = v16
        }
        v16 += v16
    }
    var v21 uint8 = 128
    var v22 uint16 = 649
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v9
}
