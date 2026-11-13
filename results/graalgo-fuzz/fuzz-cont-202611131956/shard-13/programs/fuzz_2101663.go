package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int64 }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 string; F1 bool; F2 uint8 }
type GenS7 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int; F1 int8 }
type GenS10 struct { F0 int; F1 uint64; F2 uint }
type GenS11 struct { F0 uint64; F1 int; F2 uint32 }
type GenS12 struct { F0 uint8; F1 int8; F2 int8 }
type GenS13 struct { F0 int64 }
func genFn0(p0 uint16, p1 uint16, p2 bool) bool {
    v4 := GenS1{F0: -9223372036854775808, F1: 46, F2: 427}
    var v5 string = "hello"
    v5 = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return false
}
func genFn1(p0 int, p1 uint64, p2 int) uint64 {
    v6 := GenS1{F0: 393, F1: 1, F2: 9223372036854775807}
    p0 -= (^(p0))
    var v7 int16 = 0
    for fi8 := 0; fi8 < 12; fi8++ {
        v7 = (-(v7))
        v9 := GenS2{F0: (v7 - (v7 % ((v7) | 1))), F1: 21}
        var v10 uint64 = (p1 & p1)
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p1
}

func main() {
    v1 := []uint32{}
    v2 := []int{183}
    v3 := []string{"foo bar"}
    v11 := GenS1{F0: -9223372036854775808, F1: 1, F2: -1}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v12 uint = 18446744073709551615
    v13 := GenS3{F0: 65535, F1: ""}
    fmt.Println(v12)
    switch -426 {
    case 996:
        for fi14 := 0; fi14 < 10; fi14++ {
            for fi15 := 0; fi15 < 15; fi15++ {
                v12 += (v12 / ((v12) | 1))
                if false {
                    var v16 float64 = 1e10
                    _ = v16
                }
                v17 := GenS4{F0: 23}
                _ = fi15
                _ = v17
            }
            fmt.Println(fi14)
            if len(v1) > 0 {
                v1[0] = 65535
            }
            _ = fi14
        }
    case 4294967295:
        var v18 int64 = 1
        _ = v18
    default:
        switch ("\n\t" + ("foo bar" + ("a" + (("" + ("a" + (("foo bar" + "\n\t") + ("" + "a")))) + ("" + "foo bar"))))) {
        case "":
            v12 = v12
        case "foo bar":
            v12 = 255
        default:
            v12 = v12
        }
    }
    var v19 int8 = 127
    v12 -= (v12 ^ v12)
    switch ("foo bar" + "hello") {
    case "hello":
        switch 32767 {
        case 127:
            if true {
                for fi20 := 0; fi20 < 15; fi20++ {
                    fmt.Println(v19, fi20, v12)
                    var v21 float32 = (float32((1.5 - (0.0 - 1.5))) / float32(float64((0.0 + -2.25))))
                    _ = fi20
                    _ = v21
                }
            } else {
                if ((("foo bar" + ("\n\t" + ("a" + "hello"))) + ("" + ("a" + ("a" + ("\n\t" + "a"))))) >= (("\n\t" + "hello") + ("foo bar" + ""))) {
                    switch -2147483648 {
                    case 1:
                        v12 -= (^(v12))
                    case 32767:
                        var v22 int = 127
                        _ = v22
                    case 32768:
                        if false {
                            for fi23 := 0; fi23 < 8; fi23++ {
                                var v24 bool = true
                                if ((("foo bar" + ((("\n\t" + (("hello" + "a") + "hello")) + ("\n\t" + "")) + "hello")) > "\n\t") && false) {
                                    if (167 > 32767) {
                                        for fi25 := 0; fi25 < 15; fi25++ {
                                            v26 := GenS4{F0: (v19 >> ((v12) & 63))}
                                            _ = fi25
                                            _ = v26
                                        }
                                    } else {
                                        v19 = (v19 * (v19 - 127))
                                        v12 = (v12 - (v12 * (v12 % ((2) | 1))))
                                        v12 = v12
                                    }
                                    v12 += (v12 - 728)
                                } else {
                                    var v27 int8 = 42
                                    var v28 uint8 = uint8(v12)
                                    if true {
                                        if (-256 < -1) {
                                            v29 := GenS2{F0: 0, F1: 2147483647}
                                            if (((true && false) || (-1 != 1)) || (!(v24 || (!((-(fi23)) < (fi23 | 32767)))))) {
                                                if (!(!(fi23 != fi23))) {
                                                    v2 = append(v2, fi23)
                                                    fmt.Println(fi23, v19, v28)
                                                    for fi30 := 0; fi30 < 16; fi30++ {
                                                        var v31 uint16 = 65535
                                                        _ = fi30
                                                        _ = v31
                                                    }
                                                } else {
                                                    v2 = append(v2, (fi23 * (fi23 % ((fi23) | 1))))
                                                }
                                                fmt.Println(v27, v19)
                                            } else {
                                                for fi32 := 0; fi32 < 9; fi32++ {
                                                    switch 127 {
                                                    case 2147483647:
                                                        _ = len(v2)
                                                    default:
                                                        var v33 uint64 = 4294967295
                                                        _ = v33
                                                    }
                                                    _ = fi32
                                                }
                                                v34 := GenS2{F0: -217, F1: 42}
                                                v12 = (v12 + (^(v12)))
                                                _ = v34
                                            }
                                            _ = v29
                                        }
                                    } else {
                                        v12 += (v12 + (v12 << ((v12) & 63)))
                                        for fi35 := 0; fi35 < 5; fi35++ {
                                            _ = len(v1)
                                            switch ((("" + "") + ("\n\t" + "")) + "foo bar") {
                                            case "foo bar":
                                                v28 = 0
                                            case "\n\t":
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            case "":
                                                var v36 float32 = -2.25
                                                _ = v36
                                            }
                                            if ((false && (997 >= 470)) || (-2.25 == -2.25)) {
                                                fmt.Println(v28, fi23)
                                                var v37 int = (fi23 & 9223372036854775807)
                                                _ = v37
                                            } else {
                                                v27 = (v19 - v27)
                                            }
                                            _ = fi35
                                        }
                                        fmt.Println(v19, v19, v28)
                                    }
                                    _ = v27
                                    _ = v28
                                }
                                _ = fi23
                                _ = v24
                            }
                            v19 -= (v19 >> ((v12) & 63))
                        }
                    default:
                        v19 = 100
                    }
                    if (v12 == v12) {
                        v3 = append(v3, "\n\t")
                        var v38 int32 = 846
                        fmt.Println(v12)
                        _ = v38
                    } else {
                        for fi39 := 0; fi39 < 3; fi39++ {
                            v12 += (^(v12))
                            for fi40 := 0; fi40 < 20; fi40++ {
                                fmt.Println(fi39, v12)
                                _ = fi40
                            }
                            _ = fi39
                        }
                        var v41 float32 = 1.5
                        fmt.Println(v12, v41, v19)
                        _ = v41
                    }
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v42 int16 = 710
                    v43 := GenS5{F0: v12, F1: -2147483648}
                    _ = v42
                    _ = v43
                }
                v12 = v12
            }
        case -863:
            var v44 float32 = -2.25
            _ = v44
        case -2147483648:
            v45 := GenS6{F0: "\n\t", F1: (v19 != v19), F2: 255}
            _ = v45
        default:
            v46 := GenS3{F0: 42, F1: (("" + (("hello" + "hello") + ("hello" + "foo bar"))) + "foo bar")}
            _ = v46
        }
    default:
        fmt.Println(v19, v19)
    }
    var v47 bool = false
    switch "\n\t" {
    case "a":
        var v48 uint64 = 755
        _ = v48
    default:
        var v49 uint8 = 255
        _ = v49
    }
    v12 += 18446744073709551615
    if false {
        if true {
            if v47 {
                if ((-1 > 880) && (-257 != 9223372036854775807)) {
                    if (128 <= 2) {
                        var v50 float64 = (float64(0.1) / -2.25)
                        _ = v50
                    } else {
                        for fi51 := 0; fi51 < 6; fi51++ {
                            var v52 bool = (42 == 9223372036854775807)
                            if (v12 < (^(v12))) {
                                v12 += 1
                                if len(v1) > 0 {
                                    v1[0] = 209
                                }
                            }
                            _ = fi51
                            _ = v52
                        }
                        var v53 float32 = 3.14159
                        _ = v53
                    }
                    for fi54 := 0; fi54 < 11; fi54++ {
                        for fi55 := 0; fi55 < 15; fi55++ {
                            v47 = false
                            var v56 int64 = -1
                            if (2 <= 128) {
                                var v57 string = "foo bar"
                                _ = len(v2)
                                _ = v57
                            } else {
                                v12 += v12
                                v47 = (1 == 442)
                                fmt.Println(fi55, v19, fi54)
                            }
                            _ = fi55
                            _ = v56
                        }
                        for fi58 := 0; fi58 < 1; fi58++ {
                            v59 := GenS7{F0: 473, F1: 118, F2: 2}
                            v12 -= v12
                            _ = fi58
                            _ = v59
                        }
                        _ = len(v3)
                        _ = fi54
                    }
                    for fi60 := 0; fi60 < 9; fi60++ {
                        for fi61 := 0; fi61 < 14; fi61++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v62 bool = (v47 || (fi60 >= 32767))
                            _ = fi61
                            _ = v62
                        }
                        switch -2147483648 {
                        case 4294967295:
                            if (1 == 2) {
                                v63 := GenS3{F0: 255, F1: "\n\t"}
                                for fi64 := 0; fi64 < 17; fi64++ {
                                    var v65 int8 = -128
                                    v66 := GenS5{F0: (v12 + v12), F1: 42}
                                    v65 = (^(v65))
                                    _ = fi64
                                    _ = v65
                                    _ = v66
                                }
                                v47 = (2 < 42)
                                _ = v63
                            } else {
                                var v67 uint16 = 738
                                for fi68 := 0; fi68 < 8; fi68++ {
                                    if ((v19 << (((v12 / ((806) | 1))) & 63)) != (v19 & (v19 | 0))) {
                                        var v69 string = ""
                                        _ = v69
                                    } else {
                                        v19 -= v19
                                    }
                                    if false {
                                        switch "\n\t" {
                                        case "\n\t":
                                            if (!v47) {
                                                var v70 uint64 = 18446744073709551615
                                                if len(v2) > 0 {
                                                    v2[0] = fi60
                                                }
                                                v1 = append(v1, 0)
                                                v70 = (v70 / (((^(v70))) | 1))
                                                v71 := GenS8{F0: 2}
                                                _ = v70
                                                _ = v71
                                            }
                                        case "":
                                            v72 := GenS9{F0: int(v67), F1: 42}
                                            _ = v72
                                        default:
                                            v47 = v47
                                        }
                                        var v73 string = "a"
                                        if v47 {
                                            fmt.Println(v47, fi68, v12)
                                            v2 = append(v2, 960)
                                            v19 += v19
                                        } else {
                                            fmt.Println(fi68, v47)
                                            if false {
                                                v74 := GenS9{F0: fi68, F1: (v19 >> ((v12) & 63))}
                                                fmt.Println(fi68, v19, v12)
                                                _ = len(v3)
                                                _ = v74
                                            }
                                            if true {
                                                if (63 == 4294967295) {
                                                    v75 := GenS6{F0: "foo bar", F1: v47, F2: 0}
                                                    var v76 float32 = (3.14159 - 1.5)
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    _ = v75
                                                    _ = v76
                                                }
                                                v77 := GenS10{F0: (fi68 ^ fi68), F1: 0, F2: 800}
                                                _ = v77
                                            }
                                        }
                                        _ = v73
                                    } else {
                                        v47 = (4294967295 <= (v12 ^ 1))
                                        if (!(94 > 42)) {
                                            var v78 int = 32767
                                            v67 += v67
                                            _ = v78
                                        } else {
                                            var v79 uint32 = 2147483647
                                            var v80 int8 = (^(v19))
                                            _ = v79
                                            _ = v80
                                        }
                                    }
                                    var v81 int16 = 1000
                                    _ = fi68
                                    _ = v81
                                }
                                _ = v67
                            }
                        case -9223372036854775808:
                            switch fi60 {
                            case -1:
                                switch (("a" + ("" + "a")) + "hello") {
                                case "\n\t":
                                    v47 = false
                                case "":
                                    if v47 {
                                        v3 = append(v3, ("foo bar" + ("" + "\n\t")))
                                    } else {
                                        fmt.Println(fi60)
                                    }
                                case "hello":
                                    switch "" {
                                    case "a":
                                        switch "a" {
                                        case "hello":
                                            for fi82 := 0; fi82 < 3; fi82++ {
                                                switch "\n\t" {
                                                case "a":
                                                    var v83 uint64 = 18446744073709551615
                                                    _ = v83
                                                case "":
                                                    var v84 float32 = -2.25
                                                    _ = v84
                                                case "hello":
                                                    v12 -= 42
                                                default:
                                                    v47 = true
                                                }
                                                fmt.Println(v19, v47, v19)
                                                if (!(0 >= -155)) {
                                                    v85 := GenS2{F0: 42, F1: 65535}
                                                    _ = v85
                                                }
                                                _ = fi82
                                            }
                                        case "foo bar":
                                            v86 := GenS11{F0: 77, F1: fi60, F2: 0}
                                            _ = v86
                                        case "\n\t":
                                            _ = len(v1)
                                        default:
                                            fmt.Println(v47)
                                        }
                                    case "\n\t":
                                        switch fi60 {
                                        case -111:
                                            v19 -= v19
                                        }
                                    }
                                }
                            case 533:
                                fmt.Println(fi60, v47, v47)
                            default:
                                fmt.Println(v12)
                            }
                        case 0:
                            for fi87 := 0; fi87 < 16; fi87++ {
                                v88 := GenS6{F0: "a", F1: (v47 && v47), F2: 244}
                                fmt.Println(fi87)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = fi87
                                _ = v88
                            }
                        }
                        _ = fi60
                    }
                } else {
                    var v89 bool = v47
                    var v90 uint8 = 255
                    _ = v89
                    _ = v90
                }
            }
            var v91 int8 = (v19 * v19)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v91
        }
        if v47 {
            var v92 int = 42
            v12 += v12
            for fi93 := 0; fi93 < 1; fi93++ {
                var v94 uint8 = 127
                _ = fi93
                _ = v94
            }
            _ = v92
        }
        for fi95 := 0; fi95 < 17; fi95++ {
            if v47 {
                if v47 {
                    switch 42 {
                    case 0:
                        v12 = v12
                    case 9223372036854775807:
                        var v96 uint64 = 4294967295
                        _ = v96
                    case 32767:
                        var v97 int64 = 9223372036854775807
                        _ = v97
                    default:
                        var v98 float32 = 3.14159
                        _ = v98
                    }
                    for fi99 := 0; fi99 < 19; fi99++ {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi99
                    }
                    v100 := GenS12{F0: 42, F1: 36, F2: -1}
                    _ = v100
                }
                for fi101 := 0; fi101 < 14; fi101++ {
                    if len(v1) > 0 {
                        v1[0] = 0
                    }
                    fmt.Println(fi95, fi95, fi95)
                    switch fi95 {
                    case 32768:
                        v102 := GenS8{F0: 867}
                        _ = v102
                    default:
                        var v103 uint = v12
                        _ = v103
                    }
                    _ = fi101
                }
                if ("\n\t" >= "a") {
                    v12 -= (v12 + v12)
                }
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi95
        }
    }
    switch -973 {
    case -121:
        for fi104 := 0; fi104 < 20; fi104++ {
            _ = len(v3)
            _ = fi104
        }
    case 9223372036854775807:
        v105 := GenS13{F0: -9223372036854775808}
        _ = v105
    case 128:
        v106 := GenS7{F0: 562, F1: 1, F2: 18446744073709551615}
        _ = v106
    default:
        v47 = (v47 && false)
    }
    v107 := GenS7{F0: 898, F1: 717, F2: 4294967295}
    v108 := GenS3{F0: 255, F1: ("foo bar" + (("" + "\n\t") + "\n\t"))}
    v12 -= 255
    if true {
        fmt.Println(v12, v12)
    } else {
        for fi109 := 0; fi109 < 7; fi109++ {
            if true {
                var v110 int16 = 0
                v110 -= (v110 | 42)
                for fi111 := 0; fi111 < 18; fi111++ {
                    var v112 uint64 = genFn1((fi109 - fi111), 18446744073709551615, fi109)
                    fmt.Println(v47)
                    if v47 {
                        var v113 uint8 = 1
                        v47 = v47
                        _ = v113
                    } else {
                        var v114 float32 = (1e10 - 1.5)
                        _ = len(v1)
                        var v115 int64 = 0
                        _ = v114
                        _ = v115
                    }
                    _ = fi111
                    _ = v112
                }
                _ = v110
            }
            if v47 {
                for fi116 := 0; fi116 < 4; fi116++ {
                    v2 = append(v2, fi116)
                    switch (fi116 & (fi116 % ((fi116) | 1))) {
                    case -806:
                        if v47 {
                            if v47 {
                                _ = len(v3)
                            } else {
                                switch fi109 {
                                case 793:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case 1:
                                    fmt.Println(fi109)
                                default:
                                    v19 += v19
                                }
                                v1 = append(v1, 65535)
                            }
                        }
                    case 1:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    var v117 bool = (1 >= 0)
                    _ = fi116
                    _ = v117
                }
            } else {
                var v118 uint8 = 146
                _ = v118
            }
            _ = fi109
        }
        for fi119 := 0; fi119 < 15; fi119++ {
            for fi120 := 0; fi120 < 20; fi120++ {
                v2 = append(v2, fi120)
                switch "\n\t" {
                case "foo bar":
                    switch -2147483648 {
                    case 1:
                        _ = len(v2)
                    case 32767:
                        v19 -= -121
                    case 4294967295:
                        if (("foo bar" + "a") >= "a") {
                            var v121 int = (fi119 + fi119)
                            if len(v3) > 0 {
                                v3[0] = "hello"
                            }
                            v12 += v12
                            _ = v121
                        }
                    default:
                        switch "" {
                        case "":
                            var v122 uint8 = 2
                            _ = v122
                        case "hello":
                            v123 := GenS2{F0: 1, F1: 2147483647}
                            _ = v123
                        default:
                            fmt.Println(fi119, v47)
                        }
                    }
                case "":
                    var v124 int16 = 1000
                    _ = v124
                default:
                    var v125 int64 = -16
                    _ = v125
                }
                _ = fi120
            }
            _ = fi119
        }
        var v126 int8 = (v19 + 127)
        _ = v126
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v107
    _ = v108
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v47
}
