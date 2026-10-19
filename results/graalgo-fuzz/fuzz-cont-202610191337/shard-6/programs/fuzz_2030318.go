package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 int64; F1 bool; F2 string }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 uint16; F1 uint64 }
type GenS8 struct { F0 int32; F1 uint64 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint32; F1 int64; F2 int64 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 bool; F1 uint8; F2 uint }
type GenS15 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS16 struct { F0 int64; F1 uint8 }
type GenS17 struct { F0 string }
type GenS18 struct { F0 int8; F1 uint }
type GenS19 struct { F0 int8; F1 uint32 }
type GenS20 struct { F0 uint }
type GenS21 struct { F0 string; F1 uint8 }
func genFn0(p0 uint) int {
    for fi6 := 0; fi6 < 14; fi6++ {
        fmt.Println(fi6, p0)
        p0 = p0
        _ = fi6
    }
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 uint16 = 65535
        p0 += (p0 << ((p0) & 63))
        _ = fi7
        _ = v8
    }
    fmt.Println(p0)
    var v9 int = 52
    var v10 uint8 = 39
    _ = p0
    _ = v10
    _ = v9
    return 0
}
func genFn1(p0 int64, p1 uint8) uint32 {
    v11 := GenS1{F0: 672, F1: 585}
    v12 := GenS1{F0: 151, F1: 1}
    var v13 uint64 = 128
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return 442
}

func main() {
    var v1 uint8 = 231
    var v2 uint32 = 894
    var v3 float32 = (1e10 / ((1e10 - (-2.25 * 0.0)) / float32((0.1 / -2.25))))
    var v4 float32 = 1.5
    var v5 int8 = 100
    for fi14 := 0; fi14 < 12; fi14++ {
        if (42 < 4294967295) {
            v15 := GenS1{F0: 65535, F1: 18446744073709551615}
            for fi16 := 0; fi16 < 14; fi16++ {
                v17 := GenS1{F0: 18446744073709551615, F1: 216}
                v18 := []int{2147483647, genFn0(42), fi14, genFn0(255)}
                var v19 int64 = 0
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            v5 += int8(fi14)
            _ = v15
        }
        if false {
            switch ("" + (("foo bar" + "a") + "")) {
            case "foo bar":
                fmt.Println(v3, v4, v3)
            case "\n\t":
                v1 -= v1
            default:
                for fi20 := 0; fi20 < 15; fi20++ {
                    if false {
                        var v21 uint32 = (^(v2))
                        v22 := GenS1{F0: 294, F1: 2}
                        if true {
                            var v23 int64 = -9223372036854775808
                            fmt.Println(fi20)
                            v24 := []int{(fi20 & fi14), (fi20 >> ((2) & 63)), (fi14 - 4294967295)}
                            v2 -= (v2 / (((v2 & uint32(v1))) | 1))
                            _ = v23
                            _ = v24
                        }
                        v25 := GenS1{F0: 18446744073709551615, F1: 4294967295}
                        _ = v21
                        _ = v22
                        _ = v25
                    } else {
                        if (("a" + "") == ((("\n\t" + "\n\t") + "a") + ((("foo bar" + "foo bar") + "\n\t") + (((("hello" + "foo bar") + "hello") + "\n\t") + "\n\t")))) {
                            if false {
                                var v26 int32 = 0
                                if false {
                                    fmt.Println(v26, v1)
                                    if ((fi14 >> ((72) & 63)) >= fi14) {
                                        var v27 bool = false
                                        v28 := []int{fi14, (fi14 << ((541) & 63)), (fi14 << ((65535) & 63)), 32767}
                                        if (1.5 < (((float64(v3) * (3.14159 + (1e10 * float64(((float32((1.5 * 3.14159)) / 1.5) / 3.14159))))) - 1e10) / (1e10 / 1.5))) {
                                            v5 -= (v5 - (v5 | -128))
                                            v1 = (^(v1))
                                            switch ("hello" + "\n\t") {
                                            case "\n\t":
                                                switch "" {
                                                case "a":
                                                    v28 = append(v28, fi20)
                                                default:
                                                    switch fi20 {
                                                    case 84:
                                                        v29 := GenS2{F0: v27}
                                                        _ = v29
                                                    }
                                                }
                                            case "a":
                                                if (!false) {
                                                    if ((-(v5)) >= 42) {
                                                        var v30 uint64 = 1
                                                        v4 = float32(1e10)
                                                        _ = v30
                                                    }
                                                    if (!(v4 >= (1e10 - (v4 * (3.14159 - v4))))) {
                                                        var v31 uint64 = 42
                                                        _ = v31
                                                    }
                                                }
                                            case "foo bar":
                                                _ = len(v28)
                                            }
                                        }
                                        _ = v27
                                        _ = v28
                                    }
                                    v26 = (v26 & v26)
                                    switch int(v26) {
                                    case -9223372036854775808:
                                        switch ("foo bar" + "foo bar") {
                                        case "hello":
                                            if true {
                                                v32 := GenS1{F0: 926, F1: 42}
                                                _ = v32
                                            } else {
                                                if ((v1 / ((255) | 1)) <= v1) {
                                                    var v33 int64 = 1
                                                    v26 = (v26 + int32(v2))
                                                    v34 := GenS2{F0: (!true)}
                                                    _ = v33
                                                    _ = v34
                                                }
                                            }
                                        case "foo bar":
                                            if (245 >= 1) {
                                                if true {
                                                    v5 = v5
                                                } else {
                                                    var v35 int64 = 0
                                                    _ = v35
                                                }
                                                v36 := []int{fi20, -574}
                                                _ = v36
                                            }
                                        }
                                    }
                                    switch "foo bar" {
                                    case "a":
                                        v4 = v4
                                        var v37 uint = 42
                                        _ = v37
                                    default:
                                        v26 = (v26 << ((42) & 63))
                                    }
                                    v38 := []int{(fi14 % ((-2147483648) | 1)), fi20, fi14}
                                    _ = v38
                                } else {
                                    switch 32767 {
                                    case -627:
                                        v39 := GenS2{F0: (1000 != 1)}
                                        _ = v39
                                    case 0:
                                        if true {
                                            v40 := GenS3{F0: 1, F1: 42}
                                            if false {
                                                v26 -= -1
                                                v41 := GenS4{F0: 606, F1: "\n\t"}
                                                _ = v41
                                            } else {
                                                var v42 uint32 = (v2 % ((v2) | 1))
                                                _ = v42
                                            }
                                            switch "hello" {
                                            case "\n\t":
                                                var v43 uint64 = 404
                                                _ = v43
                                            case "hello":
                                                if false {
                                                    v44 := []int{}
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        fmt.Println(v2)
                                                    case "a":
                                                        if false {
                                                            var v45 uint32 = (v2 % ((42) | 1))
                                                            v46 := GenS3{F0: 694, F1: v45}
                                                            v47 := GenS5{F0: int64(fi20), F1: true, F2: ((("foo bar" + "") + ("" + ("" + "foo bar"))) + ((("a" + "foo bar") + "a") + "\n\t"))}
                                                            _ = v45
                                                            _ = v46
                                                            _ = v47
                                                        } else {
                                                            var v48 bool = (42 >= v26)
                                                            v5 = v5
                                                            _ = v48
                                                        }
                                                    case "":
                                                        fmt.Println(v4, v4)
                                                    }
                                                    _ = v44
                                                } else {
                                                    var v49 bool = (38 == 228)
                                                    _ = v49
                                                }
                                            default:
                                                v3 = v4
                                            }
                                            _ = v40
                                        } else {
                                            var v50 uint64 = 4294967295
                                            if false {
                                                v51 := []int{}
                                                if true {
                                                    if len(v51) >= 2 {
                                                        v51 = v51[1:len(v51)]
                                                    }
                                                }
                                                v52 := GenS3{F0: 0, F1: 49}
                                                _ = v51
                                                _ = v52
                                            } else {
                                                if (!true) {
                                                    if (v1 < v1) {
                                                        var v53 uint = 255
                                                        v50 += (v50 ^ (v50 ^ (v50 % ((v50) | 1))))
                                                        switch genFn0(52) {
                                                        case -1:
                                                            switch genFn0((^(v53))) {
                                                            case 0:
                                                                fmt.Println(v5, v4)
                                                                if ((fi14 - (fi14 % ((fi20) | 1))) != -279) {
                                                                    v54 := GenS2{F0: true}
                                                                    _ = v54
                                                                } else {
                                                                    if (float64(v3) <= -2.25) {
                                                                        fmt.Println(v5, v3)
                                                                    } else {
                                                                        var v55 int = genFn0(v53)
                                                                        v1 += (v1 * (v1 >> ((v53) & 63)))
                                                                        fmt.Println(v5, v1)
                                                                        _ = v55
                                                                    }
                                                                }
                                                                v56 := GenS4{F0: 32767, F1: "foo bar"}
                                                                _ = v56
                                                            case 868:
                                                                v57 := []int{fi14, 127}
                                                                _ = v57
                                                            default:
                                                                fmt.Println(v50)
                                                            }
                                                        case 32768:
                                                            v58 := GenS6{F0: 1}
                                                            _ = v58
                                                        default:
                                                            fmt.Println(v26, v26, v50)
                                                        }
                                                        _ = v53
                                                    }
                                                    if false {
                                                        v5 -= -11
                                                        var v59 uint32 = 568
                                                        var v60 int = fi20
                                                        _ = v59
                                                        _ = v60
                                                    } else {
                                                        var v61 string = "foo bar"
                                                        v62 := GenS5{F0: -450, F1: (4294967295 < 255), F2: v61}
                                                        switch (fi20 - 32768) {
                                                        case -80:
                                                            var v63 string = v61
                                                            _ = v63
                                                        case 2147483647:
                                                            if false {
                                                                var v64 int16 = 42
                                                                if (2 != v2) {
                                                                    if ((v64 / ((0) | 1)) != (v64 / (((v64 >> ((uint(v50)) & 63))) | 1))) {
                                                                        v65 := []int{fi14, 530, genFn0(65535)}
                                                                        _ = v65
                                                                    } else {
                                                                        v66 := GenS4{F0: -32768, F1: "hello"}
                                                                        _ = v66
                                                                    }
                                                                } else {
                                                                    var v67 uint16 = 2
                                                                    var v68 uint = 18446744073709551615
                                                                    _ = v67
                                                                    _ = v68
                                                                }
                                                                _ = v64
                                                            }
                                                        case 32768:
                                                            var v69 uint32 = 2
                                                            _ = v69
                                                        default:
                                                            v70 := []int{128, (^(fi20)), (fi14 << ((42) & 63)), fi14}
                                                            _ = v70
                                                        }
                                                        if false {
                                                            if false {
                                                                if (-1 >= 42) {
                                                                    if (v61 != v61) {
                                                                        v71 := []int{}
                                                                        v26 = (v26 % (((v26 ^ (v26 / ((2147483647) | 1)))) | 1))
                                                                        if false {
                                                                            v5 += v5
                                                                            var v72 int64 = -997
                                                                            _ = v72
                                                                        }
                                                                        _ = v71
                                                                    } else {
                                                                        v5 = 100
                                                                        v1 += v1
                                                                    }
                                                                } else {
                                                                    v73 := GenS7{F0: 144, F1: v50}
                                                                    fmt.Println(v50, v1, v61)
                                                                    _ = v73
                                                                }
                                                            } else {
                                                                fmt.Println(fi14, fi20)
                                                                switch (fi14 >> ((1) & 63)) {
                                                                case 1:
                                                                    var v74 uint8 = 1
                                                                    _ = v74
                                                                case -2147483648:
                                                                    if false {
                                                                        var v75 uint8 = (v1 / ((v1) | 1))
                                                                        v76 := []int{(fi20 % ((genFn0(255)) | 1)), (-(fi20))}
                                                                        _ = v75
                                                                        _ = v76
                                                                    } else {
                                                                        v77 := GenS8{F0: (-(v26)), F1: 0}
                                                                        v78 := []int{int(v26), int(v26)}
                                                                        if false {
                                                                            var v79 uint16 = 42
                                                                            v80 := GenS8{F0: 0, F1: 412}
                                                                            if true {
                                                                                var v81 bool = false
                                                                                v82 := GenS1{F0: 670, F1: v50}
                                                                                var v83 uint8 = (v1 % ((v1) | 1))
                                                                                _ = v81
                                                                                _ = v82
                                                                                _ = v83
                                                                            }
                                                                            fmt.Println(v1, v4, v26)
                                                                            fmt.Println(v1)
                                                                            _ = v79
                                                                            _ = v80
                                                                        }
                                                                        _ = v77
                                                                        _ = v78
                                                                    }
                                                                default:
                                                                    fmt.Println(fi14, v50, v1)
                                                                }
                                                                fmt.Println(v1, fi14)
                                                                v1 -= 42
                                                                v84 := []int{(fi14 >> ((uint(v50)) & 63)), (fi20 * (fi20 ^ (fi14 | fi14))), 189}
                                                                if ((true || true) && (2 < v1)) {
                                                                    var v85 int8 = v5
                                                                    var v86 int8 = (v5 >> ((590) & 63))
                                                                    _ = v85
                                                                    _ = v86
                                                                }
                                                                fmt.Println(v3, v50)
                                                                v50 = 4294967295
                                                                if false {
                                                                    switch genFn0(uint(v1)) {
                                                                    case 1:
                                                                        v4 = (v3 * ((0.1 * 1e10) / v3))
                                                                    }
                                                                } else {
                                                                    var v87 int = (fi14 | (fi14 - (fi14 + fi20)))
                                                                    _ = v87
                                                                }
                                                                v88 := GenS9{F0: false}
                                                                v5 = v5
                                                                var v89 uint64 = v50
                                                                _ = v84
                                                                _ = v88
                                                                _ = v89
                                                            }
                                                            var v90 int8 = v5
                                                            _ = v90
                                                        } else {
                                                            switch (^(fi14)) {
                                                            case -2147483648:
                                                                fmt.Println(v2, v26, v2)
                                                            }
                                                            v91 := []int{(fi20 << ((18446744073709551615) & 63))}
                                                            _ = v91
                                                        }
                                                        switch v61 {
                                                        case "":
                                                            v92 := GenS5{F0: -217, F1: ((true || (128 > fi14)) || (false || ((v50 >> ((65535) & 63)) == (^(v50))))), F2: v61}
                                                            _ = v92
                                                        case "\n\t":
                                                            fmt.Println(v5)
                                                        case "foo bar":
                                                            fmt.Println(v50, v3, v5)
                                                        default:
                                                            var v93 uint64 = (v50 | 790)
                                                            _ = v93
                                                        }
                                                        _ = v61
                                                        _ = v62
                                                    }
                                                    var v94 uint32 = genFn1(-1, (v1 >> ((255) & 63)))
                                                    if true {
                                                        v95 := GenS8{F0: v26, F1: (v50 << ((65535) & 63))}
                                                        _ = v95
                                                    } else {
                                                        v96 := []int{genFn0(65535), fi14, genFn0(255)}
                                                        v97 := GenS10{F0: 42}
                                                        v2 += genFn1(463, (v1 & 255))
                                                        var v98 float32 = v3
                                                        _ = v96
                                                        _ = v97
                                                        _ = v98
                                                    }
                                                    var v99 uint16 = uint16(v1)
                                                    var v100 int = fi20
                                                    _ = v100
                                                    _ = v94
                                                    _ = v99
                                                }
                                            }
                                            _ = v50
                                        }
                                    case -2147483648:
                                        fmt.Println(fi20, fi20)
                                    }
                                    v101 := GenS11{F0: uint32(v1), F1: 42, F2: 20}
                                    _ = v101
                                }
                                v102 := GenS5{F0: int64(v1), F1: (!true), F2: "\n\t"}
                                _ = v102
                                _ = v26
                            } else {
                                v103 := []int{772, 4294967295}
                                v4 += ((float32(3.14159) + ((((v4 + 3.14159) - float32(3.14159)) / float32((1.5 * 3.14159))) + v3)) + 1e10)
                                v4 = v4
                                v1 -= (v1 >> ((uint(v2)) & 63))
                                _ = v103
                            }
                            v104 := []int{128, 68}
                            switch 2147483647 {
                            case 32767:
                                v1 += 32
                            case 1:
                                v104 = append(v104, (-(fi20)))
                            default:
                                v105 := GenS12{F0: true}
                                _ = v105
                            }
                            v1 = v1
                            _ = v104
                        } else {
                            var v106 bool = false
                            v5 -= (v5 ^ (-(v5)))
                            if ((^(fi14)) == fi20) {
                                var v107 uint8 = (^(v1))
                                var v108 bool = (571 >= genFn0(18446744073709551615))
                                v2 += (v2 ^ 1)
                                _ = v107
                                _ = v108
                            } else {
                                if v106 {
                                    var v109 uint = 1
                                    _ = v109
                                }
                                v110 := GenS12{F0: v106}
                                v4 = v3
                                var v111 int = (fi20 / (((fi14 >> ((1) & 63))) | 1))
                                switch ("a" + ("" + "\n\t")) {
                                case "\n\t":
                                    var v112 bool = true
                                    _ = v112
                                case "a":
                                    var v113 uint8 = (^(v1))
                                    _ = v113
                                default:
                                    if v106 {
                                        switch ("a" + "\n\t") {
                                        case "":
                                            var v114 uint64 = 1
                                            _ = v114
                                        case "hello":
                                            if false {
                                                var v115 uint16 = 1
                                                v106 = v106
                                                var v116 int32 = 0
                                                _ = v115
                                                _ = v116
                                            }
                                        default:
                                            v117 := []int{fi14, genFn0(255), fi20, v111}
                                            _ = v117
                                        }
                                        switch (^(fi20)) {
                                        case 127:
                                            v3 += v4
                                        case 9223372036854775807:
                                            v3 = ((float32(float64(v3)) * (v3 * -2.25)) - ((0.1 + float32(0.0)) / -2.25))
                                        case -573:
                                            var v118 string = ("a" + "hello")
                                            _ = v118
                                        default:
                                            v4 += v4
                                        }
                                    }
                                }
                                _ = v110
                                _ = v111
                            }
                            _ = v106
                        }
                        v119 := []int{fi14, fi14, (-(fi20)), (fi20 * (^(fi20)))}
                        if false {
                            switch (-(fi14)) {
                            case 42:
                                switch (^(fi14)) {
                                case -153:
                                    var v120 uint8 = (v1 << ((18446744073709551615) & 63))
                                    _ = v120
                                }
                            case 134:
                                v121 := GenS9{F0: true}
                                _ = v121
                            case -2147483648:
                                if ((true || (!(9223372036854775807 == 954))) || true) {
                                    fmt.Println(v4)
                                } else {
                                    fmt.Println(v2, v1, fi14)
                                    var v122 int32 = 1
                                    var v123 uint32 = 1
                                    _ = v122
                                    _ = v123
                                }
                            }
                            if len(v119) >= 2 {
                                v119 = v119[1:len(v119)]
                            }
                        }
                        if (false && (2147483647 >= -34)) {
                            var v124 bool = false
                            var v125 int8 = 127
                            switch "a" {
                            case "hello":
                                v1 += v1
                            case "\n\t":
                                if (!true) {
                                    v125 = (^(v5))
                                    var v126 float64 = -2.25
                                    _ = v126
                                } else {
                                    fmt.Println(v2, fi14)
                                }
                            case "a":
                                var v127 uint32 = 4294967295
                                _ = v127
                            }
                            _ = v124
                            _ = v125
                        } else {
                            switch (fi14 % ((genFn0(371)) | 1)) {
                            case 1:
                                v119 = append(v119, (fi14 - genFn0(1)))
                            case -604:
                                var v128 uint32 = genFn1(int64(fi20), v1)
                                _ = v128
                            case -9223372036854775808:
                                v119 = append(v119, genFn0(606))
                            }
                            var v129 uint16 = 0
                            v130 := GenS1{F0: 42, F1: 189}
                            _ = v129
                            _ = v130
                        }
                        _ = v119
                    }
                    v131 := []int{589}
                    switch -250 {
                    case 32767:
                        var v132 uint = 0
                        _ = v132
                    }
                    _ = fi20
                    _ = v131
                }
            }
            var v133 string = ("hello" + "foo bar")
            v1 += 61
            _ = v133
        } else {
            v5 += 21
        }
        _ = fi14
    }
    v134 := []int{-2147483648, 199, 0, -1}
    if false {
        var v135 uint = 4294967295
        if len(v134) > 0 {
            v134[0] = -2147483648
        }
        _ = v135
    } else {
        if true {
            if len(v134) > 0 {
                v134[0] = 4294967295
            }
            v3 += float32(float64(((0.1 - float32((3.14159 * 1e10))) / v3)))
        }
        _ = len(v134)
    }
    for fi136 := 0; fi136 < 7; fi136++ {
        if false {
            if (!true) {
                v137 := GenS11{F0: (^(v2)), F1: 42, F2: 89}
                var v138 int64 = -9223372036854775808
                v1 -= v1
                _ = v137
                _ = v138
            } else {
                var v139 uint = 1
                _ = v139
            }
            _ = len(v134)
        }
        if (!(true || false)) {
            var v140 uint8 = v1
            if len(v134) >= 2 {
                v134 = v134[1:len(v134)]
            }
            if ((v140 + 0) < v1) {
                v2 += (^(v2))
                for fi141 := 0; fi141 < 9; fi141++ {
                    fmt.Println(fi136, v1, v4)
                    _ = fi141
                }
            }
            _ = v140
        } else {
            switch "a" {
            case "\n\t":
                var v142 int8 = (v5 >> ((255) & 63))
                _ = v142
            }
            var v143 bool = ((3.14159 - 0.1) == -2.25)
            _ = v143
        }
        _ = fi136
    }
    fmt.Println(v2, v5)
    if len(v134) > 0 {
        v134[0] = genFn0(357)
    }
    if ((^(v1)) >= (v1 ^ 83)) {
        var v144 int64 = -9223372036854775808
        if len(v134) > 0 {
            v134[0] = 32767
        }
        v5 -= (v5 % ((v5) | 1))
        _ = v144
    } else {
        if (-1 == -15) {
            for fi145 := 0; fi145 < 6; fi145++ {
                var v146 bool = true
                v147 := GenS13{F0: 32767}
                _ = fi145
                _ = v146
                _ = v147
            }
        }
        var v148 float32 = (float32(1e10) - 1.5)
        _ = v148
    }
    var v149 int = genFn0(65535)
    if (1.5 >= 3.14159) {
        switch "hello" {
        case "a":
            v5 = int8(v2)
        case "":
            if len(v134) > 0 {
                v134[0] = (v149 | genFn0(42))
            }
        case "hello":
            var v150 int16 = 715
            _ = v150
        }
    }
    v151 := GenS10{F0: 42}
    v5 = (v5 | (v5 << ((0) & 63)))
    for fi152 := 0; fi152 < 15; fi152++ {
        fmt.Println(v149, v3, v4)
        v153 := GenS14{F0: (1 == 4294967295), F1: 31, F2: 65535}
        _ = fi152
        _ = v153
    }
    fmt.Println(v1, v1)
    if (true || (("" + ("" + ("foo bar" + "foo bar"))) > "\n\t")) {
        v5 -= v5
        if (v1 > v1) {
            if (65535 > uint16(v1)) {
                fmt.Println(v2, v2)
                if len(v134) > 0 {
                    v134[0] = 128
                }
                for fi154 := 0; fi154 < 9; fi154++ {
                    if true {
                        if len(v134) >= 2 {
                            v134 = v134[1:len(v134)]
                        }
                        var v155 string = "foo bar"
                        v3 -= v4
                        _ = v155
                    } else {
                        var v156 uint = 1
                        var v157 int32 = 2147483647
                        for fi158 := 0; fi158 < 4; fi158++ {
                            if true {
                                v159 := GenS15{F0: v1, F1: v5, F2: v1}
                                v134 = append(v134, fi154)
                                v156 = 672
                                _ = v159
                            }
                            _ = fi158
                        }
                        _ = v156
                        _ = v157
                    }
                    if (!true) {
                        v160 := GenS16{F0: 732, F1: v1}
                        _ = v160
                    }
                    var v161 uint32 = v2
                    _ = fi154
                    _ = v161
                }
            }
        } else {
            var v162 uint8 = (v1 << ((754) & 63))
            if len(v134) > 0 {
                v134[0] = v149
            }
            _ = v162
        }
    } else {
        if false {
            for fi163 := 0; fi163 < 3; fi163++ {
                var v164 uint64 = uint64(v1)
                _ = fi163
                _ = v164
            }
            v4 -= 1.5
            if (int64(v5) <= 42) {
                switch 128 {
                case -9223372036854775808:
                    if false {
                        v149 += genFn0(18446744073709551615)
                        if true {
                            v1 = (^(v1))
                            if (251 > genFn1(961, v1)) {
                                if true {
                                    var v165 int = v149
                                    fmt.Println(v2, v2)
                                    v5 += v5
                                    _ = v165
                                } else {
                                    if len(v134) >= 2 {
                                        v134 = v134[1:len(v134)]
                                    }
                                    v149 = v149
                                    switch ("hello" + ("foo bar" + ("hello" + (("a" + ("hello" + "foo bar")) + ("foo bar" + ((((("foo bar" + "\n\t") + "\n\t") + (("" + "foo bar") + "")) + "foo bar") + ("a" + "a"))))))) {
                                    case "hello":
                                        if len(v134) > 0 {
                                            v134[0] = genFn0(2)
                                        }
                                    case "a":
                                        fmt.Println(v5, v4, v1)
                                    case "":
                                        v2 -= (^(v2))
                                    }
                                }
                                var v166 uint32 = (v2 % ((0) | 1))
                                _ = v166
                            }
                            for fi167 := 0; fi167 < 8; fi167++ {
                                switch (fi167 ^ genFn0(217)) {
                                case -387:
                                    if (false || false) {
                                        var v168 bool = true
                                        switch "a" {
                                        case "foo bar":
                                            switch -1 {
                                            case 220:
                                                var v169 uint16 = 42
                                                _ = v169
                                            case 747:
                                                v170 := GenS17{F0: ("" + (("hello" + "") + (("hello" + "a") + "a")))}
                                                _ = v170
                                            }
                                        case "hello":
                                            var v171 uint16 = 2
                                            _ = v171
                                        }
                                        v172 := GenS16{F0: -1, F1: 0}
                                        _ = v168
                                        _ = v172
                                    }
                                case -175:
                                    v173 := GenS12{F0: false}
                                    _ = v173
                                case -2147483648:
                                    fmt.Println(v2, v1)
                                }
                                v149 -= genFn0(2)
                                fmt.Println(fi167, fi167)
                                _ = fi167
                            }
                        } else {
                            if (0 != 72) {
                                for fi174 := 0; fi174 < 3; fi174++ {
                                    _ = len(v134)
                                    _ = fi174
                                }
                                var v175 int8 = -1
                                _ = v175
                            }
                            if true {
                                switch "" {
                                case "foo bar":
                                    v149 += int(v1)
                                default:
                                    if (!false) {
                                        switch (v149 ^ v149) {
                                        case 248:
                                            if (true || false) {
                                                if len(v134) > 0 {
                                                    v134[0] = (v149 | -829)
                                                }
                                                v176 := GenS18{F0: (^(v5)), F1: 2}
                                                v177 := GenS15{F0: v1, F1: -128, F2: 0}
                                                _ = v176
                                                _ = v177
                                            }
                                        default:
                                            v149 -= (v149 + (v149 % (((v149 << ((4294967295) & 63))) | 1)))
                                        }
                                        if (0 < -9223372036854775808) {
                                            v178 := GenS19{F0: v5, F1: (v2 / ((729) | 1))}
                                            var v179 int8 = (v5 & v5)
                                            _ = v178
                                            _ = v179
                                        } else {
                                            if (true || false) {
                                                v5 = v5
                                                v149 -= (v149 ^ -16)
                                                for fi180 := 0; fi180 < 4; fi180++ {
                                                    v181 := GenS4{F0: -1, F1: (((("" + (("a" + "hello") + "")) + "hello") + "\n\t") + "foo bar")}
                                                    var v182 int16 = 848
                                                    var v183 int8 = v5
                                                    _ = fi180
                                                    _ = v181
                                                    _ = v182
                                                    _ = v183
                                                }
                                            }
                                        }
                                        if false {
                                            switch genFn0(65535) {
                                            case 535:
                                                v184 := GenS19{F0: 0, F1: 642}
                                                _ = v184
                                            case 4294967295:
                                                if (3.14159 > 0.1) {
                                                    v5 -= (v5 ^ int8(v2))
                                                    v185 := GenS20{F0: 1}
                                                    v5 -= v5
                                                    _ = v185
                                                }
                                            }
                                        } else {
                                            var v186 uint64 = 18446744073709551615
                                            var v187 float32 = (v3 + v3)
                                            _ = v186
                                            _ = v187
                                        }
                                    } else {
                                        switch "" {
                                        case "":
                                            v5 = -77
                                        default:
                                            var v188 uint8 = 0
                                            _ = v188
                                        }
                                    }
                                }
                                var v189 float64 = 3.14159
                                v4 -= 0.0
                                _ = v189
                            } else {
                                fmt.Println(v5)
                                switch (v149 % ((v149) | 1)) {
                                case 0:
                                    _ = len(v134)
                                case -1:
                                    fmt.Println(v1)
                                case 185:
                                    v4 = (float32((0.0 * 3.14159)) * -2.25)
                                }
                            }
                            var v190 uint32 = v2
                            _ = v190
                        }
                    } else {
                        var v191 float32 = 1.5
                        var v192 int16 = 42
                        _ = v191
                        _ = v192
                    }
                case 128:
                    for fi193 := 0; fi193 < 1; fi193++ {
                        v1 -= (^(v1))
                        switch "a" {
                        case "a":
                            v194 := GenS16{F0: 9223372036854775807, F1: (v1 | v1)}
                            _ = v194
                        case "hello":
                            fmt.Println(v1, v1)
                        }
                        v195 := GenS6{F0: 0}
                        _ = fi193
                        _ = v195
                    }
                case 777:
                    v1 = 42
                }
                if len(v134) >= 2 {
                    v134 = v134[1:len(v134)]
                }
                for fi196 := 0; fi196 < 13; fi196++ {
                    switch "hello" {
                    case "":
                        var v197 uint8 = (v1 * (v1 >> ((1) & 63)))
                        _ = v197
                    default:
                        fmt.Println(v1, v2)
                    }
                    fmt.Println(fi196, v5, v2)
                    switch "a" {
                    case "\n\t":
                        if (1 == 132) {
                            v5 -= (v5 >> ((255) & 63))
                            var v198 int64 = 442
                            if ((^(v149)) >= v149) {
                                if false {
                                    if len(v134) > 0 {
                                        v134[0] = genFn0(2)
                                    }
                                    v199 := GenS2{F0: false}
                                    v4 += v3
                                    _ = v199
                                }
                            } else {
                                if true {
                                    var v200 int = v149
                                    _ = v200
                                } else {
                                    v4 += (float32(-2.25) * (v4 - (float32((1e10 * 0.1)) + (float32(3.14159) - (float32(((((1e10 + float64(3.14159)) + float64(0.0)) + (1e10 + ((-2.25 * 0.0) - 1.5))) * (float64(1e10) + 1e10))) / float32(((-2.25 * 0.1) + float64(float32(0.0)))))))))
                                    v201 := GenS21{F0: "\n\t", F1: 77}
                                    if false {
                                        var v202 uint64 = 42
                                        v203 := GenS2{F0: true}
                                        v5 -= v5
                                        _ = v202
                                        _ = v203
                                    } else {
                                        var v204 uint8 = (v1 * 127)
                                        if len(v134) > 0 {
                                            v134[0] = (fi196 % (((v149 & (v149 >> ((42) & 63)))) | 1))
                                        }
                                        _ = v204
                                    }
                                    _ = v201
                                }
                                if (!(!false)) {
                                    var v205 uint = 2
                                    _ = v205
                                }
                            }
                            _ = v198
                        } else {
                            fmt.Println(v5)
                        }
                    default:
                        v3 = v3
                    }
                    _ = fi196
                }
            }
        }
        v134 = append(v134, 32767)
        fmt.Println(v3, v4, v2)
    }
    var v206 uint8 = v1
    v1 = (v206 << ((4294967295) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v134:", len(v134), v134)
    fmt.Println("v149:", v149)
    fmt.Println("v2:", v2)
    fmt.Println("v206:", v206)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v134
    _ = v149
    _ = v151
    _ = v2
    _ = v206
    _ = v3
    _ = v4
    _ = v5
}
