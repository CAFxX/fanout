package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int64; F1 bool; F2 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32; F1 int8 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 int8; F1 uint8 }
type GenS11 struct { F0 uint; F1 int8 }
type GenS12 struct { F0 int64; F1 string; F2 int8 }
type GenS13 struct { F0 int8; F1 uint; F2 uint8 }
type GenS14 struct { F0 string; F1 string; F2 string }
type GenS15 struct { F0 uint; F1 uint8 }
type GenS16 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS17 struct { F0 string; F1 int32 }
type GenS18 struct { F0 int64; F1 uint; F2 uint32 }
type GenS19 struct { F0 uint16 }
type GenS20 struct { F0 int16 }
type GenS21 struct { F0 uint64 }
type GenS22 struct { F0 uint }
type GenS23 struct { F0 int; F1 int64; F2 int }
func genFn0(p0 uint64) string {
    p0 = (p0 / ((p0) | 1))
    p0 = (p0 % (((p0 - 191)) | 1))
    _ = p0
    return "\n\t"
}
func genFn1(p0 int32) uint32 {
    switch 127 {
    case 42:
        v7 := []int{-1, 9223372036854775807, 795}
        _ = v7
    case 1:
        p0 += 2147483647
    case 61:
        v8 := []int{0, 700, -2147483648, 55}
        _ = v8
    default:
        var v9 uint64 = 870
        _ = v9
    }
    v10 := []int{128, 32767, 504}
    var v11 int8 = 0
    _ = p0
    _ = v10
    _ = v11
    return 2147483647
}

func main() {
    var v1 uint = 241
    var v2 string = (("\n\t" + "a") + ("foo bar" + ("\n\t" + "a")))
    v3 := []string{"a", ((v2 + (v2 + ("hello" + v2))) + (("" + ("hello" + v2)) + "hello"))}
    v4 := make(map[uint32]bool)
    var v5 uint64 = 710
    v6 := make(map[int]bool)
    v1 -= (v1 & (v1 ^ v1))
    v12 := GenS1{F0: v5, F1: 0}
    for fi13 := 0; fi13 < 2; fi13++ {
        v14 := GenS2{F0: 834, F1: ((-2.25 + 1.5) >= 1.5), F2: (!(0 != 2))}
        if false {
            v15 := GenS1{F0: v5, F1: 843}
            v1 += 0
            _ = v15
        } else {
            fmt.Println(fi13, v1, v1)
        }
        v5 += (v5 % ((42) | 1))
        _ = fi13
        _ = v14
    }
    v2 = v2
    for fi16 := 0; fi16 < 13; fi16++ {
        fmt.Println(v5)
        if len(v3) > 0 {
            v3[0] = "\n\t"
        }
        _ = fi16
    }
    fmt.Println(v5, v1)
    fmt.Println(v5)
    v5 = 966
    for fi17 := 0; fi17 < 1; fi17++ {
        v18 := v4[2147483648]
        v19 := v6[32767]
        for fi20 := 0; fi20 < 13; fi20++ {
            var v21 float32 = float32(-2.25)
            _ = fi20
            _ = v21
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    var v22 float32 = 1e10
    switch ((v2 + ("a" + "")) + v2) {
    case "a":
        v23 := GenS3{F0: false}
        _ = v23
    default:
        for fi24 := 0; fi24 < 17; fi24++ {
            fmt.Println(v22)
            for fi25 := 0; fi25 < 4; fi25++ {
                v22 = 0.0
                _ = fi25
            }
            _ = fi24
        }
    }
    if false {
        var v26 int64 = -761
        var v27 uint16 = 435
        switch 4294967295 {
        case -260:
            switch (((((v2 + v2) + v2) + v2) + (v2 + "")) + ("a" + (v2 + v2))) {
            case "":
                delete(v6, 2147483647)
            }
        case 9223372036854775807:
            if true {
                var v28 uint = 4294967295
                fmt.Println(v28, v2)
                for fi29 := 0; fi29 < 12; fi29++ {
                    _ = len(v3)
                    _ = len(v6)
                    _ = fi29
                }
                _ = v28
            } else {
                v30 := GenS4{F0: (v1 + 622)}
                v31 := v4[211]
                _ = v30
                _ = v31
            }
        case -973:
            if true {
                switch -305 {
                case -2147483648:
                    v22 -= float32((1e10 * ((-2.25 * -2.25) / 1.0)))
                case 9223372036854775807:
                    var v32 bool = false
                    _ = v32
                case 1:
                    fmt.Println(v5, v5, v27)
                }
            }
        }
        _ = v26
        _ = v27
    }
    if false {
        v1 = v1
        if true {
            var v33 string = v2
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v33
        } else {
            if len(v3) > 0 {
                v3[0] = ""
            }
            if (1 == uint32(v5)) {
                var v34 int32 = -378
                v35 := v6[4294967295]
                _ = v34
                _ = v35
            } else {
                for fi36 := 0; fi36 < 15; fi36++ {
                    var v37 bool = true
                    _ = fi36
                    _ = v37
                }
                switch v2 {
                case "foo bar":
                    var v38 int64 = 1
                    _ = v38
                case "":
                    var v39 string = ((v2 + v2) + ("foo bar" + ""))
                    _ = v39
                }
                if false {
                    var v40 int8 = -128
                    if false {
                        for fi41 := 0; fi41 < 10; fi41++ {
                            for fi42 := 0; fi42 < 12; fi42++ {
                                fmt.Println(v2)
                                v2 = v2
                                _ = fi42
                            }
                            if (-67 != 2147483647) {
                                v43 := GenS2{F0: -1, F1: true, F2: (0.0 == -2.25)}
                                var v44 uint64 = v5
                                _ = v43
                                _ = v44
                            }
                            _ = fi41
                        }
                    }
                    if false {
                        if false {
                            var v45 uint = v1
                            _ = v45
                        }
                        var v46 int32 = -1
                        var v47 uint = 315
                        _ = v46
                        _ = v47
                    } else {
                        v5 += v5
                        if false {
                            v22 = float32(0.0)
                            for fi48 := 0; fi48 < 15; fi48++ {
                                switch v2 {
                                case "a":
                                    for fi49 := 0; fi49 < 10; fi49++ {
                                        v40 = v40
                                        _ = fi49
                                    }
                                case "hello":
                                    for fi50 := 0; fi50 < 12; fi50++ {
                                        var v51 uint16 = 2
                                        v5 -= 17
                                        var v52 int8 = (v40 + -1)
                                        v52 = v52
                                        switch 354 {
                                        case 9223372036854775807:
                                            v22 += ((v22 / 0.1) - (float32(1e10) / ((float32(0.0) - v22) + (((v22 / (v22 + 0.0)) - v22) * 0.1))))
                                        case 4294967295:
                                            v40 += (v40 >> ((225) & 63))
                                        }
                                        v53 := GenS4{F0: 237}
                                        v1 += v1
                                        _ = fi50
                                        _ = v51
                                        _ = v52
                                        _ = v53
                                    }
                                default:
                                    if (false || false) {
                                        fmt.Println(v22, fi48, v40)
                                    } else {
                                        for fi54 := 0; fi54 < 10; fi54++ {
                                            v1 += 2
                                            _ = fi54
                                        }
                                        if true {
                                            if (v2 <= "") {
                                                _ = len(v4)
                                                if len(v3) > 0 {
                                                    v3[0] = (((v2 + v2) + v2) + (v2 + ("foo bar" + v2)))
                                                }
                                                var v55 float32 = 3.14159
                                                _ = v55
                                            } else {
                                                fmt.Println(v40, v1, v1)
                                            }
                                            v5 = 0
                                        } else {
                                            for fi56 := 0; fi56 < 5; fi56++ {
                                                var v57 int = 128
                                                if true {
                                                    v22 = v22
                                                    if ((326 <= uint16(v5)) && ("" <= (("foo bar" + v2) + v2))) {
                                                        if false {
                                                            fmt.Println(v40, v22, v1)
                                                            var v58 float64 = (1.5 / (1e10 - ((1e10 / (3.14159 + (3.14159 + 1e10))) - 1.5)))
                                                            v4[2] = true
                                                            _ = v58
                                                        } else {
                                                            switch v2 {
                                                            case "hello":
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                            case "":
                                                                var v59 int32 = -728
                                                                _ = v59
                                                            default:
                                                                if len(v3) > 0 {
                                                                    v3[0] = v2
                                                                }
                                                            }
                                                            v57 = (fi56 + fi48)
                                                            v57 = -976
                                                        }
                                                        if (false || true) {
                                                            switch (fi56 ^ (v57 >> ((0) & 63))) {
                                                            case -643:
                                                                v60 := GenS5{F0: 17}
                                                                _ = v60
                                                            }
                                                            switch -750 {
                                                            case 42:
                                                                v2 = ""
                                                            default:
                                                                var v61 uint8 = 255
                                                                _ = v61
                                                            }
                                                            var v62 uint64 = 903
                                                            _ = v62
                                                        } else {
                                                            if (!(255 <= 255)) {
                                                                switch v2 {
                                                                case "\n\t":
                                                                    switch fi56 {
                                                                    case -72:
                                                                        v63 := GenS4{F0: (v1 % ((v1) | 1))}
                                                                        _ = v63
                                                                    case 939:
                                                                        v22 -= v22
                                                                    }
                                                                case "foo bar":
                                                                    v64 := GenS6{F0: 42}
                                                                    _ = v64
                                                                case "hello":
                                                                    var v65 int64 = 195
                                                                    v66 := GenS5{F0: (fi48 * v57)}
                                                                    _ = v65
                                                                    _ = v66
                                                                }
                                                                v67 := GenS5{F0: -2147483648}
                                                                _ = v67
                                                            }
                                                            v68 := GenS3{F0: (v40 > (v40 + (v40 | (v40 * (-(v40))))))}
                                                            fmt.Println(v57, v1, v22)
                                                            var v69 uint16 = 1
                                                            _ = v68
                                                            _ = v69
                                                        }
                                                        if false {
                                                            fmt.Println(v5, fi56, v1)
                                                            if (!true) {
                                                                if true {
                                                                    var v70 float32 = (float32(1.5) * v22)
                                                                    v71 := GenS5{F0: int(v5)}
                                                                    _ = v70
                                                                    _ = v71
                                                                } else {
                                                                    var v72 int16 = 1
                                                                    var v73 int = (v57 ^ v57)
                                                                    _ = v72
                                                                    _ = v73
                                                                }
                                                            }
                                                            v74 := GenS7{F0: -2147483648, F1: 127}
                                                            _ = v74
                                                        } else {
                                                            if (false || (842 != 2)) {
                                                                fmt.Println(v2, v57, fi48)
                                                                v3 = append(v3, (v2 + v2))
                                                            } else {
                                                                fmt.Println(v40, v40, fi56)
                                                            }
                                                            v3 = append(v3, (("a" + v2) + (("" + "\n\t") + v2)))
                                                        }
                                                    } else {
                                                        v4[genFn1(42)] = (!true)
                                                        v3 = append(v3, ("" + "foo bar"))
                                                        _ = len(v4)
                                                    }
                                                }
                                                switch (^(fi56)) {
                                                case 490:
                                                    if true {
                                                        fmt.Println(v2, fi48, v5)
                                                    }
                                                case 4294967295:
                                                    var v75 int32 = 389
                                                    _ = v75
                                                default:
                                                    var v76 uint64 = v5
                                                    _ = v76
                                                }
                                                _ = fi56
                                                _ = v57
                                            }
                                            for fi77 := 0; fi77 < 15; fi77++ {
                                                var v78 int = -320
                                                v5 = (^(v5))
                                                v40 -= v40
                                                _ = fi77
                                                _ = v78
                                            }
                                        }
                                    }
                                }
                                for fi79 := 0; fi79 < 19; fi79++ {
                                    if false {
                                        v22 = v22
                                        if false {
                                            switch fi48 {
                                            case -314:
                                                var v80 float64 = (3.14159 * 1e10)
                                                _ = v80
                                            case 4294967295:
                                                fmt.Println(v1)
                                            case 9223372036854775807:
                                                var v81 int = 9223372036854775807
                                                _ = v81
                                            }
                                            v82 := GenS4{F0: v1}
                                            v5 = (v5 % ((v5) | 1))
                                            _ = v82
                                        } else {
                                            if (!((fi79 ^ (fi48 & fi79)) == fi48)) {
                                                if true {
                                                    v83 := GenS5{F0: fi48}
                                                    v22 = (float32(float64(-2.25)) * v22)
                                                    _ = len(v3)
                                                    _ = v83
                                                }
                                            }
                                            v3 = append(v3, ((v2 + "\n\t") + "\n\t"))
                                        }
                                        switch fi48 {
                                        case 32768:
                                            var v84 int64 = 227
                                            _ = v84
                                        case 4294967295:
                                            _ = len(v4)
                                        case -1:
                                            v85 := GenS8{F0: 65535}
                                            _ = v85
                                        }
                                    } else {
                                        v40 -= -1
                                        v86 := GenS7{F0: 42, F1: -104}
                                        _ = v86
                                    }
                                    var v87 bool = true
                                    if v87 {
                                        if v87 {
                                            delete(v4, genFn1(-422))
                                            v2 = v2
                                            if (v2 != (("\n\t" + v2) + (v2 + "hello"))) {
                                                if v87 {
                                                    var v88 bool = (1 <= 255)
                                                    v88 = (!(v40 <= 127))
                                                    _ = v88
                                                } else {
                                                    if v87 {
                                                        fmt.Println(v22)
                                                        v89 := GenS1{F0: (^(v5)), F1: 2147483647}
                                                        _ = v89
                                                    } else {
                                                        delete(v6, (^(fi79)))
                                                        var v90 uint32 = 2147483647
                                                        _ = v90
                                                    }
                                                    var v91 int16 = 0
                                                    _ = v91
                                                }
                                            }
                                        } else {
                                            v87 = v87
                                            if true {
                                                if false {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    var v92 float64 = (0.1 + ((3.14159 / 1.0) / float64(((v22 * v22) / ((v22 / v22) + v22)))))
                                                    var v93 float32 = v22
                                                    v94 := GenS9{F0: -9223372036854775808}
                                                    if len(v3) > 0 {
                                                        v3[0] = ""
                                                    }
                                                    var v95 uint8 = 220
                                                    v96 := GenS10{F0: -128, F1: v95}
                                                    v97 := v6[fi79]
                                                    _ = v92
                                                    _ = v93
                                                    _ = v94
                                                    _ = v95
                                                    _ = v96
                                                    _ = v97
                                                } else {
                                                    v98 := v4[4294967295]
                                                    if len(v3) > 0 {
                                                        v3[0] = (v2 + ("hello" + ""))
                                                    }
                                                    _ = v98
                                                }
                                            } else {
                                                switch "a" {
                                                case "hello":
                                                    v87 = v87
                                                case "":
                                                    v2 = v2
                                                case "a":
                                                    fmt.Println(fi79, v22)
                                                }
                                            }
                                        }
                                    } else {
                                        if ((float64(float32(3.14159)) * 3.14159) != float64((v22 / 0.1))) {
                                            v99 := GenS3{F0: (2 == 1)}
                                            if v87 {
                                                var v100 string = "a"
                                                v1 = (v1 % (((v1 % (((^(v1))) | 1))) | 1))
                                                _ = v100
                                            } else {
                                                var v101 uint64 = 4294967295
                                                _ = v101
                                            }
                                            if true {
                                                var v102 int32 = 973
                                                if (285 < 65535) {
                                                    var v103 uint64 = (v5 >> ((v1) & 63))
                                                    switch (fi79 / ((fi79) | 1)) {
                                                    case 128:
                                                        v104 := GenS7{F0: v102, F1: v40}
                                                        _ = v104
                                                    }
                                                    v103 -= 302
                                                    _ = v103
                                                } else {
                                                    _ = len(v6)
                                                    switch (fi79 ^ int(v1)) {
                                                    case 128:
                                                        v87 = v87
                                                    default:
                                                        if false {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            if false {
                                                                v87 = v87
                                                                var v105 uint8 = 1
                                                                v105 = 203
                                                                v22 += float32(1e10)
                                                                var v106 int8 = (v40 * (^(v40)))
                                                                var v107 int64 = -228
                                                                _ = v105
                                                                _ = v106
                                                                _ = v107
                                                            } else {
                                                                fmt.Println(fi48, v87)
                                                            }
                                                            _ = len(v3)
                                                        } else {
                                                            v108 := GenS11{F0: v1, F1: 1}
                                                            _ = v108
                                                        }
                                                    }
                                                    v109 := GenS12{F0: 1, F1: (("foo bar" + "hello") + (v2 + ((v2 + v2) + v2))), F2: 42}
                                                    _ = v109
                                                }
                                                fmt.Println(v102)
                                                _ = v102
                                            } else {
                                                var v110 bool = (!v87)
                                                var v111 uint32 = 2147483647
                                                v5 += (v5 ^ (v5 - v5))
                                                _ = v110
                                                _ = v111
                                            }
                                            switch 0 {
                                            case -2147483648:
                                                switch v2 {
                                                case "a":
                                                    var v112 float64 = float64(1e10)
                                                    _ = v112
                                                case "hello":
                                                    v2 = v2
                                                default:
                                                    v87 = v87
                                                }
                                            case 678:
                                                v113 := GenS13{F0: (v40 | v40), F1: 2, F2: 128}
                                                _ = v113
                                            }
                                            _ = v99
                                        }
                                        fmt.Println(v5, v5)
                                        var v114 int16 = 247
                                        _ = v114
                                    }
                                    _ = fi79
                                    _ = v87
                                }
                                _ = fi48
                            }
                            v1 -= v1
                        } else {
                            for fi115 := 0; fi115 < 7; fi115++ {
                                if false {
                                    if false {
                                        var v116 float64 = -2.25
                                        _ = v116
                                    }
                                    for fi117 := 0; fi117 < 17; fi117++ {
                                        var v118 int32 = 42
                                        v119 := GenS7{F0: v118, F1: (v40 + int8(v118))}
                                        fmt.Println(fi117, v22, fi117)
                                        v40 += (^(v40))
                                        var v120 uint32 = 73
                                        _ = fi117
                                        _ = v118
                                        _ = v119
                                        _ = v120
                                    }
                                } else {
                                    if len(v3) > 0 {
                                        v3[0] = "\n\t"
                                    }
                                    if (766 >= 1) {
                                        if false {
                                            switch -392 {
                                            case -2147483648:
                                                v1 = v1
                                            case 989:
                                                var v121 uint16 = 0
                                                _ = v121
                                            case 128:
                                                switch v2 {
                                                case "hello":
                                                    if ((v5 - (v5 | (v5 / ((v5) | 1)))) != (^(v5))) {
                                                        var v122 int16 = 1
                                                        delete(v4, genFn1(1))
                                                        _ = v122
                                                    }
                                                case "a":
                                                    v1 = v1
                                                case "":
                                                    for fi123 := 0; fi123 < 2; fi123++ {
                                                        v124 := GenS14{F0: (v2 + "\n\t"), F1: v2, F2: ("a" + (v2 + ""))}
                                                        v1 = v1
                                                        _ = fi123
                                                        _ = v124
                                                    }
                                                }
                                            }
                                            fmt.Println(v2, v2, v40)
                                            v125 := GenS15{F0: v1, F1: 255}
                                            _ = v125
                                        }
                                    }
                                    switch fi115 {
                                    case 127:
                                        fmt.Println(v40)
                                    case -9223372036854775808:
                                        if false {
                                            v126 := GenS9{F0: -1}
                                            _ = v126
                                        } else {
                                            v22 = v22
                                            v127 := v6[(-(fi115))]
                                            _ = v127
                                        }
                                    case 42:
                                        var v128 bool = true
                                        _ = v128
                                    }
                                }
                                if true {
                                    switch v2 {
                                    case "foo bar":
                                        var v129 uint16 = 2
                                        _ = v129
                                    }
                                    if (2 <= 255) {
                                        if (false && true) {
                                            v130 := GenS16{F0: 65535, F1: (v40 / ((v40) | 1)), F2: uint32(v1)}
                                            _ = v130
                                        }
                                        for fi131 := 0; fi131 < 14; fi131++ {
                                            var v132 int8 = (v40 % (((v40 / ((1) | 1))) | 1))
                                            _ = fi131
                                            _ = v132
                                        }
                                        v3 = append(v3, "\n\t")
                                    } else {
                                        v2 = (((v2 + (("" + ("hello" + "")) + v2)) + "hello") + v2)
                                        if false {
                                            if true {
                                                switch 433 {
                                                case 32768:
                                                    var v133 uint64 = v5
                                                    _ = v133
                                                default:
                                                    var v134 int64 = -332
                                                    _ = v134
                                                }
                                                var v135 uint32 = 701
                                                _ = v135
                                            }
                                            if true {
                                                var v136 string = "\n\t"
                                                v137 := GenS17{F0: v136, F1: 2147483647}
                                                v138 := GenS18{F0: -51, F1: (v1 / (((^(v1))) | 1)), F2: 2147483648}
                                                _ = v136
                                                _ = v137
                                                _ = v138
                                            } else {
                                                for fi139 := 0; fi139 < 8; fi139++ {
                                                    v2 = ("foo bar" + v2)
                                                    _ = fi139
                                                }
                                            }
                                        }
                                    }
                                    fmt.Println(v40, v22)
                                } else {
                                    switch fi115 {
                                    case 127:
                                        v140 := GenS19{F0: 1}
                                        _ = v140
                                    case -2147483648:
                                        var v141 uint = v1
                                        _ = v141
                                    case 661:
                                        v142 := GenS3{F0: (-2.25 <= 0.0)}
                                        _ = v142
                                    }
                                    fmt.Println(v5)
                                }
                                fmt.Println(fi115, v2, v40)
                                _ = fi115
                            }
                            fmt.Println(v2, v1)
                        }
                        for fi143 := 0; fi143 < 4; fi143++ {
                            v144 := GenS20{F0: 1000}
                            v2 = v2
                            for fi145 := 0; fi145 < 6; fi145++ {
                                if true {
                                    v146 := GenS8{F0: 348}
                                    var v147 uint16 = 32767
                                    v40 -= v40
                                    _ = v146
                                    _ = v147
                                } else {
                                    if ((!(0 < -1)) && (int64(v1) < 190)) {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        fmt.Println(v2, v1, v1)
                                        var v148 int32 = 2147483647
                                        _ = v148
                                    } else {
                                        _ = len(v3)
                                        v2 = ((("\n\t" + "") + "foo bar") + "hello")
                                    }
                                    fmt.Println(v2)
                                    fmt.Println(v5)
                                }
                                fmt.Println(v22)
                                var v149 int = fi143
                                _ = fi145
                                _ = v149
                            }
                            _ = fi143
                            _ = v144
                        }
                    }
                    _ = v40
                } else {
                    var v150 int8 = -128
                    var v151 float64 = 3.14159
                    v5 = v5
                    _ = v150
                    _ = v151
                }
            }
            for fi152 := 0; fi152 < 18; fi152++ {
                for fi153 := 0; fi153 < 19; fi153++ {
                    switch (fi153 % ((fi153) | 1)) {
                    case 4294967295:
                        var v154 uint8 = 1
                        if len(v3) > 0 {
                            v3[0] = (("\n\t" + "\n\t") + v2)
                        }
                        _ = v154
                    case -9223372036854775808:
                        v155 := GenS3{F0: (false || false)}
                        _ = v155
                    case 128:
                        var v156 float32 = v22
                        fmt.Println(v156, v2)
                        if (false && (true && (-9223372036854775808 == 42))) {
                            v2 = (v2 + ("a" + "hello"))
                            var v157 int64 = 42
                            v2 = v2
                            _ = v157
                        } else {
                            var v158 int8 = 0
                            v3 = append(v3, "\n\t")
                            var v159 int8 = -46
                            _ = v158
                            _ = v159
                        }
                        _ = v156
                    }
                    _ = fi153
                }
                for fi160 := 0; fi160 < 15; fi160++ {
                    v161 := GenS21{F0: (v5 << (((v1 - (v1 ^ (^(v1))))) & 63))}
                    if false {
                        fmt.Println(v5, fi152)
                        v2 = ("\n\t" + (v2 + (v2 + v2)))
                    }
                    switch (v2 + v2) {
                    case "foo bar":
                        fmt.Println(fi152)
                    }
                    var v162 uint = 18446744073709551615
                    if true {
                        switch (-(fi152)) {
                        case 4294967295:
                            v5 = v5
                        default:
                            v163 := GenS11{F0: (v1 >> ((v1) & 63)), F1: -4}
                            _ = v163
                        }
                        v3 = append(v3, v2)
                    }
                    fmt.Println(v162, v162, fi152)
                    if false {
                        var v164 uint = 255
                        _ = v164
                    } else {
                        v165 := GenS10{F0: 1, F1: 128}
                        _ = v165
                    }
                    _ = fi160
                    _ = v161
                    _ = v162
                }
                for fi166 := 0; fi166 < 13; fi166++ {
                    v167 := GenS6{F0: -9223372036854775808}
                    var v168 bool = (255 == 605)
                    _ = fi166
                    _ = v167
                    _ = v168
                }
                _ = fi152
            }
        }
    }
    _ = len(v3)
    switch -759 {
    case 32767:
        for fi169 := 0; fi169 < 4; fi169++ {
            for fi170 := 0; fi170 < 16; fi170++ {
                if false {
                    if false {
                        v171 := GenS2{F0: 1, F1: (!true), F2: true}
                        var v172 int32 = -1
                        if (v1 >= 2) {
                            v172 = v172
                            delete(v6, -9223372036854775808)
                            var v173 int16 = 32767
                            _ = v173
                        } else {
                            v174 := v6[671]
                            switch fi169 {
                            case 203:
                                if false {
                                    v175 := GenS22{F0: 2}
                                    _ = v175
                                }
                            case -2147483648:
                                v176 := v6[(fi170 - fi170)]
                                _ = v176
                            }
                            _ = v174
                        }
                        _ = v171
                        _ = v172
                    } else {
                        v177 := GenS3{F0: (32767 >= 65535)}
                        _ = v177
                    }
                    fmt.Println(fi170, v22, fi170)
                }
                if true {
                    var v178 int32 = 260
                    v178 = -817
                    _ = v178
                }
                if ("" >= "a") {
                    switch fi170 {
                    case 32768:
                        var v179 uint8 = 0
                        _ = v179
                    }
                } else {
                    v180 := GenS15{F0: v1, F1: 1}
                    fmt.Println(fi169)
                    if (!true) {
                        if (v2 < (((((v2 + v2) + "") + "a") + v2) + v2)) {
                            if true {
                                _ = len(v4)
                                if (!false) {
                                    v2 = v2
                                    switch (fi169 * (fi169 | fi170)) {
                                    case 127:
                                        v2 = "a"
                                    }
                                } else {
                                    fmt.Println(v2, v22)
                                    v22 = (v22 / 1.0)
                                }
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if true {
                                var v181 string = "\n\t"
                                v3 = append(v3, v2)
                                _ = v181
                            } else {
                                v2 = v2
                                switch (fi169 & (fi170 % ((32768) | 1))) {
                                case 518:
                                    v1 = 838
                                case 190:
                                    _ = len(v3)
                                case 128:
                                    fmt.Println(v1, fi170, v22)
                                default:
                                    var v182 float64 = 0.0
                                    _ = v182
                                }
                                _ = len(v3)
                            }
                            fmt.Println(fi170, fi169)
                        } else {
                            v22 = (v22 + 0.1)
                            if false {
                                if (!false) {
                                    if false {
                                        fmt.Println(v2)
                                        fmt.Println(v1, v5)
                                    }
                                    if true {
                                        v22 -= 3.14159
                                        v3 = append(v3, v2)
                                        var v183 uint32 = 829
                                        _ = v183
                                    }
                                    v22 = v22
                                }
                                if (0.0 < (0.0 * 0.1)) {
                                    fmt.Println(fi169, v5, v22)
                                    fmt.Println(v1, fi169, v2)
                                    v22 += float32(1e10)
                                } else {
                                    v184 := GenS11{F0: (^(v1)), F1: 0}
                                    _ = v184
                                }
                                if true {
                                    switch "" {
                                    case "foo bar":
                                        v5 += v5
                                    case "hello":
                                        var v185 float32 = 3.14159
                                        _ = v185
                                    default:
                                        fmt.Println(v1, v22, v5)
                                    }
                                    v186 := GenS19{F0: 42}
                                    delete(v4, 42)
                                    _ = v186
                                } else {
                                    v1 -= (v1 >> ((v1) & 63))
                                }
                            } else {
                                v187 := GenS22{F0: (v1 & (^(v1)))}
                                v1 -= (v1 | 18446744073709551615)
                                v22 += (v22 * v22)
                                v2 = ""
                                var v188 int = fi170
                                _ = v187
                                _ = v188
                            }
                            switch 32767 {
                            case 2147483647:
                                v189 := GenS8{F0: 0}
                                _ = v189
                            case -42:
                                var v190 float32 = ((float32(float64(v22)) / ((v22 + v22) + 1e10)) + 0.1)
                                _ = v190
                            case -939:
                                v2 = v2
                            }
                        }
                        var v191 int8 = 1
                        _ = v191
                    }
                    var v192 uint64 = (v5 | v5)
                    _ = v180
                    _ = v192
                }
                _ = fi170
            }
            v5 = (^(v5))
            _ = fi169
        }
    case 0:
        fmt.Println(v22, v1, v5)
    case -191:
        switch v2 {
        case "\n\t":
            v193 := GenS22{F0: v1}
            _ = v193
        case "foo bar":
            v194 := GenS23{F0: 127, F1: -530, F2: 1}
            _ = v194
        case "":
            _ = len(v3)
        default:
            v195 := GenS19{F0: 255}
            _ = v195
        }
    }
    var v196 bool = (("" != (("\n\t" + (v2 + "hello")) + v2)) || (!false))
    for fi197 := 0; fi197 < 11; fi197++ {
        v198 := GenS18{F0: -1, F1: (v1 >> (((v1 % (((v1 | v1)) | 1))) & 63)), F2: 985}
        v199 := GenS1{F0: 1, F1: 249}
        _ = fi197
        _ = v198
        _ = v199
    }
    fmt.Println("v1:", v1)
    fmt.Println("v196:", v196)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v196
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
