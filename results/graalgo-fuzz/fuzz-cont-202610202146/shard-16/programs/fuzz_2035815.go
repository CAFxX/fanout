package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 bool; F1 int16; F2 string }
func genFn0(p0 uint32) uint64 {
    if (false && (-53 <= 42)) {
        p0 += 42
        p0 = 631
        p0 = p0
    } else {
        p0 += 42
        v5 := []int{4294967295, -22, 32768}
        _ = v5
    }
    var v6 uint32 = p0
    switch "a" {
    case "\n\t":
        var v7 float64 = float64(0.1)
        _ = v7
    }
    _ = p0
    _ = v6
    return 4294967295
}
func genFn1(p0 uint) int64 {
    fmt.Println(p0)
    if false {
        for fi8 := 0; fi8 < 12; fi8++ {
            for fi9 := 0; fi9 < 17; fi9++ {
                v10 := GenS1{F0: 42, F1: -263}
                var v11 float64 = (((float64(-2.25) + (1e10 * (-2.25 - ((1.5 + (1.5 - 1.5)) - (1e10 / (float64(1e10) / ((float64((0.1 / ((0.1 * 3.14159) / 1.0))) + float64((-2.25 * 3.14159))) + (((((((float64((0.1 + 1e10)) * 1.5) + 1e10) + float64(0.1)) - 3.14159) * (1e10 - 0.1)) / 1.5) * 0.1)))))))) * (1.5 * 0.1)) + 1.5)
                v11 -= float64(float32((float64(1e10) - 1.5)))
                v11 = (3.14159 + 1.5)
                v11 = v11
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi8
        }
        v12 := GenS1{F0: 42, F1: 635}
        switch ("" + ("\n\t" + "foo bar")) {
        case "foo bar":
            p0 -= (p0 ^ p0)
        case "a":
            var v13 float32 = float32(0.1)
            _ = v13
        default:
            v14 := GenS2{F0: false, F1: 42, F2: "a"}
            _ = v14
        }
        _ = v12
    }
    _ = p0
    return 1
}

func main() {
    var v1 float64 = 1e10
    var v2 float64 = -2.25
    var v3 float64 = (v2 / 3.14159)
    var v4 uint8 = 127
    var v15 float32 = 1.5
    var v16 uint8 = v4
    v3 = v2
    var v17 bool = (!true)
    v18 := []int{128}
    v16 = (v16 << ((255) & 63))
    for fi19 := 0; fi19 < 15; fi19++ {
        if ((!(42 <= 42)) || (v17 && (false && true))) {
            for fi20 := 0; fi20 < 4; fi20++ {
                switch (fi20 & fi19) {
                case -326:
                    fmt.Println(v4)
                case 199:
                    switch "" {
                    case "hello":
                        if (v17 && v17) {
                            if true {
                                v2 = (v1 * 0.1)
                                v21 := GenS1{F0: 9223372036854775807, F1: 589}
                                _ = v21
                            }
                            v15 = 1e10
                        }
                    case "foo bar":
                        v16 = v4
                    default:
                        v16 -= (v4 + v4)
                    }
                default:
                    var v22 int8 = -128
                    _ = v22
                }
                _ = fi20
            }
            v23 := GenS2{F0: false, F1: -974, F2: "a"}
            fmt.Println(v15)
            _ = v23
        }
        v16 = v16
        if false {
            if len(v18) > 0 {
                v18[0] = (fi19 ^ fi19)
            }
            if len(v18) > 0 {
                v18[0] = (fi19 / (((^(fi19))) | 1))
            }
            if len(v18) > 0 {
                v18[0] = fi19
            }
        }
        _ = fi19
    }
    if len(v18) > 0 {
        v18[0] = 32768
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
