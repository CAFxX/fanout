package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 int16; F1 uint64; F2 int32 }
type GenS4 struct { F0 int64; F1 uint32; F2 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int8; F1 uint64; F2 int64 }
type GenS9 struct { F0 uint; F1 uint64 }
type GenS10 struct { F0 int32; F1 int32 }
type GenS11 struct { F0 int32; F1 int32 }
type GenS12 struct { F0 uint64; F1 uint16; F2 uint64 }
type GenS13 struct { F0 uint8; F1 uint32 }

func main() {
    var v1 bool = (!(false || (!true)))
    v2 := []string{(("\n\t" + "a") + (("\n\t" + "hello") + "foo bar")), ("hello" + "\n\t"), "a", "\n\t"}
    v3 := make(map[string]int)
    v1 = v1
    v1 = true
    v1 = v1
    v4 := GenS1{F0: "", F1: 321}
    for fi5 := 0; fi5 < 19; fi5++ {
        v1 = (280 <= 477)
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 int32 = 0
            if (-362 >= 0) {
                var v8 uint32 = 4294967295
                if (false && (((false && (!v1)) || (0 < -9223372036854775808)) && true)) {
                    v9 := GenS1{F0: "a", F1: uint16(v8)}
                    delete(v3, "\n\t")
                    var v10 int32 = -2147483648
                    var v11 uint16 = 12
                    _ = v10
                    _ = v11
                    _ = v9
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v12 float32 = 3.14159
                    _ = v12
                }
                _ = v8
            } else {
                var v13 int = fi6
                var v14 uint = 0
                _ = v13
                _ = v14
            }
            if (9223372036854775807 != (fi6 * fi6)) {
                v7 = v7
                var v15 int8 = 127
                v1 = (!true)
                if false {
                    var v16 int64 = 9223372036854775807
                    v17 := v3["foo bar"]
                    _ = v16
                    _ = v17
                } else {
                    v18 := GenS2{F0: "hello", F1: -1}
                    var v19 int64 = int64(v7)
                    _ = len(v2)
                    _ = v18
                    _ = v19
                }
                switch fi5 {
                case -526:
                    v7 += v7
                case 42:
                    if ((!v1) || v1) {
                        v15 -= (v15 - (-(v15)))
                    } else {
                        v20 := GenS3{F0: 42, F1: 18446744073709551615, F2: (-(v7))}
                        _ = v20
                    }
                case 32767:
                    switch -689 {
                    case 1:
                        switch fi5 {
                        case 9223372036854775807:
                            fmt.Println(fi6)
                        case -210:
                            v7 = v7
                        default:
                            if v1 {
                                v15 += 127
                                v7 = (v7 << ((835) & 63))
                            } else {
                                fmt.Println(fi5, fi5, fi6)
                            }
                        }
                    }
                default:
                    v21 := v3["a"]
                    _ = v21
                }
                v7 = (v7 * (v7 - v7))
                v22 := v3[("hello" + "a")]
                _ = v15
                _ = v22
            } else {
                v7 -= (v7 % ((42) | 1))
            }
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        v24 := GenS4{F0: -836, F1: 391, F2: -9223372036854775808}
        _ = fi23
        _ = v24
    }
    fmt.Println(v1)
    switch "" {
    case "\n\t":
        if v1 {
            for fi25 := 0; fi25 < 4; fi25++ {
                fmt.Println(fi25, v1)
                v26 := GenS3{F0: 42, F1: 1, F2: 0}
                _ = fi25
                _ = v26
            }
            for fi27 := 0; fi27 < 18; fi27++ {
                fmt.Println(v1)
                _ = fi27
            }
            if (!v1) {
                if v1 {
                    fmt.Println(v1)
                    v2 = append(v2, "")
                } else {
                    v28 := GenS2{F0: "\n\t", F1: 128}
                    _ = v28
                }
            } else {
                fmt.Println(v1)
                switch 127 {
                case 0:
                    if v1 {
                        var v29 string = "a"
                        for fi30 := 0; fi30 < 9; fi30++ {
                            v1 = (v1 && v1)
                            _ = fi30
                        }
                        _ = v29
                    } else {
                        v1 = (!v1)
                    }
                }
                var v31 int8 = -128
                _ = v31
            }
        } else {
            for fi32 := 0; fi32 < 12; fi32++ {
                fmt.Println(fi32, fi32)
                var v33 uint16 = 65535
                _ = fi32
                _ = v33
            }
            v1 = true
        }
    case "foo bar":
        fmt.Println(v1)
    case "hello":
        v2 = append(v2, "a")
    default:
        fmt.Println(v1)
    }
    if v1 {
        switch (("\n\t" + "hello") + "") {
        case "\n\t":
            v34 := GenS2{F0: "", F1: -9223372036854775808}
            _ = v34
        case "hello":
            if len(v2) > 0 {
                v2[0] = ""
            }
        case "foo bar":
            v1 = (42 < 1)
        }
    } else {
        for fi35 := 0; fi35 < 15; fi35++ {
            for fi36 := 0; fi36 < 4; fi36++ {
                v37 := GenS5{F0: 42}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi36
                _ = v37
            }
            var v38 int16 = -395
            _ = fi35
            _ = v38
        }
    }
    for fi39 := 0; fi39 < 8; fi39++ {
        switch "" {
        case "hello":
            for fi40 := 0; fi40 < 15; fi40++ {
                if v1 {
                    switch fi39 {
                    case 32767:
                        var v41 uint8 = 128
                        _ = v41
                    case 9223372036854775807:
                        if v1 {
                            v42 := GenS6{F0: 1}
                            switch "foo bar" {
                            case "a":
                                if len(v2) > 0 {
                                    v2[0] = (("foo bar" + "hello") + "")
                                }
                            case "\n\t":
                                var v43 int8 = -128
                                _ = v43
                            }
                            _ = v42
                        } else {
                            v44 := GenS5{F0: 9223372036854775807}
                            _ = v44
                        }
                    default:
                        v45 := GenS7{F0: -794}
                        _ = v45
                    }
                    v1 = true
                    fmt.Println(fi39)
                }
                if ((255 <= 1) || ((v1 || v1) || (1 == -788))) {
                    v46 := GenS8{F0: 1, F1: 42, F2: 453}
                    _ = v46
                } else {
                    v1 = true
                }
                var v47 int16 = 42
                var v48 bool = v1
                var v49 float64 = -2.25
                _ = fi40
                _ = v47
                _ = v48
                _ = v49
            }
        }
        if v1 {
            var v50 int = fi39
            if (!((!(v1 || true)) && (("foo bar" + ("foo bar" + (("" + "hello") + "hello"))) >= "\n\t"))) {
                for fi51 := 0; fi51 < 1; fi51++ {
                    if true {
                        fmt.Println(v50, fi51, v50)
                        switch (fi51 % ((fi39) | 1)) {
                        case 815:
                            v50 = v50
                        case -1:
                            var v52 uint8 = 151
                            _ = v52
                        default:
                            _ = len(v3)
                        }
                        if (false && (2147483648 != 4294967295)) {
                            v1 = v1
                        } else {
                            var v53 int = -2147483648
                            var v54 string = "\n\t"
                            if (-1 < 0) {
                                v55 := GenS6{F0: 9223372036854775807}
                                if len(v2) > 0 {
                                    v2[0] = "\n\t"
                                }
                                v56 := GenS1{F0: "\n\t", F1: 917}
                                _ = v55
                                _ = v56
                            } else {
                                fmt.Println(v1, fi39, fi39)
                                v1 = (v1 && (v1 || v1))
                                switch 1 {
                                case 788:
                                    switch (v54 + v54) {
                                    case "hello":
                                        var v57 uint16 = 65535
                                        _ = v57
                                    default:
                                        var v58 uint64 = 69
                                        _ = v58
                                    }
                                case 128:
                                    fmt.Println(v54)
                                default:
                                    var v59 int16 = -32768
                                    fmt.Println(v59)
                                    v59 = v59
                                    fmt.Println(v59, v1, fi39)
                                    fmt.Println(v1, v53, v50)
                                    if (-9223372036854775808 > 0) {
                                        var v60 int32 = 156
                                        _ = v60
                                    } else {
                                        delete(v3, v54)
                                    }
                                    v50 = (fi39 - (fi51 << ((382) & 63)))
                                    _ = v59
                                }
                            }
                            _ = v53
                            _ = v54
                        }
                    } else {
                        var v61 float64 = ((-2.25 / ((float64((0.1 / (1.5 + 3.14159))) * (1.5 + 1.5)) / (float64((((-2.25 * 1.5) / 1.0) - -2.25)) * (float64(-2.25) + (0.0 + -2.25))))) - 3.14159)
                        var v62 uint64 = 4294967295
                        _ = v61
                        _ = v62
                    }
                    var v63 int8 = -128
                    v50 += -2147483648
                    _ = fi51
                    _ = v63
                }
                for fi64 := 0; fi64 < 9; fi64++ {
                    if false {
                        fmt.Println(fi39, v1)
                        fmt.Println(fi39, fi39, fi64)
                    } else {
                        var v65 uint8 = 42
                        v3["\n\t"] = -509
                        if (((fi39 << ((4294967295) & 63)) < (fi39 / (((fi64 % ((v50) | 1))) | 1))) && (45 != 127)) {
                            v65 = (v65 / ((17) | 1))
                            v3[("hello" + "")] = v50
                        }
                        _ = v65
                    }
                    _ = fi64
                }
                var v66 int = v50
                _ = v66
            } else {
                for fi67 := 0; fi67 < 4; fi67++ {
                    v50 += fi39
                    v68 := GenS3{F0: 0, F1: 561, F2: 961}
                    v50 = fi67
                    _ = fi67
                    _ = v68
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = len(v3)
            }
            if (!v1) {
                var v69 int64 = -9223372036854775808
                if (false || v1) {
                    var v70 uint16 = 32767
                    var v71 int = v50
                    _ = v70
                    _ = v71
                }
                _ = v69
            }
            _ = v50
        }
        v1 = v1
        _ = fi39
    }
    for fi72 := 0; fi72 < 17; fi72++ {
        for fi73 := 0; fi73 < 2; fi73++ {
            var v74 float64 = float64(1.5)
            if (42 <= 1) {
                if (((v74 * float64(-2.25)) == (-2.25 + v74)) && v1) {
                    var v75 int = (fi73 / ((fi72) | 1))
                    v76 := GenS6{F0: 618}
                    _ = v75
                    _ = v76
                }
                switch fi72 {
                case -606:
                    var v77 int = 42
                    _ = v77
                case 9223372036854775807:
                    v74 = v74
                case 127:
                    switch (("\n\t" + "\n\t") + ("a" + "foo bar")) {
                    case "\n\t":
                        fmt.Println(v1, fi72, fi72)
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v78 := GenS8{F0: -1, F1: 393, F2: 1}
                        if (!v1) {
                            _ = len(v3)
                            v79 := GenS5{F0: -587}
                            switch "\n\t" {
                            case "\n\t":
                                v3[("\n\t" + "\n\t")] = (fi73 - fi72)
                            case "foo bar":
                                if (!((!(v1 && v1)) || v1)) {
                                    v80 := GenS7{F0: 42}
                                    switch "hello" {
                                    case "foo bar":
                                        fmt.Println(fi73, fi72, fi73)
                                    }
                                    var v81 float32 = float32(v74)
                                    _ = v80
                                    _ = v81
                                }
                            default:
                                var v82 uint64 = 0
                                _ = v82
                            }
                            _ = v79
                        }
                        if false {
                            _ = len(v3)
                            v83 := GenS2{F0: "a", F1: fi72}
                            if (861 == 1) {
                                v3[""] = fi72
                            } else {
                                v1 = false
                                var v84 string = "foo bar"
                                v85 := GenS9{F0: 0, F1: 18446744073709551615}
                                _ = v84
                                _ = v85
                            }
                            v1 = true
                            _ = len(v2)
                            _ = v83
                        } else {
                            delete(v3, "foo bar")
                            if (!(!(!v1))) {
                                v1 = true
                                _ = len(v3)
                                fmt.Println(v1)
                            }
                        }
                        v86 := GenS1{F0: "a", F1: 586}
                        v2 = append(v2, "")
                        _ = v78
                        _ = v86
                    case "hello":
                        var v87 uint32 = 296
                        _ = v87
                    case "":
                        var v88 int = (^(fi72))
                        _ = v88
                    default:
                        v1 = (v1 && ((1e10 + 1e10) <= 0.0))
                    }
                    v89 := GenS10{F0: 1, F1: 191}
                    var v90 int64 = 1
                    _ = v89
                    _ = v90
                }
            } else {
                if v1 {
                    var v91 int8 = -22
                    _ = v91
                }
            }
            _ = fi73
            _ = v74
        }
        _ = fi72
    }
    for fi92 := 0; fi92 < 3; fi92++ {
        for fi93 := 0; fi93 < 12; fi93++ {
            v1 = false
            switch "a" {
            case "hello":
                var v94 uint16 = 255
                _ = v94
            case "a":
                switch (-(fi92)) {
                case -924:
                    if (!((-77 >= -57) && v1)) {
                        if (2 <= 2) {
                            var v95 float32 = -2.25
                            v96 := GenS10{F0: 1, F1: -743}
                            fmt.Println(v95)
                            if true {
                                switch (fi92 - fi92) {
                                case -9223372036854775808:
                                    switch -344 {
                                    case 671:
                                        v97 := GenS11{F0: -871, F1: -2147483648}
                                        _ = v97
                                    default:
                                        v98 := GenS9{F0: 413, F1: 496}
                                        var v99 float32 = ((1.5 + (0.1 - (1.5 - float32(0.1)))) + (1e10 * v95))
                                        fmt.Println(v1, v99, fi93)
                                        _ = v98
                                        _ = v99
                                    }
                                    fmt.Println(fi92, v95, fi92)
                                default:
                                    switch fi93 {
                                    case -652:
                                        v3["hello"] = fi93
                                    case 2147483647:
                                        v100 := GenS12{F0: 4294967295, F1: 2, F2: 587}
                                        _ = v100
                                    default:
                                        if v1 {
                                            var v101 uint8 = 53
                                            if false {
                                                v102 := GenS13{F0: (v101 ^ (v101 % (((v101 & v101)) | 1))), F1: uint32(v101)}
                                                var v103 float32 = ((v95 + 0.0) * v95)
                                                _ = v102
                                                _ = v103
                                            } else {
                                                v95 += (1.5 - (v95 * 0.0))
                                                if v1 {
                                                    switch fi92 {
                                                    case 1:
                                                        if len(v2) > 0 {
                                                            v2[0] = "\n\t"
                                                        }
                                                    case 127:
                                                        var v104 uint = 42
                                                        v3[("foo bar" + "hello")] = (fi92 * (fi92 >> ((1) & 63)))
                                                        v104 += (v104 % (((v104 ^ (v104 | v104))) | 1))
                                                        _ = v104
                                                    default:
                                                        switch (("" + "\n\t") + "\n\t") {
                                                        case "foo bar":
                                                            v105 := GenS5{F0: 1}
                                                            _ = v105
                                                        }
                                                    }
                                                    switch (fi93 << ((18446744073709551615) & 63)) {
                                                    case -612:
                                                        var v106 string = ("hello" + ("a" + ""))
                                                        _ = v106
                                                    default:
                                                        var v107 string = ("a" + "")
                                                        _ = v107
                                                    }
                                                    v108 := GenS8{F0: -128, F1: 208, F2: -516}
                                                    _ = v108
                                                }
                                                v95 = v95
                                            }
                                            fmt.Println(v101, v1)
                                            _ = v101
                                        } else {
                                            var v109 int = -2147483648
                                            if v1 {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                var v110 float32 = 1.5
                                                if v1 {
                                                    v95 -= (((float32(float64(1.5)) + (v110 / (v95 / -2.25))) / v110) * v110)
                                                    if (-990 < 42) {
                                                        v95 -= float32((0.1 - ((0.0 - float64(((((v110 * (v110 * (0.1 / ((v110 + v95) - 3.14159)))) + v95) * v95) - v110))) - (1e10 / 0.1))))
                                                    }
                                                    var v111 int64 = 42
                                                    v112 := GenS3{F0: -32768, F1: 638, F2: -1}
                                                    if v1 {
                                                        var v113 int = 333
                                                        if v1 {
                                                            _ = len(v3)
                                                            var v114 float32 = v95
                                                            if v1 {
                                                                var v115 int = (-(v113))
                                                                v116 := GenS5{F0: v111}
                                                                var v117 uint16 = 330
                                                                _ = v115
                                                                _ = v116
                                                                _ = v117
                                                            } else {
                                                                v111 = v111
                                                                if (0 < 1000) {
                                                                    v118 := GenS11{F0: -1, F1: 59}
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    _ = v118
                                                                } else {
                                                                    var v119 int16 = 0
                                                                    _ = v119
                                                                }
                                                            }
                                                            _ = v114
                                                        }
                                                        _ = v113
                                                    }
                                                    delete(v3, "\n\t")
                                                    v120 := GenS5{F0: v111}
                                                    var v121 float32 = (v95 / 1.5)
                                                    var v122 int8 = 0
                                                    v122 = v122
                                                    _ = v111
                                                    _ = v112
                                                    _ = v120
                                                    _ = v121
                                                    _ = v122
                                                }
                                                _ = v110
                                            } else {
                                                if (1 != -95) {
                                                    v1 = (127 < -1)
                                                }
                                                v3[""] = (v109 & (v109 / ((0) | 1)))
                                            }
                                            fmt.Println(v95)
                                            _ = v109
                                        }
                                    }
                                }
                                v1 = true
                                var v123 float32 = 3.14159
                                fmt.Println(fi92)
                                var v124 uint = 693
                                _ = v123
                                _ = v124
                            }
                            if (!v1) {
                                var v125 int32 = -815
                                _ = v125
                            } else {
                                v1 = v1
                                switch 42 {
                                case -379:
                                    var v126 string = ""
                                    _ = v126
                                case -160:
                                    v127 := GenS8{F0: -117, F1: 42, F2: -1}
                                    v128 := v3["hello"]
                                    if (true && v1) {
                                        v1 = v1
                                        v95 = ((v95 - v95) + v95)
                                    }
                                    if true {
                                        var v129 uint16 = 780
                                        var v130 uint = 1
                                        _ = v129
                                        _ = v130
                                    }
                                    _ = v127
                                    _ = v128
                                }
                                var v131 int64 = 9223372036854775807
                                _ = v131
                            }
                            _ = v95
                            _ = v96
                        } else {
                            if ((128 > 42) && false) {
                                v1 = (2 == 833)
                                v1 = (false && (false && v1))
                                if (4294967295 <= 4294967295) {
                                    var v132 float64 = (1.5 * 3.14159)
                                    var v133 uint32 = 0
                                    _ = v132
                                    _ = v133
                                } else {
                                    v1 = false
                                }
                            }
                            var v134 uint32 = 2
                            _ = v134
                        }
                    }
                case -344:
                    var v135 float64 = float64((0.0 / 1.0))
                    _ = v135
                case 393:
                    v1 = (42 < 0)
                default:
                    var v136 int8 = 0
                    _ = v136
                }
            default:
                v137 := GenS12{F0: 42, F1: 32767, F2: 127}
                _ = v137
            }
            _ = fi93
        }
        var v138 uint32 = 4294967295
        _ = fi92
        _ = v138
    }
    v1 = (("" + "a") >= "a")
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
