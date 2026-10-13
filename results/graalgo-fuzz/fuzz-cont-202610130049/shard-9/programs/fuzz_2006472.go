package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 string }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16; F1 uint64 }
type GenS6 struct { F0 uint; F1 int32 }
type GenS7 struct { F0 int16; F1 int32 }
type GenS8 struct { F0 uint64; F1 string; F2 bool }
type GenS9 struct { F0 uint64; F1 int; F2 int16 }
type GenS10 struct { F0 uint8; F1 int32; F2 uint }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint32; F1 int8 }
type GenS14 struct { F0 int8; F1 int32; F2 int8 }
type GenS15 struct { F0 string; F1 uint; F2 uint }
func genFn0(p0 uint8, p1 int) uint16 {
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 string = ("hello" + "\n\t")
        v7 := GenS1{F0: "foo bar", F1: 65535, F2: v6}
        _ = fi5
        _ = v6
        _ = v7
    }
    if false {
        var v8 uint32 = 272
        if false {
            fmt.Println(p0)
            var v9 int32 = 951
            _ = v9
        } else {
            var v10 uint32 = 152
            _ = v10
        }
        switch "a" {
        case "":
            switch "\n\t" {
            case "":
                fmt.Println(v8, p0, p1)
            default:
                var v11 int8 = int8(p1)
                _ = v11
            }
        case "foo bar":
            fmt.Println(p1, p0, p1)
        case "a":
            var v12 float32 = -2.25
            _ = v12
        }
        _ = v8
    }
    p0 -= (p0 & (p0 - p0))
    for fi13 := 0; fi13 < 19; fi13++ {
        var v14 int32 = 1
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    return 260
}

func main() {
    v1 := []uint32{}
    var v2 string = ("\n\t" + "hello")
    v3 := make(map[string]int)
    var v4 uint32 = 589
    switch 860 {
    case -636:
        var v15 int32 = -2147483648
        _ = v15
    case -1:
        v16 := v3["hello"]
        _ = v16
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        var v18 float32 = (1.5 + (float32(0.1) + 3.14159))
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        if len(v1) > 0 {
            v1[0] = v4
        }
        switch fi19 {
        case 32767:
            v20 := GenS1{F0: v2, F1: 1, F2: ((v2 + (v2 + v2)) + "")}
            _ = v20
        case 32768:
            switch "foo bar" {
            case "":
                var v21 float32 = (float32(0.1) / float32(0.1))
                _ = v21
            case "foo bar":
                for fi22 := 0; fi22 < 7; fi22++ {
                    switch ("foo bar" + v2) {
                    case "foo bar":
                        var v23 int32 = 9
                        _ = v23
                    default:
                        if (52 != 255) {
                            fmt.Println(v4, v4, fi19)
                        }
                    }
                    _ = fi22
                }
            default:
                fmt.Println(fi19)
            }
        }
        for fi24 := 0; fi24 < 11; fi24++ {
            switch fi19 {
            case 552:
                v25 := GenS1{F0: ((v2 + (v2 + (v2 + v2))) + v2), F1: 2, F2: v2}
                _ = v25
            default:
                if false {
                    if false {
                        var v26 float32 = ((1.5 - ((-2.25 / 1e10) + (1.5 - float32(-2.25)))) + ((0.1 + 1.5) / -2.25))
                        if false {
                            var v27 uint16 = 856
                            v28 := GenS2{F0: 32767, F1: 1}
                            _ = v27
                            _ = v28
                        } else {
                            v1 = append(v1, (v4 ^ 42))
                            if (42 == 204) {
                                switch v2 {
                                case "hello":
                                    var v29 bool = (true && (!(42 > 255)))
                                    _ = v29
                                case "":
                                    v30 := GenS2{F0: int16(fi24), F1: 100}
                                    _ = v30
                                case "a":
                                    var v31 int64 = -9223372036854775808
                                    _ = v31
                                }
                            }
                            var v32 int64 = 42
                            _ = v32
                        }
                        var v33 string = (("a" + (("" + (v2 + v2)) + (v2 + (((v2 + v2) + "\n\t") + (v2 + v2))))) + ("foo bar" + "\n\t"))
                        _ = v26
                        _ = v33
                    }
                    v34 := GenS3{F0: 613}
                    _ = v34
                } else {
                    if (0 < 1) {
                        var v35 uint16 = 32767
                        var v36 uint8 = 1
                        v3[""] = fi24
                        _ = v35
                        _ = v36
                    } else {
                        _ = len(v3)
                        v1 = append(v1, (v4 + 4294967295))
                        if false {
                            var v37 uint = 434
                            var v38 int = (^(fi19))
                            var v39 uint = v37
                            delete(v3, ((((("hello" + v2) + v2) + "") + v2) + v2))
                            v4 += 2147483648
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                    }
                }
            }
            _ = fi24
        }
        _ = fi19
    }
    switch v2 {
    case "hello":
        var v40 string = ("foo bar" + (v2 + ("a" + v2)))
        _ = v40
    case "":
        for fi41 := 0; fi41 < 2; fi41++ {
            var v42 float32 = (float32(((0.1 - 0.1) + 3.14159)) * (-2.25 / float32(3.14159)))
            _ = fi41
            _ = v42
        }
    default:
        v3[v2] = 127
    }
    if false {
        fmt.Println(v4, v4)
        for fi43 := 0; fi43 < 13; fi43++ {
            v4 = 974
            _ = len(v3)
            switch fi43 {
            case 9223372036854775807:
                var v44 uint64 = 2
                _ = v44
            case 0:
                if (fi43 < fi43) {
                    var v45 uint32 = (v4 - 1)
                    v2 = "foo bar"
                    var v46 int16 = 1000
                    _ = v45
                    _ = v46
                }
            case 127:
                var v47 int16 = -243
                _ = v47
            default:
                if false {
                    if (((-2.25 * 0.1) / 0.1) == 3.14159) {
                        if (955 != 2) {
                            for fi48 := 0; fi48 < 11; fi48++ {
                                v2 = (v2 + v2)
                                _ = fi48
                            }
                            var v49 uint32 = (^(v4))
                            fmt.Println(v2)
                            _ = v49
                        }
                        if true {
                            fmt.Println(fi43, fi43, v2)
                            var v50 float32 = (((0.0 / -2.25) - 1e10) / 3.14159)
                            _ = v50
                        }
                        if (9223372036854775807 < 1) {
                            var v51 int = 75
                            fmt.Println(v51, fi43)
                            _ = v51
                        }
                    } else {
                        var v52 int32 = -429
                        if (true || true) {
                            for fi53 := 0; fi53 < 15; fi53++ {
                                v54 := GenS4{F0: 28}
                                v2 = "a"
                                _ = fi53
                                _ = v54
                            }
                        } else {
                            for fi55 := 0; fi55 < 7; fi55++ {
                                v4 -= 978
                                switch "" {
                                case "\n\t":
                                    v52 += v52
                                case "a":
                                    switch v2 {
                                    case "hello":
                                        switch (fi43 - fi43) {
                                        case 1:
                                            switch v2 {
                                            case "":
                                                switch 2147483647 {
                                                case 42:
                                                    var v56 int64 = 9223372036854775807
                                                    _ = v56
                                                }
                                            case "hello":
                                                var v57 uint8 = 255
                                                _ = v57
                                            }
                                        case -970:
                                            switch "a" {
                                            case "hello":
                                                if (!false) {
                                                    v4 = 689
                                                    v2 = (v2 + "a")
                                                    v58 := GenS2{F0: -32768, F1: 59}
                                                    _ = v58
                                                }
                                            default:
                                                v52 += -110
                                            }
                                        case -888:
                                            v59 := GenS5{F0: 885, F1: 18446744073709551615}
                                            _ = v59
                                        default:
                                            switch (fi43 / ((572) | 1)) {
                                            case -9223372036854775808:
                                                var v60 uint64 = 4294967295
                                                _ = v60
                                            case 1:
                                                _ = len(v1)
                                            }
                                        }
                                    case "a":
                                        v1 = append(v1, (v4 % ((v4) | 1)))
                                        switch "foo bar" {
                                        case "\n\t":
                                            v52 -= (-(v52))
                                        case "foo bar":
                                            switch v2 {
                                            case "\n\t":
                                                switch fi55 {
                                                case 127:
                                                    fmt.Println(v4, v52, fi43)
                                                case 949:
                                                    var v61 int32 = v52
                                                    _ = v61
                                                }
                                            case "a":
                                                _ = len(v3)
                                                if (v4 == v4) {
                                                    if (42 < 1) {
                                                        switch ("" + v2) {
                                                        case "":
                                                            var v62 int32 = v52
                                                            _ = v62
                                                        case "hello":
                                                            switch "hello" {
                                                            case "":
                                                                v4 += v4
                                                            case "\n\t":
                                                                switch fi55 {
                                                                case 32768:
                                                                    v4 += 0
                                                                case -92:
                                                                    if (!true) {
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        v63 := GenS6{F0: 2, F1: v52}
                                                                        _ = v63
                                                                    }
                                                                }
                                                            default:
                                                                v64 := GenS5{F0: 42, F1: 18446744073709551615}
                                                                _ = v64
                                                            }
                                                        default:
                                                            var v65 uint8 = 255
                                                            _ = v65
                                                        }
                                                        if (-32768 < 1) {
                                                            switch 797 {
                                                            case 402:
                                                                var v66 uint64 = 492
                                                                _ = v66
                                                            default:
                                                                fmt.Println(v4)
                                                            }
                                                            v4 += v4
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        } else {
                                                            _ = len(v1)
                                                            if true {
                                                                v52 = (v52 * v52)
                                                                if (!(float32((float64(((0.1 * (1.5 * (0.1 * 0.1))) * 1e10)) * 0.1)) < (0.0 / float32(1e10)))) {
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                }
                                                            }
                                                            var v67 uint32 = 647
                                                            _ = v67
                                                        }
                                                        var v68 int32 = v52
                                                        _ = v68
                                                    } else {
                                                        var v69 int = fi55
                                                        _ = v69
                                                    }
                                                } else {
                                                    v4 = (v4 | 1)
                                                    var v70 int16 = 1000
                                                    _ = v70
                                                }
                                                fmt.Println(v52, v4)
                                            }
                                        case "":
                                            var v71 int64 = 42
                                            _ = v71
                                        }
                                    case "\n\t":
                                        v2 = ((v2 + "hello") + v2)
                                    }
                                }
                                v2 = v2
                                _ = fi55
                            }
                            var v72 int64 = 42
                            for fi73 := 0; fi73 < 4; fi73++ {
                                var v74 string = ("foo bar" + ("foo bar" + v2))
                                var v75 int8 = 127
                                v75 += -124
                                _ = fi73
                                _ = v74
                                _ = v75
                            }
                            _ = v72
                        }
                        v76 := GenS1{F0: ("" + v2), F1: 878, F2: v2}
                        _ = v52
                        _ = v76
                    }
                    if false {
                        for fi77 := 0; fi77 < 8; fi77++ {
                            v1 = append(v1, 1)
                            v2 = (((v2 + "foo bar") + v2) + (("" + "hello") + v2))
                            _ = fi77
                        }
                        var v78 int8 = 42
                        _ = len(v3)
                        _ = v78
                    } else {
                        if (1 == 2) {
                            if (32768 >= fi43) {
                                v2 = v2
                                var v79 bool = true
                                if true {
                                    switch v2 {
                                    case "\n\t":
                                        var v80 float64 = -2.25
                                        _ = v80
                                    case "hello":
                                        v81 := GenS7{F0: 32767, F1: 2147483647}
                                        _ = v81
                                    default:
                                        switch v2 {
                                        case "hello":
                                            v79 = ((v4 <= 2) || v79)
                                        case "\n\t":
                                            var v82 int = (^(fi43))
                                            _ = v82
                                        }
                                    }
                                    delete(v3, (v2 + "\n\t"))
                                    if v79 {
                                        v83 := GenS3{F0: 42}
                                        _ = v83
                                    } else {
                                        if v79 {
                                            v2 = ((v2 + "hello") + ("foo bar" + v2))
                                            var v84 int16 = 46
                                            v2 = ("\n\t" + ("\n\t" + v2))
                                            _ = v84
                                        } else {
                                            var v85 string = ""
                                            _ = v85
                                        }
                                        if v79 {
                                            var v86 uint16 = 130
                                            v3[("foo bar" + ("" + ((("" + v2) + "\n\t") + "\n\t")))] = fi43
                                            switch 315 {
                                            case 127:
                                                var v87 string = v2
                                                _ = v87
                                            case 1:
                                                switch v2 {
                                                case "foo bar":
                                                    if (v4 > 0) {
                                                        v1 = append(v1, v4)
                                                    }
                                                }
                                            default:
                                                v86 -= 2
                                            }
                                            _ = v86
                                        }
                                    }
                                }
                                _ = v79
                            }
                        }
                        v88 := GenS8{F0: 535, F1: "foo bar", F2: false}
                        v2 = ("a" + (v2 + (v2 + v2)))
                        _ = v88
                    }
                    if (42 < 74) {
                        _ = len(v3)
                        var v89 float32 = 1e10
                        v89 -= v89
                        _ = v89
                    } else {
                        _ = len(v3)
                    }
                } else {
                    for fi90 := 0; fi90 < 4; fi90++ {
                        if false {
                            if false {
                                switch v2 {
                                case "":
                                    var v91 int32 = 0
                                    _ = v91
                                default:
                                    if (false && false) {
                                        if true {
                                            var v92 int64 = 0
                                            if (true || (392 == 0)) {
                                                if (((float64((0.0 + -2.25)) + -2.25) * -2.25) >= 3.14159) {
                                                    if false {
                                                        v93 := GenS2{F0: -375, F1: -128}
                                                        v92 = (v92 >> ((0) & 63))
                                                        fmt.Println(v4)
                                                        _ = v93
                                                    }
                                                    v4 = 2
                                                }
                                                switch 32768 {
                                                case -287:
                                                    v94 := GenS5{F0: uint16(v4), F1: 1}
                                                    _ = v94
                                                case 659:
                                                    v92 = -9223372036854775808
                                                case -1:
                                                    var v95 int16 = -560
                                                    _ = v95
                                                default:
                                                    fmt.Println(fi90)
                                                    v96 := GenS9{F0: 0, F1: 42, F2: 380}
                                                    _ = v96
                                                }
                                            } else {
                                                var v97 uint = 0
                                                v98 := GenS10{F0: 42, F1: 697, F2: 2}
                                                switch 533 {
                                                case 624:
                                                    v99 := GenS6{F0: 14, F1: -447}
                                                    _ = v99
                                                }
                                                _ = v97
                                                _ = v98
                                            }
                                            if (false || (!true)) {
                                                if true {
                                                    v100 := GenS7{F0: 42, F1: 248}
                                                    var v101 int = (fi90 & 128)
                                                    _ = v100
                                                    _ = v101
                                                }
                                            }
                                            _ = v92
                                        }
                                    }
                                }
                            } else {
                                if (true || (1 >= 0)) {
                                    switch "" {
                                    case "foo bar":
                                        if (!(42 == 1000)) {
                                            v4 = v4
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            if true {
                                                v102 := v3[v2]
                                                switch ("hello" + "foo bar") {
                                                case "":
                                                    v2 = v2
                                                default:
                                                    fmt.Println(v4)
                                                }
                                                var v103 int32 = 569
                                                _ = v102
                                                _ = v103
                                            }
                                        }
                                    case "a":
                                        switch fi90 {
                                        case -2147483648:
                                            v2 = v2
                                        }
                                    default:
                                        v104 := GenS9{F0: 0, F1: fi43, F2: -32768}
                                        _ = v104
                                    }
                                }
                                var v105 uint = 468
                                var v106 uint16 = 369
                                _ = v105
                                _ = v106
                            }
                            var v107 int64 = -9223372036854775808
                            _ = v107
                        } else {
                            if (false && true) {
                                v4 = (v4 ^ 65535)
                                if (-511 >= -1) {
                                    var v108 int64 = -560
                                    _ = v108
                                } else {
                                    var v109 int16 = 0
                                    v1 = append(v1, (^(v4)))
                                    v4 -= 554
                                    v2 = v2
                                    _ = v109
                                }
                                var v110 float64 = 3.14159
                                _ = v110
                            }
                            switch 310 {
                            case -709:
                                v4 = v4
                                v4 += (v4 | 208)
                            case -434:
                                v3[""] = 1
                            default:
                                v2 = v2
                                var v111 int = fi90
                                v4 += (v4 / ((v4) | 1))
                                _ = v111
                            }
                        }
                        switch v2 {
                        case "foo bar":
                            switch "" {
                            case "\n\t":
                                switch (fi43 ^ fi43) {
                                case 42:
                                    v4 = (v4 | (v4 ^ 2147483647))
                                case 32768:
                                    v112 := GenS11{F0: 127}
                                    _ = v112
                                case -39:
                                    v113 := GenS12{F0: (685 >= 697)}
                                    _ = v113
                                default:
                                    v114 := GenS12{F0: false}
                                    v115 := GenS7{F0: -1, F1: -307}
                                    _ = v114
                                    _ = v115
                                }
                            case "foo bar":
                                fmt.Println(v4, fi43, fi90)
                            }
                        case "a":
                            if len(v1) > 0 {
                                v1[0] = (v4 - 2)
                            }
                            v4 = (v4 | 2147483647)
                        default:
                            if true {
                                v2 = (v2 + (((v2 + ("" + ((((((v2 + v2) + v2) + "hello") + ((("foo bar" + v2) + (v2 + v2)) + (v2 + v2))) + v2) + v2))) + ("\n\t" + "hello")) + "foo bar"))
                                fmt.Println(v4)
                            }
                        }
                        if true {
                            var v116 uint = uint(v4)
                            if (1.5 != ((0.1 / 1.0) + 0.0)) {
                                v117 := GenS13{F0: (v4 / (((v4 >> ((v116) & 63))) | 1)), F1: 100}
                                _ = v117
                            }
                            _ = v116
                        } else {
                            if ((4294967295 <= 33) && (false && (0.1 >= 1e10))) {
                                var v118 uint64 = 4294967295
                                _ = v118
                            }
                        }
                        _ = fi90
                    }
                }
            }
            _ = fi43
        }
        if (127 >= -1) {
            var v119 int = -636
            if len(v1) > 0 {
                v1[0] = v4
            }
            _ = v119
        }
    }
    var v120 uint8 = 128
    for fi121 := 0; fi121 < 4; fi121++ {
        v2 = (v2 + (v2 + v2))
        var v122 float64 = (float64((((0.1 * 1.5) - (-2.25 * 1e10)) / float32(float64((0.1 * ((((0.0 * -2.25) / 0.1) + 3.14159) * 3.14159)))))) * (3.14159 / (0.1 - 1.5)))
        for fi123 := 0; fi123 < 17; fi123++ {
            var v124 uint64 = 143
            switch 42 {
            case -814:
                if len(v1) > 0 {
                    v1[0] = uint32(v120)
                }
            default:
                v125 := GenS14{F0: 127, F1: 42, F2: 127}
                _ = v125
            }
            v120 = 247
            _ = fi123
            _ = v124
        }
        _ = fi121
        _ = v122
    }
    if false {
        var v126 string = (v2 + ((v2 + "hello") + "a"))
        _ = v126
    } else {
        v4 += (^(v4))
        v4 += v4
    }
    if (100 != 30) {
        var v127 string = (v2 + "\n\t")
        _ = v127
    } else {
        v128 := GenS15{F0: "\n\t", F1: 255, F2: 18446744073709551615}
        _ = v128
    }
    v129 := GenS8{F0: 18446744073709551615, F1: "a", F2: (2147483647 <= 128)}
    _ = len(v3)
    var v130 bool = true
    var v131 int64 = 9223372036854775807
    var v132 int8 = 127
    for fi133 := 0; fi133 < 5; fi133++ {
        v134 := GenS14{F0: (v132 * v132), F1: -168, F2: (v132 & v132)}
        _ = fi133
        _ = v134
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v120:", v120)
    fmt.Println("v130:", v130)
    fmt.Println("v131:", v131)
    fmt.Println("v132:", v132)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v120
    _ = v129
    _ = v130
    _ = v131
    _ = v132
    _ = v2
    _ = v3
    _ = v4
}
