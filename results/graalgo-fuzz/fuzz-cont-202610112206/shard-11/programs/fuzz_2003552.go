package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int16; F2 int }
type GenS3 struct { F0 int32; F1 int64; F2 uint }
type GenS4 struct { F0 uint; F1 int64; F2 uint64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int32; F1 uint16; F2 int16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8; F1 uint8 }
type GenS9 struct { F0 int16; F1 uint8; F2 uint }
type GenS10 struct { F0 int; F1 bool; F2 string }
type GenS11 struct { F0 int; F1 uint32 }
type GenS12 struct { F0 int64; F1 int8 }
type GenS13 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS14 struct { F0 int }
func genFn0(p0 string, p1 uint64, p2 int32) int {
    v7 := []int{}
    _ = len(v7)
    v8 := GenS1{F0: 4294967295, F1: -43, F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 32768
}
func genFn1() int8 {
    for fi9 := 0; fi9 < 13; fi9++ {
        if true {
            fmt.Println(fi9)
            var v10 uint = 18446744073709551615
            v11 := GenS2{F0: 65535, F1: 1000, F2: (fi9 * (fi9 - fi9))}
            _ = v10
            _ = v11
        }
        _ = fi9
    }
    var v12 uint8 = 127
    _ = v12
    return 40
}

func main() {
    var v1 string = ("foo bar" + "foo bar")
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]uint32)
    v4 := []int{551, 755, 128}
    var v5 float64 = float64((((((((0.1 - 0.1) + (float32(1e10) + float32(((float64((float32((0.0 - 3.14159)) + 0.0)) + ((((3.14159 * -2.25) / float64(0.1)) - (0.0 / (0.1 - float64((0.0 / 0.1))))) * ((1.5 + (0.1 - ((3.14159 * 1e10) - ((0.0 - 1e10) * (0.0 + (1.5 * (float64(0.0) * ((0.0 - 1e10) - 1.5)))))))) / -2.25))) - (((0.1 * 3.14159) * 1e10) * -2.25))))) * 3.14159) + 1e10) / 1e10) * 0.1) - (-2.25 - 0.1)))
    var v6 uint16 = 255
    var v13 uint32 = 65535
    var v14 uint64 = 42
    v13 = 2
    if false {
        var v15 uint16 = v6
        var v16 bool = (true && (213 == 0))
        _ = v15
        _ = v16
    }
    v1 = ("\n\t" + v1)
    delete(v2, -565)
    v4 = append(v4, 0)
    v14 += uint64(v13)
    switch 32767 {
    case -2147483648:
        v17 := GenS1{F0: -9223372036854775808, F1: -9223372036854775808, F2: 65535}
        _ = v17
    case 32768:
        v14 = (v14 % (((v14 >> ((42) & 63))) | 1))
    default:
        v6 = 65535
    }
    v5 = (float64(1e10) - v5)
    v18 := GenS1{F0: 9223372036854775807, F1: 0, F2: v6}
    switch (v1 + (((v1 + v1) + v1) + "foo bar")) {
    case "foo bar":
        switch int(v6) {
        case -2147483648:
            fmt.Println(v14, v14)
        case -244:
            v13 = v13
        case 2147483647:
            for fi19 := 0; fi19 < 10; fi19++ {
                var v20 int8 = 1
                switch v1 {
                case "hello":
                    if (uint16(v14) == 791) {
                        switch fi19 {
                        case 42:
                            if false {
                                v5 += v5
                            }
                        case 872:
                            fmt.Println(v14, v1)
                        case -1:
                            v3[517] = v13
                        }
                    }
                case "":
                    v21 := GenS1{F0: 42, F1: -589, F2: (v6 & v6)}
                    _ = v21
                default:
                    v6 = v6
                }
                _ = fi19
                _ = v20
            }
        default:
            if false {
                v5 += 1.5
                v22 := GenS3{F0: 2147483647, F1: 505, F2: 0}
                if false {
                    if ((100 < 122) || (!(!true))) {
                        for fi23 := 0; fi23 < 13; fi23++ {
                            v24 := GenS4{F0: 42, F1: 0, F2: (^(v14))}
                            fmt.Println(v1, v6, v14)
                            _ = fi23
                            _ = v24
                        }
                        v25 := GenS5{F0: 127}
                        v6 += uint16(v14)
                        _ = v25
                    }
                    if false {
                        if false {
                            fmt.Println(v6, v14, v13)
                            for fi26 := 0; fi26 < 9; fi26++ {
                                v5 = 1e10
                                _ = len(v3)
                                fmt.Println(fi26, v6, v1)
                                _ = fi26
                            }
                        }
                    } else {
                        var v27 bool = (-2147483648 == 0)
                        _ = v27
                    }
                }
                _ = v22
            } else {
                if ((!true) || (0 <= -113)) {
                    if true {
                        var v28 float64 = 0.0
                        _ = v28
                    } else {
                        v29 := GenS3{F0: -1, F1: 0, F2: 1}
                        var v30 int16 = 42
                        for fi31 := 0; fi31 < 9; fi31++ {
                            for fi32 := 0; fi32 < 7; fi32++ {
                                var v33 int16 = v30
                                if true {
                                    fmt.Println(fi32, v1)
                                    var v34 uint64 = v14
                                    switch (fi31 + fi32) {
                                    case 9223372036854775807:
                                        v35 := GenS6{F0: -2147483648, F1: v6, F2: 800}
                                        _ = v35
                                    }
                                    _ = v34
                                } else {
                                    switch (("hello" + v1) + "a") {
                                    case "hello":
                                        switch genFn0(v1, v14, 813) {
                                        case 9223372036854775807:
                                            fmt.Println(v6, v13, v30)
                                        case 2:
                                            v14 += (v14 >> ((255) & 63))
                                        case 127:
                                            if (100 == -71) {
                                                fmt.Println(v14, v30)
                                                var v36 uint64 = v14
                                                v37 := GenS7{F0: fi31}
                                                _ = v36
                                                _ = v37
                                            } else {
                                                v38 := GenS7{F0: fi32}
                                                _ = v38
                                            }
                                        }
                                    case "\n\t":
                                        var v39 int32 = -1
                                        _ = v39
                                    case "a":
                                        v30 += 1
                                        v14 += (v14 * v14)
                                        switch -9223372036854775808 {
                                        case 989:
                                            if false {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                v30 -= v30
                                                v2[genFn0((v1 + v1), (v14 >> ((42) & 63)), -2147483648)] = v13
                                            } else {
                                                v33 -= int16(v14)
                                                var v40 bool = true
                                                _ = v40
                                            }
                                        case -476:
                                            delete(v3, (^(v13)))
                                        case 993:
                                            if ("foo bar" == "hello") {
                                                v1 = ""
                                                fmt.Println(v1)
                                                v14 = uint64(v13)
                                            } else {
                                                switch (^(fi32)) {
                                                case 127:
                                                    var v41 uint = 42
                                                    _ = v41
                                                case 32768:
                                                    if (((false || (18446744073709551615 < 18446744073709551615)) || (9 < -128)) && true) {
                                                        v42 := GenS2{F0: (v6 + 42), F1: (v30 & v30), F2: fi32}
                                                        _ = v42
                                                    } else {
                                                        v43 := GenS6{F0: 2147483647, F1: 32767, F2: 42}
                                                        var v44 int64 = 455
                                                        if ((!((^(v6)) != v6)) || true) {
                                                            v45 := GenS6{F0: -1, F1: (v6 | (v6 / ((v6) | 1))), F2: v30}
                                                            var v46 uint16 = (^(v6))
                                                            switch ("hello" + v1) {
                                                            case "\n\t":
                                                                if len(v4) >= 2 {
                                                                    v4 = v4[1:len(v4)]
                                                                }
                                                            case "":
                                                                v46 = 0
                                                            }
                                                            if (fi32 != (fi31 ^ (fi32 / (((^(fi32))) | 1)))) {
                                                                v47 := GenS2{F0: (v46 + v6), F1: 1, F2: (fi31 | (-(fi31)))}
                                                                switch (fi32 ^ fi31) {
                                                                case 128:
                                                                    v44 = (v44 << ((65535) & 63))
                                                                }
                                                                _ = v47
                                                            } else {
                                                                v13 += v13
                                                                v48 := GenS7{F0: int(v44)}
                                                                var v49 uint = 18446744073709551615
                                                                var v50 uint64 = v14
                                                                v4 = append(v4, (fi31 >> ((v49) & 63)))
                                                                var v51 bool = false
                                                                v49 -= v49
                                                                _ = v48
                                                                _ = v49
                                                                _ = v50
                                                                _ = v51
                                                            }
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                        _ = len(v3)
                                                        v5 = 3.14159
                                                        _ = v43
                                                        _ = v44
                                                    }
                                                    v13 = (v13 - (v13 % ((407) | 1)))
                                                }
                                                var v52 int = (fi31 - (fi32 % ((fi32) | 1)))
                                                _ = v52
                                            }
                                        }
                                    default:
                                        var v53 int = int(v13)
                                        _ = v53
                                    }
                                    v54 := GenS8{F0: 43, F1: 91}
                                    _ = v54
                                }
                                v33 += v30
                                _ = fi32
                                _ = v33
                            }
                            v14 -= (v14 ^ 660)
                            _ = fi31
                        }
                        _ = v29
                        _ = v30
                    }
                    var v55 uint8 = 127
                    _ = v55
                } else {
                    if (-1 < 42) {
                        v56 := GenS9{F0: 32767, F1: 127, F2: 405}
                        _ = v56
                    } else {
                        var v57 int8 = -128
                        for fi58 := 0; fi58 < 6; fi58++ {
                            if true {
                                fmt.Println(v57, v6, v6)
                            }
                            v59 := GenS7{F0: fi58}
                            _ = fi58
                            _ = v59
                        }
                        fmt.Println(v57, v57, v57)
                        _ = v57
                    }
                }
                switch v1 {
                case "foo bar":
                    fmt.Println(v5, v13, v13)
                case "\n\t":
                    var v60 float64 = v5
                    _ = v60
                default:
                    for fi61 := 0; fi61 < 12; fi61++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        for fi62 := 0; fi62 < 15; fi62++ {
                            var v63 uint8 = 1
                            if ((1.5 + (3.14159 / 1.0)) >= (1e10 - 0.1)) {
                                v64 := GenS10{F0: genFn0(v1, v14, 0), F1: (false || false), F2: "a"}
                                v65 := GenS1{F0: genFn0(("\n\t" + (("a" + ("hello" + v1)) + (("foo bar" + v1) + v1))), (v14 | 18446744073709551615), 2147483647), F1: 9223372036854775807, F2: 412}
                                v66 := GenS11{F0: -744, F1: v13}
                                _ = v64
                                _ = v65
                                _ = v66
                            } else {
                                v5 -= (((3.14159 + v5) * 1.5) + 1e10)
                                switch "\n\t" {
                                case "hello":
                                    v67 := GenS12{F0: 9223372036854775807, F1: 85}
                                    _ = v67
                                }
                                var v68 string = v1
                                var v69 int64 = 1
                                _ = v68
                                _ = v69
                            }
                            _ = fi62
                            _ = v63
                        }
                        _ = fi61
                    }
                }
                var v70 uint32 = v13
                _ = v70
            }
        }
    case "":
        for fi71 := 0; fi71 < 15; fi71++ {
            delete(v3, 584)
            fmt.Println(fi71)
            if true {
                for fi72 := 0; fi72 < 8; fi72++ {
                    v14 = v14
                    var v73 int = (fi72 / ((fi72) | 1))
                    var v74 int16 = -32768
                    fmt.Println(v6, v74, v5)
                    var v75 string = ("\n\t" + ((v1 + v1) + v1))
                    _ = fi72
                    _ = v73
                    _ = v74
                    _ = v75
                }
                v5 = v5
            } else {
                for fi76 := 0; fi76 < 14; fi76++ {
                    if true {
                        v77 := v2[genFn0((("" + ("a" + "hello")) + ("foo bar" + (v1 + "foo bar"))), (v14 ^ v14), 2147483647)]
                        if false {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            if false {
                                v13 -= 675
                                v78 := v2[42]
                                _ = v78
                            }
                        } else {
                            v79 := GenS2{F0: (v6 * v6), F1: 403, F2: fi71}
                            _ = v79
                        }
                        var v80 int64 = 42
                        _ = v77
                        _ = v80
                    } else {
                        if (v14 > 4294967295) {
                            switch ("a" + ("a" + "hello")) {
                            case "":
                                var v81 uint32 = (v13 - 2147483648)
                                _ = v81
                            case "a":
                                var v82 uint = 969
                                _ = v82
                            case "\n\t":
                                var v83 float64 = 3.14159
                                _ = v83
                            default:
                                v1 = ("\n\t" + ("a" + (v1 + (("a" + "foo bar") + ((((v1 + v1) + (v1 + (v1 + "a"))) + v1) + v1)))))
                            }
                            v6 = (^(v6))
                            v13 = v13
                            if (1.5 != float32(0.0)) {
                                v13 -= (v13 >> ((uint(v14)) & 63))
                                var v84 int = fi71
                                var v85 uint16 = (v6 * (v6 + v6))
                                v86 := GenS6{F0: int32(fi71), F1: (v85 % ((v6) | 1)), F2: 32767}
                                _ = v84
                                _ = v85
                                _ = v86
                            } else {
                                if (false || false) {
                                    v87 := GenS12{F0: -13, F1: 127}
                                    var v88 uint8 = 1
                                    _ = v87
                                    _ = v88
                                }
                            }
                        }
                        if (!false) {
                            v2[(fi71 ^ (fi76 % ((fi71) | 1)))] = v13
                            switch (v1 + "") {
                            case "a":
                                switch fi71 {
                                case 32767:
                                    v14 = v14
                                }
                            default:
                                switch -908 {
                                case 1:
                                    v89 := GenS13{F0: 87, F1: 596, F2: 1}
                                    _ = v89
                                case 50:
                                    var v90 float32 = -2.25
                                    _ = v90
                                case 32767:
                                    var v91 uint64 = 84
                                    _ = v91
                                }
                            }
                            v14 += 18446744073709551615
                        } else {
                            switch genFn0("hello", 0, 2147483647) {
                            case 2147483647:
                                _ = len(v4)
                            case 9223372036854775807:
                                v92 := GenS3{F0: 42, F1: -224, F2: 65535}
                                _ = v92
                            case 237:
                                v93 := GenS11{F0: (fi76 * (fi76 ^ 4294967295)), F1: v13}
                                _ = v93
                            default:
                                v2[genFn0("\n\t", (v14 - v14), 390)] = uint32(v6)
                            }
                            fmt.Println(v1, v1)
                            var v94 uint64 = 0
                            v94 = (v94 - 4294967295)
                            v95 := GenS5{F0: -43}
                            fmt.Println(v6, fi71, v5)
                            _ = v94
                            _ = v95
                        }
                    }
                    delete(v3, (^(v13)))
                    var v96 uint32 = (^(v13))
                    _ = fi76
                    _ = v96
                }
                var v97 uint64 = v14
                if true {
                    for fi98 := 0; fi98 < 16; fi98++ {
                        if false {
                            var v99 uint32 = v13
                            fmt.Println(v99, v13, v5)
                            v13 -= v99
                            _ = len(v4)
                            v100 := GenS2{F0: (^(v6)), F1: -1, F2: fi98}
                            if (!(!false)) {
                                fmt.Println(v14, v13)
                                switch (fi98 + (fi98 % ((fi71) | 1))) {
                                case 32767:
                                    _ = len(v4)
                                default:
                                    v6 += v6
                                }
                                if (true || (true || (!false))) {
                                    var v101 uint16 = 224
                                    switch fi71 {
                                    case -692:
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    _ = v101
                                }
                            } else {
                                var v102 float64 = v5
                                if (true && (v99 == (v13 + v13))) {
                                    v1 = ""
                                    v103 := GenS6{F0: 296, F1: v6, F2: 626}
                                    delete(v2, (fi98 / (((fi98 % ((fi98) | 1))) | 1)))
                                    _ = v103
                                }
                                _ = v102
                            }
                            _ = v100
                            _ = v99
                        }
                        fmt.Println(v13, v5, v14)
                        _ = fi98
                    }
                    for fi104 := 0; fi104 < 15; fi104++ {
                        if (((0.1 * 0.0) - 1.5) == 0.0) {
                            var v105 int8 = 42
                            fmt.Println(v1)
                            var v106 float32 = ((-2.25 + (3.14159 / (3.14159 / 1.0))) - 3.14159)
                            _ = v105
                            _ = v106
                        }
                        v5 = v5
                        v1 = v1
                        _ = fi104
                    }
                } else {
                    var v107 uint = 255
                    for fi108 := 0; fi108 < 5; fi108++ {
                        fmt.Println(v1)
                        v13 = v13
                        var v109 uint64 = (v14 | v97)
                        fmt.Println(v109, v6, fi71)
                        _ = fi108
                        _ = v109
                    }
                    _ = v107
                }
                _ = v97
            }
            _ = fi71
        }
    case "\n\t":
        for fi110 := 0; fi110 < 19; fi110++ {
            for fi111 := 0; fi111 < 8; fi111++ {
                var v112 float32 = float32((v5 / (v5 / 1.5)))
                v112 = float32(v5)
                fmt.Println(v14, fi111, fi111)
                _ = fi111
                _ = v112
            }
            _ = fi110
        }
    default:
        v2[-408] = v13
    }
    var v113 string = (v1 + (v1 + "hello"))
    var v114 uint8 = 234
    if (v14 != 42) {
        switch int(v114) {
        case 1:
            for fi115 := 0; fi115 < 15; fi115++ {
                v116 := GenS9{F0: -253, F1: (^(v114)), F2: 4294967295}
                switch (v113 + "a") {
                case "\n\t":
                    v6 = v6
                default:
                    v113 = v113
                }
                var v117 int8 = -72
                _ = fi115
                _ = v116
                _ = v117
            }
        case -1:
            var v118 uint8 = 131
            _ = v118
        }
    } else {
        _ = len(v4)
    }
    for fi119 := 0; fi119 < 3; fi119++ {
        for fi120 := 0; fi120 < 2; fi120++ {
            v121 := GenS13{F0: (^(v114)), F1: v13, F2: 23}
            if false {
                v114 += (v114 * 127)
            }
            switch int(v6) {
            case -395:
                var v122 float64 = 0.0
                _ = v122
            }
            _ = fi120
            _ = v121
        }
        _ = len(v2)
        _ = fi119
    }
    v113 = ((v1 + v1) + v113)
    for fi123 := 0; fi123 < 8; fi123++ {
        for fi124 := 0; fi124 < 10; fi124++ {
            var v125 uint32 = v13
            switch ("foo bar" + (v1 + v1)) {
            case "\n\t":
                switch "foo bar" {
                case "":
                    if true {
                        v126 := GenS12{F0: 0, F1: 113}
                        v127 := GenS8{F0: (v114 | v114), F1: 0}
                        v128 := GenS14{F0: int(v14)}
                        _ = v126
                        _ = v127
                        _ = v128
                    }
                case "foo bar":
                    fmt.Println(fi124)
                case "\n\t":
                    if (-463 > 2147483647) {
                        var v129 uint32 = v125
                        var v130 bool = false
                        _ = v129
                        _ = v130
                    }
                default:
                    fmt.Println(v113, v114)
                }
            case "a":
                v1 = "\n\t"
            }
            _ = fi124
            _ = v125
        }
        v13 += 2
        _ = fi123
    }
    fmt.Println("v1:", v1)
    fmt.Println("v113:", v113)
    fmt.Println("v114:", v114)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v113
    _ = v114
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
