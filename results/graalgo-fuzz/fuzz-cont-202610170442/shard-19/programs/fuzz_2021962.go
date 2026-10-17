package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int16; F1 int64 }
func genFn0() uint64 {
    var v7 int16 = 32767
    v8 := GenS1{F0: 2, F1: 156}
    fmt.Println(v7)
    fmt.Println(v7)
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float64 = 1e10
    v3 := []int{-400}
    v4 := []int{-1, -726, 1, -573}
    var v5 bool = true
    v6 := make(map[int]bool)
    for fi9 := 0; fi9 < 12; fi9++ {
        fmt.Println(v5, v5)
        fmt.Println(v2)
        _ = fi9
    }
    v5 = v5
    v10 := GenS1{F0: 341, F1: 1}
    for fi11 := 0; fi11 < 17; fi11++ {
        v2 = v2
        v5 = v5
        fmt.Println(v2, v2, v5)
        _ = fi11
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        v2 = (3.14159 + v2)
        v3 = append(v3, fi12)
        _ = fi12
    }
    if v5 {
        switch -9223372036854775808 {
        case -716:
            v13 := v1[42]
            _ = v13
        case -1:
            v14 := GenS2{F0: 42}
            _ = v14
        }
        var v15 uint16 = 255
        v16 := GenS3{F0: 697, F1: 9223372036854775807}
        _ = v15
        _ = v16
    } else {
        v17 := GenS1{F0: 206, F1: 762}
        v2 -= float64(1.5)
        fmt.Println(v2)
        _ = v17
    }
    v6[127] = (!(42 == 14))
    for fi18 := 0; fi18 < 7; fi18++ {
        for fi19 := 0; fi19 < 11; fi19++ {
            var v20 uint8 = 61
            _ = fi19
            _ = v20
        }
        v5 = v5
        var v21 int64 = -9223372036854775808
        _ = fi18
        _ = v21
    }
    for fi22 := 0; fi22 < 3; fi22++ {
        fmt.Println(v2, v5)
        _ = fi22
    }
    v2 = (float64((((-2.25 - (1e10 + (0.0 / (0.1 * (3.14159 / 1.0))))) * 0.1) * 1e10)) + (v2 + v2))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
