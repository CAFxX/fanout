package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint64; F2 int32 }
type GenS3 struct { F0 bool; F1 int; F2 uint }
type GenS4 struct { F0 int; F1 int8; F2 bool }
type GenS5 struct { F0 uint16; F1 int; F2 int }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 int64; F1 bool }
func genFn0() uint {
    var v6 uint8 = 255
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 5; fi8++ {
            fmt.Println(fi8, fi8, fi8)
            v9 := GenS1{F0: 132}
            _ = fi8
            _ = v9
        }
        v6 += 42
        _ = fi7
    }
    v10 := []int{-9223372036854775808, 32768, 127, -968}
    _ = v10
    _ = v6
    return 2
}
func genFn1(p0 uint64, p1 bool) int64 {
    p1 = false
    v11 := []int{-591, 1}
    _ = p0
    _ = p1
    _ = v11
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = 1.5
    v3 := make(map[int]uint32)
    v4 := make(map[int]string)
    var v5 uint64 = 18446744073709551615
    switch 4294967295 {
    case -751:
        v2 += v2
    case 128:
        delete(v3, 128)
    }
    for fi12 := 0; fi12 < 11; fi12++ {
        if false {
            v5 = 2
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 int64 = -1
                if (genFn0() > 4294967295) {
                    v15 := GenS1{F0: 4294967295}
                    if (32767 <= -892) {
                        switch ("hello" + ("\n\t" + "\n\t")) {
                        case "foo bar":
                            var v16 int64 = genFn1(699, (0 > fi13))
                            _ = v16
                        case "":
                            var v17 int = (fi13 | (-(fi12)))
                            _ = v17
                        }
                        var v18 int32 = 1
                        if true {
                            v14 += (^(v14))
                            var v19 string = "a"
                            _ = v19
                        } else {
                            if false {
                                var v20 int8 = -128
                                var v21 uint64 = v5
                                var v22 int64 = v14
                                _ = v20
                                _ = v21
                                _ = v22
                            }
                            switch fi12 {
                            case 4294967295:
                                var v23 int = (-(fi12))
                                _ = v23
                            case 0:
                                var v24 int64 = genFn1((v5 - v5), ((("foo bar" + ("\n\t" + "foo bar")) + "a") > "a"))
                                _ = v24
                            case 509:
                                switch "\n\t" {
                                case "foo bar":
                                    v25 := GenS1{F0: 2147483647}
                                    _ = v25
                                default:
                                    var v26 int64 = genFn1((^(v5)), (65535 == 710))
                                    _ = v26
                                }
                            }
                            switch 32767 {
                            case 0:
                                var v27 uint32 = 0
                                _ = v27
                            case 42:
                                var v28 int16 = -1
                                v5 = (v5 >> ((0) & 63))
                                switch 32767 {
                                case 223:
                                    if false {
                                        _ = len(v1)
                                        if ((v28 & v28) < v28) {
                                            switch "\n\t" {
                                            case "foo bar":
                                                fmt.Println(v28)
                                            case "hello":
                                                v29 := GenS2{F0: v28, F1: 4294967295, F2: 42}
                                                _ = v29
                                            }
                                            v28 += v28
                                            fmt.Println(v18, v28, v28)
                                            fmt.Println(fi12)
                                            _ = len(v4)
                                            v2 += v2
                                            fmt.Println(v5)
                                            if true {
                                                var v30 bool = (false || (v18 == v18))
                                                _ = v30
                                            } else {
                                                switch fi13 {
                                                case 1:
                                                    delete(v1, fi13)
                                                }
                                                if ((false && false) || (120 <= 0)) {
                                                    if false {
                                                        if (!true) {
                                                            v31 := GenS2{F0: (v28 / (((v28 & v28)) | 1)), F1: v5, F2: v18}
                                                            v14 += v14
                                                            var v32 uint8 = 14
                                                            var v33 float32 = v2
                                                            _ = v31
                                                            _ = v32
                                                            _ = v33
                                                        }
                                                    } else {
                                                        if true {
                                                            v2 += -2.25
                                                            v2 = ((((0.0 * 3.14159) - (float32(-2.25) + (float32((3.14159 + (-2.25 * 3.14159))) * float32(float64(v2))))) + (3.14159 / v2)) - (v2 - 0.0))
                                                            fmt.Println(v18)
                                                        } else {
                                                            fmt.Println(v14, fi12, v2)
                                                            var v34 int64 = genFn1((v5 & (v5 * v5)), (!false))
                                                            _ = v34
                                                        }
                                                        switch "\n\t" {
                                                        case "":
                                                            var v35 float64 = 0.1
                                                            _ = v35
                                                        case "\n\t":
                                                            fmt.Println(fi12, v14, fi13)
                                                        }
                                                    }
                                                    v36 := GenS2{F0: v28, F1: v5, F2: 96}
                                                    _ = v36
                                                } else {
                                                    v28 -= (v28 & v28)
                                                    fmt.Println(fi12)
                                                    if ((true && (!true)) || false) {
                                                        var v37 uint32 = 42
                                                        _ = v37
                                                    } else {
                                                        var v38 uint64 = v5
                                                        _ = v38
                                                    }
                                                }
                                                v39 := GenS2{F0: v28, F1: (v5 ^ v5), F2: v18}
                                                _ = v39
                                            }
                                            switch fi12 {
                                            case 0:
                                                var v40 int = int(v18)
                                                _ = v40
                                            case -9223372036854775808:
                                                if true {
                                                    if (!true) {
                                                        var v41 float32 = 0.1
                                                        if true {
                                                            switch "a" {
                                                            case "hello":
                                                                v42 := []int{}
                                                                _ = v42
                                                            default:
                                                                if true {
                                                                    var v43 uint = 255
                                                                    _ = v43
                                                                }
                                                            }
                                                        }
                                                        v44 := GenS1{F0: 800}
                                                        _ = v41
                                                        _ = v44
                                                    } else {
                                                        var v45 float64 = 0.0
                                                        delete(v1, -9223372036854775808)
                                                        _ = v45
                                                    }
                                                } else {
                                                    var v46 string = "foo bar"
                                                    switch "foo bar" {
                                                    case "":
                                                        if (-32768 == v28) {
                                                            var v47 float64 = (1.5 - 1e10)
                                                            switch (fi12 - (fi13 | fi13)) {
                                                            case -2147483648:
                                                                if (752 > 293) {
                                                                    delete(v1, int(v5))
                                                                }
                                                                _ = len(v4)
                                                                switch v46 {
                                                                case "hello":
                                                                    v47 -= v47
                                                                    v47 = v47
                                                                    v18 += v18
                                                                    fmt.Println(v5, v18)
                                                                    var v48 int16 = v28
                                                                    var v49 uint = 945
                                                                    v14 += int64(fi13)
                                                                    v47 += 1.5
                                                                    fmt.Println(fi13, v18)
                                                                    var v50 int16 = (v28 ^ v28)
                                                                    _ = v48
                                                                    _ = v49
                                                                    _ = v50
                                                                case "":
                                                                    var v51 int = (fi12 - fi13)
                                                                    _ = v51
                                                                default:
                                                                    v52 := GenS2{F0: v28, F1: (v5 * (^(v5))), F2: v18}
                                                                    _ = v52
                                                                }
                                                            case 0:
                                                                v53 := GenS3{F0: false, F1: (^(fi13)), F2: uint(v5)}
                                                                _ = v53
                                                            case -1:
                                                                var v54 int = fi13
                                                                _ = v54
                                                            }
                                                            var v55 int16 = (v28 << ((173) & 63))
                                                            _ = v47
                                                            _ = v55
                                                        } else {
                                                            fmt.Println(fi13, v2)
                                                            if (!false) {
                                                                fmt.Println(v18, v46)
                                                            } else {
                                                                var v56 int8 = 127
                                                                v57 := GenS2{F0: (-(v28)), F1: v5, F2: v18}
                                                                _ = v56
                                                                _ = v57
                                                            }
                                                            v58 := GenS4{F0: 801, F1: -20, F2: false}
                                                            _ = v58
                                                        }
                                                    case "foo bar":
                                                        v4[fi12] = v46
                                                    case "a":
                                                        v59 := []int{(^(fi12))}
                                                        _ = v59
                                                    }
                                                    var v60 uint32 = 4294967295
                                                    if false {
                                                        var v61 int64 = v14
                                                        _ = v61
                                                    }
                                                    _ = v46
                                                    _ = v60
                                                }
                                            case 168:
                                                v14 = v14
                                            default:
                                                if (v2 > (-2.25 + ((-2.25 / v2) + v2))) {
                                                    switch ((("a" + "foo bar") + "hello") + "foo bar") {
                                                    case "foo bar":
                                                        var v62 uint16 = 2
                                                        _ = v62
                                                    case "hello":
                                                        v14 = (v14 >> ((2) & 63))
                                                    case "\n\t":
                                                        v63 := v4[42]
                                                        _ = v63
                                                    }
                                                } else {
                                                    v28 += 1000
                                                    v14 -= -9223372036854775808
                                                    if true {
                                                        v2 += v2
                                                        var v64 uint = 255
                                                        v65 := []int{(fi13 ^ fi12), (fi12 & (fi12 & (fi13 | (fi12 >> (((v64 % (((^(v64))) | 1))) & 63))))), 4294967295, (fi12 % ((fi12) | 1))}
                                                        _ = v64
                                                        _ = v65
                                                    } else {
                                                        var v66 string = "a"
                                                        if false {
                                                            var v67 int = fi12
                                                            delete(v1, 128)
                                                            _ = v67
                                                        }
                                                        switch fi12 {
                                                        case 32768:
                                                            fmt.Println(fi12)
                                                        }
                                                        _ = v66
                                                    }
                                                }
                                            }
                                        } else {
                                            v68 := []int{(fi13 + fi13), 538, 32768, (fi12 << ((genFn0()) & 63))}
                                            _ = v68
                                        }
                                        switch ("hello" + ("" + "a")) {
                                        case "foo bar":
                                            v2 = ((float32(3.14159) * float32(1.5)) / (v2 * 3.14159))
                                        case "a":
                                            v69 := GenS2{F0: int16(v18), F1: (v5 / (((v5 * v5)) | 1)), F2: -39}
                                            if (2 != genFn0()) {
                                                fmt.Println(v2)
                                                v14 -= -9223372036854775808
                                            }
                                            _ = v69
                                        case "hello":
                                            var v70 int8 = -1
                                            _ = v70
                                        default:
                                            var v71 int = (fi12 * 987)
                                            _ = v71
                                        }
                                        if (v5 >= (^(v5))) {
                                            v72 := GenS4{F0: fi12, F1: 127, F2: (false && true)}
                                            if false {
                                                var v73 int64 = (v14 >> ((genFn0()) & 63))
                                                if false {
                                                    fmt.Println(v2, fi13, v2)
                                                    var v74 bool = false
                                                    switch "a" {
                                                    case "":
                                                        v14 -= (v14 << ((42) & 63))
                                                    case "\n\t":
                                                        var v75 int8 = 42
                                                        _ = v75
                                                    case "a":
                                                        v2 += ((v2 * v2) + v2)
                                                    }
                                                    if ((-(v73)) != (v14 * (v14 << ((52) & 63)))) {
                                                        v76 := GenS2{F0: (v28 & v28), F1: 42, F2: -98}
                                                        v1[(^(fi13))] = 0
                                                        _ = v76
                                                    } else {
                                                        fmt.Println(v14)
                                                    }
                                                    v2 += v2
                                                    _ = v74
                                                }
                                                v18 -= (v18 / ((v18) | 1))
                                                var v77 int64 = (v73 / ((v73) | 1))
                                                switch "foo bar" {
                                                case "":
                                                    switch fi13 {
                                                    case -722:
                                                        fmt.Println(v14, fi12, fi12)
                                                    case 319:
                                                        switch "a" {
                                                        case "a":
                                                            v5 = 2
                                                        case "hello":
                                                            v78 := GenS5{F0: 65535, F1: (fi13 << ((980) & 63)), F2: fi13}
                                                            _ = v78
                                                        }
                                                    default:
                                                        var v79 uint8 = 208
                                                        delete(v1, -2147483648)
                                                        v79 = v79
                                                        v80 := []int{fi12, 0, fi13}
                                                        v81 := v1[(fi13 % ((fi12) | 1))]
                                                        _ = v79
                                                        _ = v80
                                                        _ = v81
                                                    }
                                                case "a":
                                                    v82 := GenS5{F0: 291, F1: fi13, F2: -848}
                                                    _ = v82
                                                case "\n\t":
                                                    v77 = 9223372036854775807
                                                    v28 += (v28 & v28)
                                                }
                                                switch fi13 {
                                                case -345:
                                                    if (255 > 42) {
                                                        if false {
                                                            fmt.Println(v28)
                                                            v18 = (-(v18))
                                                        } else {
                                                            v73 -= v77
                                                            fmt.Println(v18, v28, fi12)
                                                            v4[(fi12 * fi12)] = ("hello" + "\n\t")
                                                        }
                                                        var v83 uint32 = 1
                                                        if (!true) {
                                                            v84 := GenS2{F0: 32767, F1: (v5 + v5), F2: v18}
                                                            v85 := GenS5{F0: 32767, F1: fi12, F2: -924}
                                                            _ = v84
                                                            _ = v85
                                                        } else {
                                                            v86 := []int{(fi12 * (fi12 & -686)), fi13, (fi12 ^ fi12)}
                                                            v86 = append(v86, (fi12 / (((fi12 + (fi12 >> ((42) & 63)))) | 1)))
                                                            if (!false) {
                                                                var v87 float64 = -2.25
                                                                _ = v87
                                                            }
                                                            _ = v86
                                                        }
                                                        v5 += uint64(v83)
                                                        fmt.Println(v73)
                                                        _ = v83
                                                    } else {
                                                        fmt.Println(v14, v18, v2)
                                                        fmt.Println(v14, v14)
                                                    }
                                                case 1:
                                                    v2 = v2
                                                default:
                                                    v14 += 1
                                                }
                                                _ = v73
                                                _ = v77
                                            }
                                            fmt.Println(v5, fi13, v18)
                                            _ = v72
                                        } else {
                                            var v88 float64 = 1.5
                                            _ = v88
                                        }
                                    }
                                case 923:
                                    v14 = v14
                                default:
                                    var v89 uint8 = 0
                                    _ = v89
                                }
                                v90 := []int{(-(fi12))}
                                _ = v28
                                _ = v90
                            case -245:
                                switch "" {
                                case "":
                                    if true {
                                        if (-1 > 32767) {
                                            var v91 float32 = float32(1.5)
                                            _ = v91
                                        } else {
                                            v92 := GenS6{F0: 681, F1: -537}
                                            v2 = v2
                                            v93 := []int{fi13}
                                            var v94 uint = 0
                                            _ = v92
                                            _ = v93
                                            _ = v94
                                        }
                                        var v95 uint16 = 65535
                                        var v96 uint32 = 418
                                        _ = v95
                                        _ = v96
                                    }
                                case "a":
                                    v97 := GenS6{F0: 42, F1: 491}
                                    _ = v97
                                }
                            }
                            var v98 uint8 = 124
                            _ = v98
                        }
                        fmt.Println(fi12, fi13, fi13)
                        _ = v18
                    } else {
                        if false {
                            switch "foo bar" {
                            case "":
                                v2 -= (v2 * (float32(1.5) + (v2 + float32(((0.1 + -2.25) * -2.25)))))
                            case "a":
                                var v99 string = ""
                                _ = v99
                            default:
                                v4[(fi13 >> ((480) & 63))] = "hello"
                            }
                        } else {
                            var v100 int = (fi13 | 1)
                            if (int32(v14) != 42) {
                                var v101 int16 = 1000
                                v5 -= v5
                                _ = v101
                            }
                            _ = v100
                        }
                    }
                    var v102 string = "foo bar"
                    _ = v102
                    _ = v15
                }
                if ("a" != "") {
                    v103 := GenS5{F0: 0, F1: -302, F2: (fi12 & (fi12 & fi13))}
                    _ = v103
                }
                _ = fi13
                _ = v14
            }
            var v104 uint32 = 42
            _ = v104
        }
        var v105 int8 = -49
        fmt.Println(v2, v105)
        _ = fi12
        _ = v105
    }
    v5 += 1
    var v106 int = 0
    fmt.Println(v106, v5)
    _ = len(v1)
    for fi107 := 0; fi107 < 18; fi107++ {
        var v108 int8 = 100
        _ = fi107
        _ = v108
    }
    switch 644 {
    case 4294967295:
        switch (((((("a" + ("foo bar" + (("foo bar" + "foo bar") + ""))) + "a") + "") + "\n\t") + (("a" + "\n\t") + "\n\t")) + ("a" + "hello")) {
        case "hello":
            _ = len(v4)
        }
    default:
        v109 := GenS2{F0: 600, F1: 622, F2: int32(v106)}
        _ = v109
    }
    if (((-4 < -120) && false) && (!(32767 < -32768))) {
        switch -9223372036854775808 {
        case 1:
            if (!(!true)) {
                if true {
                    if (true || false) {
                        fmt.Println(v2, v106, v2)
                    }
                    v110 := []int{4294967295, (-(v106))}
                    v2 += v2
                    _ = v110
                }
            } else {
                v111 := GenS7{F0: 1, F1: true}
                v112 := GenS3{F0: false, F1: 2147483647, F2: 65535}
                var v113 float64 = 3.14159
                _ = v111
                _ = v112
                _ = v113
            }
        case -1:
            v5 += 18446744073709551615
        }
        switch "foo bar" {
        case "\n\t":
            v114 := GenS4{F0: (v106 / ((v106) | 1)), F1: -128, F2: (false || (830 < -2147483648))}
            _ = v114
        case "foo bar":
            v3[v106] = 171
        case "hello":
            v5 -= 339
        default:
            switch v106 {
            case 595:
                switch "a" {
                case "":
                    v115 := []int{(v106 & v106), v106, v106, v106}
                    _ = v115
                default:
                    fmt.Println(v2, v2, v2)
                }
            case 32768:
                switch -466 {
                case 9223372036854775807:
                    v2 = (3.14159 + (v2 - (float32(0.0) - v2)))
                case -9223372036854775808:
                    v4[v106] = "hello"
                case -288:
                    v116 := GenS5{F0: 241, F1: v106, F2: (-(v106))}
                    _ = v116
                }
            case -99:
                v5 -= 4294967295
            }
        }
    } else {
        for fi117 := 0; fi117 < 18; fi117++ {
            if false {
                v5 += 18446744073709551615
                v118 := GenS1{F0: 1}
                v119 := GenS3{F0: (!(5 <= -128)), F1: (v106 % (((fi117 << ((323) & 63))) | 1)), F2: genFn0()}
                _ = v118
                _ = v119
            }
            _ = fi117
        }
    }
    v106 = v106
    for fi120 := 0; fi120 < 13; fi120++ {
        var v121 bool = false
        v4[fi120] = "a"
        _ = fi120
        _ = v121
    }
    var v122 float32 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v106:", v106)
    fmt.Println("v122:", v122)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v106
    _ = v122
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
