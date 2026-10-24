package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 uint16; F1 int32 }
type GenS5 struct { F0 uint8; F1 uint64 }
type GenS6 struct { F0 uint8; F1 uint }
type GenS7 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS8 struct { F0 int16; F1 int16 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 string; F1 int; F2 uint16 }
type GenS14 struct { F0 int8; F1 int; F2 int }
type GenS15 struct { F0 string; F1 int32 }
type GenS16 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS17 struct { F0 uint8 }
func genFn0(p0 int) uint8 {
    p0 = (p0 + -200)
    p0 = (p0 & (p0 >> ((18446744073709551615) & 63)))
    _ = p0
    return 133
}

func main() {
    v1 := []int{32768, -1, -1}
    var v2 float64 = 0.1
    var v3 uint64 = 526
    v4 := []string{}
    var v5 uint16 = 32767
    var v6 uint32 = 2
    if false {
        switch 689 {
        case -2147483648:
            switch 42 {
            case -572:
                v7 := GenS1{F0: v3}
                _ = v7
            case 117:
                var v8 int16 = 473
                _ = v8
            case -2147483648:
                if (uint16(v6) < (^(v5))) {
                    if ((v2 > (v2 * float64(0.1))) || true) {
                        switch -72 {
                        case -9223372036854775808:
                            fmt.Println(v5, v6)
                        default:
                            var v9 uint16 = 32767
                            _ = v9
                        }
                        switch 9223372036854775807 {
                        case -525:
                            if false {
                                v3 += (^(v3))
                                if len(v1) > 0 {
                                    v1[0] = 224
                                }
                                for fi10 := 0; fi10 < 16; fi10++ {
                                    if ((1.5 == 1e10) && true) {
                                        var v11 int16 = 42
                                        if (true || (!true)) {
                                            v11 = (v11 + (v11 >> ((258) & 63)))
                                        } else {
                                            v12 := GenS1{F0: v3}
                                            v13 := GenS2{F0: ("" + "\n\t"), F1: genFn0(fi10)}
                                            _ = v12
                                            _ = v13
                                        }
                                        switch "" {
                                        case "a":
                                            for fi14 := 0; fi14 < 14; fi14++ {
                                                _ = len(v1)
                                                _ = fi14
                                            }
                                        case "foo bar":
                                            v3 -= uint64(v6)
                                        case "hello":
                                            v15 := GenS2{F0: ("hello" + ("foo bar" + "a")), F1: 240}
                                            _ = v15
                                        default:
                                            v6 += 42
                                        }
                                        _ = v11
                                    }
                                    _ = fi10
                                }
                            }
                        }
                    } else {
                        var v16 int64 = 1
                        v17 := GenS3{F0: 1, F1: v3}
                        _ = v16
                        _ = v17
                    }
                    if true {
                        _ = len(v4)
                        switch -9223372036854775808 {
                        case 42:
                            v3 = v3
                        case 2147483647:
                            v6 = 63
                        }
                        _ = len(v4)
                    }
                }
            default:
                var v18 string = ("" + "foo bar")
                _ = v18
            }
        case 4294967295:
            v2 = ((v2 + v2) + (1e10 * v2))
        }
        if false {
            switch ("foo bar" + "\n\t") {
            case "foo bar":
                var v19 uint = 65535
                _ = v19
            case "a":
                _ = len(v1)
            default:
                if false {
                    v20 := GenS3{F0: 22, F1: (v3 * 355)}
                    _ = v20
                }
            }
            switch ("\n\t" + "\n\t") {
            case "hello":
                if (v5 != 65535) {
                    v4 = append(v4, "")
                } else {
                    switch "" {
                    case "":
                        v21 := GenS4{F0: (v5 << ((65535) & 63)), F1: -541}
                        _ = v21
                    case "a":
                        var v22 float64 = 0.1
                        _ = v22
                    }
                }
            case "":
                v23 := GenS1{F0: (v3 & (v3 - v3))}
                _ = v23
            case "foo bar":
                for fi24 := 0; fi24 < 13; fi24++ {
                    v25 := GenS5{F0: 42, F1: 18446744073709551615}
                    v26 := GenS1{F0: 42}
                    for fi27 := 0; fi27 < 20; fi27++ {
                        v28 := GenS6{F0: 217, F1: 4294967295}
                        var v29 bool = true
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
            for fi30 := 0; fi30 < 10; fi30++ {
                switch fi30 {
                case 0:
                    _ = len(v1)
                case 127:
                    v31 := GenS6{F0: genFn0((fi30 << ((1) & 63))), F1: 0}
                    _ = v31
                case 128:
                    var v32 float64 = -2.25
                    _ = v32
                default:
                    v2 = v2
                }
                _ = fi30
            }
        }
    } else {
        switch "\n\t" {
        case "\n\t":
            v33 := GenS1{F0: v3}
            _ = v33
        case "a":
            v5 -= (v5 & v5)
        case "foo bar":
            fmt.Println(v3, v6, v2)
        default:
            switch 571 {
            case -9223372036854775808:
                v34 := GenS4{F0: (^(v5)), F1: 1}
                _ = v34
            case -125:
                v6 = v6
            default:
                for fi35 := 0; fi35 < 18; fi35++ {
                    v2 -= (v2 + (((3.14159 + v2) + v2) / ((1.5 - ((v2 / v2) * 0.1)) + v2)))
                    _ = fi35
                }
            }
        }
    }
    if ((1.5 - -2.25) != float32(v2)) {
        if ((2 == 42) || true) {
            switch "foo bar" {
            case "a":
                v4 = append(v4, (("" + "hello") + ("a" + "\n\t")))
            case "":
                if (!false) {
                    var v36 int64 = -9223372036854775808
                    _ = v36
                } else {
                    for fi37 := 0; fi37 < 4; fi37++ {
                        v38 := GenS2{F0: ("foo bar" + ""), F1: 176}
                        v3 = 594
                        fmt.Println(v6, v6, v6)
                        _ = fi37
                        _ = v38
                    }
                    if true {
                        v3 += (v3 >> ((65535) & 63))
                    } else {
                        var v39 float64 = v2
                        _ = v39
                    }
                    v2 = v2
                }
            default:
                var v40 float32 = (3.14159 * 1.5)
                _ = v40
            }
            v41 := GenS7{F0: 0, F1: 1, F2: (v5 | v5)}
            _ = v41
        }
    } else {
        v2 = (1e10 - v2)
    }
    v3 -= v3
    var v42 string = ("\n\t" + "\n\t")
    if false {
        v43 := GenS4{F0: v5, F1: 890}
        _ = v43
    } else {
        switch v42 {
        case "\n\t":
            if (-123 != 9) {
                for fi44 := 0; fi44 < 1; fi44++ {
                    for fi45 := 0; fi45 < 17; fi45++ {
                        _ = len(v1)
                        switch ("hello" + v42) {
                        case "foo bar":
                            if (-153 != -2147483648) {
                                var v46 bool = false
                                v47 := GenS5{F0: 83, F1: (^(v3))}
                                var v48 int8 = 1
                                _ = v46
                                _ = v47
                                _ = v48
                            } else {
                                switch v42 {
                                case "foo bar":
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case "\n\t":
                                    fmt.Println(v3, v3)
                                case "":
                                    v49 := GenS8{F0: int16(v5), F1: 1}
                                    _ = v49
                                default:
                                    switch v42 {
                                    case "\n\t":
                                        var v50 bool = false
                                        _ = v50
                                    case "a":
                                        _ = len(v4)
                                        v6 = v6
                                    case "hello":
                                        v51 := GenS9{F0: "hello"}
                                        _ = v51
                                    }
                                }
                                var v52 bool = true
                                var v53 float32 = ((-2.25 * 0.1) / float32(-2.25))
                                var v54 uint32 = (v6 << ((65535) & 63))
                                _ = v52
                                _ = v53
                                _ = v54
                            }
                        case "hello":
                            if len(v1) > 0 {
                                v1[0] = (fi45 ^ fi45)
                            }
                            v5 += (v5 * (v5 / ((v5) | 1)))
                            fmt.Println(v6)
                        }
                        v5 += v5
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi45
                    }
                    if (!false) {
                        fmt.Println(fi44)
                    }
                    v6 -= (v6 ^ v6)
                    _ = fi44
                }
                v2 = ((0.0 + 0.1) * v2)
                for fi55 := 0; fi55 < 2; fi55++ {
                    v2 += ((((1e10 - (float64(1.5) + 0.0)) / v2) / v2) / v2)
                    if false {
                        var v56 uint16 = v5
                        v56 = v56
                        v57 := GenS8{F0: -32768, F1: 1}
                        _ = v56
                        _ = v57
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            if ((0.0 == 1.5) || (1 == 915)) {
                                for fi58 := 0; fi58 < 17; fi58++ {
                                    fmt.Println(v42, v42)
                                    var v59 bool = true
                                    if false {
                                        v2 = v2
                                    }
                                    _ = fi58
                                    _ = v59
                                }
                            }
                        case "a":
                            v60 := GenS10{F0: 16}
                            _ = v60
                        case "\n\t":
                            v2 = v2
                        default:
                            if ((1.5 + 0.0) <= float32(float64(3.14159))) {
                                if true {
                                    var v61 int8 = 1
                                    fmt.Println(v6, v6)
                                    if true {
                                        if true {
                                            for fi62 := 0; fi62 < 5; fi62++ {
                                                if (42 >= (^(v6))) {
                                                    v42 = v42
                                                    var v63 float64 = v2
                                                    _ = v63
                                                } else {
                                                    v42 = (v42 + "hello")
                                                    fmt.Println(fi55)
                                                    if true {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        switch v42 {
                                                        case "hello":
                                                            v64 := GenS11{F0: (v3 % (((^(v3))) | 1))}
                                                            _ = v64
                                                        case "a":
                                                            var v65 int32 = -2147483648
                                                            _ = v65
                                                        case "\n\t":
                                                            var v66 uint16 = 524
                                                            _ = v66
                                                        default:
                                                            fmt.Println(v42, v3, v42)
                                                        }
                                                    } else {
                                                        v67 := GenS1{F0: (v3 | 4294967295)}
                                                        switch "" {
                                                        case "a":
                                                            v68 := GenS8{F0: 0, F1: 42}
                                                            _ = v68
                                                        default:
                                                            fmt.Println(fi62, v6)
                                                        }
                                                        _ = v67
                                                    }
                                                }
                                                var v69 uint8 = 211
                                                _ = len(v4)
                                                _ = fi62
                                                _ = v69
                                            }
                                            fmt.Println(v6)
                                        }
                                        var v70 int32 = 646
                                        _ = v70
                                    }
                                    _ = v61
                                }
                                if (!true) {
                                    v2 = (((v2 - 1e10) - 1e10) + (v2 - ((1.5 - v2) - (v2 - 0.1))))
                                }
                                var v71 uint8 = 7
                                _ = v71
                            }
                        }
                        v42 = v42
                    }
                    _ = fi55
                }
            }
        case "":
            fmt.Println(v6, v5)
        }
    }
    v3 += v3
    switch 689 {
    case -416:
        fmt.Println(v2, v5, v3)
    case 1:
        var v72 float32 = -2.25
        _ = v72
    case 32767:
        for fi73 := 0; fi73 < 5; fi73++ {
            fmt.Println(v2, fi73)
            if false {
                var v74 int32 = 42
                _ = v74
            } else {
                if true {
                    for fi75 := 0; fi75 < 14; fi75++ {
                        fmt.Println(v5, v3)
                        v6 += 874
                        switch (v42 + "") {
                        case "foo bar":
                            v5 -= v5
                            if false {
                                fmt.Println(v42)
                            } else {
                                if true {
                                    _ = len(v1)
                                    fmt.Println(v5, v2, fi73)
                                    fmt.Println(v6)
                                    v5 -= uint16(v3)
                                } else {
                                    v3 = v3
                                    v76 := GenS3{F0: 100, F1: v3}
                                    v5 += (v5 ^ (v5 / ((v5) | 1)))
                                    _ = v76
                                }
                                v2 = (v2 + ((v2 + -2.25) - v2))
                                if false {
                                    var v77 string = v42
                                    v5 -= 742
                                    v2 -= 1.5
                                    var v78 int64 = 42
                                    _ = v77
                                    _ = v78
                                }
                            }
                            var v79 int32 = 0
                            if false {
                                v4 = append(v4, (((("\n\t" + "foo bar") + v42) + "foo bar") + ""))
                                if ((-(v79)) == 1) {
                                    v79 = 0
                                    if (v3 <= 924) {
                                        if (false || true) {
                                            v80 := GenS4{F0: (v5 - (v5 >> ((1) & 63))), F1: -1}
                                            v5 = v5
                                            var v81 int = fi75
                                            _ = v80
                                            _ = v81
                                        }
                                    } else {
                                        fmt.Println(fi75)
                                        v82 := GenS12{F0: -1}
                                        switch (("hello" + v42) + "foo bar") {
                                        case "":
                                            if len(v1) > 0 {
                                                v1[0] = (^(fi73))
                                            }
                                        case "a":
                                            var v83 uint16 = (v5 * 137)
                                            _ = v83
                                        case "hello":
                                            if (-69 < 42) {
                                                fmt.Println(v42)
                                            } else {
                                                var v84 float64 = (v2 + float64(-2.25))
                                                v3 -= v3
                                                _ = v84
                                            }
                                        default:
                                            v42 = ("foo bar" + "")
                                        }
                                        _ = v82
                                    }
                                }
                                if ((((v42 >= v42) || false) && false) || true) {
                                    v85 := GenS13{F0: ("a" + ("" + v42)), F1: (fi75 - (fi73 & int(v5))), F2: v5}
                                    _ = v85
                                } else {
                                    var v86 string = "hello"
                                    v42 = ("\n\t" + ((v86 + v86) + "\n\t"))
                                    fmt.Println(v79, v6)
                                    _ = v86
                                }
                            } else {
                                v5 = (v5 * v5)
                                v79 += (v79 % ((v79) | 1))
                            }
                            _ = v79
                        default:
                            var v87 uint16 = 1
                            _ = v87
                        }
                        _ = fi75
                    }
                    fmt.Println(v6, v6)
                }
                v6 += 404
                var v88 string = ("a" + v42)
                _ = v88
            }
            v2 = v2
            _ = fi73
        }
    default:
        switch v42 {
        case "foo bar":
            for fi89 := 0; fi89 < 5; fi89++ {
                var v90 int64 = -1
                v5 = (v5 * (v5 * v5))
                switch "" {
                case "":
                    if ((1 == 2) && false) {
                        if true {
                            v91 := GenS14{F0: -128, F1: (fi89 % ((fi89) | 1)), F2: -255}
                            for fi92 := 0; fi92 < 2; fi92++ {
                                if (fi92 >= (fi89 ^ -1)) {
                                    v42 = v42
                                    v2 += float64(float32((v2 / (v2 - 3.14159))))
                                } else {
                                    v5 += (v5 << ((0) & 63))
                                    if len(v1) > 0 {
                                        v1[0] = int(v5)
                                    }
                                    var v93 uint64 = v3
                                    _ = v93
                                }
                                fmt.Println(fi89, v90, v6)
                                if (!(((^(v5)) <= 0) || (248 == 240))) {
                                    switch v42 {
                                    case "\n\t":
                                        v94 := GenS12{F0: (v90 | v90)}
                                        _ = v94
                                    }
                                    var v95 string = ("foo bar" + v42)
                                    _ = len(v4)
                                    _ = v95
                                } else {
                                    var v96 int64 = v90
                                    var v97 bool = false
                                    var v98 uint64 = v3
                                    var v99 uint64 = (v98 - v98)
                                    v2 -= v2
                                    _ = v96
                                    _ = v97
                                    _ = v98
                                    _ = v99
                                }
                                var v100 string = "\n\t"
                                fmt.Println(v6, v100, v6)
                                if (true || (!(fi92 >= (fi89 / (((fi92 / (((fi89 | fi92)) | 1))) | 1))))) {
                                    if (0 == uint(v5)) {
                                        if false {
                                            var v101 int32 = 1
                                            v4 = append(v4, v100)
                                            _ = v101
                                        } else {
                                            if false {
                                                v3 = v3
                                                fmt.Println(v42, v100)
                                                var v102 int = (fi92 % ((fi92) | 1))
                                                _ = v102
                                            }
                                            v90 += (^(v90))
                                            var v103 int8 = 0
                                            _ = v103
                                        }
                                        v3 -= uint64(v6)
                                        var v104 uint8 = 255
                                        v105 := GenS15{F0: v100, F1: 676}
                                        _ = v104
                                        _ = v105
                                    }
                                    v106 := GenS12{F0: -751}
                                    var v107 uint32 = 65535
                                    v5 = 2
                                    if len(v1) > 0 {
                                        v1[0] = (fi92 ^ int(v3))
                                    }
                                    _ = v106
                                    _ = v107
                                } else {
                                    v100 = v100
                                }
                                v108 := GenS8{F0: 1000, F1: 32767}
                                if (-2147483648 <= 0) {
                                    switch fi92 {
                                    case -498:
                                        v109 := GenS13{F0: v42, F1: fi89, F2: (^(v5))}
                                        _ = v109
                                    case 323:
                                        var v110 uint64 = (v3 << ((629) & 63))
                                        _ = v110
                                    case 2147483647:
                                        fmt.Println(v6)
                                    }
                                } else {
                                    v111 := GenS16{F0: 174, F1: (v5 >> ((18446744073709551615) & 63)), F2: v5}
                                    var v112 int64 = (-(v90))
                                    _ = v111
                                    _ = v112
                                }
                                var v113 int64 = 221
                                fmt.Println(v42, v2)
                                if len(v4) > 0 {
                                    v4[0] = v100
                                }
                                _ = fi92
                                _ = v100
                                _ = v108
                                _ = v113
                            }
                            _ = v91
                        } else {
                            for fi114 := 0; fi114 < 18; fi114++ {
                                var v115 uint16 = v5
                                switch v42 {
                                case "\n\t":
                                    v42 = v42
                                case "hello":
                                    if false {
                                        v4 = append(v4, "")
                                        fmt.Println(fi114, v5, v115)
                                        var v116 bool = true
                                        _ = v116
                                    }
                                case "foo bar":
                                    if true {
                                        if false {
                                            if (!false) {
                                                var v117 uint64 = 18446744073709551615
                                                _ = len(v4)
                                                var v118 int16 = -401
                                                _ = v117
                                                _ = v118
                                            } else {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                v6 = 2
                                            }
                                        } else {
                                            var v119 int16 = 32767
                                            if (1 != 0) {
                                                var v120 int64 = -1
                                                v1 = append(v1, fi114)
                                                var v121 uint = 65535
                                                _ = v120
                                                _ = v121
                                            }
                                            var v122 uint64 = v3
                                            _ = v119
                                            _ = v122
                                        }
                                    }
                                }
                                _ = fi114
                                _ = v115
                            }
                            for fi123 := 0; fi123 < 4; fi123++ {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v90 = v90
                                _ = fi123
                            }
                        }
                    } else {
                        if true {
                            fmt.Println(v2, v2, v6)
                            if true {
                                v42 = v42
                            }
                            for fi124 := 0; fi124 < 17; fi124++ {
                                v125 := GenS17{F0: 223}
                                var v126 int = fi124
                                var v127 uint8 = genFn0(128)
                                _ = fi124
                                _ = v125
                                _ = v126
                                _ = v127
                            }
                        } else {
                            if len(v4) > 0 {
                                v4[0] = (("a" + v42) + v42)
                            }
                        }
                        var v128 string = v42
                        if true {
                            var v129 int8 = 1
                            var v130 int64 = v90
                            var v131 uint64 = (v3 >> ((42) & 63))
                            _ = v129
                            _ = v130
                            _ = v131
                        } else {
                            var v132 uint8 = genFn0((fi89 + fi89))
                            _ = v132
                        }
                        _ = v128
                    }
                case "\n\t":
                    switch fi89 {
                    case 161:
                        v90 = -1
                    }
                case "foo bar":
                    if (true && true) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v133 string = ("foo bar" + v42)
                        _ = v133
                    } else {
                        v42 = "foo bar"
                    }
                }
                _ = fi89
                _ = v90
            }
        }
    }
    if false {
        for fi134 := 0; fi134 < 1; fi134++ {
            var v135 int32 = 2147483647
            v1 = append(v1, (fi134 / ((fi134) | 1)))
            if len(v4) > 0 {
                v4[0] = ""
            }
            _ = fi134
            _ = v135
        }
        v5 -= uint16(v3)
    }
    v42 = v42
    fmt.Println(v3)
    var v136 uint16 = v5
    v2 = (0.0 + (v2 + (((0.0 + v2) / (v2 * v2)) / v2)))
    v3 += v3
    fmt.Println(v2, v3, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v136:", v136)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v136
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
