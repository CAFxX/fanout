package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint }
func genFn0(p0 int, p1 uint16, p2 int) uint16 {
    if (2147483647 != -2147483648) {
        var v4 float32 = 1e10
        var v5 int16 = 1000
        v4 = 0.0
        _ = v4
        _ = v5
    } else {
        p2 += p2
        if false {
            p2 += p0
        } else {
            p0 -= 2147483647
            v6 := []int{p0, p0, (p0 - 1)}
            _ = v6
        }
    }
    v7 := GenS1{F0: 2, F1: -128, F2: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 65535
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]uint32)
    v8 := GenS2{F0: 65535}
    var v9 bool = (!true)
    var v10 uint64 = 0
    v11 := GenS2{F0: 512}
    v12 := GenS1{F0: 42, F1: -23, F2: 0}
    switch 93 {
    case 32768:
        var v13 uint = 255
        _ = v13
    default:
        fmt.Println(v9, v10)
    }
    var v14 int = 4294967295
    v15 := v3[v14]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
