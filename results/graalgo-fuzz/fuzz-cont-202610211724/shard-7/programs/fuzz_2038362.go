package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int32 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 float64 = 0.0
    var v2 float32 = 1e10
    var v3 bool = (!((false || false) || ((v1 / -2.25) == float64(v2))))
    var v4 int16 = -578
    var v5 int64 = -9223372036854775808
    v6 := make(map[int]bool)
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            v1 = (v1 / float64(v2))
            v6[fi8] = v3
            _ = fi8
        }
        if ((3.14159 + v1) <= 1e10) {
            fmt.Println(v1)
            v4 = v4
        }
        fmt.Println(fi7, v4, fi7)
        _ = fi7
    }
    v2 -= (((v2 * (1.5 - 0.1)) / 1.0) - v2)
    switch "a" {
    case "\n\t":
        v9 := GenS1{F0: v4, F1: (v4 + v4), F2: 112}
        _ = v9
    default:
        v10 := GenS1{F0: v4, F1: v4, F2: -468}
        _ = v10
    }
    v4 = v4
    if v3 {
        v11 := []int{128}
        var v12 int8 = 0
        v11 = append(v11, -496)
        _ = v11
        _ = v12
    }
    if (42 != v5) {
        if (((v2 - (v2 * (v2 - v2))) - v2) == (v2 * v2)) {
            var v13 uint64 = 1
            for fi14 := 0; fi14 < 2; fi14++ {
                var v15 int8 = 0
                if v3 {
                    v16 := []int{49, -92}
                    var v17 uint64 = 18446744073709551615
                    fmt.Println(v1)
                    _ = v16
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 3; fi18++ {
                        var v19 uint32 = 42
                        _ = fi18
                        _ = v19
                    }
                }
                _ = fi14
                _ = v15
            }
            _ = v13
        }
    }
    if v3 {
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 string = ("" + "a")
            _ = fi20
            _ = v21
        }
    }
    if v3 {
        if true {
            v2 = -2.25
            v22 := GenS1{F0: v4, F1: v4, F2: 42}
            _ = v22
        } else {
            delete(v6, -9223372036854775808)
            var v23 uint16 = 1
            if v3 {
                var v24 uint64 = 0
                var v25 uint32 = 58
                _ = v24
                _ = v25
            } else {
                v4 -= (v4 / ((v4) | 1))
                var v26 uint = 1
                _ = v26
            }
            _ = v23
        }
        v27 := GenS2{F0: 430, F1: -14}
        v4 = v4
        _ = v27
    }
    if false {
        v28 := GenS3{F0: -63}
        fmt.Println(v2)
        _ = v28
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
