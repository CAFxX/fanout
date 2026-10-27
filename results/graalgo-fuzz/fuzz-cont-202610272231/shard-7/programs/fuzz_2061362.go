package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint32, p1 int16, p2 bool) int32 {
    var v7 int32 = -1
    p0 -= (p0 | 42)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 560
}
func genFn1(p0 int64, p1 uint, p2 int16) int {
    v8 := GenS1{F0: 27, F1: 42}
    v9 := GenS1{F0: 255, F1: 61}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    v1 := []uint32{1}
    v2 := make(map[uint32]bool)
    v3 := make(map[int]string)
    var v4 string = "\n\t"
    var v5 int32 = 42
    v6 := make(map[uint32]uint32)
    _ = len(v1)
    var v10 uint = 371
    v11 := GenS1{F0: uint8(v10), F1: -128}
    var v12 int64 = 42
    v5 = (v5 & genFn0(65535, 32767, (false || (2147483647 <= 65535))))
    switch "a" {
    case "\n\t":
        _ = len(v6)
    case "foo bar":
        delete(v6, 601)
    default:
        v13 := GenS2{F0: 32767}
        _ = v13
    }
    if true {
        var v14 int16 = -407
        v10 -= (v10 - (^(v10)))
        v4 = "a"
        _ = v14
    }
    v3[32767] = v4
    v10 += 4294967295
    for fi15 := 0; fi15 < 3; fi15++ {
        fmt.Println(v4)
        var v16 uint = (v10 - v10)
        switch (v4 + v4) {
        case "\n\t":
            var v17 uint8 = 79
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    var v18 uint64 = 129
    v3[-1] = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
