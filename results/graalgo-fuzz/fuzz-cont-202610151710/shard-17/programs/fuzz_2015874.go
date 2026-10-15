package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int }
type GenS2 struct { F0 uint32 }
func genFn0(p0 string, p1 int64) int16 {
    var v4 string = p0
    v5 := []int{-2147483648, 128, 32767}
    fmt.Println(p0, p1)
    fmt.Println(v4, p1)
    if (v4 >= v4) {
        var v6 float32 = 1e10
        var v7 bool = false
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 1000
}

func main() {
    var v1 uint64 = 4294967295
    var v2 uint64 = 539
    v3 := make(map[int]bool)
    v1 = v1
    v1 = v1
    var v8 int = 996
    fmt.Println(v2, v1)
    var v9 uint64 = (v2 | 42)
    v8 -= (v8 << ((798) & 63))
    v10 := GenS1{F0: 0, F1: 0, F2: v8}
    if true {
        delete(v3, (-(v8)))
        fmt.Println(v1, v1, v9)
        fmt.Println(v9)
    } else {
        var v11 float64 = 0.1
        v12 := GenS1{F0: 0, F1: 1, F2: (v8 - v8)}
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: 727}
    if true {
        _ = len(v3)
        fmt.Println(v8, v8)
    } else {
        switch "a" {
        case "a":
            var v14 float64 = 0.0
            _ = v14
        }
    }
    v15 := GenS2{F0: 0}
    v8 = v8
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
