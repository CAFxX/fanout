package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64; F1 uint32; F2 uint16 }
type GenS6 struct { F0 uint32; F1 bool }
type GenS7 struct { F0 uint16; F1 string; F2 string }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint64; F1 string; F2 int }
type GenS10 struct { F0 uint64; F1 int; F2 int64 }
type GenS11 struct { F0 string; F1 uint }
type GenS12 struct { F0 uint8; F1 uint32 }
type GenS13 struct { F0 string }
type GenS14 struct { F0 bool; F1 uint64 }
type GenS15 struct { F0 uint16 }
func genFn0(p0 bool, p1 uint64) uint32 {
    if true {
        for fi5 := 0; fi5 < 19; fi5++ {
            v6 := GenS1{F0: 477, F1: 42, F2: p1}
            p0 = (int(p1) != fi5)
            var v7 uint64 = 42
            _ = fi5
            _ = v6
            _ = v7
        }
        if false {
            for fi8 := 0; fi8 < 18; fi8++ {
                v9 := []int{fi8, 4294967295, fi8, -505}
                if (!(true || (793 < 770))) {
                    if (!(p0 && ((fi8 ^ 4294967295) != 1))) {
                        if p0 {
                            if false {
                                switch "hello" {
                                case "foo bar":
                                    switch "\n\t" {
                                    case "":
                                        var v10 int8 = int8(fi8)
                                        _ = v10
                                    default:
                                        if p0 {
                                            for fi11 := 0; fi11 < 2; fi11++ {
                                                v12 := GenS1{F0: 210, F1: 4294967295, F2: p1}
                                                _ = fi11
                                                _ = v12
                                            }
                                            _ = len(v9)
                                            switch (("\n\t" + "hello") + "\n\t") {
                                            case "foo bar":
                                                fmt.Println(p1, p1)
                                            default:
                                                fmt.Println(p0, fi8)
                                            }
                                        } else {
                                            v9 = append(v9, fi8)
                                            var v13 int64 = 9223372036854775807
                                            _ = v13
                                        }
                                    }
                                case "\n\t":
                                    for fi14 := 0; fi14 < 7; fi14++ {
                                        v15 := GenS1{F0: 631, F1: 0, F2: (p1 | 0)}
                                        _ = fi14
                                        _ = v15
                                    }
                                case "hello":
                                    v16 := GenS2{F0: (p1 << ((0) & 63))}
                                    _ = v16
                                default:
                                    var v17 int16 = -512
                                    _ = v17
                                }
                                if len(v9) > 0 {
                                    v9[0] = (fi8 / ((1) | 1))
                                }
                                v18 := GenS3{F0: (!(4294967295 != 4294967295))}
                                _ = v18
                            } else {
                                if len(v9) > 0 {
                                    v9[0] = fi8
                                }
                            }
                            var v19 bool = false
                            _ = v19
                        } else {
                            var v20 int8 = 100
                            if (!(!(3.14159 < -2.25))) {
                                var v21 int = (fi8 | 128)
                                var v22 string = ("\n\t" + "a")
                                _ = v21
                                _ = v22
                            } else {
                                if true {
                                    if len(v9) > 0 {
                                        v9[0] = fi8
                                    }
                                    for fi23 := 0; fi23 < 8; fi23++ {
                                        v24 := GenS1{F0: 78, F1: 65535, F2: 1}
                                        if p0 {
                                            v25 := GenS2{F0: 662}
                                            switch "" {
                                            case "":
                                                if len(v9) > 0 {
                                                    v9[0] = (fi8 | (-(fi23)))
                                                }
                                            }
                                            p1 += (p1 | p1)
                                            v26 := GenS3{F0: p0}
                                            v9 = append(v9, fi23)
                                            if ((float64(0.1) + 1e10) > 1e10) {
                                                if p0 {
                                                    if (113 >= 2) {
                                                        fmt.Println(fi23)
                                                        if (1.5 < float32(0.1)) {
                                                            var v27 float64 = 0.1
                                                            var v28 uint8 = 194
                                                            v29 := GenS3{F0: true}
                                                            _ = v27
                                                            _ = v28
                                                            _ = v29
                                                        } else {
                                                            fmt.Println(fi8, v20, v20)
                                                            p0 = (-1 == 1)
                                                        }
                                                        v20 += (v20 >> ((951) & 63))
                                                    } else {
                                                        switch 32768 {
                                                        case 169:
                                                            _ = len(v9)
                                                        default:
                                                            var v30 int32 = 0
                                                            _ = v30
                                                        }
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            v20 -= 1
                                                            p0 = ((fi23 >> ((2) & 63)) < 2147483647)
                                                            v31 := GenS4{F0: ""}
                                                            _ = v31
                                                        default:
                                                            v32 := GenS2{F0: p1}
                                                            _ = v32
                                                        }
                                                    }
                                                    if p0 {
                                                        if true {
                                                            var v33 int = (fi23 & fi8)
                                                            if false {
                                                                var v34 uint16 = 65535
                                                                if len(v9) > 0 {
                                                                    v9[0] = (^(fi8))
                                                                }
                                                                if false {
                                                                    p0 = p0
                                                                    v35 := GenS1{F0: 42, F1: uint32(p1), F2: (p1 % (((p1 >> ((255) & 63))) | 1))}
                                                                    _ = v35
                                                                }
                                                                _ = v34
                                                            }
                                                            var v36 uint8 = 47
                                                            _ = v33
                                                            _ = v36
                                                        } else {
                                                            v20 += -92
                                                            switch fi23 {
                                                            case 0:
                                                                var v37 uint16 = 255
                                                                _ = v37
                                                            }
                                                            if (!(126 < 0)) {
                                                                if len(v9) > 0 {
                                                                    v9[0] = fi8
                                                                }
                                                                if len(v9) > 0 {
                                                                    v9[0] = 695
                                                                }
                                                                if p0 {
                                                                    var v38 string = ("foo bar" + ("a" + "\n\t"))
                                                                    fmt.Println(fi8, p1)
                                                                    _ = len(v9)
                                                                    _ = v38
                                                                }
                                                            } else {
                                                                switch "\n\t" {
                                                                case "a":
                                                                    if p0 {
                                                                        var v39 uint = 311
                                                                        p1 -= 4294967295
                                                                        v40 := GenS5{F0: 9223372036854775807, F1: 2, F2: 1}
                                                                        _ = v39
                                                                        _ = v40
                                                                    } else {
                                                                        v20 = 109
                                                                        p0 = p0
                                                                    }
                                                                }
                                                                v20 -= 0
                                                            }
                                                        }
                                                        p1 += 42
                                                        if p0 {
                                                            _ = len(v9)
                                                            var v41 string = "\n\t"
                                                            var v42 float64 = -2.25
                                                            _ = v41
                                                            _ = v42
                                                        }
                                                        var v43 int8 = v20
                                                        switch "" {
                                                        case "hello":
                                                            if ((0 >= 359) && false) {
                                                                fmt.Println(v43)
                                                                var v44 int8 = (v43 + (v20 ^ (v43 % ((v20) | 1))))
                                                                _ = v44
                                                            }
                                                        }
                                                        _ = v43
                                                    }
                                                }
                                                var v45 uint32 = 2
                                                _ = v45
                                            } else {
                                                v46 := GenS4{F0: "\n\t"}
                                                _ = v46
                                            }
                                            _ = v25
                                            _ = v26
                                        }
                                        var v47 uint64 = p1
                                        _ = fi23
                                        _ = v24
                                        _ = v47
                                    }
                                } else {
                                    for fi48 := 0; fi48 < 15; fi48++ {
                                        p1 = 717
                                        fmt.Println(fi8, fi8, p1)
                                        _ = fi48
                                    }
                                }
                            }
                            _ = v20
                        }
                        switch (fi8 % (((^(fi8))) | 1)) {
                        case -236:
                            for fi49 := 0; fi49 < 14; fi49++ {
                                v50 := GenS6{F0: 2, F1: p0}
                                p0 = (1 != -32768)
                                _ = fi49
                                _ = v50
                            }
                        default:
                            if true {
                                fmt.Println(fi8, p1)
                                if len(v9) >= 2 {
                                    v9 = v9[1:len(v9)]
                                }
                                p0 = true
                            }
                        }
                    } else {
                        v51 := GenS7{F0: 2, F1: "foo bar", F2: ""}
                        _ = v51
                    }
                }
                _ = fi8
                _ = v9
            }
            fmt.Println(p0)
        } else {
            v52 := GenS8{F0: 0}
            fmt.Println(p0)
            _ = v52
        }
        v53 := []int{}
        _ = v53
    }
    v54 := []int{32767, 127, 4294967295, 127}
    v55 := GenS9{F0: 1, F1: "hello", F2: 161}
    var v56 float32 = 1e10
    _ = p0
    _ = p1
    _ = v54
    _ = v55
    _ = v56
    return 4294967295
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 0
    var v3 float32 = ((float32(0.1) / -2.25) + 0.0)
    v4 := make(map[string]string)
    v57 := []int{}
    v2 = v2
    v58 := GenS2{F0: 18446744073709551615}
    v3 = v3
    v1[""] = (!true)
    fmt.Println(v2, v2)
    for fi59 := 0; fi59 < 13; fi59++ {
        for fi60 := 0; fi60 < 13; fi60++ {
            _ = len(v57)
            v61 := GenS10{F0: 1, F1: (fi60 << ((1) & 63)), F2: -1}
            if len(v57) >= 2 {
                v57 = v57[1:len(v57)]
            }
            if true {
                _ = len(v57)
                v62 := GenS5{F0: 0, F1: 0, F2: 32767}
                if len(v57) > 0 {
                    v57[0] = 127
                }
                _ = v62
            }
            var v63 string = (("hello" + "\n\t") + "")
            _ = fi60
            _ = v61
            _ = v63
        }
        _ = fi59
    }
    if len(v57) > 0 {
        v57[0] = 42
    }
    if (127 == 1) {
        switch -642 {
        case 4294967295:
            if ("" > "") {
                v64 := v1[""]
                v65 := GenS9{F0: 42, F1: "", F2: 2147483647}
                _ = v64
                _ = v65
            }
        }
        v66 := GenS1{F0: 2147483648, F1: 4294967295, F2: 436}
        var v67 float64 = (1e10 * 3.14159)
        _ = v66
        _ = v67
    } else {
        fmt.Println(v2, v2)
        for fi68 := 0; fi68 < 8; fi68++ {
            v69 := GenS10{F0: 1, F1: (fi68 & 32768), F2: 1}
            switch "" {
            case "":
                for fi70 := 0; fi70 < 19; fi70++ {
                    fmt.Println(v3)
                    var v71 uint32 = 466
                    _ = fi70
                    _ = v71
                }
            }
            switch "\n\t" {
            case "\n\t":
                v72 := GenS4{F0: "foo bar"}
                _ = v72
            case "a":
                if (true && true) {
                    v73 := GenS6{F0: 65535, F1: (0 == -464)}
                    _ = v73
                } else {
                    switch fi68 {
                    case 128:
                        v57 = append(v57, 32768)
                    }
                }
            case "foo bar":
                fmt.Println(v2)
            }
            _ = fi68
            _ = v69
        }
    }
    for fi74 := 0; fi74 < 8; fi74++ {
        for fi75 := 0; fi75 < 7; fi75++ {
            v76 := GenS3{F0: (true && true)}
            fmt.Println(v3, v3)
            _ = fi75
            _ = v76
        }
        if false {
            var v77 float32 = 0.0
            _ = v77
        } else {
            for fi78 := 0; fi78 < 4; fi78++ {
                var v79 int64 = 126
                _ = fi78
                _ = v79
            }
        }
        _ = fi74
    }
    if (v3 >= (float32((0.0 - (float64(float32(1.5)) - float64(-2.25)))) + 0.0)) {
        for fi80 := 0; fi80 < 5; fi80++ {
            for fi81 := 0; fi81 < 4; fi81++ {
                if true {
                    if len(v57) >= 2 {
                        v57 = v57[1:len(v57)]
                    }
                    var v82 uint16 = 1
                    _ = v82
                }
                _ = fi81
            }
            var v83 uint16 = 2
            fmt.Println(v83, v83, v83)
            _ = fi80
            _ = v83
        }
        if true {
            for fi84 := 0; fi84 < 5; fi84++ {
                fmt.Println(v3, v3)
                v3 += v3
                _ = fi84
            }
        } else {
            v2 -= 127
        }
        delete(v4, "foo bar")
    }
    if true {
        v85 := GenS11{F0: ("\n\t" + "foo bar"), F1: 855}
        var v86 uint64 = 42
        var v87 int = 9223372036854775807
        _ = v85
        _ = v86
        _ = v87
    } else {
        for fi88 := 0; fi88 < 2; fi88++ {
            v2 += (v2 + 62)
            v2 -= v2
            for fi89 := 0; fi89 < 5; fi89++ {
                v90 := GenS2{F0: 42}
                var v91 int8 = 100
                var v92 uint32 = genFn0(false, 461)
                var v93 uint16 = 255
                if false {
                    var v94 int32 = -2147483648
                    _ = len(v57)
                    _ = v94
                }
                v92 -= genFn0((true || true), 4294967295)
                switch ("\n\t" + (("" + "a") + ("foo bar" + ""))) {
                case "a":
                    v95 := GenS5{F0: 42, F1: v92, F2: v93}
                    _ = v95
                default:
                    if ((float32(((0.0 / 1e10) - float64(v3))) / v3) != (v3 / v3)) {
                        var v96 int = fi89
                        if false {
                            var v97 int8 = (-(v91))
                            _ = v97
                        } else {
                            switch v96 {
                            case 127:
                                delete(v4, "hello")
                            case 32767:
                                v91 = 127
                            }
                            var v98 float32 = (v3 * v3)
                            var v99 int64 = -822
                            var v100 bool = (22 >= v2)
                            fmt.Println(v92, v98)
                            _ = v100
                            _ = v98
                            _ = v99
                        }
                        switch "foo bar" {
                        case "\n\t":
                            if len(v57) >= 2 {
                                v57 = v57[1:len(v57)]
                            }
                            v93 += 2
                        case "":
                            if true {
                                v101 := GenS6{F0: 2, F1: (v92 <= (v92 / ((v92) | 1)))}
                                _ = v101
                            }
                        }
                        v102 := v1[(("a" + "") + ("\n\t" + ("\n\t" + ("\n\t" + "a"))))]
                        var v103 float32 = v3
                        _ = v102
                        _ = v103
                        _ = v96
                    }
                }
                _ = fi89
                _ = v90
                _ = v91
                _ = v92
                _ = v93
            }
            _ = fi88
        }
        v104 := GenS12{F0: v2, F1: 42}
        for fi105 := 0; fi105 < 16; fi105++ {
            if (0 <= 65535) {
                for fi106 := 0; fi106 < 12; fi106++ {
                    switch ("\n\t" + "\n\t") {
                    case "":
                        fmt.Println(fi106)
                    case "foo bar":
                        if true {
                            if false {
                                switch ("\n\t" + "hello") {
                                case "hello":
                                    if len(v57) > 0 {
                                        v57[0] = (-(fi105))
                                    }
                                    v3 -= v3
                                default:
                                    fmt.Println(fi105, v3)
                                }
                                if len(v57) > 0 {
                                    v57[0] = (-(fi105))
                                }
                            } else {
                                v3 += (v3 - 0.0)
                            }
                            v3 += (1e10 * v3)
                        } else {
                            v107 := GenS12{F0: v2, F1: 496}
                            v108 := GenS13{F0: ""}
                            switch fi105 {
                            case 9223372036854775807:
                                v109 := GenS14{F0: (!(-141 > 1000)), F1: 325}
                                var v110 int8 = 0
                                fmt.Println(v3, v2)
                                _ = v109
                                _ = v110
                            case 368:
                                fmt.Println(v2)
                            }
                            _ = v107
                            _ = v108
                        }
                    default:
                        if true {
                            fmt.Println(v2, fi106)
                            switch "foo bar" {
                            case "foo bar":
                                if len(v57) > 0 {
                                    v57[0] = fi105
                                }
                            case "":
                                v111 := GenS11{F0: "", F1: 0}
                                _ = v111
                            case "\n\t":
                                var v112 int = fi105
                                _ = v112
                            default:
                                var v113 float32 = float32(1.5)
                                _ = v113
                            }
                        }
                    }
                    switch "hello" {
                    case "":
                        if (0 == 2) {
                            var v114 bool = (0 < (fi106 >> ((712) & 63)))
                            switch fi106 {
                            case 4294967295:
                                var v115 int32 = 2147483647
                                _ = v115
                            case 832:
                                var v116 uint16 = 2
                                _ = v116
                            case 128:
                                v2 = (v2 ^ (^(v2)))
                            }
                            var v117 int64 = 9223372036854775807
                            v3 = v3
                            _ = v114
                            _ = v117
                        }
                    default:
                        v3 = 1.5
                    }
                    _ = fi106
                }
            }
            if (!(!false)) {
                if len(v57) > 0 {
                    v57[0] = -325
                }
                var v118 uint32 = 597
                _ = v118
            } else {
                if (1 == -1) {
                    v3 = (v3 / v3)
                    var v119 int8 = 100
                    _ = v119
                } else {
                    v2 -= v2
                }
                for fi120 := 0; fi120 < 18; fi120++ {
                    v2 -= v2
                    _ = fi120
                }
            }
            _ = fi105
        }
        _ = v104
    }
    var v121 uint = 42
    var v122 uint32 = 65535
    var v123 bool = false
    var v124 int = 936
    for fi125 := 0; fi125 < 10; fi125++ {
        for fi126 := 0; fi126 < 17; fi126++ {
            switch "\n\t" {
            case "a":
                var v127 int32 = -2147483648
                _ = v127
            case "":
                if ((!true) || ((v123 && ((v121 ^ 4294967295) != 2)) && v123)) {
                    var v128 uint16 = 42
                    v129 := GenS15{F0: 484}
                    _ = v128
                    _ = v129
                }
            case "foo bar":
                var v130 int16 = 385
                _ = v130
            }
            v121 -= v121
            _ = fi126
        }
        fmt.Println(fi125, v121, v3)
        _ = fi125
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v121:", v121)
    fmt.Println("v122:", v122)
    fmt.Println("v123:", v123)
    fmt.Println("v124:", v124)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v57:", len(v57), v57)
    _ = v1
    _ = v121
    _ = v122
    _ = v123
    _ = v124
    _ = v2
    _ = v3
    _ = v4
    _ = v57
    _ = v58
}
