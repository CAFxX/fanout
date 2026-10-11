package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint32 }
type GenS2 struct { F0 int }
func genFn0(p0 uint32) uint {
    fmt.Println(p0)
    v6 := []int{2147483647, -775}
    var v7 int16 = 767
    fmt.Println(p0, p0)
    for fi8 := 0; fi8 < 7; fi8++ {
        fmt.Println(v7)
        _ = fi8
    }
    _ = p0
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint8, p1 uint32, p2 int64) int32 {
    p0 = (p0 - p0)
    for fi9 := 0; fi9 < 7; fi9++ {
        fmt.Println(p0, p0, p0)
        for fi10 := 0; fi10 < 1; fi10++ {
            if (1e10 < 1.5) {
                var v11 uint16 = 65535
                v12 := GenS1{F0: "hello", F1: p0, F2: (p1 / ((p1) | 1))}
                fmt.Println(p1, fi9)
                _ = v11
                _ = v12
            }
            var v13 uint32 = p1
            if (v13 > (p1 - v13)) {
                v14 := GenS1{F0: ("a" + "\n\t"), F1: 115, F2: 42}
                _ = v14
            } else {
                p0 = (p0 ^ (p0 >> ((0) & 63)))
            }
            _ = fi10
            _ = v13
        }
        _ = fi9
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        v16 := GenS2{F0: 127}
        _ = fi15
        _ = v16
    }
    switch 1 {
    case 42:
        p2 -= (p2 / ((p2) | 1))
    }
    _ = p0
    _ = p1
    _ = p2
    return 2147483647
}

func main() {
    var v1 float64 = float64((((3.14159 * (1.5 - (-2.25 + (float32(-2.25) - (3.14159 / 0.1))))) + ((3.14159 * (1.5 / 1.0)) / 0.1)) * (3.14159 / float32(1e10))))
    var v2 uint64 = 18446744073709551615
    v3 := []int{2147483647, 9223372036854775807, -2147483648, 0}
    var v4 string = "\n\t"
    v5 := make(map[string]string)
    v17 := GenS1{F0: (v4 + v4), F1: 199, F2: 470}
    for fi18 := 0; fi18 < 15; fi18++ {
        v4 = v4
        var v19 uint32 = 65535
        for fi20 := 0; fi20 < 8; fi20++ {
            if false {
                v19 = (^(v19))
                var v21 int64 = 42
                _ = v21
            }
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    for fi22 := 0; fi22 < 1; fi22++ {
        var v23 int = (fi22 | fi22)
        v24 := GenS2{F0: (fi22 ^ v23)}
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 int64 = -586
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    var v26 string = "a"
    for fi27 := 0; fi27 < 9; fi27++ {
        var v28 uint8 = 202
        _ = fi27
        _ = v28
    }
    v25 = 0
    v4 = v26
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
