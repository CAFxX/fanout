package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 bool; F1 bool; F2 int32 }
type GenS4 struct { F0 int64; F1 int32; F2 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int; F1 int32; F2 string }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 string; F1 int; F2 int32 }
type GenS9 struct { F0 uint32; F1 uint8; F2 int32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int64; F1 int16; F2 bool }
type GenS12 struct { F0 uint64; F1 uint16; F2 bool }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint8; F1 int32 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 uint8; F1 uint64 }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 bool; F1 uint }
type GenS19 struct { F0 uint32; F1 int32; F2 int32 }
type GenS20 struct { F0 int64; F1 uint32; F2 int32 }
func genFn0(p0 bool) uint16 {
    v5 := []int{359, 920}
    for fi6 := 0; fi6 < 14; fi6++ {
        p0 = p0
        _ = fi6
    }
    if p0 {
        for fi7 := 0; fi7 < 12; fi7++ {
            v5 = append(v5, fi7)
            _ = fi7
        }
        var v8 uint16 = 255
        _ = v8
    }
    switch 0 {
    case 2147483647:
        var v9 int32 = 2147483647
        _ = v9
    case -1:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case 9223372036854775807:
        if len(v5) > 0 {
            v5[0] = -1
        }
    default:
        var v10 int8 = 127
        _ = v10
    }
    _ = p0
    _ = v5
    return 32767
}
func genFn1() int64 {
    var v11 uint = 214
    for fi12 := 0; fi12 < 19; fi12++ {
        for fi13 := 0; fi13 < 8; fi13++ {
            v14 := []int{fi12}
            v11 -= (v11 ^ v11)
            var v15 string = "\n\t"
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    var v16 int32 = 497
    v17 := []int{450, -862, -736}
    _ = v11
    _ = v16
    _ = v17
    return 9223372036854775807
}

func main() {
    v1 := []string{(("" + ("hello" + "hello")) + ("a" + "\n\t"))}
    v2 := make(map[string]string)
    var v3 int8 = 42
    var v4 int32 = -1
    if true {
        fmt.Println(v3, v3)
        switch "\n\t" {
        case "\n\t":
            var v18 float64 = -2.25
            _ = v18
        default:
            v3 -= v3
        }
        var v19 int8 = v3
        _ = v19
    }
    if false {
        if (((-725 < 128) || true) && (v3 >= (v3 & 1))) {
            var v20 uint64 = 42
            switch 588 {
            case 128:
                for fi21 := 0; fi21 < 12; fi21++ {
                    v22 := GenS1{F0: 1, F1: v3}
                    var v23 uint8 = 42
                    fmt.Println(v23)
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            case 32768:
                for fi24 := 0; fi24 < 9; fi24++ {
                    v25 := v2["a"]
                    var v26 int16 = 0
                    v27 := GenS2{F0: (v26 * (v26 << ((1) & 63))), F1: 0}
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
            case -1:
                v28 := GenS1{F0: -1, F1: v3}
                _ = v28
            default:
                v29 := GenS3{F0: (42 > genFn1()), F1: false, F2: -2147483648}
                _ = v29
            }
            _ = v20
        } else {
            for fi30 := 0; fi30 < 6; fi30++ {
                if false {
                    fmt.Println(v3)
                    for fi31 := 0; fi31 < 18; fi31++ {
                        fmt.Println(fi31, fi31)
                        var v32 int = fi30
                        var v33 float64 = 1e10
                        _ = fi31
                        _ = v32
                        _ = v33
                    }
                    if (-1 <= (v4 | (^(v4)))) {
                        v34 := v2[("" + ("" + "\n\t"))]
                        v34 = v34
                        _ = v34
                    }
                } else {
                    v35 := GenS4{F0: 1, F1: (^(v4)), F2: "a"}
                    _ = len(v1)
                    v3 += v3
                    _ = v35
                }
                _ = fi30
            }
            v3 += v3
            if false {
                switch ("hello" + "") {
                case "foo bar":
                    v3 += v3
                default:
                    for fi36 := 0; fi36 < 1; fi36++ {
                        v3 = (^(v3))
                        switch 127 {
                        case 0:
                            switch "foo bar" {
                            case "":
                                fmt.Println(v3, v3, v3)
                            case "a":
                                var v37 float64 = 1.5
                                _ = v37
                            default:
                                v3 = -1
                            }
                        case -995:
                            v38 := GenS3{F0: ("hello" >= ("hello" + "hello")), F1: (v4 == (v4 & (v4 - -2147483648))), F2: 37}
                            _ = v38
                        }
                        if true {
                            if true {
                                if false {
                                    for fi39 := 0; fi39 < 4; fi39++ {
                                        _ = len(v2)
                                        v40 := v2["\n\t"]
                                        delete(v2, (v40 + ("hello" + v40)))
                                        _ = fi39
                                        _ = v40
                                    }
                                    v41 := GenS5{F0: genFn0(((!false) && false))}
                                    switch "" {
                                    case "hello":
                                        var v42 int32 = (v4 | v4)
                                        _ = v42
                                    case "\n\t":
                                        for fi43 := 0; fi43 < 11; fi43++ {
                                            if ((v3 / ((1) | 1)) < v3) {
                                                v44 := v2[(("" + "\n\t") + "hello")]
                                                v45 := GenS4{F0: -9223372036854775808, F1: 2147483647, F2: "\n\t"}
                                                _ = v44
                                                _ = v45
                                            } else {
                                                v4 -= (v4 ^ (v4 << ((18446744073709551615) & 63)))
                                                v46 := GenS6{F0: (fi36 - 32768), F1: v4, F2: ""}
                                                fmt.Println(v4, v3)
                                                v47 := GenS6{F0: fi43, F1: (v4 | (v4 - (v4 ^ 0))), F2: ""}
                                                _ = v46
                                                _ = v47
                                            }
                                            switch ("hello" + ("a" + (("" + ((("foo bar" + "") + "foo bar") + "")) + ("hello" + ("a" + "hello"))))) {
                                            case "hello":
                                                v4 = (-(v4))
                                            case "a":
                                                if false {
                                                    var v48 uint32 = 329
                                                    v3 = 54
                                                    var v49 float64 = (1e10 + 1e10)
                                                    _ = v48
                                                    _ = v49
                                                } else {
                                                    switch ("\n\t" + "\n\t") {
                                                    case "":
                                                        v4 = int32(fi43)
                                                    case "hello":
                                                        if (1e10 <= 0.0) {
                                                            if false {
                                                                _ = len(v2)
                                                                v4 -= v4
                                                            } else {
                                                                switch fi43 {
                                                                case 128:
                                                                    v50 := GenS6{F0: fi36, F1: (-(v4)), F2: "foo bar"}
                                                                    _ = v50
                                                                }
                                                                switch "" {
                                                                case "\n\t":
                                                                    var v51 int32 = -1
                                                                    _ = v51
                                                                case "hello":
                                                                    v52 := GenS7{F0: 538}
                                                                    _ = v52
                                                                default:
                                                                    v53 := GenS5{F0: 1}
                                                                    v4 = v4
                                                                    _ = v53
                                                                }
                                                                if (((3.14159 * 1.5) < -2.25) || false) {
                                                                    var v54 uint32 = 42
                                                                    _ = v54
                                                                } else {
                                                                    v55 := GenS4{F0: 993, F1: 42, F2: ("hello" + "foo bar")}
                                                                    var v56 float32 = 1.5
                                                                    var v57 int16 = -32768
                                                                    var v58 float64 = 0.0
                                                                    var v59 float32 = 0.1
                                                                    _ = v55
                                                                    _ = v56
                                                                    _ = v57
                                                                    _ = v58
                                                                    _ = v59
                                                                }
                                                            }
                                                            fmt.Println(fi43, v4)
                                                        } else {
                                                            v3 += v3
                                                            if true {
                                                                v4 = (v4 % ((v4) | 1))
                                                                v60 := GenS8{F0: ("hello" + ("a" + "\n\t")), F1: fi43, F2: v4}
                                                                v4 = v4
                                                                _ = v60
                                                            }
                                                        }
                                                    case "\n\t":
                                                        v4 -= (v4 - v4)
                                                    default:
                                                        if (((0.1 / ((0.1 / 1e10) / 1.5)) / 0.1) > (float64(-2.25) + (float64(float32((3.14159 - 1.5))) / (1e10 / 1e10)))) {
                                                            if (2 <= 962) {
                                                                var v61 uint16 = 314
                                                                _ = v61
                                                            }
                                                            switch (fi36 ^ 128) {
                                                            case 42:
                                                                if (729 <= 0) {
                                                                    if true {
                                                                        v2[""] = (("foo bar" + ("hello" + ("a" + ""))) + "hello")
                                                                        var v62 uint = 501
                                                                        _ = v62
                                                                    }
                                                                    var v63 string = (("hello" + ("\n\t" + ("a" + "a"))) + "hello")
                                                                    switch int(v4) {
                                                                    case 1:
                                                                        var v64 bool = false
                                                                        _ = v64
                                                                    case -38:
                                                                        var v65 uint64 = 4294967295
                                                                        _ = v65
                                                                    case 921:
                                                                        switch fi43 {
                                                                        case 4294967295:
                                                                            v66 := GenS8{F0: v63, F1: (fi43 >> ((31) & 63)), F2: (v4 >> ((552) & 63))}
                                                                            _ = v66
                                                                        default:
                                                                            v3 = 105
                                                                        }
                                                                    default:
                                                                        var v67 int32 = (^(v4))
                                                                        _ = v67
                                                                    }
                                                                    _ = v63
                                                                } else {
                                                                    var v68 int64 = -9223372036854775808
                                                                    var v69 int = (fi36 * -796)
                                                                    var v70 uint8 = 255
                                                                    _ = v68
                                                                    _ = v69
                                                                    _ = v70
                                                                }
                                                            case -388:
                                                                var v71 int8 = (v3 - 127)
                                                                _ = v71
                                                            }
                                                            v72 := GenS8{F0: "hello", F1: fi43, F2: v4}
                                                            _ = v72
                                                        }
                                                    }
                                                    var v73 uint = 42
                                                    fmt.Println(v4)
                                                    _ = v73
                                                }
                                            }
                                            _ = fi43
                                        }
                                    case "foo bar":
                                        v74 := GenS2{F0: 0, F1: v4}
                                        _ = v74
                                    default:
                                        if (664 < 732) {
                                            v75 := GenS7{F0: 65535}
                                            _ = v75
                                        }
                                    }
                                    _ = v41
                                }
                                if ((v4 * 1) == v4) {
                                    v76 := v2["foo bar"]
                                    _ = v76
                                }
                                if false {
                                    v77 := GenS8{F0: "a", F1: (fi36 | -989), F2: v4}
                                    v4 -= v4
                                    _ = v77
                                } else {
                                    v78 := GenS2{F0: -20, F1: (v4 % (((v4 * (-(v4)))) | 1))}
                                    var v79 int16 = -1
                                    var v80 float64 = 0.0
                                    _ = v78
                                    _ = v79
                                    _ = v80
                                }
                            } else {
                                v81 := GenS2{F0: 0, F1: int32(v3)}
                                _ = v81
                            }
                            switch fi36 {
                            case -995:
                                v82 := GenS9{F0: 2, F1: 42, F2: (v4 & (-(v4)))}
                                _ = v82
                            case 42:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            default:
                                v83 := GenS10{F0: 0}
                                _ = v83
                            }
                            for fi84 := 0; fi84 < 7; fi84++ {
                                var v85 uint32 = 42
                                if (!((false || true) && true)) {
                                    fmt.Println(v3, v85, v4)
                                    if (4294967295 == 2) {
                                        var v86 float32 = ((0.0 - (0.0 / (0.1 / 0.1))) + 0.0)
                                        if true {
                                            if false {
                                                var v87 uint32 = (^(v85))
                                                v88 := v2[""]
                                                v88 = "foo bar"
                                                v3 -= 0
                                                _ = v87
                                                _ = v88
                                            } else {
                                                if len(v1) > 0 {
                                                    v1[0] = ("hello" + "")
                                                }
                                                switch 32768 {
                                                case -9223372036854775808:
                                                    v86 = float32((((3.14159 * 0.1) + 0.1) - 0.0))
                                                }
                                                if (1e10 < 1e10) {
                                                    var v89 int16 = 0
                                                    _ = v89
                                                } else {
                                                    switch ("" + ("a" + (("foo bar" + ("" + "")) + "foo bar"))) {
                                                    case "foo bar":
                                                        var v90 int = (fi84 << ((18446744073709551615) & 63))
                                                        _ = v90
                                                    }
                                                    var v91 int16 = int16(fi36)
                                                    switch "hello" {
                                                    case "hello":
                                                        var v92 int8 = (v3 >> ((2) & 63))
                                                        _ = v92
                                                    }
                                                    switch fi84 {
                                                    case -625:
                                                        v93 := GenS1{F0: (^(v91)), F1: v3}
                                                        _ = v93
                                                    case 913:
                                                        if false {
                                                            v86 += v86
                                                        }
                                                    case -1:
                                                        var v94 int32 = 665
                                                        _ = v94
                                                    }
                                                    _ = v91
                                                }
                                                var v95 uint64 = 1
                                                _ = v95
                                            }
                                            var v96 bool = (fi84 < fi84)
                                            v97 := GenS9{F0: (^(v85)), F1: 169, F2: (v4 & -2147483648)}
                                            _ = v96
                                            _ = v97
                                        } else {
                                            var v98 int32 = (-(v4))
                                            fmt.Println(v85, v98, v4)
                                            _ = v98
                                        }
                                        _ = v86
                                    } else {
                                        switch "foo bar" {
                                        case "foo bar":
                                            if true {
                                                v99 := GenS11{F0: 1, F1: 1, F2: (32767 < 837)}
                                                _ = v99
                                            }
                                        case "a":
                                            v85 -= v85
                                        default:
                                            v3 += v3
                                        }
                                    }
                                } else {
                                    _ = len(v2)
                                    var v100 uint = 215
                                    _ = v100
                                }
                                v101 := GenS12{F0: 42, F1: 392, F2: (210 != 42)}
                                var v102 uint8 = 255
                                _ = fi84
                                _ = v101
                                _ = v102
                                _ = v85
                            }
                        }
                        _ = fi36
                    }
                }
                switch "" {
                case "\n\t":
                    var v103 int16 = 42
                    _ = v103
                case "a":
                    delete(v2, ("" + "foo bar"))
                case "foo bar":
                    if ((-870 < -896) || (false && true)) {
                        var v104 int16 = 1
                        for fi105 := 0; fi105 < 18; fi105++ {
                            switch "" {
                            case "\n\t":
                                v106 := GenS1{F0: v104, F1: -1}
                                _ = v106
                            case "foo bar":
                                v107 := GenS11{F0: genFn1(), F1: v104, F2: true}
                                _ = v107
                            case "a":
                                var v108 int32 = v4
                                _ = v108
                            }
                            v109 := GenS13{F0: (v4 ^ v4)}
                            _ = fi105
                            _ = v109
                        }
                        _ = v104
                    }
                default:
                    v110 := GenS8{F0: "foo bar", F1: 943, F2: -145}
                    _ = v110
                }
            } else {
                if false {
                    v4 -= (v4 >> ((65535) & 63))
                    for fi111 := 0; fi111 < 5; fi111++ {
                        if true {
                            v112 := GenS10{F0: 0}
                            var v113 uint16 = 32767
                            var v114 int64 = -9223372036854775808
                            _ = v112
                            _ = v113
                            _ = v114
                        } else {
                            var v115 int = fi111
                            v115 += v115
                            _ = v115
                        }
                        var v116 float64 = ((0.0 * -2.25) * 0.1)
                        var v117 int32 = 425
                        _ = fi111
                        _ = v116
                        _ = v117
                    }
                }
                switch "a" {
                case "":
                    v4 = -1
                }
                if (464 < 47) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    for fi118 := 0; fi118 < 5; fi118++ {
                        if (!(false || true)) {
                            for fi119 := 0; fi119 < 7; fi119++ {
                                v120 := GenS1{F0: 42, F1: v3}
                                if false {
                                    v4 = v4
                                    v3 = (v3 / ((v3) | 1))
                                    v1 = append(v1, ("\n\t" + "a"))
                                }
                                v3 = -128
                                v4 += v4
                                if (18446744073709551615 <= 589) {
                                    if false {
                                        if false {
                                            v3 += v3
                                            if false {
                                                v121 := GenS10{F0: 42}
                                                v3 = -21
                                                _ = v121
                                            } else {
                                                v4 = (v4 & (^(v4)))
                                            }
                                            var v122 uint64 = 18446744073709551615
                                            _ = v122
                                        } else {
                                            if true {
                                                v123 := GenS4{F0: -1, F1: -2147483648, F2: (("\n\t" + "\n\t") + "hello")}
                                                delete(v2, ("a" + ""))
                                                _ = v123
                                            }
                                            if false {
                                                var v124 int64 = 1
                                                if (false || (!(1e10 <= (-2.25 * (1.5 * -2.25))))) {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                                _ = v124
                                            } else {
                                                var v125 float64 = 3.14159
                                                if false {
                                                    if false {
                                                        fmt.Println(v125, v4, v3)
                                                    } else {
                                                        v126 := v2["\n\t"]
                                                        switch (^(fi118)) {
                                                        case -2147483648:
                                                            var v127 uint64 = 42
                                                            _ = v127
                                                        case -1:
                                                            var v128 string = ((((v126 + v126) + "hello") + v126) + "hello")
                                                            _ = v128
                                                        case 0:
                                                            fmt.Println(v4, v3)
                                                        default:
                                                            var v129 uint8 = 165
                                                            _ = v129
                                                        }
                                                        delete(v2, v126)
                                                        _ = v126
                                                    }
                                                    var v130 float64 = float64((0.1 / (3.14159 / float32((v125 * (v125 / v125))))))
                                                    v131 := GenS14{F0: 1, F1: v4}
                                                    _ = v130
                                                    _ = v131
                                                } else {
                                                    fmt.Println(v3, v125, fi119)
                                                    if (true || false) {
                                                        var v132 uint16 = 913
                                                        _ = v132
                                                    }
                                                }
                                                if true {
                                                    var v133 int32 = (v4 * v4)
                                                    _ = v133
                                                } else {
                                                    if (254 < 1) {
                                                        var v134 int64 = -1
                                                        v135 := GenS9{F0: 1, F1: 127, F2: v4}
                                                        v136 := GenS13{F0: -1}
                                                        if true {
                                                            switch (fi119 % (((fi119 | (fi118 + 1))) | 1)) {
                                                            case 0:
                                                                var v137 uint32 = 2
                                                                _ = v137
                                                            case 32768:
                                                                v4 = 42
                                                            default:
                                                                v134 = genFn1()
                                                            }
                                                        }
                                                        v134 -= genFn1()
                                                        _ = v134
                                                        _ = v135
                                                        _ = v136
                                                    }
                                                    if (!(true || (true && false))) {
                                                        if false {
                                                            var v138 int16 = 1
                                                            v125 = -2.25
                                                            _ = v138
                                                        } else {
                                                            v139 := GenS5{F0: 255}
                                                            var v140 bool = (!true)
                                                            _ = v139
                                                            _ = v140
                                                        }
                                                    } else {
                                                        var v141 bool = true
                                                        if (315 >= 255) {
                                                            v3 = (v3 % (((v3 + v3)) | 1))
                                                            v141 = v141
                                                            v141 = v141
                                                        } else {
                                                            fmt.Println(v3)
                                                            var v142 int8 = v3
                                                            fmt.Println(v3, v125)
                                                            _ = v142
                                                        }
                                                        var v143 float32 = 0.1
                                                        _ = v141
                                                        _ = v143
                                                    }
                                                    var v144 int64 = -9223372036854775808
                                                    v125 += float64((0.1 + (1e10 + ((0.0 + (((0.1 / 1.0) + 1.5) - 1e10)) / 1e10))))
                                                    if ((true || false) || true) {
                                                        fmt.Println(fi119)
                                                        v145 := GenS15{F0: "\n\t"}
                                                        _ = v145
                                                    } else {
                                                        switch ("" + (("foo bar" + "foo bar") + "\n\t")) {
                                                        case "foo bar":
                                                            if ("a" != "\n\t") {
                                                                var v146 uint16 = 533
                                                                var v147 uint16 = 1
                                                                var v148 uint32 = 2147483648
                                                                _ = v146
                                                                _ = v147
                                                                _ = v148
                                                            }
                                                            if true {
                                                                _ = len(v1)
                                                                switch "foo bar" {
                                                                case "":
                                                                    v149 := v2["a"]
                                                                    _ = v149
                                                                }
                                                            } else {
                                                                var v150 uint8 = 255
                                                                v151 := v2[("foo bar" + "a")]
                                                                _ = v150
                                                                _ = v151
                                                            }
                                                            if true {
                                                                switch "foo bar" {
                                                                case "hello":
                                                                    v4 += v4
                                                                }
                                                                if false {
                                                                    var v152 uint64 = 1
                                                                    fmt.Println(fi118, fi119)
                                                                    v3 = -1
                                                                    _ = v152
                                                                }
                                                                v3 += v3
                                                                var v153 int32 = v4
                                                                var v154 uint = 135
                                                                _ = v153
                                                                _ = v154
                                                            } else {
                                                                var v155 int16 = 1000
                                                                v125 -= 0.1
                                                                v156 := GenS16{F0: 110, F1: 42}
                                                                _ = v155
                                                                _ = v156
                                                            }
                                                        case "a":
                                                            v157 := GenS9{F0: 2, F1: 2, F2: v4}
                                                            _ = v157
                                                        case "hello":
                                                            var v158 uint64 = 1
                                                            _ = v158
                                                        default:
                                                            v3 = (v3 - (-(v3)))
                                                        }
                                                        v159 := GenS6{F0: (fi119 + 32767), F1: (v4 << ((65535) & 63)), F2: "foo bar"}
                                                        if len(v1) > 0 {
                                                            v1[0] = ((("" + "foo bar") + "hello") + "hello")
                                                        }
                                                        _ = v159
                                                    }
                                                    if true {
                                                        switch "a" {
                                                        case "foo bar":
                                                            v3 = (v3 & (v3 << ((907) & 63)))
                                                        }
                                                        v160 := GenS3{F0: false, F1: false, F2: (v4 + -232)}
                                                        var v161 uint = 141
                                                        _ = v160
                                                        _ = v161
                                                    } else {
                                                        fmt.Println(fi118, v144, fi118)
                                                        switch "" {
                                                        case "a":
                                                            v125 += 0.0
                                                        case "foo bar":
                                                            fmt.Println(v4, v125)
                                                        }
                                                        v1 = append(v1, "a")
                                                    }
                                                    var v162 bool = true
                                                    v163 := GenS10{F0: 778}
                                                    _ = v144
                                                    _ = v162
                                                    _ = v163
                                                }
                                                _ = v125
                                            }
                                        }
                                        v164 := GenS17{F0: (588 >= 65535)}
                                        v4 = (v4 - v4)
                                        _ = v164
                                    }
                                    var v165 float32 = (float32(1.5) - -2.25)
                                    if false {
                                        v3 -= -63
                                    } else {
                                        if (false && (float32((3.14159 * (3.14159 / 3.14159))) <= float32(-2.25))) {
                                            fmt.Println(fi118)
                                            v1 = append(v1, ("\n\t" + ("" + "hello")))
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        } else {
                                            delete(v2, ("foo bar" + "foo bar"))
                                        }
                                    }
                                    _ = v165
                                } else {
                                    var v166 int8 = -68
                                    fmt.Println(fi118, v3, v4)
                                    if ((fi119 > fi119) && false) {
                                        var v167 uint = 1
                                        v166 -= 100
                                        var v168 int32 = (v4 >> ((369) & 63))
                                        _ = v167
                                        _ = v168
                                    }
                                    _ = v166
                                }
                                _ = fi119
                                _ = v120
                            }
                            v2["a"] = "\n\t"
                        }
                        _ = fi118
                    }
                    _ = len(v2)
                }
            }
        }
        v169 := GenS18{F0: true, F1: 4294967295}
        switch -651 {
        case 0:
            v3 += (v3 >> ((562) & 63))
        case -213:
            var v170 uint64 = 649
            _ = v170
        default:
            v4 = v4
        }
        _ = v169
    }
    v3 = -87
    fmt.Println(v4, v4)
    for fi171 := 0; fi171 < 14; fi171++ {
        if (23 < 2) {
            switch "a" {
            case "hello":
                if (976 == 65535) {
                    v3 = v3
                    v172 := GenS10{F0: 18446744073709551615}
                    _ = v172
                }
            case "":
                fmt.Println(fi171, v4)
            case "foo bar":
                if (0 == 1) {
                    v3 = v3
                    var v173 bool = true
                    _ = v173
                } else {
                    for fi174 := 0; fi174 < 18; fi174++ {
                        var v175 uint16 = 589
                        v176 := v2[(("\n\t" + "foo bar") + "")]
                        v3 += (v3 >> ((65535) & 63))
                        _ = fi174
                        _ = v175
                        _ = v176
                    }
                }
            default:
                if (142 != 0) {
                    var v177 uint64 = 734
                    v4 = (v4 % (((v4 + (v4 ^ (v4 ^ 42)))) | 1))
                    _ = v177
                }
            }
            if true {
                if false {
                    v3 -= v3
                }
                v3 += (v3 - v3)
                if (v3 >= v3) {
                    for fi178 := 0; fi178 < 4; fi178++ {
                        if false {
                            v3 -= v3
                        } else {
                            delete(v2, (("" + ("a" + "a")) + (("a" + "hello") + ("a" + ""))))
                        }
                        v4 = (^(v4))
                        fmt.Println(v3, v3, v3)
                        _ = fi178
                    }
                }
            } else {
                for fi179 := 0; fi179 < 3; fi179++ {
                    if (191 <= 4294967295) {
                        fmt.Println(v3)
                        v3 -= v3
                        var v180 int = fi171
                        _ = len(v2)
                        _ = v180
                    }
                    var v181 uint = 4294967295
                    _ = fi179
                    _ = v181
                }
                if ((false || true) && (!false)) {
                    fmt.Println(v4, fi171, v3)
                }
            }
            for fi182 := 0; fi182 < 18; fi182++ {
                var v183 int32 = (v4 << ((4294967295) & 63))
                _ = fi182
                _ = v183
            }
        } else {
            if true {
                switch "\n\t" {
                case "a":
                    fmt.Println(v3)
                case "\n\t":
                    fmt.Println(fi171, v4)
                }
                for fi184 := 0; fi184 < 5; fi184++ {
                    var v185 int8 = (-(v3))
                    switch (("foo bar" + "hello") + "hello") {
                    case "a":
                        var v186 bool = (1 <= 2)
                        var v187 uint = 0
                        _ = v186
                        _ = v187
                    case "hello":
                        delete(v2, ("foo bar" + ""))
                    default:
                        if false {
                            if (false && (387 != 1)) {
                                if false {
                                    var v188 uint16 = 42
                                    _ = v188
                                }
                                fmt.Println(v3)
                            } else {
                                v3 = v185
                                if (1.5 != (1e10 / 1.0)) {
                                    if false {
                                        switch "foo bar" {
                                        case "foo bar":
                                            v4 = v4
                                        }
                                        var v189 float64 = (((0.1 / -2.25) + 1e10) + 0.1)
                                        if (255 >= 0) {
                                            v190 := GenS7{F0: 168}
                                            if false {
                                                v185 += (v185 % (((v185 / ((v3) | 1))) | 1))
                                                v185 -= 0
                                                var v191 uint32 = 42
                                                _ = v191
                                            } else {
                                                switch (^(fi184)) {
                                                case 651:
                                                    v192 := GenS19{F0: 2147483647, F1: (v4 % ((31) | 1)), F2: 2147483647}
                                                    _ = v192
                                                case 32767:
                                                    v193 := GenS19{F0: 65535, F1: -691, F2: 724}
                                                    _ = v193
                                                case 656:
                                                    v194 := GenS20{F0: -9223372036854775808, F1: 65535, F2: v4}
                                                    _ = v194
                                                default:
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                                var v195 int = 459
                                                switch "a" {
                                                case "\n\t":
                                                    var v196 float32 = (1.5 * 1e10)
                                                    _ = v196
                                                case "a":
                                                    var v197 float64 = v189
                                                    _ = v197
                                                case "foo bar":
                                                    fmt.Println(v185)
                                                default:
                                                    fmt.Println(v189)
                                                }
                                                _ = v195
                                            }
                                            _ = v190
                                        } else {
                                            v189 = v189
                                            v3 = (v185 + (v185 | (v185 % ((v3) | 1))))
                                            _ = len(v2)
                                        }
                                        v189 -= 0.0
                                        v185 += v185
                                        _ = v189
                                    }
                                    fmt.Println(fi171, v185, v4)
                                    fmt.Println(v3)
                                }
                                v198 := GenS15{F0: "a"}
                                fmt.Println(v3, fi184)
                                _ = v198
                            }
                            var v199 int64 = -9223372036854775808
                            _ = v199
                        } else {
                            _ = len(v1)
                            var v200 uint = 1
                            _ = v200
                        }
                        if (!("hello" >= "foo bar")) {
                            var v201 float32 = -2.25
                            v202 := GenS1{F0: -1, F1: 100}
                            _ = v201
                            _ = v202
                        }
                    }
                    _ = fi184
                    _ = v185
                }
            } else {
                if true {
                    v4 -= 42
                }
                v3 -= 127
            }
        }
        _ = fi171
    }
    switch "a" {
    case "foo bar":
        fmt.Println(v4)
    case "\n\t":
        for fi203 := 0; fi203 < 9; fi203++ {
            v3 = (v3 << ((65535) & 63))
            var v204 int32 = -2147483648
            _ = fi203
            _ = v204
        }
    case "":
        var v205 float32 = float32(-2.25)
        _ = v205
    default:
        var v206 string = (("foo bar" + ("hello" + ("foo bar" + ("hello" + ("a" + ""))))) + ("hello" + "foo bar"))
        _ = v206
    }
    switch 4294967295 {
    case -2147483648:
        v207 := v2[("foo bar" + (("hello" + "\n\t") + ("foo bar" + "\n\t")))]
        _ = v207
    case 128:
        if true {
            for fi208 := 0; fi208 < 18; fi208++ {
                fmt.Println(v4, v3)
                fmt.Println(fi208, v4)
                v3 = v3
                _ = fi208
            }
        } else {
            var v209 string = "hello"
            v4 = (v4 * -168)
            _ = v209
        }
    }
    if false {
        v3 = 37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
