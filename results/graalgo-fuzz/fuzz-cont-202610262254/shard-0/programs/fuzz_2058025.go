package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 uint8; F1 int64; F2 bool }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS7 struct { F0 uint16; F1 uint }
type GenS8 struct { F0 bool; F1 uint64; F2 int }
type GenS9 struct { F0 uint16; F1 uint }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int; F1 int8 }
type GenS13 struct { F0 bool; F1 int }
type GenS14 struct { F0 bool }
type GenS15 struct { F0 uint16; F1 string }
type GenS16 struct { F0 uint; F1 int16 }
type GenS17 struct { F0 int; F1 string; F2 int32 }
type GenS18 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS19 struct { F0 int8; F1 bool; F2 int8 }
type GenS20 struct { F0 int16; F1 int; F2 bool }
type GenS21 struct { F0 int16; F1 uint }
func genFn0(p0 int64) uint64 {
    p0 += (p0 & (p0 >> ((0) & 63)))
    p0 -= 414
    _ = p0
    return 2
}
func genFn1(p0 uint8, p1 int64, p2 int32) int32 {
    p2 -= (^(p2))
    for fi4 := 0; fi4 < 19; fi4++ {
        if (-1 <= -78) {
            var v5 int32 = (p2 >> ((18446744073709551615) & 63))
            fmt.Println(v5, v5, p2)
            _ = v5
        }
        if (!((((1.5 - 1.5) != 0.1) && (!(78 <= 111))) || (957 == 65535))) {
            p0 = 127
            fmt.Println(fi4, p1, p2)
        }
        _ = fi4
    }
    v6 := []int{812, 2147483647, 591}
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 string = ("\n\t" + ("hello" + "\n\t"))
        _ = fi7
        _ = v8
    }
    switch -9223372036854775808 {
    case -1:
        switch -9223372036854775808 {
        case 127:
            v9 := GenS1{F0: p2, F1: (p1 / ((p1) | 1))}
            _ = v9
        case 588:
            p0 -= (p0 - p0)
        default:
            var v10 float64 = float64((((((((0.0 / 1e10) - 3.14159) + 1.5) + 0.1) * float32(0.1)) * (0.0 - 0.0)) * (1.5 - 0.1)))
            _ = v10
        }
    case -28:
        switch "" {
        case "\n\t":
            p0 = (^(p0))
        case "hello":
            var v11 uint = 4294967295
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p2 | p2)
}

func main() {
    v1 := make(map[string]int)
    v2 := []string{("" + "a")}
    var v3 int = -785
    var v12 uint16 = 574
    var v13 int8 = 0
    var v14 string = "hello"
    var v15 uint32 = 1
    v16 := GenS1{F0: 697, F1: -9223372036854775808}
    v13 += (v13 - (v13 ^ (v13 % ((v13) | 1))))
    switch (^(v3)) {
    case 32767:
        fmt.Println(v12, v12)
    }
    switch -980 {
    case 1:
        var v17 uint32 = v15
        _ = v17
    case -1:
        fmt.Println(v13)
    default:
        v1[v14] = (-(v3))
    }
    if true {
        for fi18 := 0; fi18 < 3; fi18++ {
            var v19 bool = (1.5 >= 0.1)
            _ = len(v2)
            v3 = v3
            _ = fi18
            _ = v19
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch (v3 | (v3 ^ v3)) {
    case 912:
        if true {
            for fi20 := 0; fi20 < 1; fi20++ {
                fmt.Println(v3, fi20, v14)
                var v21 string = v14
                fmt.Println(fi20, v12)
                _ = fi20
                _ = v21
            }
        }
    case -766:
        var v22 int32 = 0
        _ = v22
    case 376:
        for fi23 := 0; fi23 < 8; fi23++ {
            if len(v2) > 0 {
                v2[0] = "foo bar"
            }
            v24 := GenS2{F0: 79, F1: -1, F2: true}
            if false {
                if (false || ((79 >= 2) && false)) {
                    for fi25 := 0; fi25 < 12; fi25++ {
                        if (-128 >= (v13 & (v13 >> ((316) & 63)))) {
                            var v26 uint8 = 127
                            v27 := GenS3{F0: v15, F1: 487}
                            _ = v26
                            _ = v27
                        } else {
                            v3 -= v3
                            v3 = (v3 + 561)
                        }
                        _ = fi25
                    }
                }
            } else {
                var v28 string = v14
                for fi29 := 0; fi29 < 8; fi29++ {
                    if true {
                        if len(v2) > 0 {
                            v2[0] = (v14 + (("" + "\n\t") + "\n\t"))
                        }
                        fmt.Println(fi23, fi29)
                        v3 += -1
                        fmt.Println(v28)
                    } else {
                        v13 = (v13 - v13)
                        fmt.Println(v28, v13, v3)
                    }
                    fmt.Println(fi23, v3)
                    _ = fi29
                }
                _ = v28
            }
            _ = fi23
            _ = v24
        }
    }
    if false {
        for fi30 := 0; fi30 < 12; fi30++ {
            for fi31 := 0; fi31 < 20; fi31++ {
                fmt.Println(v13, fi31)
                if false {
                    v32 := GenS4{F0: 1000, F1: v13}
                    _ = v32
                } else {
                    v33 := GenS5{F0: 52, F1: 735}
                    fmt.Println(v15)
                    _ = v33
                }
                _ = fi31
            }
            if (v14 == "") {
                if false {
                    v1["\n\t"] = -9223372036854775808
                    switch 0 {
                    case 4294967295:
                        v34 := GenS4{F0: -32768, F1: 1}
                        _ = v34
                    case 128:
                        switch 32767 {
                        case 32767:
                            var v35 float64 = float64(-2.25)
                            _ = v35
                        }
                    default:
                        var v36 string = (v14 + v14)
                        _ = v36
                    }
                    v14 = ""
                } else {
                    v13 += (v13 / (((v13 - -1)) | 1))
                }
                for fi37 := 0; fi37 < 16; fi37++ {
                    var v38 float64 = 3.14159
                    if (!(1e10 >= 3.14159)) {
                        var v39 uint8 = 105
                        v1[v14] = (^(v3))
                        var v40 int16 = 1
                        var v41 float64 = v38
                        v42 := v1[v14]
                        _ = v39
                        _ = v40
                        _ = v41
                        _ = v42
                    } else {
                        if (false && ((0.1 > 1.5) || (((!false) || ((127 >= 0) || (true && false))) && true))) {
                            v12 -= (v12 << ((385) & 63))
                        } else {
                            _ = len(v2)
                            var v43 string = "a"
                            v44 := GenS5{F0: 42, F1: -699}
                            _ = v43
                            _ = v44
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if (v38 <= (v38 * ((v38 + v38) * (v38 - (0.1 * v38))))) {
                            v38 -= ((float64(1e10) / (float64(1.5) * ((3.14159 / v38) + (float64(-2.25) - v38)))) * (1.5 / (1e10 + (-2.25 - (0.0 + v38)))))
                        } else {
                            var v45 uint64 = 42
                            _ = v45
                        }
                    }
                    v14 = v14
                    _ = fi37
                    _ = v38
                }
                switch "a" {
                case "hello":
                    if (true && true) {
                        var v46 bool = false
                        if v46 {
                            if (((!v46) || v46) && v46) {
                                if (2147483647 >= -1) {
                                    v47 := GenS2{F0: 222, F1: 9223372036854775807, F2: v46}
                                    switch fi30 {
                                    case 32768:
                                        v46 = v46
                                    case 127:
                                        if (!v46) {
                                            v48 := GenS6{F0: 808, F1: uint8(v12), F2: 0}
                                            v3 = (fi30 % ((v3) | 1))
                                            var v49 int32 = 42
                                            _ = v48
                                            _ = v49
                                        } else {
                                            fmt.Println(v15, v14, v15)
                                        }
                                    }
                                    _ = len(v1)
                                    _ = v47
                                } else {
                                    var v50 uint8 = 127
                                    _ = v50
                                }
                                for fi51 := 0; fi51 < 5; fi51++ {
                                    v52 := GenS7{F0: 1, F1: 789}
                                    v2 = append(v2, v14)
                                    var v53 float32 = (((0.1 - ((1.5 * 3.14159) * 0.0)) / 1.0) / -2.25)
                                    _ = fi51
                                    _ = v52
                                    _ = v53
                                }
                                var v54 int64 = 42
                                _ = v54
                            } else {
                                var v55 float64 = (-2.25 / 1.0)
                                _ = v55
                            }
                            v1[v14] = (^(fi30))
                            for fi56 := 0; fi56 < 5; fi56++ {
                                if (!true) {
                                    _ = len(v2)
                                    v57 := GenS2{F0: 144, F1: -1, F2: false}
                                    if v46 {
                                        v13 -= v13
                                    }
                                    var v58 float64 = 0.0
                                    if (!true) {
                                        var v59 float32 = 0.0
                                        if (v46 && (-1 >= 0)) {
                                            _ = len(v2)
                                        }
                                        _ = v59
                                    }
                                    _ = v57
                                    _ = v58
                                }
                                _ = fi56
                            }
                        }
                        var v60 int32 = 42
                        _ = v46
                        _ = v60
                    }
                default:
                    for fi61 := 0; fi61 < 1; fi61++ {
                        v62 := GenS8{F0: (18446744073709551615 > 0), F1: 2, F2: (fi30 / ((fi61) | 1))}
                        v12 += 0
                        _ = fi61
                        _ = v62
                    }
                }
            } else {
                var v63 int8 = v13
                if false {
                    for fi64 := 0; fi64 < 8; fi64++ {
                        if true {
                            v65 := GenS1{F0: -330, F1: -384}
                            var v66 bool = true
                            if true {
                                var v67 int32 = -977
                                if (128 > 1) {
                                    var v68 string = v14
                                    var v69 int16 = 602
                                    var v70 int8 = (^(v63))
                                    _ = len(v1)
                                    v71 := GenS9{F0: (v12 - v12), F1: 137}
                                    _ = v68
                                    _ = v69
                                    _ = v70
                                    _ = v71
                                } else {
                                    var v72 uint16 = (v12 % (((v12 << ((uint(v12)) & 63))) | 1))
                                    if true {
                                        if v66 {
                                            v12 += v72
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            v72 = (v12 ^ v72)
                                            if v66 {
                                                if v66 {
                                                    var v73 int8 = v63
                                                    _ = v73
                                                }
                                                if ((v12 >> ((18446744073709551615) & 63)) >= 1) {
                                                    var v74 float64 = float64((-2.25 / ((1e10 / 1.5) / (((float32(3.14159) / (1.5 + (3.14159 - 0.1))) * -2.25) + 0.1))))
                                                    if (!((-(fi30)) >= fi30)) {
                                                        switch v14 {
                                                        case "":
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        case "a":
                                                            fmt.Println(v63, v15, v74)
                                                        case "hello":
                                                            v3 -= 32768
                                                        }
                                                        fmt.Println(v3)
                                                        if (int(v15) >= (fi30 - (fi64 | -793))) {
                                                            switch fi64 {
                                                            case 595:
                                                                v2 = append(v2, v14)
                                                            case -379:
                                                                if (!(-9223372036854775808 <= fi64)) {
                                                                    fmt.Println(v72, fi30, v67)
                                                                    fmt.Println(fi30, v63, v15)
                                                                } else {
                                                                    v75 := GenS1{F0: -649, F1: 42}
                                                                    _ = v75
                                                                }
                                                            case 4294967295:
                                                                var v76 uint = 0
                                                                _ = v76
                                                            default:
                                                                if ((v74 / (0.1 + (3.14159 + v74))) > (v74 - (v74 * ((v74 - -2.25) * (0.1 - (v74 - -2.25)))))) {
                                                                    v77 := v1[v14]
                                                                    if ((!v66) && v66) {
                                                                        v14 = (((v14 + "hello") + ("a" + "foo bar")) + "")
                                                                        var v78 uint8 = 0
                                                                        delete(v1, v14)
                                                                        _ = v78
                                                                    }
                                                                    var v79 float64 = (0.1 / float64(((1e10 / 3.14159) - 3.14159)))
                                                                    _ = v77
                                                                    _ = v79
                                                                }
                                                            }
                                                        } else {
                                                            v67 -= v67
                                                        }
                                                        switch (-(fi64)) {
                                                        case -1:
                                                            switch ("" + (("hello" + v14) + v14)) {
                                                            case "\n\t":
                                                                v67 = v67
                                                            case "hello":
                                                                var v80 uint8 = 215
                                                                _ = v80
                                                            case "a":
                                                                switch "hello" {
                                                                case "hello":
                                                                    if (v74 == ((((1.5 + (0.0 + float64(3.14159))) - v74) - float64(((1e10 * 0.1) + float32(float64(3.14159))))) - v74)) {
                                                                        v81 := GenS10{F0: 104}
                                                                        if (!false) {
                                                                            if (v66 || (v66 && v66)) {
                                                                                var v82 uint8 = 127
                                                                                v82 = (v82 & 0)
                                                                                var v83 int16 = -1
                                                                                _ = v82
                                                                                _ = v83
                                                                            } else {
                                                                                delete(v1, v14)
                                                                                var v84 int16 = 42
                                                                                var v85 uint32 = 4294967295
                                                                                fmt.Println(v15)
                                                                                var v86 uint = 305
                                                                                _ = v84
                                                                                _ = v85
                                                                                _ = v86
                                                                            }
                                                                            v13 -= v13
                                                                            _ = len(v2)
                                                                        } else {
                                                                            if len(v2) > 0 {
                                                                                v2[0] = v14
                                                                            }
                                                                            v87 := GenS11{F0: v14}
                                                                            _ = v87
                                                                        }
                                                                        v88 := GenS12{F0: (-(v3)), F1: -1}
                                                                        _ = v81
                                                                        _ = v88
                                                                    } else {
                                                                        v13 = v63
                                                                    }
                                                                }
                                                            default:
                                                                if v66 {
                                                                    v89 := GenS13{F0: (42 <= -9223372036854775808), F1: (fi64 + (^(fi30)))}
                                                                    var v90 uint8 = 42
                                                                    _ = v89
                                                                    _ = v90
                                                                }
                                                            }
                                                        case 179:
                                                            var v91 string = "a"
                                                            _ = v91
                                                        }
                                                    }
                                                    _ = v74
                                                }
                                                v63 = v13
                                                if false {
                                                    v92 := v1[v14]
                                                    _ = v92
                                                }
                                            } else {
                                                var v93 float32 = 0.1
                                                var v94 int8 = 100
                                                fmt.Println(v72, v66)
                                                v67 = 42
                                                v72 = v12
                                                _ = v93
                                                _ = v94
                                            }
                                        }
                                        if (genFn0(-9223372036854775808) <= genFn0(1)) {
                                            var v95 int8 = 42
                                            _ = v95
                                        }
                                    } else {
                                        if v66 {
                                            v72 = v12
                                        }
                                    }
                                    if true {
                                        if ((v63 << ((18446744073709551615) & 63)) == v13) {
                                            var v96 int64 = 42
                                            _ = v96
                                        }
                                    }
                                    if (v66 || true) {
                                        if v66 {
                                            var v97 uint32 = v15
                                            _ = v97
                                        }
                                        _ = len(v2)
                                        v14 = "\n\t"
                                        v63 -= (v13 ^ v13)
                                        v98 := GenS14{F0: false}
                                        v72 = (v72 << ((371) & 63))
                                        switch v14 {
                                        case "\n\t":
                                            var v99 uint = 255
                                            _ = v99
                                        case "hello":
                                            switch (v3 % ((int(v15)) | 1)) {
                                            case 0:
                                                v100 := GenS8{F0: (!false), F1: 347, F2: -125}
                                                fmt.Println(v63, fi64, v12)
                                                if false {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    v67 = (-(v67))
                                                    v66 = (v15 > 648)
                                                } else {
                                                    v67 += (v67 << ((588) & 63))
                                                    v14 = v14
                                                    v2 = append(v2, (v14 + v14))
                                                    switch ("hello" + "") {
                                                    case "a":
                                                        fmt.Println(v3)
                                                    case "\n\t":
                                                        v101 := GenS15{F0: 732, F1: (v14 + (v14 + (v14 + "\n\t")))}
                                                        _ = v101
                                                    }
                                                    if (!(!(!(v63 < -1)))) {
                                                        if v66 {
                                                            var v102 uint = 255
                                                            _ = v102
                                                        } else {
                                                            switch v14 {
                                                            case "\n\t":
                                                                if (!(201 > 1)) {
                                                                    v103 := GenS16{F0: 255, F1: 42}
                                                                    var v104 uint32 = v15
                                                                    var v105 string = v14
                                                                    _ = v103
                                                                    _ = v104
                                                                    _ = v105
                                                                }
                                                            case "a":
                                                                v67 = genFn1(127, 1, v67)
                                                            case "":
                                                                v3 -= (fi64 << ((296) & 63))
                                                            default:
                                                                var v106 int16 = 0
                                                                _ = v106
                                                            }
                                                        }
                                                        v107 := GenS17{F0: fi64, F1: (v14 + "a"), F2: (v67 - v67)}
                                                        _ = v107
                                                    } else {
                                                        if (215 == uint8(v15)) {
                                                            var v108 string = "a"
                                                            var v109 string = "\n\t"
                                                            _ = v108
                                                            _ = v109
                                                        }
                                                    }
                                                    var v110 uint8 = 0
                                                    v111 := GenS18{F0: v66, F1: 252, F2: (v15 << ((2) & 63))}
                                                    _ = v110
                                                    _ = v111
                                                }
                                                _ = v100
                                            case 128:
                                                v15 = (v15 / (((v15 ^ (v15 % ((1) | 1)))) | 1))
                                            case 42:
                                                delete(v1, v14)
                                            }
                                            var v112 uint8 = 159
                                            v66 = (138 < 18446744073709551615)
                                            v3 = fi64
                                            switch (v14 + v14) {
                                            case "foo bar":
                                                var v113 uint16 = 1
                                                _ = v113
                                            case "a":
                                                var v114 int64 = -1
                                                _ = v114
                                            case "":
                                                v115 := GenS2{F0: uint8(v12), F1: 826, F2: (!(-2.25 < 1e10))}
                                                _ = v115
                                            default:
                                                v116 := GenS14{F0: ((-2147483648 == -699) && (!((v15 > (v15 & (v15 ^ (v15 | (v15 & v15))))) && v66)))}
                                                _ = v116
                                            }
                                            v117 := GenS3{F0: (v15 % ((4294967295) | 1)), F1: 67}
                                            v67 = (v67 / (((v67 + genFn1(uint8(v15), 1, (v67 << ((2) & 63))))) | 1))
                                            _ = v112
                                            _ = v117
                                        default:
                                            v13 = (v63 % (((v63 + (v13 * 127))) | 1))
                                        }
                                        _ = v98
                                    } else {
                                        var v118 uint = 42
                                        var v119 uint16 = (^(v12))
                                        var v120 int32 = genFn1(230, 42, v67)
                                        _ = v118
                                        _ = v119
                                        _ = v120
                                    }
                                    var v121 uint32 = (^(v15))
                                    var v122 bool = false
                                    _ = v121
                                    _ = v122
                                    _ = v72
                                }
                                _ = v67
                            } else {
                                var v123 uint8 = 127
                                _ = v123
                            }
                            _ = v65
                            _ = v66
                        }
                        v2 = append(v2, (v14 + ((("" + "a") + (v14 + v14)) + v14)))
                        v124 := GenS19{F0: (^(v13)), F1: (false && (float64(0.0) < 1e10)), F2: (^(v63))}
                        _ = fi64
                        _ = v124
                    }
                } else {
                    for fi125 := 0; fi125 < 2; fi125++ {
                        fmt.Println(fi125, v13, fi30)
                        _ = fi125
                    }
                    if (!false) {
                        if (!true) {
                            var v126 int16 = 0
                            v15 = v15
                            _ = v126
                        }
                        if false {
                            var v127 int = v3
                            v128 := GenS19{F0: -128, F1: (1e10 >= 0.1), F2: v13}
                            _ = v127
                            _ = v128
                        } else {
                            for fi129 := 0; fi129 < 8; fi129++ {
                                v15 = v15
                                fmt.Println(v14, fi129)
                                v130 := GenS19{F0: 63, F1: (false && false), F2: 18}
                                _ = fi129
                                _ = v130
                            }
                        }
                    } else {
                        if (32767 >= 0) {
                            v131 := GenS20{F0: 643, F1: (fi30 + (fi30 >> ((461) & 63))), F2: true}
                            v132 := GenS15{F0: (v12 << ((0) & 63)), F1: (v14 + (v14 + v14))}
                            var v133 uint16 = v12
                            _ = v131
                            _ = v132
                            _ = v133
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            for fi134 := 0; fi134 < 4; fi134++ {
                                v3 += -1
                                _ = fi134
                            }
                            var v135 int64 = 9223372036854775807
                            _ = v135
                        }
                        if (!true) {
                            var v136 uint = 65535
                            fmt.Println(fi30)
                            var v137 int = v3
                            _ = v136
                            _ = v137
                        } else {
                            if true {
                                v3 = fi30
                            } else {
                                for fi138 := 0; fi138 < 7; fi138++ {
                                    if len(v2) > 0 {
                                        v2[0] = (v14 + ((v14 + ((v14 + ("hello" + "")) + "\n\t")) + "hello"))
                                    }
                                    _ = fi138
                                }
                            }
                            if false {
                                v63 = (v13 >> ((42) & 63))
                                v15 = (v15 % ((1) | 1))
                                var v139 int8 = v63
                                _ = v139
                            }
                            var v140 uint32 = 0
                            _ = v140
                        }
                        switch (fi30 * -2147483648) {
                        case 32767:
                            v141 := GenS21{F0: -1, F1: 477}
                            _ = v141
                        default:
                            var v142 int32 = -1
                            _ = v142
                        }
                    }
                }
                var v143 int8 = v13
                _ = v143
                _ = v63
            }
            var v144 string = (("a" + v14) + v14)
            _ = fi30
            _ = v144
        }
        v3 = v3
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
