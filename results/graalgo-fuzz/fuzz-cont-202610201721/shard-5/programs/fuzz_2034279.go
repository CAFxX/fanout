package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 uint8; F1 uint64 }
type GenS4 struct { F0 uint16; F1 bool }
func genFn0(p0 uint) int32 {
    p0 = p0
    v5 := GenS1{F0: -9223372036854775808}
    _ = p0
    _ = v5
    return -2147483648
}
func genFn1(p0 uint, p1 uint64) bool {
    switch "a" {
    case "a":
        var v6 int8 = 57
        _ = v6
    case "foo bar":
        p1 -= (p1 ^ (p1 * (p1 << ((255) & 63))))
    case "\n\t":
        switch "a" {
        case "\n\t":
            v7 := GenS2{F0: (genFn0(0) == 0), F1: uint8(p0)}
            _ = v7
        default:
            var v8 int = 32768
            _ = v8
        }
    default:
        for fi9 := 0; fi9 < 20; fi9++ {
            p0 -= (p0 | p0)
            p0 = p0
            v10 := GenS2{F0: true, F1: 2}
            _ = fi9
            _ = v10
        }
    }
    var v11 uint = (p0 * 0)
    _ = p0
    _ = p1
    _ = v11
    return true
}

func main() {
    v1 := []string{}
    var v2 float64 = 0.1
    v3 := []uint32{2, 401, 2, 650}
    v4 := []string{}
    v2 += v2
    switch -9223372036854775808 {
    case 42:
        v2 += v2
    }
    v12 := GenS1{F0: 583}
    for fi13 := 0; fi13 < 14; fi13++ {
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 float32 = 0.0
            var v16 string = "\n\t"
            v17 := GenS1{F0: 1}
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        _ = fi13
    }
    v18 := GenS3{F0: 127, F1: 0}
    v19 := GenS3{F0: 127, F1: 2}
    v2 += v2
    var v20 uint8 = 255
    v21 := GenS4{F0: 0, F1: (false && true)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
