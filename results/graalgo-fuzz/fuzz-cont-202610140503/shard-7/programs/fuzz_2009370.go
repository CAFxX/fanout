package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 uint64; F1 int; F2 uint8 }
type GenS4 struct { F0 bool; F1 int64; F2 int8 }
type GenS5 struct { F0 int64; F1 int; F2 int32 }
func genFn0(p0 uint8, p1 uint32, p2 int16) uint64 {
    v4 := []int{-2147483648, 4294967295}
    var v5 uint16 = uint16(p0)
    var v6 uint64 = 541
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return (^(v6))
}

func main() {
    var v1 int = 0
    v2 := make(map[int]int)
    v3 := make(map[uint32]bool)
    if false {
        if (42 < 1000) {
            v7 := GenS1{F0: 1, F1: true}
            _ = v7
        }
        v2[(v1 + (v1 / ((0) | 1)))] = (v1 >> ((4294967295) & 63))
    }
    fmt.Println(v1)
    switch "foo bar" {
    case "\n\t":
        v8 := GenS2{F0: 18446744073709551615, F1: -48}
        _ = v8
    default:
        v9 := GenS3{F0: 520, F1: 339, F2: 1}
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := []int{(v1 % (((v1 & (^(v1)))) | 1))}
        _ = fi10
        _ = v11
    }
    fmt.Println(v1)
    var v12 int64 = 156
    v12 = (v12 + v12)
    v13 := GenS4{F0: (42 <= 42), F1: v12, F2: 127}
    v14 := GenS5{F0: v12, F1: (v1 ^ (v1 | (v1 | v1))), F2: -880}
    v15 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
