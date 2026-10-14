package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS4 struct { F0 string; F1 uint32; F2 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int16; F1 uint32; F2 int }
type GenS7 struct { F0 uint16; F1 bool; F2 int8 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 string; F1 bool; F2 uint }
type GenS10 struct { F0 int64; F1 bool; F2 uint8 }
type GenS11 struct { F0 int32; F1 uint32; F2 string }

func main() {
    var v1 uint16 = 323
    v2 := make(map[string]bool)
    var v3 uint16 = 1
    v4 := []int{int(v1), 128, 32767, 1}
    var v5 int16 = 32767
    var v6 uint = 1
    var v7 int = 9223372036854775807
    _ = len(v2)
    var v8 uint32 = 2147483647
    v2["a"] = false
    switch (("\n\t" + ("a" + ("\n\t" + "a"))) + (("a" + (("hello" + "hello") + "foo bar")) + "foo bar")) {
    case "\n\t":
        if true {
            if false {
                fmt.Println(v3, v3)
            } else {
                switch "a" {
                case "a":
                    var v9 int64 = -96
                    _ = v9
                case "":
                    var v10 uint8 = 2
                    _ = v10
                case "foo bar":
                    switch "a" {
                    case "foo bar":
                        var v11 float64 = 0.0
                        _ = v11
                    case "hello":
                        var v12 int32 = 1
                        _ = v12
                    case "a":
                        v1 = v3
                    default:
                        var v13 int16 = v5
                        _ = v13
                    }
                }
                switch "" {
                case "hello":
                    if true {
                        v6 = 335
                    }
                case "a":
                    switch (v7 & v7) {
                    case 9223372036854775807:
                        switch v7 {
                        case 4294967295:
                            v5 = (v5 + (v5 ^ v5))
                        default:
                            var v14 uint8 = 1
                            _ = v14
                        }
                    case -2147483648:
                        for fi15 := 0; fi15 < 1; fi15++ {
                            for fi16 := 0; fi16 < 3; fi16++ {
                                var v17 int64 = -1
                                var v18 int32 = int32(v17)
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                            var v19 float32 = 0.1
                            var v20 uint16 = 46
                            _ = fi15
                            _ = v19
                            _ = v20
                        }
                    case 950:
                        fmt.Println(v6, v3)
                    default:
                        var v21 uint = v6
                        _ = v21
                    }
                case "foo bar":
                    v5 += (-(v5))
                default:
                    _ = len(v4)
                }
                var v22 uint64 = 0
                _ = v22
            }
            var v23 uint16 = uint16(v8)
            _ = v23
        }
    case "":
        if (-1 >= -333) {
            var v24 float32 = 0.1
            var v25 float32 = (v24 / v24)
            v1 = 65535
            _ = v24
            _ = v25
        } else {
            var v26 uint8 = 125
            var v27 uint32 = v8
            _ = v26
            _ = v27
        }
    case "hello":
        v1 -= v3
    default:
        v28 := GenS1{F0: ""}
        _ = v28
    }
    var v29 int = (v7 + 854)
    v7 += v7
    switch ("" + "foo bar") {
    case "\n\t":
        if (0 > 1) {
            if true {
                v29 -= (v7 ^ v7)
            }
        } else {
            for fi30 := 0; fi30 < 11; fi30++ {
                var v31 int64 = 42
                v5 += int16(v31)
                if (!false) {
                    switch "hello" {
                    case "hello":
                        v6 = (^(v6))
                    default:
                        switch ("\n\t" + ("a" + "\n\t")) {
                        case "":
                            var v32 bool = false
                            _ = v32
                        case "a":
                            _ = len(v4)
                        }
                    }
                    if true {
                        for fi33 := 0; fi33 < 19; fi33++ {
                            v7 += 2147483647
                            v34 := GenS1{F0: ("a" + "foo bar")}
                            _ = fi33
                            _ = v34
                        }
                        var v35 float64 = -2.25
                        var v36 uint8 = 38
                        _ = v35
                        _ = v36
                    } else {
                        v37 := v2["hello"]
                        fmt.Println(v1, v8, fi30)
                        _ = len(v4)
                        _ = v37
                    }
                    for fi38 := 0; fi38 < 3; fi38++ {
                        v39 := GenS1{F0: "hello"}
                        _ = fi38
                        _ = v39
                    }
                } else {
                    if len(v4) > 0 {
                        v4[0] = (v7 / (((v29 + fi30)) | 1))
                    }
                    if (!true) {
                        for fi40 := 0; fi40 < 4; fi40++ {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = fi40
                        }
                        if true {
                            if len(v4) > 0 {
                                v4[0] = (v29 / (((v29 % ((-417) | 1))) | 1))
                            }
                            if ((1 <= 69) && false) {
                                v3 -= v1
                                if true {
                                    v8 = (v8 / ((v8) | 1))
                                } else {
                                    var v41 float64 = 3.14159
                                    _ = v41
                                }
                            }
                            for fi42 := 0; fi42 < 8; fi42++ {
                                if len(v4) > 0 {
                                    v4[0] = v29
                                }
                                _ = fi42
                            }
                        }
                    }
                    var v43 uint32 = (v8 / ((v8) | 1))
                    _ = v43
                }
                _ = fi30
                _ = v31
            }
            for fi44 := 0; fi44 < 20; fi44++ {
                var v45 uint64 = 0
                _ = fi44
                _ = v45
            }
        }
    default:
        v3 = v1
    }
    var v46 string = "foo bar"
    v3 += (^(v3))
    delete(v2, ("\n\t" + "hello"))
    if (!(v46 >= v46)) {
        var v47 string = (v46 + v46)
        switch int(v6) {
        case 2147483647:
            var v48 uint32 = (v8 / (((v8 | 388)) | 1))
            _ = v48
        case 1:
            v6 += v6
        }
        var v49 int = v7
        _ = v47
        _ = v49
    } else {
        for fi50 := 0; fi50 < 16; fi50++ {
            v6 = (v6 | (v6 % (((v6 - 65535)) | 1)))
            if ((1 <= 18446744073709551615) && (1 <= -2147483648)) {
                if true {
                    v7 -= 128
                    for fi51 := 0; fi51 < 3; fi51++ {
                        v5 -= (v5 >> ((v6) & 63))
                        v7 -= fi51
                        _ = fi51
                    }
                    v8 -= v8
                }
                fmt.Println(v1, fi50, v46)
                if (((v5 * v5) >= -735) || true) {
                    if true {
                        for fi52 := 0; fi52 < 4; fi52++ {
                            switch -6 {
                            case 128:
                                v2[v46] = (false && false)
                            case -593:
                                v7 -= (^(fi52))
                            case 797:
                                if (v7 == v7) {
                                    var v53 int8 = -30
                                    _ = v53
                                }
                            }
                            _ = fi52
                        }
                        for fi54 := 0; fi54 < 6; fi54++ {
                            var v55 uint64 = uint64(v6)
                            switch (-(fi50)) {
                            case 2147483647:
                                if true {
                                    if (!true) {
                                        v56 := GenS2{F0: ""}
                                        if false {
                                            v6 -= (v6 << ((v6) & 63))
                                            var v57 int = 32767
                                            v58 := GenS3{F0: (v55 * (v55 & (v55 | (v55 >> (((v6 - (v6 * v6))) & 63))))), F1: 202, F2: 128}
                                            _ = v57
                                            _ = v58
                                        } else {
                                            fmt.Println(v46, v1)
                                            if ((1 < -1) && true) {
                                                fmt.Println(v29, v3)
                                                v59 := GenS1{F0: v46}
                                                v8 = (v8 * uint32(v6))
                                                var v60 bool = false
                                                var v61 int32 = int32(v55)
                                                _ = v59
                                                _ = v60
                                                _ = v61
                                            } else {
                                                var v62 int = (fi50 ^ (fi50 + int(v6)))
                                                v6 += (v6 % ((v6) | 1))
                                                if (v3 != uint16(v55)) {
                                                    v63 := GenS3{F0: 2, F1: 250, F2: 1}
                                                    _ = v63
                                                }
                                                var v64 int16 = (^(v5))
                                                switch "\n\t" {
                                                case "hello":
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                case "\n\t":
                                                    switch -2147483648 {
                                                    case -9223372036854775808:
                                                        switch (^(v7)) {
                                                        case 2147483647:
                                                            v65 := GenS2{F0: v46}
                                                            _ = v65
                                                        }
                                                    case 127:
                                                        v66 := GenS4{F0: "", F1: (^(v8)), F2: ("hello" + ((v46 + v46) + v46))}
                                                        _ = v66
                                                    }
                                                case "foo bar":
                                                    if false {
                                                        v29 -= (fi50 - (v62 & fi50))
                                                        fmt.Println(v55, v62)
                                                    } else {
                                                        if true {
                                                            v1 = v3
                                                            var v67 uint32 = (v8 << ((732) & 63))
                                                            _ = v67
                                                        } else {
                                                            v6 += v6
                                                            if true {
                                                                if false {
                                                                    fmt.Println(v1, v55, v3)
                                                                    v6 += v6
                                                                } else {
                                                                    v4 = append(v4, fi54)
                                                                    if ((v6 ^ (v6 * (v6 * v6))) >= v6) {
                                                                        _ = len(v2)
                                                                    }
                                                                    switch (v46 + "\n\t") {
                                                                    case "a":
                                                                        var v68 string = v46
                                                                        _ = v68
                                                                    case "\n\t":
                                                                        var v69 int16 = v64
                                                                        _ = v69
                                                                    case "foo bar":
                                                                        var v70 int8 = 118
                                                                        _ = v70
                                                                    }
                                                                    var v71 uint8 = 4
                                                                    _ = v71
                                                                }
                                                            }
                                                            v29 += 4294967295
                                                        }
                                                    }
                                                default:
                                                    var v72 uint8 = 1
                                                    _ = v72
                                                }
                                                _ = v62
                                                _ = v64
                                            }
                                        }
                                        _ = v56
                                    }
                                    if (2147483648 <= v8) {
                                        var v73 bool = (3.14159 != 1.5)
                                        _ = v73
                                    }
                                    if ((!true) && false) {
                                        switch v29 {
                                        case -1:
                                            if true {
                                                v6 = (v6 & 0)
                                                if true {
                                                    var v74 int16 = (^(v5))
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    switch (^(v7)) {
                                                    case 0:
                                                        v3 = v1
                                                    }
                                                    _ = v74
                                                }
                                                var v75 float32 = 1.5
                                                _ = v75
                                            }
                                        }
                                        v76 := GenS5{F0: (v1 << ((v6) & 63))}
                                        _ = v76
                                    } else {
                                        if false {
                                            var v77 int16 = 42
                                            v1 += 65535
                                            _ = v77
                                        } else {
                                            if (-2.25 <= 1e10) {
                                                var v78 int32 = 552
                                                var v79 int16 = (v5 * v5)
                                                _ = v78
                                                _ = v79
                                            }
                                        }
                                        if (2147483647 >= 2147483647) {
                                            v8 -= 2147483647
                                            if false {
                                                v3 = (v3 + uint16(v8))
                                                switch v46 {
                                                case "\n\t":
                                                    v80 := GenS1{F0: v46}
                                                    _ = v80
                                                case "a":
                                                    _ = len(v2)
                                                case "":
                                                    if false {
                                                        v1 = (^(v3))
                                                        var v81 uint64 = (v55 >> ((v6) & 63))
                                                        _ = v81
                                                    }
                                                default:
                                                    switch v46 {
                                                    case "":
                                                        var v82 float64 = 0.0
                                                        _ = v82
                                                    case "\n\t":
                                                        v1 = (v1 * v3)
                                                    }
                                                    if ((v46 >= v46) || ((-2.25 - 0.0) <= 0.0)) {
                                                        if ((true || true) && (3.14159 <= 1.5)) {
                                                            v1 -= (v3 << ((v6) & 63))
                                                            var v83 bool = (false || true)
                                                            if (9223372036854775807 > 9223372036854775807) {
                                                                delete(v2, (v46 + v46))
                                                                var v84 uint32 = 65535
                                                                if (v29 <= -2147483648) {
                                                                    fmt.Println(v46, v6, fi54)
                                                                    v2["\n\t"] = (float64((1.5 * 3.14159)) == 3.14159)
                                                                    v85 := GenS2{F0: "\n\t"}
                                                                    v86 := GenS2{F0: "foo bar"}
                                                                    _ = v85
                                                                    _ = v86
                                                                } else {
                                                                    _ = len(v4)
                                                                }
                                                                fmt.Println(fi50)
                                                                v87 := GenS1{F0: v46}
                                                                _ = v84
                                                                _ = v87
                                                            } else {
                                                                v55 = 2
                                                                var v88 float32 = 1e10
                                                                _ = v88
                                                            }
                                                            _ = v83
                                                        }
                                                        _ = len(v4)
                                                        var v89 string = "hello"
                                                        var v90 string = v89
                                                        _ = v89
                                                        _ = v90
                                                    }
                                                }
                                                if false {
                                                    if false {
                                                        v46 = v46
                                                        var v91 uint32 = v8
                                                        _ = v91
                                                    }
                                                }
                                                v92 := GenS3{F0: uint64(v8), F1: 2, F2: 127}
                                                _ = v92
                                            }
                                            v93 := GenS6{F0: v5, F1: (v8 + (v8 | (^(v8)))), F2: (v29 & (fi50 / ((v7) | 1)))}
                                            var v94 float32 = 0.0
                                            _ = v93
                                            _ = v94
                                        } else {
                                            if len(v4) > 0 {
                                                v4[0] = v7
                                            }
                                            if true {
                                                if (-1 >= -2147483648) {
                                                    v29 = -339
                                                    v55 -= (v55 << (((v6 / ((uint(v8)) | 1))) & 63))
                                                    switch (v46 + v46) {
                                                    case "a":
                                                        fmt.Println(v7)
                                                    case "\n\t":
                                                        var v95 float64 = (((3.14159 / (float64(1.5) + float64(-2.25))) - float64(float32(-2.25))) + 0.0)
                                                        _ = v95
                                                    default:
                                                        var v96 bool = (0 <= 197)
                                                        _ = v96
                                                    }
                                                    v8 = (^(v8))
                                                    switch (v29 >> ((v6) & 63)) {
                                                    case 9223372036854775807:
                                                        var v97 float32 = 1e10
                                                        _ = v97
                                                    case 0:
                                                        var v98 string = "foo bar"
                                                        _ = v98
                                                    case 2147483647:
                                                        var v99 uint = 255
                                                        _ = v99
                                                    }
                                                } else {
                                                    v100 := GenS7{F0: 198, F1: ((v55 & 42) < v55), F2: 127}
                                                    if false {
                                                        v101 := GenS2{F0: ("foo bar" + (v46 + v46))}
                                                        fmt.Println(v3, v3)
                                                        v102 := GenS7{F0: v1, F1: (2 != 45), F2: int8(v5)}
                                                        _ = v101
                                                        _ = v102
                                                    } else {
                                                        var v103 string = "\n\t"
                                                        _ = len(v2)
                                                        _ = v103
                                                    }
                                                    v5 = v5
                                                    _ = v100
                                                }
                                                if true {
                                                    v104 := GenS8{F0: 9223372036854775807}
                                                    v5 += (v5 / ((v5) | 1))
                                                    delete(v2, "")
                                                    _ = v104
                                                }
                                                var v105 int64 = 0
                                                if ((true || (!(!false))) || false) {
                                                    v55 -= v55
                                                    var v106 int8 = 1
                                                    _ = v106
                                                }
                                                var v107 int32 = -1
                                                _ = v105
                                                _ = v107
                                            }
                                            v108 := GenS2{F0: (("" + v46) + "foo bar")}
                                            _ = v108
                                        }
                                        if true {
                                            fmt.Println(v55, fi50)
                                            var v109 uint32 = 2
                                            fmt.Println(v3)
                                            var v110 int8 = 0
                                            _ = v109
                                            _ = v110
                                        }
                                    }
                                } else {
                                    var v111 float64 = 3.14159
                                    _ = v111
                                }
                            case 0:
                                switch -111 {
                                case 127:
                                    switch int(v8) {
                                    case 128:
                                        switch v46 {
                                        case "hello":
                                            if (0.0 > 0.0) {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                v5 -= v5
                                                if false {
                                                    if (!((true && ((!(false && true)) && true)) || (!(369 >= (v6 - v6))))) {
                                                        if len(v4) > 0 {
                                                            v4[0] = -985
                                                        }
                                                    } else {
                                                        v112 := GenS9{F0: "\n\t", F1: false, F2: (^(v6))}
                                                        _ = v112
                                                    }
                                                    v1 += v3
                                                    v113 := GenS1{F0: v46}
                                                    var v114 int = (fi50 % (((v29 - (v29 + v29))) | 1))
                                                    var v115 int = v114
                                                    _ = v113
                                                    _ = v114
                                                    _ = v115
                                                } else {
                                                    var v116 bool = ((v5 & -986) >= v5)
                                                    var v117 uint16 = uint16(v8)
                                                    switch (^(fi54)) {
                                                    case -9223372036854775808:
                                                        if false {
                                                            switch v7 {
                                                            case 556:
                                                                var v118 uint = (v6 >> ((116) & 63))
                                                                _ = v118
                                                            case 189:
                                                                switch v46 {
                                                                case "\n\t":
                                                                    switch v7 {
                                                                    case 32767:
                                                                        v119 := GenS9{F0: (v46 + (((("foo bar" + v46) + (("a" + v46) + "foo bar")) + v46) + "\n\t")), F1: (!v116), F2: 908}
                                                                        _ = v119
                                                                    case 0:
                                                                        var v120 uint32 = 0
                                                                        _ = v120
                                                                    case -450:
                                                                        delete(v2, v46)
                                                                    default:
                                                                        _ = len(v4)
                                                                    }
                                                                case "a":
                                                                    if ((!(!(!v116))) || (-305 != -9223372036854775808)) {
                                                                        v3 = v3
                                                                        v121 := GenS10{F0: -9223372036854775808, F1: (float64(1e10) >= 1.5), F2: 134}
                                                                        switch v7 {
                                                                        case 1:
                                                                            if (!(!false)) {
                                                                                if len(v4) >= 2 {
                                                                                    v4 = v4[1:len(v4)]
                                                                                }
                                                                                var v122 int64 = 184
                                                                                v5 -= (v5 | int16(v55))
                                                                                _ = v122
                                                                            } else {
                                                                                switch (v46 + "hello") {
                                                                                case "foo bar":
                                                                                    var v123 bool = (!((v5 * v5) == 856))
                                                                                    _ = v123
                                                                                case "\n\t":
                                                                                    var v124 float64 = float64(1.5)
                                                                                    _ = v124
                                                                                }
                                                                                if false {
                                                                                    _ = len(v4)
                                                                                    v125 := GenS1{F0: (v46 + (v46 + (v46 + (("hello" + (v46 + ((v46 + ((v46 + v46) + "foo bar")) + (v46 + "a")))) + v46))))}
                                                                                    _ = v125
                                                                                } else {
                                                                                    v126 := GenS11{F0: 651, F1: (v8 + v8), F2: v46}
                                                                                    switch 0 {
                                                                                    case 32768:
                                                                                        switch (v46 + v46) {
                                                                                        case "":
                                                                                            v5 += v5
                                                                                        case "hello":
                                                                                            _ = len(v4)
                                                                                        case "a":
                                                                                            var v127 int8 = -1
                                                                                            _ = v127
                                                                                        default:
                                                                                            v1 -= v117
                                                                                        }
                                                                                        v7 += 0
                                                                                    case -406:
                                                                                        v116 = false
                                                                                        var v128 float64 = (((float64((3.14159 - float32((0.0 * 1.5)))) / (-2.25 - 1.5)) / 1.5) - float64(((0.1 * ((1.5 / 3.14159) + -2.25)) / (1e10 / 0.1))))
                                                                                        _ = v128
                                                                                    case 128:
                                                                                        v129 := GenS8{F0: 137}
                                                                                        _ = v129
                                                                                    default:
                                                                                        var v130 uint8 = uint8(v6)
                                                                                        _ = v130
                                                                                    }
                                                                                    if true {
                                                                                        v29 -= fi54
                                                                                        var v131 float64 = (1e10 * float64((0.1 + 1e10)))
                                                                                        switch v46 {
                                                                                        case "a":
                                                                                            var v132 string = v46
                                                                                            _ = v132
                                                                                        }
                                                                                        _ = v131
                                                                                    } else {
                                                                                        var v133 int64 = 363
                                                                                        var v134 int16 = (^(v5))
                                                                                        var v135 uint = 884
                                                                                        _ = v133
                                                                                        _ = v134
                                                                                        _ = v135
                                                                                    }
                                                                                    _ = v126
                                                                                }
                                                                                switch "\n\t" {
                                                                                case "":
                                                                                    switch "foo bar" {
                                                                                    case "":
                                                                                        var v136 uint8 = uint8(v1)
                                                                                        _ = v136
                                                                                    case "\n\t":
                                                                                        var v137 int = 127
                                                                                        v138 := GenS4{F0: v46, F1: v8, F2: ("" + "")}
                                                                                        if v116 {
                                                                                            switch "foo bar" {
                                                                                            case "\n\t":
                                                                                                if (86 > -128) {
                                                                                                    fmt.Println(v1, v55, v6)
                                                                                                    var v139 int = int(v8)
                                                                                                    v140 := GenS9{F0: (v46 + v46), F1: v116, F2: (^(v6))}
                                                                                                    _ = v139
                                                                                                    _ = v140
                                                                                                }
                                                                                            case "hello":
                                                                                                if true {
                                                                                                    if (-1 < 0) {
                                                                                                        v46 = "a"
                                                                                                    }
                                                                                                    v4 = append(v4, fi54)
                                                                                                    v55 -= (v55 - v55)
                                                                                                }
                                                                                                var v141 uint32 = v8
                                                                                                var v142 uint16 = (v1 >> (((^(v6))) & 63))
                                                                                                _ = v141
                                                                                                _ = v142
                                                                                            default:
                                                                                                var v143 uint16 = v1
                                                                                                _ = v143
                                                                                            }
                                                                                        }
                                                                                        v2[(v46 + v46)] = (!v116)
                                                                                        _ = v137
                                                                                        _ = v138
                                                                                    default:
                                                                                        var v144 float64 = float64(float32(float64(3.14159)))
                                                                                        _ = v144
                                                                                    }
                                                                                case "\n\t":
                                                                                    var v145 int64 = 42
                                                                                    _ = v145
                                                                                }
                                                                            }
                                                                        }
                                                                        _ = v121
                                                                    }
                                                                case "hello":
                                                                    if (v6 < (^(v6))) {
                                                                        var v146 int32 = 42
                                                                        if len(v4) > 0 {
                                                                            v4[0] = (v29 / ((int(v8)) | 1))
                                                                        }
                                                                        _ = v146
                                                                    } else {
                                                                        var v147 int8 = int8(v3)
                                                                        _ = len(v2)
                                                                        v148 := v2["hello"]
                                                                        _ = v147
                                                                        _ = v148
                                                                    }
                                                                }
                                                            case 42:
                                                                v6 -= (^(v6))
                                                            }
                                                            var v149 float64 = 1e10
                                                            _ = v149
                                                        } else {
                                                            if (128 >= 128) {
                                                                var v150 int = (fi50 >> (((^(v6))) & 63))
                                                                _ = v150
                                                            }
                                                            var v151 string = v46
                                                            _ = v151
                                                        }
                                                    case 32767:
                                                        fmt.Println(v1, v117, v6)
                                                    case -661:
                                                        if (!(1 > 128)) {
                                                            if v116 {
                                                                v117 = v3
                                                                fmt.Println(v46)
                                                            } else {
                                                                fmt.Println(fi50, v6)
                                                            }
                                                        }
                                                    default:
                                                        var v152 float64 = 1.5
                                                        _ = v152
                                                    }
                                                    _ = v116
                                                    _ = v117
                                                }
                                            }
                                        }
                                    }
                                default:
                                    switch (v7 - -9223372036854775808) {
                                    case 128:
                                        fmt.Println(fi50)
                                    default:
                                        var v153 uint32 = 595
                                        _ = v153
                                    }
                                }
                            }
                            if (false || true) {
                                v46 = "hello"
                                v154 := GenS10{F0: 42, F1: true, F2: 110}
                                v155 := v2[("\n\t" + v46)]
                                _ = v154
                                _ = v155
                            } else {
                                v8 += 65535
                                var v156 uint16 = (v3 | (v1 % ((157) | 1)))
                                _ = v156
                            }
                            _ = fi54
                            _ = v55
                        }
                        fmt.Println(v7)
                    } else {
                        var v157 int64 = -936
                        if (1.5 >= (1e10 - ((1.5 + float64(1e10)) + 0.0))) {
                            for fi158 := 0; fi158 < 15; fi158++ {
                                fmt.Println(fi50, v29)
                                v1 = 65535
                                var v159 int16 = (^(v5))
                                _ = fi158
                                _ = v159
                            }
                            v2[("foo bar" + ((v46 + (v46 + "a")) + v46))] = true
                            v160 := v2[v46]
                            _ = v160
                        } else {
                            var v161 int16 = -1
                            var v162 uint64 = 630
                            var v163 int64 = 577
                            _ = v161
                            _ = v162
                            _ = v163
                        }
                        _ = v157
                    }
                }
            } else {
                v8 += (^(v8))
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                if false {
                    for fi164 := 0; fi164 < 5; fi164++ {
                        if (v7 < 127) {
                            v29 -= v7
                            v2[v46] = (!(float64(-2.25) > 1.5))
                            var v165 uint = v6
                            _ = v165
                        }
                        _ = fi164
                    }
                }
            }
            _ = fi50
        }
        v3 = v3
        var v166 uint8 = 56
        _ = v166
    }
    fmt.Println(v1, v8)
    v6 = 283
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
