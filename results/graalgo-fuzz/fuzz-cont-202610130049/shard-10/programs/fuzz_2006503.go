package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int8; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint32; F1 uint }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS9 struct { F0 bool; F1 uint32; F2 bool }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint32; F1 uint8; F2 string }
type GenS12 struct { F0 int32; F1 uint16 }
type GenS13 struct { F0 int32; F1 int; F2 bool }
func genFn0(p0 int64, p1 uint) int8 {
    var v6 int16 = -32768
    v6 -= -32768
    var v7 bool = true
    p1 += 0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -128
}
func genFn1(p0 uint32) int8 {
    p0 -= (p0 << ((0) & 63))
    for fi8 := 0; fi8 < 6; fi8++ {
        if true {
            p0 = p0
            switch fi8 {
            case -2147483648:
                v9 := GenS1{F0: 42, F1: 0}
                _ = v9
            case 1:
                v10 := GenS1{F0: 1, F1: 2147483647}
                _ = v10
            case -9223372036854775808:
                if (128 >= 127) {
                    v11 := []int{}
                    switch "" {
                    case "":
                        p0 += p0
                    case "a":
                        for fi12 := 0; fi12 < 16; fi12++ {
                            var v13 uint = 0
                            _ = fi12
                            _ = v13
                        }
                    default:
                        p0 += (^(p0))
                    }
                    if false {
                        switch "hello" {
                        case "foo bar":
                            var v14 int32 = 383
                            _ = v14
                        case "hello":
                            fmt.Println(fi8, fi8)
                        case "\n\t":
                            v15 := GenS1{F0: -9223372036854775808, F1: 0}
                            _ = v15
                        default:
                            p0 = p0
                        }
                        if true {
                            p0 += (p0 - p0)
                        } else {
                            fmt.Println(p0)
                            for fi16 := 0; fi16 < 8; fi16++ {
                                if (65535 != 697) {
                                    v17 := GenS1{F0: 459, F1: 42}
                                    v18 := GenS2{F0: 2}
                                    _ = v17
                                    _ = v18
                                }
                                _ = fi16
                            }
                            var v19 uint8 = 2
                            _ = v19
                        }
                    } else {
                        v20 := GenS2{F0: p0}
                        _ = v20
                    }
                    _ = v11
                }
            }
        }
        for fi21 := 0; fi21 < 2; fi21++ {
            switch fi8 {
            case 554:
                v22 := GenS2{F0: (^(p0))}
                _ = v22
            case -184:
                if (-1 > -880) {
                    var v23 uint32 = 4294967295
                    _ = v23
                }
            case -481:
                switch (-(fi8)) {
                case 9223372036854775807:
                    var v24 int = fi8
                    _ = v24
                case -949:
                    if (42 == 1) {
                        switch "" {
                        case "hello":
                            if false {
                                v25 := GenS1{F0: 1, F1: -2147483648}
                                var v26 bool = false
                                _ = v25
                                _ = v26
                            } else {
                                if false {
                                    var v27 int = -671
                                    _ = v27
                                }
                            }
                        case "foo bar":
                            p0 = p0
                            p0 = (p0 << ((2) & 63))
                            if ((!false) || false) {
                                p0 = (p0 << ((496) & 63))
                                var v28 float32 = 3.14159
                                v29 := []int{fi21, (fi8 / (((fi21 + (fi21 << ((65535) & 63)))) | 1))}
                                _ = v28
                                _ = v29
                            }
                        }
                    }
                default:
                    var v30 uint32 = (p0 & 65535)
                    _ = v30
                }
            default:
                fmt.Println(fi8, p0)
                v31 := GenS1{F0: 1, F1: -116}
                p0 += p0
                if true {
                    var v32 int = fi8
                    _ = v32
                }
                _ = v31
            }
            _ = fi21
        }
        _ = fi8
    }
    p0 -= (p0 - 536)
    _ = p0
    return 100
}

func main() {
    var v1 bool = false
    var v2 uint = 42
    v3 := []string{"\n\t"}
    var v4 float64 = float64((0.0 * (3.14159 / -2.25)))
    var v5 int16 = 42
    for fi33 := 0; fi33 < 13; fi33++ {
        var v34 string = "foo bar"
        _ = fi33
        _ = v34
    }
    var v35 int32 = 0
    for fi36 := 0; fi36 < 8; fi36++ {
        var v37 float32 = 1.5
        var v38 uint = (v2 | v2)
        _ = fi36
        _ = v37
        _ = v38
    }
    v39 := GenS1{F0: 42, F1: v35}
    _ = len(v3)
    v2 = 5
    for fi40 := 0; fi40 < 9; fi40++ {
        fmt.Println(v5)
        _ = fi40
    }
    fmt.Println(v35)
    for fi41 := 0; fi41 < 14; fi41++ {
        v4 = 0.1
        var v42 float64 = 3.14159
        _ = fi41
        _ = v42
    }
    var v43 int64 = 1
    for fi44 := 0; fi44 < 18; fi44++ {
        var v45 string = ("\n\t" + "")
        _ = fi44
        _ = v45
    }
    switch 667 {
    case 1:
        switch 42 {
        case 4294967295:
            v1 = (!v1)
        case 32767:
            fmt.Println(v5, v5)
        case -167:
            if ((0.0 > 1.5) && false) {
                if false {
                    var v46 int64 = (v43 >> ((65535) & 63))
                    _ = v46
                } else {
                    _ = len(v3)
                    var v47 int8 = -1
                    _ = v47
                }
                for fi48 := 0; fi48 < 13; fi48++ {
                    v49 := GenS3{F0: 127, F1: true, F2: 0}
                    _ = fi48
                    _ = v49
                }
            } else {
                var v50 float64 = v4
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v50
            }
        }
    default:
        if true {
            v35 += (v35 >> ((v2) & 63))
            if ((((v1 || false) || (!v1)) || (false || v1)) || false) {
                switch "hello" {
                case "hello":
                    var v51 int32 = v35
                    _ = v51
                case "":
                    v1 = (42 > -128)
                }
            } else {
                v4 = v4
                v43 = (v43 / (((v43 - v43)) | 1))
            }
        } else {
            v52 := GenS4{F0: 32767, F1: -2147483648}
            switch 2147483647 {
            case 4294967295:
                v53 := GenS5{F0: (v43 & v43)}
                _ = v53
            default:
                if len(v3) > 0 {
                    v3[0] = ""
                }
            }
            v43 -= (v43 % ((v43) | 1))
            _ = v52
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 -= 1
    fmt.Println(v2, v5, v4)
    v4 += v4
    if v1 {
        var v54 uint8 = 1
        if len(v3) > 0 {
            v3[0] = ""
        }
        if true {
            v2 = v2
            switch (("\n\t" + ("hello" + "a")) + "\n\t") {
            case "":
                v4 += -2.25
            default:
                if len(v3) > 0 {
                    v3[0] = ("a" + "")
                }
            }
            for fi55 := 0; fi55 < 10; fi55++ {
                v3 = append(v3, "a")
                if v1 {
                    if (!(-128 <= 100)) {
                        v54 = 127
                        var v56 int8 = int8(v54)
                        _ = v56
                    } else {
                        for fi57 := 0; fi57 < 13; fi57++ {
                            if true {
                                var v58 uint32 = 42
                                _ = v58
                            } else {
                                if ((v43 <= -9223372036854775808) || (32767 != 2)) {
                                    switch "" {
                                    case "\n\t":
                                        v59 := GenS6{F0: 42, F1: v2}
                                        _ = v59
                                    case "":
                                        v43 = (v43 + 9223372036854775807)
                                    }
                                    var v60 bool = v1
                                    _ = v60
                                } else {
                                    var v61 int16 = 0
                                    if v1 {
                                        var v62 string = "hello"
                                        _ = v62
                                    } else {
                                        fmt.Println(v43, fi57, v2)
                                    }
                                    v1 = true
                                    _ = v61
                                }
                                fmt.Println(v43, fi57)
                            }
                            v4 = ((1.5 - ((0.0 * (0.0 * (v4 - v4))) - (v4 * 0.0))) + 1e10)
                            v43 = v43
                            if ((v2 * v2) > 42) {
                                var v63 int64 = (v43 << (((^(v2))) & 63))
                                v35 = (-(v35))
                                _ = v63
                            }
                            _ = fi57
                        }
                        v64 := GenS6{F0: 2147483648, F1: 18446744073709551615}
                        _ = v64
                    }
                    v65 := GenS7{F0: (v5 & -32768)}
                    v66 := GenS8{F0: 108, F1: 65535, F2: 365}
                    _ = v65
                    _ = v66
                } else {
                    for fi67 := 0; fi67 < 9; fi67++ {
                        var v68 float32 = 1e10
                        _ = fi67
                        _ = v68
                    }
                    for fi69 := 0; fi69 < 5; fi69++ {
                        var v70 int32 = 0
                        if ((v54 + v54) > 58) {
                            v5 = -307
                            if (((-2.25 * (0.1 * float32(1.5))) * 0.0) != (-2.25 / (1e10 + float32(v4)))) {
                                var v71 uint = (^(v2))
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v71
                            } else {
                                if (!v1) {
                                    if (!v1) {
                                        var v72 string = "hello"
                                        _ = v72
                                    }
                                    var v73 uint8 = (v54 - (v54 << ((v2) & 63)))
                                    v1 = v1
                                    _ = v73
                                } else {
                                    fmt.Println(v54, v70)
                                    fmt.Println(v1)
                                }
                                if (!false) {
                                    fmt.Println(v70)
                                } else {
                                    var v74 uint8 = 2
                                    _ = v74
                                }
                                if v1 {
                                    v54 -= v54
                                    switch fi69 {
                                    case -1:
                                        var v75 uint16 = 66
                                        _ = v75
                                    case 32767:
                                        v76 := GenS6{F0: 2, F1: (v2 ^ 250)}
                                        _ = v76
                                    case -9223372036854775808:
                                        var v77 uint32 = 4294967295
                                        _ = v77
                                    default:
                                        var v78 int = (fi55 / ((-833) | 1))
                                        _ = v78
                                    }
                                    v5 = v5
                                } else {
                                    v79 := GenS8{F0: genFn1(654), F1: 938, F2: 1}
                                    v80 := GenS8{F0: -124, F1: 583, F2: 244}
                                    _ = v79
                                    _ = v80
                                }
                                v5 = (v5 / ((v5) | 1))
                            }
                            fmt.Println(v43, v43, v2)
                        } else {
                            if false {
                                var v81 int8 = -1
                                var v82 bool = (true || (v43 >= -9223372036854775808))
                                if v82 {
                                    var v83 uint64 = 42
                                    var v84 int32 = (v35 / ((v35) | 1))
                                    _ = v83
                                    _ = v84
                                } else {
                                    var v85 int8 = (^(v81))
                                    fmt.Println(fi55)
                                    _ = v85
                                }
                                _ = v81
                                _ = v82
                            } else {
                                if v1 {
                                    v86 := GenS9{F0: (1 == 64), F1: 2147483647, F2: (v1 || true)}
                                    v43 = (v43 * v43)
                                    _ = v86
                                } else {
                                    if len(v3) > 0 {
                                        v3[0] = "\n\t"
                                    }
                                    var v87 bool = ((-(v35)) >= 2147483647)
                                    if false {
                                        switch -679 {
                                        case 391:
                                            if false {
                                                v88 := GenS6{F0: 2147483648, F1: v2}
                                                v89 := GenS9{F0: true, F1: 2147483648, F2: v1}
                                                _ = v88
                                                _ = v89
                                            } else {
                                                var v90 float32 = (0.0 - -2.25)
                                                v2 -= (v2 / (((v2 - (v2 - v2))) | 1))
                                                switch "hello" {
                                                case "\n\t":
                                                    if v1 {
                                                        if (103 >= -128) {
                                                            v91 := GenS10{F0: v2}
                                                            if (0 != -122) {
                                                                var v92 uint8 = v54
                                                                if (v1 || ((^(v70)) != (v70 << ((52) & 63)))) {
                                                                    if ((4294967295 <= 2) || (2 >= 1)) {
                                                                        if len(v3) > 0 {
                                                                            v3[0] = "\n\t"
                                                                        }
                                                                        var v93 uint32 = 1
                                                                        v43 = v43
                                                                        var v94 int32 = (v70 - v70)
                                                                        if v1 {
                                                                            switch (fi55 ^ (fi69 << ((65535) & 63))) {
                                                                            case 42:
                                                                                v95 := GenS8{F0: -128, F1: 65535, F2: 0}
                                                                                _ = v95
                                                                            }
                                                                            if false {
                                                                                var v96 int16 = v5
                                                                                _ = v96
                                                                            }
                                                                            v97 := GenS10{F0: v2}
                                                                            _ = v97
                                                                        }
                                                                        fmt.Println(fi55, v2)
                                                                        v5 += v5
                                                                        switch "\n\t" {
                                                                        case "a":
                                                                            if len(v3) > 0 {
                                                                                v3[0] = "a"
                                                                            }
                                                                        case "hello":
                                                                            var v98 string = ("hello" + "\n\t")
                                                                            _ = v98
                                                                        default:
                                                                            if (!true) {
                                                                                v3 = append(v3, "")
                                                                                switch (fi69 + -9223372036854775808) {
                                                                                case 42:
                                                                                    if len(v3) > 0 {
                                                                                        v3[0] = "foo bar"
                                                                                    }
                                                                                case -2147483648:
                                                                                    var v99 int8 = 42
                                                                                    _ = v99
                                                                                }
                                                                            } else {
                                                                                v87 = true
                                                                            }
                                                                            v93 = 636
                                                                        }
                                                                        _ = v93
                                                                        _ = v94
                                                                    } else {
                                                                        var v100 float64 = v4
                                                                        _ = v100
                                                                    }
                                                                    switch fi69 {
                                                                    case 128:
                                                                        var v101 uint = v2
                                                                        _ = v101
                                                                    case 0:
                                                                        if v87 {
                                                                            fmt.Println(v92, v2)
                                                                        }
                                                                    }
                                                                }
                                                                _ = v92
                                                            } else {
                                                                var v102 int = -2147483648
                                                                switch (fi69 * 1) {
                                                                case 322:
                                                                    var v103 int16 = v5
                                                                    _ = v103
                                                                default:
                                                                    v3 = append(v3, ("hello" + "foo bar"))
                                                                }
                                                                if len(v3) > 0 {
                                                                    v3[0] = (("foo bar" + "a") + (("\n\t" + "") + "a"))
                                                                }
                                                                switch ("foo bar" + "hello") {
                                                                case "":
                                                                    if (944 != 255) {
                                                                        fmt.Println(v87)
                                                                    } else {
                                                                        var v104 int16 = 1
                                                                        if (v35 == (v70 / ((v70) | 1))) {
                                                                            fmt.Println(v70)
                                                                            switch v102 {
                                                                            case -2147483648:
                                                                                v105 := GenS7{F0: v5}
                                                                                _ = v105
                                                                            case 952:
                                                                                v106 := GenS6{F0: 2147483647, F1: v2}
                                                                                _ = v106
                                                                            default:
                                                                                v87 = v1
                                                                            }
                                                                        }
                                                                        _ = v104
                                                                    }
                                                                }
                                                                _ = v102
                                                            }
                                                            fmt.Println(v4)
                                                            switch ((("foo bar" + "\n\t") + ("foo bar" + (("" + (("" + ("foo bar" + ("foo bar" + "hello"))) + ("a" + "foo bar"))) + ((("\n\t" + "a") + "\n\t") + ("" + ""))))) + "foo bar") {
                                                            case "":
                                                                switch "" {
                                                                case "\n\t":
                                                                    var v107 int16 = 118
                                                                    v5 -= 990
                                                                    _ = v107
                                                                case "a":
                                                                    v1 = (255 <= (v54 & v54))
                                                                }
                                                            }
                                                            _ = v91
                                                        }
                                                        if len(v3) > 0 {
                                                            v3[0] = ""
                                                        }
                                                        v35 += v70
                                                    } else {
                                                        v43 += v43
                                                        var v108 uint8 = v54
                                                        v43 = (^(v43))
                                                        var v109 float64 = float64(v90)
                                                        _ = v108
                                                        _ = v109
                                                    }
                                                case "":
                                                    var v110 bool = ((fi69 % (((fi69 % ((int(v5)) | 1))) | 1)) < int(v54))
                                                    _ = v110
                                                default:
                                                    switch "a" {
                                                    case "\n\t":
                                                        _ = len(v3)
                                                    default:
                                                        var v111 float32 = float32(v4)
                                                        _ = v111
                                                    }
                                                }
                                                _ = v90
                                            }
                                        case 839:
                                            v43 = v43
                                        case 32768:
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        }
                                        if true {
                                            var v112 bool = (!(((0.0 / (float32(0.1) * 0.1)) * (-2.25 + 0.0)) < (3.14159 / (float32(((((3.14159 - v4) * (((1.5 / v4) * (0.1 + v4)) / 1.0)) / 1.0) * v4)) - float32(1e10)))))
                                            switch ("\n\t" + ("\n\t" + ("" + "hello"))) {
                                            case "foo bar":
                                                v113 := GenS8{F0: int8(v43), F1: 709, F2: 2}
                                                _ = v113
                                            case "\n\t":
                                                v35 = v70
                                            case "hello":
                                                var v114 int = fi69
                                                _ = v114
                                            }
                                            if false {
                                                var v115 string = "a"
                                                v116 := GenS4{F0: 32767, F1: (fi69 % (((fi55 + (fi55 - (fi69 % ((127) | 1))))) | 1))}
                                                var v117 uint = (v2 / (((v2 / (((v2 ^ 4294967295)) | 1))) | 1))
                                                v4 = float64((0.0 - 0.1))
                                                v2 = (v117 | v2)
                                                var v118 uint32 = 1
                                                _ = v115
                                                _ = v116
                                                _ = v117
                                                _ = v118
                                            } else {
                                                v87 = v1
                                                fmt.Println(v87, v87)
                                            }
                                            _ = v112
                                        } else {
                                            v70 -= v35
                                            switch "a" {
                                            case "":
                                                if ((v87 || (255 == 0)) && v1) {
                                                    var v119 bool = false
                                                    _ = v119
                                                } else {
                                                    v70 = (v35 >> ((839) & 63))
                                                }
                                            case "hello":
                                                var v120 uint8 = v54
                                                _ = v120
                                            default:
                                                fmt.Println(v43)
                                            }
                                            if len(v3) > 0 {
                                                v3[0] = "foo bar"
                                            }
                                        }
                                    } else {
                                        fmt.Println(fi69, fi55)
                                        var v121 float64 = 1.5
                                        _ = v121
                                    }
                                    v35 = v70
                                    if false {
                                        if (2147483647 <= uint32(v54)) {
                                            var v122 uint = (^(v2))
                                            fmt.Println(v122, v70)
                                            _ = v122
                                        } else {
                                            var v123 uint = v2
                                            _ = v123
                                        }
                                        fmt.Println(fi69, v5)
                                        v124 := GenS1{F0: (v43 << (((v2 | 42)) & 63)), F1: v70}
                                        _ = v124
                                    }
                                    _ = v87
                                }
                                v3 = append(v3, "hello")
                            }
                            v1 = v1
                            var v125 int64 = 9223372036854775807
                            _ = v125
                        }
                        var v126 bool = v1
                        if len(v3) > 0 {
                            v3[0] = "a"
                        }
                        if len(v3) > 0 {
                            v3[0] = (("hello" + ("a" + "")) + (("\n\t" + "foo bar") + ("\n\t" + "hello")))
                        }
                        v127 := GenS6{F0: 117, F1: (v2 + (v2 & 750))}
                        switch "\n\t" {
                        case "foo bar":
                            var v128 float32 = (((0.0 + ((0.1 / 3.14159) / 1.5)) / (((-2.25 - 1.5) * 1.5) + 1e10)) / (0.1 * (3.14159 * 1e10)))
                            _ = v128
                        }
                        switch ("\n\t" + "foo bar") {
                        case "a":
                            switch fi55 {
                            case 2147483647:
                                switch "hello" {
                                case "":
                                    fmt.Println(v54)
                                }
                            case 697:
                                var v129 int8 = 18
                                _ = v129
                            default:
                                if ((v70 & v35) >= v35) {
                                    if len(v3) > 0 {
                                        v3[0] = "foo bar"
                                    }
                                } else {
                                    if len(v3) > 0 {
                                        v3[0] = "a"
                                    }
                                    v130 := GenS11{F0: 4294967295, F1: 0, F2: "\n\t"}
                                    _ = v130
                                }
                            }
                        }
                        v3 = append(v3, ("" + ("" + "\n\t")))
                        _ = fi69
                        _ = v126
                        _ = v127
                        _ = v70
                    }
                    for fi131 := 0; fi131 < 12; fi131++ {
                        v132 := GenS12{F0: (-(v35)), F1: 680}
                        if v1 {
                            switch fi55 {
                            case 4294967295:
                                if (!(v2 <= (v2 | 65535))) {
                                    v3 = append(v3, ("a" + "a"))
                                } else {
                                    v5 -= 951
                                    v2 += v2
                                }
                            case 699:
                                if v1 {
                                    switch int(v2) {
                                    case 4294967295:
                                        var v133 string = "hello"
                                        v134 := GenS13{F0: (v35 / ((v35) | 1)), F1: int(v5), F2: ((v35 & v35) <= int32(v5))}
                                        if v1 {
                                            v54 = v54
                                            var v135 uint8 = 1
                                            _ = v135
                                        }
                                        if v1 {
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            if true {
                                                v35 = (v35 & v35)
                                                _ = len(v3)
                                            } else {
                                                v1 = true
                                                if ((fi55 >> (((v2 / ((v2) | 1))) & 63)) < -9223372036854775808) {
                                                    if v1 {
                                                        v4 -= (v4 / v4)
                                                        if ((4294967295 > 1) || ((v35 - v35) != int32(v5))) {
                                                            var v136 uint8 = (^(v54))
                                                            switch fi131 {
                                                            case 0:
                                                                v136 = (v54 ^ v54)
                                                                var v137 int32 = -2147483648
                                                                _ = v137
                                                            case -1:
                                                                v138 := GenS4{F0: 13, F1: fi131}
                                                                _ = v138
                                                            case -2147483648:
                                                                fmt.Println(v133, v54, v5)
                                                            default:
                                                                var v139 uint8 = (^(v54))
                                                                _ = v139
                                                            }
                                                            var v140 uint8 = v136
                                                            _ = v136
                                                            _ = v140
                                                        } else {
                                                            var v141 float32 = 3.14159
                                                            fmt.Println(v141, v2)
                                                            _ = v141
                                                        }
                                                    }
                                                    v142 := GenS11{F0: 951, F1: (v54 + v54), F2: (("foo bar" + (v133 + (v133 + ("hello" + (("hello" + v133) + ("hello" + "hello")))))) + "foo bar")}
                                                    fmt.Println(fi55, v133)
                                                    if (!v1) {
                                                        switch fi131 {
                                                        case 0:
                                                            v2 += v2
                                                        case -9223372036854775808:
                                                            fmt.Println(v1)
                                                            v4 = (3.14159 + v4)
                                                        }
                                                        v3 = append(v3, v133)
                                                    }
                                                    _ = v142
                                                }
                                                var v143 string = "hello"
                                                _ = v143
                                            }
                                            v4 += (v4 / v4)
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = ((v133 + "foo bar") + v133)
                                        }
                                        _ = v133
                                        _ = v134
                                    }
                                }
                            case 2147483647:
                                v144 := GenS12{F0: (v35 >> (((v2 - v2)) & 63)), F1: 2}
                                _ = v144
                            }
                            v2 -= v2
                            v1 = v1
                        } else {
                            v54 = v54
                            var v145 uint8 = v54
                            v145 = (v145 + v54)
                            _ = v145
                        }
                        var v146 uint16 = 255
                        _ = fi131
                        _ = v132
                        _ = v146
                    }
                }
                var v147 float64 = float64(1.5)
                _ = fi55
                _ = v147
            }
        } else {
            var v148 int = 127
            _ = v148
        }
        _ = v54
    }
    v149 := GenS6{F0: 2147483647, F1: (v2 + 42)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v149
    _ = v2
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v43
    _ = v5
}
