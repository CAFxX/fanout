package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 int8; F2 uint64 }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 uint8; F1 bool; F2 int64 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 int16; F1 int32 }
type GenS11 struct { F0 int16; F1 int16 }
type GenS12 struct { F0 uint; F1 uint16 }
func genFn0(p0 bool) int {
    if (p0 && p0) {
        fmt.Println(p0)
    } else {
        p0 = true
    }
    var v7 uint8 = 255
    var v8 uint = 65535
    _ = p0
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 int16, p1 int32, p2 int8) int32 {
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 bool = ((p2 + -36) > p2)
        _ = fi9
        _ = v10
    }
    p1 -= p1
    v11 := GenS1{F0: (p1 / (((p1 >> ((685) & 63))) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return -1
}

func main() {
    var v1 int64 = -1
    v2 := make(map[int]int)
    var v3 int = 4294967295
    var v4 string = "foo bar"
    var v5 float32 = 0.1
    v6 := []int{(v3 - (v3 + -1))}
    if false {
        v12 := GenS1{F0: 0}
        if false {
            if false {
                v13 := GenS1{F0: -2147483648}
                if true {
                    for fi14 := 0; fi14 < 8; fi14++ {
                        for fi15 := 0; fi15 < 13; fi15++ {
                            v16 := GenS2{F0: 1000}
                            var v17 uint32 = 952
                            delete(v2, (fi15 >> ((uint(v17)) & 63)))
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                        v3 = fi14
                        _ = fi14
                    }
                }
                if (true || false) {
                    var v18 bool = (930 >= 0)
                    var v19 string = "\n\t"
                    _ = v18
                    _ = v19
                } else {
                    v20 := v2[128]
                    v21 := GenS2{F0: 1}
                    _ = v20
                    _ = v21
                }
                _ = v13
            }
        }
        _ = v12
    }
    if (-2147483648 < -2147483648) {
        for fi22 := 0; fi22 < 12; fi22++ {
            for fi23 := 0; fi23 < 20; fi23++ {
                var v24 int = (fi23 * genFn0((((!(42 >= 2147483647)) || false) || (127 != -1))))
                v5 = v5
                _ = fi23
                _ = v24
            }
            v25 := v2[77]
            _ = fi22
            _ = v25
        }
    } else {
        switch ("\n\t" + "foo bar") {
        case "":
            for fi26 := 0; fi26 < 18; fi26++ {
                fmt.Println(v5, v4)
                _ = fi26
            }
        }
        for fi27 := 0; fi27 < 8; fi27++ {
            switch 32768 {
            case 9223372036854775807:
                if (true && (0.0 >= 1.5)) {
                    var v28 int8 = -128
                    v3 = v3
                    _ = v28
                } else {
                    if (91 > 42) {
                        switch (fi27 + 2147483647) {
                        case -450:
                            var v29 uint64 = 0
                            _ = v29
                        case -2147483648:
                            v4 = (v4 + (("\n\t" + ((v4 + v4) + "hello")) + v4))
                        default:
                            v5 -= 3.14159
                        }
                    } else {
                        for fi30 := 0; fi30 < 14; fi30++ {
                            switch (v4 + ("\n\t" + v4)) {
                            case "":
                                v31 := v2[128]
                                _ = v31
                            case "hello":
                                if (0.1 != 0.1) {
                                    v32 := GenS3{F0: "foo bar"}
                                    var v33 int = fi27
                                    var v34 int32 = 42
                                    v33 -= genFn0(false)
                                    _ = v32
                                    _ = v33
                                    _ = v34
                                }
                            case "\n\t":
                                var v35 int16 = 32767
                                _ = v35
                            }
                            var v36 uint8 = 165
                            delete(v2, 32768)
                            _ = fi30
                            _ = v36
                        }
                        switch v3 {
                        case 9223372036854775807:
                            for fi37 := 0; fi37 < 8; fi37++ {
                                var v38 int16 = -32768
                                var v39 string = "hello"
                                _ = fi37
                                _ = v38
                                _ = v39
                            }
                        case 127:
                            v40 := GenS2{F0: 912}
                            _ = v40
                        case 128:
                            for fi41 := 0; fi41 < 16; fi41++ {
                                v42 := GenS1{F0: 1}
                                _ = fi41
                                _ = v42
                            }
                        }
                    }
                    var v43 uint8 = 77
                    switch "a" {
                    case "":
                        var v44 int8 = 58
                        _ = v44
                    case "hello":
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    default:
                        if false {
                            if false {
                                var v45 uint16 = 42
                                fmt.Println(v4, fi27)
                                _ = v45
                            } else {
                                var v46 int8 = 1
                                var v47 uint16 = 65535
                                _ = v46
                                _ = v47
                            }
                        }
                    }
                    _ = v43
                }
            }
            for fi48 := 0; fi48 < 5; fi48++ {
                v49 := GenS4{F0: -1, F1: 42, F2: 0}
                _ = len(v2)
                if true {
                    v1 += 1
                    v50 := GenS3{F0: (("\n\t" + "") + v4)}
                    v51 := GenS5{F0: (-(v1)), F1: 0}
                    _ = v50
                    _ = v51
                } else {
                    if (42 != (v1 | v1)) {
                        v5 = v5
                        switch "a" {
                        case "\n\t":
                            delete(v2, (v3 % ((32768) | 1)))
                        default:
                            if true {
                                v52 := GenS2{F0: 32767}
                                v3 += (v3 >> ((42) & 63))
                                fmt.Println(v5, v3)
                                fmt.Println(fi27, v5, fi27)
                                _ = v52
                            }
                        }
                        v4 = ("foo bar" + ("" + v4))
                    }
                    fmt.Println(v4, fi48, fi48)
                }
                _ = fi48
                _ = v49
            }
            if true {
                var v53 float32 = v5
                if false {
                    v53 += float32(0.1)
                }
                _ = v53
            }
            _ = fi27
        }
        if (genFn1(-1, -2147483648, 127) >= -790) {
            for fi54 := 0; fi54 < 17; fi54++ {
                v1 = v1
                _ = fi54
            }
        }
    }
    _ = len(v2)
    for fi55 := 0; fi55 < 14; fi55++ {
        if false {
            v56 := GenS6{F0: 144, F1: true, F2: 737}
            _ = v56
        } else {
            var v57 int8 = 70
            if (true || false) {
                if true {
                    if false {
                        for fi58 := 0; fi58 < 9; fi58++ {
                            switch v4 {
                            case "foo bar":
                                delete(v2, (fi58 << ((18446744073709551615) & 63)))
                            case "":
                                v59 := GenS7{F0: v1}
                                _ = v59
                            }
                            var v60 float32 = v5
                            _ = fi58
                            _ = v60
                        }
                    }
                }
            } else {
                for fi61 := 0; fi61 < 16; fi61++ {
                    var v62 int64 = v1
                    if (true && true) {
                        var v63 uint = 414
                        if true {
                            switch (v4 + "a") {
                            case "":
                                delete(v2, (fi55 + -1))
                            case "hello":
                                var v64 int16 = 42
                                switch ("foo bar" + (v4 + (v4 + v4))) {
                                case "":
                                    v1 -= (v1 - (v1 & 1))
                                }
                                _ = v64
                            case "foo bar":
                                _ = len(v6)
                            default:
                                if false {
                                    v65 := GenS5{F0: v62, F1: (v63 << (((v63 + 42)) & 63))}
                                    var v66 uint64 = 1
                                    v67 := GenS7{F0: -1}
                                    _ = v65
                                    _ = v66
                                    _ = v67
                                } else {
                                    v1 = v62
                                    v4 = v4
                                    fmt.Println(v57, v57)
                                }
                            }
                        }
                        _ = v63
                    } else {
                        switch v4 {
                        case "":
                            v4 = v4
                        default:
                            switch genFn0((18446744073709551615 < 2)) {
                            case -1:
                                if false {
                                    v5 = 0.0
                                } else {
                                    switch 0 {
                                    case -1:
                                        v6 = append(v6, genFn0(false))
                                    case 127:
                                        v68 := v2[genFn0(false)]
                                        _ = v68
                                    }
                                }
                            case 766:
                                v4 = "foo bar"
                            case 128:
                                v57 = (^(v57))
                            }
                        }
                    }
                    _ = fi61
                    _ = v62
                }
                v5 -= 3.14159
                var v69 int64 = (^(v1))
                _ = v69
            }
            v57 -= (v57 - v57)
            _ = v57
        }
        v4 = v4
        _ = fi55
    }
    v5 = v5
    if false {
        switch (v3 % ((v3) | 1)) {
        case 4294967295:
            for fi70 := 0; fi70 < 7; fi70++ {
                switch (((v4 + v4) + (v4 + ("foo bar" + v4))) + "a") {
                case "a":
                    switch v4 {
                    case "\n\t":
                        v2[42] = 42
                    case "a":
                        switch v4 {
                        case "hello":
                            switch (v3 << ((0) & 63)) {
                            case 877:
                                for fi71 := 0; fi71 < 8; fi71++ {
                                    v72 := GenS3{F0: (v4 + (("" + v4) + ("" + v4)))}
                                    v3 = 42
                                    var v73 int16 = -1
                                    _ = fi71
                                    _ = v72
                                    _ = v73
                                }
                            case 2147483647:
                                var v74 int32 = genFn1(80, genFn1(42, 1, 100), -1)
                                _ = v74
                            case 818:
                                v4 = v4
                            }
                        }
                    }
                case "":
                    for fi75 := 0; fi75 < 12; fi75++ {
                        var v76 uint32 = 4294967295
                        var v77 uint16 = 122
                        _ = fi75
                        _ = v76
                        _ = v77
                    }
                }
                for fi78 := 0; fi78 < 7; fi78++ {
                    v79 := GenS8{F0: 2147483647}
                    var v80 int = (^(fi78))
                    _ = fi78
                    _ = v79
                    _ = v80
                }
                v2[(fi70 + (v3 << ((0) & 63)))] = v3
                _ = fi70
            }
        case 128:
            var v81 int32 = -2147483648
            _ = v81
        default:
            v4 = v4
        }
        if ((1e10 - 0.0) >= 0.0) {
            v5 = v5
        } else {
            var v82 uint64 = 18446744073709551615
            v83 := GenS9{F0: 128}
            _ = v82
            _ = v83
        }
        v84 := GenS10{F0: 500, F1: 965}
        _ = v84
    } else {
        v1 -= v1
        fmt.Println(v1)
    }
    var v85 int64 = v1
    for fi86 := 0; fi86 < 10; fi86++ {
        if ((738 > 0) && true) {
            var v87 string = v4
            fmt.Println(v4, fi86, fi86)
            switch v87 {
            case "foo bar":
                var v88 uint = 315
                _ = v88
            default:
                for fi89 := 0; fi89 < 16; fi89++ {
                    var v90 int = (^(v3))
                    _ = fi89
                    _ = v90
                }
            }
            _ = v87
        }
        var v91 int32 = 0
        _ = fi86
        _ = v91
    }
    if false {
        v1 = v85
    }
    if false {
        v92 := GenS1{F0: genFn1(-1, 500, 100)}
        v4 = ((v4 + "") + v4)
        v5 = 1.5
        _ = v92
    } else {
        for fi93 := 0; fi93 < 15; fi93++ {
            var v94 int64 = (v85 & (v1 * v1))
            _ = fi93
            _ = v94
        }
        v4 = ""
        v3 = 2147483647
    }
    v3 += v3
    var v95 uint64 = 707
    v4 = v4
    if (!false) {
        v1 += -9223372036854775808
    }
    for fi96 := 0; fi96 < 5; fi96++ {
        if len(v6) > 0 {
            v6[0] = v3
        }
        v97 := GenS9{F0: 1}
        _ = len(v6)
        _ = fi96
        _ = v97
    }
    if true {
        if ((v85 | 9223372036854775807) >= (v85 & 0)) {
            switch v4 {
            case "\n\t":
                switch "" {
                case "hello":
                    var v98 int = v3
                    _ = v98
                case "":
                    var v99 int = (^(v3))
                    _ = v99
                default:
                    for fi100 := 0; fi100 < 4; fi100++ {
                        v1 = (v1 - (v85 - v85))
                        _ = fi100
                    }
                }
            case "":
                if false {
                    v2[(v3 - 9223372036854775807)] = v3
                    if (2 >= 42) {
                        for fi101 := 0; fi101 < 18; fi101++ {
                            v5 -= -2.25
                            if true {
                                fmt.Println(v4, v85, v3)
                            }
                            _ = fi101
                        }
                        v5 -= v5
                    } else {
                        if true {
                            if false {
                                if true {
                                    for fi102 := 0; fi102 < 9; fi102++ {
                                        v1 += (v85 & v1)
                                        var v103 int32 = -470
                                        _ = fi102
                                        _ = v103
                                    }
                                    if true {
                                        delete(v2, (v3 % (((v3 % ((9223372036854775807) | 1))) | 1)))
                                        for fi104 := 0; fi104 < 9; fi104++ {
                                            v4 = v4
                                            if true {
                                                for fi105 := 0; fi105 < 13; fi105++ {
                                                    v95 -= (v95 + (^(v95)))
                                                    v106 := GenS5{F0: v1, F1: 2}
                                                    var v107 float32 = (1.5 * (-2.25 / v5))
                                                    _ = fi105
                                                    _ = v106
                                                    _ = v107
                                                }
                                            } else {
                                                v4 = "a"
                                            }
                                            fmt.Println(fi104)
                                            _ = fi104
                                        }
                                    }
                                    v108 := GenS10{F0: 32767, F1: 42}
                                    _ = v108
                                } else {
                                    v109 := GenS11{F0: -440, F1: -842}
                                    v110 := GenS3{F0: v4}
                                    v1 -= (v85 >> ((65535) & 63))
                                    _ = v109
                                    _ = v110
                                }
                                switch v3 {
                                case -1:
                                    for fi111 := 0; fi111 < 9; fi111++ {
                                        v1 -= v85
                                        _ = fi111
                                    }
                                default:
                                    if false {
                                        for fi112 := 0; fi112 < 16; fi112++ {
                                            switch "hello" {
                                            case "":
                                                v113 := v2[v3]
                                                _ = v113
                                            default:
                                                var v114 int32 = -173
                                                _ = v114
                                            }
                                            _ = fi112
                                        }
                                        _ = len(v6)
                                    }
                                }
                            } else {
                                delete(v2, int(v95))
                            }
                        } else {
                            var v115 uint8 = 127
                            _ = v115
                        }
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                    v3 -= 128
                }
            }
            var v116 int32 = 2147483647
            v85 = (v1 ^ -1)
            _ = v116
        } else {
            for fi117 := 0; fi117 < 2; fi117++ {
                for fi118 := 0; fi118 < 4; fi118++ {
                    var v119 float64 = (1e10 + ((float64((v5 - ((v5 + (0.1 - v5)) / v5))) - -2.25) - float64((v5 / (((((v5 + 0.1) - v5) * ((3.14159 + 1.5) * float32(0.1))) + (((v5 / ((v5 + v5) - (v5 / 1e10))) * v5) / v5)) * (0.1 - ((v5 - (v5 / ((1e10 + 1.5) / 1e10))) + -2.25)))))))
                    _ = fi118
                    _ = v119
                }
                v120 := GenS12{F0: 4294967295, F1: 255}
                for fi121 := 0; fi121 < 8; fi121++ {
                    if false {
                        if (1 <= 282) {
                            fmt.Println(v1, v3)
                            var v122 int16 = -1
                            _ = v122
                        }
                        delete(v2, (fi121 - genFn0((v4 == v4))))
                    } else {
                        v123 := GenS10{F0: 1000, F1: -2147483648}
                        v4 = "hello"
                        fmt.Println(v3, v4, v5)
                        v4 = "a"
                        fmt.Println(v95, v85, fi117)
                        v124 := GenS3{F0: v4}
                        if false {
                            var v125 uint32 = 65535
                            _ = v125
                        }
                        _ = v123
                        _ = v124
                    }
                    fmt.Println(v5, v3, v3)
                    var v126 bool = (!(114 <= 0))
                    _ = fi121
                    _ = v126
                }
                _ = fi117
                _ = v120
            }
        }
        var v127 uint64 = v95
        _ = v127
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v85:", v85)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v85
    _ = v95
}
