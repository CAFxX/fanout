package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 bool }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 int8; F2 string }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 string; F1 int64; F2 uint16 }
type GenS6 struct { F0 uint8; F1 uint64; F2 int32 }
type GenS7 struct { F0 int32; F1 uint8; F2 string }
type GenS8 struct { F0 int }
type GenS9 struct { F0 string; F1 int64 }
type GenS10 struct { F0 string; F1 int }
type GenS11 struct { F0 uint16 }

func main() {
    v1 := []int{-2147483648, -677, 9223372036854775807}
    var v2 uint = 4294967295
    var v3 float64 = 3.14159
    v4 := []int{32768, 127, 9223372036854775807}
    if (-9223372036854775808 == 42) {
        for fi5 := 0; fi5 < 10; fi5++ {
            if false {
                if (("foo bar" + "foo bar") > "\n\t") {
                    var v6 float32 = 3.14159
                    v2 += v2
                    var v7 int64 = 42
                    _ = v6
                    _ = v7
                }
                v3 = (v3 / v3)
            } else {
                if true {
                    v8 := GenS1{F0: -1, F1: (0 != 440), F2: (-109 > 9)}
                    if (("a" + ("hello" + "a")) < "foo bar") {
                        v9 := GenS2{F0: 0}
                        _ = v9
                    }
                    v10 := GenS3{F0: "foo bar", F1: 1, F2: "\n\t"}
                    _ = v10
                    _ = v8
                }
                v11 := GenS4{F0: -122, F1: 1000}
                v12 := GenS1{F0: -1, F1: ((^(fi5)) < fi5), F2: true}
                _ = v11
                _ = v12
            }
            v3 = v3
            _ = fi5
        }
    }
    var v13 int = 497
    for fi14 := 0; fi14 < 18; fi14++ {
        v15 := GenS4{F0: -128, F1: 2147483647}
        if (true || true) {
            v16 := GenS3{F0: "foo bar", F1: -1, F2: "\n\t"}
            _ = v16
        }
        v3 += (1.5 * (float64(3.14159) / v3))
        _ = fi14
        _ = v15
    }
    var v17 uint = v2
    switch "a" {
    case "\n\t":
        for fi18 := 0; fi18 < 6; fi18++ {
            var v19 float32 = ((0.1 / 1.5) - (3.14159 - 1e10))
            _ = fi18
            _ = v19
        }
    case "hello":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "foo bar":
        var v20 float32 = 3.14159
        _ = v20
    }
    var v21 uint8 = 185
    var v22 uint64 = 2
    var v23 uint16 = 1
    if ((255 < (^(v2))) && (!false)) {
        if (uint(v22) != 967) {
            var v24 int64 = 114
            var v25 uint64 = (v22 + (v22 / ((v22) | 1)))
            v22 += 4294967295
            _ = v24
            _ = v25
        }
        if false {
            v26 := GenS5{F0: "foo bar", F1: 915, F2: v23}
            _ = v26
        } else {
            fmt.Println(v23)
            if (191 <= (v13 & 284)) {
                if (false || false) {
                    var v27 uint16 = 776
                    if true {
                        v28 := GenS6{F0: uint8(v2), F1: uint64(v2), F2: -1}
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 8; fi29++ {
                        fmt.Println(v13, v22)
                        var v30 bool = false
                        v2 += 65535
                        _ = fi29
                        _ = v30
                    }
                    _ = v27
                } else {
                    switch "\n\t" {
                    case "\n\t":
                        for fi31 := 0; fi31 < 16; fi31++ {
                            var v32 float64 = v3
                            for fi33 := 0; fi33 < 14; fi33++ {
                                var v34 uint16 = v23
                                if ((v21 * v21) < v21) {
                                    if (0.0 != 0.1) {
                                        if (v34 < (v23 / ((v23) | 1))) {
                                            v35 := GenS1{F0: 525, F1: (42 < 9223372036854775807), F2: true}
                                            var v36 float32 = 3.14159
                                            _ = v35
                                            _ = v36
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        v32 = (v3 / v3)
                                    }
                                    fmt.Println(v22, v34)
                                    var v37 float32 = ((float32((float64(1e10) * v3)) * 0.0) - 3.14159)
                                    _ = v37
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = 0
                                    }
                                    v32 = 3.14159
                                    if (false || (-100 > -26)) {
                                        switch 32768 {
                                        case 127:
                                            fmt.Println(v23)
                                            switch -473 {
                                            case 127:
                                                var v38 int = (fi31 & fi33)
                                                _ = v38
                                            case 1:
                                                v17 += (v17 ^ v2)
                                            }
                                        case -1:
                                            v21 -= v21
                                        case 4294967295:
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                    }
                                }
                                _ = fi33
                                _ = v34
                            }
                            _ = fi31
                            _ = v32
                        }
                    case "":
                        var v39 float32 = ((1e10 * 1.5) * 1.5)
                        _ = v39
                    case "a":
                        for fi40 := 0; fi40 < 9; fi40++ {
                            for fi41 := 0; fi41 < 18; fi41++ {
                                switch ("a" + "\n\t") {
                                case "foo bar":
                                    v13 += v13
                                case "\n\t":
                                    if (false || false) {
                                        v3 -= (v3 + 1.5)
                                        var v42 uint64 = (v22 % ((v22) | 1))
                                        _ = v42
                                    }
                                case "":
                                    v3 += 0.0
                                default:
                                    switch ("\n\t" + "hello") {
                                    case "foo bar":
                                        if true {
                                            var v43 string = "foo bar"
                                            _ = v43
                                        } else {
                                            v44 := GenS1{F0: 765, F1: false, F2: (0 <= uint16(v22))}
                                            fmt.Println(fi40, fi41)
                                            _ = v44
                                        }
                                    }
                                }
                                _ = fi41
                            }
                            var v45 uint64 = v22
                            for fi46 := 0; fi46 < 2; fi46++ {
                                switch "foo bar" {
                                case "\n\t":
                                    var v47 float64 = 3.14159
                                    var v48 float64 = (0.0 + (3.14159 - v47))
                                    var v49 float32 = float32(float64((0.1 - (-2.25 / 0.1))))
                                    _ = v47
                                    _ = v48
                                    _ = v49
                                case "a":
                                    if (false && false) {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        _ = len(v4)
                                    } else {
                                        v3 += (v3 * 1e10)
                                        v50 := GenS7{F0: 0, F1: (v21 | (^(v21))), F2: ("" + "")}
                                        if len(v1) > 0 {
                                            v1[0] = 2147483647
                                        }
                                        fmt.Println(v17, v22, v23)
                                        v51 := GenS1{F0: 32767, F1: false, F2: (v45 < v45)}
                                        _ = len(v1)
                                        v52 := GenS5{F0: ("foo bar" + ((("foo bar" + "foo bar") + "a") + "foo bar")), F1: 42, F2: 2}
                                        var v53 int32 = -564
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v50
                                        _ = v51
                                        _ = v52
                                        _ = v53
                                    }
                                default:
                                    switch "foo bar" {
                                    case "":
                                        v54 := GenS8{F0: -9223372036854775808}
                                        _ = v54
                                    case "a":
                                        switch ("foo bar" + "hello") {
                                        case "hello":
                                            var v55 uint32 = 873
                                            _ = v55
                                        case "a":
                                            v2 += (v17 + (^(v17)))
                                        default:
                                            if (true && ((v21 & (v21 << ((1) & 63))) >= 112)) {
                                                var v56 int = (fi46 * fi40)
                                                _ = v56
                                            }
                                        }
                                    case "hello":
                                        switch 42 {
                                        case 127:
                                            if (127 < v21) {
                                                v2 = v2
                                                v4 = append(v4, (v13 - fi40))
                                            } else {
                                                if false {
                                                    fmt.Println(v23)
                                                    v13 = (v13 << ((v2) & 63))
                                                    v22 = (v22 << ((v2) & 63))
                                                } else {
                                                    if (!false) {
                                                        v13 = v13
                                                        var v57 int8 = 0
                                                        if len(v4) > 0 {
                                                            v4[0] = (v13 & (fi40 - v13))
                                                        }
                                                        _ = v57
                                                    } else {
                                                        if true {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                        }
                                                        v21 -= 2
                                                    }
                                                    var v58 uint = 18446744073709551615
                                                    _ = v58
                                                }
                                                var v59 uint16 = 32767
                                                if true {
                                                    v4 = append(v4, (fi46 & 828))
                                                    if (true && true) {
                                                        var v60 float64 = 3.14159
                                                        _ = v60
                                                    }
                                                    v61 := GenS2{F0: 0}
                                                    v17 = (v17 ^ (^(v17)))
                                                    v21 += (v21 ^ 1)
                                                    _ = v61
                                                } else {
                                                    if (v3 != v3) {
                                                        if (1 >= int16(v45)) {
                                                            var v62 float64 = (v3 * (v3 - ((-2.25 / v3) + 0.1)))
                                                            if len(v1) > 0 {
                                                                v1[0] = -73
                                                            }
                                                            var v63 uint64 = v45
                                                            _ = v62
                                                            _ = v63
                                                        }
                                                        v64 := GenS9{F0: "", F1: -9223372036854775808}
                                                        _ = v64
                                                    }
                                                    var v65 bool = true
                                                    _ = len(v1)
                                                    _ = v65
                                                }
                                                _ = v59
                                            }
                                        case 2147483647:
                                            v66 := GenS9{F0: "a", F1: 9223372036854775807}
                                            _ = v66
                                        default:
                                            var v67 int64 = -954
                                            _ = v67
                                        }
                                    default:
                                        fmt.Println(v3, fi40, v2)
                                    }
                                }
                                _ = fi46
                            }
                            _ = fi40
                            _ = v45
                        }
                    default:
                        if len(v1) > 0 {
                            v1[0] = (^(v13))
                        }
                    }
                }
                if true {
                    switch (v13 << ((1) & 63)) {
                    case 9223372036854775807:
                        var v68 uint8 = (v21 | (v21 / ((v21) | 1)))
                        _ = v68
                    case 128:
                        if true {
                            _ = len(v4)
                            if (!true) {
                                v23 += 20
                                v69 := GenS5{F0: "", F1: 1, F2: v23}
                                _ = v69
                            } else {
                                var v70 uint16 = (^(v23))
                                v22 = (v22 - v22)
                                _ = v70
                            }
                        }
                    case -192:
                        var v71 float64 = float64(((-2.25 / (-2.25 - float32(v3))) - 1.5))
                        _ = v71
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    var v72 string = ("foo bar" + ("\n\t" + ""))
                    v73 := GenS9{F0: v72, F1: -382}
                    _ = v72
                    _ = v73
                } else {
                    var v74 int16 = -386
                    if (18446744073709551615 <= v2) {
                        for fi75 := 0; fi75 < 17; fi75++ {
                            v76 := GenS3{F0: ("" + "foo bar"), F1: -128, F2: "hello"}
                            switch (fi75 - 0) {
                            case 1:
                                v23 = v23
                            case 816:
                                v77 := GenS10{F0: ((((("a" + "a") + "hello") + "foo bar") + "foo bar") + "foo bar"), F1: fi75}
                                _ = v77
                            }
                            v3 = (0.1 - v3)
                            _ = fi75
                            _ = v76
                        }
                    }
                    _ = v74
                }
                for fi78 := 0; fi78 < 11; fi78++ {
                    var v79 uint = (^(v2))
                    for fi80 := 0; fi80 < 1; fi80++ {
                        switch "a" {
                        case "foo bar":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        default:
                            if len(v1) > 0 {
                                v1[0] = (v13 >> ((649) & 63))
                            }
                        }
                        _ = fi80
                    }
                    _ = fi78
                    _ = v79
                }
            } else {
                var v81 int16 = 0
                _ = v81
            }
            var v82 uint16 = v23
            _ = v82
        }
    }
    switch (-(v13)) {
    case 32767:
        switch ("\n\t" + ("foo bar" + "a")) {
        case "a":
            var v83 int64 = 538
            _ = v83
        case "":
            if len(v1) > 0 {
                v1[0] = -1
            }
        case "hello":
            if ((("a" + ("foo bar" + "foo bar")) + "foo bar") < ((("" + "\n\t") + (("foo bar" + ("" + "")) + "foo bar")) + "")) {
                for fi84 := 0; fi84 < 1; fi84++ {
                    if false {
                        for fi85 := 0; fi85 < 17; fi85++ {
                            if true {
                                _ = len(v1)
                                _ = len(v1)
                                fmt.Println(v13)
                            }
                            var v86 int8 = -128
                            v22 += (v22 & v22)
                            _ = len(v1)
                            fmt.Println(v17, v86, v21)
                            var v87 uint32 = 4294967295
                            _ = fi85
                            _ = v86
                            _ = v87
                        }
                    }
                    _ = fi84
                }
            }
        }
    case 889:
        if true {
            fmt.Println(v17, v17, v13)
            v88 := GenS2{F0: 2147483647}
            fmt.Println(v21, v3, v22)
            _ = v88
        } else {
            for fi89 := 0; fi89 < 15; fi89++ {
                _ = len(v1)
                v90 := GenS11{F0: (v23 | 32767)}
                var v91 int64 = -9223372036854775808
                _ = fi89
                _ = v90
                _ = v91
            }
            fmt.Println(v21, v17)
        }
    default:
        var v92 float32 = 0.0
        _ = v92
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) > 0 {
        v4[0] = (-(v13))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
