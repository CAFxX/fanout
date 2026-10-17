package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint64; F2 int64 }

func main() {
    var v1 float64 = -2.25
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]int)
    var v4 int64 = -170
    var v5 uint32 = 216
    var v6 uint16 = 65535
    v7 := []int{-357, -996, 32767}
    var v8 uint8 = 42
    v6 += 200
    v9 := GenS1{F0: 1, F1: v6}
    var v10 uint8 = (^(v8))
    switch "\n\t" {
    case "":
        var v11 float32 = 1.5
        _ = v11
    case "\n\t":
        if true {
            var v12 float64 = v1
            var v13 uint16 = (v6 / ((2) | 1))
            _ = v12
            _ = v13
        }
    }
    if false {
        var v14 uint64 = 4294967295
        fmt.Println(v10)
        switch "a" {
        case "hello":
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        }
        _ = v14
    } else {
        fmt.Println(v10, v8, v4)
    }
    fmt.Println(v4)
    switch "\n\t" {
    case "foo bar":
        v8 += (v10 - v10)
    case "\n\t":
        for fi15 := 0; fi15 < 2; fi15++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi15
        }
    case "":
        if len(v7) > 0 {
            v7[0] = 32767
        }
    default:
        v16 := GenS2{F0: (v5 / ((0) | 1)), F1: 4294967295, F2: v4}
        _ = v16
    }
    switch "foo bar" {
    case "a":
        v4 = (v4 - v4)
    default:
        v17 := GenS2{F0: (v5 + v5), F1: uint64(v10), F2: (v4 << ((1) & 63))}
        _ = v17
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
