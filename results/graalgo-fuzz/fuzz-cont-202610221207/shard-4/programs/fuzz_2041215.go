package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint }
func genFn0(p0 int32) int16 {
    v6 := GenS1{F0: 2, F1: 4294967295, F2: 127}
    v7 := GenS2{F0: true, F1: 469}
    _ = p0
    _ = v6
    _ = v7
    return 883
}

func main() {
    v1 := []uint32{4294967295, 2147483647}
    var v2 float64 = 0.0
    v3 := []int{-1, -521, 127, 2147483647}
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]int)
    v2 = v2
    if false {
        var v8 string = (("\n\t" + "foo bar") + "\n\t")
        if (!false) {
            if false {
                v2 = v2
            }
            for fi9 := 0; fi9 < 2; fi9++ {
                fmt.Println(v8, v2)
                _ = fi9
            }
            for fi10 := 0; fi10 < 3; fi10++ {
                var v11 float32 = 1.5
                if true {
                    fmt.Println(v8, v11)
                }
                _ = fi10
                _ = v11
            }
        } else {
            v12 := GenS3{F0: 870}
            _ = v12
        }
        _ = v8
    } else {
        v2 -= v2
    }
    v2 = v2
    v1 = append(v1, 428)
    var v13 int16 = -468
    v14 := GenS4{F0: int32(v13)}
    _ = len(v1)
    switch -22 {
    case 877:
        var v15 uint8 = 128
        _ = v15
    case -9223372036854775808:
        fmt.Println(v2)
    default:
        var v16 bool = (18446744073709551615 >= 42)
        _ = v16
    }
    v17 := GenS5{F0: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
