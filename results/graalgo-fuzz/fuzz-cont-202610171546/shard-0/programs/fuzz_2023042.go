package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint64 }

func main() {
    var v1 float64 = 1.5
    v2 := []uint32{1, 4294967295, 4294967295}
    v3 := make(map[uint32]uint32)
    v4 := []int{-807, -650}
    v5 := make(map[int]bool)
    if true {
        switch 2147483647 {
        case 2147483647:
            var v6 uint8 = 128
            _ = v6
        case 32768:
            var v7 uint = 18446744073709551615
            _ = v7
        case 902:
            v1 = v1
        default:
            for fi8 := 0; fi8 < 6; fi8++ {
                v2 = append(v2, 2147483647)
                _ = fi8
            }
        }
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 int16 = 32767
            v11 := GenS1{F0: false, F1: v10}
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 uint16 = 65535
        _ = v12
    } else {
        switch -691 {
        case 4294967295:
            var v13 float64 = v1
            _ = v13
        case -1:
            v14 := v5[-594]
            _ = v14
        default:
            v1 = (0.0 / v1)
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch -473 {
    case 128:
        for fi15 := 0; fi15 < 1; fi15++ {
            v1 = ((v1 - float64(1.5)) + v1)
            _ = fi15
        }
    case 0:
        var v16 uint = 890
        _ = v16
    case -1:
        v1 = v1
    default:
        var v17 int64 = -9223372036854775808
        _ = v17
    }
    var v18 float32 = ((float32(float64(0.0)) / 1e10) - 1.5)
    v18 = v18
    switch ("hello" + "") {
    case "\n\t":
        fmt.Println(v1, v1)
    case "a":
        v19 := GenS1{F0: false, F1: -455}
        _ = v19
    default:
        v18 += 3.14159
    }
    v20 := GenS2{F0: 42, F1: 42, F2: 4294967295}
    v21 := GenS2{F0: 185, F1: 765, F2: 18446744073709551615}
    v18 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
