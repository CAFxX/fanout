package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 uint }
type GenS4 struct { F0 uint }
func genFn0(p0 int16, p1 int16, p2 uint64) uint16 {
    v7 := GenS1{F0: 2, F1: 65535, F2: (p1 + -19)}
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 string = "foo bar"
        _ = fi8
        _ = v9
    }
    p2 += 18446744073709551615
    if ((false || (!true)) && true) {
        p1 = 1000
        p0 = (p0 % ((1000) | 1))
    } else {
        var v10 uint32 = 2147483647
        _ = v10
    }
    switch -160 {
    case -9223372036854775808:
        if true {
            v11 := GenS2{F0: 550}
            if true {
                var v12 int8 = 42
                var v13 uint = 460
                _ = v12
                _ = v13
            }
            _ = v11
        }
    default:
        p2 = (^(p2))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 572
}

func main() {
    var v1 int = 32767
    var v2 string = ""
    var v3 int = 32768
    var v4 uint64 = 2
    var v5 float64 = ((((0.0 * (((float64(0.1) + (((((3.14159 + -2.25) - (1e10 / (((3.14159 + 0.1) / 3.14159) + 0.1))) * (1.5 / 1.0)) / 1.5) / 1.5)) / (0.1 / 1.0)) - 0.1)) + 1.5) + -2.25) * ((1.5 * 0.0) - 0.0))
    var v6 float64 = v5
    v14 := GenS3{F0: 72, F1: 0}
    if (-814 > 42) {
        v15 := GenS3{F0: 175, F1: 795}
        var v16 float32 = -2.25
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        v18 := GenS3{F0: 153, F1: 588}
        _ = fi17
        _ = v18
    }
    v19 := GenS4{F0: 255}
    switch v2 {
    case "\n\t":
        for fi20 := 0; fi20 < 10; fi20++ {
            fmt.Println(v4, fi20, fi20)
            _ = fi20
        }
    }
    var v21 int32 = 2147483647
    v22 := []int{}
    if true {
        switch -9223372036854775808 {
        case 1:
            fmt.Println(v1, v2, v4)
        case -92:
            for fi23 := 0; fi23 < 16; fi23++ {
                v24 := GenS2{F0: -439}
                v25 := GenS4{F0: 42}
                _ = fi23
                _ = v24
                _ = v25
            }
        }
    } else {
        if len(v22) >= 2 {
            v22 = v22[1:len(v22)]
        }
        var v26 int16 = 0
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
