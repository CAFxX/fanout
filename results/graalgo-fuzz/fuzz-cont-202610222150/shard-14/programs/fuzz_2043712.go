package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint8; F1 string; F2 uint64 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64; F1 uint16; F2 bool }
type GenS6 struct { F0 uint; F1 int64 }
type GenS7 struct { F0 uint16; F1 uint }
func genFn0(p0 int8, p1 uint64) uint16 {
    if (42 <= 1) {
        var v5 uint32 = 2147483647
        _ = v5
    }
    var v6 int16 = 42
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            if true {
                var v9 uint32 = 37
                var v10 uint64 = p1
                if ((false || (!true)) && false) {
                    v9 = (v9 & v9)
                    var v11 uint = 2
                    _ = v11
                } else {
                    v12 := []int{}
                    fmt.Println(v6)
                    _ = v12
                }
                _ = v10
                _ = v9
            }
            _ = fi8
        }
        _ = fi7
    }
    p0 = (p0 & p0)
    var v13 uint8 = 128
    _ = p0
    _ = p1
    _ = v13
    _ = v6
    return 884
}

func main() {
    var v1 uint32 = 1
    var v2 int64 = -1
    var v3 uint32 = v1
    v4 := make(map[int]int)
    v14 := GenS1{F0: 876, F1: ("\n\t" + (("a" + ("" + "hello")) + "\n\t"))}
    switch "a" {
    case "":
        v15 := GenS2{F0: 33, F1: "", F2: 444}
        _ = v15
    case "hello":
        if true {
            for fi16 := 0; fi16 < 17; fi16++ {
                var v17 int32 = -1
                _ = fi16
                _ = v17
            }
            v1 = 1
        }
    case "foo bar":
        v18 := GenS3{F0: -9223372036854775808, F1: 0}
        _ = v18
    default:
        v19 := []int{}
        _ = v19
    }
    v2 += (v2 ^ v2)
    v20 := []int{192, 653, 42}
    for fi21 := 0; fi21 < 10; fi21++ {
        v22 := GenS3{F0: -9223372036854775808, F1: 2}
        if (1.5 == (1e10 * 3.14159)) {
            v23 := GenS4{F0: false}
            _ = v23
        } else {
            v24 := v4[fi21]
            _ = v24
        }
        _ = fi21
        _ = v22
    }
    v1 -= (v3 ^ (v3 & (^(v1))))
    v25 := GenS3{F0: -449, F1: 2}
    for fi26 := 0; fi26 < 3; fi26++ {
        var v27 float32 = 1.5
        v2 += (v2 - v2)
        _ = fi26
        _ = v27
    }
    v28 := GenS5{F0: 87, F1: 42, F2: (!(((v1 ^ (v3 - v1)) < v1) || (((true || ((3.14159 < float32(-2.25)) || true)) || false) || true)))}
    var v29 uint = 18446744073709551615
    var v30 uint8 = 217
    if true {
        v31 := GenS6{F0: 467, F1: v2}
        v32 := GenS7{F0: 42, F1: (v29 >> ((2) & 63))}
        _ = v31
        _ = v32
    }
    v2 += (v2 / ((v2) | 1))
    var v33 int = -2147483648
    v1 = (^(v3))
    v3 += v3
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v4
}
