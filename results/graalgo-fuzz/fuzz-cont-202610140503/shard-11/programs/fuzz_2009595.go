package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 bool }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint32 }

func main() {
    var v1 uint32 = 980
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]bool)
    var v4 int16 = 32767
    var v5 float64 = (0.1 * (-2.25 / (1e10 + -2.25)))
    var v6 int64 = 0
    if false {
        v7 := GenS1{F0: "hello", F1: 520, F2: 4294967295}
        _ = v7
    }
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS2{F0: -1}
        _ = fi8
        _ = v9
    }
    v10 := GenS3{F0: 1, F1: true}
    v11 := []int{-918, -1, 2147483647}
    v4 += (v4 & 64)
    v12 := GenS4{F0: 1, F1: 128, F2: v1}
    if len(v11) > 0 {
        v11[0] = 2147483647
    }
    var v13 int = 2147483647
    if (false && false) {
        var v14 uint8 = 1
        var v15 bool = true
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 int64 = int64(v4)
            switch "hello" {
            case "hello":
                v13 = (fi16 % ((v13) | 1))
            case "a":
                v15 = v15
            case "":
                var v18 int = 0
                _ = v18
            default:
                v6 -= -465
            }
            v6 += (v6 + v6)
            _ = fi16
            _ = v17
        }
        _ = v14
        _ = v15
    }
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    v11 = append(v11, (v13 ^ v13))
    _ = len(v11)
    v6 = (v6 >> ((255) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
