package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 bool }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 bool; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint16; F1 uint32; F2 uint8 }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 uint32; F1 int16 }
type GenS7 struct { F0 uint; F1 uint }
type GenS8 struct { F0 int64; F1 int; F2 uint8 }
type GenS9 struct { F0 uint64; F1 uint; F2 bool }
type GenS10 struct { F0 uint8; F1 uint16 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS13 struct { F0 int8; F1 uint64; F2 int16 }
func genFn0(p0 bool, p1 int16, p2 int16) bool {
    if (!p0) {
        var v7 uint = 2
        _ = v7
    } else {
        var v8 int16 = p2
        var v9 uint = 308
        _ = v8
        _ = v9
    }
    fmt.Println(p0, p1)
    for fi10 := 0; fi10 < 2; fi10++ {
        if p0 {
            var v11 string = "foo bar"
            for fi12 := 0; fi12 < 11; fi12++ {
                v13 := GenS1{F0: 65535, F1: 128, F2: p0}
                _ = fi12
                _ = v13
            }
            v14 := GenS1{F0: 460, F1: 127, F2: false}
            _ = v11
            _ = v14
        } else {
            for fi15 := 0; fi15 < 7; fi15++ {
                var v16 uint = 1
                p2 += (p2 | p2)
                p1 += p1
                if true {
                    if ((v16 << ((v16) & 63)) != v16) {
                        v17 := GenS2{F0: 2147483648, F1: -1}
                        p0 = (!true)
                        fmt.Println(fi10, p2)
                        _ = v17
                    } else {
                        var v18 int8 = -1
                        var v19 uint64 = 42
                        if p0 {
                            p1 = 341
                            var v20 bool = true
                            p1 += p2
                            var v21 float64 = ((-2.25 - 3.14159) / ((float64(-2.25) * ((1.5 * 0.1) * 3.14159)) - (float64(0.0) - (float64((-2.25 + 0.0)) + ((float64(1e10) - (-2.25 - (float64(1.5) + (float64(float32(((3.14159 - 1.5) / (3.14159 / 1e10)))) * 0.0)))) + -2.25)))))
                            if (1.5 > 1.5) {
                                fmt.Println(p1, v21)
                                v22 := GenS3{F0: p0, F1: 330, F2: 0}
                                fmt.Println(v19, v19)
                                _ = v22
                            }
                            var v23 bool = (!((-1 < 1) || (!((1.5 - 0.1) > 0.1))))
                            _ = v20
                            _ = v21
                            _ = v23
                        } else {
                            fmt.Println(v16)
                        }
                        switch "hello" {
                        case "foo bar":
                            v24 := GenS1{F0: 42, F1: 119, F2: (((0.1 + (0.0 * 3.14159)) * (0.0 + (1e10 + float32(float64(0.0))))) > 3.14159)}
                            _ = v24
                        }
                        _ = v18
                        _ = v19
                    }
                    v25 := []int{}
                    var v26 int8 = 127
                    _ = v25
                    _ = v26
                }
                var v27 int8 = -71
                _ = fi15
                _ = v16
                _ = v27
            }
        }
        var v28 uint8 = 0
        p0 = ((p1 / (((p1 + 276)) | 1)) > 42)
        _ = fi10
        _ = v28
    }
    p2 = p2
    p1 = (p2 * (p1 & (^(p2))))
    _ = p0
    _ = p1
    _ = p2
    return (2 < 0)
}
func genFn1(p0 string, p1 uint16, p2 uint16) uint {
    p2 = 755
    if true {
        p0 = (p0 + "hello")
        if (4294967295 != 2) {
            var v29 uint = 1
            _ = v29
        }
    } else {
        switch 1 {
        case -349:
            v30 := GenS2{F0: 65535, F1: -2147483648}
            _ = v30
        case 2147483647:
            p2 = 0
        default:
            for fi31 := 0; fi31 < 14; fi31++ {
                p2 = (p1 + p2)
                _ = fi31
            }
        }
    }
    p1 += (^(p2))
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 string = ""
    v2 := make(map[string]string)
    v3 := []int{9223372036854775807}
    var v4 bool = false
    var v5 float32 = 1e10
    v6 := make(map[uint32]string)
    v4 = ((((v1 + ("" + "foo bar")) + (v1 + v1)) + (v1 + (v1 + v1))) < (v1 + (v1 + "hello")))
    var v32 int64 = -1
    v33 := GenS3{F0: (1 < 1), F1: -9223372036854775808, F2: 691}
    v1 = ("foo bar" + "")
    for fi34 := 0; fi34 < 20; fi34++ {
        if (!v4) {
            fmt.Println(v4, v4)
            if false {
                for fi35 := 0; fi35 < 10; fi35++ {
                    var v36 int = fi35
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if (v32 == v32) {
                        delete(v2, "\n\t")
                        var v37 int = (fi35 | 4294967295)
                        v38 := GenS2{F0: 226, F1: 1}
                        if v4 {
                            _ = len(v3)
                            v36 = v37
                            if (!v4) {
                                fmt.Println(v5, v32)
                                v32 = (v32 >> ((4294967295) & 63))
                            } else {
                                switch ((v1 + ("a" + ((v1 + ("hello" + v1)) + v1))) + v1) {
                                case "":
                                    var v39 float32 = v5
                                    _ = v39
                                case "\n\t":
                                    var v40 int16 = 32767
                                    _ = v40
                                }
                            }
                            v36 -= (v37 & (v36 + v37))
                            var v41 int32 = -2147483648
                            if (false || ((1e10 < 3.14159) && v4)) {
                                v36 = (fi34 * fi35)
                            }
                            _ = v41
                        } else {
                            if ((!true) && v4) {
                                var v42 uint32 = 1
                                v5 -= v5
                                v43 := GenS2{F0: 0, F1: -2147483648}
                                delete(v2, "a")
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v42
                                _ = v43
                            } else {
                                v44 := GenS3{F0: (v4 && v4), F1: (v32 + (v32 % (((v32 - v32)) | 1))), F2: 2}
                                if ((v4 && false) && ((v32 + v32) != 42)) {
                                    v5 -= v5
                                    if ((-1 == 127) && v4) {
                                        if v4 {
                                            fmt.Println(v4, v37)
                                            switch fi34 {
                                            case 1:
                                                v5 += 3.14159
                                            case -324:
                                                var v45 int8 = 82
                                                _ = v45
                                            }
                                            var v46 int16 = -32768
                                            _ = v46
                                        }
                                        switch v1 {
                                        case "hello":
                                            v4 = (1000 != 13)
                                        case "a":
                                            v4 = true
                                        case "foo bar":
                                            var v47 float64 = 0.1
                                            _ = v47
                                        }
                                    } else {
                                        fmt.Println(v4, v1)
                                    }
                                    v4 = v4
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                } else {
                                    v2[""] = v1
                                }
                                if len(v3) > 0 {
                                    v3[0] = -9223372036854775808
                                }
                                if v4 {
                                    if true {
                                        var v48 uint32 = 0
                                        if (v4 && false) {
                                            v48 = (v48 / (((^(v48))) | 1))
                                            v49 := GenS1{F0: (v48 % ((v48) | 1)), F1: 0, F2: false}
                                            v50 := GenS3{F0: (65535 > 4294967295), F1: (v32 << ((663) & 63)), F2: 4294967295}
                                            switch (v1 + "\n\t") {
                                            case "\n\t":
                                                switch (v1 + ("hello" + v1)) {
                                                case "hello":
                                                    v51 := GenS2{F0: 1, F1: -1}
                                                    _ = v51
                                                case "":
                                                    v52 := GenS4{F0: 387, F1: (v48 ^ 2147483648), F2: 255}
                                                    _ = v52
                                                }
                                            case "foo bar":
                                                fmt.Println(v5, fi34)
                                            default:
                                                var v53 uint = 356
                                                _ = v53
                                            }
                                            _ = v49
                                            _ = v50
                                        }
                                        v3 = append(v3, (^(fi35)))
                                        if (0 < 1) {
                                            v54 := v2[v1]
                                            var v55 string = ""
                                            _ = v54
                                            _ = v55
                                        }
                                        v4 = v4
                                        _ = v48
                                    } else {
                                        v56 := GenS4{F0: 1, F1: 42, F2: 253}
                                        v4 = (4294967295 < 132)
                                        v1 = "foo bar"
                                        _ = v56
                                    }
                                }
                                _ = v44
                            }
                            fmt.Println(v37, v5, v1)
                            fmt.Println(v4, v4, v36)
                        }
                        if true {
                            var v57 int8 = -1
                            v58 := GenS2{F0: 264, F1: -2147483648}
                            switch 32767 {
                            case 641:
                                if ("foo bar" != "foo bar") {
                                    v5 = float32(1.5)
                                    v32 = (^(v32))
                                    v57 += 1
                                }
                            case 9223372036854775807:
                                v37 += 573
                            case 32768:
                                fmt.Println(v32)
                            default:
                                v3 = append(v3, (fi35 * (v37 & 329)))
                            }
                            v57 = v57
                            _ = v57
                            _ = v58
                        }
                        _ = v37
                        _ = v38
                    } else {
                        switch "" {
                        case "":
                            if true {
                                var v59 uint16 = 389
                                var v60 string = (v1 + (v1 + v1))
                                _ = v59
                                _ = v60
                            }
                        case "hello":
                            var v61 bool = (-2147483648 > 2147483647)
                            _ = v61
                        default:
                            var v62 int64 = (v32 / ((0) | 1))
                            _ = v62
                        }
                    }
                    _ = fi35
                    _ = v36
                }
            } else {
                v63 := GenS3{F0: v4, F1: int64(fi34), F2: 1}
                _ = v63
            }
        } else {
            if v4 {
                _ = len(v2)
                var v64 int32 = 1
                v65 := GenS2{F0: 2147483648, F1: 0}
                _ = v64
                _ = v65
            }
        }
        _ = fi34
    }
    v6[65535] = ("foo bar" + v1)
    var v66 int = 9223372036854775807
    var v67 int16 = -32768
    v68 := GenS5{F0: 42, F1: 100}
    v69 := GenS5{F0: 0, F1: 42}
    v67 -= v67
    if (v4 || (v1 < "\n\t")) {
        v66 -= (v66 * v66)
    } else {
        v1 = ""
        var v70 int = v66
        var v71 uint64 = 2
        _ = v70
        _ = v71
    }
    if false {
        if true {
            _ = len(v6)
            if ((v32 >= (v32 % ((-9223372036854775808) | 1))) || v4) {
                for fi72 := 0; fi72 < 13; fi72++ {
                    var v73 int16 = (v67 << ((18446744073709551615) & 63))
                    if (0 != -9223372036854775808) {
                        for fi74 := 0; fi74 < 12; fi74++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            var v75 int8 = -128
                            if false {
                                var v76 float64 = 1.5
                                if false {
                                    if (v32 < v32) {
                                        var v77 float32 = v5
                                        _ = v77
                                    } else {
                                        v73 = v73
                                        v73 -= 42
                                        v6[42] = ("" + v1)
                                    }
                                    _ = len(v2)
                                } else {
                                    switch (fi74 & (^(fi74))) {
                                    case 127:
                                        v78 := GenS6{F0: 1, F1: (v67 >> ((186) & 63))}
                                        _ = v78
                                    case -1:
                                        var v79 float64 = v76
                                        _ = v79
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = -1
                                    }
                                }
                                _ = v76
                            }
                            _ = fi74
                            _ = v75
                        }
                    } else {
                        if (0 <= 2) {
                            if ((v5 + v5) != ((v5 / v5) / 0.1)) {
                                if (v32 == 223) {
                                    var v80 uint64 = 0
                                    var v81 string = (((v1 + "hello") + ("a" + v1)) + v1)
                                    _ = v80
                                    _ = v81
                                } else {
                                    v73 = v73
                                    v82 := GenS7{F0: 65535, F1: 255}
                                    for fi83 := 0; fi83 < 8; fi83++ {
                                        var v84 uint8 = 255
                                        v67 = -503
                                        _ = fi83
                                        _ = v84
                                    }
                                    _ = v82
                                }
                                v67 -= v73
                                var v85 uint64 = 4294967295
                                _ = v85
                            } else {
                                v86 := GenS8{F0: -1, F1: v66, F2: 1}
                                v67 -= v73
                                _ = v86
                            }
                            if true {
                                if v4 {
                                    v67 -= 839
                                    for fi87 := 0; fi87 < 17; fi87++ {
                                        if false {
                                            v88 := v2[(v1 + ((((("foo bar" + "foo bar") + v1) + "") + v1) + v1))]
                                            if (!true) {
                                                v1 = (v88 + "a")
                                                v89 := GenS5{F0: 234, F1: 127}
                                                _ = v89
                                            }
                                            fmt.Println(v1, v67, v1)
                                            _ = v88
                                        }
                                        v1 = "\n\t"
                                        if (v4 && false) {
                                            _ = len(v2)
                                            var v90 float32 = (1e10 - v5)
                                            _ = v90
                                        }
                                        v32 = (v32 * 42)
                                        if (1 < -660) {
                                            if false {
                                                v67 = (v67 & int16(fi72))
                                                var v91 bool = v4
                                                var v92 int64 = v32
                                                _ = v91
                                                _ = v92
                                            } else {
                                                fmt.Println(fi87, fi87, v67)
                                                if v4 {
                                                    var v93 uint32 = 0
                                                    switch -169 {
                                                    case -1:
                                                        v94 := GenS5{F0: 127, F1: -128}
                                                        _ = v94
                                                    default:
                                                        var v95 uint = genFn1((("\n\t" + "foo bar") + v1), 32767, 42)
                                                        _ = v95
                                                    }
                                                    var v96 uint32 = (^(v93))
                                                    _ = v93
                                                    _ = v96
                                                }
                                                v97 := GenS5{F0: 160, F1: -111}
                                                if len(v3) > 0 {
                                                    v3[0] = (^(fi72))
                                                }
                                                v98 := GenS5{F0: 0, F1: 100}
                                                _ = v97
                                                _ = v98
                                            }
                                        } else {
                                            var v99 int32 = 336
                                            _ = v99
                                        }
                                        _ = fi87
                                    }
                                    v100 := GenS9{F0: 715, F1: genFn1((v1 + ("\n\t" + v1)), 59, 42), F2: ((((0.0 + 0.1) + (1e10 / (((0.1 - 0.0) / -2.25) + 3.14159))) + 1e10) == 0.0)}
                                    _ = v100
                                }
                                v66 = fi72
                            } else {
                                if (v5 <= -2.25) {
                                    _ = len(v6)
                                    v3 = append(v3, 32767)
                                }
                                v6[65535] = (("hello" + "hello") + "hello")
                            }
                            v101 := GenS4{F0: 2, F1: 2147483647, F2: 1}
                            _ = v101
                        }
                        v102 := GenS8{F0: -1, F1: (v66 % (((fi72 << ((2) & 63))) | 1)), F2: 0}
                        _ = v102
                    }
                    _ = fi72
                    _ = v73
                }
            } else {
                if (true && (false || v4)) {
                    if ((v66 - -419) < (v66 >> ((42) & 63))) {
                        var v103 uint = 65535
                        var v104 uint32 = 2147483648
                        var v105 uint = (v103 >> (((^(v103))) & 63))
                        _ = v103
                        _ = v104
                        _ = v105
                    }
                    if (100 > int8(v67)) {
                        fmt.Println(v5, v66, v67)
                        v2[v1] = v1
                    }
                    for fi106 := 0; fi106 < 10; fi106++ {
                        v5 = (float32(0.1) / (v5 * v5))
                        fmt.Println(v32)
                        v5 += v5
                        _ = fi106
                    }
                }
                if len(v3) > 0 {
                    v3[0] = v66
                }
            }
            if ((1.5 / (1.5 / 1.0)) > 1e10) {
                switch v66 {
                case 4294967295:
                    v4 = v4
                case -373:
                    var v107 int16 = v67
                    _ = v107
                case 216:
                    v108 := GenS10{F0: 189, F1: 65535}
                    _ = v108
                default:
                    v109 := GenS9{F0: 0, F1: 65535, F2: (369 != 0)}
                    _ = v109
                }
            } else {
                if (163 > 4294967295) {
                    if (false && (65535 < 42)) {
                        if true {
                            for fi110 := 0; fi110 < 10; fi110++ {
                                if false {
                                    _ = len(v3)
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    switch (fi110 / (((-(fi110))) | 1)) {
                                    case 84:
                                        v6[494] = v1
                                    case -615:
                                        v32 -= (-(v32))
                                    case -1:
                                        if ((v67 + 1) >= -760) {
                                            var v111 int = 192
                                            v4 = (v4 && ((v32 >> ((18446744073709551615) & 63)) != (v32 * v32)))
                                            if (genFn1((v1 + v1), 255, 628) == 255) {
                                                v111 = (-(v66))
                                                if ((310 >= 496) || v4) {
                                                    var v112 float64 = -2.25
                                                    switch 9223372036854775807 {
                                                    case 2147483647:
                                                        _ = len(v2)
                                                    case 1:
                                                        for fi113 := 0; fi113 < 4; fi113++ {
                                                            var v114 float64 = v112
                                                            v115 := GenS6{F0: 0, F1: v67}
                                                            v67 -= 0
                                                            _ = fi113
                                                            _ = v114
                                                            _ = v115
                                                        }
                                                    default:
                                                        for fi116 := 0; fi116 < 16; fi116++ {
                                                            switch v1 {
                                                            case "\n\t":
                                                                if true {
                                                                    var v117 int8 = 1
                                                                    v118 := GenS6{F0: 0, F1: (^(v67))}
                                                                    v119 := GenS11{F0: -2147483648}
                                                                    v111 -= (v66 + fi116)
                                                                    _ = v117
                                                                    _ = v118
                                                                    _ = v119
                                                                } else {
                                                                    var v120 uint8 = 46
                                                                    _ = v120
                                                                }
                                                            case "a":
                                                                if (false && ((v67 >> ((65535) & 63)) > 42)) {
                                                                    v5 = 0.0
                                                                    var v121 int8 = 127
                                                                    if (128 < 0) {
                                                                        if true {
                                                                            var v122 int8 = -1
                                                                            switch fi110 {
                                                                            case 103:
                                                                                v123 := GenS7{F0: genFn1((v1 + v1), 0, 42), F1: 42}
                                                                                _ = v123
                                                                            case -290:
                                                                                var v124 float32 = v5
                                                                                _ = v124
                                                                            case -592:
                                                                                switch (fi110 << ((255) & 63)) {
                                                                                case 636:
                                                                                    switch (fi116 / ((0) | 1)) {
                                                                                    case 32767:
                                                                                        v122 = v122
                                                                                    default:
                                                                                        v125 := GenS12{F0: -1, F1: 65535, F2: 942}
                                                                                        _ = v125
                                                                                    }
                                                                                }
                                                                            default:
                                                                                var v126 uint32 = 164
                                                                                _ = v126
                                                                            }
                                                                            fmt.Println(v111)
                                                                            _ = v122
                                                                        }
                                                                        v1 = "hello"
                                                                    } else {
                                                                        if v4 {
                                                                            v32 -= v32
                                                                            _ = len(v3)
                                                                            var v127 int = 127
                                                                            _ = v127
                                                                        }
                                                                        fmt.Println(v66)
                                                                    }
                                                                    _ = v121
                                                                }
                                                                var v128 float64 = -2.25
                                                                v112 = v128
                                                                _ = v128
                                                            default:
                                                                v66 -= 0
                                                                var v129 float64 = 1.5
                                                                _ = v129
                                                            }
                                                            v111 += v111
                                                            _ = fi116
                                                        }
                                                    }
                                                    _ = v112
                                                } else {
                                                    v130 := GenS12{F0: 1, F1: 716, F2: 347}
                                                    _ = v130
                                                }
                                                var v131 int = (v111 - (fi110 / ((v66) | 1)))
                                                _ = v131
                                            } else {
                                                var v132 bool = true
                                                _ = v132
                                            }
                                            _ = v111
                                        }
                                    default:
                                        var v133 int16 = (v67 - (v67 - (v67 % (((v67 - 1)) | 1))))
                                        _ = v133
                                    }
                                }
                                if len(v3) > 0 {
                                    v3[0] = v66
                                }
                                _ = fi110
                            }
                        } else {
                            var v134 int8 = -1
                            switch (-(v66)) {
                            case 4294967295:
                                if (true && v4) {
                                    fmt.Println(v66, v134)
                                    var v135 int32 = 2147483647
                                    if v4 {
                                        for fi136 := 0; fi136 < 12; fi136++ {
                                            switch v1 {
                                            case "\n\t":
                                                v3 = append(v3, v66)
                                            case "":
                                                switch -1 {
                                                case 94:
                                                    v1 = "a"
                                                case 42:
                                                    v6[4294967295] = "a"
                                                case -5:
                                                    var v137 string = "foo bar"
                                                    _ = v137
                                                }
                                            case "foo bar":
                                                switch 387 {
                                                case -9223372036854775808:
                                                    if (1 < 1) {
                                                        switch v66 {
                                                        case -48:
                                                            v6[4294967295] = v1
                                                        }
                                                        v138 := GenS13{F0: v134, F1: 6, F2: -32768}
                                                        v135 = (v135 << ((639) & 63))
                                                        _ = v138
                                                    }
                                                default:
                                                    v135 += (v135 ^ v135)
                                                }
                                            default:
                                                if ((v32 | v32) != v32) {
                                                    for fi139 := 0; fi139 < 17; fi139++ {
                                                        v140 := GenS3{F0: v4, F1: (v32 ^ 9223372036854775807), F2: 18446744073709551615}
                                                        v134 += int8(v66)
                                                        v32 = (v32 | (v32 / ((-204) | 1)))
                                                        _ = fi139
                                                        _ = v140
                                                    }
                                                }
                                            }
                                            _ = fi136
                                        }
                                        v141 := GenS10{F0: 2, F1: 65535}
                                        _ = v141
                                    }
                                    _ = v135
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                            default:
                                var v142 string = v1
                                _ = v142
                            }
                            var v143 uint32 = 4294967295
                            _ = v134
                            _ = v143
                        }
                    }
                    v4 = (1.5 > 1.5)
                }
            }
        } else {
            if len(v3) > 0 {
                v3[0] = (^(v66))
            }
        }
        v32 -= v32
        fmt.Println(v1, v67)
    } else {
        var v144 uint8 = 27
        _ = v144
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v66
    _ = v67
    _ = v68
    _ = v69
}
