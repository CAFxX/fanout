package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 string; F1 int32; F2 int32 }
type GenS5 struct { F0 bool; F1 int16; F2 int32 }
type GenS6 struct { F0 int8; F1 int8 }
type GenS7 struct { F0 int64; F1 string }
type GenS8 struct { F0 bool; F1 uint }
type GenS9 struct { F0 int64; F1 int64 }
type GenS10 struct { F0 string; F1 int; F2 uint8 }
type GenS11 struct { F0 uint64; F1 uint16 }
type GenS12 struct { F0 uint16; F1 uint32 }
type GenS13 struct { F0 int8; F1 int64 }
type GenS14 struct { F0 int32 }
type GenS15 struct { F0 int16 }
type GenS16 struct { F0 uint16; F1 int; F2 int8 }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 int8 }
type GenS19 struct { F0 int32 }
type GenS20 struct { F0 uint; F1 uint8; F2 int64 }
type GenS21 struct { F0 uint16 }
func genFn0(p0 int32, p1 int32, p2 int16) int32 {
    v5 := GenS1{F0: 127}
    var v6 int16 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -2147483648
}

func main() {
    var v1 float32 = 0.1
    var v2 int32 = -1
    v3 := []string{"foo bar", "foo bar", ("" + "hello"), "foo bar"}
    v4 := make(map[string]uint32)
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 int64 = 134
        v9 := GenS1{F0: 51}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        v11 := GenS2{F0: 42, F1: 127}
        for fi12 := 0; fi12 < 14; fi12++ {
            v13 := GenS1{F0: 42}
            if len(v3) > 0 {
                v3[0] = "hello"
            }
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    var v14 int16 = 42
    v2 = v2
    var v15 float64 = float64(1e10)
    switch -442 {
    case 128:
        var v16 int64 = 42
        _ = v16
    }
    switch "" {
    case "":
        if false {
            v17 := GenS1{F0: 0}
            if true {
                for fi18 := 0; fi18 < 9; fi18++ {
                    v15 = 3.14159
                    if ("a" != "a") {
                        for fi19 := 0; fi19 < 17; fi19++ {
                            fmt.Println(v1)
                            switch "hello" {
                            case "foo bar":
                                fmt.Println(fi18, v15)
                            case "\n\t":
                                var v20 uint8 = 2
                                _ = v20
                            default:
                                v15 -= v15
                            }
                            var v21 uint64 = 2
                            _ = fi19
                            _ = v21
                        }
                    }
                    _ = fi18
                }
                var v22 int16 = v14
                _ = v22
            } else {
                switch 4294967295 {
                case -9223372036854775808:
                    _ = len(v3)
                default:
                    v23 := GenS1{F0: 255}
                    _ = v23
                }
                switch 4294967295 {
                case 4294967295:
                    var v24 int32 = genFn0((v2 / (((v2 + int32(v14))) | 1)), v2, 0)
                    _ = v24
                case -2147483648:
                    if true {
                        if (false && true) {
                            if len(v3) > 0 {
                                v3[0] = "foo bar"
                            }
                            if false {
                                for fi25 := 0; fi25 < 5; fi25++ {
                                    if true {
                                        fmt.Println(v14, v2, v15)
                                        v4["foo bar"] = 750
                                    }
                                    _ = fi25
                                }
                                switch "hello" {
                                case "a":
                                    for fi26 := 0; fi26 < 16; fi26++ {
                                        for fi27 := 0; fi27 < 16; fi27++ {
                                            var v28 float64 = v15
                                            var v29 uint32 = 524
                                            _ = fi27
                                            _ = v28
                                            _ = v29
                                        }
                                        v1 = float32((v15 * (-2.25 / -2.25)))
                                        _ = fi26
                                    }
                                case "foo bar":
                                    if false {
                                        var v30 string = (("a" + "hello") + "hello")
                                        _ = v30
                                    } else {
                                        if true {
                                            switch -384 {
                                            case 4294967295:
                                                fmt.Println(v1, v2, v14)
                                            default:
                                                v14 -= (v14 << ((42) & 63))
                                            }
                                            switch -713 {
                                            case 0:
                                                v31 := GenS3{F0: 1, F1: 0}
                                                _ = v31
                                            case -9223372036854775808:
                                                var v32 string = "\n\t"
                                                _ = v32
                                            }
                                        }
                                        v1 -= v1
                                    }
                                default:
                                    var v33 float32 = ((v1 / (1.5 + (v1 + v1))) * ((v1 * v1) / (3.14159 * v1)))
                                    _ = v33
                                }
                            } else {
                                v15 -= (1.5 + 0.0)
                                for fi34 := 0; fi34 < 6; fi34++ {
                                    v15 = float64(v1)
                                    fmt.Println(v2, v1)
                                    v35 := GenS3{F0: 128, F1: (-(fi34))}
                                    _ = fi34
                                    _ = v35
                                }
                                for fi36 := 0; fi36 < 19; fi36++ {
                                    var v37 uint16 = 65535
                                    var v38 int8 = -100
                                    _ = fi36
                                    _ = v37
                                    _ = v38
                                }
                            }
                            for fi39 := 0; fi39 < 16; fi39++ {
                                if (1 >= 255) {
                                    if true {
                                        fmt.Println(v15, v2)
                                    }
                                    var v40 uint32 = 42
                                    _ = v40
                                }
                                switch "a" {
                                case "a":
                                    switch (-(fi39)) {
                                    case 9223372036854775807:
                                        switch (fi39 & (fi39 / ((fi39) | 1))) {
                                        case -253:
                                            var v41 uint32 = 65535
                                            _ = v41
                                        default:
                                            var v42 uint = 0
                                            _ = v42
                                        }
                                    case 2147483647:
                                        var v43 float32 = 3.14159
                                        _ = v43
                                    default:
                                        v2 += v2
                                    }
                                case "foo bar":
                                    v2 = 1
                                default:
                                    for fi44 := 0; fi44 < 15; fi44++ {
                                        switch fi44 {
                                        case -877:
                                            var v45 uint32 = 2147483648
                                            _ = v45
                                        case 1:
                                            if len(v3) > 0 {
                                                v3[0] = ("hello" + (("a" + "a") + ("hello" + "hello")))
                                            }
                                        }
                                        var v46 int32 = v2
                                        _ = fi44
                                        _ = v46
                                    }
                                }
                                for fi47 := 0; fi47 < 2; fi47++ {
                                    fmt.Println(v14)
                                    v48 := GenS3{F0: 1, F1: fi39}
                                    _ = fi47
                                    _ = v48
                                }
                                _ = fi39
                            }
                        } else {
                            fmt.Println(v2)
                            fmt.Println(v15, v1)
                        }
                        fmt.Println(v14, v1)
                    } else {
                        fmt.Println(v15, v2)
                        fmt.Println(v2, v1, v1)
                        v49 := GenS2{F0: 42, F1: -1}
                        _ = v49
                    }
                case 507:
                    if true {
                        switch "hello" {
                        case "hello":
                            if (((true || true) || false) || (-1 == -9223372036854775808)) {
                                fmt.Println(v15)
                                if (!false) {
                                    var v50 float32 = (v1 * float32((v15 / v15)))
                                    for fi51 := 0; fi51 < 13; fi51++ {
                                        for fi52 := 0; fi52 < 13; fi52++ {
                                            var v53 int = (fi52 * fi51)
                                            v2 = v2
                                            _ = fi52
                                            _ = v53
                                        }
                                        var v54 bool = (true && ((fi51 | fi51) < (fi51 % ((9223372036854775807) | 1))))
                                        v55 := GenS4{F0: ("\n\t" + ("\n\t" + "hello")), F1: v2, F2: v2}
                                        _ = fi51
                                        _ = v54
                                        _ = v55
                                    }
                                    _ = v50
                                } else {
                                    var v56 float64 = v15
                                    v57 := GenS4{F0: "a", F1: -965, F2: v2}
                                    if true {
                                        var v58 float64 = v56
                                        switch 997 {
                                        case -9223372036854775808:
                                            var v59 bool = true
                                            _ = v59
                                        case 2147483647:
                                            if false {
                                                for fi60 := 0; fi60 < 12; fi60++ {
                                                    if (372 <= (-(v2))) {
                                                        v58 = v56
                                                        var v61 uint = 2
                                                        _ = v61
                                                    }
                                                    v15 = v56
                                                    _ = fi60
                                                }
                                                fmt.Println(v58, v58)
                                                v62 := GenS1{F0: 128}
                                                _ = v62
                                            }
                                        case 4294967295:
                                            if true {
                                                delete(v4, ("foo bar" + "a"))
                                            } else {
                                                for fi63 := 0; fi63 < 14; fi63++ {
                                                    v1 = (0.0 + v1)
                                                    switch ("foo bar" + "hello") {
                                                    case "hello":
                                                        v15 = -2.25
                                                    }
                                                    var v64 uint8 = 0
                                                    _ = fi63
                                                    _ = v64
                                                }
                                                if (true || false) {
                                                    if len(v3) > 0 {
                                                        v3[0] = "\n\t"
                                                    }
                                                }
                                                var v65 int = 182
                                                _ = v65
                                            }
                                        }
                                        _ = v58
                                    } else {
                                        switch -5 {
                                        case 32768:
                                            var v66 float64 = v56
                                            _ = v66
                                        }
                                        for fi67 := 0; fi67 < 16; fi67++ {
                                            var v68 uint16 = 1
                                            switch "a" {
                                            case "foo bar":
                                                if true {
                                                    var v69 string = ((("" + ("foo bar" + "hello")) + "foo bar") + ("" + "a"))
                                                    var v70 int64 = -1
                                                    v3 = append(v3, (v69 + ("hello" + "a")))
                                                    _ = v69
                                                    _ = v70
                                                } else {
                                                    v68 += v68
                                                }
                                            default:
                                                v56 += 0.1
                                            }
                                            _ = fi67
                                            _ = v68
                                        }
                                    }
                                    _ = v56
                                    _ = v57
                                }
                            } else {
                                if (2 >= 42) {
                                    for fi71 := 0; fi71 < 10; fi71++ {
                                        switch 1 {
                                        case 42:
                                            if (!(v15 != float64(v1))) {
                                                if true {
                                                    delete(v4, "foo bar")
                                                    fmt.Println(v2)
                                                    if true {
                                                        if (((v14 <= (v14 | (v14 & v14))) && true) || (v2 <= 641)) {
                                                            for fi72 := 0; fi72 < 13; fi72++ {
                                                                v73 := GenS5{F0: true, F1: 1000, F2: 970}
                                                                fmt.Println(fi72, v1)
                                                                _ = fi72
                                                                _ = v73
                                                            }
                                                        } else {
                                                            var v74 string = "foo bar"
                                                            v1 = (v1 / 1.5)
                                                            if ((v2 ^ 42) <= (^(v2))) {
                                                                v75 := GenS6{F0: int8(v14), F1: 127}
                                                                v76 := GenS7{F0: -1, F1: v74}
                                                                v77 := GenS3{F0: 50, F1: fi71}
                                                                _ = v75
                                                                _ = v76
                                                                _ = v77
                                                            } else {
                                                                for fi78 := 0; fi78 < 4; fi78++ {
                                                                    var v79 uint64 = 19
                                                                    var v80 uint16 = 32767
                                                                    _ = fi78
                                                                    _ = v79
                                                                    _ = v80
                                                                }
                                                                delete(v4, v74)
                                                                var v81 int = 1
                                                                _ = v81
                                                            }
                                                            _ = v74
                                                        }
                                                    }
                                                }
                                                if false {
                                                    v1 += (v1 * (3.14159 * v1))
                                                    v14 += 1
                                                }
                                                if true {
                                                    v15 = 0.1
                                                    v82 := GenS1{F0: 42}
                                                    if ((v1 == v1) && false) {
                                                        v83 := v4[("foo bar" + "")]
                                                        for fi84 := 0; fi84 < 11; fi84++ {
                                                            v85 := GenS7{F0: int64(v14), F1: "foo bar"}
                                                            fmt.Println(v2, fi71)
                                                            _ = fi84
                                                            _ = v85
                                                        }
                                                        if len(v3) > 0 {
                                                            v3[0] = ("" + "a")
                                                        }
                                                        _ = v83
                                                    } else {
                                                        var v86 int32 = v2
                                                        switch (fi71 ^ (fi71 >> ((356) & 63))) {
                                                        case -1:
                                                            var v87 int = fi71
                                                            _ = v87
                                                        case -409:
                                                            v15 = v15
                                                        case -2147483648:
                                                            v2 += v86
                                                        }
                                                        var v88 int32 = v86
                                                        _ = v86
                                                        _ = v88
                                                    }
                                                    _ = v82
                                                } else {
                                                    if (false || (float64(v1) != (1.5 / (-2.25 * v15)))) {
                                                        switch fi71 {
                                                        case 0:
                                                            var v89 uint64 = 174
                                                            _ = v89
                                                        default:
                                                            v15 += 0.1
                                                        }
                                                        var v90 uint8 = 79
                                                        var v91 int16 = (v14 / (((v14 % ((v14) | 1))) | 1))
                                                        _ = v90
                                                        _ = v91
                                                    } else {
                                                        if false {
                                                            if ((!false) || (true && (true || false))) {
                                                                var v92 int16 = (v14 - v14)
                                                                if (!(("foo bar" >= (("hello" + "\n\t") + "foo bar")) || (float32(float64(v1)) > ((v1 + 3.14159) + v1)))) {
                                                                    for fi93 := 0; fi93 < 2; fi93++ {
                                                                        v1 = float32(v15)
                                                                        var v94 string = "foo bar"
                                                                        _ = fi93
                                                                        _ = v94
                                                                    }
                                                                    var v95 int = (fi71 % (((fi71 >> ((4294967295) & 63))) | 1))
                                                                    for fi96 := 0; fi96 < 20; fi96++ {
                                                                        var v97 int8 = 24
                                                                        if false {
                                                                            var v98 bool = ((((v2 << ((42) & 63)) < v2) && (!false)) && false)
                                                                            _ = v98
                                                                        }
                                                                        _ = fi96
                                                                        _ = v97
                                                                    }
                                                                    _ = v95
                                                                }
                                                                switch "foo bar" {
                                                                case "\n\t":
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                default:
                                                                    var v99 int64 = -798
                                                                    _ = v99
                                                                }
                                                                _ = v92
                                                            }
                                                        }
                                                        _ = len(v3)
                                                        if true {
                                                            for fi100 := 0; fi100 < 16; fi100++ {
                                                                v2 = v2
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                v1 -= v1
                                                                _ = fi100
                                                            }
                                                            var v101 string = ""
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            _ = v101
                                                        } else {
                                                            switch (fi71 | fi71) {
                                                            case 42:
                                                                switch (fi71 >> ((493) & 63)) {
                                                                case 351:
                                                                    var v102 int16 = v14
                                                                    _ = v102
                                                                }
                                                            default:
                                                                v2 = v2
                                                            }
                                                            for fi103 := 0; fi103 < 2; fi103++ {
                                                                fmt.Println(v14)
                                                                _ = fi103
                                                            }
                                                            var v104 int8 = -128
                                                            _ = v104
                                                        }
                                                    }
                                                    var v105 uint64 = 2
                                                    v2 = v2
                                                    _ = v105
                                                }
                                            } else {
                                                v106 := GenS7{F0: 9223372036854775807, F1: ("" + ("foo bar" + "a"))}
                                                _ = v106
                                            }
                                        case 4294967295:
                                            v4["a"] = 65535
                                        case 127:
                                            var v107 bool = false
                                            _ = v107
                                        }
                                        v4["hello"] = 316
                                        _ = fi71
                                    }
                                } else {
                                    for fi108 := 0; fi108 < 13; fi108++ {
                                        fmt.Println(v14)
                                        for fi109 := 0; fi109 < 5; fi109++ {
                                            if (778 < 0) {
                                                var v110 string = ((("\n\t" + "foo bar") + "\n\t") + ("a" + "\n\t"))
                                                if false {
                                                    v111 := GenS2{F0: 2, F1: -128}
                                                    switch fi108 {
                                                    case -1:
                                                        switch "" {
                                                        case "\n\t":
                                                            v112 := GenS4{F0: v110, F1: v2, F2: v2}
                                                            _ = v112
                                                        case "foo bar":
                                                            v113 := GenS7{F0: 290, F1: v110}
                                                            _ = v113
                                                        case "hello":
                                                            if true {
                                                                var v114 float64 = v15
                                                                _ = v114
                                                            }
                                                        }
                                                    case -649:
                                                        switch fi109 {
                                                        case 42:
                                                            v115 := GenS4{F0: v110, F1: (v2 + (v2 + (v2 + (v2 | v2)))), F2: v2}
                                                            _ = v115
                                                        case 32768:
                                                            var v116 int64 = -1
                                                            _ = v116
                                                        default:
                                                            fmt.Println(fi108)
                                                        }
                                                    case 127:
                                                        var v117 int8 = 42
                                                        _ = v117
                                                    }
                                                    if true {
                                                        fmt.Println(fi108, fi109)
                                                        var v118 float32 = v1
                                                        _ = v118
                                                    }
                                                    if false {
                                                        var v119 uint8 = 42
                                                        _ = v119
                                                    }
                                                    _ = v111
                                                }
                                                v15 += (((0.1 / -2.25) + (0.0 + ((v15 - (v15 * float64(v1))) * (((float64(v1) - (-2.25 * -2.25)) + v15) / v15)))) / 1e10)
                                                _ = v110
                                            }
                                            v120 := v4[("hello" + "a")]
                                            v120 = 4294967295
                                            if ((false && (-9223372036854775808 < 0)) || true) {
                                                var v121 int8 = 0
                                                _ = v121
                                            } else {
                                                switch fi108 {
                                                case 1:
                                                    var v122 uint8 = 127
                                                    _ = v122
                                                case -9223372036854775808:
                                                    var v123 int32 = genFn0(genFn0(genFn0((v2 * int32(v120)), 0, (v14 - (v14 * 1000))), (v2 - (v2 + (v2 ^ 42))), (v14 >> ((18446744073709551615) & 63))), v2, (v14 >> ((861) & 63)))
                                                    _ = v123
                                                }
                                            }
                                            var v124 uint64 = 1
                                            _ = fi109
                                            _ = v120
                                            _ = v124
                                        }
                                        _ = fi108
                                    }
                                    v14 -= -32768
                                    if false {
                                        if len(v3) > 0 {
                                            v3[0] = ("foo bar" + "a")
                                        }
                                    } else {
                                        switch 4294967295 {
                                        case 32767:
                                            var v125 uint16 = 42
                                            _ = v125
                                        case -9223372036854775808:
                                            var v126 int8 = 1
                                            _ = v126
                                        }
                                    }
                                }
                                switch ("foo bar" + "foo bar") {
                                case "hello":
                                    for fi127 := 0; fi127 < 13; fi127++ {
                                        var v128 uint8 = 128
                                        fmt.Println(v128, v1)
                                        var v129 int = fi127
                                        _ = fi127
                                        _ = v128
                                        _ = v129
                                    }
                                default:
                                    if true {
                                        _ = len(v3)
                                        v130 := GenS7{F0: 0, F1: ("hello" + ("foo bar" + (("foo bar" + "\n\t") + ("" + "hello"))))}
                                        _ = v130
                                    } else {
                                        switch -9223372036854775808 {
                                        case 857:
                                            var v131 int64 = -1
                                            _ = v131
                                        case 0:
                                            if len(v3) > 0 {
                                                v3[0] = "foo bar"
                                            }
                                        default:
                                            if false {
                                                var v132 int64 = 42
                                                _ = v132
                                            } else {
                                                for fi133 := 0; fi133 < 8; fi133++ {
                                                    if true {
                                                        if false {
                                                            var v134 string = ""
                                                            var v135 int = (fi133 * 32768)
                                                            for fi136 := 0; fi136 < 20; fi136++ {
                                                                _ = len(v4)
                                                                fmt.Println(v14, v1, v135)
                                                                var v137 uint = 2
                                                                _ = fi136
                                                                _ = v137
                                                            }
                                                            _ = v134
                                                            _ = v135
                                                        }
                                                        v1 = (v1 * ((0.1 + (v1 * (v1 + v1))) + v1))
                                                        if false {
                                                            fmt.Println(v1)
                                                            for fi138 := 0; fi138 < 7; fi138++ {
                                                                var v139 float64 = (v15 * 1e10)
                                                                _ = fi138
                                                                _ = v139
                                                            }
                                                        }
                                                    } else {
                                                        var v140 int64 = -1
                                                        v2 = 445
                                                        _ = v140
                                                    }
                                                    for fi141 := 0; fi141 < 5; fi141++ {
                                                        fmt.Println(fi141, fi141)
                                                        _ = fi141
                                                    }
                                                    _ = fi133
                                                }
                                            }
                                        }
                                        v1 -= 1.5
                                    }
                                }
                                v142 := GenS6{F0: 1, F1: 100}
                                _ = v142
                            }
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        if (2 > 255) {
                            switch 32767 {
                            case -9223372036854775808:
                                v15 = v15
                            case 2147483647:
                                for fi143 := 0; fi143 < 18; fi143++ {
                                    fmt.Println(v14, fi143, fi143)
                                    _ = fi143
                                }
                            case 32767:
                                if (65535 == 2) {
                                    switch ((("\n\t" + "") + "") + ("\n\t" + "")) {
                                    case "\n\t":
                                        v2 -= (v2 & (v2 / (((v2 - genFn0(1, v2, v14))) | 1)))
                                    default:
                                        if (v15 > (((v15 * 0.1) + (1.5 + v15)) * (1e10 / (-2.25 + v15)))) {
                                            if false {
                                                if (-128 <= 72) {
                                                    switch "\n\t" {
                                                    case "":
                                                        for fi144 := 0; fi144 < 5; fi144++ {
                                                            var v145 uint8 = 1
                                                            _ = fi144
                                                            _ = v145
                                                        }
                                                    case "a":
                                                        var v146 float64 = (float64(v1) / float64(v1))
                                                        _ = v146
                                                    case "\n\t":
                                                        switch ("a" + ("a" + ("foo bar" + ""))) {
                                                        case "foo bar":
                                                            v147 := GenS3{F0: 127, F1: 435}
                                                            _ = v147
                                                        case "":
                                                            v3 = append(v3, "foo bar")
                                                        case "\n\t":
                                                            for fi148 := 0; fi148 < 12; fi148++ {
                                                                for fi149 := 0; fi149 < 6; fi149++ {
                                                                    switch (("hello" + "foo bar") + ("" + ("\n\t" + "a"))) {
                                                                    case "foo bar":
                                                                        if true {
                                                                            switch (fi148 / (((fi148 + (fi148 ^ 429))) | 1)) {
                                                                            case 42:
                                                                                v15 += 1e10
                                                                            default:
                                                                                fmt.Println(fi149)
                                                                                var v150 int16 = (v14 | 32767)
                                                                                _ = v150
                                                                            }
                                                                        } else {
                                                                            fmt.Println(fi149, v1)
                                                                            v2 = v2
                                                                            v151 := GenS8{F0: false, F1: 208}
                                                                            _ = v151
                                                                        }
                                                                    case "a":
                                                                        var v152 int16 = 42
                                                                        _ = v152
                                                                    case "":
                                                                        var v153 int8 = 39
                                                                        _ = v153
                                                                    default:
                                                                        v2 += 42
                                                                    }
                                                                    _ = len(v3)
                                                                    v1 = (-2.25 - -2.25)
                                                                    if false {
                                                                        fmt.Println(v2, v14)
                                                                    } else {
                                                                        v154 := GenS7{F0: 9223372036854775807, F1: ("hello" + "foo bar")}
                                                                        _ = v154
                                                                    }
                                                                    _ = fi149
                                                                }
                                                                _ = fi148
                                                            }
                                                        default:
                                                            if true {
                                                                fmt.Println(v15, v15)
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    var v155 string = (("a" + ("hello" + "foo bar")) + "a")
                                                    v156 := GenS8{F0: true, F1: 2}
                                                    _ = v155
                                                    _ = v156
                                                }
                                                v3 = append(v3, "\n\t")
                                                for fi157 := 0; fi157 < 14; fi157++ {
                                                    var v158 int16 = (-(v14))
                                                    switch "" {
                                                    case "":
                                                        v159 := GenS2{F0: 42, F1: 0}
                                                        _ = v159
                                                    }
                                                    var v160 uint16 = 2
                                                    _ = fi157
                                                    _ = v158
                                                    _ = v160
                                                }
                                            } else {
                                                switch -89 {
                                                case 728:
                                                    if ((!(153 > 1)) && (707 == 2)) {
                                                        _ = len(v3)
                                                        v3 = append(v3, "hello")
                                                        if (!true) {
                                                            v14 = v14
                                                            if (362 < 1) {
                                                                if (!true) {
                                                                    if ((true && (0 >= 0)) || ((1 >= 2) && false)) {
                                                                        fmt.Println(v14, v2, v2)
                                                                        switch 32767 {
                                                                        case -2147483648:
                                                                            var v161 bool = ((!true) || true)
                                                                            _ = v161
                                                                        case -254:
                                                                            var v162 int16 = v14
                                                                            _ = v162
                                                                        case 166:
                                                                            v1 += (1e10 + 3.14159)
                                                                        default:
                                                                            var v163 int32 = (^(v2))
                                                                            _ = v163
                                                                        }
                                                                        v164 := v4["hello"]
                                                                        _ = v164
                                                                    }
                                                                    for fi165 := 0; fi165 < 14; fi165++ {
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        var v166 int = (fi165 % ((fi165) | 1))
                                                                        v167 := GenS6{F0: -82, F1: 127}
                                                                        _ = fi165
                                                                        _ = v166
                                                                        _ = v167
                                                                    }
                                                                    for fi168 := 0; fi168 < 6; fi168++ {
                                                                        v169 := GenS6{F0: -1, F1: 1}
                                                                        if ((18446744073709551615 > 65535) || true) {
                                                                            v170 := GenS6{F0: 1, F1: -1}
                                                                            _ = v170
                                                                        }
                                                                        _ = fi168
                                                                        _ = v169
                                                                    }
                                                                } else {
                                                                    var v171 string = ""
                                                                    _ = v171
                                                                }
                                                                var v172 int64 = 500
                                                                _ = v172
                                                            }
                                                        } else {
                                                            v1 += 1e10
                                                        }
                                                    } else {
                                                        if false {
                                                            var v173 uint64 = 42
                                                            var v174 int16 = (v14 | v14)
                                                            _ = v173
                                                            _ = v174
                                                        } else {
                                                            var v175 uint64 = 515
                                                            for fi176 := 0; fi176 < 14; fi176++ {
                                                                v177 := GenS6{F0: 90, F1: 61}
                                                                v178 := GenS7{F0: -9223372036854775808, F1: "foo bar"}
                                                                _ = fi176
                                                                _ = v177
                                                                _ = v178
                                                            }
                                                            _ = v175
                                                        }
                                                        var v179 float32 = v1
                                                        v180 := GenS1{F0: 92}
                                                        _ = v179
                                                        _ = v180
                                                    }
                                                case 32768:
                                                    var v181 string = "foo bar"
                                                    _ = v181
                                                case 2147483647:
                                                    v182 := GenS9{F0: 1, F1: 0}
                                                    _ = v182
                                                default:
                                                    var v183 uint16 = 42
                                                    _ = v183
                                                }
                                                if false {
                                                    if false {
                                                        if (!false) {
                                                            if true {
                                                                var v184 int = 128
                                                                var v185 uint8 = 127
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    if true {
                                                                        switch 4294967295 {
                                                                        case -578:
                                                                            var v186 uint8 = (^(v185))
                                                                            _ = v186
                                                                        case 127:
                                                                            fmt.Println(v15)
                                                                        }
                                                                    } else {
                                                                        v15 = 0.1
                                                                        for fi187 := 0; fi187 < 11; fi187++ {
                                                                            v1 = (v1 - 1e10)
                                                                            _ = fi187
                                                                        }
                                                                        switch "a" {
                                                                        case "a":
                                                                            var v188 int8 = 77
                                                                            _ = v188
                                                                        default:
                                                                            v189 := v4[("a" + "\n\t")]
                                                                            _ = v189
                                                                        }
                                                                    }
                                                                case "a":
                                                                    var v190 uint8 = (v185 | 0)
                                                                    _ = v190
                                                                case "":
                                                                    switch ("foo bar" + (("hello" + ("\n\t" + "a")) + "hello")) {
                                                                    case "hello":
                                                                        for fi191 := 0; fi191 < 9; fi191++ {
                                                                            if (!(-9223372036854775808 != -1)) {
                                                                                v2 = -251
                                                                            } else {
                                                                                fmt.Println(v185, v184)
                                                                                fmt.Println(v184, v2, v184)
                                                                            }
                                                                            _ = fi191
                                                                        }
                                                                    case "foo bar":
                                                                        fmt.Println(v14, v14)
                                                                    case "":
                                                                        var v192 int64 = -732
                                                                        _ = v192
                                                                    }
                                                                default:
                                                                    if (!false) {
                                                                        for fi193 := 0; fi193 < 8; fi193++ {
                                                                            var v194 uint16 = 65535
                                                                            _ = fi193
                                                                            _ = v194
                                                                        }
                                                                        switch (("a" + (("foo bar" + "") + "foo bar")) + "\n\t") {
                                                                        case "\n\t":
                                                                            v14 = -1
                                                                        default:
                                                                            if false {
                                                                                for fi195 := 0; fi195 < 15; fi195++ {
                                                                                    v196 := GenS10{F0: ("foo bar" + "\n\t"), F1: v184, F2: (v185 | (^(v185)))}
                                                                                    v1 -= (float32(((((v15 + v15) + float64(float32(((((v15 * v15) + -2.25) * (v15 + v15)) * float64(v1))))) - v15) - (0.1 / float64(((((1e10 + ((float32(-2.25) + (v1 / ((0.0 / 1e10) * v1))) * (v1 + 1e10))) - v1) / v1) / ((v1 + (v1 * ((0.1 * 1.5) + v1))) / (float32(1.5) - 1e10))))))) + v1)
                                                                                    for fi197 := 0; fi197 < 13; fi197++ {
                                                                                        var v198 uint = 1
                                                                                        v184 = (^(v184))
                                                                                        _ = fi197
                                                                                        _ = v198
                                                                                    }
                                                                                    _ = fi195
                                                                                    _ = v196
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        v199 := GenS8{F0: (2 < 0), F1: 42}
                                                                        for fi200 := 0; fi200 < 14; fi200++ {
                                                                            for fi201 := 0; fi201 < 18; fi201++ {
                                                                                v202 := GenS9{F0: -1, F1: 0}
                                                                                var v203 float32 = v1
                                                                                _ = fi201
                                                                                _ = v202
                                                                                _ = v203
                                                                            }
                                                                            _ = fi200
                                                                        }
                                                                        fmt.Println(v15, v1, v2)
                                                                        _ = v199
                                                                    }
                                                                }
                                                                _ = v184
                                                                _ = v185
                                                            }
                                                        }
                                                        if false {
                                                            var v204 bool = (true || false)
                                                            _ = v204
                                                        }
                                                    } else {
                                                        v14 = v14
                                                        for fi205 := 0; fi205 < 18; fi205++ {
                                                            v1 = v1
                                                            if len(v3) > 0 {
                                                                v3[0] = "hello"
                                                            }
                                                            _ = fi205
                                                        }
                                                    }
                                                    for fi206 := 0; fi206 < 4; fi206++ {
                                                        v1 = v1
                                                        if (v15 <= v15) {
                                                            if true {
                                                                v14 = (v14 >> ((1) & 63))
                                                                if len(v3) > 0 {
                                                                    v3[0] = ("\n\t" + "")
                                                                }
                                                            } else {
                                                                for fi207 := 0; fi207 < 19; fi207++ {
                                                                    v2 -= (v2 | (v2 + v2))
                                                                    var v208 int16 = 1000
                                                                    fmt.Println(v208, v14)
                                                                    _ = fi207
                                                                    _ = v208
                                                                }
                                                                var v209 int16 = -414
                                                                _ = v209
                                                            }
                                                        } else {
                                                            if (-61 >= 0) {
                                                                delete(v4, "")
                                                            }
                                                            v2 = -1
                                                        }
                                                        _ = fi206
                                                    }
                                                    v210 := GenS2{F0: 42, F1: -128}
                                                    _ = v210
                                                } else {
                                                    v14 = 651
                                                    v211 := GenS11{F0: 42, F1: 32767}
                                                    switch -9223372036854775808 {
                                                    case 65:
                                                        v14 -= v14
                                                    case -469:
                                                        var v212 uint = 42
                                                        _ = v212
                                                    case 32768:
                                                        v4["foo bar"] = 2
                                                    }
                                                    _ = v211
                                                }
                                            }
                                        } else {
                                            var v213 uint64 = 904
                                            var v214 int32 = v2
                                            _ = v213
                                            _ = v214
                                        }
                                    }
                                    if (0 >= 9223372036854775807) {
                                        switch 436 {
                                        case 32768:
                                            if (((!false) || ("\n\t" < ("a" + "\n\t"))) && (0 >= 517)) {
                                                if true {
                                                    v4[""] = 2
                                                } else {
                                                    v14 += v14
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        v1 -= -2.25
                                                    }
                                                    if false {
                                                        for fi215 := 0; fi215 < 13; fi215++ {
                                                            for fi216 := 0; fi216 < 7; fi216++ {
                                                                switch ("a" + "hello") {
                                                                case "":
                                                                    var v217 uint64 = 0
                                                                    var v218 float64 = v15
                                                                    _ = v217
                                                                    _ = v218
                                                                case "foo bar":
                                                                    v4["foo bar"] = 625
                                                                case "\n\t":
                                                                    var v219 int = fi216
                                                                    _ = v219
                                                                }
                                                                _ = len(v4)
                                                                if ((false || (false || (!(!true)))) || (!(v15 > v15))) {
                                                                    if false {
                                                                        var v220 bool = (false || true)
                                                                        v2 -= v2
                                                                        var v221 int32 = 0
                                                                        switch fi215 {
                                                                        case 127:
                                                                            fmt.Println(v1)
                                                                        case 629:
                                                                            v222 := GenS9{F0: 42, F1: -9223372036854775808}
                                                                            _ = v222
                                                                        case -9223372036854775808:
                                                                            v223 := GenS12{F0: 65535, F1: 2}
                                                                            _ = v223
                                                                        default:
                                                                            _ = len(v3)
                                                                        }
                                                                        switch "hello" {
                                                                        case "a":
                                                                            var v224 uint64 = 4294967295
                                                                            _ = v224
                                                                        case "":
                                                                            v225 := GenS13{F0: -128, F1: 9223372036854775807}
                                                                            _ = v225
                                                                        default:
                                                                            switch fi216 {
                                                                            case 0:
                                                                                fmt.Println(fi216)
                                                                            default:
                                                                                if (162 >= fi215) {
                                                                                    var v226 uint64 = 143
                                                                                    var v227 bool = v220
                                                                                    switch 595 {
                                                                                    case 128:
                                                                                        var v228 int32 = (v2 & (v2 / (((v2 & v2)) | 1)))
                                                                                        _ = v228
                                                                                    default:
                                                                                        _ = len(v3)
                                                                                        v229 := GenS14{F0: (v2 * 42)}
                                                                                        _ = v229
                                                                                    }
                                                                                    _ = v226
                                                                                    _ = v227
                                                                                } else {
                                                                                    v230 := GenS2{F0: 0, F1: 0}
                                                                                    var v231 string = "hello"
                                                                                    _ = v230
                                                                                    _ = v231
                                                                                }
                                                                            }
                                                                        }
                                                                        v1 -= (v1 * 0.0)
                                                                        _ = v220
                                                                        _ = v221
                                                                    }
                                                                    var v232 int64 = -1
                                                                    _ = v232
                                                                } else {
                                                                    v233 := GenS15{F0: v14}
                                                                    v234 := v4["a"]
                                                                    _ = v233
                                                                    _ = v234
                                                                }
                                                                switch 128 {
                                                                case 425:
                                                                    v235 := GenS16{F0: 65535, F1: fi215, F2: 13}
                                                                    _ = v235
                                                                case 127:
                                                                    v14 = (v14 * v14)
                                                                case 2147483647:
                                                                    if len(v3) > 0 {
                                                                        v3[0] = "a"
                                                                    }
                                                                default:
                                                                    _ = len(v4)
                                                                }
                                                                _ = fi216
                                                            }
                                                            v2 = 0
                                                            _ = fi215
                                                        }
                                                        if true {
                                                            if len(v3) > 0 {
                                                                v3[0] = (((("foo bar" + "hello") + "foo bar") + "\n\t") + "")
                                                            }
                                                            switch 128 {
                                                            case -137:
                                                                delete(v4, "a")
                                                            }
                                                            for fi236 := 0; fi236 < 19; fi236++ {
                                                                v4[("a" + "hello")] = 161
                                                                v14 -= (v14 + 1000)
                                                                _ = fi236
                                                            }
                                                        }
                                                        for fi237 := 0; fi237 < 12; fi237++ {
                                                            v238 := v4["\n\t"]
                                                            _ = fi237
                                                            _ = v238
                                                        }
                                                    } else {
                                                        fmt.Println(v2, v14, v15)
                                                        delete(v4, "")
                                                    }
                                                }
                                                v239 := GenS17{F0: 4294967295}
                                                v1 += float32(v15)
                                                _ = v239
                                            }
                                        case -9223372036854775808:
                                            var v240 uint = 1
                                            _ = v240
                                        case 0:
                                            v241 := GenS1{F0: 42}
                                            _ = v241
                                        }
                                    }
                                    if (!false) {
                                        for fi242 := 0; fi242 < 7; fi242++ {
                                            switch "foo bar" {
                                            case "\n\t":
                                                if (2147483648 < 0) {
                                                    var v243 int8 = -52
                                                    v244 := GenS18{F0: v243}
                                                    v1 += ((((v1 - ((-2.25 * 3.14159) + v1)) / ((0.0 + (float32(v15) / v1)) + float32(1.5))) * float32(v15)) * v1)
                                                    _ = v243
                                                    _ = v244
                                                } else {
                                                    v1 = 1.5
                                                    var v245 uint32 = 4294967295
                                                    v4["foo bar"] = (v245 / (((^(v245))) | 1))
                                                    _ = v245
                                                }
                                            default:
                                                v246 := GenS5{F0: (!((fi242 / (((fi242 & int(v2))) | 1)) < fi242)), F1: (v14 & v14), F2: (v2 - genFn0(772, (v2 - (v2 - (v2 >> ((255) & 63)))), (v14 % ((42) | 1))))}
                                                _ = v246
                                            }
                                            var v247 uint64 = 42
                                            v3 = append(v3, "\n\t")
                                            _ = fi242
                                            _ = v247
                                        }
                                        for fi248 := 0; fi248 < 20; fi248++ {
                                            for fi249 := 0; fi249 < 12; fi249++ {
                                                v2 = (^(v2))
                                                _ = fi249
                                            }
                                            switch ("" + "foo bar") {
                                            case "hello":
                                                v2 -= v2
                                            case "a":
                                                for fi250 := 0; fi250 < 18; fi250++ {
                                                    var v251 int = 162
                                                    if true {
                                                        var v252 int32 = (^(v2))
                                                        fmt.Println(fi250, v15)
                                                        _ = v252
                                                    }
                                                    _ = fi250
                                                    _ = v251
                                                }
                                            }
                                            for fi253 := 0; fi253 < 6; fi253++ {
                                                switch "foo bar" {
                                                case "hello":
                                                    v254 := GenS8{F0: false, F1: 1}
                                                    _ = v254
                                                case "\n\t":
                                                    if true {
                                                        v1 += v1
                                                    }
                                                }
                                                var v255 string = ""
                                                _ = fi253
                                                _ = v255
                                            }
                                            _ = fi248
                                        }
                                    }
                                } else {
                                    for fi256 := 0; fi256 < 15; fi256++ {
                                        fmt.Println(fi256)
                                        for fi257 := 0; fi257 < 19; fi257++ {
                                            var v258 string = ""
                                            fmt.Println(v2)
                                            switch fi257 {
                                            case -203:
                                                v259 := GenS19{F0: 441}
                                                _ = v259
                                            case -216:
                                                fmt.Println(v1, fi257, v14)
                                            default:
                                                _ = len(v3)
                                            }
                                            var v260 uint32 = 2
                                            if (v258 > v258) {
                                                v14 = 42
                                            }
                                            _ = fi257
                                            _ = v258
                                            _ = v260
                                        }
                                        fmt.Println(v15)
                                        _ = fi256
                                    }
                                }
                            }
                        }
                        for fi261 := 0; fi261 < 8; fi261++ {
                            switch ("foo bar" + "hello") {
                            case "\n\t":
                                if (int16(v2) != (v14 - (v14 >> ((438) & 63)))) {
                                    var v262 int16 = v14
                                    v263 := GenS20{F0: 42, F1: 211, F2: int64(fi261)}
                                    _ = v262
                                    _ = v263
                                }
                            case "hello":
                                v264 := GenS9{F0: 9223372036854775807, F1: -1}
                                _ = v264
                            case "foo bar":
                                v265 := GenS11{F0: 4294967295, F1: 2}
                                _ = v265
                            default:
                                v15 += 3.14159
                            }
                            var v266 uint = 1
                            if (271 == 18446744073709551615) {
                                v15 -= ((float64(float32((v15 / v15))) * ((v15 * float64((0.0 * 3.14159))) + (v15 - v15))) + ((0.1 / v15) + (v15 * (v15 / (v15 / ((v15 * v15) - 1.5))))))
                            }
                            _ = fi261
                            _ = v266
                        }
                    }
                }
                if (652 <= 65535) {
                    if (false || (true && true)) {
                        var v267 uint16 = 65535
                        v267 = (v267 + v267)
                        _ = v267
                    }
                } else {
                    var v268 float32 = 0.0
                    _ = v268
                }
            }
            _ = v17
        } else {
            v14 = 75
            if false {
                var v269 float32 = v1
                _ = v269
            } else {
                v2 = (v2 / (((v2 * (^(v2)))) | 1))
            }
        }
    default:
        v14 -= (v14 >> ((2) & 63))
    }
    fmt.Println(v14, v14)
    if (((v2 % (((v2 ^ (v2 / (((-(v2))) | 1)))) | 1)) >= 2147483647) || ((v1 - 1.5) < v1)) {
        var v270 uint16 = 255
        _ = v270
    } else {
        var v271 float32 = (v1 * (-2.25 - (v1 / v1)))
        var v272 int64 = 87
        _ = v271
        _ = v272
    }
    v273 := GenS21{F0: 0}
    v274 := GenS17{F0: 0}
    var v275 int64 = -74
    fmt.Println(v2, v275)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v275:", v275)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v273
    _ = v274
    _ = v275
    _ = v3
    _ = v4
}
