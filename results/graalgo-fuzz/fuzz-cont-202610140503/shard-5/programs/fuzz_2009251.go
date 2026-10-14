package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int64; F2 string }

func main() {
    v1 := make(map[int]int)
    var v2 uint32 = 2
    var v3 int32 = 42
    var v4 uint64 = 0
    fmt.Println(v4)
    v4 = (v4 % (((v4 & 660)) | 1))
    var v5 int8 = -122
    switch "a" {
    case "\n\t":
        switch -9223372036854775808 {
        case -9223372036854775808:
            v6 := GenS1{F0: 4294967295, F1: v4}
            _ = v6
        default:
            var v7 float32 = 1e10
            _ = v7
        }
    case "":
        v3 -= v3
    default:
        v8 := []int{-265, 747, 9223372036854775807}
        _ = v8
    }
    var v9 float32 = 0.1
    switch 32767 {
    case 32768:
        v5 += v5
    case -586:
        v9 = v9
    case 149:
        if (!false) {
            v4 += 2
        } else {
            v10 := GenS1{F0: 887, F1: v4}
            v9 -= -2.25
            _ = v10
        }
    default:
        v11 := GenS2{F0: uint32(v4), F1: -1, F2: ("a" + "\n\t")}
        _ = v11
    }
    v4 = v4
    var v12 uint16 = 2
    v5 -= (-(v5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
