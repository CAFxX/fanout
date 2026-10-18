package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint; F1 int32 }

func main() {
    var v1 uint = 2
    v2 := make(map[uint32]int)
    var v3 int16 = 42
    v4 := make(map[string]string)
    v5 := make(map[uint32]string)
    switch "foo bar" {
    case "":
        v6 := GenS1{F0: 42, F1: (381 > 42)}
        _ = v6
    }
    v7 := []int{}
    v8 := GenS1{F0: 42, F1: ((1 > 240) && true)}
    v9 := GenS2{F0: v1, F1: 827}
    for fi10 := 0; fi10 < 18; fi10++ {
        v3 -= (v3 & -1)
        _ = fi10
    }
    v11 := GenS2{F0: (v1 ^ (v1 >> ((136) & 63))), F1: int32(v1)}
    var v12 int32 = 1
    v1 -= 18446744073709551615
    v3 = 42
    v12 += 992
    for fi13 := 0; fi13 < 5; fi13++ {
        if true {
            var v14 uint64 = 645
            var v15 int32 = v12
            v15 -= (^(v15))
            _ = v14
            _ = v15
        }
        _ = fi13
    }
    _ = len(v7)
    var v16 uint16 = 833
    if (1.5 <= 0.1) {
        fmt.Println(v16, v12, v12)
        v12 -= int32(v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
