package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int }
type GenS2 struct { F0 int; F1 uint8 }
type GenS3 struct { F0 uint16; F1 uint16; F2 int16 }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []int{32768, 709, -9223372036854775808, 1}
    v3 := make(map[string]string)
    v4 := make(map[string]uint32)
    var v5 float32 = 0.0
    v6 := []uint32{42, 0, 4294967295, 2147483647}
    for fi7 := 0; fi7 < 7; fi7++ {
        v3[""] = "a"
        for fi8 := 0; fi8 < 15; fi8++ {
            fmt.Println(v5, v5, v5)
            var v9 float64 = 0.0
            switch (fi8 * (fi8 & (-(fi7)))) {
            case 9223372036854775807:
                v10 := GenS1{F0: 131, F1: -1, F2: (fi7 * fi7)}
                _ = v10
            default:
                v1 += v1
            }
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v11 := GenS2{F0: 2147483647, F1: 2}
    var v12 uint32 = 328
    v1 = (v1 * -9223372036854775808)
    v13 := GenS2{F0: int(v12), F1: 1}
    v14 := GenS3{F0: 1, F1: 764, F2: -726}
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 int16 = 980
        var v17 string = "a"
        _ = fi15
        _ = v16
        _ = v17
    }
    v1 = (v1 ^ (v1 % (((v1 ^ (-(v1)))) | 1)))
    v18 := GenS3{F0: 1, F1: 373, F2: 114}
    var v19 int = 127
    var v20 uint8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
