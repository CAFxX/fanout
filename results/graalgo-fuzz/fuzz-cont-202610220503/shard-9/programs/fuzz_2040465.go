package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint8 }
type GenS2 struct { F0 int; F1 int64 }

func main() {
    var v1 uint32 = 626
    v2 := make(map[int]bool)
    var v3 int8 = -128
    v4 := make(map[uint32]string)
    var v5 uint8 = 128
    var v6 int32 = 42
    var v7 float32 = ((0.0 / 3.14159) * 1.5)
    var v8 uint64 = 295
    v8 += (v8 ^ (v8 + v8))
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := []int{fi9, (fi9 >> ((60) & 63)), -318, -218}
        fmt.Println(v6, v6)
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := GenS1{F0: 631, F1: "foo bar", F2: (^(v5))}
        var v13 int = fi11
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: v1, F1: "", F2: v5}
    var v15 float32 = (((v7 + (v7 + (float32(((0.0 / -2.25) / -2.25)) + 0.1))) / (0.1 + v7)) + 1e10)
    v16 := GenS2{F0: 128, F1: int64(v5)}
    v17 := []int{830, 2147483647, -2147483648}
    var v18 uint16 = 1
    var v19 int16 = -32768
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
