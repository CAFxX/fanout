package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int64; F1 int; F2 int32 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 int32; F1 uint16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32; F1 string; F2 uint32 }
type GenS8 struct { F0 int16 }

func main() {
    var v1 uint32 = 21
    var v2 uint16 = 255
    v3 := []uint32{(v1 % ((v1) | 1)), (v1 << ((2) & 63)), (v1 / ((4294967295) | 1)), v1}
    v4 := []uint32{(^(v1))}
    var v5 int = 2147483647
    for fi6 := 0; fi6 < 8; fi6++ {
        if false {
            switch 42 {
            case 32768:
                switch "" {
                case "\n\t":
                    v7 := GenS1{F0: 18446744073709551615}
                    _ = v7
                case "":
                    switch v5 {
                    case -1:
                        var v8 string = ("a" + ((("a" + ("\n\t" + "hello")) + "foo bar") + ""))
                        _ = v8
                    case 2147483647:
                        if (100 <= 1) {
                            v5 = (^(v5))
                            var v9 int64 = 1
                            var v10 int64 = v9
                            _ = v10
                            _ = v9
                        } else {
                            if false {
                                v11 := GenS1{F0: 562}
                                v12 := GenS1{F0: 255}
                                v2 -= 988
                                _ = v11
                                _ = v12
                            } else {
                                if true {
                                    fmt.Println(v1, v1)
                                    v13 := GenS2{F0: 42, F1: 0}
                                    for fi14 := 0; fi14 < 2; fi14++ {
                                        if true {
                                            v15 := GenS2{F0: -1, F1: 2}
                                            switch (fi6 / ((fi14) | 1)) {
                                            case 42:
                                                var v16 int64 = int64(fi14)
                                                _ = v16
                                            case 2147483647:
                                                if (false || false) {
                                                    var v17 uint32 = (v1 * 65535)
                                                    var v18 bool = true
                                                    _ = v17
                                                    _ = v18
                                                }
                                            }
                                            switch (fi6 & (fi14 % ((int(v2)) | 1))) {
                                            case 32768:
                                                var v19 string = "a"
                                                _ = v19
                                            case 128:
                                                if true {
                                                    var v20 uint64 = 18446744073709551615
                                                    _ = v20
                                                }
                                                v5 = int(v1)
                                            case 643:
                                                if true {
                                                    v2 -= (v2 & v2)
                                                    var v21 int16 = -32768
                                                    v2 += (v2 | 42)
                                                    _ = v21
                                                }
                                            default:
                                                var v22 int = fi14
                                                _ = v22
                                            }
                                            _ = v15
                                        }
                                        _ = fi14
                                    }
                                    _ = v13
                                } else {
                                    if (!("hello" < "a")) {
                                        v1 = (^(v1))
                                        switch fi6 {
                                        case 128:
                                            v23 := GenS1{F0: 18446744073709551615}
                                            _ = v23
                                        case 39:
                                            v2 += 734
                                        case 0:
                                            switch "\n\t" {
                                            case "foo bar":
                                                if len(v4) > 0 {
                                                    v4[0] = v1
                                                }
                                            case "hello":
                                                v2 = 717
                                            case "\n\t":
                                                v1 += 482
                                            }
                                        default:
                                            v3 = append(v3, v1)
                                        }
                                        fmt.Println(v5, fi6)
                                    }
                                }
                                v24 := GenS2{F0: 1, F1: 255}
                                _ = v24
                            }
                        }
                    case -9223372036854775808:
                        switch fi6 {
                        case -711:
                            v2 -= (v2 % ((v2) | 1))
                        case 32767:
                            v1 = (v1 & 165)
                        case 4294967295:
                            fmt.Println(v1)
                        default:
                            v5 = 4294967295
                        }
                    default:
                        v25 := GenS3{F0: -763, F1: v5, F2: -223}
                        _ = v25
                    }
                }
            case 172:
                var v26 int32 = -64
                _ = v26
            case 793:
                switch (fi6 ^ (v5 / (((^(fi6))) | 1))) {
                case 468:
                    fmt.Println(v2)
                case 2147483647:
                    v2 = (v2 ^ v2)
                default:
                    v27 := GenS4{F0: 818, F1: 127}
                    _ = v27
                }
            default:
                fmt.Println(v2, v1)
            }
        }
        _ = len(v3)
        var v28 float32 = ((0.1 / 1.0) * (3.14159 - float32(0.1)))
        _ = fi6
        _ = v28
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        _ = len(v4)
        _ = fi29
    }
    v5 -= v5
    fmt.Println(v2, v1)
    for fi30 := 0; fi30 < 18; fi30++ {
        fmt.Println(v5, v2, v5)
        _ = fi30
    }
    for fi31 := 0; fi31 < 4; fi31++ {
        var v32 uint8 = 179
        v33 := GenS5{F0: 42, F1: v2}
        _ = fi31
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 13; fi34++ {
        var v35 int = (fi34 / (((v5 * (v5 << ((42) & 63)))) | 1))
        v36 := GenS6{F0: -128}
        _ = fi34
        _ = v35
        _ = v36
    }
    var v37 int8 = -128
    v2 -= v2
    switch ("foo bar" + "") {
    case "\n\t":
        for fi38 := 0; fi38 < 5; fi38++ {
            switch ("a" + "") {
            case "a":
                if true {
                    if false {
                        v5 += 127
                        var v39 int16 = -32768
                        _ = v39
                    } else {
                        switch (fi38 & v5) {
                        case -1:
                            switch "" {
                            case "":
                                v37 -= 122
                            case "hello":
                                v5 = fi38
                            case "\n\t":
                                for fi40 := 0; fi40 < 1; fi40++ {
                                    v1 -= (v1 * (v1 ^ (v1 ^ uint32(v2))))
                                    v2 += (v2 - 0)
                                    if (!("hello" == "")) {
                                        v4 = append(v4, (^(v1)))
                                        v1 -= (v1 / ((v1) | 1))
                                        fmt.Println(v1, v2, v5)
                                    }
                                    _ = fi40
                                }
                            default:
                                var v41 string = "foo bar"
                                _ = v41
                            }
                        case -178:
                            fmt.Println(v2)
                        default:
                            var v42 uint = 0
                            _ = v42
                        }
                        var v43 uint = 255
                        if (("hello" <= "\n\t") || (true && ((1e10 / 1.5) > (0.1 * (-2.25 + (float64((0.1 * 0.1)) * (((0.1 + 0.0) * -2.25) / (0.1 + 3.14159)))))))) {
                            for fi44 := 0; fi44 < 13; fi44++ {
                                var v45 int32 = 1
                                fmt.Println(v45)
                                _ = fi44
                                _ = v45
                            }
                            var v46 int32 = -2147483648
                            if (255 >= 235) {
                                fmt.Println(fi38, v43, v46)
                                if (fi38 != (fi38 ^ v5)) {
                                    v47 := GenS7{F0: v1, F1: "foo bar", F2: v1}
                                    _ = v47
                                }
                            } else {
                                for fi48 := 0; fi48 < 5; fi48++ {
                                    if false {
                                        v46 = (v46 >> ((18446744073709551615) & 63))
                                        fmt.Println(fi48, v37)
                                    } else {
                                        var v49 int32 = (v46 + (v46 + v46))
                                        var v50 uint64 = 4294967295
                                        switch "a" {
                                        case "a":
                                            switch 128 {
                                            case 0:
                                                var v51 float64 = 0.0
                                                _ = v51
                                            case 128:
                                                v5 -= (v5 - v5)
                                            case 9223372036854775807:
                                                v52 := GenS1{F0: v43}
                                                _ = v52
                                            default:
                                                v2 = uint16(v1)
                                            }
                                        case "foo bar":
                                            if len(v3) > 0 {
                                                v3[0] = (v1 / (((v1 & v1)) | 1))
                                            }
                                        case "hello":
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        default:
                                            _ = len(v3)
                                            switch (-(fi38)) {
                                            case 602:
                                                fmt.Println(v37, v2)
                                            }
                                        }
                                        _ = v49
                                        _ = v50
                                    }
                                    _ = fi48
                                }
                            }
                            _ = v46
                        } else {
                            if ((!false) && (-1 > 1000)) {
                                var v53 int64 = 488
                                v54 := GenS4{F0: 18446744073709551615, F1: 113}
                                v43 += (v43 / (((v43 - v43)) | 1))
                                _ = v53
                                _ = v54
                            }
                            v55 := GenS7{F0: uint32(v43), F1: "hello", F2: 2147483647}
                            _ = v55
                        }
                        _ = v43
                    }
                    v1 += v1
                    for fi56 := 0; fi56 < 6; fi56++ {
                        if (true && ((false || false) || true)) {
                            v5 = 128
                            v2 = uint16(v1)
                            var v57 float64 = float64(0.1)
                            var v58 uint8 = 2
                            v37 = (v37 * (v37 | (v37 | v37)))
                            _ = v57
                            _ = v58
                        }
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v59 int16 = -32768
                        _ = fi56
                        _ = v59
                    }
                } else {
                    v1 -= v1
                    v60 := GenS8{F0: -32768}
                    _ = v60
                }
            case "":
                v1 = (v1 - v1)
            case "hello":
                var v61 int32 = -823
                _ = v61
            }
            if (42 != -1) {
                for fi62 := 0; fi62 < 7; fi62++ {
                    var v63 uint = 42
                    switch (v5 / ((-950) | 1)) {
                    case 46:
                        if false {
                            fmt.Println(v5, fi38, v1)
                            var v64 uint16 = 42
                            _ = v64
                        }
                    }
                    _ = fi62
                    _ = v63
                }
                var v65 uint64 = 18446744073709551615
                _ = v65
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v1 = 938
            }
            _ = fi38
        }
    case "foo bar":
        v2 = (^(v2))
    case "a":
        var v66 uint16 = v2
        _ = v66
    }
    if ("" > "hello") {
        fmt.Println(v37)
        v3 = append(v3, v1)
        switch v5 {
        case -2147483648:
            for fi67 := 0; fi67 < 8; fi67++ {
                if false {
                    fmt.Println(fi67, v37, fi67)
                    var v68 int = (fi67 << ((255) & 63))
                    _ = v68
                } else {
                    switch ("foo bar" + ("\n\t" + ("foo bar" + ""))) {
                    case "hello":
                        v37 = 5
                    case "foo bar":
                        switch "\n\t" {
                        case "hello":
                            switch "a" {
                            case "a":
                                fmt.Println(fi67)
                            case "":
                                fmt.Println(v1, v37)
                            }
                        case "a":
                            var v69 bool = true
                            _ = v69
                        }
                    case "\n\t":
                        v4 = append(v4, v1)
                    }
                    if ((float32(0.0) / 1.0) > 3.14159) {
                        v4 = append(v4, (v1 >> ((42) & 63)))
                        v1 += 42
                    }
                }
                switch "\n\t" {
                case "a":
                    fmt.Println(v5, v2)
                }
                v5 = (v5 << ((18446744073709551615) & 63))
                _ = fi67
            }
        case -1:
            for fi70 := 0; fi70 < 14; fi70++ {
                if (v37 == (v37 | (-(v37)))) {
                    _ = len(v4)
                    var v71 int8 = v37
                    var v72 float32 = 1e10
                    _ = v71
                    _ = v72
                } else {
                    v2 = v2
                    var v73 uint64 = 1
                    _ = v73
                }
                v5 -= (v5 * v5)
                _ = fi70
            }
        case 127:
            v37 = (v37 + v37)
        }
    }
    var v74 int = -867
    if (!((v1 + 2) < v1)) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v37 -= (v37 >> ((625) & 63))
    }
    var v75 int16 = -1
    for fi76 := 0; fi76 < 8; fi76++ {
        v75 = (v75 + (v75 + int16(v37)))
        _ = fi76
    }
    v1 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v74:", v74)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v74
    _ = v75
}
