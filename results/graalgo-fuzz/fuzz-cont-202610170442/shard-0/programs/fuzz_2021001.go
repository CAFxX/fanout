package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 int16; F2 uint8 }

func main() {
    var v1 bool = false
    var v2 uint8 = 128
    v3 := []int{9223372036854775807, 42, 128}
    var v4 float32 = (3.14159 * 0.1)
    var v5 uint8 = v2
    var v6 uint16 = 0
    for fi7 := 0; fi7 < 15; fi7++ {
        _ = len(v3)
        _ = fi7
    }
    switch "" {
    case "":
        switch -113 {
        case -9223372036854775808:
            var v8 int32 = -871
            _ = v8
        }
    default:
        v4 += ((v4 / float32(float64((((((((v4 - ((1e10 / v4) + v4)) * v4) / v4) * ((-2.25 - float32(float64((1e10 - v4)))) * float32(1e10))) * (((((v4 - float32(0.0)) * (v4 + ((1e10 / 1.0) / (1e10 / (v4 * float32(0.1)))))) - (1.5 - 0.0)) / ((v4 - v4) + v4)) / (1e10 * (3.14159 + (((1.5 / (v4 - v4)) / (v4 - (1e10 - v4))) / v4))))) + (float32((1e10 + -2.25)) / v4)) + (0.1 / -2.25))))) - ((v4 * v4) - (0.0 - (v4 / (v4 - v4)))))
    }
    fmt.Println(v6)
    v6 = 32767
    if ((v6 * 0) <= v6) {
        v6 -= 0
        var v9 uint = 2
        var v10 float32 = ((0.1 * 0.0) / (v4 / 0.1))
        _ = v10
        _ = v9
    } else {
        for fi11 := 0; fi11 < 4; fi11++ {
            if len(v3) > 0 {
                v3[0] = fi11
            }
            _ = fi11
        }
    }
    if ((!(v1 && (1 == 255))) && v1) {
        var v12 uint16 = (v6 / ((v6) | 1))
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v12
    } else {
        fmt.Println(v2)
    }
    v13 := GenS1{F0: 100}
    v2 = (v2 & (v2 << ((18446744073709551615) & 63)))
    var v14 float32 = (((v4 - v4) / (((((-2.25 * v4) - v4) - v4) + (1e10 / 0.1)) / ((v4 - v4) / 1e10))) / (v4 + 1e10))
    _ = len(v3)
    var v15 int8 = -13
    if v1 {
        switch "\n\t" {
        case "\n\t":
            v16 := GenS1{F0: int8(v6)}
            _ = v16
        case "":
            if v1 {
                switch 127 {
                case -1:
                    if v1 {
                        v17 := GenS1{F0: (^(v15))}
                        _ = v17
                    }
                case 114:
                    if len(v3) > 0 {
                        v3[0] = int(v6)
                    }
                case 817:
                    v5 = 0
                default:
                    switch "a" {
                    case "":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                }
            }
        default:
            v3 = append(v3, 32768)
        }
        v1 = v1
        v14 += v14
    }
    v2 -= 139
    var v18 int8 = v15
    if (v1 || v1) {
        v6 += (^(v6))
        var v19 uint32 = 786
        var v20 float32 = 0.0
        _ = v19
        _ = v20
    } else {
        v5 = v5
        for fi21 := 0; fi21 < 8; fi21++ {
            fmt.Println(v1, v4, v14)
            _ = fi21
        }
    }
    if ((^(v6)) == v6) {
        v22 := GenS2{F0: -9223372036854775808, F1: 1000, F2: (v5 << ((317) & 63))}
        v3 = append(v3, -830)
        if (v1 && true) {
            if ("hello" > ("hello" + "hello")) {
                switch "foo bar" {
                case "a":
                    for fi23 := 0; fi23 < 1; fi23++ {
                        v14 = 1e10
                        for fi24 := 0; fi24 < 9; fi24++ {
                            var v25 uint32 = 2
                            _ = fi24
                            _ = v25
                        }
                        var v26 int32 = -2147483648
                        _ = fi23
                        _ = v26
                    }
                }
            }
        }
        _ = v22
    } else {
        v27 := GenS2{F0: -164, F1: 424, F2: (^(v2))}
        fmt.Println(v5, v1)
        fmt.Println(v1, v2, v1)
        _ = v27
    }
    v28 := GenS2{F0: -1, F1: -32768, F2: 255}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
