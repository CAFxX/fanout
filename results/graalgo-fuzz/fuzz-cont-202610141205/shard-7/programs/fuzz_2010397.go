package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int32; F1 int8; F2 int16 }
func genFn0(p0 int, p1 int8) uint {
    var v4 uint = 2
    for fi5 := 0; fi5 < 1; fi5++ {
        if (!(42 <= 131)) {
            var v6 uint = (v4 * 779)
            v6 = (v6 * v6)
            _ = v6
        } else {
            fmt.Println(p1, v4)
        }
        switch -2147483648 {
        case 4294967295:
            switch (fi5 + -9223372036854775808) {
            case 9223372036854775807:
                v4 += 4294967295
            case -247:
                fmt.Println(p1, fi5)
            default:
                for fi7 := 0; fi7 < 1; fi7++ {
                    v8 := []int{(p0 % (((fi7 % ((fi5) | 1))) | 1)), fi7, fi5}
                    var v9 uint32 = 788
                    switch fi5 {
                    case 1:
                        fmt.Println(fi5, v4, fi5)
                    case 128:
                        var v10 uint = 0
                        _ = v10
                    default:
                        switch "hello" {
                        case "hello":
                            v11 := GenS1{F0: 887, F1: (fi7 & 127)}
                            _ = v11
                        case "a":
                            p1 -= 0
                        }
                    }
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            }
        }
        fmt.Println(v4)
        _ = fi5
    }
    v12 := GenS1{F0: 0, F1: p0}
    var v13 int32 = -2147483648
    if (255 != 255) {
        var v14 int64 = -9223372036854775808
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    _ = v4
    return v4
}

func main() {
    var v1 int64 = 1
    var v2 int = 4294967295
    var v3 int8 = 100
    v15 := GenS1{F0: 633, F1: (^(v2))}
    v16 := GenS1{F0: -830, F1: v2}
    v2 -= (v2 * (^(v2)))
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 uint8 = 128
        _ = fi17
        _ = v18
    }
    if (true || (("a" != "") && (0.0 > 3.14159))) {
        v19 := GenS2{F0: 0, F1: -106, F2: 0}
        var v20 float64 = -2.25
        _ = v19
        _ = v20
    } else {
        v2 = (v2 - v2)
        v1 = (v1 - v1)
    }
    v1 = (-(v1))
    fmt.Println(v1, v2, v2)
    var v21 uint = 42
    fmt.Println(v1, v21, v3)
    var v22 uint8 = 127
    var v23 uint = (v21 / (((^(v21))) | 1))
    for fi24 := 0; fi24 < 15; fi24++ {
        v21 -= v23
        _ = fi24
    }
    v22 = (^(v22))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
