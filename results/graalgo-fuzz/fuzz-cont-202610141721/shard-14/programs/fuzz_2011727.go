package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int8; F1 uint8; F2 int }
type GenS3 struct { F0 int16; F1 string; F2 int8 }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint64; F1 int8 }
type GenS12 struct { F0 uint32 }

func main() {
    v1 := []uint32{2147483648, 11}
    var v2 uint8 = 2
    v3 := make(map[int]bool)
    v4 := make(map[uint32]string)
    var v5 int8 = 1
    v5 -= (v5 % (((v5 | (v5 ^ -128))) | 1))
    delete(v4, 818)
    if true {
        switch ("" + "foo bar") {
        case "foo bar":
            if true {
                switch 127 {
                case 958:
                    v6 := GenS1{F0: "\n\t", F1: (v2 & (^(v2)))}
                    _ = v6
                default:
                    if false {
                        v2 -= (^(v2))
                        v7 := GenS1{F0: "foo bar", F1: v2}
                        _ = v7
                    }
                }
                if true {
                    v8 := GenS1{F0: (("" + "foo bar") + "\n\t"), F1: v2}
                    var v9 float64 = -2.25
                    if false {
                        v1 = append(v1, uint32(v2))
                    } else {
                        var v10 float64 = v9
                        var v11 bool = (75 <= v2)
                        _ = v10
                        _ = v11
                    }
                    _ = v8
                    _ = v9
                } else {
                    for fi12 := 0; fi12 < 19; fi12++ {
                        v13 := GenS1{F0: "a", F1: 42}
                        for fi14 := 0; fi14 < 7; fi14++ {
                            var v15 uint16 = 0
                            _ = len(v3)
                            _ = fi14
                            _ = v15
                        }
                        v16 := GenS2{F0: (v5 >> ((4294967295) & 63)), F1: v2, F2: fi12}
                        _ = fi12
                        _ = v13
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 16; fi17++ {
                        var v18 int = fi17
                        var v19 float32 = 1.5
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    if (true || false) {
                        switch (("a" + "a") + "a") {
                        case "foo bar":
                            v2 += v2
                        case "":
                            for fi20 := 0; fi20 < 17; fi20++ {
                                v21 := GenS1{F0: "a", F1: v2}
                                v22 := GenS3{F0: 1000, F1: "\n\t", F2: (-(v5))}
                                _ = fi20
                                _ = v21
                                _ = v22
                            }
                        case "\n\t":
                            if (v5 >= 127) {
                                if ((-1 > -346) && false) {
                                    var v23 int64 = 1
                                    fmt.Println(v2, v5)
                                    for fi24 := 0; fi24 < 2; fi24++ {
                                        var v25 string = (("a" + "hello") + ((("a" + "") + (("\n\t" + "foo bar") + ("foo bar" + "a"))) + "\n\t"))
                                        fmt.Println(v5, v25)
                                        _ = fi24
                                        _ = v25
                                    }
                                    _ = v23
                                }
                                var v26 bool = true
                                fmt.Println(v26, v26)
                                _ = v26
                            } else {
                                if true {
                                    v5 = 42
                                }
                                var v27 int = 32767
                                switch ("hello" + "") {
                                case "":
                                    for fi28 := 0; fi28 < 5; fi28++ {
                                        var v29 float64 = (((0.0 + (3.14159 + (((0.0 / (1.5 + 3.14159)) + ((0.0 - 1.5) * (3.14159 + 1e10))) - float64((0.1 + 0.0))))) / 0.1) * (float64(0.1) / ((1e10 - (3.14159 - 0.1)) / 1.5)))
                                        var v30 uint16 = 1
                                        v4[411] = ("a" + "foo bar")
                                        _ = fi28
                                        _ = v29
                                        _ = v30
                                    }
                                }
                                _ = v27
                            }
                        }
                    }
                }
            } else {
                v31 := GenS3{F0: 1, F1: "a", F2: (-(v5))}
                v2 = 0
                _ = v31
            }
        case "hello":
            var v32 int = 2147483647
            _ = v32
        case "\n\t":
            v33 := GenS4{F0: 42, F1: 2}
            _ = v33
        }
        var v34 float32 = -2.25
        for fi35 := 0; fi35 < 7; fi35++ {
            for fi36 := 0; fi36 < 12; fi36++ {
                v2 += 1
                var v37 bool = true
                _ = fi36
                _ = v37
            }
            for fi38 := 0; fi38 < 8; fi38++ {
                var v39 uint64 = 18446744073709551615
                fmt.Println(v2)
                _ = fi38
                _ = v39
            }
            _ = fi35
        }
        _ = v34
    } else {
        fmt.Println(v2, v2)
        var v40 int64 = 0
        _ = v40
    }
    for fi41 := 0; fi41 < 15; fi41++ {
        var v42 int16 = 943
        v1 = append(v1, 0)
        var v43 int8 = 42
        _ = fi41
        _ = v42
        _ = v43
    }
    if (true || ((-2.25 * (3.14159 - float32(-2.25))) > (1e10 - 1e10))) {
        v44 := GenS2{F0: v5, F1: 234, F2: 1}
        if len(v1) > 0 {
            v1[0] = 42
        }
        fmt.Println(v5)
        _ = v44
    } else {
        if false {
            for fi45 := 0; fi45 < 8; fi45++ {
                fmt.Println(v5, fi45, v5)
                var v46 int = (fi45 + (-(fi45)))
                for fi47 := 0; fi47 < 11; fi47++ {
                    v46 = v46
                    switch 946 {
                    case 4294967295:
                        if (-9 > -35) {
                            v48 := GenS2{F0: (v5 ^ v5), F1: (v2 + v2), F2: 9223372036854775807}
                            var v49 float32 = float32(0.1)
                            v50 := GenS5{F0: v2}
                            _ = v48
                            _ = v49
                            _ = v50
                        }
                    case -2147483648:
                        fmt.Println(fi47, v5, v5)
                    case 224:
                        if false {
                            var v51 uint32 = 4294967295
                            _ = v51
                        } else {
                            var v52 int = fi45
                            if ((true && true) && false) {
                                var v53 uint32 = 2147483648
                                _ = v53
                            } else {
                                switch (("\n\t" + ("foo bar" + "")) + "a") {
                                case "":
                                    if len(v1) > 0 {
                                        v1[0] = 133
                                    }
                                }
                                v5 -= v5
                                var v54 uint8 = (^(v2))
                                _ = v54
                            }
                            if len(v1) > 0 {
                                v1[0] = 443
                            }
                            v55 := GenS3{F0: -32768, F1: ("\n\t" + ("foo bar" + (("\n\t" + "") + ""))), F2: v5}
                            _ = v52
                            _ = v55
                        }
                    }
                    _ = fi47
                }
                _ = fi45
                _ = v46
            }
            switch 4 {
            case -9223372036854775808:
                var v56 float64 = 0.0
                _ = v56
            case 127:
                for fi57 := 0; fi57 < 7; fi57++ {
                    var v58 uint32 = 905
                    _ = fi57
                    _ = v58
                }
            }
            if false {
                switch "\n\t" {
                case "hello":
                    fmt.Println(v5, v5)
                case "a":
                    v1 = append(v1, 42)
                }
                for fi59 := 0; fi59 < 7; fi59++ {
                    fmt.Println(v2, fi59)
                    _ = fi59
                }
                var v60 float64 = float64((1e10 * 1e10))
                _ = v60
            }
        }
    }
    var v61 int16 = -1
    var v62 uint = 65535
    for fi63 := 0; fi63 < 6; fi63++ {
        fmt.Println(v61)
        _ = fi63
    }
    fmt.Println(v5, v2, v2)
    if (!(42 < 1)) {
        for fi64 := 0; fi64 < 4; fi64++ {
            var v65 uint = v62
            v3[fi64] = (!true)
            fmt.Println(v5, v5, v5)
            _ = fi64
            _ = v65
        }
        v66 := GenS2{F0: v5, F1: 0, F2: 2147483647}
        v62 += v62
        _ = v66
    }
    for fi67 := 0; fi67 < 5; fi67++ {
        v68 := GenS6{F0: (v2 & v2)}
        v2 = v2
        var v69 bool = (!((int(v61) > fi67) || false))
        _ = fi67
        _ = v68
        _ = v69
    }
    for fi70 := 0; fi70 < 20; fi70++ {
        if (!false) {
            v2 -= v2
            v2 -= (v2 / ((128) | 1))
        } else {
            for fi71 := 0; fi71 < 13; fi71++ {
                v5 = v5
                var v72 uint = 405
                fmt.Println(fi71, v72)
                _ = fi71
                _ = v72
            }
        }
        _ = len(v1)
        switch -9223372036854775808 {
        case 32767:
            for fi73 := 0; fi73 < 10; fi73++ {
                if (true && true) {
                    v61 = 32767
                    if false {
                        if (!(-2 >= 649)) {
                            v2 = v2
                            if true {
                                v61 += v61
                                if (((!false) || false) || (v62 != v62)) {
                                    v74 := GenS7{F0: uint64(v62)}
                                    _ = v74
                                }
                                v5 = (-(v5))
                            }
                        }
                        var v75 int16 = -1
                        v62 -= v62
                        v75 = (v61 + v75)
                        if false {
                            v62 -= v62
                            var v76 uint = (v62 / (((v62 ^ (v62 / ((0) | 1)))) | 1))
                            _ = v76
                        }
                        _ = v75
                    }
                    if true {
                        v2 += (^(v2))
                        fmt.Println(v61)
                    } else {
                        if false {
                            switch (fi70 ^ int(v61)) {
                            case 2147483647:
                                fmt.Println(fi73, fi70)
                            case -2147483648:
                                if len(v1) > 0 {
                                    v1[0] = 2
                                }
                            case 159:
                                v62 += 255
                            default:
                                if true {
                                    v62 = v62
                                    var v77 uint64 = 42
                                    _ = v77
                                } else {
                                    var v78 uint32 = 848
                                    _ = v78
                                }
                            }
                            v5 = 127
                            if ((float32(1.5) < (-2.25 - 3.14159)) || true) {
                                v62 = (v62 | 42)
                                var v79 int = fi73
                                if len(v1) > 0 {
                                    v1[0] = 204
                                }
                                fmt.Println(v61)
                                _ = v79
                            }
                        }
                        v3[fi73] = (!false)
                        if ((fi70 / (((fi70 >> (((v62 - v62)) & 63))) | 1)) > fi70) {
                            var v80 int32 = 2147483647
                            var v81 uint = v62
                            if (42 < uint64(v2)) {
                                v5 -= 100
                            }
                            _ = v80
                            _ = v81
                        } else {
                            v5 -= -67
                            v62 -= 2
                            fmt.Println(v61, v2)
                        }
                        var v82 uint8 = (v2 & (v2 + (v2 >> ((665) & 63))))
                        _ = v82
                    }
                } else {
                    v61 = v61
                    v83 := GenS3{F0: (v61 % (((v61 % (((v61 * (v61 ^ int16(v2)))) | 1))) | 1)), F1: "hello", F2: v5}
                    v84 := GenS8{F0: -32768}
                    _ = v83
                    _ = v84
                }
                var v85 uint32 = 65535
                _ = fi73
                _ = v85
            }
        case 9223372036854775807:
            if false {
                var v86 int16 = (-(v61))
                var v87 float32 = -2.25
                if (303 <= 83) {
                    v88 := GenS3{F0: -32768, F1: "hello", F2: v5}
                    for fi89 := 0; fi89 < 6; fi89++ {
                        v90 := GenS9{F0: 4294967295}
                        switch -689 {
                        case 0:
                            if (((3.14159 - (((0.0 + 0.1) + 0.1) / 1.0)) - (1.5 / (1e10 - 3.14159))) < 0.1) {
                                var v91 float32 = (v87 * (float32(1.5) + ((v87 - (-2.25 - v87)) - ((v87 - v87) - v87))))
                                var v92 int32 = -2147483648
                                switch int(v2) {
                                case 9223372036854775807:
                                    v93 := GenS10{F0: 242}
                                    v94 := GenS10{F0: 2}
                                    _ = v93
                                    _ = v94
                                case 490:
                                    var v95 uint8 = (v2 | v2)
                                    _ = v95
                                }
                                _ = v91
                                _ = v92
                            } else {
                                var v96 bool = false
                                switch (fi89 * (fi70 >> ((0) & 63))) {
                                case 2147483647:
                                    var v97 string = "hello"
                                    _ = v97
                                default:
                                    switch "a" {
                                    case "foo bar":
                                        var v98 string = "a"
                                        _ = v98
                                    default:
                                        var v99 float64 = 1e10
                                        _ = v99
                                    }
                                }
                                switch ("a" + "hello") {
                                case "":
                                    switch "hello" {
                                    case "a":
                                        if (v61 <= v86) {
                                            fmt.Println(fi89)
                                            switch "hello" {
                                            case "foo bar":
                                                v87 = (float32((1.5 - 0.1)) * 0.0)
                                            }
                                            var v100 int32 = -969
                                            _ = v100
                                        } else {
                                            if v96 {
                                                v101 := GenS3{F0: (v61 - v61), F1: "a", F2: v5}
                                                v61 = (v86 | 1000)
                                                _ = v101
                                            } else {
                                                v3[(fi70 | 4294967295)] = false
                                                v102 := v3[fi89]
                                                fmt.Println(fi70, v86, fi89)
                                                _ = v102
                                            }
                                            v103 := v3[(fi89 | fi89)]
                                            if false {
                                                v104 := GenS3{F0: v86, F1: "a", F2: (-(v5))}
                                                fmt.Println(v62, v2, v62)
                                                _ = v104
                                            }
                                            _ = v103
                                        }
                                    case "\n\t":
                                        v105 := GenS9{F0: 4294967295}
                                        _ = v105
                                    case "foo bar":
                                        if v96 {
                                            v86 -= 32767
                                            fmt.Println(fi89, v61)
                                        } else {
                                            v106 := GenS5{F0: v2}
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            var v107 int8 = 0
                                            v108 := GenS11{F0: 529, F1: 1}
                                            _ = v106
                                            _ = v107
                                            _ = v108
                                        }
                                    }
                                case "hello":
                                    fmt.Println(v5, fi70, v5)
                                default:
                                    v5 = v5
                                }
                                _ = v96
                            }
                        case -9223372036854775808:
                            var v109 float64 = (3.14159 / 0.1)
                            _ = v109
                        case 128:
                            switch (("foo bar" + "hello") + "") {
                            case "a":
                                var v110 uint = v62
                                _ = v110
                            case "":
                                if true {
                                    v111 := GenS12{F0: 0}
                                    var v112 uint32 = 2147483647
                                    v113 := GenS3{F0: v61, F1: "\n\t", F2: v5}
                                    _ = v111
                                    _ = v112
                                    _ = v113
                                } else {
                                    v87 += v87
                                }
                            case "hello":
                                var v114 int8 = (v5 - (v5 & v5))
                                _ = v114
                            }
                        default:
                            switch "\n\t" {
                            case "":
                                v2 -= (^(v2))
                            default:
                                switch ("a" + ("hello" + "a")) {
                                case "":
                                    v115 := GenS1{F0: ("\n\t" + "hello"), F1: (v2 + (v2 >> (((^(v62))) & 63)))}
                                    _ = v115
                                case "\n\t":
                                    v116 := v4[65535]
                                    _ = v116
                                case "foo bar":
                                    v2 = (v2 + v2)
                                }
                            }
                        }
                        var v117 int16 = v61
                        _ = fi89
                        _ = v117
                        _ = v90
                    }
                    _ = v88
                }
                _ = v86
                _ = v87
            } else {
                var v118 int = fi70
                if ((false && false) && true) {
                    if (372 < 2147483647) {
                        v62 = v62
                        if false {
                            var v119 uint32 = 42
                            _ = v119
                        } else {
                            var v120 float64 = (1e10 / 1e10)
                            v118 = (v118 % ((9223372036854775807) | 1))
                            _ = v120
                        }
                        for fi121 := 0; fi121 < 9; fi121++ {
                            switch fi70 {
                            case 127:
                                var v122 int8 = 42
                                _ = v122
                            case 287:
                                delete(v3, (v118 ^ -9223372036854775808))
                            case 32768:
                                if true {
                                    v123 := v3[481]
                                    fmt.Println(v5, v2)
                                    _ = v123
                                } else {
                                    delete(v3, (v118 * fi70))
                                    var v124 int64 = 0
                                    _ = v124
                                }
                            }
                            _ = fi121
                        }
                    } else {
                        var v125 int32 = -1
                        for fi126 := 0; fi126 < 14; fi126++ {
                            v61 = (-(v61))
                            fmt.Println(fi126, v62)
                            _ = fi126
                        }
                        _ = v125
                    }
                } else {
                    v4[2] = ""
                    for fi127 := 0; fi127 < 15; fi127++ {
                        var v128 int32 = 0
                        _ = fi127
                        _ = v128
                    }
                }
                v2 -= 127
                _ = v118
            }
        case -2147483648:
            switch ("" + "\n\t") {
            case "foo bar":
                v2 = v2
            }
        default:
            switch "a" {
            case "a":
                var v129 uint64 = 4294967295
                _ = v129
            default:
                var v130 bool = (!false)
                _ = v130
            }
        }
        _ = fi70
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v61
    _ = v62
}
