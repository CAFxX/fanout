package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint32 }
func genFn0(p0 uint) bool {
    var v5 int = -1
    var v6 uint64 = 18446744073709551615
    _ = p0
    _ = v5
    _ = v6
    return true
}
func genFn1(p0 uint8) uint64 {
    p0 += (p0 >> ((2) & 63))
    v7 := []int{1}
    _ = p0
    _ = v7
    return 568
}

func main() {
    var v1 int = 9223372036854775807
    var v2 uint64 = 1
    var v3 int16 = -666
    v4 := make(map[int]bool)
    delete(v4, -85)
    var v8 int64 = -1
    _ = len(v4)
    v9 := []int{(v1 >> ((42) & 63)), 9223372036854775807, (v1 * v1), (v1 ^ 50)}
    for fi10 := 0; fi10 < 7; fi10++ {
        fmt.Println(v8)
        _ = fi10
    }
    if true {
        var v11 int16 = (^(v3))
        switch v1 {
        case 32767:
            v11 = 0
        case -9223372036854775808:
            v12 := GenS1{F0: 2, F1: 2147483648}
            _ = v12
        case 4294967295:
            v13 := GenS2{F0: 128, F1: 576, F2: 4294967295}
            _ = v13
        }
        _ = v11
    } else {
        v1 = v1
        fmt.Println(v2, v1, v1)
    }
    v8 += (v8 * (v8 + v8))
    v8 -= 9223372036854775807
    v2 += v2
    fmt.Println(v8)
    v8 = v8
    var v14 int32 = 2147483647
    var v15 int = v1
    v15 = (v1 * 128)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
