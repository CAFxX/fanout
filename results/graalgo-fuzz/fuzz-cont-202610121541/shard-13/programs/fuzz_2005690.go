package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint8; F1 uint64; F2 int32 }
type GenS3 struct { F0 int16; F1 string }

func main() {
    var v1 uint32 = 42
    v2 := make(map[int]string)
    v3 := []uint32{}
    v4 := []int{2147483647, 32767, -1}
    v5 := make(map[string]bool)
    var v6 int64 = -1
    if ((32768 > 9223372036854775807) && (2 != 255)) {
        var v7 uint16 = 0
        if len(v4) > 0 {
            v4[0] = 127
        }
        _ = v7
    }
    fmt.Println(v6, v1)
    v6 = (v6 + v6)
    switch 32768 {
    case 42:
        v1 = (v1 % ((v1) | 1))
    default:
        var v8 int8 = -128
        _ = v8
    }
    v3 = append(v3, 2147483648)
    v6 -= v6
    v9 := GenS1{F0: (255 > 0), F1: 1, F2: 0}
    for fi10 := 0; fi10 < 12; fi10++ {
        switch "foo bar" {
        case "hello":
            v1 -= 563
        case "a":
            v1 = (v1 / ((4294967295) | 1))
        case "foo bar":
            switch (fi10 | (-(fi10))) {
            case 136:
                v1 -= 42
            case 4294967295:
                v11 := GenS1{F0: false, F1: 1, F2: -657}
                _ = v11
            default:
                delete(v2, (fi10 | fi10))
            }
        default:
            var v12 bool = false
            _ = v12
        }
        if len(v4) > 0 {
            v4[0] = fi10
        }
        switch (fi10 / (((-(fi10))) | 1)) {
        case 141:
            var v13 int64 = v6
            _ = v13
        case 1:
            if (1.5 <= float32((0.1 * 1e10))) {
                for fi14 := 0; fi14 < 14; fi14++ {
                    v6 -= v6
                    v15 := GenS1{F0: false, F1: 42, F2: -2147483648}
                    if true {
                        switch "" {
                        case "":
                            if (fi10 >= fi10) {
                                var v16 int8 = 1
                                if (-2147483648 <= -626) {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                v1 = 2147483648
                                _ = v16
                            } else {
                                if (false || false) {
                                    if (0 < 255) {
                                        if false {
                                            if (0.0 >= 1e10) {
                                                switch "\n\t" {
                                                case "\n\t":
                                                    v1 = v1
                                                case "a":
                                                    v17 := GenS2{F0: 127, F1: 42, F2: -771}
                                                    _ = v17
                                                }
                                                v18 := GenS2{F0: 128, F1: 18446744073709551615, F2: -1}
                                                _ = v18
                                            }
                                            v6 -= (v6 * (v6 - v6))
                                        } else {
                                            var v19 uint = 96
                                            fmt.Println(fi14, fi10)
                                            _ = v19
                                        }
                                        v1 = v1
                                        v20 := GenS1{F0: false, F1: 65535, F2: 2147483647}
                                        _ = v20
                                    } else {
                                        if (!(false && false)) {
                                            var v21 uint32 = (v1 + (v1 >> ((65535) & 63)))
                                            v3 = append(v3, (v1 & v21))
                                            var v22 uint8 = 1
                                            if false {
                                                _ = len(v3)
                                            }
                                            switch "foo bar" {
                                            case "a":
                                                if false {
                                                    fmt.Println(fi14, v1)
                                                } else {
                                                    fmt.Println(v22, v6, fi14)
                                                    switch (fi10 / ((fi14) | 1)) {
                                                    case 620:
                                                        fmt.Println(v6)
                                                    case 32767:
                                                        v21 = v21
                                                    }
                                                }
                                            default:
                                                v2[-484] = "foo bar"
                                            }
                                            v5["hello"] = false
                                            var v23 uint64 = 2
                                            fmt.Println(fi14, v6)
                                            _ = v21
                                            _ = v22
                                            _ = v23
                                        } else {
                                            if (!false) {
                                                if ((float32((1e10 / (1e10 * ((((((((0.1 + 0.1) * (float64(-2.25) * -2.25)) - 3.14159) * -2.25) - (1.5 - (-2.25 - 0.1))) + -2.25) / float64(1.5)) * ((1.5 * (1.5 / 1.5)) + 3.14159))))) - 1.5) == (-2.25 - 0.1)) {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                            }
                                        }
                                    }
                                } else {
                                    v6 += (v6 & (v6 & (v6 + v6)))
                                    v24 := GenS3{F0: 0, F1: ("" + ((("a" + "\n\t") + "\n\t") + ((("foo bar" + ("foo bar" + ("" + "hello"))) + "\n\t") + "a")))}
                                    _ = v24
                                }
                            }
                        case "hello":
                            fmt.Println(v6, fi14)
                        case "\n\t":
                            v6 = (v6 + (v6 + v6))
                        }
                        var v25 int8 = 100
                        _ = v25
                    }
                    _ = fi14
                    _ = v15
                }
                for fi26 := 0; fi26 < 17; fi26++ {
                    v6 += (v6 * v6)
                    _ = fi26
                }
                v1 = v1
            } else {
                _ = len(v4)
            }
        case -1:
            if ((65535 != 62) || (!((v6 & 42) == 305))) {
                v6 = v6
            } else {
                var v27 uint = 517
                if (!true) {
                    v28 := GenS3{F0: 811, F1: "foo bar"}
                    v29 := GenS1{F0: true, F1: 255, F2: 2147483647}
                    v6 -= -10
                    _ = v28
                    _ = v29
                } else {
                    var v30 uint64 = 2
                    _ = v30
                }
                _ = v27
            }
        }
        _ = fi10
    }
    switch "" {
    case "a":
        for fi31 := 0; fi31 < 20; fi31++ {
            v1 = v1
            var v32 int = (fi31 | (fi31 % ((int(v6)) | 1)))
            _ = fi31
            _ = v32
        }
    case "\n\t":
        switch "hello" {
        case "":
            fmt.Println(v1, v6)
        case "hello":
            switch -1 {
            case 128:
                v1 = (^(v1))
            case 744:
                v1 += (v1 & (v1 | v1))
            }
        }
    case "":
        if (2 < 18446744073709551615) {
            v33 := GenS1{F0: (!false), F1: 426, F2: 944}
            v34 := v5[(("\n\t" + ("\n\t" + "a")) + "\n\t")]
            _ = v33
            _ = v34
        }
    }
    v1 += v1
    if false {
        if len(v3) > 0 {
            v3[0] = (v1 & 2147483647)
        }
    }
    if true {
        v1 = (v1 ^ (^(v1)))
    } else {
        v1 += v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
