package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8; F1 bool; F2 int16 }
func genFn0(p0 bool) int64 {
    p0 = p0
    var v4 float32 = 3.14159
    _ = p0
    _ = v4
    return 9223372036854775807
}

func main() {
    var v1 int8 = 127
    var v2 int = -773
    var v3 float64 = float64(1.5)
    v5 := GenS1{F0: v1, F1: 804, F2: 0}
    var v6 int16 = 1
    v2 = (v2 & 785)
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 float32 = ((1.5 + ((0.1 * ((1.5 / 3.14159) / (float32((0.0 - 1e10)) - (((0.0 / float32((v3 * (v3 + ((-2.25 - v3) - v3))))) / float32((float64(((0.0 - (1.5 / ((0.0 + -2.25) / 1.0))) + 1.5)) + (v3 / ((v3 * 3.14159) * v3))))) - 3.14159)))) / 1.5)) - 0.1)
        v1 += (v1 << ((255) & 63))
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := GenS2{F0: 1}
            if true {
                v1 += v1
                v11 := []int{}
                if (v1 <= (v1 ^ (v1 ^ v1))) {
                    var v12 uint32 = 234
                    if false {
                        fmt.Println(v8)
                    }
                    if true {
                        v1 -= v1
                    }
                    _ = v12
                } else {
                    if false {
                        switch "hello" {
                        case "":
                            v8 += ((v8 / v8) + v8)
                        default:
                            v13 := GenS3{F0: ("a" + "hello"), F1: 2147483647}
                            _ = v13
                        }
                        var v14 float32 = float32(((0.0 - ((v3 * float64(((v8 - float32((v3 / (v3 * float64(((v8 - 1.5) - 0.0)))))) / v8))) - 1e10)) - v3))
                        var v15 uint = 1
                        _ = v14
                        _ = v15
                    }
                    v16 := GenS3{F0: "a", F1: -2147483648}
                    _ = v16
                }
                _ = v11
            }
            v17 := GenS1{F0: 100, F1: 535, F2: 1}
            _ = fi9
            _ = v10
            _ = v17
        }
        _ = fi7
        _ = v8
    }
    fmt.Println(v3)
    v18 := []int{32768}
    v2 = v2
    v19 := GenS2{F0: 1}
    var v20 int16 = (v6 | 473)
    var v21 uint = 154
    var v22 int16 = (v20 / (((v20 + (-(v20)))) | 1))
    for fi23 := 0; fi23 < 10; fi23++ {
        v22 = (^(v6))
        fmt.Println(v3, v1, v21)
        _ = fi23
    }
    fmt.Println(v6)
    v24 := GenS2{F0: 127}
    fmt.Println(v3, v20, v2)
    v1 += v1
    for fi25 := 0; fi25 < 8; fi25++ {
        v26 := GenS4{F0: 2}
        var v27 uint8 = 1
        v28 := GenS3{F0: "foo bar", F1: 2147483647}
        _ = fi25
        _ = v26
        _ = v27
        _ = v28
    }
    switch "\n\t" {
    case "a":
        switch "hello" {
        case "\n\t":
            for fi29 := 0; fi29 < 14; fi29++ {
                switch v2 {
                case -871:
                    var v30 int32 = 42
                    _ = v30
                case -9223372036854775808:
                    var v31 int8 = v1
                    _ = v31
                default:
                    v20 += (v20 / ((v20) | 1))
                }
                _ = fi29
            }
        case "":
            var v32 uint16 = 1
            _ = v32
        case "hello":
            if (!false) {
                v1 += v1
            } else {
                var v33 uint64 = 93
                var v34 float64 = v3
                _ = v33
                _ = v34
            }
        default:
            v35 := GenS5{F0: -128, F1: false, F2: (v6 - (^(v22)))}
            _ = v35
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v5
    _ = v6
}
