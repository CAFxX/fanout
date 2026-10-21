package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 bool }

func main() {
    v1 := make(map[int]string)
    var v2 uint = 42
    v3 := make(map[string]bool)
    var v4 int8 = -1
    var v5 uint8 = 255
    v4 += 127
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := GenS1{F0: (v5 ^ (^(v5))), F1: uint16(v5)}
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 uint64 = 349
            v9 -= (v9 % ((826) | 1))
            v9 = (v9 << (((v2 % ((v2) | 1))) & 63))
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    fmt.Println(v4, v4, v5)
    v10 := []int{-778, 1}
    if true {
        v11 := GenS2{F0: false}
        v5 -= v5
        _ = v11
    }
    var v12 int = 127
    var v13 int64 = 42
    var v14 int8 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
