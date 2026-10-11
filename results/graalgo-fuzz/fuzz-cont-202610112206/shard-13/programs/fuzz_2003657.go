package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32 }

func main() {
    v1 := make(map[string]bool)
    var v2 int32 = 1
    v3 := make(map[uint32]string)
    v4 := make(map[int]int)
    v2 -= (v2 << ((18446744073709551615) & 63))
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 string = "a"
        v4[(fi5 ^ -642)] = 9223372036854775807
        v2 = (v2 % ((v2) | 1))
        _ = fi5
        _ = v6
    }
    var v7 int64 = 42
    v8 := GenS1{F0: 2, F1: 1, F2: v7}
    v9 := []int{302, 263}
    var v10 uint64 = 1
    v11 := GenS2{F0: 255}
    v12 := GenS2{F0: 411}
    v3[650] = "a"
    v13 := GenS3{F0: (-(v2))}
    v14 := GenS1{F0: 42, F1: 62, F2: v7}
    var v15 int16 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
