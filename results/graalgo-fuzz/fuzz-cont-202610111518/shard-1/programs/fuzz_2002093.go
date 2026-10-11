package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint8; F1 int8; F2 int64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = true
    var v3 string = "\n\t"
    v4 := []uint32{709, 7}
    _ = len(v4)
    fmt.Println(v2, v2)
    v2 = (v2 || v2)
    var v5 float64 = ((0.0 / 1e10) - 1.5)
    if v2 {
        for fi6 := 0; fi6 < 6; fi6++ {
            for fi7 := 0; fi7 < 15; fi7++ {
                v8 := GenS1{F0: 1, F1: 2}
                _ = fi7
                _ = v8
            }
            _ = fi6
        }
        var v9 uint8 = 0
        _ = v9
    }
    v4 = append(v4, 552)
    v10 := GenS1{F0: 42, F1: 540}
    v11 := GenS2{F0: v3, F1: 65535}
    v12 := GenS1{F0: 9223372036854775807, F1: 18446744073709551615}
    delete(v1, -956)
    var v13 bool = (!(-1 > 0))
    if (-1 != -1) {
        if (-729 <= 1000) {
            var v14 float32 = (((((-2.25 / 0.1) + ((3.14159 * 0.0) / -2.25)) - 3.14159) + 0.0) * 3.14159)
            var v15 int = 42
            fmt.Println(v3, v2)
            _ = v14
            _ = v15
        }
        v16 := GenS1{F0: 42, F1: 65535}
        _ = v16
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        _ = len(v1)
        _ = fi17
    }
    switch -811 {
    case 128:
        v18 := GenS3{F0: 220, F1: -1, F2: -9223372036854775808}
        _ = v18
    default:
        fmt.Println(v13)
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        v5 = v5
        _ = fi19
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        v13 = ((v3 + v3) <= (v3 + v3))
        _ = fi20
    }
    var v21 string = "hello"
    for fi22 := 0; fi22 < 4; fi22++ {
        v4 = append(v4, 2)
        v21 = "a"
        _ = fi22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
