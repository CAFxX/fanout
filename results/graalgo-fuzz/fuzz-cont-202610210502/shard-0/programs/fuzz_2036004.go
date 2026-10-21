package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 string; F2 int32 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 uint32 }

func main() {
    v1 := []int{-2147483648, 1, 732}
    var v2 int32 = 1
    var v3 bool = ((0.0 < 1e10) && ((v2 << ((1) & 63)) != (-(v2))))
    v4 := []int{32767}
    var v5 bool = true
    v3 = v5
    switch "\n\t" {
    case "hello":
        v6 := GenS1{F0: 127}
        _ = v6
    case "a":
        v7 := GenS2{F0: v2}
        _ = v7
    }
    var v8 int32 = (v2 ^ (-(v2)))
    v3 = v5
    var v9 float32 = (1e10 - (float32(1.5) - 0.1))
    var v10 int64 = int64(v8)
    v4 = append(v4, -2147483648)
    if v3 {
        v11 := GenS3{F0: 42, F1: "a", F2: (v8 >> ((565) & 63))}
        _ = v11
    }
    switch 32767 {
    case -675:
        v1 = append(v1, 255)
    default:
        if v5 {
            v12 := GenS2{F0: 23}
            _ = v12
        }
    }
    if (!true) {
        if v3 {
            var v13 int16 = 42
            _ = v13
        } else {
            for fi14 := 0; fi14 < 5; fi14++ {
                v15 := GenS4{F0: 439, F1: 2}
                v16 := GenS4{F0: 1, F1: 65535}
                v9 = float32(float64(1.5))
                _ = fi14
                _ = v15
                _ = v16
            }
            var v17 int16 = 1
            for fi18 := 0; fi18 < 4; fi18++ {
                v19 := GenS3{F0: 65535, F1: "\n\t", F2: v2}
                for fi20 := 0; fi20 < 17; fi20++ {
                    v5 = v5
                    _ = fi20
                }
                _ = fi18
                _ = v19
            }
            _ = v17
        }
        var v21 uint = 42
        for fi22 := 0; fi22 < 11; fi22++ {
            v5 = v5
            var v23 int16 = 42
            _ = fi22
            _ = v23
        }
        _ = v21
    } else {
        for fi24 := 0; fi24 < 12; fi24++ {
            var v25 float64 = 0.0
            v10 = v10
            _ = fi24
            _ = v25
        }
        _ = len(v4)
    }
    v8 = -1
    if len(v1) > 0 {
        v1[0] = -467
    }
    if len(v1) > 0 {
        v1[0] = 128
    }
    v26 := GenS5{F0: 4294967295}
    if ((!(-2147483648 == 127)) || false) {
        v9 = v9
        var v27 uint64 = 435
        _ = v27
    } else {
        fmt.Println(v2, v8, v9)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 = append(v1, -861)
    if len(v1) > 0 {
        v1[0] = 1
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
