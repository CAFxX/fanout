package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint, p1 uint8, p2 uint64) uint16 {
    var v7 uint64 = (p2 ^ uint64(p0))
    var v8 uint32 = 504
    v9 := GenS1{F0: 42}
    if ((v8 - v8) != (^(v8))) {
        p1 = (p1 << (((p0 % ((42) | 1))) & 63))
        var v10 int16 = 42
        fmt.Println(p2, p0, p0)
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return uint16(v8)
}
func genFn1(p0 uint8, p1 uint32) uint32 {
    p1 += p1
    v11 := []int{0, -9223372036854775808, -835, -861}
    p0 -= (^(p0))
    var v12 int = -176
    if true {
        if false {
            var v13 uint8 = p0
            var v14 int32 = int32(p1)
            _ = v13
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return p1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{-289, 32768, -9223372036854775808, 32767}
    var v3 int32 = -2147483648
    var v4 int = -650
    var v5 uint8 = 83
    var v6 int64 = 0
    v4 = 32768
    v15 := GenS1{F0: 127}
    var v16 uint8 = 60
    v5 -= v16
    var v17 uint8 = 255
    fmt.Println(v3, v5)
    var v18 uint8 = (v16 >> ((2) & 63))
    var v19 float32 = 0.0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
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
}
