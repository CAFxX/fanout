package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32; F1 uint }
type GenS6 struct { F0 int16; F1 string }
type GenS7 struct { F0 int32; F1 uint8 }
type GenS8 struct { F0 uint; F1 int16 }
type GenS9 struct { F0 uint16; F1 int64; F2 int8 }
type GenS10 struct { F0 int; F1 bool; F2 uint16 }
type GenS11 struct { F0 uint8; F1 uint64; F2 uint32 }
type GenS12 struct { F0 uint32 }
func genFn0() string {
    for fi4 := 0; fi4 < 9; fi4++ {
        var v5 bool = (0.1 < float32(float64(float32(-2.25))))
        for fi6 := 0; fi6 < 16; fi6++ {
            if v5 {
                v5 = v5
                var v7 int = 32768
                _ = v7
            } else {
                fmt.Println(fi4)
                fmt.Println(fi4, v5)
            }
            var v8 float32 = ((float32(0.1) / (float32(float64(-2.25)) / 0.1)) - 1.5)
            var v9 uint = 308
            switch fi4 {
            case 42:
                v10 := []int{fi6, fi6, fi6, (fi6 ^ fi6)}
                _ = v10
            case -881:
                v8 = (float32((-2.25 * 3.14159)) * -2.25)
            case -1:
                switch (fi4 & 128) {
                case 144:
                    if true {
                        v11 := GenS1{F0: 42, F1: 42}
                        _ = v11
                    }
                }
            }
            _ = fi6
            _ = v8
            _ = v9
        }
        var v12 uint64 = 4294967295
        _ = fi4
        _ = v12
        _ = v5
    }
    var v13 int32 = 2147483647
    switch (("foo bar" + "foo bar") + "\n\t") {
    case "foo bar":
        switch 9223372036854775807 {
        case 1:
            var v14 int = 1
            _ = v14
        case 128:
            fmt.Println(v13)
        }
    case "":
        for fi15 := 0; fi15 < 2; fi15++ {
            fmt.Println(fi15, v13)
            if false {
                v13 += 1
            }
            _ = fi15
        }
    case "hello":
        var v16 uint64 = 260
        _ = v16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        for fi18 := 0; fi18 < 6; fi18++ {
            v19 := []int{int(v13), 675, (fi17 % (((fi17 % ((32767) | 1))) | 1)), 1}
            if ((fi18 % (((fi18 & -778)) | 1)) >= fi17) {
                switch "\n\t" {
                case "a":
                    if len(v19) > 0 {
                        v19[0] = fi18
                    }
                }
                if (-1 != 42) {
                    if false {
                        v13 -= 42
                        var v20 uint = 870
                        v13 += v13
                        _ = v20
                    } else {
                        var v21 int64 = 1
                        var v22 int8 = -11
                        _ = v21
                        _ = v22
                    }
                    v13 = 2147483647
                    fmt.Println(fi17, v13)
                } else {
                    fmt.Println(fi17, fi18)
                    if false {
                        v23 := GenS2{F0: (v13 | 1), F1: 413}
                        v13 += v13
                        _ = v23
                    } else {
                        var v24 uint8 = 0
                        v13 -= v13
                        _ = v24
                    }
                    var v25 uint8 = 0
                    var v26 int8 = 127
                    _ = v25
                    _ = v26
                }
                v27 := GenS2{F0: -1, F1: 160}
                _ = v27
            }
            if (!false) {
                switch "hello" {
                case "foo bar":
                    var v28 uint = 606
                    _ = v28
                }
                if len(v19) >= 2 {
                    v19 = v19[1:len(v19)]
                }
                if false {
                    v13 = (v13 % (((v13 | (v13 << ((2) & 63)))) | 1))
                    var v29 bool = true
                    _ = v29
                }
            }
            _ = fi18
            _ = v19
        }
        var v30 int32 = v13
        switch "\n\t" {
        case "hello":
            if (!false) {
                switch "" {
                case "foo bar":
                    var v31 float32 = 0.0
                    _ = v31
                case "a":
                    var v32 int = (fi17 << ((1) & 63))
                    _ = v32
                }
                v33 := GenS1{F0: 2, F1: -32768}
                _ = v33
            } else {
                var v34 uint64 = 0
                _ = v34
            }
        case "\n\t":
            for fi35 := 0; fi35 < 2; fi35++ {
                v36 := []int{(^(fi35)), (^(fi17)), fi35}
                v13 += (v13 | v30)
                if true {
                    switch "foo bar" {
                    case "hello":
                        if ((v30 > (v30 << ((1) & 63))) && false) {
                            v30 += (v13 * (v13 << ((341) & 63)))
                            var v37 uint8 = 1
                            var v38 uint32 = 1
                            _ = v37
                            _ = v38
                        }
                        switch fi35 {
                        case -190:
                            if (0.0 != 0.0) {
                                fmt.Println(v30, v13, fi35)
                                var v39 int64 = -1
                                if ((1e10 < 0.0) || true) {
                                    var v40 uint64 = 18446744073709551615
                                    if false {
                                        var v41 int32 = v30
                                        v30 += (v41 << ((1) & 63))
                                        var v42 uint = 0
                                        v43 := GenS2{F0: (v30 >> (((v42 * 42)) & 63)), F1: (v40 >> (((^(v42))) & 63))}
                                        if len(v36) > 0 {
                                            v36[0] = (fi17 >> ((v42) & 63))
                                        }
                                        var v44 float32 = (0.1 - (-2.25 - 3.14159))
                                        _ = v41
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                    } else {
                                        if ((true || false) && (0 == 65535)) {
                                            if true {
                                                var v45 int32 = -913
                                                var v46 uint = 18446744073709551615
                                                _ = v45
                                                _ = v46
                                            }
                                        }
                                        v13 = v13
                                        var v47 float32 = 1e10
                                        switch fi17 {
                                        case 128:
                                            fmt.Println(fi17, fi35, fi17)
                                        }
                                        _ = v47
                                    }
                                    switch fi35 {
                                    case 32767:
                                        if false {
                                            var v48 int64 = v39
                                            var v49 int16 = 1
                                            if true {
                                                if (!(!(fi35 >= 1))) {
                                                    if (2147483647 > (v13 ^ (v13 / ((v13) | 1)))) {
                                                        v40 += (v40 % ((v40) | 1))
                                                    }
                                                    fmt.Println(v13, v30)
                                                } else {
                                                    var v50 uint32 = 75
                                                    if false {
                                                        switch "hello" {
                                                        case "hello":
                                                            v51 := GenS1{F0: 0, F1: 32767}
                                                            _ = v51
                                                        default:
                                                            if (true && true) {
                                                                v52 := GenS3{F0: 441, F1: 2}
                                                                fmt.Println(v40)
                                                                v36 = append(v36, 32767)
                                                                _ = v52
                                                            }
                                                        }
                                                        var v53 int16 = (v49 | v49)
                                                        _ = v53
                                                    }
                                                    _ = v50
                                                }
                                                var v54 string = (("a" + "foo bar") + ("\n\t" + ((("foo bar" + "") + "foo bar") + "")))
                                                v55 := GenS4{F0: 4294967295}
                                                _ = v54
                                                _ = v55
                                            } else {
                                                if (!(((3.14159 + ((1.5 - (0.0 * (-2.25 / 0.1))) * -2.25)) * 0.0) != (1.5 + 1e10))) {
                                                    if (1 > 884) {
                                                        fmt.Println(fi35, v30)
                                                        var v56 int8 = 1
                                                        _ = v56
                                                    } else {
                                                        if false {
                                                            var v57 uint16 = 42
                                                            _ = v57
                                                        } else {
                                                            fmt.Println(fi17)
                                                            if (((v13 << ((65535) & 63)) == v30) || (-1 == 1)) {
                                                                v58 := GenS4{F0: 0}
                                                                v36 = append(v36, (^(fi35)))
                                                                if len(v36) > 0 {
                                                                    v36[0] = fi17
                                                                }
                                                                _ = v58
                                                            }
                                                            var v59 float32 = (0.1 + -2.25)
                                                            _ = v59
                                                        }
                                                        v49 = v49
                                                    }
                                                    if true {
                                                        if true {
                                                            if false {
                                                                var v60 uint8 = 42
                                                                _ = v60
                                                            } else {
                                                                if false {
                                                                    var v61 uint = 748
                                                                    v62 := GenS2{F0: 1, F1: (v40 / (((v40 % ((v40) | 1))) | 1))}
                                                                    _ = v61
                                                                    _ = v62
                                                                } else {
                                                                    v49 += (v49 << ((0) & 63))
                                                                    switch "hello" {
                                                                    case "a":
                                                                        if false {
                                                                            if len(v36) >= 2 {
                                                                                v36 = v36[1:len(v36)]
                                                                            }
                                                                        } else {
                                                                            var v63 int64 = -1
                                                                            _ = v63
                                                                        }
                                                                    default:
                                                                        if len(v36) > 0 {
                                                                            v36[0] = fi17
                                                                        }
                                                                    }
                                                                    v64 := GenS3{F0: (fi17 + (fi17 | int(v49))), F1: 0}
                                                                    _ = v64
                                                                }
                                                                switch -9223372036854775808 {
                                                                case 4294967295:
                                                                    v65 := GenS5{F0: 0, F1: 255}
                                                                    _ = v65
                                                                case 187:
                                                                    fmt.Println(fi17, v30, fi17)
                                                                case 127:
                                                                    var v66 uint8 = 128
                                                                    _ = v66
                                                                default:
                                                                    v67 := GenS4{F0: 2}
                                                                    _ = v67
                                                                }
                                                            }
                                                            switch "a" {
                                                            case "a":
                                                                fmt.Println(v30, v49)
                                                            case "":
                                                                if len(v36) >= 2 {
                                                                    v36 = v36[1:len(v36)]
                                                                }
                                                            case "foo bar":
                                                                switch (fi35 - (fi17 << ((65535) & 63))) {
                                                                case 974:
                                                                    switch fi17 {
                                                                    case 278:
                                                                        v68 := GenS5{F0: 696, F1: 65535}
                                                                        _ = v68
                                                                    case 4294967295:
                                                                        var v69 float64 = ((1e10 * float64(float32(1e10))) - 0.1)
                                                                        _ = v69
                                                                    case -1:
                                                                        if (false || true) {
                                                                            var v70 bool = (!false)
                                                                            _ = v70
                                                                        }
                                                                    }
                                                                case 2147483647:
                                                                    var v71 int32 = -2147483648
                                                                    _ = v71
                                                                default:
                                                                    v72 := GenS2{F0: -651, F1: 0}
                                                                    _ = v72
                                                                }
                                                            default:
                                                                var v73 int = (fi35 >> ((uint(v40)) & 63))
                                                                _ = v73
                                                            }
                                                            _ = len(v36)
                                                        }
                                                    } else {
                                                        if (true || ((-2.25 - (-2.25 / (3.14159 - float32(1e10)))) > 0.0)) {
                                                            v74 := GenS4{F0: 2}
                                                            v49 = v49
                                                            _ = v74
                                                        } else {
                                                            var v75 uint16 = 2
                                                            _ = v75
                                                        }
                                                    }
                                                } else {
                                                    switch "foo bar" {
                                                    case "\n\t":
                                                        switch "foo bar" {
                                                        case "hello":
                                                            var v76 int = (-(fi17))
                                                            if (v30 >= (v30 | 42)) {
                                                                v49 = (v49 >> ((65535) & 63))
                                                                var v77 uint8 = 2
                                                                _ = v77
                                                            } else {
                                                                fmt.Println(v49, v13)
                                                                if true {
                                                                    var v78 uint32 = 783
                                                                    if len(v36) > 0 {
                                                                        v36[0] = fi35
                                                                    }
                                                                    _ = v78
                                                                }
                                                            }
                                                            _ = v76
                                                        case "":
                                                            v13 += (v30 + (-(v30)))
                                                        }
                                                        var v79 uint16 = 245
                                                        _ = v79
                                                    default:
                                                        v39 = v48
                                                    }
                                                    if true {
                                                        var v80 uint = 1
                                                        _ = v80
                                                    } else {
                                                        switch "hello" {
                                                        case "":
                                                            var v81 float64 = 3.14159
                                                            _ = v81
                                                        case "\n\t":
                                                            v82 := GenS4{F0: 65535}
                                                            _ = v82
                                                        case "a":
                                                            v13 = 42
                                                        default:
                                                            v83 := GenS6{F0: (v49 / (((^(v49))) | 1)), F1: "foo bar"}
                                                            _ = v83
                                                        }
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            v36 = append(v36, (fi35 ^ fi35))
                                                        case "a":
                                                            if true {
                                                                if (0 == 0) {
                                                                    var v84 float64 = 0.0
                                                                    _ = v84
                                                                } else {
                                                                    var v85 int32 = v30
                                                                    switch (fi35 - 2147483647) {
                                                                    case -1:
                                                                        if (811 == 18446744073709551615) {
                                                                            var v86 uint = 255
                                                                            var v87 int8 = -1
                                                                            _ = v86
                                                                            _ = v87
                                                                        } else {
                                                                            var v88 int64 = v39
                                                                            switch ("hello" + "foo bar") {
                                                                            case "":
                                                                                var v89 float64 = (((0.0 / 1.0) / -2.25) * (1.5 * 1.5))
                                                                                _ = v89
                                                                            }
                                                                            var v90 uint32 = 0
                                                                            var v91 float32 = 1e10
                                                                            if ((fi35 / ((fi35) | 1)) < fi17) {
                                                                                var v92 string = ""
                                                                                var v93 float64 = 1.5
                                                                                v94 := GenS1{F0: 214, F1: (v49 - v49)}
                                                                                _ = v92
                                                                                _ = v93
                                                                                _ = v94
                                                                            }
                                                                            v36 = append(v36, fi17)
                                                                            _ = v88
                                                                            _ = v90
                                                                            _ = v91
                                                                        }
                                                                    case -2147483648:
                                                                        if (!false) {
                                                                            fmt.Println(v40, fi35, fi17)
                                                                            var v95 float32 = 1e10
                                                                            _ = v95
                                                                        } else {
                                                                            fmt.Println(fi17, v13, fi35)
                                                                            v36 = append(v36, fi17)
                                                                        }
                                                                    case 1:
                                                                        v49 += 32767
                                                                    }
                                                                    _ = v85
                                                                }
                                                            } else {
                                                                if (2 > 65535) {
                                                                    if false {
                                                                        if true {
                                                                            v96 := GenS7{F0: 2147483647, F1: 255}
                                                                            _ = v96
                                                                        }
                                                                        fmt.Println(v13)
                                                                    } else {
                                                                        if true {
                                                                            var v97 string = ("a" + "\n\t")
                                                                            if len(v36) >= 2 {
                                                                                v36 = v36[1:len(v36)]
                                                                            }
                                                                            _ = v97
                                                                        } else {
                                                                            if true {
                                                                                fmt.Println(fi17, v30)
                                                                                fmt.Println(v49, v39)
                                                                                var v98 uint8 = 1
                                                                                fmt.Println(v48)
                                                                                _ = v98
                                                                            }
                                                                            var v99 uint8 = 90
                                                                            _ = v99
                                                                        }
                                                                        var v100 uint32 = 2
                                                                        _ = v100
                                                                    }
                                                                } else {
                                                                    v39 = (v39 % (((v48 & (v48 >> ((363) & 63)))) | 1))
                                                                    var v101 int64 = (v48 >> ((18446744073709551615) & 63))
                                                                    _ = v101
                                                                }
                                                            }
                                                        case "hello":
                                                            var v102 string = "hello"
                                                            _ = v102
                                                        }
                                                    }
                                                    var v103 float64 = 0.1
                                                    var v104 int16 = v49
                                                    switch "hello" {
                                                    case "a":
                                                        v39 = v39
                                                    default:
                                                        var v105 uint16 = 424
                                                        _ = v105
                                                    }
                                                    _ = v103
                                                    _ = v104
                                                }
                                                if false {
                                                    v30 += 0
                                                    v106 := GenS4{F0: 18446744073709551615}
                                                    _ = v106
                                                }
                                                v40 -= v40
                                                switch "foo bar" {
                                                case "a":
                                                    v49 += -1
                                                case "":
                                                    var v107 int32 = (v30 * (-(v13)))
                                                    _ = v107
                                                case "foo bar":
                                                    _ = len(v36)
                                                default:
                                                    var v108 int8 = 37
                                                    _ = v108
                                                }
                                                var v109 uint8 = 1
                                                _ = v109
                                            }
                                            _ = v48
                                            _ = v49
                                        }
                                    case -915:
                                        v110 := GenS4{F0: 42}
                                        _ = v110
                                    default:
                                        v36 = append(v36, 1)
                                    }
                                    _ = v40
                                }
                                if (v30 == -2147483648) {
                                    if true {
                                        if false {
                                            v39 += int64(v13)
                                            v13 -= (v13 << ((4294967295) & 63))
                                        }
                                    } else {
                                        v111 := GenS8{F0: 2, F1: 32767}
                                        v112 := GenS2{F0: (v13 / ((v13) | 1)), F1: 1}
                                        v13 = -705
                                        _ = v111
                                        _ = v112
                                    }
                                    v30 = (v13 / (((v30 | 2147483647)) | 1))
                                    var v113 float64 = 3.14159
                                    _ = v113
                                } else {
                                    fmt.Println(v13, v13)
                                }
                                if ((fi17 >= 613) || false) {
                                    var v114 string = "foo bar"
                                    if len(v36) > 0 {
                                        v36[0] = fi17
                                    }
                                    _ = v114
                                } else {
                                    v39 = (v39 << ((419) & 63))
                                }
                                _ = v39
                            }
                        default:
                            var v115 int64 = 42
                            _ = v115
                        }
                    case "\n\t":
                        if false {
                            fmt.Println(v13)
                        } else {
                            if true {
                                fmt.Println(fi35, v13, fi35)
                            }
                            v30 += (v30 & v13)
                            var v116 float32 = ((-2.25 + 3.14159) / 1e10)
                            _ = v116
                        }
                    default:
                        fmt.Println(v30, fi17)
                    }
                    var v117 int32 = 2147483647
                    v118 := GenS9{F0: 27, F1: 1, F2: -126}
                    v117 = v117
                    _ = v117
                    _ = v118
                } else {
                    v30 -= (v30 + -979)
                    fmt.Println(fi17, v13)
                }
                if (0 <= -128) {
                    if false {
                        var v119 bool = ((v13 / ((v30) | 1)) <= v13)
                        v119 = (32767 > 65535)
                        var v120 bool = (!true)
                        if len(v36) >= 2 {
                            v36 = v36[1:len(v36)]
                        }
                        var v121 uint = 1
                        var v122 bool = false
                        _ = v119
                        _ = v120
                        _ = v121
                        _ = v122
                    } else {
                        var v123 float32 = -2.25
                        _ = v123
                    }
                    fmt.Println(v13, fi35)
                } else {
                    v124 := GenS10{F0: (fi35 + 2147483647), F1: true, F2: 255}
                    v125 := GenS3{F0: -2147483648, F1: 74}
                    v126 := GenS9{F0: 0, F1: -1, F2: 1}
                    _ = v124
                    _ = v125
                    _ = v126
                }
                if len(v36) > 0 {
                    v36[0] = fi35
                }
                _ = fi35
                _ = v36
            }
        case "a":
            fmt.Println(v13)
        default:
            v30 -= 533
        }
        _ = fi17
        _ = v30
    }
    v127 := GenS11{F0: 2, F1: 1, F2: 2147483648}
    _ = v127
    _ = v13
    return ("a" + "hello")
}
func genFn1(p0 int16, p1 bool) string {
    p1 = p1
    p0 -= 1000
    p1 = (-71 != 42)
    _ = p0
    _ = p1
    return "a"
}

func main() {
    var v1 int = -679
    v2 := make(map[string]bool)
    v3 := make(map[uint32]bool)
    switch (v1 * v1) {
    case 42:
        switch "hello" {
        case "hello":
            if ("a" < ("\n\t" + (("a" + "hello") + "hello"))) {
                if ("\n\t" < ("" + "hello")) {
                    for fi128 := 0; fi128 < 4; fi128++ {
                        var v129 int32 = 263
                        fmt.Println(v1)
                        _ = fi128
                        _ = v129
                    }
                    v130 := []int{}
                    v1 += 953
                    _ = v130
                }
                v131 := GenS12{F0: 0}
                _ = v131
            } else {
                switch v1 {
                case 32768:
                    delete(v2, ("foo bar" + "foo bar"))
                case 42:
                    var v132 int8 = 42
                    _ = v132
                case 9223372036854775807:
                    if (false || (((false && false) || (-894 < 184)) && (-9223372036854775808 <= -9223372036854775808))) {
                        _ = len(v2)
                        v133 := GenS8{F0: 65535, F1: 32767}
                        for fi134 := 0; fi134 < 17; fi134++ {
                            v135 := v3[42]
                            fmt.Println(fi134, v1)
                            var v136 uint64 = 116
                            _ = fi134
                            _ = v135
                            _ = v136
                        }
                        _ = v133
                    }
                default:
                    switch (v1 + (v1 * 1)) {
                    case 448:
                        v137 := GenS2{F0: 614, F1: 1}
                        _ = v137
                    }
                }
                v1 = (v1 - (v1 / (((v1 / (((-(v1))) | 1))) | 1)))
            }
        case "":
            if true {
                var v138 int32 = -2147483648
                v139 := v2["\n\t"]
                _ = v138
                _ = v139
            } else {
                var v140 float64 = 0.1
                _ = v140
            }
        case "foo bar":
            v1 -= -2147483648
        default:
            var v141 float64 = (-2.25 / -2.25)
            _ = v141
        }
    default:
        var v142 bool = (!true)
        _ = v142
    }
    v1 -= (v1 * v1)
    if (1 < -128) {
        if (0 < -9223372036854775808) {
            v1 = v1
        } else {
            v1 = 9223372036854775807
            v143 := []int{127, v1}
            _ = v143
        }
        v1 = 4294967295
    } else {
        if (753 < 42) {
            v144 := []int{(v1 - (v1 - (v1 % ((v1) | 1)))), 9223372036854775807}
            _ = v144
        }
    }
    if false {
        var v145 int = v1
        var v146 uint = 4294967295
        _ = v145
        _ = v146
    } else {
        v147 := GenS4{F0: 1}
        v1 -= (v1 + (v1 / ((607) | 1)))
        _ = v147
    }
    var v148 int32 = -2147483648
    switch "a" {
    case "hello":
        _ = len(v3)
    }
    var v149 int8 = -1
    v148 += v148
    var v150 float64 = 0.0
    var v151 int8 = (-(v149))
    fmt.Println(v151, v149)
    v149 += v149
    fmt.Println("v1:", v1)
    fmt.Println("v148:", v148)
    fmt.Println("v149:", v149)
    fmt.Println("v150:", v150)
    fmt.Println("v151:", v151)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v148
    _ = v149
    _ = v150
    _ = v151
    _ = v2
    _ = v3
}
