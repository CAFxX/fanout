package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint }

func main() {
    v1 := []int{}
    v2 := make(map[string]string)
    var v3 uint = 18446744073709551615
    v4 := []int{}
    v5 := make(map[int]bool)
    v3 = 4294967295
    if (!true) {
        fmt.Println(v3)
    } else {
        switch -9223372036854775808 {
        case 128:
            v3 = (v3 ^ v3)
        case 827:
            fmt.Println(v3)
        }
        for fi6 := 0; fi6 < 2; fi6++ {
            v3 = v3
            v7 := GenS1{F0: -2147483648}
            _ = fi6
            _ = v7
        }
        var v8 uint16 = 1
        _ = v8
    }
    switch 32768 {
    case 796:
        v9 := GenS2{F0: uint64(v3), F1: v3}
        _ = v9
    case 32767:
        fmt.Println(v3)
    case 4294967295:
        v3 = v3
    default:
        var v10 uint64 = 4294967295
        _ = v10
    }
    v3 = v3
    v11 := GenS2{F0: 302, F1: 255}
    var v12 uint16 = 32767
    var v13 uint = (v3 + 4294967295)
    var v14 int8 = 42
    var v15 int32 = 42
    var v16 int8 = v14
    for fi17 := 0; fi17 < 4; fi17++ {
        fmt.Println(v3, fi17)
        _ = fi17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
