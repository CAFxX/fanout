package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int32 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 bool; F1 uint16; F2 int64 }
type GenS4 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS5 struct { F0 string; F1 int; F2 int16 }
type GenS6 struct { F0 int32; F1 bool }
type GenS7 struct { F0 uint32; F1 bool }
type GenS8 struct { F0 uint32; F1 int32; F2 bool }
type GenS9 struct { F0 uint; F1 int32; F2 int }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int8; F1 int32 }
type GenS12 struct { F0 uint32; F1 uint16 }
func genFn0() uint32 {
    v4 := []int{400, 650, -9223372036854775808}
    var v5 int32 = -2147483648
    _ = len(v4)
    fmt.Println(v5)
    _ = v4
    _ = v5
    return 608
}
func genFn1(p0 uint32) int8 {
    v6 := GenS1{F0: 42, F1: "hello", F2: 2147483647}
    for fi7 := 0; fi7 < 3; fi7++ {
        p0 -= p0
        v8 := []int{(fi7 * fi7), (fi7 * fi7)}
        _ = fi7
        _ = v8
    }
    p0 -= genFn0()
    switch 54 {
    case -479:
        var v9 uint16 = 225
        _ = v9
    default:
        var v10 uint8 = 0
        _ = v10
    }
    var v11 uint32 = (p0 ^ (p0 >> ((0) & 63)))
    _ = p0
    _ = v11
    _ = v6
    return -1
}

func main() {
    var v1 int32 = 0
    var v2 uint16 = 638
    var v3 uint32 = 320
    var v12 int8 = 127
    if (!(894 != 9223372036854775807)) {
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := []int{fi13, fi13}
            if (v2 <= v2) {
                v15 := GenS1{F0: 248, F1: "\n\t", F2: (v1 % ((v1) | 1))}
                _ = v15
            } else {
                if len(v14) > 0 {
                    v14[0] = (fi13 & fi13)
                }
                if false {
                    v16 := GenS2{F0: -9223372036854775808, F1: (fi13 * 0)}
                    _ = v16
                } else {
                    fmt.Println(v12, v1, v1)
                    v2 += (v2 | v2)
                }
                v3 = v3
            }
            _ = fi13
            _ = v14
        }
        v2 -= 0
    }
    var v17 uint64 = 373
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 bool = (42 < v2)
        for fi20 := 0; fi20 < 11; fi20++ {
            v1 = v1
            var v21 float64 = (0.1 - (-2.25 / ((float64((1.5 - ((float32(0.1) - 0.1) - float32(float64(3.14159))))) * 0.1) + float64(float32(-2.25)))))
            var v22 int32 = v1
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi18
        _ = v19
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        for fi24 := 0; fi24 < 16; fi24++ {
            if (genFn0() != 2147483647) {
                fmt.Println(v1, v17, fi24)
                v25 := GenS2{F0: 675, F1: fi23}
                _ = v25
            }
            _ = fi24
        }
        _ = fi23
    }
    if (!((v12 + (v12 - (v12 & genFn1(genFn0())))) > v12)) {
        switch "a" {
        case "a":
            v2 -= v2
        case "":
            fmt.Println(v1, v1)
        case "foo bar":
            for fi26 := 0; fi26 < 7; fi26++ {
                for fi27 := 0; fi27 < 6; fi27++ {
                    var v28 string = "a"
                    _ = fi27
                    _ = v28
                }
                for fi29 := 0; fi29 < 11; fi29++ {
                    if ((-2.25 == -2.25) || false) {
                        v12 = int8(fi26)
                        var v30 float64 = ((1e10 + -2.25) / 0.1)
                        var v31 int32 = (v1 << ((1) & 63))
                        _ = v30
                        _ = v31
                    }
                    if ((^(v2)) <= (v2 ^ v2)) {
                        switch "a" {
                        case "":
                            var v32 float64 = ((1.5 + (-2.25 - (1e10 * float64(0.1)))) * 0.0)
                            var v33 float64 = ((v32 / 1e10) - v32)
                            switch (fi29 / (((fi29 * 4294967295)) | 1)) {
                            case -9223372036854775808:
                                v34 := []int{}
                                v2 = (v2 << ((4294967295) & 63))
                                _ = v34
                            case 348:
                                v17 -= v17
                                var v35 bool = ((("\n\t" + "hello") + "foo bar") > ("a" + "hello"))
                                var v36 float64 = 0.1
                                _ = v35
                                _ = v36
                            case -412:
                                switch fi26 {
                                case 976:
                                    var v37 float64 = v32
                                    _ = v37
                                case 1:
                                    fmt.Println(v3, v12, fi29)
                                case 912:
                                    if true {
                                        var v38 uint32 = (v3 - (v3 % (((^(v3))) | 1)))
                                        fmt.Println(v33)
                                        v39 := []int{4294967295, (fi29 ^ fi29), fi26, int(v3)}
                                        _ = v38
                                        _ = v39
                                    }
                                }
                            default:
                                v40 := []int{}
                                _ = v40
                            }
                            _ = v32
                            _ = v33
                        case "hello":
                            if (v12 > genFn1(v3)) {
                                var v41 uint32 = 2147483648
                                var v42 bool = ((0 <= 1) || true)
                                var v43 bool = (1.5 != 0.0)
                                switch fi26 {
                                case 580:
                                    v44 := GenS1{F0: 988, F1: "", F2: v1}
                                    _ = v44
                                }
                                _ = v41
                                _ = v42
                                _ = v43
                            }
                            if (-139 >= (fi29 - fi26)) {
                                switch (fi29 * 32768) {
                                case 32768:
                                    fmt.Println(v1, v2, v17)
                                case 9223372036854775807:
                                    if true {
                                        var v45 uint = 331
                                        v2 = (v2 << (((v45 / ((v45) | 1))) & 63))
                                        fmt.Println(v12)
                                        _ = v45
                                    } else {
                                        v46 := []int{fi26, fi26}
                                        _ = v46
                                    }
                                case -388:
                                    if false {
                                        var v47 string = ""
                                        if true {
                                            if false {
                                                var v48 uint = 293
                                                switch fi26 {
                                                case -720:
                                                    if true {
                                                        if (!((v47 <= ("\n\t" + v47)) && true)) {
                                                            var v49 float64 = ((-2.25 + -2.25) / 1e10)
                                                            var v50 uint64 = v17
                                                            _ = v49
                                                            _ = v50
                                                        }
                                                        var v51 string = ((("a" + ((v47 + "a") + (("foo bar" + v47) + (v47 + v47)))) + "\n\t") + "")
                                                        v52 := []int{-293, (^(fi29)), fi29, 9223372036854775807}
                                                        _ = v51
                                                        _ = v52
                                                    } else {
                                                        v53 := GenS1{F0: v48, F1: v47, F2: int32(v12)}
                                                        if false {
                                                            var v54 string = v47
                                                            _ = v54
                                                        } else {
                                                            if false {
                                                                fmt.Println(v17, v48)
                                                                var v55 string = (v47 + v47)
                                                                _ = v55
                                                            }
                                                            var v56 uint32 = v3
                                                            var v57 float64 = 0.0
                                                            v58 := GenS1{F0: (v48 * 18446744073709551615), F1: ((v47 + (("hello" + (v47 + v47)) + "hello")) + (v47 + v47)), F2: (-(v1))}
                                                            v59 := []int{}
                                                            _ = v56
                                                            _ = v57
                                                            _ = v58
                                                            _ = v59
                                                        }
                                                        var v60 float32 = 3.14159
                                                        _ = v53
                                                        _ = v60
                                                    }
                                                case -549:
                                                    v48 -= 2
                                                }
                                                v61 := []int{}
                                                v62 := GenS1{F0: uint(v2), F1: "foo bar", F2: v1}
                                                var v63 int64 = 9223372036854775807
                                                _ = v48
                                                _ = v61
                                                _ = v62
                                                _ = v63
                                            } else {
                                                v17 = (^(v17))
                                                var v64 bool = (1 > 419)
                                                if false {
                                                    if (v64 && false) {
                                                        var v65 uint8 = 11
                                                        fmt.Println(fi26, v47, v47)
                                                        _ = v65
                                                    }
                                                }
                                                _ = v64
                                            }
                                            var v66 bool = ((!(v1 <= (v1 % ((-2147483648) | 1)))) && (v12 == v12))
                                            if ((v17 >> ((18446744073709551615) & 63)) < v17) {
                                                v67 := []int{(fi29 | fi26)}
                                                v2 += v2
                                                fmt.Println(fi29, v66, fi29)
                                                _ = v67
                                            } else {
                                                v68 := GenS2{F0: -828, F1: (fi26 % ((fi26) | 1))}
                                                v69 := GenS2{F0: int64(v12), F1: (fi29 << ((42) & 63))}
                                                v70 := GenS1{F0: 18446744073709551615, F1: (v47 + v47), F2: -1}
                                                _ = v68
                                                _ = v69
                                                _ = v70
                                            }
                                            _ = v66
                                        } else {
                                            var v71 int64 = 559
                                            var v72 int32 = 1
                                            _ = v71
                                            _ = v72
                                        }
                                        _ = v47
                                    }
                                default:
                                    var v73 uint32 = v3
                                    _ = v73
                                }
                                v17 += (v17 >> ((2) & 63))
                            }
                        case "\n\t":
                            if (true && (0.1 > 0.1)) {
                                if (("a" + "") > "") {
                                    if ((!((!true) || (v2 > (v2 + 0)))) && false) {
                                        if ((!(262 >= -1)) || true) {
                                            v74 := []int{}
                                            _ = v74
                                        } else {
                                            var v75 uint8 = 0
                                            _ = v75
                                        }
                                        if true {
                                            v76 := GenS2{F0: 42, F1: fi29}
                                            _ = v76
                                        }
                                        switch fi29 {
                                        case -2147483648:
                                            fmt.Println(fi26, fi26)
                                        case 128:
                                            if (0.0 == 3.14159) {
                                                if (!false) {
                                                    fmt.Println(v1, v2)
                                                    if false {
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            switch fi29 {
                                                            case -57:
                                                                v77 := GenS2{F0: 0, F1: fi26}
                                                                _ = v77
                                                            case -1:
                                                                var v78 float64 = 1e10
                                                                _ = v78
                                                            case 32768:
                                                                v3 -= (v3 * v3)
                                                            default:
                                                                var v79 uint16 = 0
                                                                _ = v79
                                                            }
                                                        default:
                                                            var v80 string = ""
                                                            _ = v80
                                                        }
                                                    } else {
                                                        v81 := GenS1{F0: 18446744073709551615, F1: "a", F2: 2147483647}
                                                        v1 -= int32(v2)
                                                        _ = v81
                                                    }
                                                }
                                                if (!(false && (v12 < (v12 + (v12 + v12))))) {
                                                    var v82 int16 = -1
                                                    v83 := GenS1{F0: uint(v17), F1: "\n\t", F2: v1}
                                                    var v84 int = fi29
                                                    switch -1 {
                                                    case 42:
                                                        switch (v84 >> ((407) & 63)) {
                                                        case 42:
                                                            if false {
                                                                v85 := []int{}
                                                                v86 := GenS3{F0: true, F1: (v2 | (v2 + v2)), F2: 1}
                                                                if (!true) {
                                                                    var v87 bool = (!(false || true))
                                                                    _ = v87
                                                                }
                                                                _ = v85
                                                                _ = v86
                                                            }
                                                        case -9223372036854775808:
                                                            v88 := []int{1, v84, -2147483648, 762}
                                                            _ = v88
                                                        }
                                                    default:
                                                        if false {
                                                            var v89 uint = 959
                                                            _ = v89
                                                        } else {
                                                            var v90 uint8 = 255
                                                            var v91 bool = true
                                                            switch "a" {
                                                            case "foo bar":
                                                                var v92 uint = 42
                                                                _ = v92
                                                            case "a":
                                                                var v93 int = (-(fi29))
                                                                _ = v93
                                                            case "hello":
                                                                switch -1 {
                                                                case 1:
                                                                    if v91 {
                                                                        if (4294967295 <= 0) {
                                                                            if (!true) {
                                                                                var v94 uint16 = (v2 * v2)
                                                                                _ = v94
                                                                            }
                                                                        } else {
                                                                            v95 := GenS4{F0: -1, F1: v90, F2: 139}
                                                                            _ = v95
                                                                        }
                                                                    } else {
                                                                        var v96 uint = 65535
                                                                        if false {
                                                                            v1 += v1
                                                                            if ("\n\t" > "a") {
                                                                                var v97 uint8 = v90
                                                                                _ = v97
                                                                            }
                                                                        }
                                                                        v1 = int32(v84)
                                                                        _ = v96
                                                                    }
                                                                case 128:
                                                                    var v98 int64 = -1
                                                                    _ = v98
                                                                case 4294967295:
                                                                    var v99 uint64 = (^(v17))
                                                                    _ = v99
                                                                default:
                                                                    switch 2147483647 {
                                                                    case 32768:
                                                                        var v100 int = fi29
                                                                        fmt.Println(v84)
                                                                        var v101 int = (fi26 >> ((685) & 63))
                                                                        _ = v100
                                                                        _ = v101
                                                                    default:
                                                                        var v102 int64 = 42
                                                                        _ = v102
                                                                    }
                                                                }
                                                            }
                                                            if (0 <= -9223372036854775808) {
                                                                if ((v12 % (((v12 | (-(v12)))) | 1)) > v12) {
                                                                    v17 -= 2
                                                                    switch "foo bar" {
                                                                    case "":
                                                                        fmt.Println(v91)
                                                                    case "a":
                                                                        switch fi26 {
                                                                        case -281:
                                                                            var v103 int16 = (v82 ^ v82)
                                                                            _ = v103
                                                                        default:
                                                                            var v104 int8 = -45
                                                                            _ = v104
                                                                        }
                                                                    default:
                                                                        if (((((1.5 * (0.0 * (0.0 + 0.1))) * float32((float64(float32(-2.25)) + 1.5))) < ((float32((-2.25 / 1.5)) + 1e10) / 1.0)) && v91) || (((false && v91) || v91) && v91)) {
                                                                            v105 := GenS5{F0: (("\n\t" + "hello") + ("hello" + "foo bar")), F1: (^(v84)), F2: (v82 * (v82 % ((v82) | 1)))}
                                                                            if (3.14159 > (0.0 + (3.14159 + (3.14159 / 3.14159)))) {
                                                                                var v106 bool = true
                                                                                v1 = v1
                                                                                _ = v106
                                                                            } else {
                                                                                fmt.Println(v90, v90, v82)
                                                                                v107 := []int{fi29, fi26}
                                                                                _ = v107
                                                                            }
                                                                            _ = v105
                                                                        }
                                                                    }
                                                                    var v108 uint16 = (v2 | v2)
                                                                    v109 := []int{}
                                                                    if ((127 <= (^(v90))) && v91) {
                                                                        v1 = 0
                                                                        if len(v109) > 0 {
                                                                            v109[0] = 9223372036854775807
                                                                        }
                                                                        switch v84 {
                                                                        case -2147483648:
                                                                            v82 += v82
                                                                        default:
                                                                            v108 -= (v2 + (v108 >> ((4294967295) & 63)))
                                                                        }
                                                                    }
                                                                    _ = v108
                                                                    _ = v109
                                                                }
                                                            }
                                                            _ = v90
                                                            _ = v91
                                                        }
                                                    }
                                                    v110 := GenS6{F0: v1, F1: true}
                                                    _ = v110
                                                    _ = v82
                                                    _ = v83
                                                    _ = v84
                                                } else {
                                                    v1 = (v1 - (v1 >> ((0) & 63)))
                                                    v111 := GenS1{F0: 4294967295, F1: "", F2: int32(v12)}
                                                    if (false || ((1e10 / (-2.25 / float64(-2.25))) > 3.14159)) {
                                                        var v112 int = fi29
                                                        _ = v112
                                                    } else {
                                                        var v113 int = (fi26 - fi26)
                                                        var v114 float32 = float32(((-2.25 - (0.1 + (float64(1e10) - 0.0))) - (3.14159 + (1.5 * 3.14159))))
                                                        _ = v113
                                                        _ = v114
                                                    }
                                                    _ = v111
                                                }
                                                v17 -= (^(v17))
                                                if true {
                                                    v115 := GenS7{F0: 1, F1: true}
                                                    v116 := GenS6{F0: (-(v1)), F1: (-991 == 680)}
                                                    v117 := GenS8{F0: v3, F1: v1, F2: (!true)}
                                                    _ = v115
                                                    _ = v116
                                                    _ = v117
                                                } else {
                                                    var v118 uint16 = (v2 / (((v2 % ((v2) | 1))) | 1))
                                                    v119 := []int{}
                                                    if len(v119) >= 2 {
                                                        v119 = v119[1:len(v119)]
                                                    }
                                                    _ = v118
                                                    _ = v119
                                                }
                                            } else {
                                                fmt.Println(v17, v2)
                                                var v120 int32 = (v1 * v1)
                                                _ = v120
                                            }
                                        case 2147483647:
                                            if false {
                                                v121 := []int{}
                                                _ = v121
                                            }
                                        default:
                                            if (139 <= 12) {
                                                if (false && false) {
                                                    if (genFn0() >= (v3 >> ((2) & 63))) {
                                                        var v122 int32 = (v1 << ((648) & 63))
                                                        fmt.Println(v122)
                                                        _ = v122
                                                    } else {
                                                        if false {
                                                            v123 := []int{}
                                                            _ = v123
                                                        } else {
                                                            if false {
                                                                switch (-(fi29)) {
                                                                case -2147483648:
                                                                    switch fi29 {
                                                                    case 127:
                                                                        var v124 bool = true
                                                                        _ = v124
                                                                    default:
                                                                        v125 := []int{}
                                                                        _ = v125
                                                                    }
                                                                case -9223372036854775808:
                                                                    var v126 int32 = (v1 - v1)
                                                                    var v127 int32 = v126
                                                                    if true {
                                                                        fmt.Println(v127, fi29, v3)
                                                                        v17 += v17
                                                                        v128 := []int{fi26, 0, fi26, -9223372036854775808}
                                                                        _ = v128
                                                                    } else {
                                                                        v129 := GenS9{F0: 255, F1: v126, F2: (fi26 / ((32768) | 1))}
                                                                        v130 := []int{(fi26 | 9223372036854775807), (fi26 >> ((255) & 63))}
                                                                        v131 := GenS9{F0: 910, F1: (v1 * (v1 >> ((652) & 63))), F2: (^(fi26))}
                                                                        _ = v129
                                                                        _ = v130
                                                                        _ = v131
                                                                    }
                                                                    _ = v126
                                                                    _ = v127
                                                                case 32768:
                                                                    var v132 uint8 = 0
                                                                    v133 := []int{(fi29 & fi26), (^(fi29))}
                                                                    switch ("a" + "\n\t") {
                                                                    case "a":
                                                                        var v134 float32 = 0.1
                                                                        _ = v134
                                                                    default:
                                                                        v132 = v132
                                                                    }
                                                                    _ = v132
                                                                    _ = v133
                                                                }
                                                                if true {
                                                                    if (false && false) {
                                                                        v135 := GenS9{F0: 42, F1: v1, F2: fi26}
                                                                        _ = v135
                                                                    }
                                                                }
                                                                var v136 uint64 = (v17 + v17)
                                                                _ = v136
                                                            }
                                                            v137 := []int{(fi29 - 128), fi29, (fi29 / (((^(fi29))) | 1)), fi26}
                                                            _ = v137
                                                        }
                                                        var v138 uint32 = v3
                                                        var v139 int64 = 853
                                                        _ = v138
                                                        _ = v139
                                                    }
                                                    v140 := []int{(fi26 ^ fi29), 128, fi26}
                                                    _ = v140
                                                }
                                                if false {
                                                    v141 := GenS10{F0: 943}
                                                    v142 := []int{fi26, 42}
                                                    var v143 int64 = int64(v3)
                                                    _ = v141
                                                    _ = v142
                                                    _ = v143
                                                } else {
                                                    v144 := GenS1{F0: 42, F1: ("foo bar" + ""), F2: int32(v12)}
                                                    var v145 int16 = 32767
                                                    _ = v144
                                                    _ = v145
                                                }
                                            }
                                        }
                                    } else {
                                        if ((v12 + (v12 * genFn1(genFn0()))) == genFn1((^(v3)))) {
                                            v146 := GenS11{F0: (v12 * (-(v12))), F1: v1}
                                            _ = v146
                                        } else {
                                            switch fi29 {
                                            case 0:
                                                var v147 int16 = 32767
                                                _ = v147
                                            case 32768:
                                                v3 = v3
                                            }
                                        }
                                        switch ("foo bar" + "a") {
                                        case "":
                                            v2 += (^(v2))
                                        case "foo bar":
                                            v148 := GenS8{F0: 65535, F1: v1, F2: ((v12 <= v12) || true)}
                                            _ = v148
                                        case "\n\t":
                                            v149 := []int{}
                                            _ = v149
                                        }
                                        fmt.Println(fi26)
                                        v150 := GenS1{F0: 477, F1: ("\n\t" + "\n\t"), F2: (v1 << ((997) & 63))}
                                        switch "foo bar" {
                                        case "\n\t":
                                            v151 := []int{(fi29 & 4294967295), (^(fi26))}
                                            _ = v151
                                        case "":
                                            v152 := []int{fi29, 9223372036854775807}
                                            _ = v152
                                        }
                                        _ = v150
                                    }
                                    var v153 uint32 = genFn0()
                                    if true {
                                        fmt.Println(fi26)
                                        v154 := GenS11{F0: v12, F1: v1}
                                        v155 := GenS4{F0: v1, F1: 255, F2: 0}
                                        _ = v154
                                        _ = v155
                                    }
                                    _ = v153
                                } else {
                                    if (uint8(v17) > uint8(v17)) {
                                        var v156 string = "\n\t"
                                        fmt.Println(v2, v2)
                                        _ = v156
                                    }
                                    v157 := []int{(fi26 - fi29)}
                                    _ = v157
                                }
                                v158 := GenS5{F0: "\n\t", F1: 1, F2: 1000}
                                var v159 uint8 = 196
                                _ = v158
                                _ = v159
                            } else {
                                var v160 int = 32767
                                _ = v160
                            }
                        }
                        v1 += (^(v1))
                        if (-128 <= (v12 * v12)) {
                            var v161 float64 = (1.5 + (3.14159 / 0.1))
                            if (!(255 != 1)) {
                                v162 := []int{9223372036854775807, (fi26 ^ fi26), int(v17), 1}
                                _ = v162
                            }
                            _ = v161
                        } else {
                            var v163 int64 = -1
                            _ = v163
                        }
                    }
                    _ = fi29
                }
                _ = fi26
            }
        default:
            fmt.Println(v2)
        }
        v164 := GenS12{F0: 42, F1: (^(v2))}
        _ = v164
    }
    v17 = 42
    switch "foo bar" {
    case "foo bar":
        var v165 uint = 65535
        _ = v165
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
}
