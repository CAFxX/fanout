package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int64; F1 uint32; F2 int16 }
type GenS3 struct { F0 int; F1 int; F2 uint }
type GenS4 struct { F0 uint; F1 uint8 }
func genFn0(p0 int, p1 uint64) int8 {
    fmt.Println(p1, p1)
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := GenS1{F0: p0, F1: 0}
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: p0, F1: -128}
    v9 := GenS2{F0: 0, F1: 1, F2: -363}
    p0 += 127
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 100
}
func genFn1(p0 int16, p1 int8, p2 int64) uint {
    p1 += (-(p1))
    for fi10 := 0; fi10 < 20; fi10++ {
        v11 := GenS3{F0: fi10, F1: fi10, F2: 0}
        p2 -= 9223372036854775807
        _ = fi10
        _ = v11
    }
    var v12 uint32 = 449
    fmt.Println(p1, v12)
    if true {
        v13 := []int{}
        _ = len(v13)
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int16 = 42
    v3 := []int{-1, -39, 1, 32767}
    v4 := make(map[int]uint32)
    var v5 string = ""
    var v14 int16 = (v2 | (v2 + (v2 % (((v2 + (v2 % ((v2) | 1)))) | 1))))
    fmt.Println(v14, v2, v5)
    var v15 uint16 = 32767
    var v16 int = 128
    var v17 int = (-(v16))
    var v18 uint = 1
    v4[v17] = 65535
    var v19 uint = genFn1((v2 / ((v2) | 1)), 42, 0)
    v20 := GenS4{F0: v18, F1: 127}
    var v21 int8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
