package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]string)
    v3 := []int{-504, -1}
    var v4 uint8 = 129
    var v5 uint64 = 341
    var v6 uint8 = v4
    fmt.Println(v5, v6)
    _ = len(v3)
    v2[2] = "\n\t"
    fmt.Println(v5, v5, v5)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (float32((0.1 / (-2.25 - 0.0))) <= 0.0) {
        v7 := GenS1{F0: true}
        _ = v7
    } else {
        v5 -= (v5 >> ((991) & 63))
        v8 := GenS1{F0: (!(1 >= 65535))}
        fmt.Println(v4, v6)
        _ = v8
    }
    v9 := GenS2{F0: 42, F1: 735, F2: uint16(v5)}
    v10 := GenS3{F0: 42}
    for fi11 := 0; fi11 < 16; fi11++ {
        fmt.Println(v6)
        _ = fi11
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
