package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int64; F1 int; F2 int16 }
type GenS4 struct { F0 int8; F1 uint8; F2 int32 }
type GenS5 struct { F0 bool; F1 int32; F2 uint }
type GenS6 struct { F0 int8; F1 uint16; F2 string }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint8; F1 int; F2 int16 }
type GenS9 struct { F0 uint; F1 int64 }
type GenS10 struct { F0 string }
type GenS11 struct { F0 uint32; F1 int; F2 uint64 }
type GenS12 struct { F0 uint64 }
type GenS13 struct { F0 bool; F1 int64 }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 int32 }
func genFn0(p0 int, p1 int32, p2 int) uint32 {
    fmt.Println(p0, p1)
    if (65535 >= 0) {
        v4 := []int{p0, (p2 ^ (^(p0)))}
        var v5 int64 = 42
        _ = v4
        _ = v5
    } else {
        v6 := []int{(p0 + (-(p0))), p0, p2, p2}
        _ = v6
    }
    v7 := GenS1{F0: 42, F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = ((3.14159 - ((-2.25 / 0.1) + float32(1e10))) + (0.1 / -2.25))
    v3 := make(map[int]bool)
    var v8 string = "foo bar"
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := v1[(-(fi9))]
        _ = fi9
        _ = v10
    }
    switch 9223372036854775807 {
    case 1:
        for fi11 := 0; fi11 < 12; fi11++ {
            _ = len(v1)
            var v12 int = fi11
            _ = fi11
            _ = v12
        }
    case 128:
        v13 := GenS2{F0: 1, F1: 258}
        _ = v13
    case 32767:
        var v14 uint8 = 40
        _ = v14
    default:
        var v15 uint8 = 1
        _ = v15
    }
    if true {
        if ((1 <= 42) || true) {
            var v16 uint64 = 42
            _ = v16
        }
        v2 -= float32(0.1)
    }
    var v17 bool = false
    v18 := GenS3{F0: -527, F1: 32767, F2: -32768}
    v19 := GenS2{F0: -2147483648, F1: 221}
    v2 = (1.5 / v2)
    v20 := GenS4{F0: 69, F1: 0, F2: 892}
    var v21 float32 = 0.0
    for fi22 := 0; fi22 < 10; fi22++ {
        v23 := GenS4{F0: 42, F1: 127, F2: -547}
        var v24 int = 32768
        var v25 uint16 = 2
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    var v26 uint64 = 0
    v21 = v21
    delete(v3, 42)
    for fi27 := 0; fi27 < 6; fi27++ {
        v28 := v3[fi27]
        v29 := GenS5{F0: v17, F1: 1, F2: 42}
        _ = fi27
        _ = v28
        _ = v29
    }
    if v17 {
        v17 = (65535 < 2147483647)
    } else {
        switch -38 {
        case 32768:
            if (-2 < 42) {
                v26 -= v26
            }
        case -1:
            var v30 int16 = 0
            _ = v30
        }
        v31 := GenS6{F0: 81, F1: uint16(v26), F2: ("hello" + "hello")}
        _ = v31
    }
    for fi32 := 0; fi32 < 6; fi32++ {
        for fi33 := 0; fi33 < 18; fi33++ {
            v2 += (v2 - -2.25)
            if v17 {
                v34 := GenS7{F0: 5}
                _ = v34
            }
            if (!(!v17)) {
                if (-1 == -1) {
                    switch fi33 {
                    case 1:
                        fmt.Println(fi32, v26)
                    case 127:
                        if false {
                            v26 = 18446744073709551615
                            v2 -= 0.1
                            if true {
                                v17 = (-1 < -901)
                                if v17 {
                                    var v35 uint8 = 127
                                    v3[(fi33 ^ (fi32 & (fi33 * (fi32 & (fi32 >> ((255) & 63))))))] = v17
                                    _ = v35
                                }
                                var v36 bool = (((-2.25 + v21) >= v21) || true)
                                _ = v36
                            } else {
                                var v37 int = (^(fi32))
                                _ = v37
                            }
                            var v38 uint = 255
                            _ = v38
                        }
                    case 32768:
                        fmt.Println(fi32)
                    default:
                        switch (fi32 % (((^(fi33))) | 1)) {
                        case 128:
                            var v39 bool = (!(42 < 576))
                            _ = v39
                        default:
                            v17 = v17
                        }
                        var v40 string = v8
                        _ = v40
                    }
                    fmt.Println(v21)
                    var v41 bool = true
                    _ = v41
                }
                fmt.Println(v21, v21, v21)
            } else {
                v42 := GenS8{F0: 127, F1: 127, F2: 32767}
                if (v17 && (fi32 != -2147483648)) {
                    var v43 int64 = 1
                    if v17 {
                        switch (("foo bar" + (v8 + v8)) + v8) {
                        case "\n\t":
                            switch (fi33 / ((128) | 1)) {
                            case 128:
                                switch "" {
                                case "":
                                    v1[(^(fi33))] = (!(!(fi33 == (fi32 + -2147483648))))
                                case "a":
                                    delete(v1, (fi33 / (((fi32 - -9223372036854775808)) | 1)))
                                case "\n\t":
                                    switch (fi32 + (fi33 >> ((226) & 63))) {
                                    case 0:
                                        var v44 uint64 = 302
                                        _ = v44
                                    case 32768:
                                        fmt.Println(v17)
                                    case -921:
                                        v3[fi32] = (0.0 >= float64((v2 + v2)))
                                    default:
                                        v21 = (v2 + (-2.25 / float32(0.1)))
                                    }
                                default:
                                    var v45 string = v8
                                    _ = v45
                                }
                            case 4294967295:
                                var v46 uint16 = 32767
                                _ = v46
                            default:
                                fmt.Println(v17)
                            }
                        case "foo bar":
                            switch fi32 {
                            case 2147483647:
                                v8 = v8
                            case 32768:
                                delete(v1, fi33)
                            case 0:
                                delete(v1, (fi33 / ((fi32) | 1)))
                            }
                        default:
                            v17 = (1e10 >= 1e10)
                            v47 := []int{fi33, (fi33 >> ((2) & 63)), fi32, fi33}
                            _ = v47
                        }
                    }
                    v48 := []int{-1}
                    _ = v43
                    _ = v48
                } else {
                    var v49 int16 = 0
                    var v50 string = (("\n\t" + ("\n\t" + ("hello" + (v8 + "a")))) + "a")
                    if false {
                        _ = len(v3)
                        switch v50 {
                        case "hello":
                            if true {
                                switch "" {
                                case "a":
                                    var v51 int64 = -539
                                    _ = v51
                                case "":
                                    v52 := GenS8{F0: 134, F1: 2147483647, F2: v49}
                                    _ = v52
                                }
                                if v17 {
                                    var v53 int64 = 1
                                    v26 = 18446744073709551615
                                    var v54 bool = v17
                                    if (v17 && (!false)) {
                                        v8 = v8
                                        fmt.Println(v2, v49)
                                    } else {
                                        if (v17 || false) {
                                            delete(v1, (fi32 & (fi32 >> ((264) & 63))))
                                            v2 -= v21
                                            v8 = "hello"
                                        } else {
                                            v55 := v1[(fi32 << ((65535) & 63))]
                                            switch fi32 {
                                            case 4294967295:
                                                v56 := []int{}
                                                v53 = v53
                                                if (v2 != float32((-2.25 - ((3.14159 / (3.14159 - 1e10)) * 3.14159)))) {
                                                    if ((672 == (v26 % ((v26) | 1))) || ((v26 * (v26 * (^(v26)))) >= 42)) {
                                                        var v57 uint = 4294967295
                                                        fmt.Println(v49)
                                                        _ = v57
                                                    } else {
                                                        var v58 bool = v55
                                                        v53 = (v53 & (v53 % ((v53) | 1)))
                                                        _ = v58
                                                    }
                                                    var v59 int8 = -1
                                                    _ = v59
                                                }
                                                _ = v56
                                            default:
                                                fmt.Println(fi32, v55, v50)
                                            }
                                            v50 = ((v8 + v8) + ((v50 + "foo bar") + v8))
                                            _ = v55
                                        }
                                        v60 := v1[-2147483648]
                                        _ = v60
                                    }
                                    _ = v53
                                    _ = v54
                                }
                            } else {
                                v50 = v50
                                var v61 uint32 = 1
                                var v62 uint64 = v26
                                _ = v61
                                _ = v62
                            }
                        case "a":
                            if true {
                                v21 += v2
                                if v17 {
                                    v63 := []int{fi32}
                                    v64 := GenS9{F0: 2, F1: -9223372036854775808}
                                    _ = len(v1)
                                    fmt.Println(v49, v17, fi32)
                                    _ = v63
                                    _ = v64
                                } else {
                                    fmt.Println(v50, v26, v21)
                                    if (((v49 + v49) >= (v49 | v49)) && ((v26 >> ((4294967295) & 63)) >= v26)) {
                                        v49 = v49
                                    } else {
                                        if v17 {
                                            v50 = v50
                                        }
                                    }
                                    v65 := GenS4{F0: 1, F1: 186, F2: 49}
                                    _ = v65
                                }
                                var v66 string = (((v50 + "") + v8) + v8)
                                _ = v66
                            } else {
                                var v67 float32 = 3.14159
                                v2 += 3.14159
                                if (v17 || ((v21 - v2) != 1.5)) {
                                    var v68 uint = 1
                                    v69 := v3[0]
                                    _ = v68
                                    _ = v69
                                } else {
                                    switch v8 {
                                    case "\n\t":
                                        if false {
                                            if false {
                                                delete(v1, (fi33 & (-(fi32))))
                                                delete(v1, (fi33 << ((978) & 63)))
                                                fmt.Println(v8)
                                            } else {
                                                if false {
                                                    fmt.Println(v17)
                                                    if (105 == 1) {
                                                        if (v49 == 0) {
                                                            v70 := []int{(-(fi32)), fi32, (fi33 - (fi32 / (((fi32 & fi33)) | 1)))}
                                                            _ = v70
                                                        }
                                                    } else {
                                                        var v71 bool = (false && (33 != 1))
                                                        if (v71 || v71) {
                                                            fmt.Println(fi33, v17)
                                                        } else {
                                                            switch fi32 {
                                                            case -949:
                                                                if v17 {
                                                                    fmt.Println(v2)
                                                                    var v72 int64 = -762
                                                                    _ = v72
                                                                }
                                                            case -9223372036854775808:
                                                                v71 = (!false)
                                                            case 1:
                                                                switch (fi33 | fi33) {
                                                                case -236:
                                                                    if v17 {
                                                                        v73 := GenS7{F0: 100}
                                                                        _ = v73
                                                                    } else {
                                                                        v17 = (((0.1 * 3.14159) >= 0.1) || false)
                                                                        var v74 uint = 89
                                                                        _ = v74
                                                                    }
                                                                case 2147483647:
                                                                    if v71 {
                                                                        if (v71 && (v26 == v26)) {
                                                                            v75 := []int{fi32, (fi32 * 128), (^(fi32))}
                                                                            v21 += v21
                                                                            _ = v75
                                                                        } else {
                                                                            if true {
                                                                                var v76 int16 = v49
                                                                                _ = v76
                                                                            }
                                                                            v77 := GenS10{F0: v50}
                                                                            if (!(-2.25 >= 0.0)) {
                                                                                v21 = 3.14159
                                                                                v78 := GenS11{F0: 0, F1: -9223372036854775808, F2: 914}
                                                                                _ = v78
                                                                            } else {
                                                                                if false {
                                                                                    if v17 {
                                                                                        v26 -= 765
                                                                                        v17 = false
                                                                                    }
                                                                                } else {
                                                                                    v26 += v26
                                                                                    v79 := GenS11{F0: 1, F1: fi33, F2: (v26 % ((v26) | 1))}
                                                                                    _ = v79
                                                                                }
                                                                                var v80 uint16 = 295
                                                                                _ = v80
                                                                            }
                                                                            _ = v77
                                                                        }
                                                                    } else {
                                                                        switch fi32 {
                                                                        case 459:
                                                                            var v81 uint64 = (v26 ^ (v26 ^ v26))
                                                                            _ = v81
                                                                        default:
                                                                            v82 := []int{}
                                                                            _ = v82
                                                                        }
                                                                    }
                                                                case 0:
                                                                    v49 -= (v49 & 125)
                                                                default:
                                                                    var v83 string = (v8 + v50)
                                                                    _ = v83
                                                                }
                                                            default:
                                                                if (v71 || (106 > 1)) {
                                                                    var v84 float32 = float32((float64(v21) + (1.5 * 1e10)))
                                                                    var v85 int64 = 42
                                                                    _ = v84
                                                                    _ = v85
                                                                }
                                                            }
                                                            var v86 uint8 = 128
                                                            _ = v86
                                                        }
                                                        _ = v71
                                                    }
                                                    switch -965 {
                                                    case 561:
                                                        if (!true) {
                                                            var v87 float32 = v21
                                                            _ = v87
                                                        } else {
                                                            v88 := GenS5{F0: v17, F1: -2147483648, F2: 18446744073709551615}
                                                            v89 := GenS12{F0: v26}
                                                            _ = v88
                                                            _ = v89
                                                        }
                                                    case 128:
                                                        var v90 uint8 = 2
                                                        _ = v90
                                                    case -9223372036854775808:
                                                        v91 := []int{}
                                                        _ = v91
                                                    }
                                                } else {
                                                    delete(v3, fi33)
                                                    switch (fi33 / ((-358) | 1)) {
                                                    case -837:
                                                        v92 := []int{9223372036854775807, fi33}
                                                        _ = v92
                                                    case 42:
                                                        var v93 int64 = 9223372036854775807
                                                        var v94 bool = (v17 || v17)
                                                        v94 = ((!v17) || true)
                                                        _ = v93
                                                        _ = v94
                                                    }
                                                }
                                                v17 = true
                                                v95 := GenS7{F0: 0}
                                                _ = v95
                                            }
                                            delete(v1, fi32)
                                        } else {
                                            v8 = (v50 + ("a" + (v50 + (("a" + v8) + v50))))
                                            var v96 int64 = 491
                                            var v97 string = "hello"
                                            _ = v96
                                            _ = v97
                                        }
                                        var v98 uint32 = 1
                                        if (2 <= 977) {
                                            v98 = v98
                                            v21 = v21
                                        } else {
                                            if (!true) {
                                                fmt.Println(v17)
                                            } else {
                                                var v99 int64 = -468
                                                var v100 uint8 = 127
                                                var v101 string = v50
                                                _ = v100
                                                _ = v101
                                                _ = v99
                                            }
                                        }
                                        _ = v98
                                    case "hello":
                                        v50 = ("hello" + v8)
                                        var v102 uint16 = 562
                                        v103 := GenS5{F0: v17, F1: 0, F2: 566}
                                        _ = v102
                                        _ = v103
                                    case "":
                                        switch v50 {
                                        case "":
                                            fmt.Println(fi32)
                                        }
                                    default:
                                        fmt.Println(v50, v26, v2)
                                    }
                                }
                                v104 := GenS13{F0: v17, F1: 9223372036854775807}
                                v8 = v8
                                fmt.Println(v2)
                                v8 = (v8 + v8)
                                _ = v104
                                _ = v67
                            }
                        default:
                            v26 -= v26
                        }
                        v105 := GenS2{F0: 1, F1: genFn0((fi32 ^ fi33), 2147483647, (fi33 + int(v49)))}
                        if (32767 > 2) {
                            switch (fi33 | 127) {
                            case 127:
                                v106 := GenS4{F0: 121, F1: 42, F2: 2147483647}
                                _ = v106
                            case 583:
                                v107 := GenS12{F0: 1}
                                _ = v107
                            case 1:
                                var v108 int = (fi33 / ((fi32) | 1))
                                _ = v108
                            default:
                                if ((-1 > 42) || (3.14159 > 3.14159)) {
                                    var v109 float64 = 1.5
                                    if false {
                                        v2 += v2
                                        v17 = ((65535 > 2) && (!v17))
                                        v17 = (-2147483648 <= fi33)
                                    }
                                    _ = v109
                                } else {
                                    var v110 uint64 = (v26 << ((42) & 63))
                                    v17 = (v17 && (!((v49 >> ((269) & 63)) <= 547)))
                                    v50 = v8
                                    if (((3.14159 * 3.14159) * v21) > 1.5) {
                                        if true {
                                            var v111 float32 = v2
                                            switch (^(fi33)) {
                                            case 4294967295:
                                                if (v17 || v17) {
                                                    var v112 int16 = -512
                                                    _ = v112
                                                } else {
                                                    switch fi32 {
                                                    case -2147483648:
                                                        switch -269 {
                                                        case 9223372036854775807:
                                                            if (128 > 127) {
                                                                var v113 uint64 = v26
                                                                switch "" {
                                                                case "hello":
                                                                    fmt.Println(v21)
                                                                }
                                                                v49 = (v49 / ((1000) | 1))
                                                                _ = v113
                                                            }
                                                        default:
                                                            var v114 int32 = 42
                                                            _ = v114
                                                        }
                                                    case -1:
                                                        _ = len(v3)
                                                    default:
                                                        var v115 uint8 = 42
                                                        var v116 uint32 = 4294967295
                                                        switch v8 {
                                                        case "foo bar":
                                                            var v117 float64 = -2.25
                                                            _ = v117
                                                        }
                                                        v8 = "a"
                                                        if v17 {
                                                            if (v49 == v49) {
                                                                v21 = float32((1e10 - (((1e10 + float64((float32((((((1.5 * float64((v111 / v2))) - (1.5 + float64(-2.25))) + -2.25) * 3.14159) * 1.5)) - v111))) - 1e10) + 0.1)))
                                                                switch v8 {
                                                                case "a":
                                                                    var v118 uint16 = 2
                                                                    _ = v118
                                                                case "\n\t":
                                                                    _ = len(v1)
                                                                case "hello":
                                                                    if (65535 < uint(v116)) {
                                                                        fmt.Println(fi32)
                                                                        v8 = ("foo bar" + v8)
                                                                        v17 = v17
                                                                    } else {
                                                                        if v17 {
                                                                            v2 = (v2 - 0.1)
                                                                            v8 = v8
                                                                        } else {
                                                                            var v119 uint32 = v116
                                                                            _ = v119
                                                                        }
                                                                    }
                                                                }
                                                            } else {
                                                                v8 = (v8 + "")
                                                                fmt.Println(v116, v111, v110)
                                                                v120 := GenS14{F0: v49}
                                                                _ = v120
                                                            }
                                                            fmt.Println(fi32, v116, v49)
                                                            var v121 int32 = 42
                                                            _ = v121
                                                        }
                                                        v122 := GenS14{F0: (v49 / ((v49) | 1))}
                                                        v21 += v21
                                                        _ = v115
                                                        _ = v116
                                                        _ = v122
                                                    }
                                                    switch (v8 + v8) {
                                                    case "":
                                                        var v123 int32 = int32(fi33)
                                                        _ = v123
                                                    case "foo bar":
                                                        if ((-128 <= 0) || (4294967295 < (^(v110)))) {
                                                            var v124 int = (fi32 ^ fi33)
                                                            _ = v124
                                                        } else {
                                                            var v125 int8 = 100
                                                            if ((!false) && v17) {
                                                                fmt.Println(v2, v110)
                                                                var v126 bool = true
                                                                v2 = (v111 * v111)
                                                                _ = v126
                                                            } else {
                                                                v125 = (v125 >> ((0) & 63))
                                                                if (1.5 >= 1.5) {
                                                                    v50 = v8
                                                                    _ = len(v1)
                                                                    var v127 int64 = 9223372036854775807
                                                                    _ = v127
                                                                } else {
                                                                    v128 := []int{fi32, (-(fi33)), 9223372036854775807}
                                                                    _ = v128
                                                                }
                                                            }
                                                            _ = v125
                                                        }
                                                    }
                                                }
                                            case 42:
                                                if v17 {
                                                    switch "a" {
                                                    case "hello":
                                                        switch fi33 {
                                                        case 32768:
                                                            var v129 float32 = v21
                                                            if ((v17 && (v17 || ((false && v17) && v17))) && (0 <= 42)) {
                                                                v130 := []int{fi33, (fi33 * fi32)}
                                                                if v17 {
                                                                    var v131 uint32 = 4294967295
                                                                    v17 = true
                                                                    _ = v131
                                                                }
                                                                var v132 uint8 = 2
                                                                _ = v130
                                                                _ = v132
                                                            } else {
                                                                v133 := []int{fi32, (fi32 >> ((65535) & 63))}
                                                                v134 := GenS7{F0: 96}
                                                                _ = v133
                                                                _ = v134
                                                            }
                                                            var v135 int32 = -404
                                                            _ = v129
                                                            _ = v135
                                                        case 32767:
                                                            switch "a" {
                                                            case "":
                                                                v136 := []int{-600, (fi33 >> ((2) & 63))}
                                                                var v137 int8 = 127
                                                                v21 += 1.5
                                                                _ = v136
                                                                _ = v137
                                                            case "a":
                                                                v111 += (0.1 - (v2 + (0.0 * 1.5)))
                                                            case "foo bar":
                                                                v111 = (((v111 - v111) / ((((((v2 - 1e10) + 1.5) * float32(-2.25)) + 1e10) / 1.5) + (v111 * (((-2.25 - 0.0) - v2) + v21)))) + (v21 * 1.5))
                                                            }
                                                        case 812:
                                                            v138 := GenS1{F0: 1, F1: 255}
                                                            _ = v138
                                                        default:
                                                            fmt.Println(v50, fi33)
                                                        }
                                                    default:
                                                        if (0 >= 1) {
                                                            var v139 int32 = 2147483647
                                                            v139 = v139
                                                            v111 += ((v21 * -2.25) + ((((v2 - v2) / ((v2 - (v2 - v2)) / float32((3.14159 + (1.5 + ((float64(v2) + ((float64(v111) * -2.25) - ((((-2.25 + float64(float32(((3.14159 + (3.14159 + float64(float32(1.5)))) / float64(v111))))) - 0.0) + -2.25) + 1.5))) * float64((v21 - -2.25)))))))) * (((v111 * 3.14159) / 1.0) / -2.25)) + 1e10))
                                                            _ = v139
                                                        } else {
                                                            if (!(!true)) {
                                                                var v140 int8 = -1
                                                                var v141 int32 = 1
                                                                _ = v140
                                                                _ = v141
                                                            }
                                                            v142 := GenS8{F0: 166, F1: fi32, F2: v49}
                                                            var v143 int64 = 42
                                                            _ = v142
                                                            _ = v143
                                                        }
                                                    }
                                                }
                                            case 1:
                                                v144 := GenS12{F0: v110}
                                                _ = v144
                                            }
                                            var v145 string = "foo bar"
                                            _ = v111
                                            _ = v145
                                        } else {
                                            if v17 {
                                                switch "foo bar" {
                                                case "hello":
                                                    fmt.Println(fi33, v17, v2)
                                                    if (32767 <= 962) {
                                                        v146 := GenS2{F0: -2147483648, F1: 4294967295}
                                                        var v147 uint = 42
                                                        if v17 {
                                                            v148 := GenS15{F0: -2147483648}
                                                            _ = v148
                                                        } else {
                                                            switch ("\n\t" + "a") {
                                                            case "foo bar":
                                                                fmt.Println(v2, v50)
                                                            default:
                                                                v147 = 255
                                                            }
                                                            if v17 {
                                                                delete(v3, 128)
                                                            }
                                                        }
                                                        v110 = v26
                                                        v8 = v50
                                                        _ = v146
                                                        _ = v147
                                                    }
                                                    var v149 uint64 = 0
                                                    _ = v149
                                                }
                                            } else {
                                                v1[(fi33 % ((fi33) | 1))] = (!(42 <= -96))
                                                var v150 int = fi33
                                                switch v50 {
                                                case "foo bar":
                                                    v151 := GenS13{F0: v17, F1: 9223372036854775807}
                                                    _ = v151
                                                case "\n\t":
                                                    var v152 uint8 = 1
                                                    _ = v152
                                                case "hello":
                                                    var v153 uint = 1
                                                    _ = v153
                                                }
                                                _ = v150
                                            }
                                        }
                                        fmt.Println(v2)
                                        v110 = v110
                                    }
                                    _ = v110
                                }
                            }
                        }
                        var v154 int = 9223372036854775807
                        var v155 int32 = 42
                        _ = v105
                        _ = v154
                        _ = v155
                    }
                    var v156 int = fi32
                    v157 := []int{(v156 | fi33), (v156 - 4294967295), (fi33 % (((^(fi33))) | 1)), (fi32 / ((fi32) | 1))}
                    _ = v156
                    _ = v157
                    _ = v49
                    _ = v50
                }
                var v158 int64 = 1
                if (!false) {
                    var v159 int = (fi32 + 579)
                    _ = v159
                } else {
                    fmt.Println(v158, v8, v8)
                }
                var v160 uint8 = 128
                fmt.Println(v160, v158, fi32)
                var v161 uint32 = 138
                var v162 uint16 = 2
                var v163 uint16 = v162
                _ = v158
                _ = v160
                _ = v161
                _ = v162
                _ = v163
                _ = v42
            }
            v21 -= v2
            _ = fi33
        }
        for fi164 := 0; fi164 < 12; fi164++ {
            var v165 int16 = 1000
            _ = fi164
            _ = v165
        }
        _ = fi32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v8
}
