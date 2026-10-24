package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 uint32; F2 int32 }
func genFn0(p0 bool) int8 {
    p0 = p0
    p0 = true
    var v6 float64 = 0.0
    v7 := []int{-9223372036854775808}
    fmt.Println(p0)
    _ = p0
    _ = v6
    _ = v7
    return 127
}

func main() {
    var v1 float64 = (-2.25 / 1e10)
    v2 := []uint32{42, 2, 228}
    var v3 int16 = 498
    v4 := make(map[int]int)
    v5 := make(map[string]bool)
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 uint8 = 1
        if (false && true) {
            var v10 float32 = 0.0
            switch fi8 {
            case 1:
                var v11 int64 = 9223372036854775807
                _ = v11
            default:
                var v12 int16 = 42
                _ = v12
            }
            var v13 int16 = (^(v3))
            _ = v10
            _ = v13
        } else {
            var v14 uint64 = 668
            v9 = v9
            _ = len(v4)
            _ = v14
        }
        _ = fi8
        _ = v9
    }
    switch "a" {
    case "hello":
        v15 := GenS1{F0: 255, F1: 65535, F2: 218}
        _ = v15
    case "\n\t":
        if true {
            v3 = v3
        } else {
            var v16 string = "\n\t"
            _ = v16
        }
    case "a":
        v1 = ((v1 * 1.5) / v1)
    }
    v3 -= 98
    v17 := GenS1{F0: 0, F1: 65535, F2: 1}
    if (!false) {
        v18 := GenS2{F0: 0}
        v3 -= 42
        _ = v18
    } else {
        fmt.Println(v1, v3)
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        for fi20 := 0; fi20 < 20; fi20++ {
            v1 = ((v1 * 1e10) / (0.0 - v1))
            _ = fi20
        }
        _ = len(v4)
        _ = fi19
    }
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 bool = ((true || ((fi21 ^ 127) <= (fi21 / ((-2147483648) | 1)))) || false)
        _ = len(v4)
        var v23 string = "\n\t"
        _ = fi21
        _ = v22
        _ = v23
    }
    var v24 string = ""
    if (!(0 <= 1)) {
        v2 = append(v2, 42)
        if false {
            for fi25 := 0; fi25 < 12; fi25++ {
                var v26 int32 = 0
                switch (fi25 << ((18446744073709551615) & 63)) {
                case 267:
                    fmt.Println(v3, fi25)
                case 9223372036854775807:
                    var v27 float32 = (1e10 / (0.0 - 3.14159))
                    _ = v27
                default:
                    var v28 float64 = v1
                    _ = v28
                }
                switch v24 {
                case "foo bar":
                    v29 := GenS1{F0: 1, F1: 0, F2: 128}
                    _ = v29
                }
                _ = fi25
                _ = v26
            }
        }
    }
    for fi30 := 0; fi30 < 19; fi30++ {
        v31 := GenS3{F0: ("" + (("\n\t" + "foo bar") + ("hello" + (v24 + "hello")))), F1: 889, F2: -228}
        v1 = 1.5
        if false {
            v5[v24] = false
        }
        _ = fi30
        _ = v31
    }
    var v32 uint8 = 188
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
