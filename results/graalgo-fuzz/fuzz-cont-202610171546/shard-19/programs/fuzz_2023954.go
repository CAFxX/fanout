package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 uint = 2
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    v1 = (v1 | 65535)
    if true {
        var v4 float32 = 1e10
        _ = v4
    }
    var v5 float32 = 1e10
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 float64 = (1e10 * 0.0)
        v8 := GenS1{F0: fi6, F1: 1, F2: 192}
        v1 += v1
        _ = fi6
        _ = v7
        _ = v8
    }
    v1 += (v1 ^ (v1 + 37))
    v9 := []int{-9223372036854775808}
    fmt.Println(v1, v5)
    var v10 float32 = 3.14159
    v11 := GenS2{F0: 4294967295}
    fmt.Println(v5, v1, v10)
    var v12 float32 = v5
    v1 = (v1 - (v1 << (((v1 - v1)) & 63)))
    switch "hello" {
    case "foo bar":
        var v13 uint16 = 287
        _ = v13
    default:
        v14 := GenS2{F0: 2}
        _ = v14
    }
    if false {
        fmt.Println(v10)
    } else {
        fmt.Println(v10, v5, v10)
    }
    v10 += float32((3.14159 * 0.0))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v5
    _ = v9
}
