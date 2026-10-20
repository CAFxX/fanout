package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int32 }
type GenS2 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64; F1 uint64 }
type GenS7 struct { F0 bool; F1 int64 }
type GenS8 struct { F0 string; F1 uint32 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 string; F1 uint }
type GenS11 struct { F0 uint64; F1 uint8 }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint32; F1 int64 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 uint; F1 uint16; F2 int16 }
type GenS16 struct { F0 int32; F1 bool; F2 int64 }
type GenS17 struct { F0 int64; F1 uint8 }
type GenS18 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS19 struct { F0 int16; F1 uint64 }
func genFn0(p0 uint64, p1 uint64, p2 bool) int32 {
    for fi6 := 0; fi6 < 1; fi6++ {
        p1 = p1
        _ = fi6
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 float64 = 3.14159
        p2 = ((false && (0.0 > 3.14159)) || (false && p2))
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := GenS1{F0: 149, F1: -1, F2: -2147483648}
        _ = fi9
        _ = v10
    }
    v11 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 240
}
func genFn1() int64 {
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 int64 = -579
        for fi14 := 0; fi14 < 7; fi14++ {
            if true {
                if true {
                    v13 = v13
                    var v15 int = fi14
                    v16 := []int{(fi12 >> ((18446744073709551615) & 63)), 32767, (fi14 << ((933) & 63))}
                    var v17 float64 = 0.1
                    _ = v15
                    _ = v16
                    _ = v17
                } else {
                    v13 -= v13
                    if false {
                        v18 := GenS1{F0: 42, F1: 42, F2: -1}
                        _ = v18
                    } else {
                        v13 = v13
                    }
                    v19 := []int{}
                    _ = v19
                }
                if false {
                    v20 := GenS2{F0: 4294967295, F1: false, F2: 2147483648}
                    fmt.Println(fi14, v13, fi14)
                    _ = v20
                } else {
                    var v21 int64 = 0
                    if (!(int64(fi12) != (v21 / (((v13 % ((-9223372036854775808) | 1))) | 1)))) {
                        if (("" <= "") && (!(!(988 > 219)))) {
                            var v22 uint64 = 42
                            _ = v22
                        } else {
                            if false {
                                v21 = v21
                            }
                            var v23 int8 = 1
                            v21 -= v21
                            _ = v23
                        }
                        switch 4294967295 {
                        case -9223372036854775808:
                            fmt.Println(v13, fi14, fi14)
                        case -1:
                            var v24 int64 = v13
                            var v25 uint16 = 1
                            switch "\n\t" {
                            case "":
                                v24 = v21
                                v26 := []int{(fi14 | 42)}
                                _ = v26
                            case "hello":
                                v24 -= v21
                            case "foo bar":
                                v13 = v24
                            }
                            _ = v24
                            _ = v25
                        }
                    }
                    if (false && (false || (4294967295 == 836))) {
                        v13 = v21
                    } else {
                        var v27 int64 = -9223372036854775808
                        v28 := []int{}
                        switch "\n\t" {
                        case "\n\t":
                            if false {
                                v29 := GenS3{F0: 2147483647, F1: (-(v21))}
                                var v30 int8 = -128
                                var v31 bool = true
                                _ = v29
                                _ = v30
                                _ = v31
                            } else {
                                if true {
                                    if len(v28) > 0 {
                                        v28[0] = fi12
                                    }
                                } else {
                                    _ = len(v28)
                                    if (!(2 > 925)) {
                                        v32 := GenS4{F0: 42, F1: 755, F2: genFn0(454, 2, false)}
                                        _ = v32
                                    }
                                    v33 := GenS5{F0: ("foo bar" + "")}
                                    _ = v33
                                }
                            }
                        case "":
                            if len(v28) >= 2 {
                                v28 = v28[1:len(v28)]
                            }
                        case "hello":
                            v28 = append(v28, fi12)
                        default:
                            switch "\n\t" {
                            case "foo bar":
                                v34 := GenS4{F0: 255, F1: 547, F2: genFn0(4294967295, 104, false)}
                                _ = v34
                            }
                        }
                        _ = v27
                        _ = v28
                    }
                    _ = v21
                }
            } else {
                v13 = (v13 + v13)
                v13 = (v13 & -1)
                var v35 float32 = float32(((0.1 / (1.5 / 0.1)) / 1e10))
                _ = v35
            }
            if false {
                if ((v13 | 591) == (v13 - (v13 | 9223372036854775807))) {
                    var v36 int32 = genFn0(734, 18446744073709551615, false)
                    fmt.Println(v36, v36, fi14)
                    _ = v36
                } else {
                    switch fi14 {
                    case -684:
                        var v37 int32 = 0
                        v38 := GenS6{F0: 0, F1: 0}
                        var v39 uint64 = 2
                        _ = v37
                        _ = v38
                        _ = v39
                    case 1:
                        v13 += (^(v13))
                    default:
                        if (0 > 42) {
                            var v40 int8 = -128
                            v13 = -565
                            v41 := GenS4{F0: 0, F1: 65535, F2: -1}
                            _ = v40
                            _ = v41
                        } else {
                            var v42 uint = 255
                            fmt.Println(fi14, v42)
                            _ = v42
                        }
                    }
                    v43 := GenS7{F0: true, F1: int64(fi14)}
                    _ = v43
                }
            } else {
                if false {
                    switch "\n\t" {
                    case "":
                        var v44 int16 = -32768
                        _ = v44
                    case "\n\t":
                        v13 -= (v13 + v13)
                    default:
                        var v45 int32 = -1
                        _ = v45
                    }
                    v46 := GenS8{F0: "a", F1: 478}
                    v47 := GenS7{F0: false, F1: (^(v13))}
                    _ = v46
                    _ = v47
                } else {
                    var v48 int = fi12
                    v49 := []int{(fi12 + (v48 & -2147483648)), 32767, 0}
                    if false {
                        if (304 != 531) {
                            if (("foo bar" + (((((("a" + "") + "a") + "") + (("" + "\n\t") + ("\n\t" + "a"))) + "a") + "hello")) < "a") {
                                if (18446744073709551615 < 42) {
                                    switch "\n\t" {
                                    case "":
                                        var v50 float64 = (0.0 * ((1e10 / 1.5) - 1.5))
                                        _ = v50
                                    case "a":
                                        v48 -= fi12
                                    case "foo bar":
                                        fmt.Println(v48, fi14)
                                    }
                                    var v51 uint16 = 0
                                    _ = v51
                                } else {
                                    v52 := GenS9{F0: 527}
                                    _ = v52
                                }
                            } else {
                                v13 += (-(v13))
                            }
                        } else {
                            var v53 int = (-(fi12))
                            if false {
                                v54 := GenS10{F0: (("\n\t" + "hello") + (("a" + ("\n\t" + "\n\t")) + (("hello" + "foo bar") + ("" + ("foo bar" + ""))))), F1: 1}
                                v13 += v13
                                fmt.Println(v13)
                                var v55 uint32 = 639
                                v48 = (-(fi14))
                                _ = v54
                                _ = v55
                            } else {
                                fmt.Println(v13, fi12)
                                if false {
                                    switch "" {
                                    case "":
                                        var v56 int32 = genFn0(760, 18446744073709551615, (false || ((fi12 << ((694) & 63)) == 32767)))
                                        _ = v56
                                    case "hello":
                                        v13 = v13
                                    case "\n\t":
                                        fmt.Println(v13, v13)
                                    default:
                                        var v57 uint16 = 2
                                        _ = v57
                                    }
                                    if false {
                                        v53 -= v48
                                        if false {
                                            if (42 != 18446744073709551615) {
                                                v13 = v13
                                            }
                                        } else {
                                            if len(v49) > 0 {
                                                v49[0] = (fi14 + (^(fi14)))
                                            }
                                            var v58 uint32 = 904
                                            _ = v58
                                        }
                                    } else {
                                        var v59 uint32 = 901
                                        v49 = append(v49, (v53 % (((fi14 % (((-(v53))) | 1))) | 1)))
                                        _ = v59
                                    }
                                }
                                if (0 == 183) {
                                    var v60 uint8 = 19
                                    switch (fi14 * int(v13)) {
                                    case 302:
                                        fmt.Println(v48, v60, fi12)
                                    case 50:
                                        fmt.Println(v13, v13)
                                    case 4294967295:
                                        fmt.Println(v13, v13)
                                    }
                                    switch (v48 % ((fi14) | 1)) {
                                    case -1:
                                        if (!true) {
                                            var v61 uint8 = (v60 << ((229) & 63))
                                            _ = v61
                                        }
                                    case 9223372036854775807:
                                        var v62 float64 = ((((0.0 / 3.14159) / 1.0) + float64(0.0)) * 0.1)
                                        _ = v62
                                    case -414:
                                        v63 := GenS11{F0: 953, F1: 255}
                                        _ = v63
                                    }
                                    _ = v60
                                } else {
                                    v53 -= fi12
                                }
                                v13 = v13
                                var v64 uint16 = 65535
                                _ = v64
                            }
                            fmt.Println(fi14, fi12)
                            _ = v53
                        }
                        v65 := GenS12{F0: (v13 - (^(v13)))}
                        _ = v65
                    } else {
                        v48 = (fi14 << ((255) & 63))
                        switch (v48 >> ((255) & 63)) {
                        case 4294967295:
                            var v66 uint16 = 42
                            _ = v66
                        }
                        fmt.Println(fi12, fi12, v48)
                    }
                    _ = v48
                    _ = v49
                }
                var v67 int32 = genFn0(54, 135, true)
                if (("a" == "hello") || true) {
                    if (!(0.0 != (float64(0.0) - -2.25))) {
                        if (128 == 158) {
                            v13 = (v13 & v13)
                            if false {
                                v68 := []int{(fi14 % (((fi14 & fi14)) | 1))}
                                var v69 int16 = 32767
                                _ = v68
                                _ = v69
                            }
                            v70 := []int{fi14, -1, (fi12 & fi14), (fi14 & (fi12 * 127))}
                            _ = v70
                        }
                        var v71 int16 = 1
                        v72 := []int{fi14, fi12, (fi12 ^ fi14)}
                        _ = v71
                        _ = v72
                    } else {
                        if true {
                            var v73 uint = 2
                            var v74 int32 = (v67 >> (((v73 / (((^(v73))) | 1))) & 63))
                            var v75 int8 = -1
                            var v76 float64 = 3.14159
                            v13 += 42
                            _ = v73
                            _ = v74
                            _ = v75
                            _ = v76
                        }
                    }
                    if false {
                        if false {
                            v77 := []int{fi14}
                            if (!(true || (false || true))) {
                                if len(v77) >= 2 {
                                    v77 = v77[1:len(v77)]
                                }
                                fmt.Println(fi14, v67, fi14)
                            } else {
                                fmt.Println(fi12)
                                v67 = -572
                                v78 := GenS13{F0: 329, F1: (v13 - v13)}
                                _ = v78
                            }
                            var v79 int32 = v67
                            _ = v77
                            _ = v79
                        }
                        v80 := GenS10{F0: ((("foo bar" + "foo bar") + "foo bar") + "hello"), F1: 255}
                        _ = v80
                    } else {
                        v81 := []int{2147483647, (fi14 * fi14), (fi12 / (((-(fi12))) | 1))}
                        _ = v81
                    }
                    var v82 uint = 677
                    fmt.Println(fi12, fi12, v82)
                    var v83 int64 = (v13 ^ 0)
                    _ = v82
                    _ = v83
                }
                v84 := []int{fi12, 4294967295}
                fmt.Println(v67, v13)
                _ = v67
                _ = v84
            }
            v85 := GenS3{F0: 645, F1: (v13 - 1)}
            _ = fi14
            _ = v85
        }
        _ = fi12
        _ = v13
    }
    var v86 int = -359
    _ = v86
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 426
    var v3 uint64 = 925
    var v4 uint64 = v3
    var v5 int32 = -1
    v3 = 0
    if (!true) {
        fmt.Println(v3, v5, v3)
    } else {
        fmt.Println(v4)
    }
    switch 623 {
    case 1:
        v2 += (v5 / (((v5 & -2147483648)) | 1))
    default:
        var v87 float32 = float32(1e10)
        _ = v87
    }
    for fi88 := 0; fi88 < 19; fi88++ {
        switch "" {
        case "a":
            fmt.Println(fi88)
        case "hello":
            v89 := GenS10{F0: "\n\t", F1: 268}
            _ = v89
        case "\n\t":
            var v90 uint8 = 19
            _ = v90
        }
        for fi91 := 0; fi91 < 8; fi91++ {
            var v92 uint8 = 0
            if (-2.25 == (3.14159 * -2.25)) {
                if ((!(true || true)) && true) {
                    var v93 uint16 = 2
                    _ = v93
                } else {
                    switch fi88 {
                    case -494:
                        var v94 int32 = -221
                        _ = v94
                    case 2147483647:
                        var v95 int8 = 42
                        _ = v95
                    case 0:
                        var v96 uint8 = (^(v92))
                        _ = v96
                    }
                }
            }
            v2 = v2
            _ = fi91
            _ = v92
        }
        v2 = v5
        _ = fi88
    }
    if ((1 < 128) || false) {
        if (1 >= 856) {
            v5 += (v2 & v5)
            v2 = v2
            if false {
                v2 += (v2 * v5)
                v5 -= (^(v5))
            } else {
                v97 := GenS14{F0: 0}
                _ = v97
            }
        } else {
            var v98 int = -9223372036854775808
            delete(v1, "hello")
            _ = v98
        }
        v99 := GenS15{F0: uint(v3), F1: 255, F2: 32767}
        var v100 int8 = 100
        _ = v100
        _ = v99
    }
    v101 := GenS3{F0: 919, F1: genFn1()}
    for fi102 := 0; fi102 < 2; fi102++ {
        switch 127 {
        case 264:
            v103 := GenS13{F0: 134, F1: 0}
            _ = v103
        }
        delete(v1, ("" + "hello"))
        _ = fi102
    }
    _ = len(v1)
    var v104 int32 = v5
    fmt.Println(v2, v3, v104)
    if false {
        switch ((("hello" + "a") + ("hello" + "hello")) + "") {
        case "\n\t":
            if true {
                v105 := []int{588, 560, 848}
                _ = v105
            }
        }
        v2 = (v104 * (v2 >> ((0) & 63)))
        if true {
            for fi106 := 0; fi106 < 9; fi106++ {
                var v107 uint16 = 535
                var v108 int64 = 782
                _ = fi106
                _ = v107
                _ = v108
            }
            v109 := []int{604, -262, -9223372036854775808, 128}
            var v110 uint8 = 42
            _ = v109
            _ = v110
        } else {
            for fi111 := 0; fi111 < 6; fi111++ {
                var v112 bool = (66 == 1)
                v1["foo bar"] = -17
                v4 = (v4 * v3)
                _ = fi111
                _ = v112
            }
            v113 := GenS16{F0: 143, F1: true, F2: 0}
            _ = v113
        }
    } else {
        v114 := []int{1, 128}
        v4 += 4294967295
        _ = len(v1)
        _ = v114
    }
    if true {
        if false {
            v115 := GenS11{F0: 2, F1: 127}
            for fi116 := 0; fi116 < 14; fi116++ {
                var v117 int32 = genFn0((v3 >> ((18446744073709551615) & 63)), v3, ((0.0 >= ((1.5 / (1.5 / -2.25)) + float32((1.5 * (1e10 * 1.5))))) && (-9223372036854775808 > 127)))
                if true {
                    for fi118 := 0; fi118 < 13; fi118++ {
                        if false {
                            v119 := []int{397}
                            if true {
                                switch (fi116 & fi116) {
                                case -1:
                                    v104 += v117
                                case 1:
                                    if (v4 <= (v3 / (((v4 / (((v3 ^ (^(v4)))) | 1))) | 1))) {
                                        v2 += v117
                                        switch (fi116 | 811) {
                                        case -2147483648:
                                            fmt.Println(fi118, v2, v2)
                                        case 128:
                                            if (1 <= -575) {
                                                var v120 int16 = -32768
                                                _ = v120
                                            }
                                        }
                                    } else {
                                        if false {
                                            v4 += (^(v4))
                                        } else {
                                            v119 = append(v119, (fi118 & (fi118 << ((548) & 63))))
                                            v117 -= (v117 * (^(v104)))
                                            v121 := GenS17{F0: 402, F1: 0}
                                            _ = v121
                                        }
                                        v122 := GenS18{F0: 76, F1: 385, F2: -109}
                                        _ = v122
                                    }
                                case 23:
                                    var v123 bool = true
                                    _ = v123
                                }
                                var v124 int32 = (^(v104))
                                v5 -= genFn0(v4, v4, false)
                                v2 = (v124 * (v2 - v117))
                                v125 := GenS2{F0: 42, F1: true, F2: 2}
                                _ = v124
                                _ = v125
                            } else {
                                v4 = v4
                                v126 := v1["\n\t"]
                                _ = v126
                            }
                            _ = v119
                        } else {
                            switch "foo bar" {
                            case "a":
                                fmt.Println(v5)
                            default:
                                fmt.Println(v104, v104)
                            }
                            switch "foo bar" {
                            case "\n\t":
                                v127 := []int{}
                                _ = v127
                            case "hello":
                                if true {
                                    v128 := []int{}
                                    v129 := GenS5{F0: "hello"}
                                    _ = v128
                                    _ = v129
                                }
                                fmt.Println(fi116)
                            default:
                                var v130 int8 = 1
                                _ = v130
                            }
                            v117 = 216
                        }
                        _ = fi118
                    }
                    switch 42 {
                    case -9223372036854775808:
                        v131 := GenS19{F0: 645, F1: (v4 ^ v4)}
                        _ = v131
                    case -1:
                        v117 += -981
                    }
                }
                fmt.Println(v2)
                _ = fi116
                _ = v117
            }
            if (32767 != uint16(v3)) {
                v3 = v3
                fmt.Println(v3)
            }
            _ = v115
        } else {
            _ = len(v1)
            var v132 uint = 4294967295
            var v133 int32 = v2
            _ = v132
            _ = v133
        }
        var v134 int8 = 1
        _ = v134
    } else {
        fmt.Println(v2, v5)
        var v135 int = 0
        _ = v135
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v104:", v104)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v101
    _ = v104
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
