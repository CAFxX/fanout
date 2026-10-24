package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }

func main() {
    v1 := make(map[string]string)
    v2 := []int{-502}
    var v3 int16 = 32767
    v4 := []uint32{568}
    var v5 bool = (35 <= 42)
    v6 := []uint32{1, 23}
    var v7 uint8 = 0
    if (!(true || ((^(v7)) <= (v7 - (v7 & v7))))) {
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 uint = 18446744073709551615
            _ = fi8
            _ = v9
        }
        v7 = (v7 / ((128) | 1))
    }
    v5 = false
    v7 -= v7
    v10 := GenS1{F0: v7, F1: v7}
    var v11 int16 = (^(v3))
    var v12 uint64 = 4294967295
    if v5 {
        v5 = false
        v13 := GenS1{F0: 0, F1: (v7 & (v7 % ((v7) | 1)))}
        var v14 uint64 = (^(v12))
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 5; fi15++ {
        v6 = append(v6, 65535)
        _ = len(v4)
        v12 += (v12 * (^(v12)))
        _ = fi15
    }
    v7 = (v7 - (v7 ^ (v7 | (v7 ^ v7))))
    v3 += 1000
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
