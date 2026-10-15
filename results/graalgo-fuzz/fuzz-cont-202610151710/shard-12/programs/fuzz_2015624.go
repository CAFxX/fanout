package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint32 }
type GenS2 struct { F0 int8; F1 int8; F2 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 int64; F1 uint64; F2 uint }
func genFn0(p0 uint64, p1 int64, p2 int32) uint8 {
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := GenS1{F0: 0, F1: (p1 / (((p1 | (-(p1)))) | 1)), F2: 2147483647}
        _ = fi7
        _ = v8
    }
    p2 = (p2 + -174)
    _ = p0
    _ = p1
    _ = p2
    return 2
}
func genFn1(p0 uint8) int32 {
    p0 -= 42
    p0 = p0
    fmt.Println(p0)
    _ = p0
    return 851
}

func main() {
    var v1 uint8 = 241
    var v2 int = 128
    var v3 int32 = -154
    var v4 uint64 = 755
    var v5 float64 = 0.1
    v6 := []uint32{4294967295, 37, 2147483647}
    var v9 uint16 = 32767
    v1 = v1
    var v10 bool = true
    v11 := GenS2{F0: 57, F1: 42, F2: 65535}
    var v12 uint32 = 65535
    fmt.Println(v2, v10)
    for fi13 := 0; fi13 < 15; fi13++ {
        switch (-(v2)) {
        case -480:
            fmt.Println(v10, v4, v1)
        case -9223372036854775808:
            if len(v6) > 0 {
                v6[0] = uint32(v4)
            }
        case 975:
            _ = len(v6)
        }
        for fi14 := 0; fi14 < 11; fi14++ {
            if (((2 <= uint8(v9)) || true) || (-2.25 <= (((-2.25 + 1.5) / 0.1) - (0.0 + 1.5)))) {
                switch 2147483647 {
                case 4294967295:
                    v15 := GenS1{F0: -128, F1: -1, F2: 0}
                    _ = v15
                case 9223372036854775807:
                    if true {
                        v12 = (v12 ^ (v12 + v12))
                        v10 = v10
                    } else {
                        var v16 uint32 = (v12 - v12)
                        _ = v16
                    }
                default:
                    v17 := GenS2{F0: -128, F1: 1, F2: 523}
                    _ = v17
                }
                var v18 uint8 = 255
                _ = v18
            } else {
                if v10 {
                    var v19 uint32 = (^(v12))
                    var v20 int32 = v3
                    _ = v19
                    _ = v20
                } else {
                    if false {
                        v10 = v10
                        v12 += (v12 << ((1) & 63))
                        if (v12 != (v12 ^ v12)) {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            switch "hello" {
                            case "a":
                                fmt.Println(v2, v9)
                            case "\n\t":
                                v21 := GenS3{F0: "foo bar"}
                                _ = v21
                            case "":
                                if len(v6) > 0 {
                                    v6[0] = (^(v12))
                                }
                            default:
                                v22 := GenS4{F0: fi14}
                                _ = v22
                            }
                        }
                    } else {
                        v9 = v9
                        v3 = -2147483648
                    }
                    if (false && ((v4 > (v4 + 42)) && v10)) {
                        fmt.Println(v12, v10, v1)
                        v3 = (v3 + v3)
                    }
                }
                v4 -= (v4 | v4)
            }
            v23 := GenS5{F0: 1, F1: -80}
            _ = fi14
            _ = v23
        }
        _ = fi13
    }
    if ((v3 / (((v3 / ((v3) | 1))) | 1)) == (v3 / (((v3 % ((v3) | 1))) | 1))) {
        if (v12 >= (^(v12))) {
            switch (^(v2)) {
            case 101:
                switch "a" {
                case "hello":
                    if v10 {
                        v12 = v12
                        if len(v6) > 0 {
                            v6[0] = 4294967295
                        }
                    }
                case "foo bar":
                    v2 = v2
                default:
                    v10 = ((v2 / (((v2 + v2)) | 1)) < 2147483647)
                }
            case 579:
                fmt.Println(v9, v4)
            default:
                switch (("\n\t" + "hello") + "") {
                case "":
                    v4 -= (v4 ^ (^(v4)))
                case "a":
                    for fi24 := 0; fi24 < 14; fi24++ {
                        switch "foo bar" {
                        case "foo bar":
                            _ = len(v6)
                        default:
                            var v25 int16 = -32768
                            _ = v25
                        }
                        _ = fi24
                    }
                case "hello":
                    var v26 float64 = (float64(((((((-2.25 - 3.14159) + 0.1) * 1.5) * (0.1 * 3.14159)) * float32((1e10 * 0.1))) + ((((-2.25 - ((0.0 + (1.5 * 3.14159)) + 0.1)) * (1e10 / -2.25)) * -2.25) * ((float32(((v5 - v5) / v5)) / 1.0) * 0.1)))) - -2.25)
                    _ = v26
                default:
                    var v27 int8 = -47
                    _ = v27
                }
            }
        } else {
            if (610 == -1) {
                var v28 bool = v10
                v3 = genFn1(v1)
                _ = v28
            } else {
                var v29 float32 = -2.25
                v30 := GenS6{F0: 42, F1: (^(v4)), F2: 0}
                _ = v29
                _ = v30
            }
        }
        fmt.Println(v4, v12)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    _ = len(v6)
    v12 -= v12
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
