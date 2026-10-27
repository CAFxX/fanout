package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int64; F1 int; F2 int16 }
type GenS10 struct { F0 uint16; F1 uint64 }
type GenS11 struct { F0 int64; F1 int64; F2 int8 }
type GenS12 struct { F0 uint16 }
type GenS13 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS14 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS15 struct { F0 int8; F1 int16; F2 int32 }

func main() {
    var v1 int16 = 32767
    var v2 int16 = 42
    var v3 uint32 = 2
    v4 := []uint32{42, 973}
    var v5 bool = false
    if (!(((65535 < uint(v3)) && (!(128 != -224))) || true)) {
        v1 = (v2 | (v1 * 32767))
        v2 += v2
        v6 := GenS1{F0: 941, F1: 7}
        _ = v6
    } else {
        fmt.Println(v1)
        if len(v4) > 0 {
            v4[0] = v3
        }
    }
    if (false && false) {
        fmt.Println(v5, v1)
    } else {
        v7 := GenS1{F0: -1, F1: 2}
        if len(v4) > 0 {
            v4[0] = (v3 >> ((1) & 63))
        }
        v8 := GenS2{F0: 0}
        _ = v7
        _ = v8
    }
    v5 = (2 >= 759)
    if true {
        for fi9 := 0; fi9 < 19; fi9++ {
            fmt.Println(v5)
            fmt.Println(v1)
            if (v5 || (v5 && false)) {
                for fi10 := 0; fi10 < 14; fi10++ {
                    _ = len(v4)
                    var v11 float64 = 1.5
                    v3 = v3
                    _ = fi10
                    _ = v11
                }
            }
            _ = fi9
        }
        var v12 int16 = (v2 / ((0) | 1))
        _ = v12
    }
    if (21 <= 4294967295) {
        if (true && v5) {
            fmt.Println(v2)
            for fi13 := 0; fi13 < 10; fi13++ {
                if false {
                    for fi14 := 0; fi14 < 10; fi14++ {
                        if ((v2 >> ((807) & 63)) >= (v2 << ((4294967295) & 63))) {
                            if v5 {
                                fmt.Println(v5, v5, fi13)
                                if (v3 > (^(v3))) {
                                    if len(v4) > 0 {
                                        v4[0] = (v3 << ((42) & 63))
                                    }
                                    v1 = (-(v2))
                                    var v15 int8 = 42
                                    _ = v15
                                } else {
                                    var v16 uint8 = 0
                                    fmt.Println(v5)
                                    fmt.Println(v3, v3, v1)
                                    _ = v16
                                }
                                var v17 float64 = 3.14159
                                _ = v17
                            }
                            if true {
                                v4 = append(v4, 0)
                                if (!v5) {
                                    v5 = false
                                    var v18 int8 = -1
                                    var v19 uint8 = 109
                                    var v20 uint16 = 333
                                    var v21 int = 128
                                    if (true || (!false)) {
                                        var v22 int = (fi14 & v21)
                                        _ = v22
                                    }
                                    _ = v18
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                } else {
                                    v3 = (v3 / (((v3 & v3)) | 1))
                                    var v23 int = (fi14 >> ((53) & 63))
                                    v3 = 2
                                    fmt.Println(v2)
                                    _ = v23
                                }
                                switch "" {
                                case "\n\t":
                                    if len(v4) > 0 {
                                        v4[0] = v3
                                    }
                                case "foo bar":
                                    var v24 uint16 = 255
                                    _ = v24
                                case "a":
                                    if v5 {
                                        var v25 int64 = 9223372036854775807
                                        _ = v25
                                    } else {
                                        var v26 int64 = 1
                                        var v27 float32 = 0.1
                                        _ = v26
                                        _ = v27
                                    }
                                }
                            }
                            v1 = 32767
                        } else {
                            v1 = (v1 | v2)
                        }
                        var v28 int16 = (v2 & v2)
                        _ = fi14
                        _ = v28
                    }
                    switch fi13 {
                    case 253:
                        if v5 {
                            switch "" {
                            case "":
                                switch 0 {
                                case -2147483648:
                                    v5 = v5
                                default:
                                    var v29 uint64 = 0
                                    _ = v29
                                }
                            case "a":
                                switch (-(fi13)) {
                                case 9223372036854775807:
                                    v3 = (v3 << ((4294967295) & 63))
                                case 1:
                                    v30 := GenS2{F0: -358}
                                    _ = v30
                                case 32767:
                                    v31 := GenS3{F0: 0}
                                    _ = v31
                                }
                            case "hello":
                                if (false && false) {
                                    if v5 {
                                        var v32 int = 2147483647
                                        _ = v32
                                    } else {
                                        var v33 uint8 = 1
                                        if len(v4) > 0 {
                                            v4[0] = v3
                                        }
                                        v34 := GenS1{F0: -2147483648, F1: 481}
                                        _ = v33
                                        _ = v34
                                    }
                                    _ = len(v4)
                                }
                            default:
                                var v35 uint32 = (^(v3))
                                _ = v35
                            }
                            switch "a" {
                            case "\n\t":
                                switch fi13 {
                                case -580:
                                    v36 := GenS3{F0: 42}
                                    _ = v36
                                default:
                                    v5 = v5
                                }
                            case "":
                                for fi37 := 0; fi37 < 9; fi37++ {
                                    v38 := GenS4{F0: 1, F1: -9223372036854775808}
                                    var v39 int64 = -9223372036854775808
                                    v2 -= v2
                                    _ = fi37
                                    _ = v38
                                    _ = v39
                                }
                            default:
                                v3 = (v3 | v3)
                            }
                            switch (fi13 + (fi13 * fi13)) {
                            case 0:
                                var v40 uint64 = 614
                                _ = v40
                            case -713:
                                v2 = (^(v1))
                            default:
                                v5 = (v5 && v5)
                            }
                        } else {
                            var v41 uint = 981
                            var v42 bool = v5
                            if (uint8(v41) != 194) {
                                var v43 uint16 = 255
                                var v44 int64 = int64(v41)
                                _ = v43
                                _ = v44
                            }
                            _ = v41
                            _ = v42
                        }
                    case 773:
                        v5 = (v3 == 1)
                    }
                }
                if (-705 > -1) {
                    v3 = (v3 & v3)
                    if len(v4) > 0 {
                        v4[0] = (v3 >> ((65535) & 63))
                    }
                } else {
                    if (v1 < v2) {
                        var v45 uint32 = 2
                        _ = v45
                    } else {
                        fmt.Println(v2, v2)
                    }
                    for fi46 := 0; fi46 < 19; fi46++ {
                        v5 = (1 == 0)
                        v5 = (4294967295 <= 42)
                        v47 := GenS1{F0: 2147483647, F1: 589}
                        if false {
                            _ = len(v4)
                            switch fi13 {
                            case -963:
                                v1 += v1
                            default:
                                v48 := GenS4{F0: 379, F1: -1}
                                _ = v48
                            }
                        } else {
                            switch (("\n\t" + "a") + "a") {
                            case "\n\t":
                                v5 = ((-104 == int64(v3)) || v5)
                            case "a":
                                v5 = (!(-6 > -1))
                            }
                            if len(v4) > 0 {
                                v4[0] = v3
                            }
                            if (1 != 2147483647) {
                                _ = len(v4)
                                if v5 {
                                    v2 = -417
                                    var v49 int64 = 42
                                    _ = v49
                                }
                            } else {
                                var v50 uint8 = 127
                                v51 := GenS5{F0: 0, F1: (!((3.14159 / 1.5) == (-2.25 + (-2.25 * 3.14159))))}
                                _ = v50
                                _ = v51
                            }
                        }
                        v3 = 594
                        _ = fi46
                        _ = v47
                    }
                }
                fmt.Println(fi13, fi13, fi13)
                _ = fi13
            }
        } else {
            var v52 bool = v5
            v5 = true
            if ((v52 || (0 > 2147483647)) || true) {
                var v53 uint = 695
                if len(v4) > 0 {
                    v4[0] = (v3 ^ 2147483647)
                }
                _ = v53
            } else {
                var v54 int32 = 1
                for fi55 := 0; fi55 < 17; fi55++ {
                    v56 := GenS6{F0: 1, F1: (int8(fi55) >= 100), F2: 0}
                    _ = fi55
                    _ = v56
                }
                _ = v54
            }
            _ = v52
        }
        var v57 uint32 = v3
        _ = v57
    } else {
        switch 2147483647 {
        case 32768:
            fmt.Println(v5, v2)
        case -9223372036854775808:
            v58 := GenS7{F0: 127}
            _ = v58
        case 32767:
            if ((9223372036854775807 > 128) && (false || (-9223372036854775808 >= 2147483647))) {
                for fi59 := 0; fi59 < 16; fi59++ {
                    var v60 float64 = (-2.25 - (((float64((3.14159 * -2.25)) * 0.0) * 1e10) + 0.1))
                    fmt.Println(v2, v1, v60)
                    _ = fi59
                    _ = v60
                }
                v3 = (^(v3))
            }
        }
        v4 = append(v4, v3)
        for fi61 := 0; fi61 < 15; fi61++ {
            for fi62 := 0; fi62 < 19; fi62++ {
                v3 += v3
                if (!v5) {
                    v5 = (65535 != 65535)
                }
                switch ("foo bar" + "foo bar") {
                case "foo bar":
                    if true {
                        var v63 uint16 = 808
                        fmt.Println(v1, fi62)
                        switch fi62 {
                        case -9223372036854775808:
                            if len(v4) > 0 {
                                v4[0] = v3
                            }
                            fmt.Println(fi61, v3, fi61)
                            var v64 int64 = -9223372036854775808
                            _ = v64
                        case 128:
                            switch ((("hello" + "hello") + ("foo bar" + ("a" + "\n\t"))) + "") {
                            case "hello":
                                var v65 int = (fi62 / ((fi62) | 1))
                                _ = v65
                            case "foo bar":
                                if v5 {
                                    fmt.Println(v5, v1)
                                    fmt.Println(v3, v3, v5)
                                    v5 = v5
                                } else {
                                    var v66 uint16 = v63
                                    _ = v66
                                }
                            default:
                                var v67 uint8 = 179
                                _ = v67
                            }
                            if len(v4) > 0 {
                                v4[0] = v3
                            }
                            var v68 uint16 = v63
                            _ = v68
                        case 42:
                            v1 = v2
                        }
                        _ = v63
                    } else {
                        v4 = append(v4, v3)
                        v1 -= (v1 % (((v2 % ((v2) | 1))) | 1))
                    }
                case "\n\t":
                    switch fi62 {
                    case 42:
                        var v69 bool = (v5 || (-631 != -1))
                        _ = v69
                    case -2147483648:
                        v4 = append(v4, (v3 + v3))
                    case 32767:
                        switch ("a" + ("a" + ((("hello" + "") + "foo bar") + "\n\t"))) {
                        case "":
                            var v70 int64 = 533
                            _ = v70
                        case "a":
                            _ = len(v4)
                        default:
                            v5 = (2 >= 2)
                        }
                    default:
                        if true {
                            v71 := GenS8{F0: 42}
                            v72 := GenS9{F0: 348, F1: (fi61 * 0), F2: (v2 * v2)}
                            v3 -= (v3 & (^(v3)))
                            _ = v71
                            _ = v72
                        } else {
                            switch fi62 {
                            case 32767:
                                if true {
                                    if v5 {
                                        v2 = v1
                                        v5 = true
                                    } else {
                                        fmt.Println(v3, fi61)
                                    }
                                } else {
                                    var v73 uint16 = 42
                                    var v74 int32 = 377
                                    if v5 {
                                        v75 := GenS7{F0: 1}
                                        fmt.Println(v2)
                                        if false {
                                            var v76 int64 = 772
                                            if (42 <= 1) {
                                                v77 := GenS7{F0: 255}
                                                v73 = v73
                                                fmt.Println(fi61, v3)
                                                _ = v77
                                            } else {
                                                v4 = append(v4, 693)
                                            }
                                            _ = v76
                                        }
                                        _ = v75
                                    }
                                    _ = v73
                                    _ = v74
                                }
                            case 128:
                                if (42 < -1) {
                                    fmt.Println(fi61, fi62)
                                    v78 := GenS9{F0: 0, F1: (^(fi61)), F2: (v2 & (-(v1)))}
                                    _ = v78
                                } else {
                                    v79 := GenS6{F0: 2, F1: (0 == 656), F2: 42}
                                    if v5 {
                                        var v80 int64 = -1
                                        if (1000 <= v1) {
                                            switch ("" + ("a" + "foo bar")) {
                                            case "hello":
                                                var v81 uint = 255
                                                _ = v81
                                            case "\n\t":
                                                switch (fi61 >> ((1) & 63)) {
                                                case -1:
                                                    v1 += (v2 >> ((255) & 63))
                                                default:
                                                    var v82 int64 = (^(v80))
                                                    _ = v82
                                                }
                                            case "a":
                                                v4 = append(v4, (v3 - (^(v3))))
                                            }
                                            v3 += v3
                                            v3 = (v3 & 2)
                                        }
                                        if (!(v5 && (!v5))) {
                                            if false {
                                                switch fi62 {
                                                case -2147483648:
                                                    fmt.Println(fi61, v5)
                                                }
                                            }
                                        } else {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            v83 := GenS3{F0: 42}
                                            _ = v83
                                        }
                                        switch 42 {
                                        case 0:
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        _ = v80
                                    }
                                    if v5 {
                                        fmt.Println(v2, v5, v2)
                                        fmt.Println(v3, v5, v3)
                                    }
                                    _ = v79
                                }
                            case 264:
                                v5 = (4294967295 == 18446744073709551615)
                            default:
                                var v84 bool = false
                                _ = v84
                            }
                            if (42 <= 115) {
                                if len(v4) > 0 {
                                    v4[0] = 65535
                                }
                            }
                        }
                    }
                default:
                    if v5 {
                        switch fi61 {
                        case 42:
                            v2 = (v1 % ((76) | 1))
                        case 4294967295:
                            v2 = (v1 & v2)
                        case -416:
                            v4 = append(v4, v3)
                        default:
                            var v85 string = ((("hello" + "\n\t") + "hello") + "hello")
                            _ = v85
                        }
                    }
                    v3 = (v3 - (^(v3)))
                }
                _ = fi62
            }
            switch "" {
            case "":
                if v5 {
                    switch "hello" {
                    case "a":
                        fmt.Println(v3, v3, v3)
                    case "foo bar":
                        fmt.Println(v3)
                    case "\n\t":
                        var v86 uint8 = 127
                        _ = v86
                    default:
                        if true {
                            v87 := GenS10{F0: 999, F1: 2}
                            _ = v87
                        }
                    }
                } else {
                    switch ("hello" + "hello") {
                    case "\n\t":
                        switch fi61 {
                        case 1:
                            if true {
                                v88 := GenS8{F0: -9223372036854775808}
                                for fi89 := 0; fi89 < 17; fi89++ {
                                    var v90 int8 = 127
                                    v90 += (v90 ^ v90)
                                    if (!(("hello" + "a") < ((("\n\t" + "foo bar") + "") + ("a" + "hello")))) {
                                        if (!("a" > ("\n\t" + ("\n\t" + "\n\t")))) {
                                            switch (fi61 % (((fi61 >> ((65535) & 63))) | 1)) {
                                            case -874:
                                                fmt.Println(v5, fi89)
                                            case 32767:
                                                v5 = (v90 > int8(fi61))
                                                if (false && true) {
                                                    fmt.Println(v5, fi89, v3)
                                                    var v91 string = ("" + (("foo bar" + ((((("foo bar" + "hello") + "hello") + "") + ((("foo bar" + ((("hello" + "hello") + "foo bar") + "foo bar")) + "a") + "a")) + "a")) + "hello"))
                                                    var v92 float64 = 0.1
                                                    v3 = v3
                                                    _ = v91
                                                    _ = v92
                                                } else {
                                                    fmt.Println(v3)
                                                    v3 -= 0
                                                    switch fi89 {
                                                    case 1:
                                                        v1 = (v1 / (((v2 / (((v2 + v1)) | 1))) | 1))
                                                    case -9223372036854775808:
                                                        var v93 uint8 = 226
                                                        _ = v93
                                                    case -339:
                                                        switch fi89 {
                                                        case -604:
                                                            v4 = append(v4, 1)
                                                        default:
                                                            v94 := GenS11{F0: 42, F1: 1, F2: v90}
                                                            _ = v94
                                                        }
                                                    }
                                                }
                                                if ((v5 || (v5 || (!(((("foo bar" + "foo bar") <= ("\n\t" + "a")) && (766 >= -1)) && (42 == 141))))) || v5) {
                                                    if (false || (1 <= v2)) {
                                                        v95 := GenS12{F0: 2}
                                                        _ = v95
                                                    }
                                                }
                                                var v96 int32 = -109
                                                _ = v96
                                            default:
                                                if (!true) {
                                                    if ((v90 & 100) > (-(v90))) {
                                                        if (0.0 < (float32((((((float64(3.14159) * (float64(((1.5 * (((0.0 / ((0.1 * (-2.25 + 1.5)) * 1.5)) / (0.1 / float32(1e10))) / 1.0)) * 0.1)) * 1.5)) + (0.1 - 1e10)) * float64(0.0)) / (-2.25 - (float64(1e10) + 3.14159))) * 3.14159)) / 0.1)) {
                                                            var v97 float32 = 3.14159
                                                            _ = v97
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                    } else {
                                        v1 = (v2 + v1)
                                    }
                                    var v98 uint16 = 1
                                    if v5 {
                                        if v5 {
                                            v2 -= (-(v2))
                                            var v99 int64 = 42
                                            _ = v99
                                        } else {
                                            var v100 int32 = 2147483647
                                            _ = v100
                                        }
                                        var v101 uint16 = 32767
                                        _ = v101
                                    } else {
                                        switch 0 {
                                        case 4294967295:
                                            v5 = v5
                                        case 49:
                                            v102 := GenS13{F0: 378, F1: -657, F2: 255}
                                            _ = v102
                                        }
                                        fmt.Println(fi61)
                                    }
                                    _ = len(v4)
                                    v1 += v1
                                    _ = len(v4)
                                    v98 = 2
                                    if (!v5) {
                                        v2 -= (-(v1))
                                        fmt.Println(fi61, fi61, v3)
                                    }
                                    _ = fi89
                                    _ = v90
                                    _ = v98
                                }
                                _ = v88
                            }
                        default:
                            if v5 {
                                var v103 uint16 = 691
                                switch (fi61 ^ 42) {
                                case -813:
                                    v2 = v1
                                case -2147483648:
                                    v1 += int16(fi61)
                                case 2147483647:
                                    for fi104 := 0; fi104 < 15; fi104++ {
                                        v3 -= (v3 / ((788) | 1))
                                        _ = fi104
                                    }
                                }
                                var v105 int8 = -128
                                _ = v103
                                _ = v105
                            } else {
                                if v5 {
                                    var v106 float32 = (3.14159 - ((float32(0.0) / float32(float64(3.14159))) * 3.14159))
                                    _ = v106
                                }
                                var v107 uint64 = 925
                                _ = v107
                            }
                        }
                    default:
                        var v108 uint16 = 144
                        _ = v108
                    }
                    _ = len(v4)
                    v109 := GenS13{F0: 2, F1: (-(v1)), F2: 42}
                    _ = v109
                }
            case "a":
                var v110 float64 = 3.14159
                _ = v110
            default:
                if ((v5 || (389 < 1)) && false) {
                    if v5 {
                        v111 := GenS11{F0: 42, F1: -1, F2: 127}
                        v112 := GenS11{F0: 0, F1: -9223372036854775808, F2: -128}
                        var v113 int8 = 101
                        _ = v111
                        _ = v112
                        _ = v113
                    } else {
                        if true {
                            var v114 int8 = -61
                            if (v3 != 2147483648) {
                                v115 := GenS6{F0: 42, F1: (!false), F2: 65535}
                                _ = v115
                            } else {
                                fmt.Println(v2, v3, v1)
                                switch ((("" + "") + "\n\t") + ("hello" + "\n\t")) {
                                case "\n\t":
                                    if v5 {
                                        _ = len(v4)
                                        v3 = v3
                                    } else {
                                        fmt.Println(v3, v1)
                                    }
                                case "":
                                    var v116 string = ""
                                    _ = v116
                                }
                                var v117 bool = true
                                _ = v117
                            }
                            switch fi61 {
                            case 32767:
                                v5 = true
                            }
                            _ = v114
                        }
                    }
                } else {
                    var v118 bool = v5
                    var v119 float64 = (0.1 * ((0.1 / 1.0) + (float64(0.0) / 0.1)))
                    for fi120 := 0; fi120 < 1; fi120++ {
                        if false {
                            fmt.Println(v118)
                            fmt.Println(v119, fi120)
                            switch "a" {
                            case "foo bar":
                                v2 += -32768
                            case "":
                                var v121 int32 = int32(v3)
                                var v122 string = ("a" + ("\n\t" + ""))
                                fmt.Println(fi61, v122, v2)
                                var v123 uint = 65535
                                v118 = v118
                                _ = v121
                                _ = v122
                                _ = v123
                            default:
                                fmt.Println(v119)
                            }
                        } else {
                            if (71 > -1) {
                                var v124 uint8 = 2
                                if true {
                                    v4 = append(v4, (v3 % (((v3 - (^(v3)))) | 1)))
                                } else {
                                    if (((v3 ^ 65535) != 2147483648) || ((4294967295 > 65535) || ((fi120 < fi120) && (42 > 255)))) {
                                        v119 -= (v119 * v119)
                                    }
                                    var v125 uint = 523
                                    v126 := GenS7{F0: 42}
                                    switch fi61 {
                                    case 9223372036854775807:
                                        _ = len(v4)
                                    }
                                    _ = v125
                                    _ = v126
                                }
                                fmt.Println(fi61)
                                _ = v124
                            } else {
                                v5 = (4294967295 < 255)
                                var v127 uint32 = 2147483647
                                switch (((("foo bar" + "\n\t") + "\n\t") + ("foo bar" + "a")) + "") {
                                case "foo bar":
                                    v128 := GenS12{F0: 246}
                                    _ = v128
                                case "a":
                                    var v129 int64 = -299
                                    _ = v129
                                case "\n\t":
                                    switch (fi120 >> ((0) & 63)) {
                                    case 32768:
                                        var v130 int = (^(fi61))
                                        _ = v130
                                    case 9223372036854775807:
                                        switch 2147483647 {
                                        case 76:
                                            v3 = v127
                                        case 128:
                                            var v131 uint = 65535
                                            _ = v131
                                        case 1:
                                            if v118 {
                                                if (!v5) {
                                                    if v5 {
                                                        switch "hello" {
                                                        case "foo bar":
                                                            var v132 uint32 = 2147483648
                                                            _ = v132
                                                        default:
                                                            v1 = v1
                                                        }
                                                        fmt.Println(v5, v5)
                                                        if v5 {
                                                            v2 += v2
                                                            v1 -= -214
                                                        } else {
                                                            v133 := GenS14{F0: 4294967295, F1: v3, F2: -1}
                                                            _ = v133
                                                        }
                                                        var v134 bool = (!v118)
                                                        _ = v134
                                                    }
                                                }
                                                var v135 bool = (!((217 == 19) && true))
                                                _ = v135
                                            } else {
                                                var v136 uint16 = 2
                                                v1 = 42
                                                var v137 float64 = -2.25
                                                _ = v136
                                                _ = v137
                                            }
                                            if len(v4) > 0 {
                                                v4[0] = 1
                                            }
                                            v4 = append(v4, (v3 ^ v127))
                                        }
                                    case 32767:
                                        v138 := GenS4{F0: 0, F1: 1}
                                        _ = v138
                                    default:
                                        fmt.Println(v118)
                                    }
                                }
                                v3 -= (v127 / ((65535) | 1))
                                _ = v127
                            }
                        }
                        _ = fi120
                    }
                    _ = v118
                    _ = v119
                }
            }
            if v5 {
                fmt.Println(fi61, fi61)
                fmt.Println(v5, fi61, v3)
            }
            _ = fi61
        }
    }
    if (false || (((v3 ^ v3) <= 427) || (!(42 != -342)))) {
        var v139 uint8 = 255
        switch 426 {
        case -2147483648:
            v2 += v1
        }
        v139 = v139
        _ = v139
    }
    v1 = (v1 + 32767)
    switch (("a" + "") + "\n\t") {
    case "foo bar":
        if v5 {
            v1 -= v2
            switch 0 {
            case 1:
                v2 += (v2 >> ((453) & 63))
            case -9223372036854775808:
                v140 := GenS15{F0: -69, F1: 42, F2: -686}
                _ = v140
            }
        }
    default:
        if (!(!v5)) {
            var v141 uint32 = v3
            _ = v141
        }
    }
    fmt.Println(v2)
    v4 = append(v4, (v3 & (v3 << ((18446744073709551615) & 63))))
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
