package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 int32; F1 string }
type GenS5 struct { F0 bool; F1 string; F2 uint16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint64; F1 uint16; F2 bool }
type GenS9 struct { F0 uint8; F1 uint; F2 int }
type GenS10 struct { F0 string; F1 uint64 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 string; F1 uint }
type GenS13 struct { F0 bool; F1 int }
type GenS14 struct { F0 bool; F1 uint }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 int; F1 string }
type GenS17 struct { F0 bool; F1 int; F2 uint }
type GenS18 struct { F0 uint }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 string; F1 int }
func genFn0(p0 int, p1 uint32, p2 uint16) int16 {
    v4 := GenS1{F0: int(p1), F1: 42, F2: ""}
    if (-2.25 > float32(0.1)) {
        v5 := []int{p0, p0, p0, -1}
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return -1
}
func genFn1() int16 {
    switch ("a" + "foo bar") {
    case "\n\t":
        var v6 int16 = genFn0(9223372036854775807, 1, 65535)
        _ = v6
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: (fi7 - 0), F1: 1, F2: "\n\t"}
        var v9 string = ((("" + "\n\t") + (("hello" + "hello") + "foo bar")) + "")
        fmt.Println(fi7)
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 127, F1: 1000, F2: "\n\t"}
    var v11 int32 = -569
    if false {
        if ((1.5 / 3.14159) >= -2.25) {
            fmt.Println(v11)
            switch 0 {
            case -2147483648:
                v11 += v11
            default:
                for fi12 := 0; fi12 < 17; fi12++ {
                    if true {
                        for fi13 := 0; fi13 < 4; fi13++ {
                            v11 = -2147483648
                            var v14 int64 = 9223372036854775807
                            if false {
                                if (0.1 <= 3.14159) {
                                    if ((0.1 > (3.14159 + (((1e10 / 3.14159) / 0.1) + (1e10 + float64(1.5))))) || (-32768 > 42)) {
                                        if (42 == -54) {
                                            v14 -= 9223372036854775807
                                            fmt.Println(v11, fi12, v11)
                                            var v15 int32 = v11
                                            if false {
                                                var v16 int8 = 42
                                                v11 = (v11 << ((4294967295) & 63))
                                                switch "foo bar" {
                                                case "\n\t":
                                                    fmt.Println(fi12, v11)
                                                case "hello":
                                                    switch "hello" {
                                                    case "hello":
                                                        var v17 int8 = v16
                                                        _ = v17
                                                    }
                                                default:
                                                    fmt.Println(v15, v16, v14)
                                                }
                                                _ = v16
                                            } else {
                                                var v18 string = "a"
                                                v19 := GenS2{F0: false}
                                                v20 := []int{fi12}
                                                _ = v18
                                                _ = v19
                                                _ = v20
                                            }
                                            var v21 uint8 = 10
                                            v14 += v14
                                            v22 := GenS1{F0: (fi13 & 1), F1: -32768, F2: ((("" + ("a" + (("foo bar" + "\n\t") + ""))) + ((("hello" + "\n\t") + ("" + "foo bar")) + ("hello" + "\n\t"))) + ("foo bar" + ""))}
                                            _ = v15
                                            _ = v21
                                            _ = v22
                                        }
                                    }
                                    v14 = (v14 % ((v14) | 1))
                                    fmt.Println(v11)
                                }
                            } else {
                                fmt.Println(fi13)
                                if false {
                                    v11 += v11
                                    v14 = (v14 & -983)
                                    var v23 uint8 = 128
                                    if ((-(v14)) != (v14 << ((uint(v23)) & 63))) {
                                        v24 := GenS1{F0: (fi13 % (((fi13 >> ((196) & 63))) | 1)), F1: 0, F2: ("hello" + "")}
                                        var v25 int64 = v14
                                        _ = v24
                                        _ = v25
                                    } else {
                                        if false {
                                            v23 = v23
                                            switch "hello" {
                                            case "\n\t":
                                                switch (^(fi12)) {
                                                case 42:
                                                    var v26 string = "foo bar"
                                                    _ = v26
                                                case 674:
                                                    v27 := GenS3{F0: 18446744073709551615, F1: 98}
                                                    _ = v27
                                                case -9223372036854775808:
                                                    v28 := GenS1{F0: -802, F1: 0, F2: (("\n\t" + (("hello" + "") + "foo bar")) + "a")}
                                                    _ = v28
                                                }
                                            case "a":
                                                switch fi12 {
                                                case 1:
                                                    switch "a" {
                                                    case "foo bar":
                                                        fmt.Println(fi12, v11, fi13)
                                                    case "\n\t":
                                                        if true {
                                                            switch (fi12 & fi13) {
                                                            case 2147483647:
                                                                switch "" {
                                                                case "foo bar":
                                                                    if false {
                                                                        if ((653 <= 735) || true) {
                                                                            v11 += (v11 << ((255) & 63))
                                                                        }
                                                                        v29 := GenS3{F0: 2, F1: -128}
                                                                        fmt.Println(v11)
                                                                        _ = v29
                                                                    }
                                                                case "\n\t":
                                                                    v30 := []int{(-(fi12)), 596, fi13, (fi12 % (((^(fi12))) | 1))}
                                                                    v11 += (v11 * -2147483648)
                                                                    v11 = (v11 & (^(v11)))
                                                                    _ = v30
                                                                }
                                                            case 127:
                                                                var v31 bool = false
                                                                _ = v31
                                                            case -9223372036854775808:
                                                                v23 += (v23 * (v23 & (^(v23))))
                                                            default:
                                                                fmt.Println(fi12, fi12, v11)
                                                            }
                                                            if (v11 >= v11) {
                                                                if (1 != -1) {
                                                                    v32 := GenS3{F0: 18446744073709551615, F1: 1}
                                                                    _ = v32
                                                                } else {
                                                                    v11 -= v11
                                                                    fmt.Println(v14, v11, v14)
                                                                }
                                                            }
                                                        }
                                                    case "a":
                                                        if ("hello" == "\n\t") {
                                                            v33 := GenS2{F0: (!((false && true) && true))}
                                                            if (!false) {
                                                                v34 := []int{fi13, (fi13 ^ fi12)}
                                                                var v35 int64 = v14
                                                                _ = v34
                                                                _ = v35
                                                            } else {
                                                                v14 = -944
                                                                var v36 float32 = (1e10 * (1e10 * 1e10))
                                                                switch fi13 {
                                                                case -519:
                                                                    v23 = (v23 / ((v23) | 1))
                                                                default:
                                                                    v37 := []int{}
                                                                    _ = v37
                                                                }
                                                                _ = v36
                                                            }
                                                            _ = v33
                                                        }
                                                    }
                                                case 32768:
                                                    v38 := GenS1{F0: 803, F1: 225, F2: ""}
                                                    _ = v38
                                                }
                                            case "hello":
                                                var v39 int8 = 100
                                                _ = v39
                                            default:
                                                var v40 uint64 = 4294967295
                                                _ = v40
                                            }
                                        }
                                        fmt.Println(v11)
                                        v23 = 128
                                        var v41 float32 = (float32((0.0 / (1.5 - 0.0))) * 3.14159)
                                        var v42 int64 = 9223372036854775807
                                        switch (^(fi13)) {
                                        case 0:
                                            v43 := GenS4{F0: v11, F1: "\n\t"}
                                            _ = v43
                                        default:
                                            v44 := GenS5{F0: true, F1: "\n\t", F2: 871}
                                            _ = v44
                                        }
                                        v45 := []int{(^(fi13)), (fi13 / ((fi12) | 1)), (fi13 + (-(fi13)))}
                                        _ = v41
                                        _ = v42
                                        _ = v45
                                    }
                                    _ = v23
                                }
                            }
                            v46 := []int{(fi13 - -483), (-(fi13)), int(v11)}
                            v47 := GenS6{F0: (!true)}
                            _ = fi13
                            _ = v14
                            _ = v46
                            _ = v47
                        }
                        if true {
                            v11 += (v11 ^ 308)
                        } else {
                            switch fi12 {
                            case -9223372036854775808:
                                v48 := GenS6{F0: false}
                                _ = v48
                            case 9223372036854775807:
                                for fi49 := 0; fi49 < 11; fi49++ {
                                    v50 := []int{9223372036854775807, 32768}
                                    var v51 string = "a"
                                    _ = fi49
                                    _ = v50
                                    _ = v51
                                }
                            case 32767:
                                var v52 int16 = -195
                                _ = v52
                            }
                            switch fi12 {
                            case 0:
                                if (!((true && true) && true)) {
                                    var v53 uint32 = 2147483647
                                    v53 = 948
                                    _ = v53
                                } else {
                                    var v54 uint8 = 0
                                    v55 := GenS3{F0: 4294967295, F1: -1}
                                    _ = v54
                                    _ = v55
                                }
                            default:
                                switch ("" + "foo bar") {
                                case "foo bar":
                                    v56 := GenS7{F0: 1}
                                    _ = v56
                                case "":
                                    if false {
                                        for fi57 := 0; fi57 < 17; fi57++ {
                                            fmt.Println(v11, fi12, fi57)
                                            v58 := GenS2{F0: (11 < 0)}
                                            v59 := GenS6{F0: false}
                                            _ = fi57
                                            _ = v58
                                            _ = v59
                                        }
                                        switch ("hello" + "\n\t") {
                                        case "":
                                            if ((false && ("hello" >= ((("hello" + ("\n\t" + ("foo bar" + "hello"))) + ("foo bar" + "hello")) + ((("\n\t" + (("" + "") + "a")) + "\n\t") + "")))) || (128 != 24)) {
                                                for fi60 := 0; fi60 < 12; fi60++ {
                                                    switch "hello" {
                                                    case "":
                                                        var v61 float32 = 0.0
                                                        _ = v61
                                                    case "\n\t":
                                                        fmt.Println(v11, fi60)
                                                    }
                                                    _ = fi60
                                                }
                                                v11 = v11
                                            } else {
                                                for fi62 := 0; fi62 < 3; fi62++ {
                                                    v63 := GenS8{F0: 455, F1: 255, F2: (!((0.0 - (3.14159 - (1e10 + -2.25))) > (0.0 + -2.25)))}
                                                    switch -2147483648 {
                                                    case -9223372036854775808:
                                                        v64 := GenS9{F0: 2, F1: 1, F2: -9223372036854775808}
                                                        _ = v64
                                                    case 808:
                                                        v65 := []int{(^(fi12))}
                                                        _ = v65
                                                    default:
                                                        v66 := GenS10{F0: "", F1: 1}
                                                        _ = v66
                                                    }
                                                    if true {
                                                        var v67 int16 = -1
                                                        switch "foo bar" {
                                                        case "":
                                                            var v68 int32 = (v11 << ((2) & 63))
                                                            _ = v68
                                                        case "foo bar":
                                                            fmt.Println(fi62)
                                                        default:
                                                            var v69 int32 = v11
                                                            _ = v69
                                                        }
                                                        v70 := []int{-2147483648}
                                                        _ = v67
                                                        _ = v70
                                                    } else {
                                                        var v71 int32 = (v11 + -374)
                                                        var v72 uint = 2
                                                        _ = v71
                                                        _ = v72
                                                    }
                                                    _ = fi62
                                                    _ = v63
                                                }
                                            }
                                        case "foo bar":
                                            var v73 uint8 = 247
                                            _ = v73
                                        case "\n\t":
                                            switch 1 {
                                            case -80:
                                                var v74 uint16 = 127
                                                _ = v74
                                            case -302:
                                                for fi75 := 0; fi75 < 14; fi75++ {
                                                    switch (fi75 - fi75) {
                                                    case 32767:
                                                        v76 := GenS11{F0: 607}
                                                        if true {
                                                            if true {
                                                                v77 := []int{fi12, fi75, (fi12 ^ 653)}
                                                                _ = v77
                                                            }
                                                            switch (fi75 % ((1) | 1)) {
                                                            case 92:
                                                                if (fi75 != -90) {
                                                                    v78 := GenS12{F0: ("\n\t" + "foo bar"), F1: 582}
                                                                    v79 := []int{(^(fi75)), (fi12 & 4294967295), 127, (fi12 & (fi12 - (fi75 ^ (-(fi75)))))}
                                                                    v11 += v11
                                                                    _ = v78
                                                                    _ = v79
                                                                } else {
                                                                    v80 := GenS13{F0: true, F1: fi12}
                                                                    if (!((fi75 & -419) >= (fi75 ^ fi12))) {
                                                                        v81 := GenS7{F0: -32768}
                                                                        fmt.Println(fi12)
                                                                        _ = v81
                                                                    }
                                                                    fmt.Println(fi75, fi75)
                                                                    _ = v80
                                                                }
                                                                if (!true) {
                                                                    var v82 uint8 = 245
                                                                    _ = v82
                                                                } else {
                                                                    v83 := []int{fi75, 2147483647, fi75, 128}
                                                                    v11 -= (v11 & v11)
                                                                    v84 := GenS7{F0: genFn0(fi12, 125, 262)}
                                                                    _ = v83
                                                                    _ = v84
                                                                }
                                                                var v85 bool = (!(653 != 4294967295))
                                                                _ = v85
                                                            case -665:
                                                                var v86 int = (fi75 + -800)
                                                                _ = v86
                                                            case 2147483647:
                                                                if false {
                                                                    v11 = int32(fi75)
                                                                }
                                                            }
                                                            v11 = (v11 % (((v11 % ((v11) | 1))) | 1))
                                                        } else {
                                                            v11 += v11
                                                        }
                                                        _ = v76
                                                    case 32768:
                                                        v11 -= -521
                                                    case -1:
                                                        var v87 int16 = 0
                                                        _ = v87
                                                    }
                                                    var v88 uint32 = 65535
                                                    _ = fi75
                                                    _ = v88
                                                }
                                            default:
                                                var v89 uint32 = 4294967295
                                                _ = v89
                                            }
                                        default:
                                            var v90 uint16 = 255
                                            _ = v90
                                        }
                                        v91 := []int{}
                                        _ = v91
                                    }
                                default:
                                    if ((float64(0.0) / 1.5) < 3.14159) {
                                        switch fi12 {
                                        case -2147483648:
                                            if true {
                                                v11 += -2147483648
                                                v11 = (v11 * v11)
                                                for fi92 := 0; fi92 < 17; fi92++ {
                                                    var v93 uint = 42
                                                    if (!(827 < -1)) {
                                                        fmt.Println(fi92, fi92)
                                                    } else {
                                                        v94 := GenS14{F0: false, F1: v93}
                                                        v95 := []int{}
                                                        _ = v94
                                                        _ = v95
                                                    }
                                                    fmt.Println(v11, v93, v93)
                                                    switch (fi92 ^ fi12) {
                                                    case 42:
                                                        switch fi92 {
                                                        case -584:
                                                            v96 := []int{(^(fi12)), (fi12 - -1), fi12, (fi12 | fi12)}
                                                            _ = v96
                                                        }
                                                    case -967:
                                                        v93 -= 489
                                                    }
                                                    fmt.Println(v93, v11, fi12)
                                                    v93 = (v93 * 255)
                                                    _ = fi92
                                                    _ = v93
                                                }
                                            }
                                        default:
                                            if ((fi12 ^ -9223372036854775808) > (fi12 - fi12)) {
                                                v97 := GenS14{F0: false, F1: 38}
                                                _ = v97
                                            } else {
                                                v98 := []int{(fi12 + fi12), (^(fi12)), (^(fi12))}
                                                if (990 <= 0) {
                                                    v11 += v11
                                                    switch (("" + "") + "\n\t") {
                                                    case "\n\t":
                                                        v99 := GenS15{F0: 0}
                                                        _ = v99
                                                    case "":
                                                        var v100 float64 = (-2.25 + ((-2.25 + ((-2.25 / ((-2.25 - 0.1) + 3.14159)) + float64(-2.25))) - (3.14159 / 1.0)))
                                                        _ = v100
                                                    }
                                                    var v101 uint32 = 4294967295
                                                    _ = v101
                                                } else {
                                                    v102 := GenS16{F0: 1, F1: "foo bar"}
                                                    if false {
                                                        fmt.Println(v11)
                                                        var v103 int64 = -9223372036854775808
                                                        _ = v103
                                                    } else {
                                                        var v104 float64 = 1.5
                                                        var v105 uint16 = 32767
                                                        _ = v104
                                                        _ = v105
                                                    }
                                                    for fi106 := 0; fi106 < 10; fi106++ {
                                                        var v107 uint64 = 4294967295
                                                        v108 := GenS9{F0: 183, F1: 0, F2: fi12}
                                                        if true {
                                                            var v109 int8 = 1
                                                            var v110 bool = (!false)
                                                            v111 := GenS1{F0: fi12, F1: 32767, F2: ("\n\t" + ("foo bar" + "a"))}
                                                            _ = v109
                                                            _ = v110
                                                            _ = v111
                                                        } else {
                                                            var v112 bool = ((((!false) || ((!false) || true)) && true) && (true || (!(false && false))))
                                                            _ = v112
                                                        }
                                                        _ = fi106
                                                        _ = v107
                                                        _ = v108
                                                    }
                                                    _ = v102
                                                }
                                                var v113 uint16 = 65535
                                                _ = v113
                                                _ = v98
                                            }
                                        }
                                    }
                                }
                            }
                            if false {
                                for fi114 := 0; fi114 < 12; fi114++ {
                                    v11 -= v11
                                    var v115 float64 = float64((-2.25 * 1e10))
                                    var v116 float32 = 3.14159
                                    v11 = v11
                                    _ = fi114
                                    _ = v115
                                    _ = v116
                                }
                                fmt.Println(v11, fi12)
                            }
                        }
                        for fi117 := 0; fi117 < 6; fi117++ {
                            v11 = v11
                            v118 := GenS16{F0: (fi12 / ((fi12) | 1)), F1: "foo bar"}
                            var v119 int = 0
                            _ = fi117
                            _ = v118
                            _ = v119
                        }
                    } else {
                        for fi120 := 0; fi120 < 14; fi120++ {
                            v11 = (v11 + -996)
                            _ = fi120
                        }
                        if true {
                            v121 := []int{(^(fi12))}
                            _ = v121
                        } else {
                            v11 -= (v11 + (^(v11)))
                            var v122 int32 = 908
                            var v123 int32 = 2147483647
                            _ = v122
                            _ = v123
                        }
                        v124 := GenS6{F0: true}
                        _ = v124
                    }
                    if ((1 >= (v11 + (v11 & v11))) || true) {
                        var v125 uint32 = 65535
                        v126 := GenS4{F0: v11, F1: ("a" + "foo bar")}
                        _ = v125
                        _ = v126
                    } else {
                        v11 = (-(v11))
                    }
                    fmt.Println(v11, v11)
                    _ = fi12
                }
            }
        }
    } else {
        for fi127 := 0; fi127 < 15; fi127++ {
            v128 := []int{(fi127 << ((1) & 63)), (fi127 / ((fi127) | 1)), (fi127 & (fi127 - fi127))}
            v11 += -2147483648
            switch -107 {
            case 189:
                fmt.Println(fi127)
            }
            _ = fi127
            _ = v128
        }
        var v129 int8 = 127
        v130 := GenS11{F0: 18446744073709551615}
        _ = v129
        _ = v130
    }
    _ = v10
    _ = v11
    return 1
}

func main() {
    var v1 uint8 = 2
    var v2 uint64 = 805
    var v3 int32 = 1
    for fi131 := 0; fi131 < 17; fi131++ {
        v132 := []int{(fi131 / (((fi131 + fi131)) | 1))}
        v133 := GenS6{F0: false}
        var v134 int = fi131
        _ = fi131
        _ = v132
        _ = v133
        _ = v134
    }
    var v135 float32 = float32(float64(3.14159))
    if false {
        switch -2147483648 {
        case -460:
            var v136 float32 = float32(float64(0.1))
            _ = v136
        case 161:
            if false {
                for fi137 := 0; fi137 < 11; fi137++ {
                    for fi138 := 0; fi138 < 4; fi138++ {
                        v139 := []int{128}
                        var v140 int32 = (v3 ^ v3)
                        _ = fi138
                        _ = v139
                        _ = v140
                    }
                    if false {
                        var v141 uint8 = uint8(v2)
                        _ = v141
                    } else {
                        v142 := GenS17{F0: false, F1: int(v2), F2: 0}
                        if (v3 <= (v3 + 2147483647)) {
                            var v143 uint64 = v2
                            fmt.Println(v135)
                            v143 = 0
                            _ = v143
                        }
                        v144 := GenS1{F0: (^(fi137)), F1: -326, F2: "\n\t"}
                        _ = v142
                        _ = v144
                    }
                    switch "hello" {
                    case "hello":
                        v3 = v3
                    default:
                        fmt.Println(v2)
                    }
                    _ = fi137
                }
            }
        }
    }
    v135 -= (float32(float64((((float32((0.0 * (0.1 + 1.5))) / float32(3.14159)) * 0.0) - (0.1 + v135)))) - (3.14159 - 0.0))
    switch -560 {
    case 127:
        var v145 string = ("a" + "a")
        _ = v145
    case 4294967295:
        switch 188 {
        case -9223372036854775808:
            var v146 int64 = -1
            _ = v146
        case 42:
            for fi147 := 0; fi147 < 5; fi147++ {
                for fi148 := 0; fi148 < 17; fi148++ {
                    v149 := []int{(-(fi147)), fi148}
                    _ = fi148
                    _ = v149
                }
                fmt.Println(v2)
                _ = fi147
            }
        default:
            var v150 uint8 = 128
            _ = v150
        }
    }
    for fi151 := 0; fi151 < 14; fi151++ {
        v152 := []int{}
        if false {
            v2 = (v2 >> ((42) & 63))
            if true {
                v3 += -1
                fmt.Println(v2, v1)
                v153 := GenS18{F0: 18446744073709551615}
                _ = v153
            }
            for fi154 := 0; fi154 < 13; fi154++ {
                v155 := GenS10{F0: "", F1: (v2 & uint64(v1))}
                v156 := GenS19{F0: 0}
                _ = fi154
                _ = v155
                _ = v156
            }
        } else {
            _ = len(v152)
            if len(v152) >= 2 {
                v152 = v152[1:len(v152)]
            }
            for fi157 := 0; fi157 < 5; fi157++ {
                switch fi151 {
                case 32767:
                    if len(v152) >= 2 {
                        v152 = v152[1:len(v152)]
                    }
                case 73:
                    fmt.Println(v3)
                case 32768:
                    switch fi151 {
                    case 127:
                        v158 := GenS8{F0: v2, F1: 0, F2: true}
                        _ = v158
                    case 1:
                        var v159 int8 = -1
                        _ = v159
                    case 128:
                        v160 := GenS13{F0: (true || true), F1: int(v2)}
                        _ = v160
                    }
                    v1 = 1
                }
                if true {
                    var v161 int64 = 42
                    v162 := GenS20{F0: "\n\t", F1: (-(fi151))}
                    _ = len(v152)
                    _ = v161
                    _ = v162
                } else {
                    v2 = v2
                    fmt.Println(v135)
                    var v163 int32 = int32(fi151)
                    _ = v163
                }
                _ = fi157
            }
        }
        _ = fi151
        _ = v152
    }
    if ("" > "a") {
        v1 = v1
    } else {
        v2 += (v2 << ((18446744073709551615) & 63))
    }
    v1 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v135:", v135)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v135
    _ = v2
    _ = v3
}
