package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 int16; F2 int64 }
type GenS5 struct { F0 bool; F1 string; F2 bool }
type GenS6 struct { F0 bool; F1 bool; F2 string }
type GenS7 struct { F0 int32; F1 int8; F2 int16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 bool; F1 uint64; F2 uint }
type GenS12 struct { F0 bool; F1 int; F2 string }
type GenS13 struct { F0 uint8 }
type GenS14 struct { F0 uint; F1 uint32; F2 int64 }
func genFn0(p0 int32, p1 int, p2 int16) int16 {
    p0 = p0
    p2 -= (p2 % ((p2) | 1))
    p1 = p1
    v4 := GenS1{F0: 0, F1: false}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return (p2 + p2)
}
func genFn1(p0 int32, p1 int16) int {
    v5 := GenS2{F0: 0}
    var v6 uint64 = 143
    if ((-(p0)) >= (^(p0))) {
        fmt.Println(p0, v6)
    } else {
        var v7 int32 = (p0 >> ((0) & 63))
        _ = v7
    }
    if false {
        var v8 uint64 = (v6 << ((4294967295) & 63))
        for fi9 := 0; fi9 < 2; fi9++ {
            switch (fi9 + fi9) {
            case -571:
                v8 = 459
            }
            _ = fi9
        }
        v10 := GenS2{F0: -128}
        _ = v10
        _ = v8
    } else {
        var v11 int8 = -1
        var v12 bool = true
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 4294967295
}

func main() {
    var v1 int = 140
    v2 := []int{127}
    v3 := []string{"\n\t", "foo bar"}
    if (((("hello" + "foo bar") + (("a" + "") + "foo bar")) >= ("" + ("a" + "foo bar"))) || (-1 != -128)) {
        fmt.Println(v1)
        for fi13 := 0; fi13 < 13; fi13++ {
            if len(v2) > 0 {
                v2[0] = v1
            }
            v14 := GenS3{F0: 4294967295}
            _ = fi13
            _ = v14
        }
        switch (v1 / ((v1) | 1)) {
        case -923:
            v15 := GenS4{F0: 880, F1: -32768, F2: -489}
            _ = v15
        default:
            fmt.Println(v1)
        }
    }
    switch ("" + "foo bar") {
    case "foo bar":
        v16 := GenS3{F0: 2}
        _ = v16
    }
    if (255 == 32767) {
        fmt.Println(v1)
        v1 = 4294967295
        switch v1 {
        case 42:
            v1 = (v1 - genFn1(-2147483648, 0))
        default:
            v1 -= 79
        }
    }
    switch "a" {
    case "a":
        if (false || ((42 < -91) || (127 <= -1))) {
            if true {
                for fi17 := 0; fi17 < 10; fi17++ {
                    if false {
                        v18 := GenS1{F0: 42, F1: true}
                        for fi19 := 0; fi19 < 5; fi19++ {
                            fmt.Println(fi17)
                            var v20 float32 = -2.25
                            _ = fi19
                            _ = v20
                        }
                        fmt.Println(v1, v1)
                        _ = v18
                    } else {
                        v1 = v1
                        if ((3.14159 <= 1.5) && false) {
                            if true {
                                if false {
                                    if (true || true) {
                                        switch (v1 | fi17) {
                                        case 127:
                                            v21 := GenS5{F0: true, F1: "\n\t", F2: false}
                                            _ = v21
                                        case -2147483648:
                                            for fi22 := 0; fi22 < 8; fi22++ {
                                                v23 := GenS6{F0: true, F1: true, F2: ""}
                                                v2 = append(v2, fi17)
                                                var v24 float32 = ((-2.25 / 1.5) / 0.1)
                                                switch (("" + "") + "\n\t") {
                                                case "foo bar":
                                                    if (float64((v24 - v24)) < float64((float32(0.1) / float32(1.5)))) {
                                                        var v25 uint32 = 2147483648
                                                        var v26 int = -2147483648
                                                        _ = v25
                                                        _ = v26
                                                    } else {
                                                        var v27 uint8 = 0
                                                        _ = v27
                                                    }
                                                }
                                                v28 := GenS1{F0: 65535, F1: false}
                                                v29 := GenS4{F0: -2147483648, F1: -32768, F2: 485}
                                                var v30 int32 = -2147483648
                                                v31 := GenS3{F0: 1}
                                                if (2147483647 >= 2) {
                                                    fmt.Println(fi22, v1, v24)
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    v30 += -514
                                                } else {
                                                    if len(v2) > 0 {
                                                        v2[0] = (fi17 & (v1 + fi17))
                                                    }
                                                    v32 := GenS7{F0: (^(v30)), F1: -128, F2: 189}
                                                    var v33 uint16 = 303
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    _ = v32
                                                    _ = v33
                                                }
                                                switch ((("foo bar" + "") + ((("hello" + "hello") + (("a" + "") + "hello")) + "hello")) + "hello") {
                                                case "\n\t":
                                                    fmt.Println(v24)
                                                case "":
                                                    if ((127 < 2) && ((0.0 / 1.0) == -2.25)) {
                                                        var v34 float32 = float32((((-2.25 / 3.14159) + 1e10) - 3.14159))
                                                        var v35 int = -2147483648
                                                        switch -968 {
                                                        case 392:
                                                            switch genFn1(v30, -1) {
                                                            case -1:
                                                                _ = len(v2)
                                                            }
                                                        case 1:
                                                            if false {
                                                                var v36 bool = (!(-41 == -1))
                                                                var v37 int64 = 149
                                                                var v38 uint = 1
                                                                _ = v36
                                                                _ = v37
                                                                _ = v38
                                                            } else {
                                                                v1 -= genFn1((v30 % ((871) | 1)), 1000)
                                                                v39 := GenS2{F0: 40}
                                                                fmt.Println(v24, v34)
                                                                if len(v2) > 0 {
                                                                    v2[0] = 1
                                                                }
                                                                var v40 uint8 = 0
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                var v41 uint32 = 887
                                                                _ = v39
                                                                _ = v40
                                                                _ = v41
                                                            }
                                                        default:
                                                            v42 := GenS8{F0: 127}
                                                            _ = v42
                                                        }
                                                        _ = v34
                                                        _ = v35
                                                    }
                                                }
                                                v1 += fi22
                                                var v43 float64 = 0.1
                                                _ = fi22
                                                _ = v23
                                                _ = v24
                                                _ = v28
                                                _ = v29
                                                _ = v30
                                                _ = v31
                                                _ = v43
                                            }
                                        }
                                        v44 := GenS8{F0: 255}
                                        _ = v44
                                    } else {
                                        for fi45 := 0; fi45 < 2; fi45++ {
                                            var v46 uint = 0
                                            v2 = append(v2, -230)
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = fi45
                                            _ = v46
                                        }
                                        v1 -= (v1 * -40)
                                    }
                                    if true {
                                        switch "foo bar" {
                                        case "":
                                            v1 -= (v1 / ((fi17) | 1))
                                        case "a":
                                            for fi47 := 0; fi47 < 15; fi47++ {
                                                v1 += (-(fi47))
                                                var v48 int64 = 9223372036854775807
                                                _ = fi47
                                                _ = v48
                                            }
                                        case "hello":
                                            v1 -= v1
                                        default:
                                            switch fi17 {
                                            case 32767:
                                                v49 := GenS7{F0: 853, F1: -128, F2: -1000}
                                                _ = v49
                                            case 1:
                                                v50 := GenS4{F0: -180, F1: 1, F2: 9223372036854775807}
                                                _ = v50
                                            default:
                                                v2 = append(v2, fi17)
                                            }
                                        }
                                        var v51 uint64 = 2
                                        var v52 float64 = 0.1
                                        _ = v51
                                        _ = v52
                                    }
                                }
                            } else {
                                v1 = (v1 ^ genFn1(42, 1000))
                                if true {
                                    _ = len(v2)
                                } else {
                                    switch (v1 * fi17) {
                                    case -984:
                                        var v53 uint16 = 1
                                        _ = v53
                                    }
                                    v54 := GenS9{F0: 1}
                                    v1 = 42
                                    _ = v54
                                }
                                v1 -= (fi17 - v1)
                            }
                            if ((!((-563 >= 766) && (!(65535 == 4294967295)))) && (0.1 < 1.5)) {
                                var v55 float32 = 3.14159
                                v1 += v1
                                _ = v55
                            }
                        }
                    }
                    _ = fi17
                }
                if true {
                    if (1e10 != 1e10) {
                        var v56 uint32 = 787
                        _ = v56
                    } else {
                        fmt.Println(v1)
                    }
                } else {
                    switch (v1 - (^(v1))) {
                    case 127:
                        v57 := GenS8{F0: 140}
                        _ = v57
                    case 128:
                        if true {
                            if true {
                                var v58 uint = 417
                                var v59 float32 = 3.14159
                                fmt.Println(v1)
                                _ = v58
                                _ = v59
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = "hello"
                                }
                            }
                            for fi60 := 0; fi60 < 12; fi60++ {
                                for fi61 := 0; fi61 < 18; fi61++ {
                                    var v62 int16 = 469
                                    if (918 != 1) {
                                        v62 = (v62 % ((-188) | 1))
                                        var v63 uint32 = 490
                                        _ = v63
                                    }
                                    _ = fi61
                                    _ = v62
                                }
                                switch ("\n\t" + "\n\t") {
                                case "a":
                                    switch "\n\t" {
                                    case "a":
                                        v1 = 32767
                                    case "\n\t":
                                        if false {
                                            if (("\n\t" + "") == "") {
                                                var v64 uint32 = 740
                                                _ = v64
                                            } else {
                                                v2 = append(v2, 128)
                                            }
                                            for fi65 := 0; fi65 < 2; fi65++ {
                                                v2 = append(v2, genFn1(-1, 1))
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                _ = fi65
                                            }
                                            var v66 float64 = (0.1 / 1.5)
                                            _ = v66
                                        } else {
                                            v2 = append(v2, v1)
                                        }
                                    default:
                                        for fi67 := 0; fi67 < 4; fi67++ {
                                            if (false && (4294967295 != 65535)) {
                                                if false {
                                                    var v68 int32 = 42
                                                    fmt.Println(fi60, fi60, v1)
                                                    _ = v68
                                                } else {
                                                    v2 = append(v2, (^(fi60)))
                                                }
                                                v1 += (fi60 & (fi67 * -386))
                                                var v69 uint32 = 65535
                                                _ = v69
                                            }
                                            fmt.Println(fi60, v1, fi67)
                                            switch "foo bar" {
                                            case "foo bar":
                                                var v70 int64 = 42
                                                _ = v70
                                            default:
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                            if false {
                                                v1 -= (fi67 - 4294967295)
                                                if ((fi67 | v1) >= (fi67 + v1)) {
                                                    v71 := GenS10{F0: 42}
                                                    var v72 int64 = -761
                                                    if (false && false) {
                                                        v1 -= genFn1(1, 1000)
                                                    } else {
                                                        switch 32768 {
                                                        case 127:
                                                            var v73 uint64 = 2
                                                            _ = v73
                                                        case -2147483648:
                                                            if true {
                                                                v2 = append(v2, -679)
                                                            } else {
                                                                v2 = append(v2, (^(fi67)))
                                                                var v74 float32 = (0.0 - -2.25)
                                                                v74 = 0.0
                                                                _ = v74
                                                            }
                                                        default:
                                                            var v75 uint8 = 0
                                                            _ = v75
                                                        }
                                                        var v76 uint64 = 4294967295
                                                        _ = v76
                                                    }
                                                    _ = v71
                                                    _ = v72
                                                }
                                                fmt.Println(fi67, fi60, v1)
                                                v1 += (v1 / (((v1 * genFn1(-731, 1))) | 1))
                                                if false {
                                                    if (!(-513 < -9223372036854775808)) {
                                                        var v77 string = (("a" + (("foo bar" + (("" + ("a" + "hello")) + "foo bar")) + "a")) + "\n\t")
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            _ = len(v3)
                                                            var v78 string = v77
                                                            _ = v78
                                                        default:
                                                            var v79 int = (fi60 >> ((799) & 63))
                                                            _ = v79
                                                        }
                                                        _ = v77
                                                    } else {
                                                        var v80 int64 = 9223372036854775807
                                                        v81 := GenS4{F0: -1, F1: -1, F2: (v80 * (v80 | -230))}
                                                        v82 := GenS11{F0: false, F1: 136, F2: 375}
                                                        _ = v80
                                                        _ = v81
                                                        _ = v82
                                                    }
                                                }
                                                v83 := GenS12{F0: (860 != 2), F1: (-(v1)), F2: "a"}
                                                var v84 int32 = 0
                                                _ = v83
                                                _ = v84
                                            } else {
                                                var v85 uint32 = 2147483647
                                                switch "" {
                                                case "hello":
                                                    fmt.Println(v1, v85, fi67)
                                                    switch (fi67 << ((116) & 63)) {
                                                    case -20:
                                                        v85 -= v85
                                                    }
                                                    v2 = append(v2, v1)
                                                    var v86 uint32 = 2147483647
                                                    _ = v86
                                                default:
                                                    if false {
                                                        fmt.Println(v85, fi60)
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                }
                                                v85 = 42
                                                if (1 <= fi60) {
                                                    v87 := GenS1{F0: 796, F1: (v85 < v85)}
                                                    _ = v87
                                                } else {
                                                    _ = len(v2)
                                                }
                                                _ = v85
                                            }
                                            _ = fi67
                                        }
                                    }
                                case "\n\t":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                default:
                                    v1 -= (fi60 * (v1 | genFn1(int32(fi60), 1000)))
                                }
                                for fi88 := 0; fi88 < 18; fi88++ {
                                    if (genFn0(998, 1, 42) != 172) {
                                        if ("foo bar" != ("a" + "hello")) {
                                            if (186 > 415) {
                                                _ = len(v3)
                                            } else {
                                                v1 -= fi60
                                                var v89 int32 = -520
                                                _ = v89
                                            }
                                        }
                                        fmt.Println(fi88, fi88, fi88)
                                        fmt.Println(v1, fi88, fi60)
                                    }
                                    v1 = -809
                                    v3 = append(v3, "\n\t")
                                    if ((^(v1)) >= v1) {
                                        if (true || (127 <= -367)) {
                                            if (749 < 4294967295) {
                                                v1 = 256
                                                var v90 int8 = -1
                                                _ = v90
                                            } else {
                                                if true {
                                                    var v91 int64 = 866
                                                    _ = v91
                                                } else {
                                                    fmt.Println(v1)
                                                    var v92 int = (fi88 & -9223372036854775808)
                                                    _ = v92
                                                }
                                                v1 = fi60
                                                v1 = (v1 % ((genFn1(0, -773)) | 1))
                                            }
                                            v93 := GenS9{F0: 2}
                                            _ = v93
                                        }
                                        v1 = (fi88 << ((4294967295) & 63))
                                    } else {
                                        var v94 uint = 351
                                        switch genFn1(1, -398) {
                                        case -9223372036854775808:
                                            switch ("\n\t" + "hello") {
                                            case "":
                                                if (!(!false)) {
                                                    v94 = v94
                                                    var v95 uint = v94
                                                    v94 = 628
                                                    _ = v95
                                                } else {
                                                    var v96 int64 = 562
                                                    switch ("foo bar" + "foo bar") {
                                                    case "":
                                                        var v97 string = "a"
                                                        _ = v97
                                                    case "\n\t":
                                                        var v98 int32 = -324
                                                        _ = v98
                                                    case "hello":
                                                        var v99 float32 = (0.0 + -2.25)
                                                        _ = v99
                                                    }
                                                    if (!false) {
                                                        var v100 float64 = 0.0
                                                        var v101 int64 = v96
                                                        v102 := GenS13{F0: 2}
                                                        _ = v100
                                                        _ = v101
                                                        _ = v102
                                                    }
                                                    _ = v96
                                                }
                                            case "foo bar":
                                                var v103 int8 = 1
                                                var v104 int = (fi60 % (((v1 & 32767)) | 1))
                                                _ = v103
                                                _ = v104
                                            case "hello":
                                                v1 -= (fi60 / ((fi60) | 1))
                                                var v105 int32 = 42
                                                _ = v105
                                            }
                                        default:
                                            var v106 int16 = 266
                                            _ = v106
                                        }
                                        var v107 int32 = 1
                                        _ = v107
                                        _ = v94
                                    }
                                    fmt.Println(fi88, v1)
                                    v108 := GenS13{F0: 128}
                                    fmt.Println(v1, fi88)
                                    fmt.Println(fi88, v1, fi60)
                                    if (127 < 99) {
                                        v1 += fi88
                                        switch 128 {
                                        case 250:
                                            if (-481 != -2147483648) {
                                                if len(v3) > 0 {
                                                    v3[0] = "a"
                                                }
                                            }
                                        case 32768:
                                            switch "\n\t" {
                                            case "\n\t":
                                                var v109 uint8 = 128
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                var v110 int64 = -1
                                                _ = v109
                                                _ = v110
                                            default:
                                                var v111 float32 = 1e10
                                                _ = v111
                                            }
                                        case 127:
                                            v112 := GenS3{F0: 42}
                                            _ = v112
                                        }
                                    }
                                    _ = fi88
                                    _ = v108
                                }
                                _ = fi60
                            }
                            var v113 int8 = 26
                            _ = v113
                        } else {
                            var v114 int64 = 1
                            _ = v114
                        }
                    default:
                        var v115 uint = 4294967295
                        _ = v115
                    }
                    var v116 uint16 = 255
                    _ = v116
                }
            } else {
                v1 -= genFn1(1, -32768)
                v117 := GenS4{F0: 1, F1: 0, F2: 42}
                v1 = genFn1(-286, -1)
                _ = v117
            }
            var v118 int64 = -1
            _ = v118
        }
    case "foo bar":
        v1 -= genFn1(28, 0)
    case "hello":
        if (!false) {
            var v119 uint16 = 0
            v119 += v119
            v120 := GenS1{F0: v119, F1: (0.1 != (float64(((float32(float64(0.0)) + (3.14159 - (1.5 + 0.0))) / 1.5)) * (0.1 + -2.25)))}
            _ = v119
            _ = v120
        } else {
            switch -961 {
            case 128:
                fmt.Println(v1)
            case -9223372036854775808:
                var v121 bool = true
                _ = v121
            case -2147483648:
                switch "foo bar" {
                case "\n\t":
                    var v122 float32 = 3.14159
                    _ = v122
                case "foo bar":
                    v1 = (v1 * (v1 - v1))
                default:
                    for fi123 := 0; fi123 < 13; fi123++ {
                        v124 := GenS14{F0: 448, F1: 2147483647, F2: 9223372036854775807}
                        var v125 float32 = (1e10 - 1e10)
                        var v126 uint32 = 2147483647
                        _ = fi123
                        _ = v124
                        _ = v125
                        _ = v126
                    }
                }
            default:
                fmt.Println(v1)
            }
            var v127 string = "a"
            _ = v127
        }
    }
    if ((!(755 != -1)) && true) {
        for fi128 := 0; fi128 < 19; fi128++ {
            v2 = append(v2, v1)
            var v129 uint8 = 255
            _ = fi128
            _ = v129
        }
        fmt.Println(v1)
    } else {
        for fi130 := 0; fi130 < 3; fi130++ {
            for fi131 := 0; fi131 < 1; fi131++ {
                v132 := GenS10{F0: 657}
                fmt.Println(fi130, fi131)
                _ = fi131
                _ = v132
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v133 := GenS10{F0: 1}
            _ = fi130
            _ = v133
        }
    }
    var v134 int64 = 317
    switch ("foo bar" + ("\n\t" + ("hello" + "\n\t"))) {
    case "hello":
        var v135 int32 = 0
        _ = v135
    default:
        var v136 float64 = 1e10
        _ = v136
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v134:", v134)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v134
    _ = v2
    _ = v3
}
