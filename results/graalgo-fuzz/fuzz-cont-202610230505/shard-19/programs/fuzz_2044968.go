package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint32 }
func genFn0(p0 int16, p1 uint8) uint {
    v6 := []int{32768}
    var v7 uint64 = 299
    p1 = p1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 string, p1 uint8) string {
    switch 2147483647 {
    case -923:
        if (-1 == -767) {
            fmt.Println(p1, p0)
        }
    case 1:
        p1 = 255
    case -1:
        v8 := GenS1{F0: (p1 / (((^(p1))) | 1)), F1: -692}
        _ = v8
    }
    switch -9223372036854775808 {
    case 2147483647:
        p1 = 0
    case 42:
        p0 = ("hello" + "foo bar")
    case -633:
        for fi9 := 0; fi9 < 8; fi9++ {
            p1 -= p1
            var v10 int64 = 876
            _ = fi9
            _ = v10
        }
    default:
        p0 = "\n\t"
    }
    var v11 uint8 = 42
    var v12 int16 = 1
    for fi13 := 0; fi13 < 7; fi13++ {
        v14 := []int{(-(fi13))}
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return p0
}

func main() {
    v1 := []int{434, -9223372036854775808, -1, 4294967295}
    var v2 int64 = -1
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]bool)
    v5 := make(map[string]string)
    v15 := GenS2{F0: v2, F1: 94}
    v16 := GenS3{F0: 545, F1: -1, F2: 182}
    if len(v1) > 0 {
        v1[0] = 42
    }
    var v17 uint32 = 42
    v17 -= 1
    v2 += (v2 ^ (v2 | (v2 << ((genFn0(0, 127)) & 63))))
    var v18 uint = 810
    var v19 uint32 = v17
    fmt.Println(v17, v17, v2)
    v2 -= -1
    fmt.Println(v19, v2)
    var v20 uint8 = 42
    _ = len(v1)
    var v21 uint = v18
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
