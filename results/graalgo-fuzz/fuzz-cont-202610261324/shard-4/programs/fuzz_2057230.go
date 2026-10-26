package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int32; F1 string; F2 bool }

func main() {
    v1 := []int{-243, 1, -264}
    var v2 int8 = -103
    v3 := make(map[string]string)
    var v4 uint32 = 1
    var v5 int16 = 42
    v6 := make(map[string]uint32)
    v4 = (v4 >> ((4294967295) & 63))
    fmt.Println(v2, v2, v4)
    if false {
        fmt.Println(v2, v2, v2)
    } else {
        for fi7 := 0; fi7 < 20; fi7++ {
            v5 += v5
            _ = fi7
        }
        for fi8 := 0; fi8 < 17; fi8++ {
            for fi9 := 0; fi9 < 15; fi9++ {
                v10 := GenS1{F0: -826, F1: 1}
                fmt.Println(v5, v4, v2)
                v2 = (v2 % ((v2) | 1))
                _ = fi9
                _ = v10
            }
            v5 -= v5
            var v11 int = (fi8 - fi8)
            _ = fi8
            _ = v11
        }
    }
    var v12 int32 = 443
    v13 := GenS2{F0: v12, F1: "hello", F2: true}
    fmt.Println(v2, v5)
    v5 += v5
    v4 += (^(v4))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
