package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 string; F1 int64; F2 uint64 }
func genFn0(p0 bool, p1 uint64) int {
    var v6 int8 = -1
    v7 := GenS1{F0: (v6 * v6), F1: 0, F2: 403}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 32767
}

func main() {
    var v1 uint32 = 566
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]bool)
    var v4 uint8 = 42
    var v5 int64 = 42
    v8 := []int{}
    v5 -= -437
    var v9 uint = 65535
    v1 -= (v1 % (((v1 << (((v9 - 1)) & 63))) | 1))
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 int16 = 0
        _ = fi10
        _ = v11
    }
    _ = len(v8)
    v9 = (v9 >> ((v9) & 63))
    fmt.Println(v1, v5)
    var v12 uint = (^(v9))
    v13 := GenS2{F0: -1, F1: 100}
    v14 := GenS3{F0: "\n\t", F1: v5, F2: 4294967295}
    if true {
        v12 = 255
        _ = len(v8)
    } else {
        v15 := GenS1{F0: 127, F1: v4, F2: v9}
        v9 -= v12
        fmt.Println(v5, v5, v12)
        _ = v15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
