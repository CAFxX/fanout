package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 uint16; F2 bool }
func genFn0(p0 int32, p1 int64) int32 {
    var v7 uint32 = 2147483647
    v8 := []int{}
    v7 = (v7 / ((v7) | 1))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p0
}
func genFn1(p0 int16) int8 {
    v9 := GenS1{F0: 255}
    var v10 uint16 = 32767
    v11 := []int{}
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return -64
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]int)
    var v3 bool = false
    var v4 uint64 = 868
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]bool)
    if v3 {
        v4 -= 18446744073709551615
    }
    var v12 uint8 = 125
    v13 := GenS1{F0: 65535}
    var v14 int8 = genFn1(38)
    v15 := GenS2{F0: (^(v4))}
    v16 := GenS3{F0: "foo bar"}
    v12 -= 2
    switch 1 {
    case 9223372036854775807:
        v17 := GenS4{F0: v4, F1: 32767, F2: (v3 && (1 < 42))}
        _ = v17
    case 128:
        var v18 float64 = 0.1
        _ = v18
    case 4294967295:
        v19 := GenS4{F0: (v4 | 4294967295), F1: 255, F2: v3}
        _ = v19
    default:
        v3 = (255 > 0)
    }
    v14 += -1
    var v20 int16 = 1000
    v3 = (v12 != 147)
    var v21 uint32 = 2147483647
    fmt.Println(v12)
    var v22 uint32 = (^(v21))
    _ = len(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
