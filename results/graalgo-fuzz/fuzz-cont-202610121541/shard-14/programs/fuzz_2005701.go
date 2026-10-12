package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint; F2 int }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 bool; F1 int8 }
type GenS6 struct { F0 uint32; F1 int }
type GenS7 struct { F0 int; F1 int8; F2 int16 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS10 struct { F0 string; F1 int64 }
func genFn0(p0 int64, p1 uint16) int16 {
    var v7 uint = 1
    if false {
        if false {
            v8 := GenS1{F0: 2147483647, F1: 255}
            _ = v8
        }
        v9 := GenS1{F0: 42, F1: (^(p1))}
        _ = v9
    } else {
        v7 += v7
        p0 = (p0 & 1)
        v10 := []int{32767, -733}
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v7
    return 1000
}
func genFn1(p0 string, p1 uint32, p2 int8) int64 {
    if false {
        fmt.Println(p0, p2)
    } else {
        if (32767 == 598) {
            var v11 float32 = 0.0
            _ = v11
        }
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 int32 = 407
            v14 := []int{fi12, 0}
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    fmt.Println(p0, p2, p0)
    switch "hello" {
    case "hello":
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 int16 = genFn0(-1, 1)
            _ = fi15
            _ = v16
        }
    case "a":
        switch (p0 + (p0 + p0)) {
        case "foo bar":
            var v17 uint = 1
            _ = v17
        case "\n\t":
            var v18 uint8 = 255
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        if (233 > 0) {
            v20 := []int{32768, (fi19 | fi19)}
            _ = v20
        } else {
            var v21 float32 = 1e10
            var v22 bool = (!true)
            _ = v21
            _ = v22
        }
        p2 = p2
        for fi23 := 0; fi23 < 19; fi23++ {
            switch (("foo bar" + "foo bar") + p0) {
            case "foo bar":
                v24 := GenS2{F0: 611, F1: 18446744073709551615, F2: (fi23 * (^(fi19)))}
                _ = v24
            case "":
                var v25 int64 = -9223372036854775808
                v26 := GenS2{F0: 1, F1: 65535, F2: fi19}
                _ = v25
                _ = v26
            case "\n\t":
                v27 := []int{(fi19 | (fi23 ^ fi19))}
                _ = v27
            }
            if (!true) {
                var v28 int32 = 0
                var v29 uint8 = uint8(p1)
                fmt.Println(fi23, fi19, p1)
                if (genFn0(-1, 32767) <= 1) {
                    v30 := []int{32767}
                    var v31 float64 = ((0.0 * float64(3.14159)) + (float64(float32(0.1)) + 3.14159))
                    _ = v30
                    _ = v31
                } else {
                    if (42 <= -1) {
                        var v32 int32 = v28
                        if ((!((0 <= -1) || (!false))) && (!true)) {
                            var v33 uint = 18446744073709551615
                            _ = v33
                        }
                        var v34 int8 = 1
                        var v35 uint64 = 198
                        _ = v32
                        _ = v34
                        _ = v35
                    }
                    if (115 <= v29) {
                        switch (fi23 ^ -303) {
                        case 54:
                            if ((^(v28)) < v28) {
                                var v36 bool = true
                                var v37 float32 = (float32(-2.25) / (-2.25 + 1.5))
                                switch (-(fi23)) {
                                case 128:
                                    switch "foo bar" {
                                    case "foo bar":
                                        var v38 uint16 = 32767
                                        _ = v38
                                    case "a":
                                        v39 := GenS1{F0: (p1 + (^(p1))), F1: 2}
                                        _ = v39
                                    case "\n\t":
                                        var v40 uint32 = (^(p1))
                                        _ = v40
                                    default:
                                        switch ("a" + "hello") {
                                        case "\n\t":
                                            v28 -= v28
                                        case "foo bar":
                                            p2 = p2
                                        }
                                    }
                                    fmt.Println(fi23)
                                    switch (fi23 + -712) {
                                    case 9223372036854775807:
                                        var v41 int8 = p2
                                        _ = v41
                                    case -2147483648:
                                        v29 = v29
                                    default:
                                        v42 := []int{(fi19 >> ((500) & 63)), fi19, 4294967295, (fi23 ^ (fi19 + fi19))}
                                        _ = v42
                                    }
                                case -2147483648:
                                    v28 -= v28
                                case 32768:
                                    var v43 uint = 1
                                    _ = v43
                                }
                                _ = v36
                                _ = v37
                            } else {
                                if true {
                                    v28 = (v28 << ((44) & 63))
                                } else {
                                    v44 := GenS3{F0: (v28 >> ((400) & 63)), F1: 65535}
                                    _ = v44
                                }
                                p1 = (p1 - (p1 % (((p1 + p1)) | 1)))
                            }
                        case 128:
                            v45 := GenS1{F0: (p1 >> ((2) & 63)), F1: 32767}
                            _ = v45
                        case 42:
                            switch "" {
                            case "a":
                                p0 = ((p0 + p0) + "a")
                            case "":
                                if ((p2 * p2) < (p2 - (-(p2)))) {
                                    var v46 bool = true
                                    _ = v46
                                }
                            default:
                                var v47 uint16 = 149
                                _ = v47
                            }
                        default:
                            p1 = 1
                        }
                        fmt.Println(p0, fi23, fi23)
                        var v48 float32 = -2.25
                        var v49 uint64 = 85
                        _ = v48
                        _ = v49
                    }
                    if false {
                        var v50 uint8 = v29
                        if false {
                            switch fi19 {
                            case 32767:
                                if true {
                                    var v51 uint16 = 891
                                    var v52 string = p0
                                    var v53 int32 = (^(v28))
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                } else {
                                    var v54 int32 = -16
                                    v50 = v50
                                    var v55 float32 = (0.1 - (1e10 + -2.25))
                                    _ = v54
                                    _ = v55
                                }
                            case -2147483648:
                                v56 := GenS1{F0: p1, F1: 0}
                                _ = v56
                            default:
                                v57 := GenS1{F0: (p1 ^ (p1 >> ((244) & 63))), F1: 1}
                                _ = v57
                            }
                            v58 := []int{(fi23 ^ (fi19 & (^(fi19)))), fi23, fi19}
                            v58 = append(v58, fi19)
                            _ = v58
                        }
                        var v59 uint32 = 65535
                        _ = v50
                        _ = v59
                    }
                }
                _ = v28
                _ = v29
            } else {
                var v60 uint8 = 127
                if true {
                    if false {
                        p2 = 42
                        v61 := GenS1{F0: (p1 % ((uint32(v60)) | 1)), F1: 0}
                        v62 := GenS2{F0: 0, F1: uint(p1), F2: fi19}
                        _ = v61
                        _ = v62
                    } else {
                        switch (fi19 & fi19) {
                        case 127:
                            switch (^(fi19)) {
                            case 42:
                                var v63 int = 127
                                _ = v63
                            case -1:
                                fmt.Println(p2, p2, fi19)
                            case 793:
                                switch p0 {
                                case "foo bar":
                                    if true {
                                        v64 := []int{}
                                        _ = v64
                                    } else {
                                        v65 := []int{fi23, fi23, (fi19 + fi23)}
                                        _ = v65
                                    }
                                case "":
                                    var v66 uint64 = 4294967295
                                    _ = v66
                                default:
                                    if (false && true) {
                                        if (false && false) {
                                            var v67 uint8 = (v60 + v60)
                                            _ = v67
                                        } else {
                                            v68 := []int{}
                                            var v69 string = "\n\t"
                                            fmt.Println(v69)
                                            _ = v68
                                            _ = v69
                                        }
                                        v70 := GenS4{F0: 2}
                                        fmt.Println(p0, fi19)
                                        _ = v70
                                    }
                                }
                            default:
                                p1 = (^(p1))
                            }
                            v71 := GenS1{F0: 493, F1: 0}
                            v72 := GenS1{F0: (p1 & (p1 + p1)), F1: 434}
                            switch fi19 {
                            case 127:
                                if false {
                                    p0 = ""
                                    var v73 int64 = 1
                                    switch fi19 {
                                    case -9223372036854775808:
                                        fmt.Println(p1, p2)
                                    case 2147483647:
                                        var v74 float64 = ((0.1 * 0.0) + -2.25)
                                        _ = v74
                                    case -2147483648:
                                        p2 = -98
                                    }
                                    _ = v73
                                }
                            case 942:
                                if (-648 < 1) {
                                    v75 := []int{fi19, fi19}
                                    _ = len(v75)
                                    v76 := GenS4{F0: 65535}
                                    _ = v75
                                    _ = v76
                                }
                            case 32767:
                                p2 = 0
                            default:
                                v77 := GenS4{F0: 882}
                                _ = v77
                            }
                            v78 := []int{14, (fi19 << ((4294967295) & 63)), fi23}
                            p1 = (p1 / ((2147483647) | 1))
                            fmt.Println(p1, p2, fi23)
                            _ = v71
                            _ = v72
                            _ = v78
                        }
                        var v79 bool = (true || (255 > 18446744073709551615))
                        if (681 > 2) {
                            var v80 uint32 = p1
                            v79 = (0 == 1000)
                            _ = v80
                        }
                        var v81 uint16 = 146
                        _ = v79
                        _ = v81
                    }
                    if true {
                        switch p0 {
                        case "foo bar":
                            switch fi23 {
                            case 128:
                                var v82 int32 = 496
                                _ = v82
                            case -590:
                                if false {
                                    if (1.5 < 0.0) {
                                        var v83 float64 = (((-2.25 + (0.1 / 3.14159)) / 3.14159) * float64(-2.25))
                                        _ = v83
                                    }
                                    if false {
                                        var v84 int = 127
                                        p0 = "hello"
                                        _ = v84
                                    } else {
                                        var v85 uint32 = (p1 + (p1 | p1))
                                        _ = v85
                                    }
                                    var v86 int64 = 42
                                    p0 = "\n\t"
                                    _ = v86
                                }
                            }
                        case "\n\t":
                            p2 = p2
                        default:
                            p2 -= p2
                        }
                        v87 := GenS5{F0: false, F1: 42}
                        v60 += v60
                        var v88 uint = 255
                        _ = v87
                        _ = v88
                    }
                } else {
                    v89 := []int{fi19}
                    v89 = append(v89, (fi19 | 128))
                    v90 := GenS6{F0: p1, F1: fi19}
                    _ = v89
                    _ = v90
                }
                _ = v60
            }
            v91 := []int{fi23, fi23, fi19, (fi19 ^ (fi23 ^ (-(fi23))))}
            _ = fi23
            _ = v91
        }
        _ = fi19
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]string)
    var v3 int16 = 0
    var v4 float32 = 1.5
    var v5 float32 = 1.5
    var v6 int32 = 0
    fmt.Println(v6)
    var v92 int16 = v3
    var v93 float64 = 3.14159
    var v94 int32 = (v6 << ((4294967295) & 63))
    v95 := []int{}
    v96 := GenS6{F0: 883, F1: 42}
    var v97 bool = false
    v92 += v92
    v97 = v97
    v98 := GenS7{F0: 1, F1: 110, F2: v92}
    fmt.Println(v92, v6)
    if len(v95) >= 2 {
        v95 = v95[1:len(v95)]
    }
    for fi99 := 0; fi99 < 14; fi99++ {
        v93 = (-2.25 / (v93 + (float64(v5) / ((0.0 - v93) - (v93 * ((v93 / v93) + 3.14159))))))
        var v100 float32 = v4
        _ = fi99
        _ = v100
    }
    for fi101 := 0; fi101 < 18; fi101++ {
        v94 -= 0
        _ = len(v95)
        if v97 {
            if (((v97 && (23 >= 32767)) || v97) && v97) {
                fmt.Println(v94)
                if (genFn0(0, 2) >= genFn0(genFn1("", 42, 0), 326)) {
                    fmt.Println(v3, v94, v6)
                } else {
                    var v102 uint16 = 781
                    var v103 int8 = -64
                    v92 = int16(v102)
                    _ = v102
                    _ = v103
                }
                v4 = ((((v5 / v5) / v4) + (((1e10 + 1e10) - (float32(float64(v4)) + -2.25)) + v5)) - ((0.1 * (v5 * 0.1)) + v4))
            } else {
                var v104 int16 = 42
                var v105 int32 = (v94 >> ((255) & 63))
                fmt.Println(v94, v3, v3)
                _ = v104
                _ = v105
            }
            if (!(!(((v93 / float64((v5 / (v5 - (1e10 / v5))))) - float64(v4)) != -2.25))) {
                var v106 int32 = v94
                fmt.Println(v92)
                _ = v106
            }
        }
        _ = fi101
    }
    for fi107 := 0; fi107 < 19; fi107++ {
        v108 := GenS5{F0: (v3 < v3), F1: 7}
        if len(v95) > 0 {
            v95[0] = (^(fi107))
        }
        switch fi107 {
        case -2147483648:
            var v109 uint = 0
            _ = v109
        default:
            v110 := GenS8{F0: 4294967295}
            _ = v110
        }
        _ = fi107
        _ = v108
    }
    v6 -= (v94 / ((1) | 1))
    for fi111 := 0; fi111 < 20; fi111++ {
        if (0.1 != (v4 - 1.5)) {
            v112 := GenS9{F0: 1, F1: 272, F2: 2}
            _ = v112
        }
        for fi113 := 0; fi113 < 14; fi113++ {
            var v114 int16 = (v3 - 704)
            v3 += 32767
            if (!false) {
                v115 := GenS10{F0: "", F1: -1}
                v4 = ((v5 * ((float32(0.0) * 1.5) + 3.14159)) / 1.5)
                v95 = append(v95, fi113)
                v116 := GenS5{F0: ((((-1 == 42) && v97) && (42 >= 42)) || v97), F1: 100}
                _ = v115
                _ = v116
            }
            _ = fi113
            _ = v114
        }
        for fi117 := 0; fi117 < 18; fi117++ {
            if v97 {
                v97 = true
            }
            switch (fi111 * (fi111 / ((fi117) | 1))) {
            case 32768:
                fmt.Println(v97)
            case 32767:
                fmt.Println(v3, v94)
            case 756:
                var v118 bool = v97
                _ = v118
            default:
                if v97 {
                    if (!v97) {
                        var v119 int32 = (v6 + (-(v94)))
                        _ = v119
                    }
                    if (float32((((float64((v4 * v4)) / (0.1 / v93)) * (0.1 / v93)) - (v93 - float64(3.14159)))) > (v4 - v5)) {
                        var v120 uint8 = 28
                        _ = v120
                    }
                } else {
                    var v121 float32 = (3.14159 / v5)
                    switch fi117 {
                    case 1:
                        if v97 {
                            var v122 float32 = (((float32((0.1 * 1.5)) - ((v4 / (-2.25 - 3.14159)) - v4)) / -2.25) * 0.0)
                            delete(v2, ("\n\t" + ("\n\t" + "foo bar")))
                            _ = v122
                        } else {
                            v123 := GenS4{F0: 2}
                            _ = v123
                        }
                    case -347:
                        fmt.Println(v3, v4, v94)
                        switch "hello" {
                        case "\n\t":
                            var v124 float32 = v4
                            _ = v124
                        case "a":
                            switch fi117 {
                            case 127:
                                v95 = append(v95, (fi111 + (fi117 & (fi117 / ((32768) | 1)))))
                            case -2147483648:
                                v121 -= v121
                            }
                        default:
                            v97 = true
                            if true {
                                var v125 uint32 = 2147483647
                                v121 += 3.14159
                                v4 -= (v121 * v121)
                                _ = v125
                            }
                            fmt.Println(v4)
                            if (v4 == (float32(((v93 / v93) * (1e10 / v93))) - (v4 / (0.1 / v5)))) {
                                if v97 {
                                    if (!true) {
                                        v97 = (true && true)
                                        if false {
                                            var v126 int8 = -128
                                            _ = v126
                                        } else {
                                            v93 += ((v93 * (3.14159 - (-2.25 * ((v93 / 1.5) + v93)))) - v93)
                                            var v127 int8 = 1
                                            switch 2147483647 {
                                            case 721:
                                                var v128 int = (fi111 >> ((65535) & 63))
                                                _ = v128
                                            default:
                                                var v129 uint64 = 42
                                                _ = v129
                                            }
                                            v130 := GenS4{F0: 0}
                                            _ = v127
                                            _ = v130
                                        }
                                    } else {
                                        if (v97 || (v97 && (127 > 1))) {
                                            if (v3 != 32767) {
                                                if v97 {
                                                    v3 = v3
                                                    v92 = v3
                                                    var v131 int32 = (v6 + (v94 ^ (v94 + v6)))
                                                    _ = v131
                                                } else {
                                                    var v132 int = 32767
                                                    v92 -= int16(v6)
                                                    var v133 uint8 = 255
                                                    _ = v132
                                                    _ = v133
                                                }
                                            }
                                            if v97 {
                                                if v97 {
                                                    v5 = v5
                                                    if len(v95) > 0 {
                                                        v95[0] = -294
                                                    }
                                                    v5 = 3.14159
                                                    switch -37 {
                                                    case 0:
                                                        v121 -= v4
                                                    case 42:
                                                        v6 += (v94 * (v6 << ((827) & 63)))
                                                    case 4294967295:
                                                        v93 += v93
                                                    default:
                                                        var v134 string = "a"
                                                        _ = v134
                                                    }
                                                }
                                                fmt.Println(v94)
                                            } else {
                                                var v135 string = "a"
                                                var v136 int16 = v92
                                                _ = v135
                                                _ = v136
                                            }
                                        }
                                    }
                                    v92 += v3
                                    if true {
                                        switch (fi111 | (^(fi111))) {
                                        case 529:
                                            if (!(!v97)) {
                                                var v137 uint8 = 2
                                                _ = v137
                                            }
                                        }
                                        delete(v2, ("foo bar" + "a"))
                                    } else {
                                        fmt.Println(v3, v92, v94)
                                        var v138 uint32 = 880
                                        _ = v138
                                    }
                                }
                            } else {
                                v139 := GenS7{F0: (fi117 - fi117), F1: 125, F2: v92}
                                var v140 uint64 = 1
                                _ = v139
                                _ = v140
                            }
                        }
                        fmt.Println(v92)
                        switch "\n\t" {
                        case "hello":
                            v141 := GenS4{F0: 18446744073709551615}
                            _ = v141
                        }
                        v4 = v121
                    }
                    var v142 uint16 = 32767
                    _ = v121
                    _ = v142
                }
            }
            var v143 float32 = 1e10
            if len(v95) > 0 {
                v95[0] = fi111
            }
            _ = fi117
            _ = v143
        }
        _ = fi111
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    fmt.Println("slicelen v95:", len(v95), v95)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v92
    _ = v93
    _ = v94
    _ = v95
    _ = v96
    _ = v97
    _ = v98
}
