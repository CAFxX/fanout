package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint; F2 int }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 int32; F1 uint64; F2 int32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 string; F1 int32 }
type GenS8 struct { F0 int32; F1 uint; F2 uint }
type GenS9 struct { F0 uint16; F1 uint8; F2 string }
type GenS10 struct { F0 int16; F1 uint }
type GenS11 struct { F0 int; F1 uint64; F2 bool }
type GenS12 struct { F0 uint64; F1 uint32; F2 int }
type GenS13 struct { F0 int8; F1 uint64; F2 int32 }
type GenS14 struct { F0 uint16; F1 uint16; F2 int }
type GenS15 struct { F0 int16; F1 int16 }
type GenS16 struct { F0 string; F1 uint; F2 uint8 }
func genFn0(p0 bool, p1 bool) uint8 {
    v5 := GenS1{F0: 4294967295, F1: 735}
    if (1.5 < 1e10) {
        if (0.1 > 1.5) {
            var v6 uint64 = 0
            var v7 int32 = -1
            _ = v6
            _ = v7
        } else {
            var v8 uint32 = 4294967295
            p0 = p0
            _ = v8
        }
    } else {
        v9 := []int{177, 32768, 32767, 127}
        if p1 {
            v10 := GenS1{F0: 1, F1: 0}
            _ = v10
        } else {
            fmt.Println(p1)
            fmt.Println(p1, p0)
            p1 = p1
        }
        _ = v9
    }
    v11 := GenS1{F0: 491, F1: 32767}
    p1 = p1
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v11
    _ = v5
    return 68
}
func genFn1() uint16 {
    if ((100 <= -57) || (42 < 797)) {
        switch "\n\t" {
        case "hello":
            var v12 uint16 = 42
            _ = v12
        }
        var v13 int = 1
        fmt.Println(v13)
        _ = v13
    } else {
        v14 := []int{128, -1, -479, -889}
        if len(v14) > 0 {
            v14[0] = 9223372036854775807
        }
        if (!(1.5 > float32((0.1 / ((1.5 - (((1.5 * ((1e10 + float64(1e10)) - 1e10)) - 1e10) * 1.5)) * 1.5))))) {
            v15 := GenS1{F0: 42, F1: 1000}
            if true {
                var v16 uint = 154
                _ = v16
            }
            _ = v15
        }
        _ = v14
    }
    var v17 int32 = 1
    var v18 int8 = 1
    _ = v17
    _ = v18
    return 0
}

func main() {
    v1 := []uint32{2147483648, 4294967295, 2147483647}
    v2 := make(map[int]bool)
    v3 := []int{145}
    v4 := make(map[uint32]bool)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch -2147483648 {
    case 128:
        v19 := GenS1{F0: 1, F1: -575}
        _ = v19
    case 32767:
        if false {
            var v20 int64 = 54
            _ = v20
        }
    case 32768:
        for fi21 := 0; fi21 < 15; fi21++ {
            switch fi21 {
            case -431:
                var v22 int = 32768
                _ = v22
            case -2147483648:
                for fi23 := 0; fi23 < 18; fi23++ {
                    var v24 uint16 = 1
                    _ = fi23
                    _ = v24
                }
            case 32767:
                for fi25 := 0; fi25 < 14; fi25++ {
                    if (!false) {
                        var v26 int8 = -128
                        _ = v26
                    } else {
                        var v27 int32 = 42
                        _ = v27
                    }
                    var v28 int = fi25
                    if false {
                        v29 := GenS2{F0: 42, F1: 4294967295, F2: (fi21 << ((0) & 63))}
                        switch "foo bar" {
                        case "foo bar":
                            var v30 uint32 = 784
                            _ = v30
                        }
                        var v31 float32 = 0.1
                        _ = v29
                        _ = v31
                    } else {
                        var v32 uint8 = 1
                        _ = v32
                    }
                    _ = len(v4)
                    var v33 uint = 415
                    _ = fi25
                    _ = v28
                    _ = v33
                }
            default:
                if (-612 >= 32767) {
                    for fi34 := 0; fi34 < 20; fi34++ {
                        var v35 uint = 255
                        if (65535 < genFn1()) {
                            if true {
                                switch fi34 {
                                case 974:
                                    switch fi34 {
                                    case 127:
                                        v36 := GenS2{F0: 2, F1: v35, F2: (fi34 | fi21)}
                                        _ = v36
                                    case -569:
                                        var v37 int = (fi21 >> ((v35) & 63))
                                        _ = v37
                                    }
                                default:
                                    v35 = 675
                                }
                                if (0 >= genFn1()) {
                                    var v38 int16 = -150
                                    _ = len(v1)
                                    if true {
                                        var v39 int16 = (v38 | 32767)
                                        var v40 float64 = 0.1
                                        _ = v39
                                        _ = v40
                                    } else {
                                        v38 += v38
                                        fmt.Println(v38)
                                        switch "hello" {
                                        case "a":
                                            v38 = v38
                                        case "":
                                            if false {
                                                fmt.Println(fi34, fi34)
                                            } else {
                                                v41 := GenS3{F0: 843, F1: v38}
                                                fmt.Println(v38, v35, v35)
                                                _ = v41
                                            }
                                        case "hello":
                                            v3 = append(v3, (-(fi21)))
                                        }
                                    }
                                    v38 = v38
                                    _ = v38
                                }
                                if true {
                                    v42 := GenS2{F0: 0, F1: v35, F2: (fi34 * (fi21 - -2147483648))}
                                    switch fi34 {
                                    case 127:
                                        fmt.Println(fi21, fi21)
                                    case 4294967295:
                                        var v43 float64 = ((1e10 + 3.14159) / (-2.25 - (3.14159 * (1.5 * (-2.25 / 1.0)))))
                                        _ = v43
                                    case 42:
                                        var v44 int16 = -451
                                        _ = v44
                                    }
                                    var v45 uint64 = 1
                                    if true {
                                        var v46 int64 = 0
                                        v2[42] = (!true)
                                        v35 -= 4294967295
                                        _ = v46
                                    }
                                    _ = v42
                                    _ = v45
                                }
                                if false {
                                    v47 := GenS1{F0: 0, F1: 0}
                                    var v48 uint64 = 18446744073709551615
                                    _ = v47
                                    _ = v48
                                } else {
                                    fmt.Println(v35)
                                    v49 := GenS4{F0: -1, F1: 110}
                                    _ = v49
                                }
                                if (false && (!(1 >= 4294967295))) {
                                    fmt.Println(v35, fi21)
                                    switch fi21 {
                                    case 32768:
                                        v35 = 865
                                    case 32767:
                                        v35 -= v35
                                    case 532:
                                        var v50 uint = (v35 - v35)
                                        _ = v50
                                    }
                                    var v51 uint32 = 1
                                    _ = v51
                                } else {
                                    _ = len(v4)
                                    switch "foo bar" {
                                    case "":
                                        v35 = 42
                                    case "foo bar":
                                        v35 = v35
                                    }
                                    var v52 float64 = 0.0
                                    _ = len(v3)
                                    if false {
                                        v52 = -2.25
                                        _ = len(v2)
                                    }
                                    _ = v52
                                }
                            }
                            var v53 bool = (4294967295 == 42)
                            if (-2147483648 >= 2147483647) {
                                var v54 int32 = 42
                                _ = v54
                            }
                            _ = v53
                        } else {
                            if false {
                                var v55 int64 = -398
                                v4[861] = (false || false)
                                v35 = v35
                                _ = v55
                            } else {
                                if true {
                                    switch 9223372036854775807 {
                                    case 1:
                                        switch "foo bar" {
                                        case "":
                                            v56 := v4[936]
                                            _ = v56
                                        default:
                                            var v57 bool = (("a" + "hello") <= ("" + ""))
                                            _ = v57
                                        }
                                    }
                                    var v58 uint8 = 127
                                    _ = v58
                                }
                                var v59 bool = true
                                if v59 {
                                    v35 = v35
                                    switch fi21 {
                                    case 4294967295:
                                        switch "a" {
                                        case "hello":
                                            fmt.Println(v35, fi34)
                                        case "foo bar":
                                            v60 := GenS5{F0: -2147483648, F1: 42, F2: 1}
                                            _ = v60
                                        case "a":
                                            if (9 != 127) {
                                                switch fi21 {
                                                case 1:
                                                    switch ("" + ("\n\t" + "\n\t")) {
                                                    case "hello":
                                                        v61 := GenS2{F0: 657, F1: 2, F2: fi21}
                                                        _ = v61
                                                    case "\n\t":
                                                        if v59 {
                                                            if len(v1) > 0 {
                                                                v1[0] = 2147483648
                                                            }
                                                        }
                                                    case "a":
                                                        v35 = 4294967295
                                                    }
                                                case 162:
                                                    v2[fi34] = (v59 && (v59 && v59))
                                                default:
                                                    var v62 uint64 = 970
                                                    _ = v62
                                                }
                                                v35 = v35
                                                v59 = v59
                                                v63 := GenS2{F0: 1, F1: v35, F2: (-(fi21))}
                                                _ = v63
                                            }
                                            var v64 int = 4294967295
                                            fmt.Println(v64, fi34, fi21)
                                            _ = v64
                                        default:
                                            var v65 bool = v59
                                            _ = v65
                                        }
                                    case 0:
                                        v59 = (155 != uint32(v35))
                                    default:
                                        v1 = append(v1, 0)
                                    }
                                    fmt.Println(fi34, v59, v59)
                                }
                                _ = v59
                            }
                        }
                        _ = fi34
                        _ = v35
                    }
                    var v66 uint16 = 487
                    var v67 uint = 4294967295
                    _ = v66
                    _ = v67
                } else {
                    if true {
                        fmt.Println(fi21)
                    } else {
                        switch "foo bar" {
                        case "foo bar":
                            switch fi21 {
                            case -318:
                                v68 := GenS4{F0: -1, F1: 1}
                                _ = v68
                            case 353:
                                v69 := GenS6{F0: -1}
                                _ = v69
                            default:
                                var v70 int64 = -87
                                _ = v70
                            }
                        case "hello":
                            for fi71 := 0; fi71 < 1; fi71++ {
                                if false {
                                    var v72 uint = 255
                                    switch 32767 {
                                    case 4294967295:
                                        v72 = 1
                                    case -1:
                                        v72 -= 2
                                    }
                                    _ = v72
                                } else {
                                    if (702 <= 4294967295) {
                                        var v73 uint8 = 255
                                        v74 := GenS7{F0: "foo bar", F1: -186}
                                        _ = v73
                                        _ = v74
                                    } else {
                                        if false {
                                            v75 := GenS7{F0: (("foo bar" + ((("a" + (("hello" + ("\n\t" + ("" + ""))) + ("hello" + "a"))) + (("" + "foo bar") + "hello")) + "a")) + "foo bar"), F1: 224}
                                            var v76 float64 = 1e10
                                            switch "hello" {
                                            case "foo bar":
                                                fmt.Println(fi71)
                                            case "hello":
                                                var v77 uint8 = genFn0(false, true)
                                                _ = v77
                                            case "":
                                                fmt.Println(v76, fi71)
                                            }
                                            _ = v75
                                            _ = v76
                                        } else {
                                            var v78 int32 = 2147483647
                                            _ = v78
                                        }
                                        v79 := v2[(fi21 * fi71)]
                                        var v80 uint16 = 42
                                        v80 -= 255
                                        v1 = append(v1, 635)
                                        _ = v79
                                        _ = v80
                                    }
                                }
                                var v81 int32 = 0
                                if true {
                                    if (503 == 159) {
                                        v81 = v81
                                        if (false || (false || false)) {
                                            _ = len(v4)
                                            switch "hello" {
                                            case "a":
                                                switch (fi71 / (((fi21 & fi71)) | 1)) {
                                                case 42:
                                                    if (18446744073709551615 > 18446744073709551615) {
                                                        fmt.Println(fi21, fi21, v81)
                                                        fmt.Println(fi21, v81)
                                                    } else {
                                                        var v82 int16 = -810
                                                        var v83 int8 = -1
                                                        _ = len(v2)
                                                        _ = v82
                                                        _ = v83
                                                    }
                                                case -681:
                                                    fmt.Println(fi21, fi21, fi71)
                                                    var v84 string = "foo bar"
                                                    switch (fi71 / ((fi21) | 1)) {
                                                    case 127:
                                                        var v85 uint = 2
                                                        _ = v85
                                                    default:
                                                        v84 = v84
                                                    }
                                                    _ = v84
                                                }
                                            default:
                                                v81 += (v81 | (v81 ^ (v81 + int32(fi71))))
                                                v81 = -1
                                            }
                                            v81 += (^(v81))
                                            if true {
                                                switch ("" + "hello") {
                                                case "":
                                                    fmt.Println(v81)
                                                case "foo bar":
                                                    var v86 int32 = v81
                                                    _ = v86
                                                default:
                                                    if false {
                                                        v87 := GenS3{F0: 550, F1: -1}
                                                        v81 += 42
                                                        v88 := v2[fi21]
                                                        if (1 == 2) {
                                                            if v88 {
                                                                switch "\n\t" {
                                                                case "a":
                                                                    var v89 uint64 = 4294967295
                                                                    _ = v89
                                                                case "":
                                                                    v81 = (v81 & v81)
                                                                case "foo bar":
                                                                    var v90 uint16 = 0
                                                                    _ = v90
                                                                default:
                                                                    if false {
                                                                        v91 := GenS7{F0: "foo bar", F1: 2147483647}
                                                                        _ = v91
                                                                    } else {
                                                                        switch "" {
                                                                        case "a":
                                                                            if v88 {
                                                                                switch (("foo bar" + (("hello" + "\n\t") + "\n\t")) + ("foo bar" + (("\n\t" + "") + ""))) {
                                                                                case "hello":
                                                                                    var v92 int = (^(fi71))
                                                                                    _ = v92
                                                                                case "\n\t":
                                                                                    var v93 uint = 255
                                                                                    _ = v93
                                                                                default:
                                                                                    var v94 int32 = v81
                                                                                    _ = v94
                                                                                }
                                                                                switch "\n\t" {
                                                                                case "foo bar":
                                                                                    switch "\n\t" {
                                                                                    case "\n\t":
                                                                                        var v95 string = "foo bar"
                                                                                        var v96 int8 = 1
                                                                                        v97 := GenS8{F0: 670, F1: 18446744073709551615, F2: 2}
                                                                                        _ = v95
                                                                                        _ = v96
                                                                                        _ = v97
                                                                                    default:
                                                                                        v98 := v4[2]
                                                                                        _ = v98
                                                                                    }
                                                                                case "a":
                                                                                    var v99 float64 = -2.25
                                                                                    _ = v99
                                                                                }
                                                                            } else {
                                                                                delete(v2, fi71)
                                                                                v81 += v81
                                                                                var v100 float64 = -2.25
                                                                                _ = v100
                                                                            }
                                                                        case "\n\t":
                                                                            v88 = (!v88)
                                                                        }
                                                                        if v88 {
                                                                            v101 := GenS9{F0: 255, F1: genFn0((!(1e10 > 1.5)), (!(v88 && (2147483647 >= v81)))), F2: "foo bar"}
                                                                            v102 := GenS10{F0: -1, F1: 1}
                                                                            v103 := GenS3{F0: 1, F1: -286}
                                                                            _ = v101
                                                                            _ = v102
                                                                            _ = v103
                                                                        }
                                                                        v104 := GenS5{F0: 42, F1: 2, F2: v81}
                                                                        fmt.Println(v81, fi21, v81)
                                                                        _ = v104
                                                                    }
                                                                }
                                                                v81 += v81
                                                                if (0 > 2) {
                                                                    v88 = (false && (fi21 < (fi71 >> ((18446744073709551615) & 63))))
                                                                    if len(v3) > 0 {
                                                                        v3[0] = fi71
                                                                    }
                                                                    if (((1.5 / (float32(-2.25) / 1.0)) >= 3.14159) || false) {
                                                                        v81 = v81
                                                                        v105 := GenS10{F0: 463, F1: 590}
                                                                        var v106 uint = 509
                                                                        v106 -= v106
                                                                        _ = v105
                                                                        _ = v106
                                                                    }
                                                                    if true {
                                                                        v88 = true
                                                                        if (v88 || v88) {
                                                                            fmt.Println(v88, v88)
                                                                            switch (fi71 << ((4294967295) & 63)) {
                                                                            case -9223372036854775808:
                                                                                var v107 int8 = 42
                                                                                _ = v107
                                                                            case 32768:
                                                                                if v88 {
                                                                                    if (42 > 42) {
                                                                                        v88 = (!("a" <= "hello"))
                                                                                    } else {
                                                                                        if v88 {
                                                                                            v81 += (v81 / (((v81 >> ((600) & 63))) | 1))
                                                                                        }
                                                                                        var v108 string = "a"
                                                                                        _ = v108
                                                                                    }
                                                                                }
                                                                            case 127:
                                                                                if ((v81 | v81) != -1) {
                                                                                    fmt.Println(v88)
                                                                                    v1 = append(v1, 934)
                                                                                } else {
                                                                                    if false {
                                                                                        v88 = ((-2147483648 == fi21) || v88)
                                                                                    }
                                                                                    v109 := GenS11{F0: fi21, F1: 175, F2: v88}
                                                                                    if v88 {
                                                                                        switch "hello" {
                                                                                        case "a":
                                                                                            switch ("\n\t" + "hello") {
                                                                                            case "foo bar":
                                                                                                if len(v3) > 0 {
                                                                                                    v3[0] = (fi71 + (fi71 + fi21))
                                                                                                }
                                                                                            case "hello":
                                                                                                switch "\n\t" {
                                                                                                case "foo bar":
                                                                                                    fmt.Println(fi21, fi21, v81)
                                                                                                case "a":
                                                                                                    v88 = (-1 <= -32768)
                                                                                                }
                                                                                            default:
                                                                                                var v110 uint8 = 42
                                                                                                _ = v110
                                                                                            }
                                                                                        case "\n\t":
                                                                                            switch (fi21 - (fi71 * fi21)) {
                                                                                            case 141:
                                                                                                var v111 uint = 723
                                                                                                _ = v111
                                                                                            case -420:
                                                                                                delete(v2, 2147483647)
                                                                                            case 0:
                                                                                                v88 = (0 >= 42)
                                                                                            }
                                                                                        case "hello":
                                                                                            var v112 int32 = 2147483647
                                                                                            _ = v112
                                                                                        }
                                                                                        var v113 uint = 65535
                                                                                        fmt.Println(fi71, v81)
                                                                                        _ = v113
                                                                                    } else {
                                                                                        if ((v88 || (!((!v88) && ((1e10 - -2.25) >= 0.0)))) && v88) {
                                                                                            var v114 int8 = 127
                                                                                            switch fi71 {
                                                                                            case -2147483648:
                                                                                                v114 += (v114 - 91)
                                                                                                v114 += (^(v114))
                                                                                            case 264:
                                                                                                var v115 int64 = 814
                                                                                                _ = v115
                                                                                            case 32768:
                                                                                                switch ("" + "hello") {
                                                                                                case "foo bar":
                                                                                                    var v116 string = ""
                                                                                                    _ = v116
                                                                                                case "a":
                                                                                                    v117 := GenS12{F0: 482, F1: 83, F2: 42}
                                                                                                    _ = v117
                                                                                                }
                                                                                            default:
                                                                                                fmt.Println(v88, fi21)
                                                                                            }
                                                                                            _ = v114
                                                                                        }
                                                                                        v118 := GenS13{F0: -114, F1: 31, F2: (v81 | v81)}
                                                                                        if false {
                                                                                            v88 = (v88 || v88)
                                                                                            var v119 uint16 = 355
                                                                                            _ = v119
                                                                                        } else {
                                                                                            v88 = v88
                                                                                            if (fi71 >= fi71) {
                                                                                                fmt.Println(fi21)
                                                                                            }
                                                                                            if (!v88) {
                                                                                                var v120 uint32 = 2147483647
                                                                                                if true {
                                                                                                    v121 := GenS9{F0: 0, F1: 0, F2: "hello"}
                                                                                                    if (!v88) {
                                                                                                        switch "a" {
                                                                                                        case "hello":
                                                                                                            switch (fi71 - fi21) {
                                                                                                            case 670:
                                                                                                                v81 = v81
                                                                                                            case 128:
                                                                                                                var v122 uint8 = 252
                                                                                                                _ = v122
                                                                                                            case -2147483648:
                                                                                                                if ((^(v120)) == (v120 + (v120 + v120))) {
                                                                                                                    if (3.14159 < (1e10 * 3.14159)) {
                                                                                                                        fmt.Println(fi21, fi71, fi71)
                                                                                                                    } else {
                                                                                                                        v81 += 696
                                                                                                                        fmt.Println(v88, v88)
                                                                                                                        fmt.Println(v88, fi21)
                                                                                                                        if false {
                                                                                                                            fmt.Println(fi21, fi71, fi21)
                                                                                                                            v81 = v81
                                                                                                                        } else {
                                                                                                                            var v123 float64 = float64((((-2.25 / (((-2.25 + 1.5) * float32(3.14159)) + (0.1 + 0.0))) * 0.1) * (-2.25 / (((1.5 - 0.1) - 1.5) - 0.1))))
                                                                                                                            var v124 int8 = 100
                                                                                                                            if len(v3) > 0 {
                                                                                                                                v3[0] = 32768
                                                                                                                            }
                                                                                                                            if v88 {
                                                                                                                                var v125 float64 = 3.14159
                                                                                                                                switch "a" {
                                                                                                                                case "a":
                                                                                                                                    v88 = (-1 <= int16(v124))
                                                                                                                                case "foo bar":
                                                                                                                                    v123 = 0.1
                                                                                                                                case "":
                                                                                                                                    if (v88 || (v125 != ((0.0 - 1e10) + 0.0))) {
                                                                                                                                        var v126 int8 = 1
                                                                                                                                        v81 = (v81 | -911)
                                                                                                                                        v123 = v125
                                                                                                                                        _ = v126
                                                                                                                                    }
                                                                                                                                }
                                                                                                                                fmt.Println(fi21)
                                                                                                                                _ = v125
                                                                                                                            } else {
                                                                                                                                v123 = (v123 / v123)
                                                                                                                            }
                                                                                                                            v88 = (1 == 964)
                                                                                                                            _ = v123
                                                                                                                            _ = v124
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v127 := GenS7{F0: "", F1: (v81 * (v81 << ((2) & 63)))}
                                                                                                                    _ = v127
                                                                                                                }
                                                                                                            default:
                                                                                                                v88 = false
                                                                                                                var v128 float32 = ((0.1 - float32((((0.1 - 0.1) + -2.25) + (-2.25 - 0.0)))) * 1.5)
                                                                                                                var v129 string = ""
                                                                                                                var v130 int64 = 0
                                                                                                                _ = v128
                                                                                                                _ = v129
                                                                                                                _ = v130
                                                                                                            }
                                                                                                        case "a":
                                                                                                            v131 := GenS11{F0: (fi21 | fi71), F1: 0, F2: false}
                                                                                                            _ = v131
                                                                                                        }
                                                                                                        var v132 int32 = (v81 + (-(v81)))
                                                                                                        _ = v132
                                                                                                    } else {
                                                                                                        v2[(^(fi71))] = (128 != 19)
                                                                                                        v81 = int32(v120)
                                                                                                        v81 = 59
                                                                                                    }
                                                                                                    _ = v121
                                                                                                }
                                                                                                if (-2.25 <= 3.14159) {
                                                                                                    _ = len(v1)
                                                                                                    var v133 uint16 = 63
                                                                                                    _ = v133
                                                                                                } else {
                                                                                                    v2[(fi71 ^ fi21)] = (!v88)
                                                                                                }
                                                                                                _ = v120
                                                                                            } else {
                                                                                                if (2 != 909) {
                                                                                                    v81 -= 2147483647
                                                                                                    var v134 int8 = -128
                                                                                                    v88 = ("a" == (("a" + "foo bar") + "\n\t"))
                                                                                                    _ = v134
                                                                                                } else {
                                                                                                    fmt.Println(v88)
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        _ = v118
                                                                                    }
                                                                                    _ = v109
                                                                                }
                                                                            default:
                                                                                _ = len(v1)
                                                                            }
                                                                            var v135 uint32 = 752
                                                                            _ = v135
                                                                        } else {
                                                                            v136 := GenS14{F0: genFn1(), F1: 1, F2: -2147483648}
                                                                            _ = v136
                                                                        }
                                                                        v137 := GenS9{F0: 255, F1: 2, F2: "\n\t"}
                                                                        var v138 string = ("foo bar" + "hello")
                                                                        _ = v137
                                                                        _ = v138
                                                                    }
                                                                    var v139 uint64 = 1
                                                                    _ = v139
                                                                }
                                                            } else {
                                                                delete(v2, fi71)
                                                                v140 := GenS15{F0: -32768, F1: 0}
                                                                var v141 int64 = -9223372036854775808
                                                                _ = v140
                                                                _ = v141
                                                            }
                                                            switch (fi71 << ((2) & 63)) {
                                                            case 127:
                                                                v88 = (true || (!false))
                                                            }
                                                        } else {
                                                            switch fi21 {
                                                            case -9223372036854775808:
                                                                var v142 uint64 = 277
                                                                _ = v142
                                                            }
                                                            v143 := v4[2]
                                                            var v144 uint8 = 42
                                                            var v145 uint8 = (v144 + v144)
                                                            _ = v143
                                                            _ = v144
                                                            _ = v145
                                                        }
                                                        var v146 uint = 4294967295
                                                        _ = v146
                                                        _ = v87
                                                        _ = v88
                                                    }
                                                }
                                                if true {
                                                    var v147 float32 = (-2.25 * (1.5 + 0.0))
                                                    _ = v147
                                                } else {
                                                    var v148 int32 = -2147483648
                                                    if true {
                                                        delete(v4, 2147483647)
                                                        var v149 uint32 = 700
                                                        _ = v149
                                                    }
                                                    _ = v148
                                                }
                                            } else {
                                                v150 := GenS14{F0: 2, F1: 0, F2: (fi71 << ((255) & 63))}
                                                _ = v150
                                            }
                                            v81 -= v81
                                            v2[fi71] = true
                                            fmt.Println(v81, v81)
                                            var v151 int = 32768
                                            _ = v151
                                        }
                                    }
                                } else {
                                    v81 = (v81 | -630)
                                }
                                _ = fi71
                                _ = v81
                            }
                        }
                    }
                    var v152 uint64 = 4294967295
                    _ = v152
                }
            }
            _ = fi21
        }
    }
    v153 := GenS16{F0: (("foo bar" + "") + "foo bar"), F1: 255, F2: 2}
    v154 := GenS16{F0: "hello", F1: 692, F2: 165}
    var v155 int16 = 0
    var v156 string = "\n\t"
    v155 = (v155 - (v155 | (^(v155))))
    v156 = v156
    fmt.Println(v156)
    v155 = v155
    v156 = v156
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v155:", v155)
    fmt.Println("v156:", v156)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v153
    _ = v154
    _ = v155
    _ = v156
    _ = v2
    _ = v3
    _ = v4
}
