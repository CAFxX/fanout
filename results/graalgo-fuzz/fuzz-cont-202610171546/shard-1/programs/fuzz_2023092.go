package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint; F1 uint; F2 bool }
func genFn0(p0 string, p1 int8) uint32 {
    var v4 int64 = 42
    switch 243 {
    case 1:
        p0 = ((((p0 + "a") + "a") + p0) + ("foo bar" + p0))
    case 32767:
        if false {
            v4 = (v4 * (v4 % ((v4) | 1)))
        }
    }
    fmt.Println(v4)
    _ = p0
    _ = p1
    _ = v4
    return 42
}
func genFn1() string {
    var v5 uint16 = 2
    fmt.Println(v5)
    v6 := []int{893, 1}
    v7 := GenS1{F0: 244, F1: 1}
    _ = v5
    _ = v6
    _ = v7
    return "a"
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 659
    v3 := []uint32{2147483647, 0}
    var v8 uint32 = 4294967295
    var v9 uint32 = (^(v8))
    var v10 uint32 = 0
    var v11 int8 = -72
    v12 := GenS1{F0: 2, F1: 462}
    v10 += 18
    switch 2147483647 {
    case 2147483647:
        fmt.Println(v11, v2)
    case -9223372036854775808:
        fmt.Println(v9, v10, v2)
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        switch fi13 {
        case 1:
            v14 := GenS2{F0: 530, F1: 4294967295, F2: true}
            _ = v14
        default:
            fmt.Println(v8)
        }
        switch (fi13 ^ 32768) {
        case 392:
            var v15 uint64 = 232
            _ = v15
        case 9223372036854775807:
            v10 = 932
        case -913:
            v2 = v2
        }
        var v16 int32 = 2147483647
        _ = fi13
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
