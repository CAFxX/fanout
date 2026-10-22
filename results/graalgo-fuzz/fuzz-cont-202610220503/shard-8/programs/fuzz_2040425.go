package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int64; F1 uint; F2 string }
func genFn0(p0 int16, p1 uint32) uint64 {
    v4 := []int{132}
    if (!false) {
        fmt.Println(p1, p0)
        v5 := GenS1{F0: p1, F1: 1}
        var v6 int8 = -128
        _ = v5
        _ = v6
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = v4
    return 274
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]uint32)
    var v3 int = -1
    v3 += 2147483647
    v7 := GenS1{F0: 741, F1: 2}
    var v8 int16 = -863
    v3 -= v3
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 float64 = float64((1e10 - 3.14159))
        v11 := []int{v3}
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS1{F0: 42, F1: 18446744073709551615}
    var v13 bool = true
    var v14 uint = 42
    var v15 int64 = -1
    v15 += v15
    var v16 float64 = -2.25
    v17 := GenS2{F0: (1.5 < 1.5)}
    var v18 uint8 = 42
    v19 := GenS3{F0: (-(v15)), F1: 255, F2: ""}
    var v20 float32 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v7
    _ = v8
}
