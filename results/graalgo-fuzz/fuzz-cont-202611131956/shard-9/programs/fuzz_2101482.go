package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 bool; F1 int32; F2 int8 }
type GenS3 struct { F0 int64; F1 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 int; F2 bool }
type GenS6 struct { F0 uint; F1 uint8 }
type GenS7 struct { F0 int64; F1 int16 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 int16; F1 uint8 }
type GenS11 struct { F0 int; F1 int16; F2 uint8 }
type GenS12 struct { F0 int16; F1 int32 }

func main() {
    var v1 int = -2147483648
    var v2 string = ("foo bar" + "a")
    v3 := []int{(v1 & -713)}
    var v4 string = "\n\t"
    var v5 int32 = 1
    var v6 int32 = (-(v5))
    var v7 uint = 42
    fmt.Println(v6, v6, v5)
    v8 := GenS1{F0: (128 != 128), F1: v7}
    if (-32768 != 0) {
        var v9 int16 = -1
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v9
    }
    fmt.Println(v6, v1, v5)
    v2 = ""
    v7 -= v7
    v7 = v7
    _ = len(v3)
    switch "hello" {
    case "":
        if len(v3) > 0 {
            v3[0] = -9223372036854775808
        }
    case "a":
        var v10 float32 = 0.1
        _ = v10
    case "foo bar":
        if (!(((float64(1e10) + float64(0.1)) / 1.0) < float64((float32((-2.25 + ((1e10 - 3.14159) + (0.1 / 3.14159)))) - 0.0)))) {
            _ = len(v3)
        } else {
            if (!(true || (42 > 62))) {
                switch v1 {
                case -2147483648:
                    var v11 uint32 = 2147483648
                    _ = v11
                case 410:
                    v6 = v5
                }
            }
            v6 -= 42
        }
    default:
        switch (v4 + v2) {
        case "\n\t":
            if (((v4 + v4) < v2) && false) {
                v6 = 1
                if true {
                    var v12 uint8 = 42
                    fmt.Println(v6, v7, v5)
                    for fi13 := 0; fi13 < 19; fi13++ {
                        if len(v3) > 0 {
                            v3[0] = (fi13 << ((v7) & 63))
                        }
                        var v14 int16 = 0
                        _ = fi13
                        _ = v14
                    }
                    _ = v12
                } else {
                    if len(v3) > 0 {
                        v3[0] = -335
                    }
                    switch v2 {
                    case "a":
                        fmt.Println(v4, v5)
                    case "foo bar":
                        v2 = v2
                    case "":
                        if true {
                            if (-63 >= -119) {
                                v15 := GenS2{F0: true, F1: (v6 * v6), F2: -1}
                                switch v4 {
                                case "foo bar":
                                    v1 = (v1 | v1)
                                case "":
                                    for fi16 := 0; fi16 < 5; fi16++ {
                                        for fi17 := 0; fi17 < 4; fi17++ {
                                            var v18 int64 = 370
                                            v3 = append(v3, fi17)
                                            v5 = 42
                                            _ = fi17
                                            _ = v18
                                        }
                                        for fi19 := 0; fi19 < 7; fi19++ {
                                            if (false || true) {
                                                if len(v3) > 0 {
                                                    v3[0] = v1
                                                }
                                                switch ((((v4 + "\n\t") + v2) + (v2 + ((v2 + v2) + v4))) + v2) {
                                                case "a":
                                                    var v20 int16 = 42
                                                    _ = v20
                                                default:
                                                    if (101 < -1) {
                                                        var v21 int32 = v5
                                                        v22 := GenS1{F0: false, F1: v7}
                                                        _ = v21
                                                        _ = v22
                                                    }
                                                }
                                                v23 := GenS1{F0: ((v6 != (v6 / (((v6 << ((v7) & 63))) | 1))) || (42 < 2147483647)), F1: v7}
                                                _ = v23
                                            }
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            if true {
                                                v24 := GenS1{F0: true, F1: v7}
                                                _ = v24
                                            }
                                            _ = fi19
                                        }
                                        v25 := GenS3{F0: -1, F1: 128}
                                        _ = fi16
                                        _ = v25
                                    }
                                }
                                switch (v1 * 1) {
                                case -599:
                                    if (("a" + (v4 + v2)) >= ((v4 + v2) + v4)) {
                                        if false {
                                            for fi26 := 0; fi26 < 8; fi26++ {
                                                var v27 int64 = -1
                                                if false {
                                                    v1 = (^(v1))
                                                }
                                                _ = fi26
                                                _ = v27
                                            }
                                            fmt.Println(v1, v4, v1)
                                            v3 = append(v3, v1)
                                        } else {
                                            switch "\n\t" {
                                            case "\n\t":
                                                var v28 string = "\n\t"
                                                _ = v28
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = v1
                                            }
                                        }
                                        var v29 float32 = (((-2.25 - (1.5 / (3.14159 + 0.0))) - 3.14159) * (1e10 / float32(1e10)))
                                        _ = v29
                                    }
                                default:
                                    if false {
                                        switch (v1 ^ v1) {
                                        case 639:
                                            var v30 int = (v1 % ((int(v7)) | 1))
                                            _ = v30
                                        }
                                        v6 = (v5 >> ((76) & 63))
                                        v31 := GenS4{F0: 388}
                                        _ = v31
                                    }
                                }
                                _ = v15
                            } else {
                                if false {
                                    var v32 bool = false
                                    v1 = v1
                                    fmt.Println(v5, v7)
                                    _ = v32
                                } else {
                                    if true {
                                        if (v2 != v2) {
                                            _ = len(v3)
                                            if (v7 != (v7 * (v7 * v7))) {
                                                if false {
                                                    for fi33 := 0; fi33 < 14; fi33++ {
                                                        fmt.Println(fi33)
                                                        if (!false) {
                                                            var v34 int8 = 127
                                                            if true {
                                                                for fi35 := 0; fi35 < 6; fi35++ {
                                                                    var v36 float32 = (0.0 - ((1.5 - 1e10) - 1e10))
                                                                    _ = fi35
                                                                    _ = v36
                                                                }
                                                            } else {
                                                                switch (-(fi33)) {
                                                                case 32767:
                                                                    fmt.Println(v5, v1, fi33)
                                                                case -1:
                                                                    var v37 uint64 = 931
                                                                    _ = v37
                                                                case 32768:
                                                                    switch v1 {
                                                                    case -2147483648:
                                                                        switch v4 {
                                                                        case "a":
                                                                            var v38 int8 = (^(v34))
                                                                            _ = v38
                                                                        case "hello":
                                                                            if true {
                                                                                var v39 int16 = 1
                                                                                _ = v39
                                                                            }
                                                                        case "foo bar":
                                                                            v4 = v4
                                                                        default:
                                                                            v40 := GenS5{F0: 255, F1: v1, F2: ((v34 & 127) <= v34)}
                                                                            _ = v40
                                                                        }
                                                                    case 769:
                                                                        if (32767 <= 947) {
                                                                            if true {
                                                                                var v41 int64 = -976
                                                                                switch fi33 {
                                                                                case 32767:
                                                                                    v2 = (v2 + "\n\t")
                                                                                }
                                                                                var v42 int64 = 452
                                                                                _ = v41
                                                                                _ = v42
                                                                            } else {
                                                                                v43 := GenS6{F0: (^(v7)), F1: 127}
                                                                                _ = v43
                                                                            }
                                                                            var v44 bool = ((206 < 42) || ((65535 > 486) || ((v7 << (((^(v7))) & 63)) <= 2)))
                                                                            for fi45 := 0; fi45 < 13; fi45++ {
                                                                                fmt.Println(v44, v44, v5)
                                                                                if true {
                                                                                    v3 = append(v3, fi45)
                                                                                    _ = len(v3)
                                                                                    var v46 int64 = 42
                                                                                    _ = v46
                                                                                } else {
                                                                                    var v47 uint8 = 172
                                                                                    _ = v47
                                                                                }
                                                                                switch ((v4 + v4) + v4) {
                                                                                case "foo bar":
                                                                                    fmt.Println(v34, v7, v34)
                                                                                case "hello":
                                                                                    v1 += (fi45 * (fi45 / (((v1 >> ((2) & 63))) | 1)))
                                                                                }
                                                                                _ = fi45
                                                                            }
                                                                            _ = v44
                                                                        } else {
                                                                            for fi48 := 0; fi48 < 17; fi48++ {
                                                                                fmt.Println(v6)
                                                                                _ = fi48
                                                                            }
                                                                            v49 := GenS7{F0: -231, F1: -621}
                                                                            if true {
                                                                                v50 := GenS8{F0: 4294967295}
                                                                                _ = v50
                                                                            }
                                                                            _ = v49
                                                                        }
                                                                    default:
                                                                        if len(v3) > 0 {
                                                                            v3[0] = v1
                                                                        }
                                                                    }
                                                                }
                                                                var v51 uint8 = 255
                                                                var v52 int16 = -717
                                                                _ = v51
                                                                _ = v52
                                                            }
                                                            v3 = append(v3, 9223372036854775807)
                                                            _ = v34
                                                        }
                                                        v6 -= (^(v5))
                                                        _ = fi33
                                                    }
                                                } else {
                                                    switch v2 {
                                                    case "a":
                                                        fmt.Println(v4)
                                                    }
                                                    if (480 == 42) {
                                                        v6 -= v6
                                                        for fi53 := 0; fi53 < 2; fi53++ {
                                                            for fi54 := 0; fi54 < 8; fi54++ {
                                                                if len(v3) > 0 {
                                                                    v3[0] = fi53
                                                                }
                                                                _ = fi54
                                                            }
                                                            _ = fi53
                                                        }
                                                    }
                                                }
                                            }
                                            if true {
                                                switch (v4 + (v4 + v2)) {
                                                case "\n\t":
                                                    v2 = "foo bar"
                                                }
                                            }
                                        }
                                        for fi55 := 0; fi55 < 4; fi55++ {
                                            var v56 int8 = 42
                                            _ = fi55
                                            _ = v56
                                        }
                                    } else {
                                        for fi57 := 0; fi57 < 5; fi57++ {
                                            v6 += (v6 ^ (v5 + v5))
                                            _ = fi57
                                        }
                                    }
                                    v7 -= 2
                                }
                                v3 = append(v3, (v1 >> (((v7 - v7)) & 63)))
                            }
                            v2 = v4
                            if (v7 <= (v7 | (v7 * v7))) {
                                v7 = 883
                            }
                        } else {
                            fmt.Println(v7)
                            v2 = v2
                            v58 := GenS9{F0: -301}
                            _ = v58
                        }
                    default:
                        v59 := GenS5{F0: 4294967295, F1: v1, F2: false}
                        _ = v59
                    }
                }
            } else {
                _ = len(v3)
            }
        }
    }
    if true {
        v1 = (v1 - -1)
        switch (v2 + "foo bar") {
        case "foo bar":
            var v60 uint32 = 2147483648
            _ = v60
        default:
            fmt.Println(v5)
        }
    }
    for fi61 := 0; fi61 < 8; fi61++ {
        v4 = v4
        switch fi61 {
        case 1:
            for fi62 := 0; fi62 < 4; fi62++ {
                v63 := GenS3{F0: 0, F1: 127}
                _ = fi62
                _ = v63
            }
        case -543:
            if false {
                switch ("" + (v4 + "hello")) {
                case "hello":
                    v5 += v6
                case "":
                    var v64 uint = v7
                    _ = v64
                default:
                    var v65 uint64 = 42
                    _ = v65
                }
                for fi66 := 0; fi66 < 14; fi66++ {
                    fmt.Println(fi66)
                    _ = fi66
                }
            }
        case 128:
            fmt.Println(v7, v4, fi61)
        }
        switch "" {
        case "a":
            fmt.Println(fi61, v1)
        case "hello":
            for fi67 := 0; fi67 < 7; fi67++ {
                switch fi67 {
                case -2147483648:
                    v68 := GenS2{F0: (!false), F1: (-(v5)), F2: -1}
                    _ = v68
                default:
                    v69 := GenS10{F0: -32768, F1: 42}
                    _ = v69
                }
                var v70 float64 = (1e10 - 1e10)
                if ("\n\t" >= v4) {
                    v71 := GenS11{F0: fi61, F1: -1, F2: uint8(v7)}
                    _ = len(v3)
                    _ = len(v3)
                    v2 = "\n\t"
                    v7 -= v7
                    _ = v71
                }
                _ = fi67
                _ = v70
            }
        }
        _ = fi61
    }
    var v72 uint64 = 4294967295
    if false {
        var v73 uint = (^(v7))
        v74 := GenS12{F0: 1000, F1: 1}
        _ = v73
        _ = v74
    } else {
        v72 = (v72 + v72)
        v3 = append(v3, v1)
        if (3.14159 <= 3.14159) {
            v5 -= (v6 << ((v7) & 63))
            v75 := GenS4{F0: (v72 % ((537) | 1))}
            _ = v75
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v72
    _ = v8
}
