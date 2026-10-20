package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint32; F2 int }

func main() {
    v1 := make(map[int]int)
    var v2 int64 = -192
    v3 := make(map[string]string)
    var v4 int = -661
    v5 := []string{}
    var v6 uint = 42
    fmt.Println(v4, v6, v4)
    var v7 uint32 = 42
    fmt.Println(v7, v6)
    v2 -= v2
    v1[127] = (-(v4))
    var v8 int64 = (v2 ^ v2)
    fmt.Println(v4)
    var v9 uint64 = 963
    switch (v4 * v4) {
    case 128:
        v10 := GenS1{F0: 0, F1: 692, F2: 1}
        _ = v10
    default:
        if true {
            var v11 int = v4
            _ = v11
        } else {
            if true {
                switch (v4 << ((v6) & 63)) {
                case 9223372036854775807:
                    fmt.Println(v2, v6)
                default:
                    v12 := GenS1{F0: -32768, F1: -648, F2: 0}
                    _ = v12
                }
                v13 := GenS1{F0: 32767, F1: 483, F2: 0}
                v8 = (v2 & (v8 & int64(v6)))
                _ = v13
            }
        }
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        for fi15 := 0; fi15 < 19; fi15++ {
            switch "foo bar" {
            case "foo bar":
                v6 = v6
            case "a":
                if true {
                    var v16 uint64 = v9
                    _ = v16
                }
                var v17 float64 = ((-2.25 + (((1e10 + 0.1) + float64(0.1)) / 1.0)) * (1.5 - 0.0))
                var v18 int = (fi14 & 9223372036854775807)
                _ = v17
                _ = v18
            }
            _ = fi15
        }
        _ = fi14
    }
    var v19 uint = (v6 % ((490) | 1))
    fmt.Println(v19, v19)
    switch "a" {
    case "":
        fmt.Println(v2, v7, v2)
    default:
        v20 := GenS2{F0: 42, F1: 1, F2: -940}
        _ = v20
    }
    var v21 float32 = 0.0
    v2 -= v8
    for fi22 := 0; fi22 < 3; fi22++ {
        for fi23 := 0; fi23 < 1; fi23++ {
            v24 := v1[(fi23 ^ (v4 / (((fi23 >> ((565) & 63))) | 1)))]
            _ = fi23
            _ = v24
        }
        _ = fi22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
