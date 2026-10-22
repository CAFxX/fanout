package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 uint64; F1 uint8; F2 uint }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int; F1 bool; F2 uint32 }
type GenS9 struct { F0 string; F1 uint8; F2 bool }
func genFn0(p0 int16, p1 bool, p2 uint64) int64 {
    v7 := []int{1, 0}
    var v8 uint = 2
    v8 = (v8 + 985)
    var v9 uint8 = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    var v1 uint16 = 209
    v2 := make(map[string]bool)
    var v3 uint = 255
    var v4 float32 = 3.14159
    v5 := make(map[int]string)
    v6 := []uint32{252, 42, 448, 42}
    v5[-9223372036854775808] = (("hello" + "\n\t") + "hello")
    if true {
        for fi10 := 0; fi10 < 8; fi10++ {
            if ((false && (!false)) || false) {
                for fi11 := 0; fi11 < 3; fi11++ {
                    v12 := GenS1{F0: (^(v1))}
                    v13 := GenS1{F0: v1}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                var v14 int32 = 2147483647
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v14
            }
            v4 -= v4
            fmt.Println(v3, v4)
            _ = fi10
        }
        var v15 uint16 = (v1 & v1)
        v16 := GenS1{F0: 65535}
        _ = v15
        _ = v16
    } else {
        var v17 uint32 = 2147483648
        fmt.Println(v17, v17)
        _ = v17
    }
    var v18 float64 = ((1.5 * float64(((v4 - ((float32(3.14159) + float32(0.1)) * -2.25)) + 3.14159))) - 1e10)
    for fi19 := 0; fi19 < 6; fi19++ {
        v20 := GenS1{F0: (v1 >> ((181) & 63))}
        _ = fi19
        _ = v20
    }
    var v21 string = "a"
    v6 = append(v6, 2)
    for fi22 := 0; fi22 < 11; fi22++ {
        _ = len(v2)
        _ = fi22
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        switch -2147483648 {
        case 128:
            for fi24 := 0; fi24 < 13; fi24++ {
                if false {
                    if (!false) {
                        v25 := v2[(v21 + "")]
                        _ = v25
                    }
                    var v26 uint16 = 2
                    v18 = (float64(v4) + 0.0)
                    var v27 int64 = 1
                    _ = v26
                    _ = v27
                } else {
                    v28 := GenS1{F0: v1}
                    var v29 int = -2147483648
                    switch "" {
                    case "foo bar":
                        if ((!false) || (1 != 42)) {
                            v29 = (fi24 - fi24)
                            var v30 uint32 = 0
                            _ = v30
                        }
                    case "\n\t":
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    case "hello":
                        v6 = append(v6, 42)
                    default:
                        v21 = "a"
                    }
                    _ = v28
                    _ = v29
                }
                var v31 int16 = 22
                switch (fi23 >> (((v3 & (v3 & 1))) & 63)) {
                case -38:
                    fmt.Println(v3, v1, v3)
                case 128:
                    v3 = 2
                case 1:
                    v1 -= v1
                }
                var v32 uint32 = 997
                _ = fi24
                _ = v31
                _ = v32
            }
        case 413:
            var v33 int8 = 1
            _ = v33
        case 445:
            fmt.Println(v21)
        default:
            switch v21 {
            case "":
                switch v21 {
                case "hello":
                    var v34 float64 = (-2.25 / v18)
                    _ = v34
                case "\n\t":
                    v35 := GenS2{F0: -128}
                    _ = v35
                case "a":
                    v36 := v2[v21]
                    _ = v36
                }
            case "foo bar":
                switch v21 {
                case "a":
                    if false {
                        if false {
                            var v37 int64 = 0
                            var v38 uint = (v3 | (v3 >> (((v3 % (((v3 | (^(v3)))) | 1))) & 63)))
                            fmt.Println(v18, v1)
                            _ = v37
                            _ = v38
                        } else {
                            var v39 uint8 = uint8(v3)
                            switch ((v21 + "\n\t") + ((v21 + v21) + v21)) {
                            case "\n\t":
                                for fi40 := 0; fi40 < 4; fi40++ {
                                    v3 -= v3
                                    _ = fi40
                                }
                            case "hello":
                                if true {
                                    v21 = "\n\t"
                                    v3 -= (v3 / ((v3) | 1))
                                }
                            case "":
                                var v41 uint64 = 240
                                _ = v41
                            }
                            _ = v39
                        }
                        v21 = "\n\t"
                        var v42 int8 = int8(v1)
                        _ = v42
                    }
                case "":
                    if true {
                        v43 := GenS1{F0: (v1 | v1)}
                        _ = v43
                    }
                default:
                    var v44 bool = false
                    _ = v44
                }
            case "hello":
                v3 = (^(v3))
            }
        }
        _ = fi23
    }
    fmt.Println(v4, v4, v18)
    v4 = v4
    if false {
        _ = len(v6)
    }
    if true {
        v45 := GenS3{F0: 2, F1: -1, F2: 2}
        v1 = 414
        _ = v45
    } else {
        var v46 int32 = -2147483648
        for fi47 := 0; fi47 < 12; fi47++ {
            var v48 uint64 = 962
            v5[fi47] = v21
            _ = fi47
            _ = v48
        }
        _ = v46
    }
    v49 := GenS2{F0: 1}
    for fi50 := 0; fi50 < 17; fi50++ {
        var v51 uint64 = 0
        _ = fi50
        _ = v51
    }
    v52 := GenS1{F0: (v1 << (((v3 | (v3 ^ (v3 - v3)))) & 63))}
    for fi53 := 0; fi53 < 19; fi53++ {
        fmt.Println(v1, v4, v1)
        for fi54 := 0; fi54 < 16; fi54++ {
            var v55 bool = (v1 < v1)
            if v55 {
                v55 = v55
                if v55 {
                    if false {
                        switch "a" {
                        case "a":
                            v3 = (v3 / (((v3 + (v3 >> ((2) & 63)))) | 1))
                        case "":
                            switch -664 {
                            case -2147483648:
                                switch "a" {
                                case "\n\t":
                                    var v56 int = (^(fi54))
                                    _ = v56
                                case "a":
                                    v18 = (((float64(v4) + (v18 + 0.0)) - (v18 + v18)) - 3.14159)
                                case "foo bar":
                                    switch (-(fi53)) {
                                    case 398:
                                        v18 += v18
                                    case 0:
                                        var v57 uint = (v3 | 542)
                                        _ = v57
                                    case 32768:
                                        v58 := GenS4{F0: 611, F1: (!false)}
                                        switch fi53 {
                                        case 9223372036854775807:
                                            switch fi54 {
                                            case 1:
                                                v59 := GenS3{F0: 491, F1: 42, F2: 127}
                                                _ = v59
                                            case 726:
                                                _ = len(v6)
                                            case 9223372036854775807:
                                                v4 -= float32(v18)
                                            }
                                        case 714:
                                            v60 := GenS5{F0: 433, F1: 128, F2: (v3 >> (((v3 & v3)) & 63))}
                                            _ = v60
                                        case 2147483647:
                                            switch v21 {
                                            case "":
                                                v21 = v21
                                            case "a":
                                                var v61 int16 = 32767
                                                _ = v61
                                            case "hello":
                                                var v62 int16 = 999
                                                _ = v62
                                            default:
                                                var v63 uint32 = 0
                                                _ = v63
                                            }
                                        }
                                        var v64 int32 = -2147483648
                                        v65 := GenS2{F0: 0}
                                        v66 := v2["foo bar"]
                                        v67 := GenS6{F0: 29}
                                        _ = v58
                                        _ = v64
                                        _ = v65
                                        _ = v66
                                        _ = v67
                                    }
                                default:
                                    if (341 >= 1) {
                                        v68 := GenS3{F0: (v1 / ((255) | 1)), F1: -2147483648, F2: 1}
                                        fmt.Println(v18)
                                        _ = v68
                                    }
                                    v69 := GenS7{F0: v21}
                                    v70 := GenS5{F0: 0, F1: 255, F2: 4294967295}
                                    _ = v69
                                    _ = v70
                                }
                            case -1:
                                var v71 float64 = v18
                                _ = v71
                            case 1:
                                v18 -= v18
                            default:
                                var v72 uint = (v3 / ((v3) | 1))
                                if v55 {
                                    v73 := GenS3{F0: v1, F1: 2147483647, F2: 23}
                                    v1 += (v1 + v1)
                                    switch ("hello" + v21) {
                                    case "a":
                                        var v74 int = (-(fi53))
                                        _ = v74
                                    }
                                    _ = v73
                                } else {
                                    fmt.Println(v55, v21)
                                    v5[fi53] = ""
                                    var v75 int8 = -90
                                    if v55 {
                                        v76 := GenS5{F0: 4294967295, F1: 0, F2: (v72 | 812)}
                                        var v77 float32 = v4
                                        v78 := GenS5{F0: 4294967295, F1: 87, F2: 110}
                                        _ = v76
                                        _ = v77
                                        _ = v78
                                    }
                                    v79 := GenS2{F0: v75}
                                    _ = v75
                                    _ = v79
                                }
                                _ = v72
                            }
                        }
                        v6 = append(v6, 65535)
                        v80 := GenS6{F0: 681}
                        _ = v80
                    } else {
                        delete(v2, ("\n\t" + v21))
                    }
                    var v81 uint = uint(v1)
                    fmt.Println(v1, v81, v1)
                    _ = v81
                } else {
                    if v55 {
                        v82 := GenS8{F0: 42, F1: v55, F2: 2147483648}
                        fmt.Println(v18, v1, v55)
                        var v83 int64 = genFn0(1, false, 1)
                        _ = v82
                        _ = v83
                    } else {
                        delete(v2, ("\n\t" + v21))
                        if (42 == 4294967295) {
                            var v84 int64 = 42
                            switch "\n\t" {
                            case "foo bar":
                                if ((((!false) || (v3 >= (^(v3)))) || v55) || ((v1 > (v1 ^ (^(v1)))) && (v3 <= 4294967295))) {
                                    v84 = (v84 % ((7) | 1))
                                    var v85 uint = 42
                                    switch ("a" + v21) {
                                    case "foo bar":
                                        switch "\n\t" {
                                        case "foo bar":
                                            fmt.Println(v84, v1, v1)
                                        }
                                    case "a":
                                        v1 -= 1
                                    case "hello":
                                        v1 += v1
                                    }
                                    _ = v85
                                } else {
                                    var v86 uint16 = (^(v1))
                                    v21 = v21
                                    v86 = 42
                                    _ = v86
                                }
                            case "\n\t":
                                fmt.Println(fi54, v55, v4)
                            default:
                                switch "\n\t" {
                                case "a":
                                    v87 := GenS1{F0: v1}
                                    _ = v87
                                case "hello":
                                    v88 := GenS9{F0: "foo bar", F1: 42, F2: true}
                                    var v89 int = (-(fi53))
                                    _ = v88
                                    _ = v89
                                case "":
                                    var v90 float64 = v18
                                    _ = v90
                                }
                            }
                            var v91 int32 = 2147483647
                            _ = v84
                            _ = v91
                        } else {
                            delete(v5, (fi54 + 0))
                            var v92 uint8 = 255
                            var v93 uint32 = 438
                            _ = v92
                            _ = v93
                        }
                    }
                }
            } else {
                var v94 string = ("\n\t" + ("foo bar" + (v21 + ("foo bar" + "a"))))
                if len(v6) > 0 {
                    v6[0] = 210
                }
                _ = v94
            }
            switch "hello" {
            case "\n\t":
                switch ("" + "hello") {
                case "\n\t":
                    switch 910 {
                    case 865:
                        var v95 uint8 = 1
                        _ = v95
                    case 1:
                        fmt.Println(v4, v55)
                    default:
                        v55 = v55
                    }
                case "foo bar":
                    switch "foo bar" {
                    case "\n\t":
                        v1 += 1
                    case "hello":
                        var v96 float64 = (0.0 * ((v18 * v18) + (1.5 / (((-2.25 + 0.0) * 1e10) / 3.14159))))
                        _ = v96
                    case "":
                        var v97 uint = v3
                        _ = v97
                    default:
                        if v55 {
                            v4 = 0.0
                            if ((v3 * 255) == v3) {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                v4 = -2.25
                                _ = len(v6)
                            } else {
                                v18 += v18
                                v1 = 1
                            }
                        }
                    }
                case "hello":
                    if v55 {
                        switch (v21 + v21) {
                        case "foo bar":
                            fmt.Println(fi54, v3, v4)
                        case "":
                            switch v21 {
                            case "a":
                                var v98 string = ((v21 + v21) + v21)
                                _ = v98
                            case "foo bar":
                                v99 := v5[fi53]
                                _ = v99
                            case "hello":
                                var v100 string = "hello"
                                _ = v100
                            }
                        case "hello":
                            if ((v3 / ((751) | 1)) <= v3) {
                                var v101 uint32 = 357
                                _ = v101
                            }
                        default:
                            _ = len(v6)
                        }
                        delete(v2, ((v21 + ("\n\t" + "hello")) + (v21 + (v21 + v21))))
                        var v102 int8 = 1
                        v1 -= 609
                        v55 = v55
                        _ = v102
                    } else {
                        if false {
                            v3 = v3
                            v3 = (^(v3))
                            fmt.Println(fi54)
                            v2[v21] = (1 < -32768)
                        }
                    }
                    v21 = (v21 + (v21 + (("hello" + ((v21 + "a") + ("foo bar" + v21))) + (("\n\t" + "a") + "hello"))))
                    var v103 uint64 = 18446744073709551615
                    _ = v103
                }
            case "hello":
                switch -198 {
                case 42:
                    fmt.Println(v18)
                case 280:
                    v55 = (127 == 100)
                case 4294967295:
                    var v104 int64 = genFn0(-1, (42 <= 100), 4294967295)
                    _ = v104
                }
            case "foo bar":
                if v55 {
                    var v105 string = v21
                    v3 -= (v3 - v3)
                    _ = v105
                } else {
                    var v106 float64 = -2.25
                    if ((v1 ^ v1) != v1) {
                        if v55 {
                            _ = len(v6)
                            v4 += v4
                            switch fi53 {
                            case 228:
                                switch fi54 {
                                case 487:
                                    v107 := GenS8{F0: 32767, F1: v55, F2: 2147483648}
                                    _ = v107
                                case -463:
                                    if len(v6) > 0 {
                                        v6[0] = 661
                                    }
                                case -2147483648:
                                    if ((v1 | v1) == (v1 << (((v3 / (((v3 + (v3 - (v3 << (((^(v3))) & 63))))) | 1))) & 63))) {
                                        delete(v5, 158)
                                        v108 := GenS4{F0: 32767, F1: v55}
                                        delete(v5, (fi53 & 42))
                                        _ = v108
                                    } else {
                                        v18 += v18
                                        if ((42 != 0) && v55) {
                                            var v109 uint8 = 1
                                            if ((v3 << (((v3 & 647)) & 63)) <= (^(v3))) {
                                                v6 = append(v6, 1)
                                            }
                                            fmt.Println(v3, fi53, v4)
                                            _ = v109
                                        }
                                        if v55 {
                                            var v110 float64 = 1e10
                                            _ = v110
                                        } else {
                                            var v111 int = -354
                                            _ = len(v2)
                                            _ = v111
                                        }
                                    }
                                default:
                                    var v112 int = fi53
                                    switch v112 {
                                    case 284:
                                        v106 += (v18 - (3.14159 - 1e10))
                                    default:
                                        var v113 bool = (42 < 42)
                                        var v114 int8 = 93
                                        _ = v113
                                        _ = v114
                                    }
                                    v106 += v106
                                    var v115 int8 = 45
                                    var v116 bool = false
                                    _ = v112
                                    _ = v115
                                    _ = v116
                                }
                            case -661:
                                if (((-2.25 * v4) < (v4 / 0.1)) || (-32768 >= 1)) {
                                    var v117 int = (fi53 ^ fi54)
                                    var v118 int16 = 1
                                    v119 := GenS9{F0: "hello", F1: 42, F2: (v55 || true)}
                                    v5[-9223372036854775808] = v21
                                    _ = v117
                                    _ = v118
                                    _ = v119
                                } else {
                                    switch (fi53 + (fi53 & (fi53 - 1))) {
                                    case 128:
                                        var v120 int16 = -32768
                                        _ = v120
                                    }
                                    v4 += (0.0 * ((v4 / v4) * v4))
                                    if (0 > 1) {
                                        var v121 uint64 = 4294967295
                                        v4 += -2.25
                                        v121 = v121
                                        _ = v121
                                    } else {
                                        v21 = ""
                                    }
                                    v18 = 0.1
                                }
                            default:
                                _ = len(v6)
                            }
                        }
                    }
                    v4 = (0.1 + (-2.25 / v4))
                    _ = v106
                }
            default:
                v122 := GenS8{F0: (^(fi54)), F1: (v55 && (v55 && false)), F2: 515}
                _ = v122
            }
            _ = fi54
            _ = v55
        }
        _ = fi53
    }
    fmt.Println(v18, v4)
    fmt.Println(v21, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v52
    _ = v6
}
