package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 int32; F2 int }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32; F1 uint8 }
type GenS8 struct { F0 int8; F1 string; F2 uint }
type GenS9 struct { F0 uint64; F1 uint8 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int8 }
type GenS12 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS13 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint32, p1 int64, p2 bool) int8 {
    fmt.Println(p2)
    p2 = p2
    if p2 {
        p1 += 0
        fmt.Println(p0, p2, p0)
        for fi6 := 0; fi6 < 4; fi6++ {
            v7 := GenS1{F0: p0, F1: p0}
            _ = fi6
            _ = v7
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return -46
}
func genFn1(p0 string, p1 uint) int {
    for fi8 := 0; fi8 < 5; fi8++ {
        if false {
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 bool = true
                _ = fi9
                _ = v10
            }
            var v11 uint64 = 1
            _ = v11
        }
        _ = fi8
    }
    v12 := GenS1{F0: 4294967295, F1: 65535}
    p0 = ("foo bar" + (p0 + p0))
    var v13 float32 = (((1e10 - 0.0) - 1.5) - 0.0)
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return -441
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 128
    var v3 int32 = 357
    var v4 uint32 = 65535
    v5 := make(map[string]uint32)
    v14 := GenS1{F0: 877, F1: 42}
    if false {
        v2 += v2
        var v15 int16 = 0
        v16 := GenS1{F0: 1, F1: v4}
        _ = v15
        _ = v16
    } else {
        var v17 int = 767
        _ = v17
    }
    fmt.Println(v3)
    v18 := GenS2{F0: -55}
    v19 := []int{-2147483648, 1}
    switch -90 {
    case 127:
        if false {
            if false {
                v3 += (v3 << ((934) & 63))
                var v20 int8 = -128
                _ = v20
            } else {
                var v21 uint32 = (^(v4))
                _ = v21
            }
            var v22 int32 = v3
            v23 := GenS2{F0: 122}
            _ = v22
            _ = v23
        }
    case -2147483648:
        v24 := GenS1{F0: 65535, F1: (v4 % ((2) | 1))}
        _ = v24
    default:
        v3 += v3
    }
    fmt.Println(v4, v4, v2)
    switch "" {
    case "hello":
        var v25 string = ("foo bar" + "")
        _ = v25
    case "\n\t":
        for fi26 := 0; fi26 < 17; fi26++ {
            switch (fi26 >> ((132) & 63)) {
            case 32768:
                v27 := GenS3{F0: (-(v3)), F1: (v3 << ((18446744073709551615) & 63)), F2: genFn1(("foo bar" + "foo bar"), 255)}
                _ = v27
            case -735:
                for fi28 := 0; fi28 < 17; fi28++ {
                    v3 = v3
                    _ = fi28
                }
            default:
                v29 := GenS1{F0: 479, F1: v4}
                _ = v29
            }
            v2 += (v2 % ((v2) | 1))
            v30 := GenS2{F0: genFn0((v4 << ((306) & 63)), -9223372036854775808, (2 == 42))}
            _ = fi26
            _ = v30
        }
    }
    v31 := GenS4{F0: (((("foo bar" + "a") + "hello") + "\n\t") + ""), F1: -12}
    v4 = v4
    for fi32 := 0; fi32 < 15; fi32++ {
        v3 = v3
        _ = fi32
    }
    if true {
        if false {
            switch ("a" + "a") {
            case "":
                for fi33 := 0; fi33 < 20; fi33++ {
                    v34 := GenS2{F0: -30}
                    switch "a" {
                    case "hello":
                        if true {
                            if true {
                                v35 := GenS5{F0: (v4 << ((42) & 63))}
                                _ = v35
                            } else {
                                v3 = -2147483648
                                var v36 uint64 = uint64(v2)
                                v19 = append(v19, fi33)
                                _ = v36
                            }
                        } else {
                            v2 = v2
                        }
                    case "\n\t":
                        fmt.Println(v2, fi33, v3)
                    case "a":
                        var v37 int16 = 1000
                        _ = v37
                    }
                    if (false && true) {
                        if (((0.1 + -2.25) - 0.0) == 3.14159) {
                            if (2147483647 <= v3) {
                                v38 := v1[fi33]
                                if false {
                                    if ((0.1 + 0.0) < 0.0) {
                                        switch ("a" + "a") {
                                        case "":
                                            v39 := GenS3{F0: v3, F1: v3, F2: fi33}
                                            _ = v39
                                        }
                                        if true {
                                            v3 = (v3 & -1)
                                        } else {
                                            if v38 {
                                                if v38 {
                                                    if (796 == 0) {
                                                        var v40 int64 = 42
                                                        v5["foo bar"] = (v4 >> ((253) & 63))
                                                        for fi41 := 0; fi41 < 20; fi41++ {
                                                            fmt.Println(v40)
                                                            var v42 int64 = (v40 ^ v40)
                                                            if v38 {
                                                                if v38 {
                                                                    switch "\n\t" {
                                                                    case "foo bar":
                                                                        var v43 int64 = (v40 << ((42) & 63))
                                                                        _ = v43
                                                                    }
                                                                    v44 := GenS1{F0: 0, F1: (v4 ^ (v4 << ((67) & 63)))}
                                                                    _ = v44
                                                                } else {
                                                                    v4 += (^(v4))
                                                                }
                                                                if (!(-32768 == -1)) {
                                                                    fmt.Println(v38)
                                                                    v2 = v2
                                                                    v45 := GenS5{F0: 4294967295}
                                                                    _ = v45
                                                                } else {
                                                                    v2 = 128
                                                                    if false {
                                                                        v40 = -1
                                                                    }
                                                                    var v46 int32 = 2147483647
                                                                    _ = v46
                                                                }
                                                                var v47 int16 = 698
                                                                switch "a" {
                                                                case "a":
                                                                    fmt.Println(v40, fi33, fi41)
                                                                case "":
                                                                    var v48 float64 = 1.5
                                                                    _ = v48
                                                                default:
                                                                    v4 += 1
                                                                }
                                                                _ = v47
                                                            }
                                                            _ = fi41
                                                            _ = v42
                                                        }
                                                        _ = v40
                                                    }
                                                }
                                                if v38 {
                                                    v3 = v3
                                                    var v49 uint32 = (v4 | (v4 | 42))
                                                    _ = v49
                                                } else {
                                                    fmt.Println(v4, fi33)
                                                }
                                                var v50 float32 = float32((0.1 - 1e10))
                                                _ = v50
                                            } else {
                                                _ = len(v19)
                                            }
                                            v4 -= 2
                                        }
                                        if v38 {
                                            _ = len(v1)
                                            v38 = ((42 <= 2) && true)
                                            fmt.Println(v3)
                                        } else {
                                            v4 -= v4
                                            for fi51 := 0; fi51 < 6; fi51++ {
                                                var v52 int16 = 612
                                                v1[fi51] = true
                                                var v53 uint8 = 20
                                                fmt.Println(v53)
                                                if (!(v53 > (v2 >> ((551) & 63)))) {
                                                    v38 = (345 > 4294967295)
                                                    if (true && (!false)) {
                                                        if len(v19) > 0 {
                                                            v19[0] = genFn1(("hello" + (("" + (("\n\t" + "a") + ("foo bar" + ("foo bar" + "foo bar")))) + ("a" + ""))), 4294967295)
                                                        }
                                                        v54 := GenS1{F0: 197, F1: (v4 ^ (^(v4)))}
                                                        _ = v54
                                                    } else {
                                                        fmt.Println(v2, fi51)
                                                        if v38 {
                                                            var v55 float64 = 1.5
                                                            _ = v55
                                                        } else {
                                                            if v38 {
                                                                fmt.Println(v52, v2)
                                                            } else {
                                                                switch "foo bar" {
                                                                case "foo bar":
                                                                    if true {
                                                                        if (!v38) {
                                                                            var v56 uint16 = 2
                                                                            _ = v56
                                                                        } else {
                                                                            if len(v19) >= 2 {
                                                                                v19 = v19[1:len(v19)]
                                                                            }
                                                                            var v57 uint16 = 369
                                                                            switch ("foo bar" + "a") {
                                                                            case "a":
                                                                                var v58 float32 = 3.14159
                                                                                _ = v58
                                                                            case "hello":
                                                                                var v59 string = "hello"
                                                                                _ = v59
                                                                            case "foo bar":
                                                                                if v38 {
                                                                                    var v60 string = "a"
                                                                                    _ = v60
                                                                                }
                                                                            }
                                                                            _ = len(v19)
                                                                            if v38 {
                                                                                v38 = v38
                                                                                var v61 int = fi33
                                                                                _ = v61
                                                                            } else {
                                                                                v4 += (^(v4))
                                                                                switch (((("a" + "\n\t") + "\n\t") + (("\n\t" + "") + "")) + ("hello" + ((("foo bar" + ("hello" + "hello")) + "foo bar") + "a"))) {
                                                                                case "foo bar":
                                                                                    if false {
                                                                                        var v62 string = "a"
                                                                                        v63 := GenS6{F0: 9223372036854775807}
                                                                                        _ = v62
                                                                                        _ = v63
                                                                                    } else {
                                                                                        var v64 int8 = -17
                                                                                        var v65 int64 = 1
                                                                                        switch ("\n\t" + (("hello" + "\n\t") + "\n\t")) {
                                                                                        case "":
                                                                                            var v66 int = (fi51 - 32768)
                                                                                            v2 = (^(v53))
                                                                                            v53 += 2
                                                                                            _ = len(v19)
                                                                                            v67 := GenS5{F0: (^(v4))}
                                                                                            _ = v66
                                                                                            _ = v67
                                                                                        case "hello":
                                                                                            switch "a" {
                                                                                            case "a":
                                                                                                var v68 bool = (0 <= 18446744073709551615)
                                                                                                _ = v68
                                                                                            }
                                                                                        case "\n\t":
                                                                                            v69 := GenS7{F0: int32(v65), F1: v2}
                                                                                            _ = v69
                                                                                        default:
                                                                                            var v70 uint64 = 2
                                                                                            _ = v70
                                                                                        }
                                                                                        _ = v64
                                                                                        _ = v65
                                                                                    }
                                                                                case "hello":
                                                                                    var v71 int64 = 0
                                                                                    _ = v71
                                                                                }
                                                                            }
                                                                            _ = v57
                                                                        }
                                                                        var v72 uint = 576
                                                                        v73 := GenS8{F0: -1, F1: "", F2: 127}
                                                                        _ = v72
                                                                        _ = v73
                                                                    } else {
                                                                        var v74 int32 = v3
                                                                        fmt.Println(fi33, fi33, v74)
                                                                        _ = v74
                                                                    }
                                                                case "":
                                                                    var v75 uint16 = 65535
                                                                    _ = v75
                                                                case "hello":
                                                                    var v76 int64 = 506
                                                                    _ = v76
                                                                }
                                                                var v77 uint64 = 532
                                                                _ = v77
                                                            }
                                                        }
                                                        v2 = (^(v2))
                                                        var v78 int = (fi51 - (fi51 - (fi51 & genFn1("a", 1))))
                                                        _ = v78
                                                    }
                                                }
                                                _ = fi51
                                                _ = v52
                                                _ = v53
                                            }
                                            var v79 int8 = -1
                                            _ = v79
                                        }
                                    }
                                    if (0 == 1000) {
                                        if len(v19) >= 2 {
                                            v19 = v19[1:len(v19)]
                                        }
                                    }
                                }
                                switch (("a" + ("a" + "hello")) + ("" + "")) {
                                case "a":
                                    if (((!((^(v4)) <= 0)) && ((^(v4)) != v4)) || v38) {
                                        for fi80 := 0; fi80 < 3; fi80++ {
                                            if (false && ((!v38) && ((!false) || false))) {
                                                if v38 {
                                                    fmt.Println(fi33, fi33, fi33)
                                                }
                                            }
                                            _ = fi80
                                        }
                                        var v81 int8 = 127
                                        var v82 bool = v38
                                        _ = v81
                                        _ = v82
                                    } else {
                                        v83 := GenS2{F0: genFn0(4294967295, 42, v38)}
                                        fmt.Println(v2)
                                        _ = v83
                                    }
                                default:
                                    var v84 uint16 = 255
                                    _ = v84
                                }
                                _ = v38
                            } else {
                                for fi85 := 0; fi85 < 4; fi85++ {
                                    if true {
                                        switch 127 {
                                        case -875:
                                            if false {
                                                if false {
                                                    fmt.Println(v2, v3, v4)
                                                    fmt.Println(fi33)
                                                    v19 = append(v19, -1)
                                                } else {
                                                    var v86 uint16 = 42
                                                    fmt.Println(v4)
                                                    var v87 int = fi33
                                                    _ = v86
                                                    _ = v87
                                                }
                                            } else {
                                                _ = len(v5)
                                                v2 += 127
                                                v3 = (v3 / ((v3) | 1))
                                            }
                                        case 2147483647:
                                            v88 := GenS6{F0: 9223372036854775807}
                                            _ = v88
                                        case 128:
                                            switch (("hello" + "\n\t") + ("\n\t" + "")) {
                                            case "a":
                                                v89 := GenS5{F0: (v4 / ((2) | 1))}
                                                _ = v89
                                            case "foo bar":
                                                if ((-2.25 + -2.25) >= 3.14159) {
                                                    _ = len(v19)
                                                    var v90 bool = (127 >= 128)
                                                    if v90 {
                                                        switch 9223372036854775807 {
                                                        case 42:
                                                            v91 := v1[32767]
                                                            _ = v91
                                                        default:
                                                            v4 = 65535
                                                        }
                                                        if (v90 && ((("\n\t" + "foo bar") + ("\n\t" + (((((("hello" + "foo bar") + "") + "hello") + "foo bar") + "foo bar") + ""))) < "hello")) {
                                                            v90 = true
                                                            var v92 int16 = 42
                                                            _ = v92
                                                        }
                                                        if (!v90) {
                                                            v93 := GenS3{F0: (v3 - v3), F1: v3, F2: fi85}
                                                            var v94 int32 = 1
                                                            _ = v93
                                                            _ = v94
                                                        } else {
                                                            var v95 float32 = ((((3.14159 * ((1e10 * (((0.0 + float32((-2.25 * 0.1))) - (float32(0.1) - 1.5)) / 0.1)) / 1.5)) * (3.14159 / 1e10)) - (-2.25 - (1e10 - 3.14159))) - (float32(0.1) / 0.1))
                                                            v2 += 255
                                                            if v90 {
                                                                var v96 int64 = 9223372036854775807
                                                                v97 := GenS5{F0: v4}
                                                                _ = v96
                                                                _ = v97
                                                            } else {
                                                                v19 = append(v19, (^(fi33)))
                                                            }
                                                            fmt.Println(v95, v2)
                                                            _ = v95
                                                        }
                                                        var v98 uint32 = 65535
                                                        fmt.Println(v98, fi85, v2)
                                                        _ = v98
                                                    } else {
                                                        v3 -= v3
                                                        var v99 bool = (true || ("\n\t" > "foo bar"))
                                                        if len(v19) > 0 {
                                                            v19[0] = (fi85 * genFn1("", 4294967295))
                                                        }
                                                        _ = v99
                                                    }
                                                    _ = v90
                                                }
                                            case "hello":
                                                v2 = v2
                                            }
                                        default:
                                            v19 = append(v19, fi33)
                                        }
                                        fmt.Println(v3, v3)
                                        v3 = (v3 % ((int32(fi85)) | 1))
                                    } else {
                                        var v100 int = (fi33 >> ((1) & 63))
                                        v101 := v1[-175]
                                        _ = v100
                                        _ = v101
                                    }
                                    var v102 uint16 = 665
                                    if false {
                                        var v103 int32 = int32(fi33)
                                        v104 := GenS8{F0: -1, F1: "foo bar", F2: 167}
                                        v103 += (^(v3))
                                        _ = v103
                                        _ = v104
                                    } else {
                                        v105 := GenS9{F0: 42, F1: (v2 + v2)}
                                        v2 = (v2 % (((v2 * (v2 + (v2 + (v2 / ((35) | 1)))))) | 1))
                                        switch fi33 {
                                        case 4294967295:
                                            if (true && false) {
                                                var v106 bool = true
                                                var v107 int64 = 1
                                                v108 := GenS2{F0: 26}
                                                _ = v106
                                                _ = v107
                                                _ = v108
                                            }
                                        default:
                                            var v109 float32 = -2.25
                                            _ = v109
                                        }
                                        _ = v105
                                    }
                                    _ = fi85
                                    _ = v102
                                }
                            }
                            switch -1 {
                            case 9223372036854775807:
                                switch fi33 {
                                case -366:
                                    if ((((1.5 / 3.14159) * (3.14159 * 0.1)) + 0.1) >= 0.0) {
                                        var v110 bool = true
                                        v111 := GenS10{F0: 0}
                                        fmt.Println(v3)
                                        _ = v110
                                        _ = v111
                                    }
                                default:
                                    v112 := GenS2{F0: -126}
                                    _ = v112
                                }
                            case -9223372036854775808:
                                if (266 > 4294967295) {
                                    v4 = v4
                                } else {
                                    for fi113 := 0; fi113 < 16; fi113++ {
                                        v2 -= v2
                                        var v114 uint64 = 1
                                        var v115 uint32 = v4
                                        _ = fi113
                                        _ = v114
                                        _ = v115
                                    }
                                    for fi116 := 0; fi116 < 15; fi116++ {
                                        v117 := GenS3{F0: (v3 << ((273) & 63)), F1: 42, F2: (fi116 & 109)}
                                        _ = fi116
                                        _ = v117
                                    }
                                }
                            case -964:
                                for fi118 := 0; fi118 < 7; fi118++ {
                                    var v119 float64 = float64(1e10)
                                    _ = fi118
                                    _ = v119
                                }
                            }
                        }
                    } else {
                        var v120 string = ""
                        _ = v120
                    }
                    _ = fi33
                    _ = v34
                }
            case "foo bar":
                var v121 int8 = -34
                _ = v121
            case "\n\t":
                if false {
                    _ = len(v19)
                } else {
                    fmt.Println(v4, v3)
                }
            default:
                switch 1 {
                case 42:
                    fmt.Println(v3)
                case 32767:
                    var v122 int8 = 42
                    _ = v122
                }
            }
        }
        switch 2147483647 {
        case 127:
            if (int32(v2) <= (v3 / ((v3) | 1))) {
                if false {
                    switch ("hello" + (("\n\t" + "foo bar") + "a")) {
                    case "\n\t":
                        v123 := GenS4{F0: "hello", F1: int8(v3)}
                        _ = v123
                    case "hello":
                        v4 = v4
                    case "":
                        var v124 uint = 42
                        _ = v124
                    }
                } else {
                    if (42 <= 235) {
                        v3 = (v3 / (((v3 + (v3 ^ v3))) | 1))
                        var v125 float32 = 1e10
                        _ = v125
                    }
                    if false {
                        switch 2147483647 {
                        case 2147483647:
                            fmt.Println(v2)
                        }
                    }
                    for fi126 := 0; fi126 < 10; fi126++ {
                        var v127 int64 = 42
                        var v128 float64 = float64(1e10)
                        switch 1 {
                        case 9223372036854775807:
                            var v129 uint8 = v2
                            _ = v129
                        case -751:
                            _ = len(v19)
                        case 128:
                            v3 -= (v3 % ((v3) | 1))
                        }
                        _ = fi126
                        _ = v127
                        _ = v128
                    }
                }
                if (1 > 475) {
                    fmt.Println(v4)
                } else {
                    for fi130 := 0; fi130 < 11; fi130++ {
                        fmt.Println(v3, fi130)
                        _ = fi130
                    }
                    var v131 float32 = (((float32((float64(((3.14159 - float32(1.5)) + float32(((((0.1 + (0.0 - (-2.25 + -2.25))) * float64(1.5)) * 0.1) / 1.5)))) / 0.1)) - 1e10) * float32(3.14159)) - -2.25)
                    for fi132 := 0; fi132 < 17; fi132++ {
                        v133 := GenS11{F0: 42}
                        v131 -= (float32(3.14159) + v131)
                        _ = fi132
                        _ = v133
                    }
                    _ = v131
                }
                v3 = v3
            } else {
                switch "hello" {
                case "a":
                    v4 += 2
                case "":
                    if ((0.1 / -2.25) <= 1e10) {
                        var v134 float64 = 1e10
                        v134 -= ((v134 - (v134 * float64((((0.1 / float32((v134 * (v134 - (v134 / v134))))) / (float32(v134) / (-2.25 - ((3.14159 * -2.25) / (1.5 / -2.25))))) + (1.5 / 1.0))))) / (-2.25 - 3.14159))
                        _ = v134
                    }
                case "hello":
                    var v135 uint64 = 544
                    _ = v135
                }
                fmt.Println(v4, v4)
            }
        case 0:
            v136 := GenS6{F0: 82}
            _ = v136
        }
    }
    var v137 bool = (v2 == (v2 / ((v2) | 1)))
    for fi138 := 0; fi138 < 19; fi138++ {
        if true {
            for fi139 := 0; fi139 < 11; fi139++ {
                var v140 string = "a"
                _ = fi139
                _ = v140
            }
            v2 += (v2 | (v2 * v2))
            fmt.Println(fi138, v2)
        }
        if true {
            for fi141 := 0; fi141 < 6; fi141++ {
                v142 := GenS12{F0: v4, F1: v4, F2: 127}
                var v143 int32 = (^(v3))
                _ = fi141
                _ = v142
                _ = v143
            }
            var v144 uint8 = (^(v2))
            var v145 int8 = -122
            _ = v144
            _ = v145
        } else {
            v137 = v137
            v146 := GenS5{F0: (^(v4))}
            _ = v146
        }
        var v147 uint16 = 1
        _ = fi138
        _ = v147
    }
    if v137 {
        v4 = 65535
        v148 := GenS3{F0: (v3 | (v3 ^ (-(v3)))), F1: v3, F2: -508}
        v4 = (v4 + (^(v4)))
        _ = v148
    } else {
        v137 = (1000 <= -3)
        v149 := GenS12{F0: (v4 + (v4 / ((v4) | 1))), F1: v4, F2: 42}
        var v150 float32 = 1e10
        _ = v149
        _ = v150
    }
    if v137 {
        fmt.Println(v3)
        v137 = true
        if false {
            v137 = v137
        }
    } else {
        if v137 {
            if (!false) {
                var v151 string = "hello"
                if v137 {
                    var v152 int32 = (v3 & 1)
                    fmt.Println(v3)
                    var v153 int16 = -153
                    _ = v152
                    _ = v153
                } else {
                    for fi154 := 0; fi154 < 6; fi154++ {
                        var v155 uint32 = v4
                        v156 := GenS12{F0: (^(v155)), F1: v4, F2: 127}
                        v3 -= v3
                        _ = fi154
                        _ = v155
                        _ = v156
                    }
                    var v157 int64 = -1
                    if (1e10 > -2.25) {
                        if v137 {
                            fmt.Println(v137, v2, v157)
                        }
                    } else {
                        for fi158 := 0; fi158 < 1; fi158++ {
                            for fi159 := 0; fi159 < 16; fi159++ {
                                if (!v137) {
                                    v160 := GenS13{F0: 1, F1: 1}
                                    var v161 int32 = 0
                                    v151 = "a"
                                    _ = v160
                                    _ = v161
                                }
                                switch fi159 {
                                case 1:
                                    var v162 string = (v151 + (v151 + v151))
                                    _ = v162
                                case 32767:
                                    var v163 bool = v137
                                    _ = v163
                                default:
                                    var v164 int64 = (-(v157))
                                    _ = v164
                                }
                                _ = fi159
                            }
                            fmt.Println(v4, v2, fi158)
                            _ = fi158
                        }
                    }
                    _ = v157
                }
                _ = v151
            }
            var v165 uint16 = 950
            var v166 uint8 = 24
            _ = v165
            _ = v166
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v137:", v137)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v137
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
