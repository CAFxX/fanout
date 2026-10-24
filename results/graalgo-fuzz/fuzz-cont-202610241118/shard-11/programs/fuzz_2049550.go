package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 string; F1 bool; F2 bool }

func main() {
    var v1 uint16 = 0
    v2 := []int{-103, 32767, 1}
    v3 := make(map[string]uint32)
    v4 := make(map[string]bool)
    var v5 int64 = 1
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 int8 = 42
        _ = fi6
        _ = v7
    }
    _ = len(v2)
    for fi8 := 0; fi8 < 7; fi8++ {
        fmt.Println(v5, v1, v5)
        v9 := GenS1{F0: v1}
        fmt.Println(v5, v1, fi8)
        _ = fi8
        _ = v9
    }
    var v10 uint8 = 0
    v1 -= v1
    v1 -= 24
    if false {
        fmt.Println(v5, v10)
        _ = len(v4)
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        for fi12 := 0; fi12 < 11; fi12++ {
            switch fi12 {
            case 128:
                var v13 bool = ((fi12 ^ (fi12 / (((fi11 / ((0) | 1))) | 1))) <= (fi12 + (fi11 >> ((255) & 63))))
                _ = v13
            case 9223372036854775807:
                v14 := GenS2{F0: ("foo bar" + "a")}
                _ = v14
            case 178:
                v15 := GenS3{F0: 1}
                _ = v15
            default:
                if ((!(true || false)) || (float32(-2.25) <= 0.0)) {
                    v16 := GenS4{F0: 34, F1: int(v10)}
                    v17 := GenS5{F0: "foo bar", F1: (!false), F2: (false && (v10 > (^(v10))))}
                    _ = v16
                    _ = v17
                }
            }
            _ = fi12
        }
        _ = fi11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
