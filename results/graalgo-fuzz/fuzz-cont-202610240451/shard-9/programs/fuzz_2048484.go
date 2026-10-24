package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int8; F1 string; F2 uint16 }
func genFn0(p0 uint, p1 string) uint8 {
    var v5 uint8 = 195
    v6 := GenS1{F0: (p0 / ((uint(v5)) | 1))}
    if false {
        v7 := []int{-9223372036854775808, -541, 4294967295, 1}
        v7 = append(v7, 2147483647)
        _ = v7
    }
    v8 := []int{127}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v8
    return (v5 ^ (v5 | v5))
}
func genFn1(p0 string, p1 uint64, p2 uint32) int8 {
    var v9 uint32 = uint32(p1)
    var v10 bool = (-51 >= 125)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return -61
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = 131
    v3 := []uint32{2147483647, 1}
    var v4 uint8 = 35
    v11 := GenS1{F0: 42}
    v2 -= (^(v2))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v12 := GenS2{F0: 124, F1: v4, F2: 255}
    if true {
        fmt.Println(v4, v2)
    }
    var v13 string = "hello"
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 string = v13
        delete(v1, 1)
        v2 -= fi14
        _ = fi14
        _ = v15
    }
    var v16 uint64 = 0
    v17 := GenS2{F0: 1, F1: genFn0(626, v13), F2: 178}
    switch 1 {
    case 124:
        v18 := GenS3{F0: 42, F1: ("hello" + "foo bar"), F2: 1}
        _ = v18
    case -9223372036854775808:
        var v19 uint = 42
        _ = v19
    case 4294967295:
        var v20 bool = (0 == 0)
        _ = v20
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        v2 = 127
        _ = fi21
    }
    var v22 uint16 = 42
    v23 := GenS2{F0: (v22 & 469), F1: 2, F2: (v22 % (((v22 + (^(v22)))) | 1))}
    var v24 uint64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
