package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int64 }
func genFn0(p0 uint) uint16 {
    v5 := []int{-639, 42, 1, 351}
    p0 = (p0 << ((p0) & 63))
    fmt.Println(p0)
    _ = p0
    _ = v5
    return 42
}
func genFn1(p0 string, p1 bool) uint16 {
    switch "a" {
    case "":
        for fi6 := 0; fi6 < 11; fi6++ {
            var v7 uint = 2
            _ = fi6
            _ = v7
        }
    default:
        p0 = "\n\t"
    }
    p1 = (148 >= 18446744073709551615)
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 int32 = 2147483647
    var v2 float32 = 1e10
    v3 := make(map[int]uint32)
    v4 := []int{0, -1, -503, -655}
    v1 += (v1 * (v1 << ((616) & 63)))
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 float64 = (0.1 + 1e10)
        fmt.Println(v9, v9)
        var v10 uint16 = 42
        _ = fi8
        _ = v10
        _ = v9
    }
    v2 = float32(((0.1 * 3.14159) * -2.25))
    v2 = v2
    switch 42 {
    case 128:
        var v11 uint32 = 42
        _ = v11
    }
    fmt.Println(v1, v1)
    fmt.Println(v1, v1)
    if len(v4) > 0 {
        v4[0] = -370
    }
    v12 := GenS1{F0: (v1 | 42), F1: "", F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
