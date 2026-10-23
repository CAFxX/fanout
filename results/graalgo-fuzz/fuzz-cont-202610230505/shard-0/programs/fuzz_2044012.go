package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 int8 }
type GenS7 struct { F0 int16; F1 uint8 }
type GenS8 struct { F0 string; F1 uint32 }
type GenS9 struct { F0 int16; F1 int32; F2 int8 }
type GenS10 struct { F0 int64; F1 int64; F2 bool }
type GenS11 struct { F0 int32; F1 uint16 }
type GenS12 struct { F0 uint8; F1 int8 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 int16; F1 uint8; F2 string }

func main() {
    var v1 int8 = 100
    v2 := make(map[string]bool)
    v3 := []int{127, 42}
    delete(v2, ("" + ((("foo bar" + "\n\t") + "foo bar") + "")))
    v4 := GenS1{F0: 2, F1: 1, F2: 1}
    var v5 uint = 1
    var v6 int16 = 1000
    var v7 int64 = -9223372036854775808
    for fi8 := 0; fi8 < 11; fi8++ {
        switch -1 {
        case 340:
            v1 = v1
        case 758:
            var v9 int64 = 0
            _ = v9
        default:
            _ = len(v3)
        }
        switch "foo bar" {
        case "foo bar":
            if (1.5 > 3.14159) {
                if true {
                    v10 := GenS1{F0: v5, F1: 2, F2: v7}
                    _ = v10
                }
                var v11 int16 = (^(v6))
                _ = v11
            }
        }
        var v12 bool = (!false)
        _ = fi8
        _ = v12
    }
    var v13 float64 = float64(((0.1 - (0.0 - 1.5)) + 1.5))
    for fi14 := 0; fi14 < 11; fi14++ {
        for fi15 := 0; fi15 < 14; fi15++ {
            v1 = -128
            var v16 string = "\n\t"
            fmt.Println(v16)
            _ = fi15
            _ = v16
        }
        if true {
            fmt.Println(v13, v1)
        }
        _ = fi14
    }
    v17 := GenS1{F0: (v5 * v5), F1: 217, F2: v7}
    v18 := GenS1{F0: 65535, F1: 615, F2: (v7 + 42)}
    for fi19 := 0; fi19 < 6; fi19++ {
        fmt.Println(v5, v5, v13)
        var v20 uint64 = 18446744073709551615
        _ = fi19
        _ = v20
    }
    switch "foo bar" {
    case "\n\t":
        var v21 uint64 = 268
        _ = v21
    case "hello":
        var v22 uint8 = 2
        _ = v22
    default:
        if true {
            for fi23 := 0; fi23 < 14; fi23++ {
                _ = len(v2)
                _ = fi23
            }
            for fi24 := 0; fi24 < 15; fi24++ {
                for fi25 := 0; fi25 < 2; fi25++ {
                    switch ("" + "") {
                    case "\n\t":
                        if ((v1 < v1) && true) {
                            var v26 string = ("hello" + "hello")
                            _ = v26
                        }
                    }
                    if true {
                        delete(v2, "a")
                    } else {
                        var v27 uint64 = 968
                        var v28 int16 = v6
                        _ = v27
                        _ = v28
                    }
                    var v29 int64 = 0
                    var v30 int16 = v6
                    v13 = 1.5
                    var v31 bool = true
                    _ = fi25
                    _ = v29
                    _ = v30
                    _ = v31
                }
                for fi32 := 0; fi32 < 5; fi32++ {
                    if (false && ((-1 < (v1 | v1)) || (v5 != 1))) {
                        if len(v3) > 0 {
                            v3[0] = fi24
                        }
                        switch (("a" + "hello") + ("a" + ("foo bar" + ("a" + "a")))) {
                        case "foo bar":
                            fmt.Println(fi32, fi24)
                        }
                        switch "a" {
                        case "":
                            var v33 bool = (((-2.25 < (1.5 + (-2.25 * 1.5))) && false) && (2147483647 == -381))
                            if false {
                                switch "foo bar" {
                                case "":
                                    if (v33 || ((fi24 * fi24) == (fi24 >> ((255) & 63)))) {
                                        v34 := GenS1{F0: v5, F1: 2, F2: v7}
                                        fmt.Println(fi24, v6, fi32)
                                        _ = v34
                                    } else {
                                        var v35 int8 = 1
                                        _ = v35
                                    }
                                case "foo bar":
                                    v33 = v33
                                case "hello":
                                    var v36 uint8 = 128
                                    _ = v36
                                default:
                                    if v33 {
                                        v13 -= 0.0
                                        var v37 int = (fi24 << (((^(v5))) & 63))
                                        v38 := GenS1{F0: 4294967295, F1: 2, F2: v7}
                                        v39 := GenS1{F0: 18446744073709551615, F1: 916, F2: 9223372036854775807}
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                    } else {
                                        var v40 float32 = 0.0
                                        v41 := GenS1{F0: 817, F1: 255, F2: (v7 - (v7 - int64(v6)))}
                                        _ = v40
                                        _ = v41
                                    }
                                    v7 = v7
                                    switch (fi24 % ((fi32) | 1)) {
                                    case 0:
                                        var v42 uint8 = 2
                                        _ = v42
                                    case -2147483648:
                                        if v33 {
                                            v13 = float64(1.5)
                                            _ = len(v2)
                                            if true {
                                                switch 32767 {
                                                case -612:
                                                    if (18446744073709551615 >= 4294967295) {
                                                        v1 += 127
                                                        fmt.Println(fi24, v1)
                                                    }
                                                case -2147483648:
                                                    switch "" {
                                                    case "a":
                                                        v3 = append(v3, 32768)
                                                    default:
                                                        if (((("a" + "") + "foo bar") + "a") <= "hello") {
                                                            if ((v13 * float64(0.1)) < v13) {
                                                                v43 := GenS1{F0: v5, F1: 1, F2: -623}
                                                                v33 = true
                                                                _ = v43
                                                            } else {
                                                                v7 = (v7 / ((962) | 1))
                                                                v44 := GenS2{F0: (v1 & (^(v1)))}
                                                                _ = v44
                                                            }
                                                        }
                                                    }
                                                case 42:
                                                    v45 := v2["a"]
                                                    _ = v45
                                                }
                                                switch (fi32 & fi32) {
                                                case 522:
                                                    switch ("\n\t" + ("" + "\n\t")) {
                                                    case "a":
                                                        if v33 {
                                                            var v46 int32 = 116
                                                            _ = v46
                                                        }
                                                    case "hello":
                                                        if ((150 < 1) || true) {
                                                            v3 = append(v3, (fi32 >> (((v5 - v5)) & 63)))
                                                            var v47 string = "a"
                                                            v7 += v7
                                                            _ = v47
                                                        } else {
                                                            var v48 uint64 = 2
                                                            v1 = (v1 << (((v5 ^ v5)) & 63))
                                                            v5 = 255
                                                            var v49 float64 = ((0.1 + v13) / v13)
                                                            _ = v48
                                                            _ = v49
                                                        }
                                                    case "foo bar":
                                                        if false {
                                                            var v50 int32 = int32(v7)
                                                            v51 := v2[""]
                                                            _ = v50
                                                            _ = v51
                                                        }
                                                    default:
                                                        fmt.Println(v5, v6, v13)
                                                    }
                                                }
                                                if v33 {
                                                    _ = len(v3)
                                                    var v52 uint16 = 255
                                                    v2[""] = v33
                                                    var v53 uint64 = 1
                                                    switch (fi32 - fi24) {
                                                    case 565:
                                                        v33 = (-1 > 42)
                                                    case 437:
                                                        fmt.Println(v7, v13, v52)
                                                        var v54 int8 = v1
                                                        _ = v54
                                                    }
                                                    v2[("hello" + "\n\t")] = false
                                                    _ = v52
                                                    _ = v53
                                                } else {
                                                    if ((v5 + v5) > (v5 + v5)) {
                                                        var v55 int = int(v7)
                                                        var v56 int64 = (v7 % ((v7) | 1))
                                                        var v57 string = "foo bar"
                                                        _ = v55
                                                        _ = v56
                                                        _ = v57
                                                    } else {
                                                        fmt.Println(v13, v7, fi24)
                                                    }
                                                }
                                                fmt.Println(v6)
                                                v58 := GenS3{F0: v6, F1: 127}
                                                switch fi32 {
                                                case -1:
                                                    v1 = 1
                                                default:
                                                    v13 = 0.0
                                                }
                                                _ = v58
                                            } else {
                                                v13 = (float64(-2.25) / v13)
                                                v13 += 0.0
                                                var v59 float32 = 1.5
                                                v7 = v7
                                                var v60 uint = 42
                                                _ = v59
                                                _ = v60
                                            }
                                        }
                                    case 128:
                                        var v61 int32 = -2147483648
                                        _ = v61
                                    }
                                    v7 = (^(v7))
                                    var v62 int8 = 100
                                    v63 := GenS2{F0: 42}
                                    fmt.Println(v62, fi24)
                                    var v64 int = int(v6)
                                    var v65 uint = 505
                                    fmt.Println(v64, v13, v6)
                                    _ = v62
                                    _ = v63
                                    _ = v64
                                    _ = v65
                                }
                                switch fi24 {
                                case 599:
                                    v5 -= v5
                                case 0:
                                    v66 := GenS4{F0: int(v7)}
                                    _ = v66
                                case -1:
                                    if (!(255 == 42)) {
                                        fmt.Println(v13)
                                        var v67 uint32 = 2147483648
                                        v68 := GenS2{F0: (-(v1))}
                                        v69 := GenS5{F0: (^(fi32))}
                                        _ = v67
                                        _ = v68
                                        _ = v69
                                    } else {
                                        v70 := GenS6{F0: 18446744073709551615, F1: 48}
                                        v13 -= 0.1
                                        v13 += (float64((float32((((0.0 / ((float64(1.5) / v13) * 0.1)) * (v13 / 1.0)) * 3.14159)) / 1.5)) * v13)
                                        var v71 int64 = v7
                                        _ = v70
                                        _ = v71
                                    }
                                default:
                                    v72 := GenS7{F0: v6, F1: 128}
                                    _ = v72
                                }
                                if (1.5 <= 1.5) {
                                    v5 += 1
                                    if true {
                                        switch (fi24 * int(v6)) {
                                        case 935:
                                            v1 += (v1 | v1)
                                        }
                                        v73 := GenS8{F0: "a", F1: 2147483648}
                                        _ = v73
                                    } else {
                                        v5 += v5
                                        if true {
                                            fmt.Println(v5, v13, v7)
                                            v74 := GenS4{F0: (fi24 ^ (^(fi24)))}
                                            _ = v74
                                        }
                                    }
                                    v13 = v13
                                } else {
                                    delete(v2, "foo bar")
                                }
                            }
                            _ = v33
                        default:
                            if ((1e10 + 3.14159) < 0.1) {
                                if (true && false) {
                                    if (66 < 35) {
                                        v75 := GenS2{F0: v1}
                                        if true {
                                            v6 = (v6 + v6)
                                            v76 := GenS3{F0: v6, F1: 128}
                                            var v77 float32 = 0.1
                                            _ = v76
                                            _ = v77
                                        } else {
                                            v13 = (1.5 - 0.0)
                                            var v78 uint32 = 396
                                            _ = v78
                                        }
                                        var v79 int64 = 1
                                        _ = v75
                                        _ = v79
                                    } else {
                                        v6 -= 1
                                        switch "" {
                                        case "hello":
                                            if true {
                                                if true {
                                                    v80 := GenS1{F0: v5, F1: 791, F2: v7}
                                                    var v81 int = (fi24 << ((65535) & 63))
                                                    _ = v80
                                                    _ = v81
                                                } else {
                                                    v6 -= v6
                                                    var v82 int32 = int32(fi24)
                                                    _ = v82
                                                }
                                                v6 += v6
                                                v83 := v2[(("a" + ("foo bar" + ("\n\t" + ""))) + "")]
                                                _ = v83
                                            } else {
                                                v84 := GenS9{F0: v6, F1: 182, F2: 124}
                                                if (!(127 == uint8(v5))) {
                                                    v6 += 0
                                                } else {
                                                    if false {
                                                        if len(v3) >= 2 {
                                                            v3 = v3[1:len(v3)]
                                                        }
                                                        if true {
                                                            var v85 bool = false
                                                            var v86 uint16 = 641
                                                            _ = v85
                                                            _ = v86
                                                        }
                                                        v87 := GenS10{F0: 0, F1: v7, F2: ((((0.0 + ((v13 + (((v13 + (v13 + v13)) + float64(((((float32(v13) * -2.25) + ((float32((1.5 - 0.0)) + 1e10) / float32((((v13 * 0.0) / float64(1e10)) / (v13 * (0.0 + (((1.5 * float64(0.1)) - (((v13 * (0.1 - (0.0 * v13))) / ((v13 - v13) / float64(1.5))) - float64((float32(v13) / (1e10 + ((float32(float64(float32((((0.0 + (3.14159 - -2.25)) * float64((((-2.25 / 0.1) + 0.0) * 0.1))) / 1.0)))) / 3.14159) / float32(float64(((((1e10 - 3.14159) + (1.5 - -2.25)) - float32(float64(0.1))) + (1e10 * float32(0.1))))))))))) + (3.14159 - v13)))))))) / (1e10 - -2.25)) + 0.1))) - -2.25)) - v13)) / v13) + v13) >= v13)}
                                                        _ = v87
                                                    } else {
                                                        fmt.Println(fi32)
                                                        fmt.Println(v7, fi24)
                                                    }
                                                    v88 := GenS2{F0: v1}
                                                    _ = v88
                                                }
                                                switch "\n\t" {
                                                case "hello":
                                                    v89 := GenS6{F0: 178, F1: 40}
                                                    _ = v89
                                                case "":
                                                    v13 -= -2.25
                                                case "a":
                                                    v5 -= (v5 * v5)
                                                }
                                                if false {
                                                    var v90 uint16 = 65535
                                                    _ = v90
                                                }
                                                v3 = append(v3, (fi32 & (-(fi32))))
                                                var v91 int64 = v7
                                                _ = v84
                                                _ = v91
                                            }
                                        case "a":
                                            v5 -= 65535
                                        }
                                        var v92 int16 = (-(v6))
                                        delete(v2, ("foo bar" + ("" + "foo bar")))
                                        v92 += (v92 * (v92 ^ v92))
                                        if (false && (!(999 <= 274))) {
                                            v6 += 42
                                            var v93 float64 = (float64(float32(v13)) + v13)
                                            if true {
                                                var v94 float32 = ((0.0 * 3.14159) / 1.5)
                                                v94 = ((v94 + v94) / float32(v93))
                                                _ = v94
                                            }
                                            _ = v93
                                        } else {
                                            if (!(!(!(!true)))) {
                                                v95 := GenS1{F0: 0, F1: 0, F2: v7}
                                                var v96 uint8 = 38
                                                fmt.Println(v13, v6, fi32)
                                                _ = v95
                                                _ = v96
                                            }
                                            delete(v2, ((((("foo bar" + ("" + ("" + "hello"))) + "\n\t") + "a") + "hello") + ("foo bar" + "hello")))
                                        }
                                        _ = v92
                                    }
                                    v5 -= v5
                                    if (!(true || (v6 != (v6 << (((v5 | v5)) & 63))))) {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    } else {
                                        v3 = append(v3, (^(fi24)))
                                    }
                                    if false {
                                        v97 := GenS9{F0: v6, F1: -2147483648, F2: v1}
                                        var v98 int = int(v7)
                                        var v99 int32 = 1
                                        _ = v97
                                        _ = v98
                                        _ = v99
                                    }
                                }
                                switch "foo bar" {
                                case "a":
                                    v7 = 791
                                case "hello":
                                    var v100 uint64 = 2
                                    _ = v100
                                }
                                switch fi32 {
                                case 42:
                                    if ((v5 - (v5 / ((v5) | 1))) >= (^(v5))) {
                                        var v101 int = (fi24 / ((fi32) | 1))
                                        if (((0.0 - ((v13 + float64(1.5)) / v13)) <= v13) || (!false)) {
                                            switch 1 {
                                            case -1:
                                                if false {
                                                    var v102 uint8 = 128
                                                    _ = v102
                                                }
                                            case -9223372036854775808:
                                                v7 = 42
                                            case 2147483647:
                                                var v103 float64 = (v13 + v13)
                                                _ = v103
                                            }
                                            if false {
                                                v104 := GenS10{F0: v7, F1: 0, F2: (-354 >= (v6 << (((v5 - v5)) & 63)))}
                                                fmt.Println(v5, fi24, fi24)
                                                var v105 int8 = (v1 + (v1 ^ (v1 + (-(v1)))))
                                                var v106 bool = (fi32 <= fi32)
                                                var v107 int64 = v7
                                                _ = v104
                                                _ = v105
                                                _ = v106
                                                _ = v107
                                            } else {
                                                v6 -= (v6 - 0)
                                            }
                                            v6 = v6
                                            v108 := GenS11{F0: 940, F1: 32767}
                                            _ = v108
                                        } else {
                                            v5 -= 42
                                            var v109 uint32 = 2147483648
                                            switch 297 {
                                            case 245:
                                                var v110 uint16 = 194
                                                _ = v110
                                            }
                                            v1 = v1
                                            _ = v109
                                        }
                                        _ = v101
                                    } else {
                                        v13 -= ((v13 - v13) - ((v13 * float64(((0.0 + (0.1 / 1.5)) / 1.0))) - (v13 + v13)))
                                        v111 := GenS4{F0: fi24}
                                        if true {
                                            v3 = append(v3, (^(fi32)))
                                            var v112 uint8 = 0
                                            _ = v112
                                        } else {
                                            v7 = v7
                                            var v113 string = "a"
                                            var v114 int8 = v1
                                            if true {
                                                var v115 int16 = 0
                                                v116 := GenS12{F0: 112, F1: v114}
                                                if (!(-1 >= -2147483648)) {
                                                    if true {
                                                        if (!true) {
                                                            var v117 float64 = v13
                                                            v117 -= v13
                                                            _ = v117
                                                        }
                                                        var v118 int64 = v7
                                                        v119 := GenS6{F0: 265, F1: v114}
                                                        _ = v118
                                                        _ = v119
                                                    }
                                                    if true {
                                                        switch 128 {
                                                        case 756:
                                                            v1 += v1
                                                        }
                                                        var v120 uint16 = 42
                                                        if false {
                                                            if false {
                                                                var v121 int64 = v7
                                                                _ = v121
                                                            } else {
                                                                fmt.Println(v5, v7)
                                                            }
                                                            fmt.Println(v13, v5)
                                                            v122 := GenS13{F0: v120}
                                                            _ = v122
                                                        } else {
                                                            if ((192 == 255) && ((v1 | v114) >= v114)) {
                                                                if false {
                                                                    v123 := GenS2{F0: v1}
                                                                    fmt.Println(v113)
                                                                    _ = v123
                                                                } else {
                                                                    var v124 uint8 = 255
                                                                    v5 -= 18446744073709551615
                                                                    _ = v124
                                                                }
                                                                if true {
                                                                    if false {
                                                                        if (65535 < v120) {
                                                                            if false {
                                                                                v5 = (v5 & 588)
                                                                            } else {
                                                                                switch fi24 {
                                                                                case -867:
                                                                                    v114 -= 42
                                                                                default:
                                                                                    delete(v2, ((v113 + v113) + v113))
                                                                                    var v125 uint8 = 255
                                                                                    fmt.Println(v125)
                                                                                    var v126 int16 = (v6 | (^(v6)))
                                                                                    _ = v125
                                                                                    _ = v126
                                                                                }
                                                                                v127 := GenS14{F0: 35}
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                var v128 float64 = ((float64((1e10 - (((float32(v13) + 0.1) * (0.0 - (1.5 - 1.5))) + float32(v13)))) + (v13 / 3.14159)) * v13)
                                                                                _ = len(v2)
                                                                                _ = v127
                                                                                _ = v128
                                                                            }
                                                                            fmt.Println(v6)
                                                                            v3 = append(v3, (fi24 >> ((uint(v120)) & 63)))
                                                                            var v129 uint32 = 2
                                                                            _ = v129
                                                                        }
                                                                        var v130 uint64 = 18446744073709551615
                                                                        _ = v130
                                                                    } else {
                                                                        v114 += (v114 - v1)
                                                                    }
                                                                    switch (v113 + "foo bar") {
                                                                    case "":
                                                                        var v131 int64 = 161
                                                                        _ = v131
                                                                    default:
                                                                        v2[v113] = true
                                                                    }
                                                                }
                                                            }
                                                            v132 := GenS6{F0: 4294967295, F1: (v114 / ((0) | 1))}
                                                            v1 -= (v114 << (((v5 | 103)) & 63))
                                                            v133 := GenS14{F0: 42}
                                                            _ = v132
                                                            _ = v133
                                                        }
                                                        fmt.Println(v115)
                                                        switch (fi24 << ((2) & 63)) {
                                                        case -144:
                                                            if (((v120 & v120) < (v120 % ((v120) | 1))) && (!false)) {
                                                                if ((!(false || ((fi32 == fi32) && ((v13 < (((v13 - v13) - (0.0 / (float64(1.5) - v13))) * v13)) || false)))) && (!(true && true))) {
                                                                    if (true || true) {
                                                                        switch (fi32 >> (((v5 / (((v5 % ((v5) | 1))) | 1))) & 63)) {
                                                                        case 9223372036854775807:
                                                                            v7 = v7
                                                                        case -4:
                                                                            v134 := v2["a"]
                                                                            _ = v134
                                                                        case 128:
                                                                            if false {
                                                                                v135 := v2[(((v113 + v113) + (v113 + "\n\t")) + "foo bar")]
                                                                                v2["foo bar"] = true
                                                                                _ = v135
                                                                            }
                                                                        default:
                                                                            v1 = (v114 | (v114 + v114))
                                                                        }
                                                                        var v136 int8 = (v1 / ((42) | 1))
                                                                        var v137 int64 = 0
                                                                        _ = v136
                                                                        _ = v137
                                                                    }
                                                                    var v138 float32 = 1.5
                                                                    v6 = 506
                                                                    _ = v138
                                                                } else {
                                                                    switch v113 {
                                                                    case "foo bar":
                                                                        var v139 uint8 = 24
                                                                        _ = v139
                                                                    case "":
                                                                        var v140 int32 = 2147483647
                                                                        _ = v140
                                                                    case "\n\t":
                                                                        fmt.Println(v7)
                                                                    default:
                                                                        var v141 bool = true
                                                                        _ = v141
                                                                    }
                                                                    fmt.Println(v5, v7, v5)
                                                                    if (uint32(v120) <= 420) {
                                                                        var v142 bool = true
                                                                        _ = v142
                                                                    }
                                                                }
                                                                var v143 int32 = 0
                                                                var v144 uint32 = 1
                                                                _ = v143
                                                                _ = v144
                                                            }
                                                        default:
                                                            if false {
                                                                fmt.Println(fi32, v7)
                                                            } else {
                                                                var v145 uint64 = uint64(v5)
                                                                v146 := GenS13{F0: 36}
                                                                if (false || (v5 >= (^(v5)))) {
                                                                    fmt.Println(fi24)
                                                                    v147 := v2[v113]
                                                                    if (301 <= 2147483648) {
                                                                        v2[("\n\t" + v113)] = false
                                                                        v5 -= 4294967295
                                                                    }
                                                                    _ = v147
                                                                } else {
                                                                    switch v113 {
                                                                    case "foo bar":
                                                                        _ = len(v2)
                                                                    case "a":
                                                                        v148 := GenS13{F0: v120}
                                                                        _ = v148
                                                                    }
                                                                }
                                                                _ = v145
                                                                _ = v146
                                                            }
                                                        }
                                                        if true {
                                                            var v149 uint32 = 554
                                                            v150 := GenS2{F0: -128}
                                                            _ = v149
                                                            _ = v150
                                                        } else {
                                                            var v151 int32 = -2147483648
                                                            _ = v151
                                                        }
                                                        v2[(v113 + v113)] = (true || (!true))
                                                        _ = v120
                                                    } else {
                                                        if (!(v13 <= 1e10)) {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            if (328 > 616) {
                                                                if (false && false) {
                                                                    if (!false) {
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        v2[("foo bar" + ("hello" + ("a" + v113)))] = ((1e10 * 0.0) > 1.5)
                                                                        fmt.Println(v7, v115, v13)
                                                                        var v152 float64 = (v13 + v13)
                                                                        var v153 int16 = -185
                                                                        _ = v152
                                                                        _ = v153
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        v154 := GenS10{F0: (v7 * int64(v5)), F1: (-(v7)), F2: false}
                                                                        fmt.Println(v115, v6, fi32)
                                                                        v155 := GenS12{F0: 128, F1: (^(v1))}
                                                                        delete(v2, v113)
                                                                        var v156 float64 = v13
                                                                        _ = v154
                                                                        _ = v155
                                                                        _ = v156
                                                                    }
                                                                }
                                                                v157 := GenS7{F0: v115, F1: 42}
                                                                v7 = -1
                                                                _ = v157
                                                            } else {
                                                                fmt.Println(v5, v115)
                                                                v7 += (v7 ^ v7)
                                                                if len(v3) > 0 {
                                                                    v3[0] = fi24
                                                                }
                                                            }
                                                        } else {
                                                            fmt.Println(v113, v13)
                                                        }
                                                        v115 += v115
                                                    }
                                                    delete(v2, v113)
                                                } else {
                                                    v158 := GenS3{F0: v6, F1: 2}
                                                    _ = v158
                                                }
                                                _ = v115
                                                _ = v116
                                            }
                                            if true {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                if ((v113 + (((((v113 + v113) + (v113 + "\n\t")) + v113) + (v113 + "\n\t")) + v113)) < ((v113 + "\n\t") + "a")) {
                                                    if (false && (!(v7 != 42))) {
                                                        v13 = ((v13 + float64(-2.25)) * -2.25)
                                                        v159 := GenS12{F0: 255, F1: v114}
                                                        _ = v159
                                                    }
                                                    fmt.Println(v1, v114, v113)
                                                }
                                                _ = len(v3)
                                            } else {
                                                var v160 int8 = v114
                                                _ = v160
                                            }
                                            _ = v113
                                            _ = v114
                                        }
                                        _ = v111
                                    }
                                case 32767:
                                    v7 += 588
                                case 128:
                                    v161 := GenS4{F0: 927}
                                    _ = v161
                                }
                                var v162 uint64 = 2
                                var v163 float32 = 0.1
                                var v164 int8 = (v1 & 127)
                                fmt.Println(v5)
                                _ = v162
                                _ = v163
                                _ = v164
                            }
                        }
                        fmt.Println(v1)
                        v165 := GenS10{F0: (v7 * (v7 + v7)), F1: -9223372036854775808, F2: false}
                        v166 := GenS15{F0: v6, F1: 128, F2: "foo bar"}
                        v7 += 1
                        _ = v165
                        _ = v166
                    }
                    var v167 uint64 = 4294967295
                    v168 := GenS7{F0: -32768, F1: 1}
                    _ = fi32
                    _ = v167
                    _ = v168
                }
                _ = fi24
            }
            switch "\n\t" {
            case "":
                for fi169 := 0; fi169 < 18; fi169++ {
                    var v170 int32 = 0
                    _ = fi169
                    _ = v170
                }
            case "foo bar":
                for fi171 := 0; fi171 < 4; fi171++ {
                    for fi172 := 0; fi172 < 17; fi172++ {
                        v2["hello"] = (!false)
                        var v173 int8 = v1
                        _ = fi172
                        _ = v173
                    }
                    _ = fi171
                }
            case "a":
                v5 -= v5
            }
        }
    }
    for fi174 := 0; fi174 < 19; fi174++ {
        v175 := v2["a"]
        delete(v2, ("hello" + ("\n\t" + (("hello" + (("foo bar" + (((("" + "a") + ("hello" + "")) + ("" + "hello")) + "hello")) + "hello")) + ""))))
        _ = fi174
        _ = v175
    }
    var v176 uint8 = 127
    if len(v3) > 0 {
        v3[0] = 9223372036854775807
    }
    v176 = v176
    v3 = append(v3, 42)
    v177 := GenS11{F0: 2147483647, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v176:", v176)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v17
    _ = v176
    _ = v177
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
