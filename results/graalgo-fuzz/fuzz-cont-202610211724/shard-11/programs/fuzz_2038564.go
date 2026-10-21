package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 bool; F1 int32 }
func genFn0(p0 uint64) int8 {
    v6 := GenS1{F0: 0, F1: 483}
    var v7 uint64 = p0
    var v8 int = 841
    for fi9 := 0; fi9 < 17; fi9++ {
        fmt.Println(p0)
        fmt.Println(v7)
        var v10 uint8 = 0
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 42
}
func genFn1(p0 string) uint8 {
    var v11 uint8 = 128
    v12 := []int{-420, 790, -654}
    if false {
        var v13 int8 = 31
        var v14 int64 = 0
        switch 9223372036854775807 {
        case 42:
            p0 = p0
        case -203:
            if len(v12) > 0 {
                v12[0] = 0
            }
        default:
            var v15 int = -1
            _ = v15
        }
        _ = v13
        _ = v14
    } else {
        var v16 uint = 901
        switch "hello" {
        case "a":
            p0 = (((p0 + p0) + p0) + p0)
        case "hello":
            switch 889 {
            case 2147483647:
                v17 := GenS2{F0: -714, F1: p0}
                _ = v17
            case 984:
                p0 = p0
            }
        default:
            var v18 int = 4294967295
            _ = v18
        }
        _ = v16
    }
    if len(v12) > 0 {
        v12[0] = 9223372036854775807
    }
    v19 := GenS2{F0: 199, F1: p0}
    _ = p0
    _ = v11
    _ = v12
    _ = v19
    return v11
}

func main() {
    var v1 int8 = -1
    var v2 int16 = 1000
    v3 := []string{}
    var v4 int64 = -29
    v5 := []uint32{0, 2147483647}
    fmt.Println(v2, v4)
    v3 = append(v3, "\n\t")
    v20 := GenS3{F0: false, F1: -1}
    _ = len(v3)
    v2 -= 927
    v4 -= (v4 & (v4 % ((v4) | 1)))
    v4 += v4
    var v21 float64 = float64(((1e10 / (1e10 * 1e10)) / -2.25))
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 uint8 = 255
        v2 = 42
        v4 = (v4 << ((18446744073709551615) & 63))
        _ = fi22
        _ = v23
    }
    v5 = append(v5, 42)
    var v24 uint64 = 18446744073709551615
    v25 := GenS1{F0: 42, F1: 2147483647}
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
