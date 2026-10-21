package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int; F1 uint64; F2 bool }
type GenS4 struct { F0 uint }
func genFn0(p0 uint16, p1 bool) uint {
    p1 = (false || (65535 != p0))
    p1 = (((3.14159 - 1.5) - ((((1e10 / 1.5) - 0.1) + -2.25) - 1.5)) >= (1.5 - 1e10))
    if false {
        fmt.Println(p1, p1)
        p0 = (p0 >> ((65535) & 63))
        fmt.Println(p0, p0)
    } else {
        var v6 int32 = 59
        _ = v6
    }
    v7 := GenS1{F0: (78 != 111), F1: 127, F2: (("\n\t" + "\n\t") + ("foo bar" + ""))}
    _ = p0
    _ = p1
    _ = v7
    return 739
}
func genFn1(p0 int32, p1 int16, p2 uint64) int8 {
    fmt.Println(p1, p0)
    v8 := GenS2{F0: 2}
    p0 = p0
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 42
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]string)
    var v3 int16 = 1000
    v4 := make(map[uint32]uint32)
    var v5 int64 = -9223372036854775808
    v5 -= (v5 % ((9223372036854775807) | 1))
    v9 := []int{127, 42, 127, 0}
    switch 127 {
    case -2147483648:
        v10 := GenS3{F0: 2147483647, F1: 2, F2: true}
        _ = v10
    case 9223372036854775807:
        var v11 int = 0
        _ = v11
    case 4294967295:
        var v12 float32 = (0.1 - -2.25)
        _ = v12
    default:
        v13 := GenS4{F0: 65535}
        _ = v13
    }
    v5 = v5
    if ((!false) && false) {
        var v14 int32 = 2147483647
        var v15 uint32 = 1
        _ = v14
        _ = v15
    } else {
        v3 = v3
        _ = len(v9)
    }
    v5 += 0
    v3 -= (v3 / ((v3) | 1))
    var v16 int64 = -798
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
