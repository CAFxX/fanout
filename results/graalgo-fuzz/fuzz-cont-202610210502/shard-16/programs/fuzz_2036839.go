package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 bool; F1 int64; F2 uint }
func genFn0(p0 string, p1 int) bool {
    if true {
        var v4 uint8 = 0
        _ = v4
    }
    switch p1 {
    case -382:
        v5 := GenS1{F0: 42, F1: 32767}
        _ = v5
    case -9223372036854775808:
        var v6 bool = false
        _ = v6
    default:
        var v7 float32 = (3.14159 * -2.25)
        _ = v7
    }
    _ = p0
    _ = p1
    return (2 < 0)
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    var v8 int16 = 1000
    fmt.Println(v8)
    for fi9 := 0; fi9 < 8; fi9++ {
        v10 := GenS2{F0: false, F1: 42, F2: 358}
        _ = fi9
        _ = v10
    }
    var v11 bool = true
    for fi12 := 0; fi12 < 12; fi12++ {
        fmt.Println(v11, v11, v11)
        var v13 uint64 = 18446744073709551615
        _ = fi12
        _ = v13
    }
    v8 = (v8 / ((v8) | 1))
    v8 = (v8 | v8)
    var v14 uint8 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v8
}
