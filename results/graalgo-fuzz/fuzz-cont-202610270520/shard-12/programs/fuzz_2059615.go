package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int; F1 string; F2 bool }

func main() {
    v1 := []uint32{0, 4294967295}
    var v2 int32 = -2147483648
    v3 := make(map[int]bool)
    var v4 int32 = (v2 - v2)
    var v5 int64 = 703
    v1 = append(v1, 0)
    v6 := GenS1{F0: 148, F1: v5}
    var v7 float32 = 0.0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v8 uint8 = 42
    var v9 string = "foo bar"
    switch v9 {
    case "\n\t":
        var v10 string = v9
        _ = v10
    case "foo bar":
        if (!false) {
            var v11 float64 = ((((-2.25 + 3.14159) / 1e10) * 0.0) * (((float64((-2.25 * 3.14159)) + (-2.25 + (((-2.25 + 0.0) + 1e10) - (3.14159 + 0.1)))) / ((float64(-2.25) * 3.14159) * float64((v7 + v7)))) - 0.0))
            _ = len(v1)
            v8 -= (^(v8))
            _ = v11
        } else {
            v5 += 1
        }
    case "":
        fmt.Println(v5, v7)
    default:
        v12 := GenS1{F0: 42, F1: v5}
        _ = v12
    }
    if true {
        v13 := GenS2{F0: 161, F1: v9, F2: true}
        var v14 bool = true
        switch (v9 + v9) {
        case "a":
            for fi15 := 0; fi15 < 20; fi15++ {
                v1 = append(v1, 0)
                v2 = v4
                v16 := GenS1{F0: 0, F1: (v5 * v5)}
                _ = fi15
                _ = v16
            }
        default:
            v17 := GenS1{F0: 1, F1: v5}
            _ = v17
        }
        _ = v13
        _ = v14
    } else {
        var v18 int8 = 100
        var v19 float64 = 1.5
        var v20 uint8 = (v8 / ((v8) | 1))
        _ = v18
        _ = v19
        _ = v20
    }
    v8 = (v8 & (v8 | 248))
    for fi21 := 0; fi21 < 13; fi21++ {
        var v22 uint16 = 42
        switch (v9 + v9) {
        case "hello":
            var v23 uint = 42
            _ = v23
        case "a":
            switch (fi21 | (fi21 >> ((255) & 63))) {
            case 1:
                for fi24 := 0; fi24 < 14; fi24++ {
                    switch -32 {
                    case 32768:
                        fmt.Println(v2, fi21)
                    case 2147483647:
                        if len(v1) > 0 {
                            v1[0] = 2147483648
                        }
                    case -714:
                        fmt.Println(v22)
                    default:
                        fmt.Println(fi21)
                    }
                    _ = fi24
                }
            case 127:
                if len(v1) > 0 {
                    v1[0] = 233
                }
            }
        }
        if (false || true) {
            var v25 int32 = (-(v2))
            _ = v25
        }
        _ = fi21
        _ = v22
    }
    v8 = (v8 ^ (v8 * 88))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
