package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 uint64; F1 int8 }
type GenS5 struct { F0 int; F1 uint32; F2 string }
type GenS6 struct { F0 uint32; F1 string; F2 uint64 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int64; F1 uint16 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint; F1 int16 }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 uint32 }
func genFn0(p0 int32, p1 uint16, p2 uint32) uint8 {
    var v4 int = -282
    var v5 uint = 1
    v6 := GenS1{F0: (p1 << ((v5) & 63)), F1: v5}
    if true {
        for fi7 := 0; fi7 < 19; fi7++ {
            if ((!(774 != (p1 ^ p1))) || true) {
                v8 := GenS1{F0: p1, F1: uint(p1)}
                _ = v8
            } else {
                v9 := []int{(fi7 | fi7), (v4 << ((74) & 63)), fi7}
                _ = v9
            }
            v5 = (v5 ^ (v5 << ((255) & 63)))
            v10 := GenS1{F0: (p1 ^ (p1 >> (((v5 ^ v5)) & 63))), F1: (^(v5))}
            _ = fi7
            _ = v10
        }
        if false {
            if ((47 != 124) && (v5 > 258)) {
                var v11 uint64 = 42
                _ = v11
            }
            v12 := []int{(v4 + v4), (v4 ^ (v4 << ((449) & 63))), (v4 & v4), (v4 & v4)}
            _ = v12
        }
    } else {
        fmt.Println(p0, v5, v4)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    v13 := GenS2{F0: 255}
    switch 127 {
    case 815:
        var v14 bool = (false || (!true))
        _ = v14
    case -2147483648:
        v15 := []int{-332, 0, 42}
        _ = v15
    case 42:
        v16 := []int{32768, 0, 4294967295, 4294967295}
        _ = v16
    default:
        for fi17 := 0; fi17 < 17; fi17++ {
            if false {
                v18 := []int{(fi17 - 324), (fi17 % ((fi17) | 1)), fi17}
                _ = v18
            } else {
                var v19 int32 = 42
                _ = v19
            }
            fmt.Println(fi17)
            v20 := GenS3{F0: "a", F1: 564}
            _ = fi17
            _ = v20
        }
    }
    delete(v1, 0)
    var v21 int16 = -657
    fmt.Println(v21)
    v22 := GenS2{F0: 0}
    if false {
        if false {
            v23 := []int{}
            var v24 uint = 995
            _ = v23
            _ = v24
        } else {
            v25 := GenS4{F0: 18446744073709551615, F1: 1}
            if true {
                for fi26 := 0; fi26 < 2; fi26++ {
                    var v27 int16 = v21
                    switch (("hello" + "a") + "\n\t") {
                    case "hello":
                        var v28 string = "hello"
                        _ = v28
                    case "\n\t":
                        v21 = (v27 % ((638) | 1))
                    case "a":
                        v29 := GenS4{F0: 4294967295, F1: 109}
                        _ = v29
                    default:
                        if false {
                            switch (fi26 / (((fi26 % ((1) | 1))) | 1)) {
                            case -170:
                                for fi30 := 0; fi30 < 1; fi30++ {
                                    var v31 int64 = -9223372036854775808
                                    _ = fi30
                                    _ = v31
                                }
                            default:
                                for fi32 := 0; fi32 < 18; fi32++ {
                                    v33 := GenS2{F0: 2}
                                    switch ("a" + "foo bar") {
                                    case "foo bar":
                                        switch "hello" {
                                        case "a":
                                            switch "\n\t" {
                                            case "":
                                                v34 := []int{fi32, (fi32 | fi32)}
                                                _ = v34
                                            case "\n\t":
                                                var v35 uint64 = 4294967295
                                                _ = v35
                                            default:
                                                if (!(((("" + ("hello" + (("" + ("a" + ("hello" + "foo bar"))) + "\n\t"))) + "foo bar") + (("" + ("foo bar" + "\n\t")) + "hello")) > "hello")) {
                                                    if true {
                                                        switch "hello" {
                                                        case "\n\t":
                                                            v21 -= (v27 ^ 336)
                                                        case "a":
                                                            v21 = v27
                                                        case "hello":
                                                            v27 = (v27 | v21)
                                                        default:
                                                            if true {
                                                                var v36 uint64 = 2
                                                                var v37 int16 = (^(v21))
                                                                _ = v36
                                                                _ = v37
                                                            }
                                                        }
                                                        v21 = v21
                                                        if ((4294967295 >= 4294967295) || true) {
                                                            var v38 float64 = 0.0
                                                            var v39 bool = (float32(1.5) < (float32(v38) * 1e10))
                                                            v39 = (42 != 1)
                                                            _ = v38
                                                            _ = v39
                                                        } else {
                                                            var v40 uint32 = 2
                                                            fmt.Println(v21)
                                                            _ = v40
                                                        }
                                                    }
                                                } else {
                                                    v21 -= (v27 % (((-(v27))) | 1))
                                                }
                                            }
                                            if false {
                                                v27 = -1
                                                if (true && (!(42 <= 290))) {
                                                    var v41 int16 = v21
                                                    switch fi32 {
                                                    case 32768:
                                                        var v42 uint = 0
                                                        v43 := GenS3{F0: "", F1: -3}
                                                        if true {
                                                            fmt.Println(v42)
                                                            var v44 uint = v42
                                                            v44 += 4294967295
                                                            _ = v44
                                                        }
                                                        var v45 int64 = -9223372036854775808
                                                        _ = v42
                                                        _ = v43
                                                        _ = v45
                                                    default:
                                                        var v46 float64 = float64((-2.25 / 1e10))
                                                        v21 -= v21
                                                        v47 := v1[455]
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                    _ = v41
                                                }
                                            } else {
                                                v21 += v27
                                            }
                                        }
                                    case "hello":
                                        var v48 int8 = -1
                                        _ = v48
                                    default:
                                        v49 := GenS2{F0: 255}
                                        v27 -= v27
                                        if false {
                                            var v50 uint64 = 18446744073709551615
                                            v51 := v3[fi32]
                                            var v52 bool = ("" < ((("a" + "\n\t") + "hello") + v51))
                                            v53 := []int{32767, (fi32 & fi26)}
                                            _ = v50
                                            _ = v51
                                            _ = v52
                                            _ = v53
                                        }
                                        _ = v49
                                    }
                                    _ = fi32
                                    _ = v33
                                }
                            }
                            fmt.Println(fi26, fi26)
                            var v54 bool = (!false)
                            _ = v54
                        }
                    }
                    if true {
                        fmt.Println(fi26, v21, v27)
                    }
                    _ = fi26
                    _ = v27
                }
                v55 := GenS3{F0: ("hello" + "hello"), F1: 733}
                for fi56 := 0; fi56 < 16; fi56++ {
                    if (42 == 2) {
                        v21 = v21
                        if (1 <= 537) {
                            for fi57 := 0; fi57 < 6; fi57++ {
                                fmt.Println(v21, fi56, fi56)
                                if true {
                                    if (127 == 0) {
                                        var v58 int8 = 1
                                        _ = v58
                                    } else {
                                        fmt.Println(fi56, fi56, fi57)
                                        if true {
                                            if (!((fi56 == 730) && true)) {
                                                if false {
                                                    _ = len(v1)
                                                } else {
                                                    if true {
                                                        var v59 float32 = 1e10
                                                        _ = v59
                                                    }
                                                }
                                                if false {
                                                    fmt.Println(v21, fi56, fi57)
                                                    var v60 int16 = -536
                                                    fmt.Println(v21, fi57, v60)
                                                    v60 += (v21 | (v21 >> ((0) & 63)))
                                                    v21 -= v60
                                                    var v61 uint = 2
                                                    var v62 int = (fi57 >> ((v61) & 63))
                                                    _ = v60
                                                    _ = v61
                                                    _ = v62
                                                } else {
                                                    var v63 int64 = 0
                                                    v63 = (-(v63))
                                                    v64 := GenS3{F0: "\n\t", F1: v63}
                                                    _ = v63
                                                    _ = v64
                                                }
                                                var v65 int8 = -1
                                                _ = v65
                                            } else {
                                                var v66 int8 = -53
                                                v21 = v21
                                                _ = v66
                                            }
                                            if true {
                                                var v67 uint8 = 124
                                                var v68 uint64 = 18446744073709551615
                                                v68 = (v68 << ((255) & 63))
                                                _ = v67
                                                _ = v68
                                            } else {
                                                var v69 int8 = 70
                                                v70 := GenS5{F0: (fi56 | (fi57 & (fi57 - fi56))), F1: 1, F2: "foo bar"}
                                                if (-518 <= fi56) {
                                                    var v71 uint16 = 2
                                                    _ = v71
                                                }
                                                _ = v69
                                                _ = v70
                                            }
                                        } else {
                                            var v72 uint8 = 42
                                            v73 := []int{(fi56 << ((4294967295) & 63)), 9223372036854775807, 32768, (fi56 | 949)}
                                            if false {
                                                var v74 uint32 = 1
                                                _ = v74
                                            }
                                            _ = v72
                                            _ = v73
                                        }
                                    }
                                    _ = len(v2)
                                } else {
                                    var v75 int32 = 1
                                    var v76 uint32 = 1
                                    v77 := []int{}
                                    _ = v75
                                    _ = v76
                                    _ = v77
                                }
                                var v78 uint = 1
                                _ = fi57
                                _ = v78
                            }
                            var v79 uint8 = 106
                            if false {
                                v21 -= v21
                                if (128 <= (fi56 * 930)) {
                                    v79 = (^(v79))
                                    var v80 int8 = -93
                                    v81 := []int{fi56, (fi56 % (((fi56 & (fi56 + (fi56 | fi56)))) | 1)), (fi56 + 729)}
                                    _ = v80
                                    _ = v81
                                }
                            } else {
                                v21 = v21
                            }
                            _ = v79
                        }
                    }
                    switch fi56 {
                    case 970:
                        v21 = 32767
                    case 266:
                        for fi82 := 0; fi82 < 8; fi82++ {
                            fmt.Println(fi56)
                            v83 := []int{(fi56 * fi56)}
                            switch "\n\t" {
                            case "hello":
                                var v84 uint32 = 42
                                _ = v84
                            case "foo bar":
                                v85 := GenS2{F0: 1}
                                _ = v85
                            case "\n\t":
                                var v86 int = fi56
                                v21 += (v21 % (((-(v21))) | 1))
                                if false {
                                    if true {
                                        v86 -= (-(fi56))
                                        if false {
                                            v21 = (^(v21))
                                            fmt.Println(v21, v86, fi82)
                                            if (65535 == 771) {
                                                if true {
                                                    v86 -= 1
                                                    v87 := GenS6{F0: 2, F1: ((("\n\t" + "\n\t") + "hello") + "hello"), F2: 42}
                                                    var v88 uint8 = 2
                                                    _ = v87
                                                    _ = v88
                                                }
                                                if (((^(v86)) >= -764) || true) {
                                                    if len(v83) > 0 {
                                                        v83[0] = v86
                                                    }
                                                    if true {
                                                        if false {
                                                            fmt.Println(v21)
                                                            if true {
                                                                fmt.Println(fi82, fi56)
                                                                var v89 int64 = 9223372036854775807
                                                                _ = v89
                                                            }
                                                        } else {
                                                            if true {
                                                                var v90 string = "foo bar"
                                                                var v91 float64 = -2.25
                                                                _ = v90
                                                                _ = v91
                                                            } else {
                                                                fmt.Println(v86)
                                                                var v92 int8 = 1
                                                                _ = v92
                                                            }
                                                        }
                                                        var v93 int = -2147483648
                                                        var v94 float64 = 1e10
                                                        v83 = append(v83, -9223372036854775808)
                                                        var v95 int = -213
                                                        _ = v93
                                                        _ = v94
                                                        _ = v95
                                                    } else {
                                                        v86 = fi82
                                                    }
                                                } else {
                                                    v86 += fi82
                                                    if (((1e10 + ((((0.0 * (((1.5 - ((0.1 / 1.0) + 0.0)) - 1.5) / 1.0)) - (3.14159 * float32(float64(((1e10 / 1e10) + ((float32(0.1) * 1.5) + -2.25)))))) + 0.1) - -2.25)) - float32((0.1 / ((float64(1e10) * 1e10) / 1.0)))) > ((0.1 - 1.5) * 0.1)) {
                                                        fmt.Println(v21, v86, fi56)
                                                    }
                                                }
                                                if false {
                                                    v21 -= v21
                                                } else {
                                                    if true {
                                                        if true {
                                                            v21 -= (v21 + v21)
                                                            var v96 uint16 = 1
                                                            var v97 uint16 = (v96 / ((32767) | 1))
                                                            _ = v96
                                                            _ = v97
                                                        } else {
                                                            v98 := GenS7{F0: 393}
                                                            v99 := GenS5{F0: fi82, F1: 666, F2: ("hello" + "\n\t")}
                                                            var v100 float64 = (float64(0.1) - -2.25)
                                                            _ = v100
                                                            _ = v98
                                                            _ = v99
                                                        }
                                                        v101 := GenS6{F0: 2, F1: "foo bar", F2: 1}
                                                        _ = v101
                                                    } else {
                                                        switch fi82 {
                                                        case -9223372036854775808:
                                                            switch "\n\t" {
                                                            case "":
                                                                v102 := GenS3{F0: "", F1: -374}
                                                                _ = v102
                                                            default:
                                                                v21 = (v21 * 1)
                                                            }
                                                        case 4294967295:
                                                            if (!(!false)) {
                                                                var v103 uint8 = 237
                                                                fmt.Println(fi56, v21)
                                                                fmt.Println(fi82, v86, fi56)
                                                                _ = v103
                                                            } else {
                                                                switch (^(v86)) {
                                                                case 127:
                                                                    v104 := v3[fi56]
                                                                    _ = v104
                                                                case 394:
                                                                    if len(v83) > 0 {
                                                                        v83[0] = (-(v86))
                                                                    }
                                                                case -9223372036854775808:
                                                                    var v105 uint32 = 357
                                                                    _ = v105
                                                                }
                                                                if (!false) {
                                                                    var v106 uint16 = 32767
                                                                    if false {
                                                                        fmt.Println(fi56, fi56)
                                                                        v21 += 1000
                                                                        var v107 float32 = 3.14159
                                                                        _ = v107
                                                                    } else {
                                                                        var v108 int16 = v21
                                                                        _ = v108
                                                                    }
                                                                    switch fi82 {
                                                                    case 1:
                                                                        var v109 string = "foo bar"
                                                                        _ = v109
                                                                    case 127:
                                                                        if len(v83) >= 2 {
                                                                            v83 = v83[1:len(v83)]
                                                                        }
                                                                    default:
                                                                        v106 = v106
                                                                    }
                                                                    _ = v106
                                                                }
                                                            }
                                                        }
                                                    }
                                                    if (!false) {
                                                        var v110 int = v86
                                                        _ = v110
                                                    }
                                                    v111 := GenS3{F0: ("a" + ("foo bar" + ("" + (("" + "") + ("" + ""))))), F1: 42}
                                                    _ = v111
                                                }
                                                switch "\n\t" {
                                                case "\n\t":
                                                    fmt.Println(v21, v21)
                                                    if (false && false) {
                                                        var v112 uint = 255
                                                        _ = v112
                                                    }
                                                case "":
                                                    var v113 int = (v86 / (((fi82 >> ((755) & 63))) | 1))
                                                    _ = v113
                                                default:
                                                    v114 := GenS8{F0: -9223372036854775808, F1: 42}
                                                    _ = v114
                                                }
                                                v21 = (v21 * 32767)
                                                var v115 bool = (genFn0(-592, 0, 2) >= genFn0(1, 2, 2))
                                                _ = v115
                                            }
                                        }
                                    } else {
                                        if (255 <= 618) {
                                            v21 += v21
                                        }
                                        switch ("\n\t" + "") {
                                        case "a":
                                            v116 := GenS9{F0: 42}
                                            _ = v116
                                        case "foo bar":
                                            var v117 float32 = 1.5
                                            _ = v117
                                        case "hello":
                                            if (!(float64(-2.25) == 0.1)) {
                                                v118 := GenS6{F0: 818, F1: "foo bar", F2: 4294967295}
                                                fmt.Println(fi82)
                                                _ = v118
                                            } else {
                                                v21 = (v21 & 1)
                                            }
                                        }
                                        v21 -= -1
                                    }
                                    switch (fi82 + (v86 ^ fi56)) {
                                    case 388:
                                        switch ("" + (("a" + "foo bar") + "a")) {
                                        case "a":
                                            v21 = (v21 | v21)
                                        }
                                    case 32768:
                                        if (2 <= 0) {
                                            v21 += (v21 << ((1) & 63))
                                            if len(v83) >= 2 {
                                                v83 = v83[1:len(v83)]
                                            }
                                        } else {
                                            v86 = (fi82 + (v86 | 32767))
                                        }
                                    }
                                    v21 = v21
                                    fmt.Println(fi82, v21, fi82)
                                    v119 := GenS10{F0: 2147483647}
                                    if true {
                                        switch "\n\t" {
                                        case "":
                                            v86 += (fi56 | v86)
                                        case "foo bar":
                                            v120 := GenS11{F0: 65535, F1: v21}
                                            _ = v120
                                        case "a":
                                            v86 = 4294967295
                                        }
                                        v121 := GenS12{F0: 2147483648}
                                        _ = v121
                                    } else {
                                        if (false && false) {
                                            switch (fi56 >> ((1) & 63)) {
                                            case -678:
                                                if len(v83) > 0 {
                                                    v83[0] = (fi82 % ((fi56) | 1))
                                                }
                                            case 4294967295:
                                                fmt.Println(v21, v86)
                                            }
                                            var v122 uint = 318
                                            v86 += v86
                                            _ = v122
                                        }
                                        var v123 int8 = 1
                                        var v124 int64 = 42
                                        v125 := GenS10{F0: 65535}
                                        var v126 int32 = 0
                                        _ = v123
                                        _ = v124
                                        _ = v125
                                        _ = v126
                                    }
                                    _ = v119
                                }
                                _ = v86
                            }
                            _ = fi82
                            _ = v83
                        }
                    }
                    v127 := GenS13{F0: -128}
                    _ = fi56
                    _ = v127
                }
                _ = v55
            } else {
                v128 := []int{9223372036854775807}
                _ = v128
            }
            if (4294967295 == 1) {
                if (2147483648 <= 1) {
                    fmt.Println(v21)
                } else {
                    var v129 bool = false
                    _ = v129
                }
                fmt.Println(v21)
                fmt.Println(v21)
            } else {
                var v130 int = -188
                v21 += -1
                _ = v130
            }
            _ = v25
        }
    } else {
        v131 := []int{}
        v21 = v21
        _ = v131
    }
    v132 := GenS14{F0: 42}
    fmt.Println(v21)
    v21 = -32768
    var v133 float64 = (1e10 - ((((3.14159 - 3.14159) / 3.14159) / 1.0) / (1.5 / 1.5)))
    v134 := GenS14{F0: 312}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v133:", v133)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v132
    _ = v133
    _ = v134
    _ = v2
    _ = v21
    _ = v22
    _ = v3
}
