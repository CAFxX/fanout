package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int16; F2 string }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 uint64; F1 uint8; F2 string }
type GenS6 struct { F0 string; F1 uint64; F2 string }
type GenS7 struct { F0 int; F1 uint32 }
type GenS8 struct { F0 int64; F1 uint8; F2 int32 }
type GenS9 struct { F0 uint8; F1 uint8 }
type GenS10 struct { F0 uint64; F1 uint16 }
type GenS11 struct { F0 bool; F1 int }
type GenS12 struct { F0 int8; F1 uint }
type GenS13 struct { F0 uint8; F1 uint8; F2 int16 }
func genFn0() uint {
    var v5 int64 = 0
    v6 := []int{852, 2147483647, -642, 0}
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS1{F0: 703}
        for fi9 := 0; fi9 < 6; fi9++ {
            v5 = v5
            _ = fi9
        }
        v5 -= v5
        _ = fi7
        _ = v8
    }
    for fi10 := 0; fi10 < 4; fi10++ {
        switch "a" {
        case "":
            if (771 >= -2147483648) {
                v11 := GenS2{F0: "", F1: int16(fi10), F2: "hello"}
                if (true && true) {
                    var v12 int64 = v5
                    var v13 float32 = 0.1
                    _ = v12
                    _ = v13
                } else {
                    v14 := GenS1{F0: fi10}
                    v5 += (-(v5))
                    _ = v14
                }
                v5 = (v5 - (^(v5)))
                _ = v11
            } else {
                for fi15 := 0; fi15 < 10; fi15++ {
                    var v16 uint32 = 593
                    if (false || true) {
                        if true {
                            v5 += (^(v5))
                            if (!(!false)) {
                                fmt.Println(v5, fi15, v16)
                            }
                        } else {
                            v5 -= v5
                            var v17 int8 = 1
                            switch "hello" {
                            case "hello":
                                if (true || true) {
                                    v18 := GenS2{F0: "foo bar", F1: 32767, F2: "a"}
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    if true {
                                        v19 := GenS1{F0: int(v5)}
                                        _ = v19
                                    } else {
                                        if len(v6) > 0 {
                                            v6[0] = fi10
                                        }
                                        fmt.Println(fi10, fi15, v16)
                                        if len(v6) > 0 {
                                            v6[0] = 1
                                        }
                                    }
                                    _ = v18
                                } else {
                                    switch "foo bar" {
                                    case "":
                                        var v20 uint64 = 0
                                        _ = v20
                                    }
                                    if len(v6) > 0 {
                                        v6[0] = (fi15 - (fi10 >> ((42) & 63)))
                                    }
                                }
                            }
                            switch "a" {
                            case "foo bar":
                                v17 -= -1
                            case "":
                                switch (("a" + "\n\t") + "a") {
                                case "\n\t":
                                    fmt.Println(fi15, v5, v5)
                                }
                            default:
                                var v21 uint32 = v16
                                _ = v21
                            }
                            if false {
                                var v22 int8 = v17
                                if (false || false) {
                                    v17 += (v22 - 16)
                                    if (!(false && false)) {
                                        v5 = -1
                                        v17 -= 1
                                        v6 = append(v6, fi10)
                                    } else {
                                        if true {
                                            v6 = append(v6, (fi15 | (^(fi15))))
                                            fmt.Println(fi15)
                                            fmt.Println(v22, v16, v22)
                                        } else {
                                            if true {
                                                var v23 int32 = -302
                                                v24 := GenS3{F0: ("\n\t" + "a"), F1: 692}
                                                var v25 int = (^(fi10))
                                                _ = v23
                                                _ = v24
                                                _ = v25
                                            } else {
                                                v6 = append(v6, (fi10 * fi10))
                                            }
                                        }
                                        switch (("\n\t" + "") + "\n\t") {
                                        case "a":
                                            v26 := GenS4{F0: "foo bar", F1: 218}
                                            _ = v26
                                        default:
                                            var v27 bool = (156 == 65535)
                                            _ = v27
                                        }
                                        v17 = 100
                                    }
                                    v16 = (v16 / ((213) | 1))
                                } else {
                                    var v28 uint16 = 2
                                    if true {
                                        var v29 bool = false
                                        var v30 uint16 = v28
                                        if (42 <= -311) {
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            v17 = 1
                                            switch fi15 {
                                            case -2147483648:
                                                if v29 {
                                                    v28 = (^(v30))
                                                    switch (-(fi10)) {
                                                    case 4294967295:
                                                        v30 += (v30 ^ (v28 | (v28 / (((v28 ^ uint16(v16))) | 1))))
                                                    }
                                                    var v31 int = (fi10 - fi15)
                                                    _ = v31
                                                } else {
                                                    if true {
                                                        var v32 uint8 = 255
                                                        v5 += (v5 ^ v5)
                                                        v33 := GenS2{F0: "", F1: -258, F2: ""}
                                                        _ = v32
                                                        _ = v33
                                                    } else {
                                                        v34 := GenS5{F0: 2, F1: 255, F2: "hello"}
                                                        v16 += 282
                                                        v30 -= v30
                                                        _ = v34
                                                    }
                                                    if v29 {
                                                        if ("foo bar" <= (("a" + ("\n\t" + "")) + "a")) {
                                                            v28 += 0
                                                            v35 := GenS2{F0: "a", F1: 32767, F2: "\n\t"}
                                                            _ = v35
                                                        }
                                                    }
                                                }
                                            case 410:
                                                if v29 {
                                                    v36 := GenS6{F0: "", F1: 863, F2: "a"}
                                                    v6 = append(v6, 1)
                                                    v37 := GenS6{F0: "foo bar", F1: 18446744073709551615, F2: "foo bar"}
                                                    _ = v36
                                                    _ = v37
                                                }
                                            case 54:
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                            default:
                                                switch ("\n\t" + "a") {
                                                case "foo bar":
                                                    if v29 {
                                                        v6 = append(v6, 599)
                                                    } else {
                                                        v6 = append(v6, fi10)
                                                        v30 = 1
                                                        switch fi10 {
                                                        case 0:
                                                            if v29 {
                                                                v6 = append(v6, -201)
                                                                switch 4294967295 {
                                                                case 4294967295:
                                                                    v17 = v22
                                                                case 823:
                                                                    switch "a" {
                                                                    case "\n\t":
                                                                        fmt.Println(v30, v16)
                                                                    case "a":
                                                                        fmt.Println(v29, v5, v28)
                                                                    case "":
                                                                        v38 := GenS6{F0: "a", F1: 2, F2: "hello"}
                                                                        _ = v38
                                                                    }
                                                                }
                                                            } else {
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                                var v39 string = (("foo bar" + "") + "a")
                                                                v5 -= (v5 << ((42) & 63))
                                                                _ = v39
                                                            }
                                                        }
                                                    }
                                                case "a":
                                                    var v40 bool = false
                                                    _ = v40
                                                default:
                                                    var v41 bool = (!v29)
                                                    _ = v41
                                                }
                                            }
                                        } else {
                                            fmt.Println(fi15, v16, v16)
                                            switch fi10 {
                                            case -1:
                                                v16 += v16
                                            case 831:
                                                v30 += v30
                                            }
                                            if (4294967295 != 4294967295) {
                                                fmt.Println(fi15, v22, v17)
                                                if v29 {
                                                    switch "hello" {
                                                    case "foo bar":
                                                        if false {
                                                            var v42 string = "hello"
                                                            _ = v42
                                                        }
                                                    case "a":
                                                        if v29 {
                                                            var v43 float32 = (-2.25 + 3.14159)
                                                            var v44 uint8 = 2
                                                            _ = v43
                                                            _ = v44
                                                        }
                                                    case "":
                                                        if false {
                                                            v16 = v16
                                                            v45 := GenS7{F0: (-(fi10)), F1: 532}
                                                            _ = v45
                                                        }
                                                    default:
                                                        var v46 string = ((("a" + "a") + ("\n\t" + "")) + "hello")
                                                        _ = v46
                                                    }
                                                    _ = len(v6)
                                                    fmt.Println(fi15, v28, fi10)
                                                    v29 = (!((v30 + v28) != (v30 << ((42) & 63))))
                                                } else {
                                                    v29 = v29
                                                    v16 += (v16 & (v16 - (v16 * (v16 ^ v16))))
                                                    var v47 int32 = -703
                                                    _ = v47
                                                }
                                                v48 := GenS2{F0: "hello", F1: 32767, F2: "hello"}
                                                v29 = v29
                                                _ = v48
                                            }
                                        }
                                        _ = v29
                                        _ = v30
                                    }
                                    _ = v28
                                }
                                if (false || false) {
                                    v16 = 4294967295
                                } else {
                                    var v49 float32 = 0.0
                                    _ = v49
                                }
                                _ = v22
                            } else {
                                switch -1 {
                                case 0:
                                    var v50 int64 = (v5 - (v5 >> ((509) & 63)))
                                    _ = v50
                                case 127:
                                    if false {
                                        var v51 uint32 = (v16 * v16)
                                        _ = v51
                                    } else {
                                        var v52 bool = (2 != 2)
                                        var v53 int8 = v17
                                        _ = len(v6)
                                        _ = v52
                                        _ = v53
                                    }
                                case -791:
                                    v16 -= 4294967295
                                default:
                                    var v54 string = "hello"
                                    _ = v54
                                }
                                if ((^(v16)) > (v16 * (v16 / (((v16 + v16)) | 1)))) {
                                    if (4294967295 <= -2147483648) {
                                        if true {
                                            switch "\n\t" {
                                            case "hello":
                                                if true {
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v55 float32 = 0.1
                                                            _ = v55
                                                        default:
                                                            if ("" == "\n\t") {
                                                                var v56 uint64 = 875
                                                                _ = v56
                                                            } else {
                                                                if false {
                                                                    v17 -= v17
                                                                } else {
                                                                    v16 -= (v16 ^ v16)
                                                                    if (0 == 65535) {
                                                                        if len(v6) >= 2 {
                                                                            v6 = v6[1:len(v6)]
                                                                        }
                                                                        v16 += v16
                                                                    } else {
                                                                        v16 = (v16 >> ((1) & 63))
                                                                    }
                                                                }
                                                                if len(v6) >= 2 {
                                                                    v6 = v6[1:len(v6)]
                                                                }
                                                                if len(v6) > 0 {
                                                                    v6[0] = fi10
                                                                }
                                                            }
                                                            var v57 uint8 = 127
                                                            v6 = append(v6, fi10)
                                                            _ = v57
                                                        }
                                                    case "hello":
                                                        var v58 bool = false
                                                        _ = v58
                                                    case "":
                                                        v6 = append(v6, (fi10 - fi15))
                                                    default:
                                                        v5 -= v5
                                                    }
                                                    switch 32767 {
                                                    case 2147483647:
                                                        var v59 int = (fi10 / ((-461) | 1))
                                                        _ = v59
                                                    case -1:
                                                        switch (fi10 + fi15) {
                                                        case -441:
                                                            var v60 uint8 = 198
                                                            _ = v60
                                                        case -1:
                                                            switch fi10 {
                                                            case 128:
                                                                var v61 uint = 255
                                                                _ = v61
                                                            }
                                                        case 2147483647:
                                                            if true {
                                                                var v62 uint32 = (v16 & 42)
                                                                v16 = v62
                                                                switch (fi10 / ((12) | 1)) {
                                                                case 1:
                                                                    var v63 uint = 0
                                                                    v5 -= v5
                                                                    _ = v63
                                                                case 128:
                                                                    var v64 int = (fi10 | fi10)
                                                                    _ = v64
                                                                }
                                                                if true {
                                                                    var v65 uint = 42
                                                                    v65 += v65
                                                                    v66 := GenS7{F0: fi15, F1: (v62 / (((v16 * v62)) | 1))}
                                                                    _ = v65
                                                                    _ = v66
                                                                }
                                                                var v67 uint8 = 128
                                                                _ = v62
                                                                _ = v67
                                                            }
                                                        }
                                                    case 32767:
                                                        if false {
                                                            switch "hello" {
                                                            case "foo bar":
                                                                if (1 < v16) {
                                                                    var v68 int32 = -1
                                                                    v16 -= (v16 % ((v16) | 1))
                                                                    _ = v68
                                                                }
                                                            }
                                                            switch (^(fi15)) {
                                                            case -181:
                                                                var v69 bool = true
                                                                _ = v69
                                                            case 58:
                                                                var v70 uint16 = 860
                                                                _ = v70
                                                            case -9223372036854775808:
                                                                var v71 uint8 = 60
                                                                _ = v71
                                                            default:
                                                                fmt.Println(fi10)
                                                            }
                                                            fmt.Println(v17, v5)
                                                            v5 = v5
                                                        }
                                                    default:
                                                        var v72 int32 = -323
                                                        _ = v72
                                                    }
                                                    if true {
                                                        fmt.Println(fi15, v16, v16)
                                                    } else {
                                                        v73 := GenS8{F0: (v5 & 1), F1: 255, F2: -1}
                                                        switch fi15 {
                                                        case 1:
                                                            switch (("foo bar" + "foo bar") + "foo bar") {
                                                            case "hello":
                                                                _ = len(v6)
                                                            case "foo bar":
                                                                v74 := GenS9{F0: 0, F1: 0}
                                                                _ = v74
                                                            default:
                                                                var v75 int32 = 42
                                                                _ = v75
                                                            }
                                                        }
                                                        v16 += (v16 >> ((86) & 63))
                                                        if true {
                                                            var v76 int16 = 1
                                                            var v77 int64 = v5
                                                            var v78 float64 = ((1e10 + 0.1) / 1.5)
                                                            _ = v76
                                                            _ = v77
                                                            _ = v78
                                                        }
                                                        _ = v73
                                                    }
                                                    v79 := GenS4{F0: "", F1: (^(v16))}
                                                    switch "a" {
                                                    case "":
                                                        switch "\n\t" {
                                                        case "":
                                                            fmt.Println(fi15, fi10)
                                                        case "a":
                                                            v80 := GenS6{F0: "", F1: 244, F2: "a"}
                                                            v6 = append(v6, (fi10 ^ 4294967295))
                                                            v81 := GenS3{F0: (((("a" + "\n\t") + "hello") + "foo bar") + "\n\t"), F1: 721}
                                                            v82 := GenS5{F0: 42, F1: 2, F2: ("hello" + "a")}
                                                            _ = v80
                                                            _ = v81
                                                            _ = v82
                                                        }
                                                    default:
                                                        var v83 uint = 825
                                                        _ = v83
                                                    }
                                                    _ = v79
                                                } else {
                                                    fmt.Println(v5)
                                                    if false {
                                                        if true {
                                                            if false {
                                                                fmt.Println(fi15, v5)
                                                                var v84 uint32 = (v16 & v16)
                                                                _ = v84
                                                            } else {
                                                                switch fi15 {
                                                                case -1:
                                                                    var v85 uint16 = 32767
                                                                    v5 = int64(v17)
                                                                    _ = v85
                                                                case -373:
                                                                    v17 += 0
                                                                default:
                                                                    switch 4294967295 {
                                                                    case 4294967295:
                                                                        var v86 int32 = -901
                                                                        _ = v86
                                                                    case 9223372036854775807:
                                                                        v16 = v16
                                                                    }
                                                                }
                                                                v87 := GenS3{F0: "\n\t", F1: 424}
                                                                _ = v87
                                                            }
                                                            var v88 uint16 = 255
                                                            _ = v88
                                                        } else {
                                                            var v89 int32 = -381
                                                            var v90 uint8 = 0
                                                            _ = v89
                                                            _ = v90
                                                        }
                                                        var v91 int = 9223372036854775807
                                                        fmt.Println(fi10)
                                                        if (!true) {
                                                            var v92 uint32 = 1
                                                            var v93 int8 = -1
                                                            _ = v92
                                                            _ = v93
                                                        }
                                                        var v94 int16 = 0
                                                        if (false && ((65535 < 336) || (!false))) {
                                                            var v95 uint8 = 1
                                                            _ = v95
                                                        } else {
                                                            var v96 string = "a"
                                                            v97 := GenS10{F0: 1, F1: uint16(v16)}
                                                            _ = v96
                                                            _ = v97
                                                        }
                                                        var v98 string = "hello"
                                                        v5 -= 42
                                                        _ = v91
                                                        _ = v94
                                                        _ = v98
                                                    }
                                                }
                                            default:
                                                var v99 bool = (42 < -57)
                                                var v100 int = (fi10 >> ((42) & 63))
                                                _ = v100
                                                _ = v99
                                            }
                                        } else {
                                            v101 := GenS1{F0: -966}
                                            var v102 uint32 = (v16 / (((v16 >> ((197) & 63))) | 1))
                                            _ = v101
                                            _ = v102
                                        }
                                    } else {
                                        fmt.Println(v17, v17)
                                        switch ("hello" + "") {
                                        case "":
                                            if (false && true) {
                                                var v103 float32 = 1.5
                                                var v104 string = "hello"
                                                fmt.Println(v17, v104)
                                                v105 := GenS11{F0: (42 > -2147483648), F1: (fi15 | -52)}
                                                _ = v103
                                                _ = v104
                                                _ = v105
                                            } else {
                                                v17 = v17
                                                switch -9223372036854775808 {
                                                case 0:
                                                    if false {
                                                        var v106 uint32 = 2147483647
                                                        switch "" {
                                                        case "\n\t":
                                                            switch ("hello" + "a") {
                                                            case "foo bar":
                                                                var v107 int8 = -1
                                                                _ = v107
                                                            case "\n\t":
                                                                v106 += (v106 / (((v16 & v106)) | 1))
                                                                var v108 float32 = (3.14159 - -2.25)
                                                                switch fi10 {
                                                                case -1:
                                                                    v17 -= (v17 >> ((42) & 63))
                                                                }
                                                                _ = v108
                                                            }
                                                        case "hello":
                                                            v106 -= 681
                                                        default:
                                                            v6 = append(v6, (^(fi10)))
                                                        }
                                                        _ = v106
                                                    }
                                                default:
                                                    switch "hello" {
                                                    case "hello":
                                                        switch fi15 {
                                                        case 2147483647:
                                                            v16 += v16
                                                        default:
                                                            if true {
                                                                if (false && false) {
                                                                    v109 := GenS12{F0: (v17 - (v17 << ((18446744073709551615) & 63))), F1: 716}
                                                                    _ = v109
                                                                }
                                                            }
                                                        }
                                                    }
                                                    var v110 uint16 = 255
                                                    _ = v110
                                                }
                                                fmt.Println(fi15, fi10)
                                                v111 := GenS13{F0: 124, F1: 127, F2: 880}
                                                v112 := GenS3{F0: "hello", F1: 534}
                                                _ = v111
                                                _ = v112
                                            }
                                        case "hello":
                                            switch fi15 {
                                            case 36:
                                                if false {
                                                    fmt.Println(fi15)
                                                } else {
                                                    v5 += 9223372036854775807
                                                }
                                            case -2147483648:
                                                v113 := GenS4{F0: ((("\n\t" + "\n\t") + "") + ""), F1: (^(v16))}
                                                _ = v113
                                            }
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            switch (fi10 / ((42) | 1)) {
                                            case -942:
                                                var v114 int64 = 9223372036854775807
                                                _ = v114
                                            case 128:
                                                v115 := GenS5{F0: 755, F1: 255, F2: "\n\t"}
                                                _ = v115
                                            case 4294967295:
                                                var v116 uint16 = 65535
                                                _ = v116
                                            default:
                                                var v117 uint64 = 620
                                                _ = v117
                                            }
                                        case "foo bar":
                                            v16 = 65535
                                        }
                                    }
                                }
                            }
                            var v118 bool = (18446744073709551615 < 0)
                            _ = v118
                            _ = v17
                        }
                        var v119 uint64 = 909
                        if (!false) {
                            if (-1 > (v5 * v5)) {
                                _ = len(v6)
                                var v120 uint = 208
                                v120 -= 640
                                _ = v120
                            }
                        }
                        _ = v119
                    } else {
                        if len(v6) > 0 {
                            v6[0] = fi10
                        }
                        var v121 uint = 18446744073709551615
                        _ = v121
                    }
                    var v122 int32 = 2147483647
                    _ = fi15
                    _ = v122
                    _ = v16
                }
                _ = len(v6)
                if true {
                    var v123 bool = false
                    fmt.Println(v123, v5)
                    _ = v123
                }
            }
        default:
            v124 := GenS6{F0: "\n\t", F1: 1, F2: ("" + "foo bar")}
            _ = v124
        }
        _ = fi10
    }
    for fi125 := 0; fi125 < 6; fi125++ {
        v5 -= v5
        _ = fi125
    }
    _ = v5
    _ = v6
    return 2
}

func main() {
    v1 := []string{("" + "a"), ""}
    v2 := []uint32{4294967295}
    var v3 uint = 1
    v4 := []string{"foo bar", "foo bar", (("a" + "foo bar") + ""), (("\n\t" + "") + "a")}
    var v126 float64 = float64((0.0 - -2.25))
    var v127 int = 0
    switch v127 {
    case 32767:
        var v128 float64 = v126
        _ = v128
    case -1:
        var v129 float32 = 3.14159
        _ = v129
    case 4294967295:
        v127 = (v127 ^ 1)
    default:
        var v130 uint8 = 128
        _ = v130
    }
    var v131 bool = (false && false)
    v126 = v126
    v126 = v126
    v132 := GenS4{F0: ("hello" + "a"), F1: 1}
    for fi133 := 0; fi133 < 2; fi133++ {
        fmt.Println(v127, v127, v131)
        var v134 string = ""
        _ = fi133
        _ = v134
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v126:", v126)
    fmt.Println("v127:", v127)
    fmt.Println("v131:", v131)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v126
    _ = v127
    _ = v131
    _ = v132
    _ = v2
    _ = v3
    _ = v4
}
