package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 bool; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int32; F1 int; F2 uint }
func genFn0(p0 uint8, p1 uint) uint8 {
    v7 := GenS1{F0: "a", F1: -1}
    v8 := GenS1{F0: "", F1: 144}
    v9 := []int{}
    if len(v9) > 0 {
        v9[0] = -28
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return p0
}
func genFn1(p0 int16) int32 {
    v10 := GenS2{F0: false, F1: 42, F2: 0}
    if (711 > -1) {
        var v11 float32 = (3.14159 + 3.14159)
        v11 = 1e10
        v12 := GenS3{F0: 98, F1: -1, F2: 18446744073709551615}
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v10
    return 0
}

func main() {
    var v1 int32 = 2147483647
    var v2 float64 = 1e10
    v3 := make(map[uint32]bool)
    var v4 uint16 = 98
    var v5 float64 = (((v2 - -2.25) / v2) / 1e10)
    var v6 uint64 = 541
    if (-1 <= 1) {
        fmt.Println(v5, v5, v5)
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        var v14 uint16 = v4
        _ = fi13
        _ = v14
    }
    v2 = (v5 / v2)
    var v15 float32 = -2.25
    fmt.Println(v6)
    v16 := []int{32768, 42, 0}
    var v17 float64 = v5
    fmt.Println(v4, v5)
    v18 := GenS4{F0: 775}
    for fi19 := 0; fi19 < 5; fi19++ {
        _ = len(v3)
        _ = fi19
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        fmt.Println(v17)
        switch "foo bar" {
        case "a":
            v16 = append(v16, 127)
        case "foo bar":
            v17 = (v17 + 1.5)
        }
        v1 = (v1 % ((0) | 1))
        _ = fi20
    }
    v21 := v3[2147483647]
    fmt.Println(v5, v15)
    v21 = (42 != 0)
    v22 := GenS5{F0: 2147483647, F1: 32767, F2: 697}
    v21 = v21
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
