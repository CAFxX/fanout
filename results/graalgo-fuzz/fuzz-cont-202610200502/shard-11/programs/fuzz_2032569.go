package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
func genFn0(p0 uint64) int64 {
    fmt.Println(p0)
    if (540 > 42) {
        if true {
            for fi5 := 0; fi5 < 16; fi5++ {
                fmt.Println(fi5)
                _ = fi5
            }
        }
        v6 := GenS1{F0: 0, F1: 65535}
        _ = v6
    }
    switch 42 {
    case 9223372036854775807:
        v7 := GenS1{F0: (p0 & (p0 ^ (^(p0)))), F1: 32767}
        _ = v7
    default:
        fmt.Println(p0)
    }
    v8 := []int{}
    for fi9 := 0; fi9 < 9; fi9++ {
        var v10 int16 = -1
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v8
    return -1
}

func main() {
    v1 := []int{}
    var v2 int32 = 818
    var v3 uint64 = 2
    v4 := make(map[uint32]string)
    v11 := GenS1{F0: (v3 | (v3 | (v3 * (^(v3))))), F1: 42}
    fmt.Println(v3, v2)
    var v12 float32 = -2.25
    var v13 int = -195
    v13 += v13
    var v14 int8 = 42
    fmt.Println(v3)
    v1 = append(v1, (v13 & v13))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
