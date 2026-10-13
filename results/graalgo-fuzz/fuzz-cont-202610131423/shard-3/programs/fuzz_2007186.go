package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 uint8; F2 int32 }
type GenS6 struct { F0 uint; F1 uint16; F2 bool }
type GenS7 struct { F0 int32; F1 int32; F2 string }
type GenS8 struct { F0 int; F1 uint8 }
type GenS9 struct { F0 bool; F1 uint32; F2 int }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 bool; F1 bool }
type GenS13 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS14 struct { F0 uint8; F1 uint64 }
func genFn0() uint {
    var v6 int16 = 1000
    var v7 int8 = 42
    v8 := []int{-385, -879, 340}
    v9 := GenS1{F0: 32767, F1: -462, F2: 1}
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 792
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]uint32)
    var v3 bool = false
    var v4 int = 9223372036854775807
    var v5 uint8 = 1
    v1 += ((float64(1.5) + float64(-2.25)) / 0.1)
    var v10 uint = 42
    var v11 string = "\n\t"
    v12 := v2[(v4 << ((v10) & 63))]
    v13 := GenS1{F0: 65535, F1: -1, F2: 1}
    v14 := GenS2{F0: v12, F1: -699, F2: 255}
    v10 = v10
    switch v11 {
    case "":
        var v15 float32 = (float32((0.1 * v1)) + 0.0)
        _ = v15
    default:
        if v3 {
            if true {
                v16 := GenS2{F0: (v12 / ((uint32(v10)) | 1)), F1: -878, F2: 596}
                _ = v16
            } else {
                v11 = "\n\t"
            }
        }
    }
    delete(v2, (v4 << ((2) & 63)))
    v1 += (3.14159 - 3.14159)
    if v3 {
        if false {
            if v3 {
                for fi17 := 0; fi17 < 14; fi17++ {
                    v5 -= v5
                    _ = fi17
                }
            }
            var v18 uint16 = 267
            _ = v18
        } else {
            switch 1 {
            case -295:
                if (1 >= 127) {
                    switch ("a" + "") {
                    case "\n\t":
                        v19 := []int{(v4 << (((v10 + (v10 - 4294967295))) & 63)), 32768, v4, (v4 & (-(v4)))}
                        _ = v19
                    }
                }
            case 376:
                if v3 {
                    for fi20 := 0; fi20 < 12; fi20++ {
                        v21 := GenS2{F0: (v12 & (v12 >> ((genFn0()) & 63))), F1: 1, F2: 65535}
                        _ = fi20
                        _ = v21
                    }
                    if (v1 != (v1 - v1)) {
                        v5 = 127
                        var v22 int32 = 1
                        _ = v22
                    } else {
                        var v23 int = (v4 & (-(v4)))
                        var v24 bool = ((636 > 2) || (-128 <= -98))
                        v25 := GenS1{F0: 842, F1: int32(v5), F2: 32767}
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                    var v26 int = v4
                    _ = v26
                } else {
                    var v27 uint32 = 2147483647
                    switch (v11 + v11) {
                    case "hello":
                        fmt.Println(v10, v12)
                    case "\n\t":
                        var v28 uint32 = v27
                        _ = v28
                    case "a":
                        switch v11 {
                        case "\n\t":
                            v10 += (v10 / ((genFn0()) | 1))
                        case "foo bar":
                            v27 = v27
                        }
                    default:
                        for fi29 := 0; fi29 < 7; fi29++ {
                            switch (fi29 ^ (fi29 * fi29)) {
                            case 32768:
                                var v30 float32 = float32(v1)
                                _ = v30
                            default:
                                v10 = v10
                            }
                            v12 -= 306
                            for fi31 := 0; fi31 < 15; fi31++ {
                                var v32 bool = ((^(v12)) != uint32(v5))
                                v33 := GenS1{F0: 980, F1: 48, F2: 789}
                                fmt.Println(v32)
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                            _ = fi29
                        }
                    }
                    for fi34 := 0; fi34 < 12; fi34++ {
                        for fi35 := 0; fi35 < 5; fi35++ {
                            v3 = (false || (!v3))
                            if (!true) {
                                v10 = v10
                            }
                            _ = len(v2)
                            if false {
                                v36 := []int{fi34, v4, fi34, v4}
                                _ = v36
                            } else {
                                if (0 == 0) {
                                    v11 = v11
                                } else {
                                    fmt.Println(fi34, fi34)
                                    v37 := v2[(^(v4))]
                                    var v38 int16 = 1000
                                    fmt.Println(v10, v3)
                                    var v39 uint16 = 2
                                    _ = v37
                                    _ = v38
                                    _ = v39
                                }
                            }
                            v4 = v4
                            _ = fi35
                        }
                        for fi40 := 0; fi40 < 15; fi40++ {
                            if ((!(!v3)) || v3) {
                                var v41 uint64 = 2
                                if v3 {
                                    v41 = (^(v41))
                                    v42 := GenS3{F0: int32(v4), F1: (678 >= 2)}
                                    v43 := GenS2{F0: v12, F1: -1, F2: 246}
                                    _ = v42
                                    _ = v43
                                } else {
                                    if true {
                                        switch -9223372036854775808 {
                                        case 381:
                                            var v44 int64 = 368
                                            _ = v44
                                        case 42:
                                            if v3 {
                                                var v45 float32 = 3.14159
                                                _ = v45
                                            } else {
                                                if true {
                                                    switch (v4 | fi40) {
                                                    case 2147483647:
                                                        if (-1 <= 1) {
                                                            var v46 uint16 = 407
                                                            _ = v46
                                                        } else {
                                                            if (false && false) {
                                                                var v47 uint64 = v41
                                                                v11 = ((((v11 + "hello") + v11) + (v11 + "\n\t")) + "foo bar")
                                                                var v48 int8 = 100
                                                                _ = v47
                                                                _ = v48
                                                            }
                                                            v49 := GenS4{F0: (^(v4))}
                                                            v50 := v2[(v4 & (-(v4)))]
                                                            _ = v49
                                                            _ = v50
                                                        }
                                                    case 1:
                                                        if v3 {
                                                            if v3 {
                                                                v41 = (v41 % ((v41) | 1))
                                                                if (2 > 32767) {
                                                                    var v51 int8 = 42
                                                                    var v52 uint = v10
                                                                    fmt.Println(v27, v12, fi34)
                                                                    _ = v51
                                                                    _ = v52
                                                                } else {
                                                                    var v53 bool = (42 <= (^(v5)))
                                                                    _ = v53
                                                                }
                                                                var v54 uint = (v10 | genFn0())
                                                                _ = v54
                                                            }
                                                            v55 := GenS5{F0: -1, F1: v5, F2: 75}
                                                            _ = v55
                                                        } else {
                                                            if v3 {
                                                                if false {
                                                                    switch (v11 + ("\n\t" + v11)) {
                                                                    case "\n\t":
                                                                        if ("\n\t" != "hello") {
                                                                            v12 = uint32(v5)
                                                                            switch v11 {
                                                                            case "hello":
                                                                                if (0 != -95) {
                                                                                    fmt.Println(v5, v3)
                                                                                    var v56 uint32 = v12
                                                                                    v57 := GenS6{F0: genFn0(), F1: 1, F2: v3}
                                                                                    _ = v56
                                                                                    _ = v57
                                                                                } else {
                                                                                    var v58 uint64 = 2
                                                                                    var v59 uint16 = 1
                                                                                    _ = v58
                                                                                    _ = v59
                                                                                }
                                                                            case "\n\t":
                                                                                v60 := []int{369, 265, 9223372036854775807}
                                                                                _ = v60
                                                                            }
                                                                            v61 := GenS7{F0: 42, F1: 0, F2: v11}
                                                                            _ = v61
                                                                        } else {
                                                                            v3 = v3
                                                                            var v62 int16 = 42
                                                                            _ = v62
                                                                        }
                                                                    case "a":
                                                                        var v63 int8 = 3
                                                                        _ = v63
                                                                    case "hello":
                                                                        v64 := []int{(^(fi40)), (fi40 | fi34), 9223372036854775807, v4}
                                                                        _ = v64
                                                                    default:
                                                                        if (!(((v11 + ("hello" + "")) != v11) && false)) {
                                                                            switch v4 {
                                                                            case -255:
                                                                                v65 := []int{v4, (fi40 + (fi40 + (fi34 >> ((uint(v41)) & 63))))}
                                                                                _ = v65
                                                                            }
                                                                            switch "\n\t" {
                                                                            case "\n\t":
                                                                                delete(v2, (^(fi34)))
                                                                            case "hello":
                                                                                fmt.Println(v41)
                                                                            case "a":
                                                                                if v3 {
                                                                                    fmt.Println(v5)
                                                                                    switch ((v11 + v11) + "") {
                                                                                    case "\n\t":
                                                                                        var v66 uint = genFn0()
                                                                                        _ = v66
                                                                                    case "hello":
                                                                                        if (v3 && (0 < -89)) {
                                                                                            v67 := GenS4{F0: fi34}
                                                                                            fmt.Println(v11, v41, v11)
                                                                                            _ = v67
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    v4 -= int(v10)
                                                                                    var v68 int = fi34
                                                                                    _ = v68
                                                                                }
                                                                            default:
                                                                                if (32767 <= -32768) {
                                                                                    if true {
                                                                                        var v69 int64 = 0
                                                                                        if v3 {
                                                                                            if (v5 != 0) {
                                                                                                fmt.Println(fi40)
                                                                                            } else {
                                                                                                _ = len(v2)
                                                                                                if false {
                                                                                                    if (1e10 <= (v1 - float64((0.0 - 0.1)))) {
                                                                                                        v70 := GenS6{F0: v10, F1: 231, F2: ((v69 * v69) != v69)}
                                                                                                        fmt.Println(v4, v3, v5)
                                                                                                        _ = v70
                                                                                                    } else {
                                                                                                        if ((v3 && (-128 > 102)) || ((v3 && v3) && v3)) {
                                                                                                            var v71 uint = v10
                                                                                                            fmt.Println(v12, v41, fi40)
                                                                                                            v72 := GenS2{F0: 867, F1: 1000, F2: 2}
                                                                                                            v69 = -915
                                                                                                            v41 -= (v41 - (v41 << (((v10 % ((v71) | 1))) & 63)))
                                                                                                            _ = v71
                                                                                                            _ = v72
                                                                                                        } else {
                                                                                                            v69 += v69
                                                                                                        }
                                                                                                    }
                                                                                                } else {
                                                                                                    var v73 uint = (v10 % ((753) | 1))
                                                                                                    var v74 uint16 = 301
                                                                                                    v12 = v27
                                                                                                    _ = v73
                                                                                                    _ = v74
                                                                                                }
                                                                                                v10 -= v10
                                                                                            }
                                                                                            var v75 float64 = 0.1
                                                                                            _ = v75
                                                                                        } else {
                                                                                            v10 = (v10 * genFn0())
                                                                                            var v76 string = v11
                                                                                            _ = v76
                                                                                        }
                                                                                        v77 := GenS1{F0: uint16(v12), F1: -266, F2: uint16(v27)}
                                                                                        _ = v69
                                                                                        _ = v77
                                                                                    } else {
                                                                                        v10 = uint(v27)
                                                                                        _ = len(v2)
                                                                                        v78 := v2[v4]
                                                                                        _ = v78
                                                                                    }
                                                                                    v79 := GenS8{F0: (v4 - (fi34 * (v4 & fi40))), F1: v5}
                                                                                    v27 += (v12 & 2147483648)
                                                                                    _ = v79
                                                                                } else {
                                                                                    v41 += (v41 | v41)
                                                                                    v80 := []int{fi40, (^(fi40)), (fi40 + 817), (v4 / (((fi40 + v4)) | 1))}
                                                                                    var v81 uint = (v10 + (v10 << ((genFn0()) & 63)))
                                                                                    fmt.Println(v81)
                                                                                    if len(v80) >= 2 {
                                                                                        v80 = v80[1:len(v80)]
                                                                                    }
                                                                                    var v82 int32 = -660
                                                                                    switch v4 {
                                                                                    case 4294967295:
                                                                                        switch 32767 {
                                                                                        case 422:
                                                                                            var v83 float32 = -2.25
                                                                                            if true {
                                                                                                _ = len(v80)
                                                                                            } else {
                                                                                                var v84 bool = (v5 <= (v5 | v5))
                                                                                                v81 = 65535
                                                                                                _ = v84
                                                                                            }
                                                                                            var v85 float32 = 0.0
                                                                                            _ = v83
                                                                                            _ = v85
                                                                                        case 63:
                                                                                            var v86 int16 = 42
                                                                                            _ = v86
                                                                                        case 4294967295:
                                                                                            v27 += uint32(v41)
                                                                                        default:
                                                                                            var v87 int = 4294967295
                                                                                            _ = v87
                                                                                        }
                                                                                        var v88 int8 = -59
                                                                                        if true {
                                                                                            if false {
                                                                                                switch "hello" {
                                                                                                case "\n\t":
                                                                                                    var v89 int = fi40
                                                                                                    _ = v89
                                                                                                case "hello":
                                                                                                    if v3 {
                                                                                                        if (uint64(v10) != 42) {
                                                                                                            if true {
                                                                                                                var v90 uint16 = 0
                                                                                                                _ = v90
                                                                                                            }
                                                                                                        } else {
                                                                                                            if (!((v82 >= -428) && (true && (((^(v88)) == v88) || (!((uint(v41) < v10) || (305 <= 65535))))))) {
                                                                                                                if v3 {
                                                                                                                    v88 -= v88
                                                                                                                }
                                                                                                                v91 := GenS9{F0: true, F1: (v12 & (v12 * 42)), F2: fi34}
                                                                                                                _ = v91
                                                                                                            } else {
                                                                                                                var v92 string = v11
                                                                                                                _ = v92
                                                                                                            }
                                                                                                        }
                                                                                                        var v93 string = v11
                                                                                                        if (804 > (v41 / ((1) | 1))) {
                                                                                                            if false {
                                                                                                                v94 := GenS10{F0: 1000}
                                                                                                                if ((v88 / ((-1) | 1)) <= (-(v88))) {
                                                                                                                    v4 = v4
                                                                                                                }
                                                                                                                _ = v94
                                                                                                            }
                                                                                                        } else {
                                                                                                            var v95 uint32 = 42
                                                                                                            _ = v95
                                                                                                        }
                                                                                                        _ = v93
                                                                                                    }
                                                                                                case "a":
                                                                                                    v82 -= 604
                                                                                                default:
                                                                                                    var v96 uint32 = 12
                                                                                                    _ = v96
                                                                                                }
                                                                                                v97 := GenS11{F0: 146}
                                                                                                v3 = v3
                                                                                                _ = v97
                                                                                            }
                                                                                            _ = len(v80)
                                                                                            _ = len(v2)
                                                                                        }
                                                                                        _ = v88
                                                                                    case 740:
                                                                                        switch fi40 {
                                                                                        case 32768:
                                                                                            var v98 uint8 = 135
                                                                                            _ = v98
                                                                                        }
                                                                                    default:
                                                                                        var v99 float64 = (v1 * v1)
                                                                                        _ = v99
                                                                                    }
                                                                                    switch (v11 + v11) {
                                                                                    case "a":
                                                                                        v27 = 699
                                                                                    default:
                                                                                        var v100 uint = (v81 << (((^(v81))) & 63))
                                                                                        _ = v100
                                                                                    }
                                                                                    _ = v80
                                                                                    _ = v81
                                                                                    _ = v82
                                                                                }
                                                                            }
                                                                            switch "a" {
                                                                            case "a":
                                                                                v27 = 0
                                                                            case "hello":
                                                                                fmt.Println(v4, v27, v12)
                                                                            case "\n\t":
                                                                                v10 = 588
                                                                            }
                                                                        }
                                                                    }
                                                                    v11 = (v11 + ((v11 + v11) + "foo bar"))
                                                                    if v3 {
                                                                        v101 := v2[fi40]
                                                                        v102 := GenS9{F0: (!false), F1: v27, F2: v4}
                                                                        switch (v11 + ("foo bar" + (v11 + ("a" + ((v11 + v11) + v11))))) {
                                                                        case "":
                                                                            var v103 bool = v3
                                                                            _ = v103
                                                                        case "hello":
                                                                            var v104 uint64 = 455
                                                                            _ = v104
                                                                        case "a":
                                                                            fmt.Println(v1, fi40)
                                                                        default:
                                                                            switch ("" + v11) {
                                                                            case "hello":
                                                                                switch "\n\t" {
                                                                                case "foo bar":
                                                                                    v105 := GenS5{F0: 225, F1: v5, F2: -122}
                                                                                    _ = v105
                                                                                case "":
                                                                                    v106 := GenS3{F0: 2147483647, F1: (v11 >= ((v11 + "") + (("a" + (v11 + ("" + (((v11 + "a") + (v11 + "a")) + (("foo bar" + ("\n\t" + ((v11 + v11) + "a"))) + ((v11 + (v11 + "a")) + "foo bar")))))) + v11)))}
                                                                                    _ = v106
                                                                                default:
                                                                                    var v107 float64 = 0.1
                                                                                    _ = v107
                                                                                }
                                                                            case "":
                                                                                var v108 int8 = 42
                                                                                _ = v108
                                                                            default:
                                                                                v11 = (v11 + (v11 + (v11 + ("foo bar" + "a"))))
                                                                            }
                                                                        }
                                                                        v5 = v5
                                                                        var v109 uint32 = v12
                                                                        _ = v101
                                                                        _ = v102
                                                                        _ = v109
                                                                    }
                                                                    if ((v5 ^ (v5 & (v5 ^ v5))) != (^(v5))) {
                                                                        v4 += fi40
                                                                    } else {
                                                                        v12 = (v12 << ((genFn0()) & 63))
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    case -657:
                                                        v4 -= 32767
                                                        var v110 int64 = -1
                                                        _ = v110
                                                    default:
                                                        if (v3 || v3) {
                                                            v11 = v11
                                                        } else {
                                                            if (9 != -1) {
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    v111 := GenS5{F0: 2147483647, F1: 93, F2: 2147483647}
                                                                    v3 = (float32((-2.25 * 1e10)) != float32(v1))
                                                                    _ = v111
                                                                default:
                                                                    delete(v2, (^(fi34)))
                                                                }
                                                            }
                                                            var v112 uint = genFn0()
                                                            fmt.Println(fi34, v5)
                                                            _ = v112
                                                        }
                                                    }
                                                    v113 := GenS6{F0: 1, F1: 65535, F2: ((v11 + v11) <= "a")}
                                                    var v114 string = ""
                                                    delete(v2, fi34)
                                                    v115 := GenS12{F0: v3, F1: v3}
                                                    v114 = "hello"
                                                    _ = v113
                                                    _ = v114
                                                    _ = v115
                                                }
                                                if true {
                                                    if (!(46 > 42)) {
                                                        v116 := GenS13{F0: uint32(v10), F1: 42, F2: (v5 % (((v5 << ((genFn0()) & 63))) | 1))}
                                                        _ = v116
                                                    }
                                                    _ = len(v2)
                                                    var v117 int8 = int8(v41)
                                                    _ = v117
                                                } else {
                                                    fmt.Println(v12, fi40)
                                                    if (18446744073709551615 == v41) {
                                                        v2[fi40] = (^(v27))
                                                        v10 -= 42
                                                    } else {
                                                        v4 += (fi34 * v4)
                                                        v27 = (^(v12))
                                                        var v118 uint32 = (^(v12))
                                                        _ = v118
                                                    }
                                                    var v119 int = 42
                                                    _ = v119
                                                }
                                                v120 := GenS14{F0: 2, F1: (v41 | (v41 | (v41 % (((v41 / ((v41) | 1))) | 1))))}
                                                _ = v120
                                            }
                                        default:
                                            var v121 int32 = 2147483647
                                            _ = v121
                                        }
                                        var v122 int32 = -530
                                        v1 = v1
                                        _ = v122
                                    } else {
                                        var v123 string = ("hello" + "")
                                        v124 := GenS9{F0: v3, F1: v27, F2: (fi34 + (fi34 * 223))}
                                        _ = v123
                                        _ = v124
                                    }
                                }
                                _ = v41
                            }
                            fmt.Println(v12, v11)
                            if v3 {
                                v125 := GenS14{F0: (^(v5)), F1: 945}
                                _ = v125
                            } else {
                                switch v4 {
                                case -9223372036854775808:
                                    var v126 int32 = 1
                                    _ = v126
                                default:
                                    var v127 int = 286
                                    _ = v127
                                }
                                v128 := GenS2{F0: v27, F1: 336, F2: 255}
                                _ = v128
                            }
                            _ = fi40
                        }
                        _ = fi34
                    }
                    _ = v27
                }
            default:
                v1 -= v1
            }
            var v129 float64 = 0.1
            _ = v129
        }
    } else {
        v130 := GenS13{F0: 1, F1: 2147483647, F2: 26}
        v2[(v4 + (^(v4)))] = v12
        v3 = v3
        _ = v130
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
