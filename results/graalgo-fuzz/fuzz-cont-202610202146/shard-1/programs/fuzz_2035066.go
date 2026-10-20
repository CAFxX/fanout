package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }

func main() {
    var v1 uint32 = 336
    var v2 int32 = -505
    var v3 uint = 1
    v4 := make(map[string]bool)
    v1 += v1
    v5 := GenS1{F0: -473, F1: -222}
    for fi6 := 0; fi6 < 1; fi6++ {
        v3 -= (v3 - (v3 % (((v3 % ((v3) | 1))) | 1)))
        _ = fi6
    }
    switch -554 {
    case -853:
        v7 := []int{}
        _ = v7
    case 32767:
        fmt.Println(v3, v2)
    case -762:
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 uint64 = 0
            _ = fi8
            _ = v9
        }
    default:
        v4["a"] = (!true)
    }
    v10 := []int{}
    v2 -= v2
    var v11 uint64 = 0
    var v12 uint64 = 412
    fmt.Println(v1, v3, v11)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
