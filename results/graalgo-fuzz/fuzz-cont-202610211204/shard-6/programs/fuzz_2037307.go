package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint64 }
func genFn0(p0 uint8, p1 string, p2 bool) int {
    p0 = p0
    var v5 float64 = float64(float32(1.5))
    var v6 int = 1
    var v7 int64 = 105
    fmt.Println(v7, v5, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return v6
}
func genFn1(p0 int64) uint8 {
    fmt.Println(p0)
    switch "" {
    case "a":
        p0 = (p0 / ((9223372036854775807) | 1))
    case "hello":
        p0 = 476
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 uint8 = 206
            v11 := []int{fi9, fi9, -38}
            v10 = (v10 + v10)
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    v12 := GenS1{F0: 18446744073709551615, F1: 514}
    var v13 uint = 295
    _ = p0
    _ = v12
    _ = v13
    return 255
}

func main() {
    v1 := []int{9223372036854775807, 588, 0}
    var v2 uint64 = 4294967295
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]bool)
    var v14 uint = 18446744073709551615
    v2 = (v2 >> ((v14) & 63))
    fmt.Println(v14, v14)
    switch "" {
    case "hello":
        v2 = uint64(v14)
    case "a":
        if true {
            v2 = v2
            v15 := v3[754]
            for fi16 := 0; fi16 < 12; fi16++ {
                v17 := GenS2{F0: -1, F1: (^(v2))}
                _ = fi16
                _ = v17
            }
            _ = v15
        }
    }
    if false {
        switch "foo bar" {
        case "":
            var v18 float32 = 1e10
            _ = v18
        default:
            v14 = (v14 & v14)
        }
    }
    switch 0 {
    case 2147483647:
        var v19 uint64 = 89
        _ = v19
    case 42:
        v20 := GenS1{F0: (v2 << ((v14) & 63)), F1: 255}
        _ = v20
    case 32767:
        var v21 uint16 = 2
        _ = v21
    }
    if false {
        v1 = append(v1, 32768)
    } else {
        fmt.Println(v2, v2)
    }
    v2 += 480
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
