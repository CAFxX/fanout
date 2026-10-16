package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int; F1 uint16; F2 uint }
type GenS4 struct { F0 uint32; F1 int32; F2 int8 }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 int32; F1 int16; F2 int8 }
type GenS7 struct { F0 uint32; F1 string }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint8; F1 uint64 }
type GenS11 struct { F0 uint32 }
func genFn0() uint {
    v6 := []int{2147483647}
    var v7 int32 = 0
    var v8 string = "a"
    v7 = (v7 % ((v7) | 1))
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := []uint32{}
    v2 := []uint32{0, 65535, 0}
    var v3 uint64 = 596
    var v4 int32 = 0
    var v5 bool = true
    for fi9 := 0; fi9 < 1; fi9++ {
        v5 = v5
        switch fi9 {
        case 128:
            switch fi9 {
            case -1:
                var v10 uint16 = 2
                _ = v10
            default:
                var v11 uint = 0
                _ = v11
            }
        case 42:
            for fi12 := 0; fi12 < 19; fi12++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v3 = 1
                _ = fi12
            }
        case 989:
            switch "\n\t" {
            case "foo bar":
                var v13 uint8 = 127
                _ = v13
            case "hello":
                fmt.Println(v3, v3, v5)
            case "":
                switch (fi9 & 127) {
                case 431:
                    fmt.Println(v4)
                case -9223372036854775808:
                    var v14 uint64 = (v3 * 202)
                    _ = v14
                case 2147483647:
                    v5 = (1 > int16(v3))
                default:
                    if len(v2) > 0 {
                        v2[0] = 4294967295
                    }
                }
            }
        }
        _ = fi9
    }
    v15 := GenS1{F0: 665, F1: (!v5)}
    if v5 {
        switch 32767 {
        case 32767:
            v16 := GenS2{F0: -359, F1: 255}
            _ = v16
        case 9223372036854775807:
            var v17 float32 = 0.0
            _ = v17
        case -1:
            if len(v1) > 0 {
                v1[0] = 65535
            }
        default:
            var v18 int = 1
            _ = v18
        }
    }
    fmt.Println(v4)
    if (v5 || (!(v5 || v5))) {
        if (42 >= 42) {
            for fi19 := 0; fi19 < 17; fi19++ {
                v3 = (v3 >> ((255) & 63))
                _ = fi19
            }
        } else {
            switch 4294967295 {
            case 0:
                if false {
                    v20 := GenS3{F0: 800, F1: 621, F2: 255}
                    if ((v3 <= v3) && v5) {
                        v21 := GenS4{F0: 2147483648, F1: 991, F2: -128}
                        var v22 uint64 = 2
                        _ = v21
                        _ = v22
                    } else {
                        if (((18446744073709551615 != 4294967295) || v5) && false) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            switch "foo bar" {
                            case "hello":
                                if (19 != -128) {
                                    for fi23 := 0; fi23 < 4; fi23++ {
                                        var v24 float64 = (((0.0 - float64(3.14159)) * (float64(-2.25) * (1.5 - -2.25))) / ((0.1 + 1.5) / float64(3.14159)))
                                        v4 -= v4
                                        _ = fi23
                                        _ = v24
                                    }
                                    v2 = append(v2, 461)
                                }
                            }
                        }
                        if (!("" == ("foo bar" + ("foo bar" + "hello")))) {
                            v25 := GenS5{F0: 32767, F1: 612}
                            if len(v2) > 0 {
                                v2[0] = uint32(v3)
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = v25
                        }
                        switch "hello" {
                        case "hello":
                            var v26 int16 = -1
                            _ = v26
                        default:
                            v27 := GenS2{F0: 0, F1: 129}
                            _ = v27
                        }
                    }
                    v4 += (v4 << ((1) & 63))
                    _ = v20
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        var v28 int16 = int16(v4)
                        _ = v28
                    case "":
                        for fi29 := 0; fi29 < 6; fi29++ {
                            if v5 {
                                if ((v5 || v5) && (!true)) {
                                    fmt.Println(v3, v4, v5)
                                } else {
                                    v5 = v5
                                    fmt.Println(v4)
                                    switch 4294967295 {
                                    case 42:
                                        for fi30 := 0; fi30 < 1; fi30++ {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            var v31 string = ("foo bar" + ("hello" + ("\n\t" + "hello")))
                                            var v32 float32 = (1.5 + 3.14159)
                                            _ = fi30
                                            _ = v31
                                            _ = v32
                                        }
                                    case 4294967295:
                                        for fi33 := 0; fi33 < 5; fi33++ {
                                            v1 = append(v1, 1)
                                            v4 += (v4 >> ((255) & 63))
                                            _ = fi33
                                        }
                                    }
                                }
                                for fi34 := 0; fi34 < 17; fi34++ {
                                    if (v3 < 42) {
                                        fmt.Println(fi34, fi34)
                                        var v35 uint32 = 2
                                        v2 = append(v2, v35)
                                        var v36 int64 = 9223372036854775807
                                        _ = v35
                                        _ = v36
                                    } else {
                                        if false {
                                            if true {
                                                if v5 {
                                                    var v37 uint16 = 42
                                                    v37 = v37
                                                    _ = v37
                                                }
                                            }
                                            if (-9223372036854775808 > 519) {
                                                switch fi34 {
                                                case 42:
                                                    v4 = 1
                                                }
                                                v38 := GenS3{F0: fi34, F1: 2, F2: 0}
                                                _ = v38
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = 4294967295
                                            }
                                            v2 = append(v2, 131)
                                        } else {
                                            switch fi29 {
                                            case -2147483648:
                                                switch (fi34 & fi29) {
                                                case 128:
                                                    switch ("\n\t" + "hello") {
                                                    case "foo bar":
                                                        _ = len(v1)
                                                    default:
                                                        if (v5 && true) {
                                                            v3 += (v3 * (v3 | (v3 & v3)))
                                                            if v5 {
                                                                v5 = (2 == 42)
                                                                var v39 float32 = -2.25
                                                                switch (fi29 | (fi29 + 128)) {
                                                                case 128:
                                                                    var v40 uint16 = 0
                                                                    _ = v40
                                                                case 1:
                                                                    var v41 int8 = -1
                                                                    v3 = (v3 / (((v3 - 2)) | 1))
                                                                    var v42 int16 = -32768
                                                                    _ = v41
                                                                    _ = v42
                                                                case 9223372036854775807:
                                                                    fmt.Println(fi29)
                                                                default:
                                                                    if (int16(fi34) < -399) {
                                                                        v39 += 3.14159
                                                                        fmt.Println(fi34, v3, v39)
                                                                        var v43 int = fi29
                                                                        _ = v43
                                                                    } else {
                                                                        var v44 string = "hello"
                                                                        _ = v44
                                                                    }
                                                                }
                                                                _ = v39
                                                            } else {
                                                                v45 := GenS6{F0: (v4 << ((107) & 63)), F1: 532, F2: 0}
                                                                switch ("\n\t" + "") {
                                                                case "hello":
                                                                    var v46 uint32 = 2147483648
                                                                    _ = v46
                                                                case "a":
                                                                    v4 += (v4 % ((v4) | 1))
                                                                default:
                                                                    var v47 uint = 255
                                                                    _ = v47
                                                                }
                                                                _ = v45
                                                            }
                                                        }
                                                    }
                                                case 42:
                                                    v2 = append(v2, 2147483647)
                                                }
                                            }
                                            if false {
                                                var v48 bool = ((609 >= 65535) && v5)
                                                v4 = v4
                                                _ = v48
                                            }
                                            if ((42 < (v3 + (v3 ^ v3))) && false) {
                                                if false {
                                                    v2 = append(v2, 2147483647)
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    var v49 uint = 1
                                                    _ = v49
                                                } else {
                                                    if (!(128 == (fi29 - fi29))) {
                                                        v1 = append(v1, 65535)
                                                        v5 = (v5 || (-537 != -273))
                                                    }
                                                    v1 = append(v1, 1)
                                                    var v50 int8 = -103
                                                    _ = v50
                                                }
                                            }
                                        }
                                        v51 := GenS7{F0: 788, F1: "\n\t"}
                                        v3 -= (^(v3))
                                        _ = v51
                                    }
                                    var v52 bool = (v5 || (v5 && false))
                                    _ = fi34
                                    _ = v52
                                }
                            } else {
                                fmt.Println(v4, v3, v3)
                            }
                            switch fi29 {
                            case 186:
                                v5 = ((^(fi29)) <= fi29)
                            case 42:
                                if (-2147483648 != v4) {
                                    if ("" >= "\n\t") {
                                        if (2 == 255) {
                                            switch (fi29 * fi29) {
                                            case 4294967295:
                                                _ = len(v2)
                                            }
                                            var v53 bool = (1.5 < 0.0)
                                            var v54 string = "a"
                                            _ = v53
                                            _ = v54
                                        }
                                        for fi55 := 0; fi55 < 16; fi55++ {
                                            var v56 uint32 = 2
                                            _ = fi55
                                            _ = v56
                                        }
                                    } else {
                                        v57 := GenS8{F0: (v3 % (((v3 / (((v3 / ((2) | 1))) | 1))) | 1))}
                                        v5 = (32767 > -859)
                                        var v58 int16 = 1000
                                        _ = v57
                                        _ = v58
                                    }
                                } else {
                                    _ = len(v2)
                                    switch "" {
                                    case "hello":
                                        if len(v1) > 0 {
                                            v1[0] = 2147483648
                                        }
                                    case "foo bar":
                                        for fi59 := 0; fi59 < 5; fi59++ {
                                            v4 = 2147483647
                                            v5 = (118 < 1)
                                            _ = fi59
                                        }
                                    default:
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    var v60 float32 = 3.14159
                                    _ = v60
                                }
                            case 1:
                                v4 -= v4
                            default:
                                if v5 {
                                    for fi61 := 0; fi61 < 10; fi61++ {
                                        v5 = ("hello" < (("\n\t" + ("\n\t" + (((("\n\t" + "") + ("foo bar" + "foo bar")) + "a") + "hello"))) + "a"))
                                        _ = fi61
                                    }
                                    v62 := GenS9{F0: 128}
                                    _ = v62
                                } else {
                                    if ((v3 >> ((255) & 63)) >= v3) {
                                        var v63 int64 = 219
                                        fmt.Println(v4)
                                        _ = v63
                                    }
                                    v3 = (^(v3))
                                }
                            }
                            if v5 {
                                var v64 int16 = 0
                                for fi65 := 0; fi65 < 17; fi65++ {
                                    switch 4294967295 {
                                    case -1:
                                        fmt.Println(v64, fi65)
                                    default:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    var v66 uint64 = (v3 | (^(v3)))
                                    v67 := GenS10{F0: 42, F1: v66}
                                    _ = fi65
                                    _ = v66
                                    _ = v67
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v64
                            }
                            _ = fi29
                        }
                    }
                    for fi68 := 0; fi68 < 20; fi68++ {
                        var v69 float64 = 3.14159
                        _ = fi68
                        _ = v69
                    }
                    v4 += 0
                }
            case 82:
                var v70 uint64 = v3
                _ = v70
            case 128:
                var v71 uint64 = (v3 | 1)
                _ = v71
            default:
                for fi72 := 0; fi72 < 7; fi72++ {
                    switch 2147483647 {
                    case -9223372036854775808:
                        var v73 uint8 = 255
                        _ = v73
                    case 2147483647:
                        fmt.Println(fi72, v5)
                    case 32767:
                        v74 := GenS4{F0: 665, F1: v4, F2: 0}
                        _ = v74
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    _ = fi72
                }
            }
            for fi75 := 0; fi75 < 15; fi75++ {
                var v76 uint64 = (v3 | (v3 - (v3 << ((65535) & 63))))
                if (139 >= 0) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    switch "a" {
                    case "":
                        if v5 {
                            for fi77 := 0; fi77 < 17; fi77++ {
                                if (32767 == 982) {
                                    var v78 string = "foo bar"
                                    v76 -= (v3 & v3)
                                    _ = v78
                                } else {
                                    var v79 bool = v5
                                    _ = v79
                                }
                                _ = fi77
                            }
                            if (101 != 65535) {
                                var v80 int64 = -454
                                _ = v80
                            }
                        }
                    case "\n\t":
                        switch "a" {
                        case "":
                            switch (fi75 << ((255) & 63)) {
                            case -527:
                                if (0 > 0) {
                                    fmt.Println(fi75, v5, v4)
                                }
                            default:
                                var v81 float32 = 0.1
                                _ = v81
                            }
                        case "\n\t":
                            v3 = (v3 % (((v76 << ((255) & 63))) | 1))
                        case "hello":
                            var v82 int = (fi75 % ((fi75) | 1))
                            _ = v82
                        }
                    case "foo bar":
                        switch (fi75 >> ((65535) & 63)) {
                        case 32768:
                            v3 = 268
                        case 128:
                            var v83 uint16 = 65535
                            _ = v83
                        default:
                            switch (fi75 >> ((genFn0()) & 63)) {
                            case 0:
                                v1 = append(v1, 65535)
                            }
                        }
                    default:
                        if len(v2) > 0 {
                            v2[0] = 2147483648
                        }
                    }
                }
                _ = fi75
                _ = v76
            }
        }
        var v84 float32 = 0.0
        fmt.Println(v4)
        _ = v84
    } else {
        var v85 string = "hello"
        _ = v85
    }
    var v86 uint8 = 127
    var v87 string = (((("hello" + "") + "a") + (("a" + ("" + "foo bar")) + "hello")) + "hello")
    switch 0 {
    case -2147483648:
        var v88 float64 = -2.25
        _ = v88
    case -41:
        for fi89 := 0; fi89 < 15; fi89++ {
            for fi90 := 0; fi90 < 3; fi90++ {
                fmt.Println(v3, fi90)
                switch fi90 {
                case -955:
                    v87 = v87
                case 372:
                    v86 = (v86 + v86)
                }
                v91 := GenS3{F0: 721, F1: 1, F2: 255}
                _ = fi90
                _ = v91
            }
            _ = fi89
        }
    }
    switch ("foo bar" + v87) {
    case "":
        var v92 bool = ((^(v3)) < (^(v3)))
        _ = v92
    case "a":
        v93 := GenS11{F0: 2147483647}
        _ = v93
    }
    var v94 int = -352
    for fi95 := 0; fi95 < 14; fi95++ {
        fmt.Println(v86, v94, v86)
        switch v87 {
        case "":
            v3 += 559
        default:
            fmt.Println(v5, v5, v3)
        }
        _ = fi95
    }
    if ((v86 >= (v86 << ((0) & 63))) && (true && v5)) {
        var v96 int32 = (v4 / ((v4) | 1))
        _ = v96
    } else {
        _ = len(v2)
    }
    switch (v87 + "hello") {
    case "hello":
        var v97 float64 = float64(1.5)
        _ = v97
    case "a":
        v98 := GenS2{F0: 42, F1: 42}
        _ = v98
    }
    var v99 float32 = (-2.25 / 1e10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v86:", v86)
    fmt.Println("v87:", v87)
    fmt.Println("v94:", v94)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v86
    _ = v87
    _ = v94
    _ = v99
}
