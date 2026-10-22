package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }

func main() {
    var v1 uint32 = 65535
    var v2 int32 = 478
    v3 := []int{994}
    v4 := make(map[uint32]string)
    v5 := make(map[int]string)
    var v6 float64 = (float64((1.5 - float32(((1e10 + (1e10 - 3.14159)) - (float64(3.14159) / 1.5))))) * 1e10)
    for fi7 := 0; fi7 < 6; fi7++ {
        if false {
            v2 += (v2 + 2147483647)
            for fi8 := 0; fi8 < 17; fi8++ {
                if ((v2 & v2) <= -581) {
                    var v9 string = ""
                    v10 := GenS1{F0: -1, F1: 255}
                    v6 = v6
                    _ = v10
                    _ = v9
                }
                var v11 uint = 255
                if false {
                    v12 := v5[0]
                    v1 = (^(v1))
                    _ = v12
                } else {
                    var v13 string = "a"
                    if false {
                        var v14 bool = false
                        v6 = 1e10
                        _ = v14
                    }
                    var v15 int8 = -128
                    _ = v13
                    _ = v15
                }
                _ = fi8
                _ = v11
            }
        } else {
            v1 -= (v1 % (((v1 << ((65535) & 63))) | 1))
        }
        for fi16 := 0; fi16 < 20; fi16++ {
            var v17 int = fi16
            switch "foo bar" {
            case "\n\t":
                v18 := GenS1{F0: -1, F1: 437}
                _ = v18
            case "":
                fmt.Println(fi16, v2)
            case "hello":
                var v19 int8 = 127
                _ = v19
            default:
                v17 -= -387
            }
            var v20 uint64 = 2
            _ = fi16
            _ = v17
            _ = v20
        }
        _ = fi7
    }
    switch 127 {
    case 4294967295:
        v4[(v1 << ((18446744073709551615) & 63))] = "\n\t"
    case 200:
        if true {
            if len(v3) > 0 {
                v3[0] = 127
            }
            v21 := GenS1{F0: 9223372036854775807, F1: 65535}
            var v22 uint8 = 118
            _ = v21
            _ = v22
        }
    }
    v4[v1] = "hello"
    for fi23 := 0; fi23 < 2; fi23++ {
        v6 = ((3.14159 - v6) + 1e10)
        v1 = 2147483647
        _ = fi23
    }
    v24 := GenS1{F0: 177, F1: 870}
    fmt.Println(v6)
    v25 := v5[-36]
    var v26 int16 = -32768
    fmt.Println(v6, v6)
    var v27 float32 = (-2.25 + ((0.0 + (3.14159 - 1e10)) + -2.25))
    var v28 uint32 = v1
    fmt.Println(v2, v2, v28)
    if (128 == 1) {
        for fi29 := 0; fi29 < 11; fi29++ {
            fmt.Println(fi29)
            _ = fi29
        }
    }
    v30 := GenS1{F0: 9223372036854775807, F1: 973}
    switch (v25 + v25) {
    case "\n\t":
        fmt.Println(v27, v28, v6)
    case "":
        v1 += (v28 | v1)
    case "hello":
        v31 := GenS1{F0: 1, F1: 18446744073709551615}
        _ = v31
    default:
        var v32 string = "a"
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
