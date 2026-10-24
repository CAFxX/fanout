package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 int64; F1 int64; F2 uint64 }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 int; F1 bool }
type GenS7 struct { F0 uint32; F1 uint }
type GenS8 struct { F0 int; F1 uint }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint32; F1 int16 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int; F1 uint; F2 int }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int }
func genFn0() bool {
    v5 := []int{128, 32768}
    switch "foo bar" {
    case "a":
        v5 = append(v5, 32768)
    default:
        switch 284 {
        case 128:
            v6 := GenS1{F0: ("a" + "hello"), F1: 9223372036854775807, F2: 18446744073709551615}
            _ = v6
        case 545:
            var v7 int = -223
            _ = v7
        }
    }
    var v8 int16 = -1
    var v9 int64 = -1
    _ = v5
    _ = v8
    _ = v9
    return false
}
func genFn1(p0 int, p1 uint32) uint64 {
    switch "foo bar" {
    case "":
        var v10 int64 = 0
        _ = v10
    default:
        p1 = (p1 / ((p1) | 1))
    }
    p0 = -9223372036854775808
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 uint = 2
    v2 := make(map[int]int)
    var v3 float64 = 0.0
    v4 := make(map[uint32]bool)
    for fi11 := 0; fi11 < 11; fi11++ {
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 int = fi11
            var v14 uint16 = 174
            v15 := []int{(fi12 | (fi11 ^ (fi11 ^ fi11))), (fi11 * (v13 | (-(v13)))), 587, -2147483648}
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi11
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        v3 = float64(float32(v3))
        if true {
            var v17 uint64 = genFn1(fi16, 65535)
            var v18 float32 = 0.1
            if (false && true) {
                var v19 string = "a"
                v20 := []int{(fi16 / ((-931) | 1)), 42, 1}
                var v21 bool = true
                _ = v19
                _ = v20
                _ = v21
            } else {
                v22 := GenS1{F0: "a", F1: int64(v17), F2: v1}
                switch "hello" {
                case "foo bar":
                    for fi23 := 0; fi23 < 12; fi23++ {
                        if false {
                            v18 += v18
                            var v24 uint32 = 633
                            v25 := []int{(fi16 ^ fi16)}
                            var v26 int32 = 0
                            fmt.Println(v17, v18, v1)
                            v26 = (^(v26))
                            if false {
                                v27 := GenS2{F0: (v26 | v26), F1: 373}
                                var v28 float64 = ((1.5 - ((v3 * float64((v18 + 0.0))) / -2.25)) + (v3 + ((float64(v18) - v3) + 1e10)))
                                _ = v27
                                _ = v28
                            }
                            switch (fi23 * (fi16 | (fi23 << ((v1) & 63)))) {
                            case 9223372036854775807:
                                if true {
                                    var v29 string = ("\n\t" + ("foo bar" + "\n\t"))
                                    _ = v29
                                } else {
                                    v17 = (v17 & (v17 - (^(v17))))
                                    v30 := GenS1{F0: "a", F1: 1, F2: v1}
                                    if len(v25) >= 2 {
                                        v25 = v25[1:len(v25)]
                                    }
                                    v31 := GenS1{F0: (("a" + "a") + ""), F1: 186, F2: v1}
                                    _ = v30
                                    _ = v31
                                }
                            }
                            var v32 int64 = 1
                            _ = v24
                            _ = v25
                            _ = v26
                            _ = v32
                        } else {
                            var v33 float32 = ((v18 + v18) + float32(v3))
                            _ = v33
                        }
                        v34 := GenS1{F0: "hello", F1: -9223372036854775808, F2: (v1 - v1)}
                        if true {
                            v35 := []int{(fi23 - fi23), 963, (^(fi23))}
                            _ = v35
                        }
                        _ = fi23
                        _ = v34
                    }
                default:
                    v36 := v4[0]
                    _ = v36
                }
                _ = v22
            }
            _ = v17
            _ = v18
        }
        switch ("foo bar" + "hello") {
        case "":
            switch "a" {
            case "":
                v37 := v2[(fi16 * 1)]
                _ = v37
            case "\n\t":
                var v38 int16 = 1000
                _ = v38
            case "hello":
                v39 := []int{(fi16 ^ fi16)}
                _ = v39
            default:
                v1 = v1
            }
        default:
            for fi40 := 0; fi40 < 17; fi40++ {
                v1 += (^(v1))
                _ = fi40
            }
        }
        _ = fi16
    }
    v1 += (v1 / ((v1) | 1))
    switch 32767 {
    case 127:
        switch -1 {
        case -830:
            if true {
                if false {
                    var v41 int = -700
                    for fi42 := 0; fi42 < 7; fi42++ {
                        for fi43 := 0; fi43 < 17; fi43++ {
                            v41 = (fi42 / ((v41) | 1))
                            v44 := GenS2{F0: -1, F1: 42}
                            if false {
                                v45 := GenS3{F0: -2147483648, F1: 32767}
                                _ = v45
                            }
                            if true {
                                fmt.Println(v3, v3)
                                v46 := v4[4294967295]
                                var v47 uint8 = 255
                                _ = v46
                                _ = v47
                            } else {
                                v3 = v3
                                fmt.Println(v3, v3, fi43)
                                if true {
                                    var v48 uint8 = 2
                                    v3 -= 1.5
                                    v49 := GenS4{F0: 42, F1: 678, F2: 772}
                                    _ = v48
                                    _ = v49
                                }
                            }
                            fmt.Println(fi43)
                            switch fi42 {
                            case -81:
                                var v50 int = fi43
                                _ = v50
                            case -9223372036854775808:
                                switch 32768 {
                                case -2147483648:
                                    var v51 float64 = (float64(0.0) * 1e10)
                                    _ = v51
                                case 0:
                                    var v52 uint16 = 0
                                    _ = v52
                                case -9223372036854775808:
                                    var v53 float32 = -2.25
                                    _ = v53
                                default:
                                    switch 4294967295 {
                                    case 128:
                                        v41 -= fi43
                                    case -848:
                                        v3 = (0.1 - v3)
                                    case 655:
                                        v54 := GenS4{F0: 0, F1: 1, F2: 642}
                                        var v55 bool = (false || (false && (false && false)))
                                        switch fi43 {
                                        case 2147483647:
                                            v41 = (fi42 % ((v41) | 1))
                                        case 128:
                                            var v56 int32 = -2147483648
                                            var v57 float64 = ((0.0 + ((((1.5 + (v3 + (v3 * v3))) - (((1.5 * (v3 + v3)) * ((v3 / (v3 * v3)) * v3)) + (v3 / v3))) * (1e10 + (1.5 - v3))) / v3)) / v3)
                                            v58 := GenS5{F0: (("hello" + "\n\t") + ""), F1: 255}
                                            _ = v56
                                            _ = v57
                                            _ = v58
                                        default:
                                            if v55 {
                                                fmt.Println(v3, v1, fi43)
                                                v55 = (-1 == -2147483648)
                                                delete(v2, v41)
                                                var v59 float64 = v3
                                                switch "\n\t" {
                                                case "\n\t":
                                                    var v60 float64 = v59
                                                    _ = v60
                                                default:
                                                    v61 := GenS6{F0: fi43, F1: v55}
                                                    v3 -= v3
                                                    var v62 int16 = int16(v1)
                                                    _ = v61
                                                    _ = v62
                                                }
                                                v63 := []int{v41, (v41 + (^(fi42)))}
                                                fmt.Println(v41, v3)
                                                var v64 int32 = -86
                                                _ = v59
                                                _ = v63
                                                _ = v64
                                            }
                                        }
                                        _ = v54
                                        _ = v55
                                    default:
                                        if (((42 < 42) || true) || false) {
                                            if false {
                                                var v65 uint64 = 4294967295
                                                switch (fi42 / ((fi42) | 1)) {
                                                case -1:
                                                    v41 -= v41
                                                }
                                                v41 -= 2147483647
                                                fmt.Println(v1)
                                                _ = v65
                                            }
                                            var v66 int16 = int16(v1)
                                            _ = v66
                                        } else {
                                            v67 := v2[v41]
                                            fmt.Println(v41, fi42)
                                            fmt.Println(v3, v1)
                                            _ = v67
                                        }
                                    }
                                }
                                v68 := GenS1{F0: "hello", F1: -969, F2: (v1 + (v1 >> ((0) & 63)))}
                                _ = v68
                            }
                            if true {
                                _ = len(v2)
                                if ((v41 ^ (v41 - v41)) >= (fi42 - -9223372036854775808)) {
                                    fmt.Println(fi43, fi43)
                                    var v69 uint8 = 0
                                    if (-923 > -9223372036854775808) {
                                        if true {
                                            var v70 int64 = 42
                                            switch ((("a" + ("a" + "a")) + (("a" + "a") + ("" + "\n\t"))) + "foo bar") {
                                            case "\n\t":
                                                v71 := GenS5{F0: ((("a" + "") + (("foo bar" + ("" + "foo bar")) + "a")) + "\n\t"), F1: v1}
                                                _ = v71
                                            case "a":
                                                fmt.Println(v70)
                                            }
                                            delete(v4, 1)
                                            switch ("" + "a") {
                                            case "a":
                                                if (!true) {
                                                    if (!(false || ((1.5 - float32(v3)) < (0.0 / 1.0)))) {
                                                        var v72 int32 = 0
                                                        _ = v72
                                                    } else {
                                                        var v73 int8 = 100
                                                        v1 = 18446744073709551615
                                                        switch ("" + "a") {
                                                        case "hello":
                                                            if true {
                                                                var v74 float64 = (v3 / v3)
                                                                _ = len(v2)
                                                                if (int64(v1) != (v70 * (v70 % ((v70) | 1)))) {
                                                                    if false {
                                                                        v4[482] = (!false)
                                                                    } else {
                                                                        var v75 float64 = v74
                                                                        _ = v75
                                                                    }
                                                                    _ = len(v2)
                                                                    if false {
                                                                        v74 -= (v74 + 1.5)
                                                                        if (!(true && (uint8(v1) != 1))) {
                                                                            v76 := GenS7{F0: 0, F1: v1}
                                                                            v77 := []int{}
                                                                            _ = v76
                                                                            _ = v77
                                                                        } else {
                                                                            v1 += 263
                                                                            v78 := GenS3{F0: -1, F1: 32767}
                                                                            _ = v78
                                                                        }
                                                                        switch 127 {
                                                                        case 128:
                                                                            var v79 uint64 = 2
                                                                            _ = v79
                                                                        case 32768:
                                                                            switch v41 {
                                                                            case -937:
                                                                                fmt.Println(v70)
                                                                            }
                                                                        }
                                                                        var v80 uint32 = 1
                                                                        _ = v80
                                                                    }
                                                                } else {
                                                                    var v81 float64 = float64(0.0)
                                                                    v3 = (3.14159 / 0.1)
                                                                    v81 += 0.0
                                                                    _ = v81
                                                                }
                                                                _ = v74
                                                            } else {
                                                                v73 = v73
                                                                v82 := GenS5{F0: "a", F1: 18446744073709551615}
                                                                _ = v82
                                                            }
                                                        case "\n\t":
                                                            if (!true) {
                                                                v83 := []int{(fi43 >> ((255) & 63)), 684, v41}
                                                                if len(v83) > 0 {
                                                                    v83[0] = 42
                                                                }
                                                                _ = v83
                                                            }
                                                            v84 := []int{v41, 1, (v41 * 1)}
                                                            _ = v84
                                                        case "a":
                                                            switch ("" + "a") {
                                                            case "foo bar":
                                                                v85 := []int{v41, fi42, 32767}
                                                                _ = v85
                                                            case "":
                                                                switch 32767 {
                                                                case 127:
                                                                    var v86 uint32 = 42
                                                                    _ = v86
                                                                case 1:
                                                                    v69 -= 128
                                                                case -923:
                                                                    if false {
                                                                        v3 += v3
                                                                        v69 -= (v69 & (v69 + 127))
                                                                    } else {
                                                                        var v87 int8 = 1
                                                                        var v88 int64 = 424
                                                                        _ = v87
                                                                        _ = v88
                                                                    }
                                                                default:
                                                                    v70 -= 9223372036854775807
                                                                }
                                                            default:
                                                                if true {
                                                                    fmt.Println(fi42, v41, v69)
                                                                } else {
                                                                    fmt.Println(v73)
                                                                    v89 := GenS8{F0: (fi43 - (fi43 + (fi43 | (fi42 >> (((v1 * v1)) & 63))))), F1: 994}
                                                                    switch ("hello" + "foo bar") {
                                                                    case "\n\t":
                                                                        v69 += 128
                                                                    case "hello":
                                                                        var v90 uint8 = (v69 % (((^(v69))) | 1))
                                                                        _ = v90
                                                                    default:
                                                                        fmt.Println(v41)
                                                                    }
                                                                    _ = v89
                                                                }
                                                                fmt.Println(v70)
                                                            }
                                                        default:
                                                            v41 = fi42
                                                        }
                                                        _ = v73
                                                    }
                                                    if true {
                                                        v1 -= v1
                                                        if ((0 <= v1) || (!(-61 != -128))) {
                                                            v91 := GenS8{F0: v41, F1: 4294967295}
                                                            _ = v91
                                                        } else {
                                                            v92 := []int{(v41 << (((v1 / (((v1 ^ v1)) | 1))) & 63))}
                                                            var v93 int32 = -2147483648
                                                            v70 += (-(v70))
                                                            v70 += (v70 & -797)
                                                            _ = v92
                                                            _ = v93
                                                        }
                                                    }
                                                    var v94 float32 = 1.5
                                                    _ = v94
                                                } else {
                                                    v41 += 127
                                                    var v95 bool = (false || false)
                                                    _ = v95
                                                }
                                                v96 := GenS9{F0: (v70 & v70)}
                                                fmt.Println(v69, fi43)
                                                _ = v96
                                            case "\n\t":
                                                v97 := GenS9{F0: v70}
                                                _ = v97
                                            default:
                                                v98 := []int{-6}
                                                _ = v98
                                            }
                                            v99 := []int{}
                                            v4[2] = (false || true)
                                            _ = v70
                                            _ = v99
                                        } else {
                                            fmt.Println(fi43, v69)
                                            var v100 int64 = -1
                                            _ = v100
                                        }
                                        if false {
                                            if true {
                                                v101 := GenS1{F0: "", F1: -9223372036854775808, F2: 2}
                                                var v102 float32 = 1e10
                                                _ = v101
                                                _ = v102
                                            } else {
                                                v41 = 128
                                                var v103 int16 = -198
                                                delete(v4, 400)
                                                v103 = (v103 << ((v1) & 63))
                                                fmt.Println(v41, v41, v103)
                                                _ = v103
                                            }
                                        }
                                        if false {
                                            v2[(fi42 + (^(fi43)))] = (v41 >> ((v1) & 63))
                                            if true {
                                                v104 := GenS10{F0: 98, F1: -857}
                                                _ = v104
                                            } else {
                                                var v105 int8 = 127
                                                _ = v105
                                            }
                                            v106 := []int{32768, (fi42 & (-(fi42)))}
                                            _ = v106
                                        } else {
                                            v107 := GenS7{F0: 4294967295, F1: (^(v1))}
                                            var v108 string = ("a" + "foo bar")
                                            var v109 float64 = v3
                                            _ = v107
                                            _ = v108
                                            _ = v109
                                        }
                                        fmt.Println(v3, fi43, fi43)
                                    } else {
                                        var v110 float64 = (v3 / 1.5)
                                        var v111 uint32 = 2147483647
                                        var v112 int32 = -2147483648
                                        _ = v110
                                        _ = v111
                                        _ = v112
                                    }
                                    v113 := GenS11{F0: 32767}
                                    v69 += (^(v69))
                                    var v114 bool = (100 == 127)
                                    fmt.Println(v114)
                                    v115 := GenS11{F0: 669}
                                    _ = v113
                                    _ = v114
                                    _ = v115
                                    _ = v69
                                } else {
                                    switch (fi43 << ((v1) & 63)) {
                                    case 1:
                                        fmt.Println(v3, v41)
                                        if (!false) {
                                            var v116 string = ""
                                            v3 += v3
                                            _ = v116
                                        } else {
                                            v41 = v41
                                        }
                                        v117 := v4[445]
                                        v41 += v41
                                        if ((477 != 42) && ((!v117) || v117)) {
                                            if (!(42 > 1)) {
                                                fmt.Println(v3, v41)
                                                var v118 uint64 = genFn1((v41 + 4294967295), 1)
                                                v119 := []int{9223372036854775807}
                                                _ = v118
                                                _ = v119
                                            }
                                        } else {
                                            var v120 uint16 = 1
                                            _ = v120
                                        }
                                        fmt.Println(v3, fi43, fi43)
                                        v121 := GenS12{F0: fi42, F1: v1, F2: fi43}
                                        var v122 uint = (^(v1))
                                        _ = v117
                                        _ = v121
                                        _ = v122
                                    default:
                                        fmt.Println(v41, v41, fi42)
                                    }
                                    var v123 uint = v1
                                    var v124 float32 = -2.25
                                    _ = v123
                                    _ = v124
                                }
                            } else {
                                switch v41 {
                                case 468:
                                    var v125 uint = v1
                                    _ = v125
                                case 32767:
                                    if ((127 != 255) || true) {
                                        var v126 int64 = 9223372036854775807
                                        v41 = (v41 & 577)
                                        v41 += (fi43 & fi42)
                                        _ = v126
                                    } else {
                                        if (!false) {
                                            v127 := GenS12{F0: fi43, F1: (v1 * (v1 | 208)), F2: fi43}
                                            _ = v127
                                        }
                                        if (((("\n\t" + "hello") + ("a" + ("hello" + (("hello" + "\n\t") + "")))) < "") && (968 <= 65535)) {
                                            switch fi43 {
                                            case 8:
                                                delete(v2, (fi43 | (^(fi43))))
                                            case 1:
                                                var v128 float64 = v3
                                                _ = v128
                                            default:
                                                v129 := GenS13{F0: 1}
                                                _ = v129
                                            }
                                        }
                                        var v130 int16 = 1
                                        _ = v130
                                    }
                                case 9223372036854775807:
                                    if ((100 != 85) || false) {
                                        v131 := []int{}
                                        v3 = (((1.5 * (-2.25 / (((v3 * (1.5 * (v3 / v3))) / ((v3 - v3) + ((0.1 - (3.14159 + v3)) * (v3 - (v3 * 1e10))))) / (((3.14159 * (1.5 * v3)) * (v3 * -2.25)) - 0.0)))) - (0.0 / v3)) / (v3 * v3))
                                        _ = v131
                                    } else {
                                        v1 = v1
                                        fmt.Println(fi43, fi42, v1)
                                    }
                                default:
                                    fmt.Println(fi42, v3)
                                }
                            }
                            _ = fi43
                            _ = v44
                        }
                        v3 = (v3 / 1e10)
                        for fi132 := 0; fi132 < 8; fi132++ {
                            v41 = (v41 ^ (v41 & -9223372036854775808))
                            v41 += fi132
                            _ = fi132
                        }
                        _ = fi42
                    }
                    _ = v41
                }
            } else {
                var v133 int64 = -1
                fmt.Println(v1, v133, v1)
                for fi134 := 0; fi134 < 13; fi134++ {
                    v1 = (v1 % ((v1) | 1))
                    switch -1 {
                    case 9223372036854775807:
                        if false {
                            var v135 uint = (v1 >> ((v1) & 63))
                            _ = v135
                        }
                    case 32767:
                        var v136 float64 = 0.0
                        _ = v136
                    case 32768:
                        for fi137 := 0; fi137 < 16; fi137++ {
                            var v138 string = "foo bar"
                            var v139 int8 = 127
                            _ = fi137
                            _ = v138
                            _ = v139
                        }
                    default:
                        if true {
                            v1 -= (v1 >> ((875) & 63))
                        } else {
                            if ("" == ("hello" + ("foo bar" + ("a" + ("foo bar" + "a"))))) {
                                v133 = v133
                            }
                            v140 := GenS10{F0: 2147483647, F1: 100}
                            _ = v140
                        }
                    }
                    v141 := []int{}
                    _ = fi134
                    _ = v141
                }
                _ = v133
            }
        case 42:
            switch ("foo bar" + ("\n\t" + "hello")) {
            case "hello":
                var v142 uint64 = 4294967295
                _ = v142
            case "\n\t":
                for fi143 := 0; fi143 < 8; fi143++ {
                    var v144 uint = (v1 ^ v1)
                    v144 = v144
                    fmt.Println(v3, v3, fi143)
                    _ = fi143
                    _ = v144
                }
            default:
                switch ("a" + "") {
                case "a":
                    v1 -= 1
                case "foo bar":
                    var v145 string = "hello"
                    _ = v145
                case "":
                    switch "foo bar" {
                    case "foo bar":
                        switch 545 {
                        case -2147483648:
                            v1 = 255
                        case 907:
                            v146 := GenS8{F0: int(v1), F1: (v1 % ((42) | 1))}
                            _ = v146
                        default:
                            switch 1 {
                            case -2147483648:
                                fmt.Println(v1)
                            case 4294967295:
                                if (1 != 0) {
                                    v3 = 3.14159
                                }
                            case 127:
                                var v147 bool = true
                                _ = v147
                            }
                        }
                    case "hello":
                        fmt.Println(v3, v3)
                    }
                default:
                    v148 := GenS14{F0: -448}
                    _ = v148
                }
            }
        default:
            v3 = v3
        }
    }
    var v149 bool = true
    for fi150 := 0; fi150 < 12; fi150++ {
        var v151 int8 = -21
        switch "\n\t" {
        case "":
            var v152 uint8 = 127
            _ = v152
        }
        _ = fi150
        _ = v151
    }
    fmt.Println(v3, v149)
    var v153 int = 9223372036854775807
    fmt.Println(v153, v149)
    var v154 int16 = -63
    fmt.Println("v1:", v1)
    fmt.Println("v149:", v149)
    fmt.Println("v153:", v153)
    fmt.Println("v154:", v154)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v149
    _ = v153
    _ = v154
    _ = v2
    _ = v3
    _ = v4
}
