package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint64; F1 uint16; F2 int64 }
type GenS3 struct { F0 uint }
func genFn0(p0 int32, p1 int32) uint {
    v4 := []int{-1}
    v5 := GenS1{F0: "a", F1: 0}
    var v6 int = 9223372036854775807
    switch (v6 - (v6 * (-(v6)))) {
    case 9223372036854775807:
        v6 -= 244
    default:
        var v7 bool = false
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 18446744073709551615
}

func main() {
    v1 := []uint32{65535}
    var v2 int16 = -32768
    var v3 int64 = 1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v8 := GenS2{F0: 1, F1: 839, F2: v3}
    v3 = (v3 / ((v3) | 1))
    var v9 int = 127
    v1 = append(v1, 1)
    if false {
        var v10 int16 = -739
        _ = v10
    }
    _ = len(v1)
    _ = len(v1)
    v3 = 1
    var v11 bool = false
    v9 += 2147483647
    v12 := GenS3{F0: 4294967295}
    fmt.Println(v9, v9, v3)
    var v13 uint32 = 207
    switch v9 {
    case -427:
        if len(v1) > 0 {
            v1[0] = (v13 | 2147483648)
        }
    }
    v3 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
