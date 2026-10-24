package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32; F2 int64 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 bool; F1 uint8 }
func genFn0(p0 uint32) uint {
    fmt.Println(p0)
    v6 := GenS1{F0: 1, F1: 42, F2: -344}
    v7 := []int{127, 1, 706, 32768}
    v8 := GenS2{F0: 127, F1: 381}
    if false {
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 int8 = -128
            var v11 float64 = (((1.5 * (3.14159 / 1.5)) - (-2.25 + 1e10)) * 1.5)
            _ = fi9
            _ = v10
            _ = v11
        }
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 377
}

func main() {
    v1 := []uint32{0, 300}
    var v2 int16 = 1
    v3 := []int{4294967295, 32768, 32768}
    v4 := make(map[int]bool)
    var v5 bool = true
    v12 := GenS2{F0: 108, F1: 2}
    v13 := GenS3{F0: 1, F1: true}
    var v14 int = 375
    fmt.Println(v2, v5)
    var v15 int16 = (v2 << ((308) & 63))
    if len(v3) > 0 {
        v3[0] = 998
    }
    v16 := GenS4{F0: (320 > 65535), F1: 1}
    var v17 uint32 = 1
    v1 = append(v1, 209)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
