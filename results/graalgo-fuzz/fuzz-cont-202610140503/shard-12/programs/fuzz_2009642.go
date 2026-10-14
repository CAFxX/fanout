package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int16; F1 int32; F2 uint }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int16; F1 string }
type GenS10 struct { F0 uint64; F1 string; F2 int8 }
type GenS11 struct { F0 uint; F1 int32; F2 int64 }
type GenS12 struct { F0 int32; F1 uint8 }
type GenS13 struct { F0 bool; F1 bool; F2 int8 }
type GenS14 struct { F0 uint; F1 uint }
type GenS15 struct { F0 int8 }

func main() {
    var v1 uint16 = 1
    v2 := make(map[int]int)
    var v3 float64 = 0.1
    v4 := make(map[string]int)
    fmt.Println(v3, v3)
    v5 := GenS1{F0: (v1 % ((32767) | 1))}
    fmt.Println(v3)
    if true {
        var v6 uint16 = v1
        v7 := GenS1{F0: (v6 | 0)}
        _ = v6
        _ = v7
    } else {
        var v8 uint8 = 228
        if (-2.25 == 1.5) {
            v8 -= 170
            fmt.Println(v3, v3, v8)
        } else {
            if (!(44 > 492)) {
                v9 := []int{203, 4294967295, 32767, 266}
                _ = v9
            }
            if ((^(v8)) <= 2) {
                v10 := GenS1{F0: 0}
                for fi11 := 0; fi11 < 13; fi11++ {
                    var v12 int16 = 0
                    _ = fi11
                    _ = v12
                }
                var v13 float64 = (-2.25 / (((float64(0.1) + v3) - 3.14159) / (v3 / 0.1)))
                _ = v10
                _ = v13
            } else {
                delete(v2, 127)
            }
        }
        _ = v8
    }
    v14 := GenS1{F0: v1}
    switch 4294967295 {
    case 634:
        for fi15 := 0; fi15 < 1; fi15++ {
            for fi16 := 0; fi16 < 20; fi16++ {
                v1 = v1
                _ = fi16
            }
            v4[(("\n\t" + "a") + ("foo bar" + ("hello" + "\n\t")))] = (fi15 * (fi15 >> ((4294967295) & 63)))
            var v17 int8 = 1
            _ = fi15
            _ = v17
        }
    case 2147483647:
        _ = len(v4)
    case 9223372036854775807:
        if (18446744073709551615 != uint64(v1)) {
            var v18 string = ("\n\t" + "a")
            v19 := []int{373, -2147483648, -1, 445}
            _ = v18
            _ = v19
        } else {
            if false {
                var v20 uint = 112
                v1 -= (^(v1))
                _ = v20
            }
            fmt.Println(v3, v3)
            v21 := []int{42, -658, 0}
            _ = v21
        }
    }
    v3 = -2.25
    fmt.Println(v3)
    switch -1 {
    case -102:
        var v22 bool = true
        _ = v22
    case 2147483647:
        _ = len(v4)
    case 4294967295:
        v3 += v3
    default:
        var v23 int16 = 1000
        _ = v23
    }
    v3 += (v3 - v3)
    switch 32768 {
    case -226:
        v24 := []int{}
        _ = v24
    case 32767:
        switch 146 {
        case 32768:
            var v25 int16 = -105
            _ = v25
        default:
            if false {
                fmt.Println(v3, v3)
            }
        }
    case 0:
        for fi26 := 0; fi26 < 11; fi26++ {
            switch ("a" + "") {
            case "a":
                fmt.Println(v3, fi26, fi26)
            case "\n\t":
                if (42 <= 42) {
                    var v27 int32 = 0
                    _ = v27
                } else {
                    if true {
                        switch "hello" {
                        case "foo bar":
                            if (false || (42 > 0)) {
                                switch (fi26 ^ (fi26 ^ (fi26 >> ((4294967295) & 63)))) {
                                case 1:
                                    var v28 int = (^(fi26))
                                    _ = v28
                                case -41:
                                    if false {
                                        for fi29 := 0; fi29 < 13; fi29++ {
                                            if ((float32(0.1) * 1e10) != 1.5) {
                                                var v30 bool = false
                                                var v31 uint8 = 128
                                                _ = v30
                                                _ = v31
                                            } else {
                                                var v32 float32 = 0.0
                                                v3 = ((0.1 - v3) - 0.0)
                                                v32 = (v32 / float32(1.5))
                                                _ = v32
                                            }
                                            if true {
                                                if (!false) {
                                                    v1 -= v1
                                                    v3 = (0.0 + v3)
                                                    fmt.Println(fi29, v1, fi26)
                                                }
                                            }
                                            if false {
                                                v33 := GenS1{F0: (v1 & (v1 ^ v1))}
                                                fmt.Println(v1, v1, v3)
                                                var v34 uint = 36
                                                var v35 uint16 = v1
                                                if (false || true) {
                                                    v36 := GenS2{F0: 42, F1: 255}
                                                    _ = v36
                                                }
                                                fmt.Println(v34, fi26, v35)
                                                _ = v33
                                                _ = v34
                                                _ = v35
                                            } else {
                                                if (fi26 < fi29) {
                                                    if true {
                                                        var v37 float32 = 0.1
                                                        v1 = 246
                                                        _ = v37
                                                    } else {
                                                        switch fi26 {
                                                        case -1:
                                                            v38 := GenS3{F0: v1, F1: 639}
                                                            _ = v38
                                                        default:
                                                            fmt.Println(fi29, v3)
                                                        }
                                                    }
                                                    var v39 uint32 = 65535
                                                    _ = v39
                                                }
                                            }
                                            _ = fi29
                                        }
                                        var v40 int8 = 65
                                        _ = v40
                                    } else {
                                        for fi41 := 0; fi41 < 11; fi41++ {
                                            if false {
                                                fmt.Println(fi26)
                                                if ("a" != "a") {
                                                    var v42 float64 = (v3 + 3.14159)
                                                    _ = v42
                                                } else {
                                                    v43 := v4[(("hello" + "\n\t") + ("a" + "hello"))]
                                                    var v44 uint64 = 4294967295
                                                    _ = v43
                                                    _ = v44
                                                }
                                                var v45 int64 = 892
                                                _ = v45
                                            }
                                            _ = fi41
                                        }
                                        switch ((("foo bar" + "") + "") + "") {
                                        case "foo bar":
                                            v3 -= (float64(0.0) / 1.5)
                                        case "":
                                            v46 := GenS1{F0: (^(v1))}
                                            _ = v46
                                        default:
                                            for fi47 := 0; fi47 < 6; fi47++ {
                                                v48 := GenS3{F0: v1, F1: 255}
                                                if (-1 <= -725) {
                                                    v3 = -2.25
                                                }
                                                v1 = 32767
                                                _ = fi47
                                                _ = v48
                                            }
                                        }
                                    }
                                case -9223372036854775808:
                                    v3 += (v3 / (float64(1e10) * 1.5))
                                default:
                                    switch fi26 {
                                    case 4294967295:
                                        fmt.Println(v3, fi26, v1)
                                    case -233:
                                        if (-32768 <= 1) {
                                            switch "" {
                                            case "foo bar":
                                                for fi49 := 0; fi49 < 8; fi49++ {
                                                    v50 := GenS4{F0: 127}
                                                    v51 := GenS5{F0: 42}
                                                    _ = fi49
                                                    _ = v50
                                                    _ = v51
                                                }
                                            case "a":
                                                fmt.Println(fi26, fi26, v1)
                                            case "hello":
                                                var v52 uint = 931
                                                _ = v52
                                            default:
                                                switch (fi26 / (((-(fi26))) | 1)) {
                                                case 32768:
                                                    if (!(!(((-1 < int32(fi26)) || (-54 >= 14)) && (65535 != 0)))) {
                                                        switch fi26 {
                                                        case 9223372036854775807:
                                                            v1 -= (v1 & 32767)
                                                        case -9223372036854775808:
                                                            v53 := GenS2{F0: 1, F1: 18446744073709551615}
                                                            _ = v53
                                                        default:
                                                            v54 := GenS6{F0: 0}
                                                            _ = v54
                                                        }
                                                        var v55 int64 = 532
                                                        _ = v55
                                                    }
                                                case 0:
                                                    var v56 int32 = 1
                                                    _ = v56
                                                default:
                                                    var v57 int32 = 42
                                                    _ = v57
                                                }
                                            }
                                            v58 := []int{128, fi26}
                                            _ = v58
                                        } else {
                                            if false {
                                                if (540 >= 42) {
                                                    v1 += 255
                                                }
                                                v1 += 42
                                            } else {
                                                if (484 >= 10) {
                                                    v59 := GenS3{F0: 255, F1: 4294967295}
                                                    _ = v59
                                                } else {
                                                    for fi60 := 0; fi60 < 2; fi60++ {
                                                        fmt.Println(fi26, fi60)
                                                        var v61 int64 = 9223372036854775807
                                                        _ = fi60
                                                        _ = v61
                                                    }
                                                    var v62 uint8 = 186
                                                    _ = v62
                                                }
                                                for fi63 := 0; fi63 < 16; fi63++ {
                                                    var v64 bool = true
                                                    _ = fi63
                                                    _ = v64
                                                }
                                            }
                                            if (1 < 4294967295) {
                                                if ((true && (false && false)) && (false && (-2.25 < -2.25))) {
                                                    v1 = (v1 * (v1 << ((4294967295) & 63)))
                                                    for fi65 := 0; fi65 < 9; fi65++ {
                                                        var v66 int = fi65
                                                        _ = fi65
                                                        _ = v66
                                                    }
                                                    v67 := GenS3{F0: 235, F1: 757}
                                                    _ = v67
                                                }
                                                var v68 uint = 18446744073709551615
                                                _ = v68
                                            } else {
                                                v3 += (0.0 / v3)
                                                v69 := GenS5{F0: 9223372036854775807}
                                                fmt.Println(v1, v1, fi26)
                                                _ = v69
                                            }
                                            v3 += 0.0
                                        }
                                    }
                                }
                                v1 += (v1 | v1)
                                if (!(189 < 0)) {
                                    v3 -= float64(1.5)
                                    v70 := GenS1{F0: v1}
                                    var v71 int = 32767
                                    _ = v70
                                    _ = v71
                                }
                            } else {
                                var v72 uint32 = 42
                                _ = v72
                            }
                        }
                        v1 = v1
                    } else {
                        for fi73 := 0; fi73 < 1; fi73++ {
                            v1 = 2
                            _ = fi73
                        }
                        if false {
                            v3 -= (v3 / (v3 - v3))
                        } else {
                            if (!false) {
                                for fi74 := 0; fi74 < 10; fi74++ {
                                    fmt.Println(fi26, v1)
                                    _ = fi74
                                }
                                if (2 > 2) {
                                    v1 = v1
                                    for fi75 := 0; fi75 < 17; fi75++ {
                                        if (-2147483648 > fi26) {
                                            switch fi75 {
                                            case 0:
                                                var v76 string = "foo bar"
                                                _ = v76
                                            case 42:
                                                var v77 uint32 = 2
                                                _ = v77
                                            default:
                                                if (!(v1 != 2)) {
                                                    v78 := GenS7{F0: -32768, F1: -1, F2: 0}
                                                    v79 := []int{fi75, fi26}
                                                    var v80 int64 = 42
                                                    _ = v78
                                                    _ = v79
                                                    _ = v80
                                                } else {
                                                    _ = len(v2)
                                                    var v81 bool = false
                                                    _ = v81
                                                }
                                            }
                                            var v82 uint64 = 2
                                            if false {
                                                if false {
                                                    _ = len(v4)
                                                    var v83 uint64 = 219
                                                    if (true || true) {
                                                        v82 += (v82 / ((536) | 1))
                                                        var v84 uint64 = v83
                                                        switch (("\n\t" + "foo bar") + "hello") {
                                                        case "a":
                                                            v84 = v83
                                                        case "hello":
                                                            if false {
                                                                v85 := []int{}
                                                                if false {
                                                                    if false {
                                                                        v84 -= 167
                                                                        switch fi26 {
                                                                        case 4294967295:
                                                                            v83 -= (^(v83))
                                                                        case -2147483648:
                                                                            fmt.Println(v83)
                                                                        case -9223372036854775808:
                                                                            var v86 string = "hello"
                                                                            _ = v86
                                                                        default:
                                                                            var v87 int32 = 0
                                                                            v88 := GenS2{F0: 255, F1: 1}
                                                                            v83 += v84
                                                                            if true {
                                                                                v89 := GenS8{F0: "a"}
                                                                                _ = v89
                                                                            }
                                                                            if ((!false) || (true || (!false))) {
                                                                                var v90 int8 = 127
                                                                                var v91 int16 = 32767
                                                                                _ = v90
                                                                                _ = v91
                                                                            } else {
                                                                                v83 -= v83
                                                                                if (false || false) {
                                                                                    if len(v85) >= 2 {
                                                                                        v85 = v85[1:len(v85)]
                                                                                    }
                                                                                } else {
                                                                                    v83 = v82
                                                                                    v1 += (v1 * (v1 | (^(v1))))
                                                                                    fmt.Println(v84)
                                                                                }
                                                                                fmt.Println(fi26)
                                                                            }
                                                                            _ = v87
                                                                            _ = v88
                                                                        }
                                                                    }
                                                                    v3 = v3
                                                                    fmt.Println(v3, v84, v84)
                                                                }
                                                                v92 := GenS9{F0: -32768, F1: "\n\t"}
                                                                _ = v85
                                                                _ = v92
                                                            } else {
                                                                var v93 float64 = float64(1.5)
                                                                v94 := GenS10{F0: 2, F1: "a", F2: -1}
                                                                if (4294967295 <= 65535) {
                                                                    v95 := []int{fi26, (^(fi26)), fi75, (fi75 & (^(fi26)))}
                                                                    var v96 uint64 = v82
                                                                    v97 := GenS4{F0: 127}
                                                                    v3 = (v93 + v3)
                                                                    v98 := GenS3{F0: (v1 & 274), F1: 1}
                                                                    v1 -= v1
                                                                    fmt.Println(v3)
                                                                    _ = v95
                                                                    _ = v96
                                                                    _ = v97
                                                                    _ = v98
                                                                } else {
                                                                    v82 += 0
                                                                    if ("hello" == "hello") {
                                                                        v99 := []int{}
                                                                        _ = v99
                                                                    }
                                                                }
                                                                v100 := GenS11{F0: 2, F1: 482, F2: 0}
                                                                switch (fi75 | 2147483647) {
                                                                case 127:
                                                                    v101 := []int{42, 1}
                                                                    _ = v101
                                                                case 2147483647:
                                                                    var v102 uint16 = v1
                                                                    _ = v102
                                                                case -859:
                                                                    fmt.Println(v84, v82)
                                                                default:
                                                                    var v103 float32 = ((1e10 + 1.5) + (3.14159 * 0.0))
                                                                    _ = v103
                                                                }
                                                                var v104 uint = 255
                                                                _ = v100
                                                                _ = v104
                                                                _ = v93
                                                                _ = v94
                                                            }
                                                        default:
                                                            var v105 int8 = 100
                                                            _ = v105
                                                        }
                                                        _ = v84
                                                    }
                                                    _ = v83
                                                }
                                            }
                                            _ = v82
                                        } else {
                                            v106 := GenS3{F0: 361, F1: 65535}
                                            var v107 int = (fi75 << ((116) & 63))
                                            v107 -= (v107 | -1)
                                            _ = v106
                                            _ = v107
                                        }
                                        if (!(32767 <= -32768)) {
                                            v4["foo bar"] = (fi26 - fi75)
                                        }
                                        if (42 == 42) {
                                            var v108 bool = false
                                            var v109 int = fi75
                                            _ = v108
                                            _ = v109
                                        } else {
                                            if (true && false) {
                                                v110 := GenS10{F0: 4294967295, F1: "\n\t", F2: -43}
                                                var v111 int64 = 42
                                                _ = len(v2)
                                                v3 += -2.25
                                                if (88 < 1) {
                                                    v112 := GenS12{F0: 215, F1: 127}
                                                    v113 := []int{(fi75 & (fi26 / ((fi26) | 1))), fi75, 32768}
                                                    v114 := GenS13{F0: true, F1: false, F2: -128}
                                                    _ = v112
                                                    _ = v113
                                                    _ = v114
                                                }
                                                _ = v110
                                                _ = v111
                                            }
                                            if (true || (!false)) {
                                                if ("\n\t" < ("hello" + "a")) {
                                                    fmt.Println(v1, fi75, v3)
                                                    if (false || true) {
                                                        var v115 uint8 = 42
                                                        switch (fi75 / (((fi26 * (fi75 | 48))) | 1)) {
                                                        case 127:
                                                            if false {
                                                                v1 -= (^(v1))
                                                            } else {
                                                                v115 = (^(v115))
                                                            }
                                                        case 0:
                                                            v1 -= (v1 * (v1 | v1))
                                                        case 489:
                                                            var v116 int8 = 42
                                                            _ = v116
                                                        default:
                                                            fmt.Println(fi75)
                                                        }
                                                        v115 -= 128
                                                        switch 128 {
                                                        case 42:
                                                            var v117 uint16 = (^(v1))
                                                            _ = v117
                                                        case 128:
                                                            v115 -= v115
                                                        }
                                                        _ = v115
                                                    } else {
                                                        if ((!true) && (42 == 4294967295)) {
                                                            v118 := GenS1{F0: v1}
                                                            _ = v118
                                                        }
                                                        var v119 float64 = v3
                                                        _ = v119
                                                    }
                                                    if true {
                                                        v3 -= (1e10 + (v3 / float64(float32((0.0 - (float64(-2.25) * (v3 / ((-2.25 + v3) * v3))))))))
                                                        switch (fi75 ^ fi26) {
                                                        case -2147483648:
                                                            switch fi26 {
                                                            case 1:
                                                                v120 := []int{(-(fi26))}
                                                                _ = v120
                                                            case 128:
                                                                var v121 float32 = ((float32((v3 * (float64(0.1) / 1.0))) / 1.0) - (0.1 / 1.0))
                                                                _ = v121
                                                            }
                                                        case 129:
                                                            var v122 uint = 255
                                                            _ = v122
                                                        case 716:
                                                            v123 := GenS11{F0: 4294967295, F1: 673, F2: 0}
                                                            _ = v123
                                                        }
                                                        delete(v4, (("" + "\n\t") + ""))
                                                    }
                                                } else {
                                                    if (-1 <= -285) {
                                                        if true {
                                                            v3 = (v3 * v3)
                                                            switch fi75 {
                                                            case 42:
                                                                fmt.Println(v3)
                                                                fmt.Println(fi75, v1)
                                                                var v124 uint = 18446744073709551615
                                                                switch fi26 {
                                                                case 32767:
                                                                    switch "\n\t" {
                                                                    case "hello":
                                                                        var v125 uint16 = (v1 ^ 65535)
                                                                        _ = v125
                                                                    case "":
                                                                        v126 := GenS12{F0: 211, F1: 1}
                                                                        _ = v126
                                                                    case "foo bar":
                                                                        var v127 int8 = 42
                                                                        _ = v127
                                                                    }
                                                                default:
                                                                    var v128 int32 = 2147483647
                                                                    _ = v128
                                                                }
                                                                v129 := []int{0}
                                                                _ = v124
                                                                _ = v129
                                                            case 488:
                                                                switch ("\n\t" + ((("\n\t" + ("a" + "hello")) + "a") + "\n\t")) {
                                                                case "hello":
                                                                    var v130 bool = (!true)
                                                                    _ = v130
                                                                case "\n\t":
                                                                    fmt.Println(fi26, fi26, v1)
                                                                case "a":
                                                                    v1 = v1
                                                                }
                                                            default:
                                                                delete(v2, fi26)
                                                            }
                                                        } else {
                                                            v1 = v1
                                                        }
                                                        v1 = v1
                                                    }
                                                    switch -2147483648 {
                                                    case 942:
                                                        switch 32768 {
                                                        case 2147483647:
                                                            v131 := GenS14{F0: 307, F1: 18446744073709551615}
                                                            _ = v131
                                                        case 127:
                                                            fmt.Println(v1, v3, fi26)
                                                        case 590:
                                                            v132 := v2[127]
                                                            _ = v132
                                                        default:
                                                            var v133 bool = true
                                                            _ = v133
                                                        }
                                                    case 128:
                                                        v134 := []int{(fi26 + 32767), 9223372036854775807}
                                                        _ = v134
                                                    }
                                                    delete(v2, fi75)
                                                }
                                                fmt.Println(v3)
                                                v135 := GenS15{F0: -67}
                                                _ = v135
                                            } else {
                                                v3 = v3
                                                var v136 uint32 = 2147483647
                                                _ = v136
                                            }
                                        }
                                        _ = fi75
                                    }
                                    var v137 uint64 = uint64(v1)
                                    _ = v137
                                } else {
                                    var v138 float32 = 0.1
                                    v3 += float64(v138)
                                    v3 += (-2.25 * v3)
                                    _ = v138
                                }
                                for fi139 := 0; fi139 < 4; fi139++ {
                                    fmt.Println(v1, fi26)
                                    if (v3 <= ((v3 + 1e10) + 1.5)) {
                                        v140 := []int{(fi26 - 32768), (fi139 - 4294967295)}
                                        _ = v140
                                    } else {
                                        var v141 int64 = 42
                                        _ = v141
                                    }
                                    if true {
                                        fmt.Println(v1)
                                        switch fi26 {
                                        case 2147483647:
                                            switch (("hello" + "hello") + "") {
                                            case "hello":
                                                var v142 int = 200
                                                _ = v142
                                            case "\n\t":
                                                var v143 bool = ("a" >= ("hello" + "hello"))
                                                _ = v143
                                            default:
                                                fmt.Println(fi139)
                                            }
                                        case -1:
                                            v144 := GenS10{F0: 323, F1: "foo bar", F2: -128}
                                            _ = v144
                                        }
                                    }
                                    _ = fi139
                                }
                            }
                        }
                    }
                    v1 -= 42
                    var v145 uint8 = 255
                    _ = v145
                }
            }
            v1 += v1
            if true {
                switch ("hello" + ("foo bar" + "hello")) {
                case "hello":
                    var v146 int = (-(fi26))
                    _ = v146
                case "foo bar":
                    if true {
                        var v147 int32 = 0
                        _ = v147
                    } else {
                        var v148 uint = 18446744073709551615
                        _ = v148
                    }
                default:
                    v3 = (v3 - (v3 * float64((3.14159 / (3.14159 * 0.1)))))
                }
            }
            _ = fi26
        }
    }
    var v149 bool = (2 <= 1)
    for fi150 := 0; fi150 < 2; fi150++ {
        v149 = (18446744073709551615 > 18446744073709551615)
        fmt.Println(fi150)
        v3 += float64(0.0)
        _ = fi150
    }
    delete(v4, "")
    fmt.Println(v149, v149)
    fmt.Println("v1:", v1)
    fmt.Println("v149:", v149)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v149
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
