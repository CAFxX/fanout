package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
func genFn0(p0 string, p1 int16) int64 {
    v5 := []int{440}
    var v6 uint32 = 0
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 uint64 = 42
    var v2 int64 = -995
    var v3 uint = 18446744073709551615
    v4 := make(map[uint32]uint32)
    v3 -= v3
    _ = len(v4)
    v1 = 2
    v3 -= 65535
    switch 2147483647 {
    case -2147483648:
        v1 -= v1
    case 128:
        v3 += v3
    case 0:
        fmt.Println(v2)
    default:
        var v7 int = -9223372036854775808
        _ = v7
    }
    v8 := []int{9223372036854775807}
    for fi9 := 0; fi9 < 16; fi9++ {
        var v10 int32 = 1
        v11 := GenS1{F0: 42, F1: 18446744073709551615}
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int64 = v2
    fmt.Println(v1, v3, v1)
    var v13 uint = 897
    v3 = (v13 / (((v13 - (v13 - 65535))) | 1))
    var v14 bool = true
    var v15 uint64 = (v1 - v1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
