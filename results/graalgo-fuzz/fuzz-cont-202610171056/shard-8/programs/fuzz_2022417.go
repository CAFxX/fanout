package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int }
type GenS2 struct { F0 int8; F1 int64; F2 bool }
type GenS3 struct { F0 int; F1 int32; F2 int32 }
type GenS4 struct { F0 uint; F1 string; F2 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int; F1 uint; F2 int }
type GenS7 struct { F0 string; F1 uint16; F2 uint32 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 string; F1 bool }
type GenS10 struct { F0 string; F1 uint16; F2 int64 }
type GenS11 struct { F0 uint64; F1 int32 }
func genFn0() bool {
    var v6 string = "a"
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 18; fi8++ {
            switch fi7 {
            case 2147483647:
                v6 = v6
            case 0:
                var v9 int = fi8
                _ = v9
            case -2147483648:
                var v10 int16 = -1
                _ = v10
            default:
                fmt.Println(fi8, fi8)
            }
            fmt.Println(fi7, v6)
            switch v6 {
            case "\n\t":
                v11 := GenS1{F0: false, F1: 32767, F2: fi8}
                _ = v11
            case "a":
                if true {
                    v12 := GenS1{F0: (2147483647 != -871), F1: -2147483648, F2: 9223372036854775807}
                    fmt.Println(v6)
                    _ = v12
                }
            case "foo bar":
                var v13 int8 = -44
                _ = v13
            default:
                var v14 uint8 = 255
                _ = v14
            }
            _ = fi8
        }
        v15 := []int{(^(fi7)), (fi7 & (fi7 << ((18446744073709551615) & 63)))}
        v6 = v6
        _ = fi7
        _ = v15
    }
    _ = v6
    return false
}
func genFn1(p0 uint32, p1 uint16, p2 uint32) string {
    var v16 uint = 65535
    if (-748 > 42) {
        for fi17 := 0; fi17 < 17; fi17++ {
            switch ("" + "a") {
            case "hello":
                v18 := []int{fi17, 532}
                _ = v18
            case "foo bar":
                p0 = p0
            case "a":
                for fi19 := 0; fi19 < 1; fi19++ {
                    if (-2.25 == 1.5) {
                        switch (fi17 >> ((v16) & 63)) {
                        case -438:
                            if true {
                                p1 += 2
                                switch "foo bar" {
                                case "hello":
                                    if ((967 >= -1) && true) {
                                        if true {
                                            fmt.Println(p1, p1)
                                            var v20 int16 = 1000
                                            _ = v20
                                        } else {
                                            switch "foo bar" {
                                            case "":
                                                if (int32(p2) != -2147483648) {
                                                    switch "" {
                                                    case "foo bar":
                                                        v21 := []int{}
                                                        _ = v21
                                                    case "":
                                                        v22 := []int{128, fi19, 0, fi19}
                                                        _ = v22
                                                    default:
                                                        p2 += p2
                                                    }
                                                    v16 -= v16
                                                    if (!((-415 == 0) || (!true))) {
                                                        v23 := GenS2{F0: 127, F1: 1, F2: ((p1 >> ((v16) & 63)) == (p1 / ((p1) | 1)))}
                                                        v24 := GenS3{F0: int(p1), F1: 265, F2: 0}
                                                        switch ("" + (("foo bar" + "a") + "\n\t")) {
                                                        case "":
                                                            p0 += (p0 ^ (p2 * p0))
                                                        case "foo bar":
                                                            p2 = p0
                                                        case "a":
                                                            switch 224 {
                                                            case 9223372036854775807:
                                                                v16 = 18446744073709551615
                                                            case 0:
                                                                var v25 int8 = -128
                                                                _ = v25
                                                            default:
                                                                v26 := GenS2{F0: 100, F1: 1, F2: false}
                                                                _ = v26
                                                            }
                                                        default:
                                                            fmt.Println(v16)
                                                        }
                                                        switch fi19 {
                                                        case 499:
                                                            v27 := GenS3{F0: 4294967295, F1: 2147483647, F2: 42}
                                                            _ = v27
                                                        case -1:
                                                            var v28 int64 = -298
                                                            _ = v28
                                                        default:
                                                            p0 += (p0 - (^(p0)))
                                                        }
                                                        _ = v23
                                                        _ = v24
                                                    }
                                                }
                                            case "\n\t":
                                                p1 -= (p1 % ((uint16(p2)) | 1))
                                            case "a":
                                                switch 128 {
                                                case 105:
                                                    if true {
                                                        fmt.Println(p0, v16, p1)
                                                        v29 := GenS4{F0: (^(v16)), F1: (("" + "a") + "a"), F2: -874}
                                                        _ = v29
                                                    } else {
                                                        if false {
                                                            var v30 uint16 = 1
                                                            _ = v30
                                                        }
                                                        v31 := []int{fi17, 802, (fi17 + (^(fi19)))}
                                                        if (-128 > 102) {
                                                            var v32 int = fi19
                                                            p0 = 2
                                                            _ = v32
                                                        }
                                                        _ = v31
                                                    }
                                                case 4294967295:
                                                    v33 := []int{0, (fi17 & (fi17 & fi19)), (^(fi17))}
                                                    _ = v33
                                                case -931:
                                                    switch fi17 {
                                                    case 1:
                                                        v34 := GenS5{F0: (!(true && (!(1 <= int32(fi19)))))}
                                                        _ = v34
                                                    default:
                                                        p2 -= p2
                                                    }
                                                default:
                                                    v35 := GenS6{F0: fi17, F1: (v16 - v16), F2: fi19}
                                                    _ = v35
                                                }
                                            default:
                                                v16 = (^(v16))
                                            }
                                            switch 32767 {
                                            case 0:
                                                v36 := GenS7{F0: "", F1: p1, F2: uint32(v16)}
                                                _ = v36
                                            case -1:
                                                p0 = 91
                                            default:
                                                v37 := []int{9223372036854775807, (fi19 + fi19)}
                                                _ = v37
                                            }
                                            switch fi19 {
                                            case 364:
                                                v38 := []int{1, 374, (fi19 * (-(fi17))), fi17}
                                                _ = v38
                                            default:
                                                p0 -= (^(p2))
                                            }
                                        }
                                    }
                                case "":
                                    if (!false) {
                                        var v39 float32 = (((1.5 - 3.14159) * 1e10) * 0.1)
                                        _ = v39
                                    }
                                case "a":
                                    v40 := []int{(fi17 & -2147483648), fi17, (fi19 * 1)}
                                    _ = v40
                                default:
                                    fmt.Println(p2)
                                }
                            }
                        case 128:
                            if false {
                                var v41 int16 = 42
                                var v42 int64 = 1
                                _ = v41
                                _ = v42
                            }
                            v43 := []int{fi19, 127}
                            p0 -= p2
                            _ = v43
                        case 165:
                            p0 = (p2 * p0)
                        }
                    }
                    fmt.Println(fi17, p1)
                    v44 := GenS2{F0: 1, F1: 0, F2: false}
                    fmt.Println(p0, p1, fi17)
                    v45 := GenS6{F0: (fi19 << ((1) & 63)), F1: v16, F2: fi17}
                    _ = fi19
                    _ = v44
                    _ = v45
                }
            }
            _ = fi17
        }
    } else {
        p0 = p2
        p0 += p2
    }
    p2 -= p0
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return ("\n\t" + "foo bar")
}

func main() {
    var v1 float64 = float64(1.5)
    var v2 float64 = ((float64(0.0) * 1e10) + (v1 + v1))
    v3 := make(map[uint32]bool)
    var v4 float32 = ((1e10 - -2.25) * 1.5)
    var v5 int16 = 1000
    v2 = 0.0
    if true {
        if (v5 == (v5 % ((v5) | 1))) {
            v4 = v4
            var v46 string = "hello"
            if (9223372036854775807 >= 1) {
                _ = len(v3)
                if (895 == 2147483647) {
                    v47 := []int{4294967295, 1, -1}
                    var v48 uint64 = 2
                    var v49 uint32 = 2147483647
                    _ = v47
                    _ = v48
                    _ = v49
                } else {
                    v4 -= (float32(float64(v4)) + float32((0.1 / 1e10)))
                }
                v50 := GenS6{F0: 775, F1: 65535, F2: 127}
                _ = v50
            } else {
                if true {
                    v51 := []int{}
                    _ = v51
                } else {
                    for fi52 := 0; fi52 < 19; fi52++ {
                        if false {
                            var v53 uint8 = 127
                            fmt.Println(v4)
                            _ = v53
                        }
                        for fi54 := 0; fi54 < 11; fi54++ {
                            var v55 int = -497
                            _ = len(v3)
                            if false {
                                switch "" {
                                case "foo bar":
                                    var v56 uint8 = 73
                                    _ = v56
                                case "hello":
                                    fmt.Println(v46)
                                default:
                                    v57 := GenS8{F0: 42}
                                    _ = v57
                                }
                                if false {
                                    var v58 int = fi54
                                    if ((fi52 != v55) || (!(2 > 128))) {
                                        var v59 int64 = -1
                                        v60 := GenS9{F0: "a", F1: (false || (42 != 312))}
                                        _ = v59
                                        _ = v60
                                    }
                                    v61 := GenS1{F0: true, F1: (v55 % ((4294967295) | 1)), F2: 142}
                                    _ = v58
                                    _ = v61
                                } else {
                                    if false {
                                        v62 := []int{fi52, (fi54 + (v55 / ((9223372036854775807) | 1))), v55}
                                        _ = v62
                                    }
                                    var v63 int = fi52
                                    var v64 uint16 = 255
                                    v65 := GenS1{F0: false, F1: v63, F2: 119}
                                    _ = v63
                                    _ = v64
                                    _ = v65
                                }
                                fmt.Println(v2, v4, v46)
                            } else {
                                fmt.Println(v2, v5)
                                if false {
                                    if (true || (-1 < 44)) {
                                        v4 -= (v4 + float32(v2))
                                    }
                                    v5 -= 800
                                    fmt.Println(v5)
                                    var v66 int8 = 0
                                    if true {
                                        v46 = "a"
                                        v55 = fi54
                                        _ = len(v3)
                                        var v67 int = fi52
                                        _ = v67
                                    } else {
                                        var v68 float64 = 0.0
                                        if (106 <= 74) {
                                            var v69 string = ("" + ("" + "hello"))
                                            switch fi52 {
                                            case -2147483648:
                                                if false {
                                                    delete(v3, 2147483648)
                                                    var v70 uint64 = 18446744073709551615
                                                    _ = v70
                                                }
                                            case -224:
                                                var v71 uint16 = 774
                                                _ = v71
                                            case 42:
                                                switch (-(v55)) {
                                                case 0:
                                                    v5 = int16(v55)
                                                case -806:
                                                    var v72 bool = (false && ((fi54 << ((0) & 63)) == v55))
                                                    _ = v72
                                                default:
                                                    var v73 uint32 = 4294967295
                                                    if true {
                                                        var v74 int8 = (-(v66))
                                                        switch v69 {
                                                        case "":
                                                            delete(v3, v73)
                                                        }
                                                        v5 -= v5
                                                        _ = v74
                                                    }
                                                    var v75 int = (fi52 | 2147483647)
                                                    _ = v73
                                                    _ = v75
                                                }
                                            }
                                            _ = v69
                                        } else {
                                            v66 = 42
                                            if true {
                                                switch v46 {
                                                case "foo bar":
                                                    if false {
                                                        v3[65535] = false
                                                        switch fi54 {
                                                        case 0:
                                                            v76 := GenS2{F0: (v66 | (-(v66))), F1: 114, F2: (!(-639 != 2147483647))}
                                                            _ = v76
                                                        case 743:
                                                            if (0 >= 4294967295) {
                                                                v55 = (fi54 / ((v55) | 1))
                                                                var v77 float32 = v4
                                                                switch (v46 + "foo bar") {
                                                                case "a":
                                                                    if false {
                                                                        v78 := []int{(v55 / ((-133) | 1)), fi52, 128, int(v5)}
                                                                        fmt.Println(v77, v1)
                                                                        _ = v78
                                                                    }
                                                                default:
                                                                    if (213 <= 42) {
                                                                        if (v1 > float64((0.1 - v4))) {
                                                                            if true {
                                                                                v1 = float64((float32((v68 * v68)) - v77))
                                                                            }
                                                                        }
                                                                    } else {
                                                                        var v79 int32 = -2147483648
                                                                        v4 -= (1e10 - v4)
                                                                        _ = len(v3)
                                                                        if ((false && true) && ("hello" > (v46 + v46))) {
                                                                            v80 := GenS9{F0: v46, F1: true}
                                                                            v46 = "a"
                                                                            _ = v80
                                                                        } else {
                                                                            fmt.Println(v77, v4, v4)
                                                                        }
                                                                        _ = v79
                                                                    }
                                                                }
                                                                _ = v77
                                                            } else {
                                                                if (!false) {
                                                                    var v81 uint16 = 65535
                                                                    var v82 float64 = v2
                                                                    _ = v81
                                                                    _ = v82
                                                                }
                                                                if (-905 <= (v55 * (fi54 << ((2) & 63)))) {
                                                                    v1 = 0.0
                                                                    var v83 uint16 = 828
                                                                    fmt.Println(v55)
                                                                    _ = v83
                                                                } else {
                                                                    switch ("" + ("foo bar" + (((((v46 + "foo bar") + "hello") + v46) + (("hello" + v46) + v46)) + v46))) {
                                                                    case "a":
                                                                        v55 = 32768
                                                                    }
                                                                    v66 -= v66
                                                                    var v84 int64 = -1
                                                                    _ = v84
                                                                }
                                                            }
                                                        case 42:
                                                            delete(v3, 42)
                                                        }
                                                        var v85 int = 42
                                                        _ = v85
                                                    } else {
                                                        v5 = v5
                                                        var v86 int8 = (v66 ^ v66)
                                                        _ = v86
                                                    }
                                                case "":
                                                    v3[299] = false
                                                case "\n\t":
                                                    fmt.Println(v1, v2)
                                                default:
                                                    fmt.Println(v1, v4, v4)
                                                }
                                                v5 += v5
                                            } else {
                                                if true {
                                                    var v87 int16 = 32767
                                                    v88 := []int{128, (fi54 | (fi54 / ((4294967295) | 1))), fi54, (v55 / ((fi54) | 1))}
                                                    _ = v87
                                                    _ = v88
                                                }
                                            }
                                            v68 -= ((float64((v4 - (1.5 - v4))) * float64(float32(v68))) - float64(1e10))
                                            var v89 bool = (1 < 42)
                                            _ = v89
                                        }
                                        var v90 int64 = -155
                                        var v91 string = v46
                                        _ = v68
                                        _ = v90
                                        _ = v91
                                    }
                                    _ = v66
                                } else {
                                    v3[395] = true
                                }
                            }
                            _ = fi54
                            _ = v55
                        }
                        v3[2147483647] = (v1 == v2)
                        _ = fi52
                    }
                    fmt.Println(v5, v1)
                    v92 := GenS4{F0: 65535, F1: v46, F2: -2147483648}
                    _ = v92
                }
            }
            _ = v46
        } else {
            if false {
                if true {
                    if ((42 == 829) || false) {
                        if true {
                            v3[347] = (false || false)
                            fmt.Println(v4, v4)
                        }
                    } else {
                        var v93 uint16 = 65535
                        v94 := GenS9{F0: ("\n\t" + "\n\t"), F1: (2 < v93)}
                        _ = v93
                        _ = v94
                    }
                    v5 = (-(v5))
                } else {
                    if false {
                        fmt.Println(v4, v1, v2)
                        delete(v3, 65535)
                        v95 := GenS8{F0: -1}
                        _ = v95
                    }
                    switch 2147483647 {
                    case 42:
                        v96 := GenS9{F0: "a", F1: (995 > -262)}
                        _ = v96
                    case -2147483648:
                        v3[42] = (false && (0 != 61))
                    case -311:
                        switch 4294967295 {
                        case 727:
                            v97 := GenS10{F0: "foo bar", F1: 32767, F2: -9223372036854775808}
                            _ = v97
                        case -9223372036854775808:
                            for fi98 := 0; fi98 < 11; fi98++ {
                                var v99 uint16 = 728
                                _ = fi98
                                _ = v99
                            }
                        default:
                            fmt.Println(v5, v2)
                        }
                    }
                }
                fmt.Println(v2, v4, v4)
                v2 = 1e10
            } else {
                v2 = -2.25
                v5 -= -32768
            }
            delete(v3, 737)
        }
        switch 127 {
        case -9223372036854775808:
            fmt.Println(v2)
        case 1:
            if false {
                if false {
                    v100 := []int{127}
                    switch 1 {
                    case -1:
                        v101 := v3[2147483647]
                        _ = v101
                    default:
                        if len(v100) > 0 {
                            v100[0] = -9223372036854775808
                        }
                    }
                    switch ("hello" + "foo bar") {
                    case "a":
                        var v102 uint32 = 42
                        _ = v102
                    case "\n\t":
                        v4 = 0.1
                    case "foo bar":
                        for fi103 := 0; fi103 < 13; fi103++ {
                            var v104 int64 = 9223372036854775807
                            _ = fi103
                            _ = v104
                        }
                    default:
                        v105 := v3[2147483648]
                        _ = v105
                    }
                    _ = v100
                }
                fmt.Println(v5, v5)
                if (v1 != 3.14159) {
                    if (false || (934 > 42)) {
                        fmt.Println(v5, v1)
                        var v106 float32 = v4
                        var v107 uint = 18446744073709551615
                        _ = v106
                        _ = v107
                    }
                    for fi108 := 0; fi108 < 10; fi108++ {
                        if (false && (3.14159 > 0.1)) {
                            v4 -= ((v4 - (0.0 / float32(v1))) + v4)
                        } else {
                            if false {
                                if (42 != 168) {
                                    v109 := GenS8{F0: -1}
                                    for fi110 := 0; fi110 < 1; fi110++ {
                                        v111 := GenS3{F0: 1, F1: 2147483647, F2: -2147483648}
                                        if ((v1 <= float64(3.14159)) || true) {
                                            delete(v3, 42)
                                            v112 := GenS11{F0: 819, F1: 42}
                                            if true {
                                                var v113 uint64 = 18446744073709551615
                                                switch "\n\t" {
                                                case "foo bar":
                                                    if (32767 != 32767) {
                                                        var v114 float32 = (v4 + 3.14159)
                                                        _ = v114
                                                    } else {
                                                        v2 -= 3.14159
                                                    }
                                                default:
                                                    v115 := GenS10{F0: "hello", F1: 0, F2: int64(v5)}
                                                    _ = v115
                                                }
                                                _ = v113
                                            }
                                            _ = v112
                                        } else {
                                            v116 := GenS3{F0: -9223372036854775808, F1: 1, F2: 2147483647}
                                            var v117 uint64 = 1
                                            switch ("hello" + "\n\t") {
                                            case "hello":
                                                var v118 float32 = v4
                                                _ = v118
                                            default:
                                                fmt.Println(v2)
                                            }
                                            _ = v116
                                            _ = v117
                                        }
                                        fmt.Println(fi110, fi110)
                                        _ = fi110
                                        _ = v111
                                    }
                                    v119 := GenS5{F0: ((42 <= 0) || (true && (!false)))}
                                    _ = v109
                                    _ = v119
                                } else {
                                    fmt.Println(v4)
                                    var v120 uint64 = 42
                                    _ = v120
                                }
                                v121 := []int{fi108, (fi108 | -204), fi108, 128}
                                v5 -= (v5 >> ((42) & 63))
                                _ = v121
                            } else {
                                v122 := GenS7{F0: ("\n\t" + ("a" + "foo bar")), F1: 155, F2: 4294967295}
                                _ = v122
                            }
                            for fi123 := 0; fi123 < 17; fi123++ {
                                var v124 uint64 = 661
                                var v125 uint64 = v124
                                v2 -= 1.5
                                _ = fi123
                                _ = v124
                                _ = v125
                            }
                            if (!true) {
                                var v126 float64 = v2
                                _ = v126
                            } else {
                                v127 := v3[42]
                                _ = v127
                            }
                        }
                        v128 := GenS1{F0: (false || true), F1: fi108, F2: fi108}
                        v129 := GenS5{F0: (54 != 2147483647)}
                        _ = fi108
                        _ = v128
                        _ = v129
                    }
                } else {
                    v2 -= v2
                    fmt.Println(v4, v1)
                    var v130 uint = 18446744073709551615
                    _ = v130
                }
            }
        }
        v131 := GenS7{F0: "foo bar", F1: 756, F2: 0}
        _ = v131
    }
    v132 := GenS10{F0: "a", F1: 2, F2: -9223372036854775808}
    if true {
        var v133 int32 = 42
        _ = v133
    }
    _ = len(v3)
    v134 := GenS1{F0: true, F1: 42, F2: 2147483647}
    var v135 bool = (!false)
    v2 -= float64(v4)
    if (!((v5 + (-(v5))) == (v5 & (v5 / (((v5 - (v5 ^ v5))) | 1))))) {
        v136 := GenS1{F0: v135, F1: -1, F2: -294}
        var v137 int8 = 0
        _ = v136
        _ = v137
    }
    v138 := GenS10{F0: ((((((((("" + "a") + "") + "\n\t") + "") + "") + "") + "a") + ("a" + (("\n\t" + "hello") + ""))) + ""), F1: 1, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v135:", v135)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v132
    _ = v134
    _ = v135
    _ = v138
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
